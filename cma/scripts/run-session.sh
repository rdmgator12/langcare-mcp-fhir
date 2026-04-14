#!/usr/bin/env bash
# Run a CMA agent session — interactive or single-prompt mode
#
# Usage:
#   ./run-session.sh <agent-id> <env-id> <vault-id>              # interactive REPL
#   ./run-session.sh <agent-id> <env-id> <vault-id> "prompt"     # single prompt
set -euo pipefail

AGENT_ID="${1:?Usage: run-session.sh <agent-id> <env-id> <vault-id> [\"prompt\"]}"
ENV_ID="${2:?}"
VAULT_ID="${3:?}"
PROMPT="${4:-}"

SCRIPT=$(mktemp /tmp/run_session_XXXXXX.py)
trap "rm -f $SCRIPT" EXIT

cat > "$SCRIPT" << 'PYEOF'
import sys, json, os, time, textwrap
import urllib.request, urllib.error

api_key = os.environ["ANTHROPIC_API_KEY"]
agent_id, env_id, vault_id, initial_prompt = sys.argv[1], sys.argv[2], sys.argv[3], sys.argv[4]

BASE = "https://api.anthropic.com"
HEADERS = {
    "x-api-key": api_key,
    "anthropic-version": "2023-06-01",
    "anthropic-beta": "managed-agents-2026-04-01",
    "content-type": "application/json",
}

# ── ANSI colors ───────────────────────────────────────────────────────────────
RESET   = "\033[0m"
BOLD    = "\033[1m"
DIM     = "\033[2m"
BLINK   = "\033[5m"
CYAN    = "\033[36m"
GREEN   = "\033[32m"
YELLOW  = "\033[33m"
BLUE    = "\033[34m"
MAGENTA = "\033[35m"
WHITE   = "\033[97m"
BG_DARK = "\033[48;5;235m"

def c(text, *codes):
    return "".join(codes) + text + RESET

# ── Agent-specific example questions ─────────────────────────────────────────
EXAMPLES = {
    "Medication Management": [
        "Show active medications for patient ID <id>",
        "Check drug interactions for patient ID <id>",
        "Run Beers Criteria screening for patient ID <id>",
        "Assess medication adherence for patient ID <id>",
        "Evaluate opioid risk for patient ID <id>",
    ],
    "Care Coordination": [
        "Create a discharge plan for patient ID <id>",
        "Identify care gaps for patient ID <id>",
        "List open referrals for patient ID <id>",
        "Summarize transitions of care for patient ID <id>",
        "What follow-up tasks are pending for patient ID <id>?",
    ],
    "Clinical Decision Support": [
        "Calculate qSOFA score for patient ID <id>",
        "Assess cardiovascular risk for patient ID <id>",
        "Evaluate VTE prophylaxis need for patient ID <id>",
        "Run fall risk assessment for patient ID <id>",
        "Calculate CURB-65 score for patient ID <id>",
    ],
    "Clinical Triage": [
        "Summarize chief complaint for patient ID <id>",
        "Review vitals and flag abnormals for patient ID <id>",
        "What is the acuity level for patient ID <id>?",
        "Check for sepsis indicators for patient ID <id>",
        "Give a clinical summary for patient ID <id>",
    ],
    "Documentation": [
        "Write a SOAP note for patient ID <id>",
        "Generate a discharge summary for patient ID <id>",
        "Create a progress note for patient ID <id>",
        "Draft a referral letter for patient ID <id>",
        "Write a procedure note for patient ID <id>",
    ],
    "Lab & Diagnostics": [
        "Review recent lab results for patient ID <id>",
        "Flag any critical values for patient ID <id>",
        "Interpret the diabetes panel for patient ID <id>",
        "Review pre-op labs for patient ID <id>",
        "Assess renal function for patient ID <id>",
    ],
    "Patient Data": [
        "Show demographics for patient ID <id>",
        "Review allergy list for patient ID <id>",
        "Give a full clinical summary for patient ID <id>",
        "Check insurance coverage for patient ID <id>",
        "Audit the problem list for patient ID <id>",
    ],
    "Population Health": [
        "Show chronic disease registry for my panel",
        "Check immunization status for patient ID <id>",
        "Review preventive care compliance for patient ID <id>",
        "What quality measures are unmet for patient ID <id>?",
        "Give a panel overview for my patient population",
    ],
    "Specialty Care": [
        "Assess chronic pain management for patient ID <id>",
        "Review mental health status for patient ID <id>",
        "Summarize oncology treatment for patient ID <id>",
        "Check pediatric growth percentiles for patient ID <id>",
        "Review prenatal care for patient ID <id>",
    ],
}

def api(method, path, body=None):
    req = urllib.request.Request(
        f"{BASE}{path}",
        data=json.dumps(body).encode() if body else None,
        headers=HEADERS,
        method=method,
    )
    try:
        with urllib.request.urlopen(req) as r:
            return json.load(r)
    except urllib.error.HTTPError as e:
        print(c(f"\nAPI error {e.code}: {e.read().decode()}", BOLD, YELLOW), file=sys.stderr)
        sys.exit(1)

def send_message(session_id, text):
    api("POST", f"/v1/sessions/{session_id}/events", {
        "events": [{"type": "user.message", "content": [{"type": "text", "text": text}]}]
    })

def wait_for_idle(session_id, timeout=120):
    deadline = time.time() + timeout
    seen_running = False
    sys.stdout.write(c("  Thinking", DIM, CYAN))
    sys.stdout.flush()
    while time.time() < deadline:
        time.sleep(0.5)
        s = api("GET", f"/v1/sessions/{session_id}")
        status = s["status"]
        if status == "running":
            if not seen_running:
                seen_running = True
            sys.stdout.write(c(".", DIM, CYAN))
            sys.stdout.flush()
        elif status == "idle" and seen_running:
            sys.stdout.write("\n")
            return True
        elif status == "error":
            sys.stdout.write("\n")
            print(c("Session error.", BOLD, YELLOW), file=sys.stderr)
            return False
    sys.stdout.write("\n")
    print(c("Timed out.", BOLD, YELLOW), file=sys.stderr)
    return False

def fetch_response(session_id, after_event_id=None):
    path = f"/v1/sessions/{session_id}/events"
    if after_event_id:
        path += f"?after={after_event_id}"
    events = api("GET", path).get("data", [])
    parts = []
    last_id = None
    for e in events:
        last_id = e.get("id")
        if e.get("type") == "agent.message":
            for block in e.get("content", []):
                if block.get("type") == "text":
                    parts.append(("text", block["text"]))
        elif e.get("type") in ("agent.tool_use", "agent.mcp_tool_use"):
            parts.append(("tool", e.get("name", "")))
    return parts, last_id

def print_response(parts):
    for kind, value in parts:
        if kind == "tool":
            print(c(f"  ⚙  {value}", DIM, MAGENTA))
        elif kind == "text":
            # Wrap long lines for readability
            print(c(value, WHITE))

def print_banner(agent_info):
    full_name = agent_info.get("name", "LangCare Agent")
    # Strip "LangCare — " prefix to get just the agent speciality
    agent_name = full_name.replace("LangCare — ", "").replace("LangCare - ", "").strip()
    desc = agent_info.get("description", "").strip()
    width = 64

    print()
    print(c("╔" + "═" * (width - 2) + "╗", BOLD, CYAN))
    print(c("║" + "  LangCare: A Claude Managed Agent".ljust(width - 2) + "║", BOLD, CYAN))
    print(c("║" + f"  {agent_name}".ljust(width - 2) + "║", BOLD, WHITE))
    print(c("╠" + "═" * (width - 2) + "╣", BOLD, CYAN))
    if desc:
        for line in textwrap.wrap(desc, width - 4):
            print(c("║" + f"  {line}".ljust(width - 2) + "║", DIM, CYAN))
        print(c("║" + " " * (width - 2) + "║", DIM, CYAN))

    # Match agent name to examples
    examples = []
    for key, qs in EXAMPLES.items():
        if key.lower() in agent_name.lower():
            examples = qs
            break

    if examples:
        print(c("║" + "  Example questions:".ljust(width - 2) + "║", BOLD, CYAN))
        for q in examples:
            line = f"    • {q}"
            print(c("║" + line[:width - 2].ljust(width - 2) + "║", DIM, CYAN))

    print(c("╚" + "═" * (width - 2) + "╝", BOLD, CYAN))
    print()

# ── Fetch agent info ──────────────────────────────────────────────────────────

agent_info = api("GET", f"/v1/agents/{agent_id}")
print_banner(agent_info)

# ── Create session ────────────────────────────────────────────────────────────

session = api("POST", "/v1/sessions", {
    "agent": agent_id,
    "environment_id": env_id,
    "vault_ids": [vault_id],
})
session_id = session["id"]
print(c(f"  Session : {session_id}", DIM))
print(c(f"  Portal  : claude.ai → Sessions", DIM))
print()

# ── Single prompt mode ────────────────────────────────────────────────────────

if initial_prompt:
    print(c("You: ", BOLD, GREEN) + initial_prompt)
    send_message(session_id, initial_prompt)
    if wait_for_idle(session_id):
        parts, _ = fetch_response(session_id)
        print()
        print(c("Agent:", BOLD, BLUE))
        print_response(parts)
    sys.exit(0)

# ── Interactive mode ──────────────────────────────────────────────────────────

print(c("  Type your message and press Enter. /quit to exit, /id for session ID.", DIM))
print(c("─" * 64, DIM))

last_event_id = None

while True:
    try:
        sys.stdout.write(c("\nYou › ", BOLD, GREEN) + BLINK + "█" + RESET + " ")
        sys.stdout.flush()
        user_input = input("").strip()
    except (EOFError, KeyboardInterrupt):
        print(c("\nExiting.", DIM))
        break

    if not user_input:
        continue
    if user_input == "/quit":
        print(c("Exiting.", DIM))
        break
    if user_input == "/id":
        print(c(f"  Session: {session_id}", DIM))
        continue

    send_message(session_id, user_input)

    if not wait_for_idle(session_id):
        break

    parts, last_event_id = fetch_response(session_id, after_event_id=last_event_id)
    print()
    print(c("Agent:", BOLD, BLUE))
    print_response(parts)
    print(c("─" * 64, DIM))

PYEOF

python3 "$SCRIPT" "$AGENT_ID" "$ENV_ID" "$VAULT_ID" "$PROMPT"
