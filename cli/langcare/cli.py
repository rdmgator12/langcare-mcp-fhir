"""LangCare CLI — FHIR tools over MCP for AI agents."""

import json
import sys
import click
from langcare.mcp_client import MCPClient, MCPError


def output_error(error_type: str, code: int, message: str, exit_code: int):
    """Print structured error JSON to stdout and exit."""
    print(json.dumps({
        "error": error_type,
        "code": code,
        "message": message
    }))
    sys.exit(exit_code)


def run_tool(ctx, tool_name: str, arguments: dict):
    """Execute a tool call and print result. Used by all fhir subcommands."""
    server = ctx.obj["server"]
    token = ctx.obj["token"]

    if not token:
        output_error("auth_error", 0, "LANGCARE_API_KEY is not set", exit_code=2)

    client = MCPClient(server=server, token=token)

    try:
        result = client.run(tool_name, arguments)
        print(json.dumps(result))
        sys.exit(0)

    except MCPError as e:
        exit_code = {
            "auth_error":    2,
            "network_error": 3,
            "session_error": 1,
            "fhir_error":    1,
            "mcp_error":     1,
            "server_error":  1,
            "parse_error":   1,
        }.get(e.error_type, 1)
        output_error(e.error_type, e.code, e.message, exit_code)


# ── Root command ──────────────────────────────────────────────────────────────

@click.group()
@click.option(
    "--server",
    envvar="LANGCARE_SERVER",
    default="http://localhost:8080",
    show_default=True,
    help="LangCare MCP server URL [env: LANGCARE_SERVER]",
)
@click.option(
    "--token",
    envvar="LANGCARE_API_KEY",
    default="",
    help="Bearer token for MCP authentication [env: LANGCARE_API_KEY]",
)
@click.version_option(version="1.0.0", prog_name="langcare")
@click.pass_context
def cli(ctx, server, token):
    """LangCare CLI — FHIR tools over MCP.

    Designed to be called by AI agents as a subprocess tool.
    Always outputs JSON to stdout.

    \b
    The LangCare MCP server can be running locally on your machine
    or remotely on Fly.io. Default is http://localhost:8080 (local).

    \b
    Environment variables:
      LANGCARE_SERVER    MCP server URL (default: http://localhost:8080)
      LANGCARE_API_KEY   Bearer token (required)

    \b
    Examples (local server):
      langcare fhir search Patient --query "name=John"
      langcare fhir read Patient 123

    \b
    Examples (remote Fly.io server):
      LANGCARE_SERVER=https://langcare-mcp-prod.fly.dev \\
      langcare fhir search Patient --query "name=John"
    """
    ctx.ensure_object(dict)
    ctx.obj["server"] = server
    ctx.obj["token"] = token


# ── fhir subgroup ─────────────────────────────────────────────────────────────

@cli.group()
def fhir():
    """FHIR resource operations (search, read, create, update)."""
    pass


# ── fhir search ───────────────────────────────────────────────────────────────

@fhir.command()
@click.argument("resource_type")
@click.option(
    "--query", "-q",
    default="",
    help="FHIR search query parameters (e.g. 'patient=123&category=laboratory')",
)
@click.pass_context
def search(ctx, resource_type, query):
    """Search FHIR resources with query parameters.

    \b
    Examples:
      langcare fhir search Patient --query "name=John&birthdate=gt1990-01-01"
      langcare fhir search Observation --query "patient=123&category=laboratory&_sort=-date"
      langcare fhir search MedicationRequest --query "patient=123&status=active"
      langcare fhir search AllergyIntolerance --query "patient=123"
      langcare fhir search Condition --query "patient=123&clinical-status=active"
    """
    run_tool(ctx, "fhir_search", {
        "resourceType": resource_type,
        "queryParams": query,
    })


# ── fhir read ─────────────────────────────────────────────────────────────────

@fhir.command()
@click.argument("resource_type")
@click.argument("resource_id")
@click.pass_context
def read(ctx, resource_type, resource_id):
    """Read a FHIR resource by type and ID.

    \b
    Examples:
      langcare fhir read Patient 123
      langcare fhir read Observation abc-456
      langcare fhir read MedicationRequest xyz-789
    """
    run_tool(ctx, "fhir_read", {
        "resourceType": resource_type,
        "id": resource_id,
    })


# ── fhir create ───────────────────────────────────────────────────────────────

@fhir.command()
@click.argument("resource_type")
@click.option(
    "--data", "-d",
    required=True,
    help="FHIR resource JSON string or @filename to read from file",
)
@click.pass_context
def create(ctx, resource_type, data):
    """Create a new FHIR resource.

    \b
    Examples:
      langcare fhir create Observation --data '{"resourceType":"Observation","status":"final",...}'
      langcare fhir create Patient --data @patient.json
    """
    if data.startswith("@"):
        filepath = data[1:]
        try:
            with open(filepath) as f:
                data = f.read()
        except FileNotFoundError:
            output_error("invalid_args", 0, f"File not found: {filepath}", exit_code=4)

    try:
        resource = json.loads(data)
    except json.JSONDecodeError as e:
        output_error("invalid_args", 0, f"Invalid JSON: {e}", exit_code=4)

    run_tool(ctx, "fhir_create", {
        "resourceType": resource_type,
        "resource": resource,
    })


# ── fhir update ───────────────────────────────────────────────────────────────

@fhir.command()
@click.argument("resource_type")
@click.argument("resource_id")
@click.option(
    "--data", "-d",
    required=True,
    help="Updated FHIR resource JSON string or @filename to read from file",
)
@click.pass_context
def update(ctx, resource_type, resource_id, data):
    """Update an existing FHIR resource.

    \b
    Examples:
      langcare fhir update Patient 123 --data '{"resourceType":"Patient","id":"123",...}'
      langcare fhir update MedicationRequest abc --data @med.json
    """
    if data.startswith("@"):
        filepath = data[1:]
        try:
            with open(filepath) as f:
                data = f.read()
        except FileNotFoundError:
            output_error("invalid_args", 0, f"File not found: {filepath}", exit_code=4)

    try:
        resource = json.loads(data)
    except json.JSONDecodeError as e:
        output_error("invalid_args", 0, f"Invalid JSON: {e}", exit_code=4)

    run_tool(ctx, "fhir_update", {
        "resourceType": resource_type,
        "id": resource_id,
        "resource": resource,
    })


if __name__ == "__main__":
    cli()
