# OpenEMR Backend Services (JWT Bearer) Authentication Guide

This document describes how to set up SMART on FHIR Backend Services authentication for an OpenEMR FHIR API client used with the LangCare MCP FHIR Server.

OpenEMR's `client_credentials` + `private_key_jwt` flow follows the [SMART on FHIR Backend Services specification](https://hl7.org/fhir/smart-app-launch/backend-services.html) — the same family of auth that EPIC uses. The mechanics are nearly identical to EPIC: you generate an RSA key pair, register the public key (as a JWKS) with OpenEMR, and your application signs short-lived JWT assertions with the private key to obtain access tokens.

---

## Table of Contents

1. [Overview](#overview)
2. [Differences from EPIC](#differences-from-epic)
3. [Prerequisites](#prerequisites)
4. [Step 1: Generate RSA Key Pair](#step-1-generate-rsa-key-pair)
5. [Step 2: Create JWKS with the Helper Script](#step-2-create-jwks-with-the-helper-script)
6. [Step 3: Register the API Client in OpenEMR](#step-3-register-the-api-client-in-openemr)
7. [Step 4: Enable the API Client](#step-4-enable-the-api-client)
8. [Step 5: Configure LangCare MCP FHIR Server](#step-5-configure-langcare-mcp-fhir-server)
9. [Step 6: Run and Verify](#step-6-run-and-verify)
10. [Step 7: Point Claude Desktop at the OpenEMR Config](#step-7-point-claude-desktop-at-the-openemr-config)
11. [Authentication Flow](#authentication-flow)
12. [Troubleshooting](#troubleshooting)
13. [References](#references)

---

## Overview

```
┌─────────────┐                                    ┌──────────────┐
│  LangCare   │                                    │   OpenEMR    │
│  MCP FHIR   │                                    │  OAuth2 +    │
│  Server     │                                    │  FHIR APIs   │
│             │                                    │              │
│ Sign JWT    │──── client_assertion (RS384) ────▶│  Verify with │
│ with        │                                    │  registered  │
│ private key │                                    │  JWKS        │
│             │◀──── access_token (60s TTL) ──────│              │
│             │                                    │              │
│ Use bearer  │──── GET /apis/default/fhir/... ──▶│              │
│ token       │                                    │              │
└─────────────┘                                    └──────────────┘
```

The LangCare MCP FHIR Server's OpenEMR provider lives in `internal/fhir/providers/openemr.go` and handles JWT signing, token caching (refresh when <10s remain), and bearer-token injection automatically.

---

## Differences from EPIC

| Aspect | EPIC | OpenEMR |
|---|---|---|
| Spec | SMART Backend Services | SMART Backend Services |
| Algorithm | RS384 | RS384 (required) |
| Key registration | EPIC App Orchard (JWKS upload or JWKS URL) | OpenEMR Admin UI (paste JWKS into client config) |
| Access token TTL | ~1 hour | **60 seconds** (no refresh tokens) |
| JWT assertion TTL | ≤ 5 minutes | ≤ 5 minutes |
| Token endpoint | `https://<host>/interconnect-fhir-oauth/oauth2/token` | `https://<host>/oauth2/default/token` |
| FHIR base URL | `https://<host>/interconnect-fhir-oauth/api/FHIR/R4` | `https://<host>/apis/default/fhir` |
| Allowed scopes | `system/*`, `user/*`, `patient/*` | **`system/*` only** for this flow |

Because OpenEMR access tokens expire in 60 seconds, the OpenEMR provider refreshes tokens whenever fewer than 10 seconds remain (vs. EPIC's 5-minute buffer). This is handled automatically.

---

## Prerequisites

- An OpenEMR instance with the FHIR API and SMART on FHIR Backend Services enabled
- Admin access to OpenEMR (Administration → System → API Clients)
- `openssl` installed locally
- A built copy of the LangCare MCP FHIR Server (`make build`)

---

## Step 1: Generate RSA Key Pair

OpenEMR's documentation recommends a 4096-bit RSA key.

```bash
mkdir -p keys/openemr
cd keys/openemr

# Generate a 4096-bit RSA private key
openssl genrsa -out private.key 4096

# Extract the matching public key
openssl rsa -in private.key -pubout -out public.pem
```

**Keep `private.key` secret.** Anyone with that file can impersonate your client.

Verify the keys:

```bash
openssl rsa -in private.key -text -noout | head
openssl rsa -pubin -in public.pem -text -noout | head
```

The OpenEMR provider supports both PKCS1 and PKCS8 PEM formats (same as EPIC), so no conversion is required.

---

## Step 2: Create JWKS with the Helper Script

The repo ships with `scripts/create_jwks_openemr.sh`, which converts `public.pem` into a JWKS that OpenEMR can consume.

```bash
cd scripts
chmod +x create_jwks_openemr.sh
./create_jwks_openemr.sh
```

The script:

1. Reads `../keys/openemr/public.pem`
2. Extracts the RSA modulus and base64url-encodes it
3. Writes `../keys/openemr/jwks.json` with `kty=RSA`, `use=sig`, `alg=RS384`, `kid=openemr-key-001`
4. Prints next steps

The resulting JWKS looks like:

```json
{
  "keys": [
    {
      "kty": "RSA",
      "kid": "openemr-key-001",
      "use": "sig",
      "alg": "RS384",
      "n": "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78...",
      "e": "AQAB"
    }
  ]
}
```

> **Important:** the `kid` value in this JWKS must match the `key_id` field in your LangCare config (Step 5). The script defaults both to `openemr-key-001`. If you change one, change both.

To use a different Key ID, edit `KEY_ID` near the top of `scripts/create_jwks_openemr.sh`.

---

## Step 3: Register the API Client in OpenEMR

1. Log in to OpenEMR as an administrator.
2. Navigate to **Administration → System → API Clients**.
3. Click **Register New API Client**.
4. Fill in:

   | Field | Value |
   |---|---|
   | **Client Name** | e.g. `LangCare MCP FHIR` |
   | **Client Type** | `Confidential` |
   | **Grant Types** | `Client Credentials` |
   | **Authentication Method** | `private_key_jwt` *(required for this flow)* |
   | **Scopes** | System-level scopes you need, e.g. `system/Patient.read system/Observation.read system/Condition.read system/MedicationRequest.read system/AllergyIntolerance.read system/Encounter.read system/Practitioner.read system/Organization.read` |
   | **JSON Web Key Set** | Paste the **entire contents** of `keys/openemr/jwks.json` from Step 2 |

5. Click **Save**.
6. **Copy the generated Client ID** — you will need it for the config file. It looks something like `Lsi_SwHsgGVdcUYb5-R4-5Tr_4cjQovbEx-K0gpGmY0`.

> **Scope rule:** the scopes your application requests at token time must be a subset of the scopes registered on the API client. Only `system/*` scopes are allowed for `client_credentials` + `private_key_jwt` — `user/*` and `patient/*` scopes require a different OAuth flow.

---

## Step 4: Enable the API Client

Newly registered clients are disabled by default.

1. Go back to **Administration → System → API Clients**.
2. Find your new client in the list.
3. Click **Enable**.

The client is now ready to authenticate.

---

## Step 5: Configure LangCare MCP FHIR Server

Copy the example config and edit it:

```bash
cp configs/config.openemr.example.yaml configs/config.local.openemr.yaml
```

Set the four OpenEMR-specific fields:

```yaml
fhir_server:
  provider: "openemr"
  base_url: "https://your-openemr-host/apis/default/fhir"

  openemr:
    client_id: "Lsi_SwHsgGVdcUYb5-R4-5Tr_4cjQovbEx-K0gpGmY0"   # from Step 3
    private_key_path: "/absolute/path/to/keys/openemr/private.key"
    token_url: "https://your-openemr-host/oauth2/default/token"
    key_id: "openemr-key-001"   # must match the kid in jwks.json
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"
      - "system/Condition.read"
      - "system/MedicationRequest.read"
      - "system/AllergyIntolerance.read"
      - "system/Encounter.read"
      - "system/Practitioner.read"
      - "system/Organization.read"
```

**Critical rules:**

- `token_url` becomes the JWT `aud` claim and is validated **byte-for-byte** by OpenEMR. Match scheme, host, port, and path exactly with what OpenEMR expects to see.
- `key_id` must match the `kid` field of the JWK in OpenEMR's registered JWKS.
- `client_id` must match the value OpenEMR generated in Step 3 — both `iss` and `sub` claims are sent as this value.

For a local OpenEMR running plain HTTP, use `http://` for both `base_url` and `token_url` — the Go HTTP client refuses to talk TLS to a non-TLS server. OpenEMR may need to be put in dev/insecure mode to accept non-HTTPS OAuth requests.

---

## Step 6: Run and Verify

```bash
make build
./bin/langcare-mcp-fhir -config configs/config.local.openemr.yaml
```

A successful start logs something like:

```
INFO creating fhir client provider=openemr base_url=https://...
INFO initializing openemr provider base_url=... client_id=...
INFO openemr access token refreshed expires_in=60
INFO openemr provider authenticated successfully
```

The provider pre-authenticates at startup, so any credential or registration mistake fails fast with a clear error.

---

## Step 7: Point Claude Desktop at the OpenEMR Config

Switching Claude Desktop from EPIC (or any other provider) to OpenEMR is purely a config-file swap — the binary is the same. Edit your Claude Desktop config:

- **macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Linux:** `~/.config/Claude/claude_desktop_config.json`
- **Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

Change the `-config` argument to point at your OpenEMR YAML:

```json
{
  "mcpServers": {
    "langcare-mcp-fhir": {
      "command": "/absolute/path/to/bin/langcare-mcp-fhir",
      "args": [
        "-config",
        "/absolute/path/to/configs/config.local.openemr.yaml"
      ]
    }
  }
}
```

If you installed via npm, the `command` is just `langcare-mcp-fhir` (or use `npx @langcare/langcare-mcp-fhir`):

```json
{
  "mcpServers": {
    "langcare-mcp-fhir": {
      "command": "langcare-mcp-fhir",
      "args": ["-config", "/absolute/path/to/configs/config.local.openemr.yaml"]
    }
  }
}
```

**Then fully quit and reopen Claude Desktop** (not just close the window — use Cmd-Q on macOS). Claude only re-reads `claude_desktop_config.json` and re-spawns MCP servers on a full restart.

You can verify the switch worked by asking Claude something like *"List the MCP tools you have available"* and then running a `fhir_search` against a known OpenEMR Patient. The MCP server's startup log line `provider=openemr` is also visible in Claude Desktop's MCP server log (`~/Library/Logs/Claude/mcp-server-langcare-mcp-fhir.log` on macOS).

> **Tip:** Keep separate config files per backend (`config.local.epic.yaml`, `config.local.openemr.yaml`, …) and switch by editing only the `-config` arg. You can also register multiple MCP server entries under different names (e.g. `langcare-fhir-epic`, `langcare-fhir-openemr`) pointing at different configs to have both backends available simultaneously in Claude.

---

## Authentication Flow

1. **Build JWT claims** (per token request):
   ```json
   {
     "iss": "<client_id>",
     "sub": "<client_id>",
     "aud": "https://your-openemr-host/oauth2/default/token",
     "iat": 1700000000,
     "exp": 1700000300,
     "jti": "<random-uuid>"
   }
   ```
2. **Build JWT header**:
   ```json
   { "alg": "RS384", "kid": "openemr-key-001", "typ": "JWT" }
   ```
3. **Sign** with the RSA private key (RS384) → `client_assertion`.
4. **POST** to the token endpoint:
   ```
   POST /oauth2/default/token
   Content-Type: application/x-www-form-urlencoded

   grant_type=client_credentials
   &client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer
   &client_assertion=<signed JWT>
   &scope=system/Patient.read system/Observation.read ...
   ```
5. **OpenEMR verifies** the signature using the JWK matching `kid`, validates `iss`/`sub`/`aud`/`exp`/`jti`, then issues a 60-second bearer token:
   ```json
   {
     "token_type": "Bearer",
     "expires_in": 60,
     "access_token": "eyJ0eXAiOiJKV1Qi...",
     "scope": "system/Patient.read system/Observation.read"
   }
   ```
6. **Use the bearer token** on FHIR requests:
   ```
   GET /apis/default/fhir/Patient/123
   Authorization: Bearer eyJ0eXAiOiJKV1Qi...
   Accept: application/fhir+json
   ```
7. **Auto-refresh:** the provider re-runs steps 1-5 transparently when the cached token has <10 seconds left.

---

## Troubleshooting

### `http: server gave HTTP response to HTTPS client`

Your `base_url` or `token_url` uses `https://` but OpenEMR is serving plain HTTP. Switch both to `http://` for local dev, or front OpenEMR with a TLS proxy.

### `invalid_client` / `Invalid client assertion signature`

OpenEMR could not verify the JWT signature. Causes:

- The private key in `private_key_path` does not match the public key in the registered JWKS. Re-derive the public key:
  ```bash
  openssl rsa -in private.key -pubout
  ```
  and re-run `create_jwks_openemr.sh`.
- The `kid` in your config (`key_id`) does not match the `kid` in OpenEMR's registered JWKS.
- The JWKS was edited after registration but not re-saved in OpenEMR.

### `invalid_grant: audience claim mismatch`

The `aud` claim in the JWT (= `token_url` in your config) does not exactly match what OpenEMR expects. Common mistakes: trailing slash, wrong port, `http` vs `https`, missing `/oauth2/default/token` path.

### `invalid_scope` / 403 on FHIR requests

The scope you requested is not registered on the API client in OpenEMR. Edit the API client in the OpenEMR admin UI and add the missing scope.

### `Authentication Method` field missing in OpenEMR admin

The OpenEMR instance is too old or does not have SMART on FHIR Backend Services enabled. Enable it under `Administration → Globals → Connectors` (look for FHIR / OAuth2 / SMART settings) and re-check after a reload.

### Tokens expire in the middle of long operations

This is expected — OpenEMR access tokens are 60 seconds with no refresh tokens. The provider's automatic refresh handles it; if you see auth failures during long FHIR exports, increase log level to `debug` to confirm refresh is firing:

```yaml
logging:
  level: "debug"
```

### `failed to load private key`

The path in `private_key_path` is wrong, the file is not readable by the process, or the PEM is corrupted. Use an absolute path and verify with `openssl rsa -in <path> -text -noout`.

---

## References

- OpenEMR API & FHIR documentation: https://github.com/openemr/openemr/blob/master/API_README.md
- SMART on FHIR Backend Services: https://hl7.org/fhir/smart-app-launch/backend-services.html
- RFC 7523 (JWT Bearer client authentication): https://datatracker.ietf.org/doc/html/rfc7523
- LangCare provider implementation: `internal/fhir/providers/openemr.go`
- Example config: `configs/config.openemr.example.yaml`
- JWKS helper script: `scripts/create_jwks_openemr.sh`
- Related: [`docs/EPIC-APP-SECURITY.md`](EPIC-APP-SECURITY.md)
