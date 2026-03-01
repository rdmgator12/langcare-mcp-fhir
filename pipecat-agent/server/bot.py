import json
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
from pipecat.transcriptions.language import Language
from pipecat.services.mcp_service import MCPClient
from pipecat.runner.types import (
    DailyDialinRequest,
    DailyRunnerArguments,
    RunnerArguments,
    SmallWebRTCRunnerArguments,
)
from pipecat.transports.base_transport import TransportParams
from pipecat.transports.daily.transport import DailyDialinSettings, DailyParams, DailyTransport

try:
    from pipecat.transports.smallwebrtc.transport import SmallWebRTCTransport
except Exception:
    SmallWebRTCTransport = None
from mcp.client.session import ClientSession
from mcp.client.session_group import StreamableHttpParameters
from mcp.client.streamable_http import streamablehttp_client

from prompts.base_prompt import build_clinical_prompt


async def lookup_patient_by_phone(phone: str, server_params: StreamableHttpParameters) -> dict | None:
    """Search FHIR for a patient matching this phone number via MCP."""
    # URL-encode the + sign for FHIR search (HAPI requires %2B)
    from urllib.parse import quote
    phone_encoded = quote(phone, safe="")

    try:
        async with streamablehttp_client(**server_params.model_dump()) as (read, write, _):
            async with ClientSession(read, write) as session:
                await session.initialize()
                result = await session.call_tool(
                    "fhir_search",
                    {"resourceType": "Patient", "queryParams": f"phone={phone_encoded}"},
                )

        if not result or not result.content:
            logger.info(f"No MCP content for phone lookup: {phone_digits}")
            return None

        bundle_text = result.content[0].text
        bundle = json.loads(bundle_text)

        entries = bundle.get("entry", [])
        if not entries:
            logger.info(f"No patient found for phone: {phone_digits}")
            return None

        patient = entries[0].get("resource", {})
        names = patient.get("name", [{}])
        given = names[0].get("given", [""])
        family = names[0].get("family", "")
        patient_name = f"{given[0]} {family}".strip()

        return {
            "patient_fhir_id": patient["id"],
            "patient_name": patient_name,
            "patient_birthdate": patient.get("birthDate", ""),
        }
    except Exception as e:
        logger.error(f"Patient phone lookup failed: {e}")
        return None


async def bot(args: RunnerArguments):
    """Main bot entry point. Called by PipeCat Cloud or local runner."""
    use_daily = isinstance(args, DailyRunnerArguments)
    use_webrtc = isinstance(args, SmallWebRTCRunnerArguments)

    if not use_daily and not use_webrtc:
        logger.error(f"Unsupported runner args: {type(args).__name__}. "
                     "Use -t daily (needs DAILY_API_KEY) or -t webrtc.")
        return

    body = args.body or {}

    # Default test patient for local WebRTC testing when no body is provided
    if use_webrtc and not body:
        # Scenario 1 — Pre-authenticated (default):**
        # body = {
        #     "patient_fhir_id": "d886a934-5568-42b3-9324-0f0b05fc018c",
        #     "patient_name": "Patricia Ann Williams",
        #     "verified": True,
        # }
        #Scenario 2 — DOB verification:**
        #The agent will ask you to confirm the date of birth. Say the correct DOB to unlock clinical access.
        body = {
            "caller_phone": "+15039264538",
            "patient_context": {
                "patient_fhir_id": "d886a934-5568-42b3-9324-0f0b05fc018c",
                "patient_name": "Patricia Ann Williams",
                "patient_birthdate": "1958-11-08",
            },
        }
        #Scenario 3 — Unknown caller:**
        # body = {}

        logger.info("No body provided — using default test patient")

    if use_daily:
        logger.info(f"Bot started (Daily): room={args.room_url}")
    else:
        logger.info("Bot started (SmallWebRTC)")
    logger.info(f"Session data: {body}")

    # ── MCP Client (needed for phone lookup and tool registration) ──

    mcp_url = os.getenv("LANGCARE_MCP_URL", "https://langcare-mcp-dev.fly.dev/mcp")
    mcp_headers = {}
    mcp_api_key = os.getenv("LANGCARE_API_KEY")
    if mcp_api_key:
        mcp_headers["Authorization"] = f"Bearer {mcp_api_key}"

    mcp_server_params = StreamableHttpParameters(
        url=mcp_url,
        headers=mcp_headers if mcp_headers else None,
    )
    mcp = MCPClient(server_params=mcp_server_params)

    # ── Detect dial-in (PSTN) vs WebRTC ──

    dialin_request = None
    if use_daily:
        try:
            dialin_request = DailyDialinRequest.model_validate(body)
            logger.info(f"Dial-in call detected from: {dialin_request.dialin_settings.From}")
        except Exception:
            pass

    if dialin_request:
        # PSTN dial-in: look up patient by caller phone number
        caller_phone = dialin_request.dialin_settings.From or "unknown"
        logger.info(f"Looking up patient for phone: {caller_phone}")
        patient_context = await lookup_patient_by_phone(caller_phone, mcp_server_params)
        verified = False
        patient_fhir_id = None
        patient_name = None
    else:
        # WebRTC or Daily (non-dialin): extract session context from body
        patient_context = body.get("patient_context")
        caller_phone = body.get("caller_phone", "unknown")
        verified = body.get("verified", False)
        patient_fhir_id = body.get("patient_fhir_id")
        patient_name = body.get("patient_name")

        # If caller_phone provided but no patient_context, look up in FHIR
        if not patient_context and not verified and caller_phone != "unknown":
            logger.info(f"Looking up patient for phone: {caller_phone}")
            patient_context = await lookup_patient_by_phone(caller_phone, mcp_server_params)

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
            f"You are a healthcare voice assistant on a live phone call "
            f"with a caller from {caller_phone}. No patient account matches "
            f"this phone number.\n\n"
            f"Speak directly to the caller. Tell them you were not able to "
            f"find an account associated with their phone number and suggest "
            f"they call the main office to set up their account or resolve "
            f"the issue. Be warm and brief.\n\n"
            f"Do NOT narrate your actions or describe what you are doing. "
            f"Do NOT access or search any patient records. "
            f"Do NOT repeat yourself."
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

    stt = DeepgramSTTService(
        api_key=os.getenv("DEEPGRAM_API_KEY"),
        language=os.getenv("VOICE_LANGUAGE", "en"),
    )

    tts = CartesiaTTSService(
        api_key=os.getenv("CARTESIA_API_KEY"),
        voice_id=os.getenv("CARTESIA_VOICE_ID", "71a7ad14-091c-4e8e-a314-022ece01c121"),
        model=os.getenv("CARTESIA_MODEL", "sonic-2"),
        params=CartesiaTTSService.InputParams(
            language=Language(os.getenv("VOICE_LANGUAGE", "en")),
        ),
    )

    if use_daily:
        daily_params = DailyParams(
            audio_in_enabled=True,
            audio_out_enabled=True,
        )

        if dialin_request:
            daily_params.api_key = dialin_request.daily_api_key
            daily_params.api_url = dialin_request.daily_api_url
            daily_params.dialin_settings = DailyDialinSettings(
                call_id=dialin_request.dialin_settings.call_id,
                call_domain=dialin_request.dialin_settings.call_domain,
            )

        transport = DailyTransport(
            args.room_url,
            args.token,
            "Healthcare Assistant",
            daily_params,
        )
    else:
        transport = SmallWebRTCTransport(
            webrtc_connection=args.webrtc_connection,
            params=TransportParams(
                audio_in_enabled=True,
                audio_out_enabled=True,
            ),
        )

    # ── MCP Tools (auto-discovered from LangCare) ──

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

    if use_daily:
        @transport.event_handler("on_first_participant_joined")
        async def on_first_participant_joined(transport, participant):
            await task.queue_frames([LLMRunFrame()])

        @transport.event_handler("on_participant_left")
        async def on_participant_left(transport, participant, reason):
            await task.cancel()
    else:
        @transport.event_handler("on_client_connected")
        async def on_client_connected(transport, client):
            await task.queue_frames([LLMRunFrame()])

        @transport.event_handler("on_client_disconnected")
        async def on_client_disconnected(transport, client):
            await task.cancel()

    runner = PipelineRunner()
    await runner.run(task)


if __name__ == "__main__":
    from pipecat.runner.run import main

    main()
