"""MCP HTTP client — handles initialize + tools/call session flow."""

import json
import httpx


class MCPError(Exception):
    """MCP or FHIR error with structured data."""
    def __init__(self, message: str, code: int = 0, error_type: str = "mcp_error"):
        self.message = message
        self.code = code
        self.error_type = error_type
        super().__init__(message)


class MCPClient:
    """Stateless MCP client — initialize + tools/call per invocation."""

    def __init__(self, server: str, token: str):
        self.server = server.rstrip("/")
        self.token = token
        self.mcp_url = f"{self.server}/mcp"

    def _headers(self, session_id: str = None) -> dict:
        headers = {
            "Authorization": f"Bearer {self.token}",
            "Content-Type": "application/json",
            "Accept": "application/json, text/event-stream",
        }
        if session_id:
            headers["Mcp-Session-Id"] = session_id
        return headers

    def _parse_response(self, resp: httpx.Response) -> dict:
        """Handle both application/json and text/event-stream responses."""
        content_type = resp.headers.get("content-type", "")

        if "text/event-stream" in content_type:
            # Parse SSE — find the data: line with our result
            for line in resp.text.splitlines():
                if line.startswith("data: "):
                    return json.loads(line[6:])
            raise MCPError("No data in SSE response", error_type="parse_error")
        else:
            return resp.json()

    def initialize(self) -> str | None:
        """
        Send MCP initialize request.
        Returns session_id if server provides one, else None.
        Session ID comes from HTTP response HEADER Mcp-Session-Id, not the body.
        Raises MCPError on failure.
        """
        try:
            resp = httpx.post(
                self.mcp_url,
                headers=self._headers(),
                json={
                    "jsonrpc": "2.0",
                    "method": "initialize",
                    "params": {
                        "protocolVersion": "2025-03-26",
                        "clientInfo": {
                            "name": "langcare-cli",
                            "version": "1.0.0"
                        },
                        "capabilities": {}
                    },
                    "id": 1
                },
                timeout=30.0
            )
        except httpx.ConnectError:
            raise MCPError(
                f"Connection refused to {self.server} — is the LangCare MCP server running?",
                error_type="network_error"
            )
        except httpx.TimeoutException:
            raise MCPError(
                f"Request timed out connecting to {self.server}",
                error_type="network_error"
            )

        if resp.status_code == 401:
            raise MCPError("Invalid bearer token", code=401, error_type="auth_error")
        if resp.status_code == 403:
            raise MCPError("Forbidden — check your API token", code=403, error_type="auth_error")
        if resp.status_code >= 400:
            raise MCPError(
                f"Server error {resp.status_code}",
                code=resp.status_code,
                error_type="server_error"
            )

        # Session ID is in the HTTP HEADER not the response body
        session_id = resp.headers.get("Mcp-Session-Id")

        data = self._parse_response(resp)
        if "error" in data:
            raise MCPError(
                data["error"].get("message", "Initialize failed"),
                code=data["error"].get("code", 0),
                error_type="mcp_error"
            )

        return session_id

    def call_tool(self, tool_name: str, arguments: dict, session_id: str = None) -> dict:
        """
        Call an MCP tool.
        Returns the parsed FHIR result as a dict.
        Raises MCPError on failure.
        """
        try:
            resp = httpx.post(
                self.mcp_url,
                headers=self._headers(session_id),
                json={
                    "jsonrpc": "2.0",
                    "method": "tools/call",
                    "params": {
                        "name": tool_name,
                        "arguments": arguments
                    },
                    "id": 2
                },
                timeout=60.0  # FHIR queries can be slow
            )
        except httpx.ConnectError:
            raise MCPError(
                f"Connection refused to {self.server} — is the LangCare MCP server running?",
                error_type="network_error"
            )
        except httpx.TimeoutException:
            raise MCPError(
                "Request timed out — FHIR query took too long",
                error_type="network_error"
            )

        if resp.status_code == 401:
            raise MCPError("Invalid bearer token", code=401, error_type="auth_error")
        if resp.status_code == 403:
            raise MCPError("Forbidden", code=403, error_type="auth_error")
        if resp.status_code == 404:
            raise MCPError("Session expired — retry", code=404, error_type="session_error")
        if resp.status_code >= 400:
            raise MCPError(
                f"Server error {resp.status_code}",
                code=resp.status_code,
                error_type="server_error"
            )

        data = self._parse_response(resp)

        if "error" in data:
            raise MCPError(
                data["error"].get("message", "Tool call failed"),
                code=data["error"].get("code", 0),
                error_type="fhir_error"
            )

        # Extract FHIR data from MCP result envelope
        # FHIR data is in result.content[0].text as a JSON string — must json.loads() it
        try:
            content = data["result"]["content"][0]["text"]
            return json.loads(content)
        except (KeyError, IndexError, json.JSONDecodeError):
            # Return raw result if structure is unexpected
            return data.get("result", data)

    def run(self, tool_name: str, arguments: dict) -> dict:
        """
        Full flow: initialize → tools/call.
        This is what every CLI command calls.
        Re-initializes on every invocation — CLI is stateless by design.
        """
        session_id = self.initialize()
        return self.call_tool(tool_name, arguments, session_id)
