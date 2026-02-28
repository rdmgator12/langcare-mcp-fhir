import os

from dotenv import load_dotenv
from loguru import logger

load_dotenv(override=True)

from pipecat.adapters.schemas.function_schema import FunctionSchema
from pipecat.adapters.schemas.tools_schema import ToolsSchema
from pipecat.audio.vad.silero import SileroVADAnalyzer
from pipecat.audio.vad.vad_analyzer import VADParams
from pipecat.frames.frames import LLMRunFrame
from pipecat.pipeline.pipeline import Pipeline
from pipecat.pipeline.runner import PipelineRunner
from pipecat.pipeline.task import PipelineTask
from pipecat.processors.aggregators.llm_context import LLMContext
from pipecat.processors.aggregators.llm_response_universal import (
    LLMContextAggregatorPair,
    LLMUserAggregatorParams,
)
from pipecat.services.anthropic.llm import AnthropicLLMService
from pipecat.services.cartesia.tts import CartesiaTTSService
from pipecat.services.deepgram.stt import DeepgramSTTService
from pipecat.services.llm_service import FunctionCallParams
from pipecat.services.mcp_service import MCPClient
from pipecat.runner.types import DailyRunnerArguments, RunnerArguments
from pipecat.transports.daily.transport import DailyParams, DailyTransport
from mcp.client.session_group import StreamableHttpParameters

from prompts.base_prompt import build_clinical_prompt


async def bot(args: RunnerArguments):
    """Main bot entry point. Called by PipeCat Cloud or local runner."""
    if not isinstance(args, DailyRunnerArguments):
        logger.error(f"Expected DailyRunnerArguments, got {type(args).__name__}. "
                     "Ensure DAILY_API_KEY is set.")
        return

    room_url = args.room_url
    token = args.token
    body = args.body or {}

    logger.info(f"Bot started: room={room_url}")
    logger.info(f"Session data: {body}")

    # Extract session context
    patient_context = body.get("patient_context")
    caller_phone = body.get("caller_phone", "unknown")
    verified = body.get("verified", False)
    patient_fhir_id = body.get("patient_fhir_id")
    patient_name = body.get("patient_name")

    # ── Determine initial prompt based on auth state ──

    if verified and patient_fhir_id:
        # Pre-authenticated (web app, mobile, clinician portal)
        system_prompt = build_clinical_prompt(patient_fhir_id, patient_name)
        needs_verification = False

    elif patient_context:
        # Inbound phone — known caller, needs DOB verification
        patient_name = patient_context["patient_name"]
        patient_birthdate = patient_context["patient_birthdate"]
        patient_fhir_id = patient_context["patient_fhir_id"]
        needs_verification = True

        system_prompt = (
            f"You are a healthcare voice assistant. A caller is on the line "
            f"from {caller_phone}.\n\n"
            f"Based on their phone number, this caller appears to be "
            f"{patient_name} (Patient/{patient_fhir_id}).\n\n"
            f"VERIFICATION STEP (required before any clinical access):\n"
            f"1. Greet the caller warmly.\n"
            f'2. Say: "I see you\'re calling from a number we have on file. '
            f'For security, can you please confirm your date of birth?"\n'
            f"3. The caller should provide a date. Compare it against: "
            f"{patient_birthdate}\n"
            f'   - Accept common formats: "March 15 1980", "3/15/80", '
            f'"03-15-1980", etc.\n'
            f"   - If it matches, call the verify_patient function.\n"
            f'   - If it doesn\'t match, say "I\'m sorry, that doesn\'t match '
            f'our records. Could you try again?" Allow up to 3 attempts.\n'
            f'   - After 3 failed attempts, say "I\'m unable to verify your '
            f'identity. Please contact our office directly." and end the call.\n\n'
            f"IMPORTANT:\n"
            f"- Do NOT access any clinical data until verify_patient has been called.\n"
            f"- Do NOT tell the caller what birthdate you have on file.\n"
            f"- Do NOT mention the patient's name until they confirm their DOB.\n"
        )

    else:
        # Unknown caller — no phone match found
        needs_verification = False
        system_prompt = (
            f"You are a healthcare voice assistant. A caller is on the line "
            f"from {caller_phone}, but we could not find a patient matching "
            f"this phone number.\n\n"
            f'Politely explain: "I wasn\'t able to find an account associated '
            f"with your phone number. You may need to call our main office to "
            f"set up your account, or if you believe this is an error, please "
            f'contact us at our main office number."\n\n'
            f"Then end the call. Do NOT attempt to look up or access any "
            f"patient records."
        )

    # ── Services ──

    llm = AnthropicLLMService(
        api_key=os.getenv("ANTHROPIC_API_KEY"),
        model="claude-sonnet-4-5-20250929",
        params=AnthropicLLMService.InputParams(
            enable_prompt_caching=True,
            temperature=0.3,
        ),
    )

    stt = DeepgramSTTService(api_key=os.getenv("DEEPGRAM_API_KEY"))

    tts = CartesiaTTSService(
        api_key=os.getenv("CARTESIA_API_KEY"),
        voice_id="71a7ad14-091c-4e8e-a314-022ece01c121",
    )

    transport = DailyTransport(
        room_url,
        token,
        "Healthcare Assistant",
        DailyParams(
            audio_in_enabled=True,
            audio_out_enabled=True,
        ),
    )

    # ── MCP Tools (auto-discovered from LangCare) ──

    mcp_url = os.getenv("LANGCARE_MCP_URL", "https://langcare-mcp-dev.fly.dev/mcp")
    mcp_headers = {}
    mcp_api_key = os.getenv("LANGCARE_API_KEY")
    if mcp_api_key:
        mcp_headers["Authorization"] = f"Bearer {mcp_api_key}"

    mcp = MCPClient(
        server_params=StreamableHttpParameters(
            url=mcp_url,
            headers=mcp_headers if mcp_headers else None,
        )
    )
    mcp_tools = await mcp.register_tools(llm)

    # ── Verification function (only for inbound phone) ──

    tools_list = list(mcp_tools.standard_tools)

    if needs_verification:
        verify_patient_tool = FunctionSchema(
            name="verify_patient",
            description=(
                "Call this ONLY when the caller has confirmed their date of birth "
                "and it matches the expected value. This unlocks clinical access "
                "to their medical records."
            ),
            properties={},
            required=[],
        )
        tools_list.append(verify_patient_tool)

        async def verify_patient(params: FunctionCallParams):
            """Transition from verification to clinical mode."""
            clinical_prompt = build_clinical_prompt(patient_fhir_id, patient_name)
            # Replace context messages in-place for clean transition
            context.messages[:] = [
                {"role": "system", "content": clinical_prompt},
            ]
            await params.result_callback(
                f"Identity verified. You are now speaking with {patient_name}. "
                f"Greet them by first name and ask how you can help today."
            )

        llm.register_function(
            "verify_patient",
            verify_patient,
            cancel_on_interruption=False,
        )

    # ── Context + Pipeline ──

    all_tools = ToolsSchema(standard_tools=tools_list)
    messages = [{"role": "system", "content": system_prompt}]
    context = LLMContext(messages=messages, tools=all_tools)

    user_aggregator, assistant_aggregator = LLMContextAggregatorPair(
        context,
        user_params=LLMUserAggregatorParams(
            vad_analyzer=SileroVADAnalyzer(params=VADParams(stop_secs=0.2)),
        ),
    )

    pipeline = Pipeline(
        [
            transport.input(),
            stt,
            user_aggregator,
            llm,
            tts,
            transport.output(),
            assistant_aggregator,
        ]
    )

    task = PipelineTask(pipeline)

    @transport.event_handler("on_first_participant_joined")
    async def on_first_participant_joined(transport, participant):
        await task.queue_frames([LLMRunFrame()])

    @transport.event_handler("on_participant_left")
    async def on_participant_left(transport, participant, reason):
        await task.cancel()

    runner = PipelineRunner()
    await runner.run(task)


if __name__ == "__main__":
    from pipecat.runner.run import main

    main()
