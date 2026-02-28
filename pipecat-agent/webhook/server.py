# Webhook server for inbound phone calls.
# Dependencies: fastapi, uvicorn, httpx
# Run: uvicorn webhook.server:app --port 8080

import os

import httpx
from fastapi import FastAPI, Request
from fastapi.responses import PlainTextResponse

app = FastAPI()

LANGCARE_MCP_URL = os.getenv("LANGCARE_MCP_URL", "https://langcare-mcp-dev.fly.dev")
LANGCARE_API_KEY = os.getenv("LANGCARE_API_KEY")
PIPECAT_API_KEY = os.getenv("PIPECAT_API_KEY")
PIPECAT_AGENT_URL = os.getenv(
    "PIPECAT_AGENT_URL",
    "https://api.pipecat.daily.co/agents/healthcare-voice-agent/start",
)


async def lookup_patient_by_phone(phone: str) -> dict | None:
    """Search FHIR for a patient matching this phone number."""
    phone_digits = phone.lstrip("+1") if phone.startswith("+1") else phone.lstrip("+")

    headers = {"Content-Type": "application/json"}
    if LANGCARE_API_KEY:
        headers["Authorization"] = f"Bearer {LANGCARE_API_KEY}"

    async with httpx.AsyncClient() as client:
        resp = await client.post(
            f"{LANGCARE_MCP_URL}/mcp",
            headers=headers,
            json={
                "jsonrpc": "2.0",
                "id": 1,
                "method": "tools/call",
                "params": {
                    "name": "fhir_search",
                    "arguments": {
                        "resourceType": "Patient",
                        "queryParams": f"phone={phone_digits}",
                    },
                },
            },
        )
        result = resp.json()

    # Parse FHIR Bundle from MCP response
    try:
        content = result.get("result", {}).get("content", [])
        if not content:
            return None

        import json

        fhir_data = json.loads(content[0].get("text", "{}"))

        if fhir_data.get("resourceType") != "Bundle":
            return None

        entries = fhir_data.get("entry", [])
        if not entries:
            return None

        patient = entries[0].get("resource", {})
        if patient.get("resourceType") != "Patient":
            return None

        names = patient.get("name", [{}])
        given = names[0].get("given", ["Unknown"])
        family = names[0].get("family", "Unknown")

        return {
            "patient_fhir_id": patient["id"],
            "patient_name": f"{given[0]} {family}",
            "patient_birthdate": patient.get("birthDate", ""),
        }
    except (KeyError, IndexError, TypeError):
        return None


async def start_pipecat_agent(agent_data: dict):
    """Start a PipeCat Cloud agent session with the given data."""
    headers = {"Content-Type": "application/json"}
    if PIPECAT_API_KEY:
        headers["Authorization"] = f"Bearer {PIPECAT_API_KEY}"

    async with httpx.AsyncClient() as client:
        resp = await client.post(
            PIPECAT_AGENT_URL,
            headers=headers,
            json={
                "data": agent_data,
                "use_daily": True,
            },
        )
        return resp.json()


@app.post("/call")
async def handle_inbound_call(request: Request):
    """Twilio/Daily PSTN webhook for incoming calls."""
    body = await request.form()
    caller_phone = body.get("From", "")
    call_sid = body.get("CallSid", "")

    patient_context = await lookup_patient_by_phone(caller_phone)

    agent_data = {
        "caller_phone": caller_phone,
        "call_sid": call_sid,
        "patient_context": patient_context,
    }

    await start_pipecat_agent(agent_data)

    twiml = """<?xml version="1.0" encoding="UTF-8"?>
    <Response>
        <Play loop="10">https://your-cdn.com/hold-music.mp3</Play>
    </Response>"""
    return PlainTextResponse(content=twiml, media_type="text/xml")
