import type { App, McpUiHostContext } from "@modelcontextprotocol/ext-apps";
import { useApp, useHostStyles } from "@modelcontextprotocol/ext-apps/react";
import type { CallToolResult } from "@modelcontextprotocol/sdk/types.js";
import { StrictMode, useCallback, useEffect, useRef, useState } from "react";
import { createRoot } from "react-dom/client";
import "./global.css";

// ---------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------

const RESOURCE_TYPES = [
  "Patient",
  "Observation",
  "Condition",
  "MedicationRequest",
  "MedicationStatement",
  "AllergyIntolerance",
  "Procedure",
  "Encounter",
  "DiagnosticReport",
  "Immunization",
  "CarePlan",
  "CareTeam",
  "DocumentReference",
  "Practitioner",
  "Organization",
  "Location",
  "Device",
  "Coverage",
  "Claim",
  "ExplanationOfBenefit",
] as const;

const SEARCH_PRESETS: Record<string, { label: string; params: string }[]> = {
  Patient: [
    { label: "By name", params: "name=Smith" },
    { label: "By birthdate", params: "birthdate=gt1990-01-01" },
    { label: "By identifier", params: "identifier=12345" },
  ],
  Observation: [
    { label: "Vitals", params: "category=vital-signs" },
    { label: "Labs", params: "category=laboratory" },
    { label: "By LOINC", params: "code=http://loinc.org|2339-0" },
  ],
  Condition: [
    { label: "Active", params: "clinical-status=active" },
    { label: "By category", params: "category=encounter-diagnosis" },
  ],
  MedicationRequest: [
    { label: "Active", params: "status=active" },
    { label: "By patient", params: "patient=example" },
  ],
};

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

function extractResultData(result: CallToolResult): unknown | null {
  const textItem = result.content?.find(
    (c: { type: string }) => c.type === "text"
  );
  if (!textItem || !("text" in textItem)) return null;
  try {
    return JSON.parse((textItem as { text: string }).text);
  } catch {
    return (textItem as { text: string }).text;
  }
}

/** Pull resource ID from a FHIR resource or fullUrl */
function extractResourceId(
  entry: Record<string, unknown>
): { resourceType: string; id: string } | null {
  const resource = entry.resource as Record<string, unknown> | undefined;
  if (resource?.id && resource?.resourceType) {
    return {
      resourceType: resource.resourceType as string,
      id: resource.id as string,
    };
  }
  return null;
}

// ---------------------------------------------------------------------------
// Root wrapper
// ---------------------------------------------------------------------------

function FhirExplorerRoot() {
  const [hostContext, setHostContext] = useState<McpUiHostContext | undefined>();
  const [initialResult, setInitialResult] = useState<CallToolResult | null>(
    null
  );

  const { app, error } = useApp({
    appInfo: { name: "FHIR Explorer", version: "1.0.0" },
    capabilities: {},
    onAppCreated: (app) => {
      app.ontoolresult = (result) => {
        setInitialResult(result);
      };
      app.ontoolinput = () => {};
      app.onteardown = async () => ({});
      app.onerror = console.error;
      app.onhostcontextchanged = (ctx) => {
        setHostContext((prev) => ({ ...prev, ...ctx }));
      };
    },
  });

  useHostStyles(app);

  useEffect(() => {
    if (app) {
      setHostContext(app.getHostContext());
    }
  }, [app]);

  if (error) return <div className="error-banner">{error.message}</div>;
  if (!app) return <div className="loading">Connecting...</div>;

  return (
    <FhirExplorer
      app={app}
      initialResult={initialResult}
      hostContext={hostContext}
    />
  );
}

// ---------------------------------------------------------------------------
// Main component
// ---------------------------------------------------------------------------

interface FhirExplorerProps {
  app: App;
  initialResult: CallToolResult | null;
  hostContext?: McpUiHostContext;
}

function FhirExplorer({ app, initialResult, hostContext }: FhirExplorerProps) {
  const [resourceType, setResourceType] = useState("Patient");
  const [queryParams, setQueryParams] = useState("");
  const [loading, setLoading] = useState(false);
  const [searchResult, setSearchResult] = useState<unknown | null>(null);
  const [detailResult, setDetailResult] = useState<unknown | null>(null);
  const [detailLoading, setDetailLoading] = useState(false);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);
  const [view, setView] = useState<"search" | "detail">("search");
  const resultRef = useRef<HTMLDivElement>(null);

  // Handle initial tool result (from host-initiated call)
  useEffect(() => {
    if (initialResult) {
      const data = extractResultData(initialResult);
      if (data) {
        setSearchResult(data);
        setView("search");
      }
    }
  }, [initialResult]);

  const handleSearch = useCallback(async () => {
    setLoading(true);
    setErrorMsg(null);
    setSearchResult(null);
    setDetailResult(null);
    setView("search");

    try {
      const result = await app.callServerTool({
        name: "fhir_search",
        arguments: {
          resourceType,
          ...(queryParams ? { queryParams } : {}),
        },
      });

      if (result.isError) {
        const msg = extractResultData(result);
        setErrorMsg(typeof msg === "string" ? msg : JSON.stringify(msg));
      } else {
        setSearchResult(extractResultData(result));
      }
    } catch (e) {
      setErrorMsg(e instanceof Error ? e.message : String(e));
    } finally {
      setLoading(false);
    }
  }, [app, resourceType, queryParams]);

  const handleReadResource = useCallback(
    async (resType: string, id: string) => {
      setDetailLoading(true);
      setErrorMsg(null);
      setView("detail");

      try {
        const result = await app.callServerTool({
          name: "fhir_read",
          arguments: { resourceType: resType, id },
        });

        if (result.isError) {
          const msg = extractResultData(result);
          setErrorMsg(typeof msg === "string" ? msg : JSON.stringify(msg));
        } else {
          setDetailResult(extractResultData(result));
        }
      } catch (e) {
        setErrorMsg(e instanceof Error ? e.message : String(e));
      } finally {
        setDetailLoading(false);
      }
    },
    [app]
  );

  const handlePreset = useCallback(
    (params: string) => {
      setQueryParams(params);
    },
    []
  );

  const safeArea = hostContext?.safeAreaInsets;

  return (
    <main
      className="explorer"
      style={{
        paddingTop: safeArea?.top,
        paddingRight: safeArea?.right,
        paddingBottom: safeArea?.bottom,
        paddingLeft: safeArea?.left,
      }}
    >
      {/* Header */}
      <header className="header">
        <h1 className="title">FHIR Explorer</h1>
        <span className="subtitle">LangCare MCP</span>
      </header>

      {/* Search form */}
      <section className="search-form">
        <div className="form-row">
          <select
            className="select"
            value={resourceType}
            onChange={(e) => setResourceType(e.target.value)}
          >
            {RESOURCE_TYPES.map((rt) => (
              <option key={rt} value={rt}>
                {rt}
              </option>
            ))}
          </select>

          <input
            className="input"
            type="text"
            placeholder="Query params (e.g. name=Smith&birthdate=gt1990-01-01)"
            value={queryParams}
            onChange={(e) => setQueryParams(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") handleSearch();
            }}
          />

          <button
            className="btn btn-primary"
            onClick={handleSearch}
            disabled={loading}
          >
            {loading ? "Searching..." : "Search"}
          </button>
        </div>

        {/* Presets */}
        {SEARCH_PRESETS[resourceType] && (
          <div className="presets">
            <span className="presets-label">Presets:</span>
            {SEARCH_PRESETS[resourceType].map((preset) => (
              <button
                key={preset.params}
                className="btn btn-preset"
                onClick={() => handlePreset(preset.params)}
              >
                {preset.label}
              </button>
            ))}
          </div>
        )}
      </section>

      {/* Error */}
      {errorMsg && <div className="error-banner">{errorMsg}</div>}

      {/* Results */}
      <section className="results" ref={resultRef}>
        {view === "detail" && (
          <button
            className="btn btn-back"
            onClick={() => setView("search")}
          >
            Back to results
          </button>
        )}

        {view === "search" && searchResult && (
          <BundleView
            data={searchResult}
            onReadResource={handleReadResource}
          />
        )}

        {view === "detail" && detailLoading && (
          <div className="loading">Loading resource...</div>
        )}

        {view === "detail" && !detailLoading && detailResult && (
          <JsonView data={detailResult} />
        )}
      </section>
    </main>
  );
}

// ---------------------------------------------------------------------------
// Bundle view — shows search results as a list of entries
// ---------------------------------------------------------------------------

interface BundleViewProps {
  data: unknown;
  onReadResource: (resourceType: string, id: string) => void;
}

function BundleView({ data, onReadResource }: BundleViewProps) {
  if (!data || typeof data !== "object") {
    return <JsonView data={data} />;
  }

  const bundle = data as Record<string, unknown>;
  const entries = bundle.entry as Record<string, unknown>[] | undefined;
  const total = bundle.total as number | undefined;

  if (!entries || !Array.isArray(entries)) {
    return <JsonView data={data} />;
  }

  return (
    <div className="bundle">
      <div className="bundle-header">
        <span className="bundle-count">
          {entries.length} result{entries.length !== 1 ? "s" : ""}
          {total !== undefined && total !== entries.length
            ? ` (${total} total)`
            : ""}
        </span>
      </div>

      <div className="entry-list">
        {entries.map((entry, idx) => {
          const ref = extractResourceId(entry);
          const resource = entry.resource as Record<string, unknown> | undefined;

          return (
            <div key={idx} className="entry-card">
              <div className="entry-header">
                <span className="entry-type">
                  {(resource?.resourceType as string) || "Unknown"}
                </span>
                <span className="entry-id">{(resource?.id as string) || ""}</span>
                {ref && (
                  <button
                    className="btn btn-read"
                    onClick={() =>
                      onReadResource(ref.resourceType, ref.id)
                    }
                  >
                    Read
                  </button>
                )}
              </div>
              <ResourceSummary resource={resource} />
            </div>
          );
        })}
      </div>
    </div>
  );
}

// ---------------------------------------------------------------------------
// Resource summary — shows key fields inline
// ---------------------------------------------------------------------------

function ResourceSummary({
  resource,
}: {
  resource: Record<string, unknown> | undefined;
}) {
  if (!resource) return null;

  const rt = resource.resourceType as string;
  const fields: { label: string; value: string }[] = [];

  if (rt === "Patient") {
    const names = resource.name as { text?: string; family?: string; given?: string[] }[] | undefined;
    if (names?.[0]) {
      const n = names[0];
      fields.push({
        label: "Name",
        value: n.text || [n.family, ...(n.given || [])].filter(Boolean).join(", "),
      });
    }
    if (resource.birthDate) fields.push({ label: "DOB", value: resource.birthDate as string });
    if (resource.gender) fields.push({ label: "Gender", value: resource.gender as string });
  } else if (rt === "Observation") {
    const code = resource.code as { text?: string; coding?: { display?: string }[] } | undefined;
    if (code?.text || code?.coding?.[0]?.display) {
      fields.push({ label: "Code", value: code.text || code.coding![0].display! });
    }
    if (resource.status) fields.push({ label: "Status", value: resource.status as string });
  } else if (rt === "Condition") {
    const code = resource.code as { text?: string; coding?: { display?: string }[] } | undefined;
    if (code?.text || code?.coding?.[0]?.display) {
      fields.push({ label: "Condition", value: code.text || code.coding![0].display! });
    }
    const cs = resource.clinicalStatus as { coding?: { code?: string }[] } | undefined;
    if (cs?.coding?.[0]?.code) {
      fields.push({ label: "Status", value: cs.coding[0].code });
    }
  } else if (rt === "MedicationRequest") {
    const med = resource.medicationCodeableConcept as { text?: string; coding?: { display?: string }[] } | undefined;
    if (med?.text || med?.coding?.[0]?.display) {
      fields.push({ label: "Medication", value: med.text || med.coding![0].display! });
    }
    if (resource.status) fields.push({ label: "Status", value: resource.status as string });
  }

  if (fields.length === 0) return null;

  return (
    <div className="summary-fields">
      {fields.map((f) => (
        <span key={f.label} className="summary-field">
          <span className="field-label">{f.label}:</span> {f.value}
        </span>
      ))}
    </div>
  );
}

// ---------------------------------------------------------------------------
// JSON view — formatted, collapsible JSON
// ---------------------------------------------------------------------------

function JsonView({ data }: { data: unknown }) {
  const [collapsed, setCollapsed] = useState<Set<string>>(new Set());

  const toggle = useCallback((path: string) => {
    setCollapsed((prev) => {
      const next = new Set(prev);
      if (next.has(path)) next.delete(path);
      else next.add(path);
      return next;
    });
  }, []);

  return (
    <pre className="json-view">
      <JsonNode data={data} path="$" depth={0} collapsed={collapsed} onToggle={toggle} />
    </pre>
  );
}

interface JsonNodeProps {
  data: unknown;
  path: string;
  depth: number;
  collapsed: Set<string>;
  onToggle: (path: string) => void;
}

function JsonNode({ data, path, depth, collapsed, onToggle }: JsonNodeProps) {
  const indent = "  ".repeat(depth);
  const innerIndent = "  ".repeat(depth + 1);

  if (data === null) return <span className="json-null">null</span>;
  if (typeof data === "boolean")
    return <span className="json-bool">{String(data)}</span>;
  if (typeof data === "number")
    return <span className="json-number">{data}</span>;
  if (typeof data === "string")
    return <span className="json-string">"{data}"</span>;

  if (Array.isArray(data)) {
    if (data.length === 0) return <span>{"[]"}</span>;

    const isCollapsed = collapsed.has(path);

    return (
      <span>
        <span className="json-bracket clickable" onClick={() => onToggle(path)}>
          [{isCollapsed ? `... ${data.length} items` : ""}
        </span>
        {!isCollapsed && (
          <>
            {"\n"}
            {data.map((item, i) => (
              <span key={i}>
                {innerIndent}
                <JsonNode
                  data={item}
                  path={`${path}[${i}]`}
                  depth={depth + 1}
                  collapsed={collapsed}
                  onToggle={onToggle}
                />
                {i < data.length - 1 ? "," : ""}
                {"\n"}
              </span>
            ))}
            {indent}
          </>
        )}
        <span className="json-bracket">{isCollapsed ? "" : "]"}</span>
        {isCollapsed && <span className="json-bracket">]</span>}
      </span>
    );
  }

  if (typeof data === "object") {
    const entries = Object.entries(data as Record<string, unknown>);
    if (entries.length === 0) return <span>{"{}"}</span>;

    const isCollapsed = collapsed.has(path);

    return (
      <span>
        <span className="json-bracket clickable" onClick={() => onToggle(path)}>
          {"{"}{isCollapsed ? `... ${entries.length} keys` : ""}
        </span>
        {!isCollapsed && (
          <>
            {"\n"}
            {entries.map(([key, val], i) => (
              <span key={key}>
                {innerIndent}
                <span className="json-key">"{key}"</span>
                {": "}
                <JsonNode
                  data={val}
                  path={`${path}.${key}`}
                  depth={depth + 1}
                  collapsed={collapsed}
                  onToggle={onToggle}
                />
                {i < entries.length - 1 ? "," : ""}
                {"\n"}
              </span>
            ))}
            {indent}
          </>
        )}
        <span className="json-bracket">{isCollapsed ? "" : "}"}</span>
        {isCollapsed && <span className="json-bracket">{"}"}</span>}
      </span>
    );
  }

  return <span>{String(data)}</span>;
}

// ---------------------------------------------------------------------------
// Mount
// ---------------------------------------------------------------------------

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <FhirExplorerRoot />
  </StrictMode>
);
