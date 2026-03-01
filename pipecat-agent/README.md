# LangCare Voice Agent

Real-time voice AI agent that connects patients to their electronic health records over WebRTC or phone. Built on [PipeCat](https://docs.pipecat.ai/), powered by Claude, with FHIR R4 data access via [LangCare MCP](https://github.com/langcare/langcare-mcp-fhir).

## Architecture

```
Audio In → DeepGram STT → Claude (Anthropic) → Cartesia TTS → Audio Out
                               ↓
                          MCP Tool Calls
                               ↓
                    LangCare MCP Server (Fly.io)
                               ↓
                      EMR (Epic / Cerner / HAPI FHIR)
```

Five layers, three auth boundaries:

| Layer | Component | Role |
|-------|-----------|------|
| 1 | Client | Browser (WebRTC), phone (PSTN), or mobile app |
| 2 | PipeCat | Voice pipeline — STT, LLM orchestration, TTS |
| 3 | Claude | Reasoning, tool calling, clinical Q&A |
| 4 | LangCare MCP | FHIR R4 operations over Streamable HTTP |
| 5 | EMR | Epic, Cerner, GCP Healthcare API, or HAPI FHIR |

## How It Works

1. Patient connects via browser or phone
2. Speech is transcribed by DeepGram
3. Claude processes the request, calls FHIR tools when needed (lab results, medications, conditions, etc.)
4. FHIR data is fetched from the EMR via LangCare MCP server
5. Claude synthesizes a patient-friendly response
6. Cartesia converts the response to natural speech

## Authentication Flows

The agent supports three caller states:

| State | How It Works | Example |
|-------|-------------|---------|
| **Pre-authenticated** | Backend passes `verified: true` + patient FHIR ID. Agent skips to clinical mode. | Web app with SSO, mobile with biometrics |
| **Known caller, needs verification** | Phone number matches a patient. Agent asks for DOB, verifies, then unlocks clinical access via `verify_patient()`. | Inbound phone call |
| **Unknown caller** | No patient match. Agent politely declines access. | Unrecognized phone number |

## Project Structure

```
pipecat-agent/
├── server/
│   ├── bot.py                 # PipeCat agent entry point
│   ├── prompts/
│   │   └── base_prompt.py     # System prompts (clinical mode)
│   ├── pyproject.toml         # Python dependencies
│   ├── .env.example           # Required environment variables
│   ├── Dockerfile             # PipeCat Cloud deployment
│   └── pcc-deploy.toml        # PipeCat Cloud config
├── webhook/
│   └── server.py              # FastAPI webhook for inbound phone calls
├── CLAUDE.md                  # Full design spec for Claude Code
└── README.md                  # This file
```

## Prerequisites

- Python 3.11 - 3.13
- [uv](https://docs.astral.sh/uv/) (package manager)
- API keys for: Anthropic, DeepGram, Cartesia
- Daily API key (only for Daily transport mode)

## Setup

```bash
cd pipecat-agent/server

# Create venv and install dependencies
uv venv
source .venv/bin/activate
uv sync

# Copy and fill in your API keys
cp .env.example .env
```

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `ANTHROPIC_API_KEY` | Yes | Claude LLM |
| `DEEPGRAM_API_KEY` | Yes | Speech-to-text |
| `CARTESIA_API_KEY` | Yes | Text-to-speech |
| `DAILY_API_KEY` | Daily mode only | Daily.co transport |
| `LANGCARE_MCP_URL` | No | MCP server URL (defaults to `https://langcare-mcp-dev.fly.dev/mcp`) |
| `LANGCARE_API_KEY` | No | Bearer token for MCP server auth |

## Local Testing

The agent supports two transport modes and three authentication scenarios for local development.

### Test Scenarios

The agent has three authentication paths. Each can be tested with either transport mode.

| Scenario | What Happens | Payload Key |
|----------|-------------|-------------|
| **Pre-authenticated** | Agent greets patient by name, immediately has clinical access to FHIR tools | `verified: true` + `patient_fhir_id` |
| **DOB verification** | Agent asks caller to confirm date of birth before granting clinical access | `patient_context` with `patient_birthdate` |
| **Unknown caller** | Agent politely declines — no FHIR access | No `patient_context`, no `verified` |

---

### Option A: WebRTC (no Daily account needed)

Runs a local WebRTC server with a browser-based UI. No Daily API key required.

```bash
# Install WebRTC dependencies (one-time)
uv pip install 'pipecat-ai[webrtc]' pipecat-ai-small-webrtc-prebuilt

# Start the agent
source .venv/bin/activate
python bot.py -t webrtc 
```

Open **http://localhost:7860** in your browser and click the mic button.

The prebuilt browser UI cannot pass patient context to the agent. In WebRTC mode, the agent uses a hardcoded test patient defined in `bot.py` (search for `"No body provided"`). To switch between test scenarios, edit the fallback `body`:

**Scenario 1 — Pre-authenticated (default):**

```python
body = {
    "patient_fhir_id": "d886a934-5568-42b3-9324-0f0b05fc018c",
    "patient_name": "Patricia Ann Williams",
    "verified": True,
}
```

**Scenario 2 — DOB verification:**

```python
body = {
    "caller_phone": "+15039264538",
    "patient_context": {
        "patient_fhir_id": "d886a934-5568-42b3-9324-0f0b05fc018c",
        "patient_name": "Patricia Ann Williams",
        "patient_birthdate": "1958-11-08",
    },
}
```

The agent will ask you to confirm the date of birth. Say the correct DOB to unlock clinical access.

**Scenario 3 — Unknown caller:**

```python
body = {}
```

The agent will explain that no account was found and decline access.

In production, patient context is always passed by the backend (via Daily transport or PipeCat Cloud), never hardcoded.

---

### Option B: Daily Transport (WebRTC via Daily.co)

Uses Daily.co rooms for WebRTC. Requires a `DAILY_API_KEY`. Patient context is passed via curl — no code changes needed to switch scenarios.

```bash
# Ensure DAILY_API_KEY is exported (dotenv alone won't work — the runner reads it before bot.py loads)
export DAILY_API_KEY=your-key-here

source .venv/bin/activate
python bot.py -t daily
```

Then start a session with curl. The response includes a `dailyRoom` URL — open it in your browser to join the voice session.

**Important:** The `"createDailyRoom": true` field is required in all curl commands below. Without it, the runner won't create a Daily room and the bot will fail.

**Scenario 1 — Pre-authenticated (skips DOB verification):**

```bash
curl -X POST http://localhost:7860/start \
  -H "Content-Type: application/json" \
  -d '{
    "createDailyRoom": true,
    "body": {
      "patient_fhir_id": "d886a934-5568-42b3-9324-0f0b05fc018c",
      "patient_name": "Patricia Ann Williams",
      "verified": true
    }
  }'
```

**Scenario 2 — DOB verification:**

```bash
curl -X POST http://localhost:7860/start \
  -H "Content-Type: application/json" \
  -d '{
    "createDailyRoom": true,
    "body": {
      "caller_phone": "+15039264538",
      "patient_context": {
        "patient_fhir_id": "d886a934-5568-42b3-9324-0f0b05fc018c",
        "patient_name": "Patricia Ann Williams",
        "patient_birthdate": "1958-11-08"
      }
    }
  }'
```

The agent will ask for DOB. Say "November 8, 1958" to pass verification and unlock clinical access.

**Scenario 3 — Unknown caller:**

```bash
curl -X POST http://localhost:7860/start \
  -H "Content-Type: application/json" \
  -d '{
    "createDailyRoom": true,
    "body": {
      "caller_phone": "+15559999999"
    }
  }'
```

The agent will explain that no account was found for this phone number.

---

### Option C: Daily with PSTN Dial-In (Phone Calls)

For actual phone call testing with a Daily phone number:

```bash
export DAILY_API_KEY=your-key-here
source .venv/bin/activate
python bot.py -t daily --dialin
```

In a separate terminal, expose the webhook:

```bash
ngrok http 7860
```

Configure your Daily phone number webhook to: `https://your-ngrok-url.ngrok.io/daily-dialin-webhook`

When a call comes in, the agent automatically looks up the caller's phone number in FHIR, then asks for DOB verification before granting clinical access.

## What You Can Ask

Once authenticated, the agent can answer questions like:

- "What were my last lab results?"
- "What medications am I on?"
- "Do I have any allergies?"
- "What was my last blood pressure?"
- "Is my A1C in normal range?"
- "When is my next appointment?"
- "What conditions do I have on file?"

The agent queries FHIR resources (Observation, MedicationRequest, Condition, AllergyIntolerance, Encounter, Appointment) and presents the data in natural spoken language.

## Deployment (PipeCat Cloud)

```bash
# Install pipecat CLI (requires Python 3.13 or lower — 3.14 is not yet supported)
uv tool install pipecat-ai-cli --python 3.13

# Auth to pipecat cloud
pipecat cloud auth login
```

### Configure Secrets

```bash
# Agent API keys
pipecat cloud secrets set langcare-secrets \
  DEEPGRAM_API_KEY=... \
  CARTESIA_API_KEY=... \
  ANTHROPIC_API_KEY=... \
  LANGCARE_MCP_URL=https://langcare-mcp-dev.fly.dev/mcp \
  LANGCARE_API_KEY=...
```

### Configure Docker Registry

The deploy requires image pull credentials. Create a Docker Hub access token at **Account Settings > Personal access tokens**, then:

```bash
# Create image pull secret (prompts for Docker Hub username + access token)
pipecat cloud secrets image-pull-secret langcare-image-creds https://index.docker.io/v1/
```

Add to `pcc-deploy.toml`:

```toml
image_credentials = "langcare-image-creds"
```

Or skip credentials for public images by passing `--no-credentials` to the deploy command.

### Build and Deploy

```bash
cd pipecat-agent/server
docker login
pipecat cloud docker build-push
pipecat cloud deploy
```

### Verify

```bash
pipecat cloud agent status langcare-voice-agent
```

## Alternate Deployment Architectures

PipeCat is a framework, not tied to Daily or any specific cloud. The pipeline structure, MCP FHIR integration, authentication flows, and clinical prompts remain the same regardless of deployment target. Only the transport layer and AI services change.

### Default: PipeCat Cloud + Daily

```
Phone / Browser → Daily (WebRTC/PSTN) → PipeCat Cloud
                                             ↓
                                     DeepGram STT → Claude → Cartesia TTS
                                             ↓
                                     LangCare MCP → EMR
```

- `bot.py` deploys via Docker with `pipecat cloud deploy`
- The `webhook/` directory is **not needed** — here's why:

**Why no webhook?** PipeCat's runner has a built-in `/daily-dialin-webhook` endpoint that handles inbound PSTN calls automatically. When a phone call comes in, Daily sends the call event directly to PipeCat Cloud, which creates a room, starts a bot instance, and passes the caller's phone number and call metadata to `bot()` via `args.body` as a `DailyDialinRequest`. The bot itself then does the FHIR patient lookup by phone number and handles DOB verification — no intermediate webhook server needed. In PipeCat Cloud, this is fully managed. For local development, running `python bot.py -t daily --dialin` enables the same built-in webhook at `http://localhost:7860/daily-dialin-webhook`.

### Enterprise: Genesys + GCP

For contact center deployments using Genesys and Google Cloud:

```
Genesys → AudioHook (WebSocket) → PipeCat Bot on GCP (Cloud Run / GKE)
                                       ↓
                                   Google STT → Gemini → Google TTS
                                       ↓
                                   LangCare MCP → EMR
```

What changes in `bot.py`:

| Layer | Default | Genesys + GCP |
|-------|---------|---------------|
| Transport | `DailyTransport` | `WebSocketTransport` |
| STT | DeepGram | Google STT |
| LLM | Claude (Anthropic) | Gemini (Google) |
| TTS | Cartesia | Google TTS |
| Hosting | PipeCat Cloud | GCP Cloud Run / GKE |
| MCP | LangCare on Fly.io | LangCare on Fly.io (unchanged) |

Install the Google services:

```bash
pip install 'pipecat-ai[google,websocket]'
```

The pipeline structure is identical:

```
transport.input() → STT → user_aggregator → LLM → TTS → transport.output() → assistant_aggregator
```

In this architecture, the `webhook/` directory becomes relevant. You deploy a FastAPI webhook server on GCP that:

1. Receives inbound call events from Genesys
2. Looks up the patient by phone number via MCP
3. Starts a PipeCat bot instance with WebSocket transport
4. Returns the WebSocket URL to Genesys for audio bridging

### Other Combinations

PipeCat supports mixing and matching any combination of:

- **Transport**: Daily, SmallWebRTC, WebSocket
- **STT**: DeepGram, Google, Azure, Whisper
- **LLM**: Claude, Gemini, OpenAI, any OpenAI-compatible API
- **TTS**: Cartesia, Google, ElevenLabs, Azure

The MCP FHIR layer, authentication flows, verification logic, and clinical prompts are all transport- and service-agnostic. Swap services by changing the constructor in `bot.py` — no other code changes needed.

## MCP Tools

Four FHIR tools are auto-discovered from the LangCare MCP server at startup:

| Tool | Description |
|------|-------------|
| `fhir_read` | Read a specific FHIR resource by type and ID |
| `fhir_search` | Search FHIR resources with query parameters |
| `fhir_create` | Create a new FHIR resource |
| `fhir_update` | Update an existing FHIR resource |

All FHIR queries are scoped to the authenticated patient's ID. The agent cannot access other patients' records within a session.

## Pipeline Order

The PipeCat pipeline must follow this exact sequence:

```
transport.input() → STT → user_aggregator → LLM → TTS → transport.output() → assistant_aggregator
```

`user_aggregator` must precede the LLM (feeds transcribed speech into context). `assistant_aggregator` must follow `transport.output()` (collects assistant responses after audio is sent).
