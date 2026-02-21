# MCP Apps — Interactive UIs for LangCare MCP FHIR

MCP Apps are interactive, React-based UIs that run directly inside MCP-capable hosts (Claude Desktop, etc.). Each app is a self-contained single-file HTML bundle embedded into the Go binary at compile time, served as an MCP Resource, and triggered by a dedicated MCP Tool. This repo includes two reference apps — **FHIR Explorer** and **Patient Chart Review** — demonstrating the pattern end-to-end.

<p align="center">
  <img src="patient-chart-review.png" alt="Patient Chart Review — clinical dashboard with vitals trends, conditions, medications, and labs" width="660" />
  <br />
  <em>Patient Chart Review — clinical dashboard running inside Claude Desktop</em>
</p>

## How MCP Apps Work

```
┌──────────────────────────────────────────────────────────────────────┐
│  MCP Host (Claude Desktop)                                           │
│                                                                      │
│  ┌──────────────────────────────────────────────────────────────┐    │
│  │  MCP App (React UI)                                          │    │
│  │                                                              │    │
│  │  app.callServerTool("fhir_search", {...})  ──────────────┐   │    │
│  │  app.callServerTool("fhir_read", {...})    ──────────┐   │   │    │
│  └──────────────────────────────────────────────────────│───│───┘    │
│                                                         │   │        │
└─────────────────────────────────────────────────────────│───│────────┘
                                                          │   │
                              MCP Protocol                │   │
                                                          v   v
┌──────────────────────────────────────────────────────────────────────┐
│  LangCare MCP FHIR Server (Go)                                       │
│                                                                      │
│  ┌────────────┐  ┌────────────┐  ┌─────────────┐  ┌──────────────┐  │
│  │ fhir_read  │  │fhir_search │  │ fhir_create │  │ fhir_update  │  │
│  └─────┬──────┘  └─────┬──────┘  └──────┬──────┘  └──────┬───────┘  │
│        └───────────────┬┘               │                 │          │
│                        v                v                 v          │
│                   FHIR Client (OAuth2 / JWT)                         │
│                        │                                             │
└────────────────────────│─────────────────────────────────────────────┘
                         v
                 FHIR R4 Server (EPIC, Cerner, GCP, etc.)
```

The key insight: **apps call FHIR tools directly through the MCP protocol** via `app.callServerTool()`. No LLM is involved in data fetching. The host renders the UI; the app handles all data logic deterministically.

## Built-in Apps

### FHIR Explorer (`fhir_explorer`)

Interactive FHIR resource browser. Search, read, create, and update any of 20+ FHIR R4 resource types with full JSON detail views, search presets, and syntax-highlighted output.

**Source:** `apps/fhir-explorer/src/app.tsx`

### Patient Chart Review (`patient_chart_review`)

Clinical dashboard for reviewing a patient's medical record. Includes:
- Patient search and demographics banner
- Four data cards: Active Conditions, Medications, Vitals, Labs
- Vitals Trends dashboard with SVG charts for Blood Pressure and Weight over selectable date ranges (1M/3M/6M/1Y)
- Click-to-expand detail panels for any item (fetches full resource via `fhir_read`)

**Source:** `apps/patient-chart-review/src/app.tsx`

## Architecture

### Build Pipeline

```
apps/{app-name}/src/app.tsx          # React + TypeScript source
        │
        v  (Vite + vite-plugin-singlefile)
apps/dist/{app-name}.html            # Single-file HTML (all JS/CSS inlined)
        │
        v  (scripts/build-apps.sh copies)
internal/apps/dist/{app-name}.html   # Go embed source directory
        │
        v  (go:embed at compile time)
Binary (langcare-mcp-fhir)           # HTML embedded in Go binary
        │
        v  (server.go registerApps)
MCP Resource + MCP Tool              # Served to MCP hosts at runtime
```

### Key Files

| File | Role |
|------|------|
| `apps/package.json` | Shared dependencies: React 19, `@modelcontextprotocol/ext-apps`, Vite 6 |
| `apps/vite.config.ts` | Dynamic Vite config; reads `INPUT`/`OUTDIR` env vars, uses `vite-plugin-singlefile` |
| `apps/tsconfig.json` | TypeScript config (ESNext, strict, React JSX transform) |
| `scripts/build-apps.sh` | Iterates app directories, builds each with Vite, copies to `internal/apps/dist/` |
| `internal/apps/embed.go` | `go:embed dist/*` directive — embeds all HTML bundles into the binary |
| `internal/apps/registry.go` | `DefaultApps` array defining each app's metadata (name, tool name, resource URI, HTML file) |
| `internal/mcp/server.go` | `registerApps()` — registers each app as an MCP Resource + linked MCP Tool |

### Registration Flow (server.go)

For each app in `DefaultApps`:

1. **Load HTML** from the embedded filesystem (`apps.LoadAppHTML`)
2. **Register MCP Resource** with MIME type `text/html;profile=mcp-app` — the `ReadResource` handler returns the HTML
3. **Register MCP Tool** with `_meta.ui.resourceUri` pointing to the resource — when an MCP host calls this tool, it fetches the linked resource and renders the UI

```go
// Resource registration
s.mcpServer.AddResource(&mcp.Resource{
    URI:      "ui://patient-chart-review/app.html",
    MIMEType: "text/html;profile=mcp-app",
    // ...
}, readHandler)

// Tool registration with UI link
s.mcpServer.AddTool(&mcp.Tool{
    Name: "patient_chart_review",
    Meta: mcp.Meta{
        "ui": map[string]interface{}{
            "resourceUri": "ui://patient-chart-review/app.html",
        },
    },
    // ...
}, callHandler)
```

### App SDK Pattern

Every app follows the same React pattern using `@modelcontextprotocol/ext-apps`:

```tsx
import type { App } from "@modelcontextprotocol/ext-apps";
import { useApp, useHostStyles } from "@modelcontextprotocol/ext-apps/react";

function MyAppRoot() {
  const { app, error } = useApp({
    appInfo: { name: "My App", version: "1.0.0" },
    capabilities: {},
    onAppCreated: (app) => {
      app.ontoolresult = (result) => { /* handle initial tool result */ };
      app.ontoolinput = () => {};
      app.onteardown = async () => ({});
      app.onerror = console.error;
      app.onhostcontextchanged = (ctx) => { /* handle theme/layout changes */ };
    },
  });

  useHostStyles(app); // Inherit host theme

  // Call FHIR tools directly — no LLM round-trip
  const result = await app.callServerTool({
    name: "fhir_search",
    arguments: { resourceType: "Patient", queryParams: "name=Smith" },
  });
}
```

`app.callServerTool()` calls back into the MCP server's tool handlers through the host. The response is a standard `CallToolResult` with `content` array containing text items (JSON-stringified FHIR responses).

## Building Apps

### Prerequisites

- Node.js 18+
- npm

### Build All Apps

```bash
make apps
# or directly:
bash scripts/build-apps.sh
```

This:
1. Installs npm dependencies in `apps/`
2. Builds each app subdirectory into a single-file HTML bundle
3. Copies bundles to `internal/apps/dist/` for Go embedding

### Build Full Binary (Apps + Go)

```bash
make build       # Builds apps first, then Go binary
make build-all   # Builds apps + cross-platform Go binaries
```

### Build Go Only (Skip Apps)

When you're only changing Go code and apps haven't changed:

```bash
make build-go
```

### Clean App Artifacts

```bash
make apps-clean
```

## Adding a New App

### 1. Create the App Directory

```bash
mkdir -p apps/my-new-app/src
```

Create `apps/my-new-app/index.html`:

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>My New App</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="./src/app.tsx"></script>
  </body>
</html>
```

Create `apps/my-new-app/src/global.css` with your styles.

Create `apps/my-new-app/src/app.tsx`:

```tsx
import type { App } from "@modelcontextprotocol/ext-apps";
import { useApp, useHostStyles } from "@modelcontextprotocol/ext-apps/react";
import type { CallToolResult } from "@modelcontextprotocol/sdk/types.js";
import { StrictMode, useCallback, useEffect, useState } from "react";
import { createRoot } from "react-dom/client";
import "./global.css";

// Helper to extract text content from CallToolResult
function extractResultData(result: CallToolResult): unknown | null {
  const textItem = result.content?.find((c: { type: string }) => c.type === "text");
  if (!textItem || !("text" in textItem)) return null;
  try {
    return JSON.parse((textItem as { text: string }).text);
  } catch {
    return (textItem as { text: string }).text;
  }
}

function MyAppRoot() {
  const { app, error } = useApp({
    appInfo: { name: "My New App", version: "1.0.0" },
    capabilities: {},
    onAppCreated: (app) => {
      app.ontoolresult = () => {};
      app.ontoolinput = () => {};
      app.onteardown = async () => ({});
      app.onerror = console.error;
    },
  });

  useHostStyles(app);

  if (error) return <div>{error.message}</div>;
  if (!app) return <div>Connecting...</div>;

  return <MyApp app={app} />;
}

function MyApp({ app }: { app: App }) {
  const [data, setData] = useState<unknown>(null);

  const fetchData = useCallback(async () => {
    const result = await app.callServerTool({
      name: "fhir_search",
      arguments: {
        resourceType: "Patient",
        queryParams: "name=Smith&_count=5",
      },
    });
    setData(extractResultData(result));
  }, [app]);

  useEffect(() => { fetchData(); }, [fetchData]);

  return (
    <main>
      <h1>My New App</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </main>
  );
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <MyAppRoot />
  </StrictMode>
);
```

### 2. Register the App in Go

Add an entry to `DefaultApps` in `internal/apps/registry.go`:

```go
{
    Name:        "my-new-app",
    ToolName:    "my_new_app",
    ToolDesc:    "Description for the LLM explaining when to use this tool.",
    ResourceURI: "ui://my-new-app/app.html",
    Description: "Human-readable description of the app.",
    HTMLFile:    "my-new-app.html",
},
```

The `ToolDesc` field matters — it's what the LLM sees when deciding which tool to call. Write it from the LLM's perspective: when should it open this UI?

### 3. Build and Test

```bash
make build       # Builds apps + Go binary
```

Restart Claude Desktop. The new tool will appear in the MCP tools list. When the LLM calls it, the host renders your app.

### 4. (Optional) Fly.io Dockerfile

If deploying to Fly.io, the `fly/Dockerfile` already runs `scripts/build-apps.sh` as part of the multi-stage build. No changes needed unless you add system-level dependencies.

## App Development Tips

### Calling FHIR Tools

All apps use the same 4 server tools. The response is always a `CallToolResult` with JSON text:

```tsx
// Search
const result = await app.callServerTool({
  name: "fhir_search",
  arguments: { resourceType: "Observation", queryParams: "patient=Patient/123&category=vital-signs" },
});

// Read
const result = await app.callServerTool({
  name: "fhir_read",
  arguments: { resourceType: "Patient", id: "123" },
});

// Create
const result = await app.callServerTool({
  name: "fhir_create",
  arguments: { resourceType: "Observation", resource: JSON.stringify(observationObj) },
});

// Update
const result = await app.callServerTool({
  name: "fhir_update",
  arguments: { resourceType: "Patient", id: "123", resource: JSON.stringify(patientObj) },
});
```

### Parsing FHIR Bundles

Search results come back as FHIR Bundles. Extract entries:

```tsx
function extractBundleEntries(data: unknown): unknown[] {
  if (!data || typeof data !== "object") return [];
  const bundle = data as Record<string, unknown>;
  const entries = bundle.entry as Record<string, unknown>[] | undefined;
  if (!entries || !Array.isArray(entries)) return [];
  return entries.map((e) => e.resource).filter(Boolean);
}
```

### Host Context and Theming

Use `useHostStyles(app)` to inherit the host's theme. Use CSS custom properties (`light-dark()`, `color-scheme: light dark`) for automatic dark mode support. The host provides safe area insets via `app.getHostContext()` for proper padding.

### Single-File Constraint

`vite-plugin-singlefile` inlines all JS, CSS, and assets into the HTML file. This means:

- **No external imports at runtime** — everything must be bundled
- **No dynamic imports** — code splitting won't work
- **Images must be inlined** — use SVG or base64, not external URLs
- **Keep bundle size reasonable** — avoid large charting libraries; inline SVG works well

### Concurrent Tool Calls

MCP tool calls are async. Fire independent calls in parallel with `Promise.allSettled`:

```tsx
const configs = [
  { name: "fhir_search", arguments: { resourceType: "Condition", queryParams: "..." } },
  { name: "fhir_search", arguments: { resourceType: "MedicationRequest", queryParams: "..." } },
];
await Promise.allSettled(configs.map((c) => app.callServerTool(c)));
```

Be aware of FHIR server rate limits. If fetching many resources, stagger requests or reduce `_count` parameters.

## Directory Structure

```
apps/
├── README.md                        # This file
├── package.json                     # Shared deps (React 19, MCP Apps SDK, Vite)
├── vite.config.ts                   # Dynamic Vite config (INPUT/OUTDIR env vars)
├── tsconfig.json                    # Shared TypeScript config
├── fhir-explorer/
│   ├── index.html                   # Entry point
│   └── src/
│       ├── app.tsx                  # React app (585 lines)
│       └── global.css               # Styles (318 lines)
└── patient-chart-review/
    ├── index.html                   # Entry point
    └── src/
        ├── app.tsx                  # React app (1448 lines)
        └── global.css               # Styles (906 lines)
```

Build output lands in `internal/apps/dist/` and gets embedded into the Go binary via `internal/apps/embed.go`.
