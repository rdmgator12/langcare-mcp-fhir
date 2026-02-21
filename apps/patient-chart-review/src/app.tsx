import type { App, McpUiHostContext } from "@modelcontextprotocol/ext-apps";
import { useApp, useHostStyles } from "@modelcontextprotocol/ext-apps/react";
import type { CallToolResult } from "@modelcontextprotocol/sdk/types.js";
import { StrictMode, useCallback, useEffect, useState } from "react";
import { createRoot } from "react-dom/client";
import "./global.css";

// ---------------------------------------------------------------------------
// Types
// ---------------------------------------------------------------------------

interface PatientInfo {
  id: string;
  name: string;
  initials: string;
  dob: string;
  gender: string;
  mrn: string;
}

interface CardData {
  loading: boolean;
  error: string | null;
  items: unknown[];
}

type CardKey =
  | "conditions"
  | "medications"
  | "vitals"
  | "labs";

interface SearchConfig {
  key: CardKey;
  resourceType: string;
  params: string;
}

interface ExpandedState {
  loading: boolean;
  error: string | null;
  resource: Record<string, unknown> | null;
}

// Vitals Trends types
type TrendRange = "1m" | "3m" | "6m" | "1y";

const TREND_RANGES: { key: TrendRange; label: string; months: number }[] = [
  { key: "1m", label: "1M", months: 1 },
  { key: "3m", label: "3M", months: 3 },
  { key: "6m", label: "6M", months: 6 },
  { key: "1y", label: "1Y", months: 12 },
];

const LOINC_BP = "85354-9";
const LOINC_WEIGHT = "29463-7";
const LOINC_SYSTOLIC = "8480-6";
const LOINC_DIASTOLIC = "8462-4";

interface TrendPoint {
  date: string;
  timestamp: number;
}

interface BPPoint extends TrendPoint {
  systolic: number;
  diastolic: number;
}

interface WeightPoint extends TrendPoint {
  value: number;
  unit: string;
}

interface TrendData {
  loading: boolean;
  error: string | null;
  bp: BPPoint[];
  weight: WeightPoint[];
}

// ---------------------------------------------------------------------------
// Card metadata — icons, labels
// ---------------------------------------------------------------------------

const CARD_META: Record<CardKey, { icon: string; label: string; emptyIcon: string }> = {
  conditions:  { icon: "\u2764",   label: "Active Conditions",  emptyIcon: "\u2714\uFE0F" },
  medications: { icon: "\uD83D\uDC8A", label: "Medications",        emptyIcon: "\uD83D\uDC8A" },
  vitals:      { icon: "\uD83D\uDCC8", label: "Vitals",             emptyIcon: "\uD83D\uDCC9" },
  labs:        { icon: "\uD83E\uDDEA", label: "Recent Labs",        emptyIcon: "\uD83E\uDDEA" },
};

// ---------------------------------------------------------------------------
// Search configs factory
// ---------------------------------------------------------------------------

function getSearchConfigs(patientId: string): SearchConfig[] {
  const pid = `Patient/${patientId}`;
  return [
    { key: "conditions", resourceType: "Condition", params: `patient=${pid}&clinical-status=active` },
    { key: "medications", resourceType: "MedicationRequest", params: `patient=${pid}&status=active` },
    { key: "vitals", resourceType: "Observation", params: `patient=${pid}&category=vital-signs&_sort=-date&_count=10` },
    { key: "labs", resourceType: "Observation", params: `patient=${pid}&category=laboratory&_sort=-date&_count=10` },
  ];
}

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

function extractBundleEntries(data: unknown): unknown[] {
  if (!data || typeof data !== "object") return [];
  const bundle = data as Record<string, unknown>;
  const entries = bundle.entry as Record<string, unknown>[] | undefined;
  if (!entries || !Array.isArray(entries)) return [];
  return entries.map((e) => e.resource).filter(Boolean);
}

function formatPatientName(resource: Record<string, unknown>): string {
  const names = resource.name as
    | { text?: string; family?: string; given?: string[] }[]
    | undefined;
  if (!names?.[0]) return "Unknown";
  const n = names[0];
  if (n.text) return n.text;
  const parts = [...(n.given || []), n.family].filter(Boolean);
  return parts.join(" ") || "Unknown";
}

function getInitials(name: string): string {
  const parts = name.split(/\s+/).filter(Boolean);
  if (parts.length === 0) return "?";
  if (parts.length === 1) return parts[0][0].toUpperCase();
  return (parts[0][0] + parts[parts.length - 1][0]).toUpperCase();
}

function extractMRN(resource: Record<string, unknown>): string {
  const identifiers = resource.identifier as
    | { type?: { coding?: { code?: string }[] }; value?: string }[]
    | undefined;
  if (!identifiers) return "";
  const mrn = identifiers.find(
    (id) =>
      id.type?.coding?.some((c) => c.code === "MR") || id.value
  );
  return mrn?.value || identifiers[0]?.value || "";
}

function extractCodeDisplay(
  codeableConcept: unknown
): string {
  if (!codeableConcept || typeof codeableConcept !== "object") return "";
  const cc = codeableConcept as {
    text?: string;
    coding?: { display?: string; code?: string }[];
  };
  return cc.text || cc.coding?.[0]?.display || cc.coding?.[0]?.code || "";
}

function formatDate(dateStr: unknown): string {
  if (!dateStr || typeof dateStr !== "string") return "";
  return dateStr.slice(0, 10);
}

// ---------------------------------------------------------------------------
// Vitals Trend helpers
// ---------------------------------------------------------------------------

function hasLoincCode(resource: Record<string, unknown>, loincCode: string): boolean {
  const code = resource.code as { coding?: { code?: string; system?: string }[] } | undefined;
  if (!code?.coding) return false;
  return code.coding.some((c) => c.code === loincCode);
}

function extractBPPoints(resources: unknown[]): BPPoint[] {
  const points: BPPoint[] = [];
  for (const r of resources) {
    const res = r as Record<string, unknown>;
    if (!hasLoincCode(res, LOINC_BP)) continue;
    const components = res.component as
      | { code?: { coding?: { code?: string }[] }; valueQuantity?: { value?: number } }[]
      | undefined;
    if (!components) continue;
    let systolic: number | undefined;
    let diastolic: number | undefined;
    for (const c of components) {
      const codes = c.code?.coding;
      if (!codes) continue;
      if (codes.some((cd) => cd.code === LOINC_SYSTOLIC) && c.valueQuantity?.value !== undefined) {
        systolic = c.valueQuantity.value;
      }
      if (codes.some((cd) => cd.code === LOINC_DIASTOLIC) && c.valueQuantity?.value !== undefined) {
        diastolic = c.valueQuantity.value;
      }
    }
    if (systolic === undefined || diastolic === undefined) continue;
    const dateStr = (res.effectiveDateTime as string) || (res.issued as string) || "";
    if (!dateStr) continue;
    points.push({ date: dateStr.slice(0, 10), timestamp: new Date(dateStr).getTime(), systolic, diastolic });
  }
  return points.sort((a, b) => a.timestamp - b.timestamp);
}

function extractWeightPoints(resources: unknown[]): WeightPoint[] {
  const points: WeightPoint[] = [];
  for (const r of resources) {
    const res = r as Record<string, unknown>;
    if (!hasLoincCode(res, LOINC_WEIGHT)) continue;
    const vq = res.valueQuantity as { value?: number; unit?: string } | undefined;
    if (vq?.value === undefined) continue;
    const dateStr = (res.effectiveDateTime as string) || (res.issued as string) || "";
    if (!dateStr) continue;
    points.push({ date: dateStr.slice(0, 10), timestamp: new Date(dateStr).getTime(), value: vq.value, unit: vq.unit || "kg" });
  }
  return points.sort((a, b) => a.timestamp - b.timestamp);
}

function computePolylinePoints(
  values: number[],
  timestamps: number[],
  width: number,
  height: number,
  padding: { top: number; right: number; bottom: number; left: number },
  yMin: number,
  yMax: number,
): string {
  const plotW = width - padding.left - padding.right;
  const plotH = height - padding.top - padding.bottom;
  const yRange = yMax - yMin || 1;
  const tMin = timestamps[0];
  const tRange = (timestamps[timestamps.length - 1] - tMin) || 1;
  return values
    .map((v, i) => {
      const x = padding.left + ((timestamps[i] - tMin) / tRange) * plotW;
      const y = padding.top + plotH - ((v - yMin) / yRange) * plotH;
      return `${x.toFixed(1)},${y.toFixed(1)}`;
    })
    .join(" ");
}

function getStartDate(range: TrendRange): string {
  const months = TREND_RANGES.find((r) => r.key === range)!.months;
  const d = new Date();
  d.setMonth(d.getMonth() - months);
  return d.toISOString().slice(0, 10);
}

// ---------------------------------------------------------------------------
// Root wrapper
// ---------------------------------------------------------------------------

function ChartReviewRoot() {
  const [hostContext, setHostContext] = useState<McpUiHostContext | undefined>();
  const [initialResult, setInitialResult] = useState<CallToolResult | null>(
    null
  );

  const { app, error } = useApp({
    appInfo: { name: "Patient Chart Review", version: "1.0.0" },
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
    <ChartReview
      app={app}
      initialResult={initialResult}
      hostContext={hostContext}
    />
  );
}

// ---------------------------------------------------------------------------
// Main component
// ---------------------------------------------------------------------------

interface ChartReviewProps {
  app: App;
  initialResult: CallToolResult | null;
  hostContext?: McpUiHostContext;
}

function ChartReview({ app, initialResult, hostContext }: ChartReviewProps) {
  const [view, setView] = useState<"search" | "chart">("search");
  const [searchQuery, setSearchQuery] = useState("");
  const [searching, setSearching] = useState(false);
  const [patients, setPatients] = useState<Record<string, unknown>[]>([]);
  const [searchError, setSearchError] = useState<string | null>(null);
  const [selectedPatient, setSelectedPatient] = useState<PatientInfo | null>(
    null
  );
  const [cards, setCards] = useState<Record<CardKey, CardData>>({
    conditions: { loading: false, error: null, items: [] },
    medications: { loading: false, error: null, items: [] },
    vitals: { loading: false, error: null, items: [] },
    labs: { loading: false, error: null, items: [] },
  });
  const [activeToolCall, setActiveToolCall] = useState<string | null>(null);

  // Handle initial tool result
  useEffect(() => {
    if (initialResult) {
      const data = extractResultData(initialResult);
      if (data && typeof data === "object") {
        const entries = extractBundleEntries(data);
        if (entries.length > 0) {
          setPatients(entries as Record<string, unknown>[]);
        }
      }
    }
  }, [initialResult]);

  const handleSearch = useCallback(async () => {
    if (!searchQuery.trim()) return;
    setSearching(true);
    setSearchError(null);
    setPatients([]);

    try {
      const result = await app.callServerTool({
        name: "fhir_search",
        arguments: {
          resourceType: "Patient",
          queryParams: `name=${encodeURIComponent(searchQuery.trim())}`,
        },
      });

      if (result.isError) {
        const msg = extractResultData(result);
        setSearchError(typeof msg === "string" ? msg : JSON.stringify(msg));
      } else {
        const data = extractResultData(result);
        const entries = extractBundleEntries(data);
        setPatients(entries as Record<string, unknown>[]);
      }
    } catch (e) {
      setSearchError(e instanceof Error ? e.message : String(e));
    } finally {
      setSearching(false);
    }
  }, [app, searchQuery]);

  const runCardSearch = useCallback(
    async (config: SearchConfig) => {
      try {
        const result = await app.callServerTool({
          name: "fhir_search",
          arguments: { resourceType: config.resourceType, queryParams: config.params },
        });
        if (result.isError) {
          const msg = extractResultData(result);
          setCards((prev) => ({
            ...prev,
            [config.key]: {
              loading: false,
              error: typeof msg === "string" ? msg : JSON.stringify(msg),
              items: [],
            },
          }));
        } else {
          const data = extractResultData(result);
          const items = extractBundleEntries(data);
          setCards((prev) => ({
            ...prev,
            [config.key]: { loading: false, error: null, items },
          }));
        }
      } catch (e) {
        setCards((prev) => ({
          ...prev,
          [config.key]: {
            loading: false,
            error: e instanceof Error ? e.message : String(e),
            items: [],
          },
        }));
      }
    },
    [app]
  );

  const loadChart = useCallback(
    async (patient: Record<string, unknown>) => {
      const name = formatPatientName(patient);
      const info: PatientInfo = {
        id: patient.id as string,
        name,
        initials: getInitials(name),
        dob: formatDate(patient.birthDate),
        gender: (patient.gender as string) || "",
        mrn: extractMRN(patient),
      };
      setSelectedPatient(info);
      setView("chart");

      const configs = getSearchConfigs(info.id);

      const initCards: Record<CardKey, CardData> = {
        conditions: { loading: true, error: null, items: [] },
        medications: { loading: true, error: null, items: [] },
        vitals: { loading: true, error: null, items: [] },
        labs: { loading: true, error: null, items: [] },
      };
      setCards(initCards);

      await Promise.allSettled(configs.map(runCardSearch));
    },
    [runCardSearch]
  );

  const refreshCard = useCallback(
    async (cardKey: CardKey) => {
      if (!selectedPatient) return;
      const configs = getSearchConfigs(selectedPatient.id);
      const config = configs.find((c) => c.key === cardKey);
      if (!config) return;

      setCards((prev) => ({
        ...prev,
        [cardKey]: { loading: true, error: null, items: [] },
      }));
      setActiveToolCall(`fhir_search ${config.resourceType}`);

      try {
        await runCardSearch(config);
      } finally {
        setActiveToolCall(null);
      }
    },
    [selectedPatient, runCardSearch]
  );

  const handleBack = useCallback(() => {
    setView("search");
    setSelectedPatient(null);
  }, []);

  const safeArea = hostContext?.safeAreaInsets;

  return (
    <main
      className="app"
      style={{
        paddingTop: safeArea?.top,
        paddingRight: safeArea?.right,
        paddingBottom: safeArea?.bottom,
        paddingLeft: safeArea?.left,
      }}
    >
      <header className="header">
        <span className="header-icon">{"\uD83D\uDCCB"}</span>
        <h1 className="title">Patient Chart Review</h1>
        <span className="subtitle">LangCare MCP</span>
      </header>

      {view === "search" && (
        <PatientSearchView
          query={searchQuery}
          onQueryChange={setSearchQuery}
          onSearch={handleSearch}
          searching={searching}
          patients={patients}
          error={searchError}
          onSelect={loadChart}
        />
      )}

      {view === "chart" && selectedPatient && (
        <ChartView
          app={app}
          patient={selectedPatient}
          cards={cards}
          onBack={handleBack}
          onRefresh={refreshCard}
          activeToolCall={activeToolCall}
          setActiveToolCall={setActiveToolCall}
        />
      )}
    </main>
  );
}

// ---------------------------------------------------------------------------
// Patient Search View
// ---------------------------------------------------------------------------

interface PatientSearchViewProps {
  query: string;
  onQueryChange: (q: string) => void;
  onSearch: () => void;
  searching: boolean;
  patients: Record<string, unknown>[];
  error: string | null;
  onSelect: (patient: Record<string, unknown>) => void;
}

function PatientSearchView({
  query,
  onQueryChange,
  onSearch,
  searching,
  patients,
  error,
  onSelect,
}: PatientSearchViewProps) {
  return (
    <>
      <section className="search-form">
        <div className="form-row">
          <input
            className="input"
            type="text"
            placeholder="Search by patient name..."
            value={query}
            onChange={(e) => onQueryChange(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") onSearch();
            }}
          />
          <button
            className="btn btn-primary"
            onClick={onSearch}
            disabled={searching || !query.trim()}
          >
            {searching ? "Searching..." : "Search"}
          </button>
        </div>
      </section>

      {error && <div className="error-banner">{error}</div>}

      {patients.length === 0 && !searching && !error && (
        <div className="search-empty">
          <span className="search-empty-icon">{"\uD83D\uDC64"}</span>
          <span className="search-empty-text">Search for a patient to review their chart</span>
        </div>
      )}

      {patients.length > 0 && (
        <div className="patient-list">
          <div className="result-count">
            {patients.length} patient{patients.length !== 1 ? "s" : ""} found
          </div>
          {patients.map((p, i) => {
            const name = formatPatientName(p);
            return (
              <div
                key={(p.id as string) || i}
                className="patient-row"
                onClick={() => onSelect(p)}
              >
                <div className="patient-avatar">{getInitials(name)}</div>
                <div className="patient-info">
                  <span className="patient-name">{name}</span>
                  <div className="patient-meta">
                    {p.birthDate && (
                      <span className="patient-detail">
                        <span className="patient-detail-icon">{"\uD83D\uDCC5"}</span>
                        {formatDate(p.birthDate)}
                      </span>
                    )}
                    {p.gender && (
                      <span className="patient-detail">
                        {(p.gender as string) === "male" ? "\u2642\uFE0F" : (p.gender as string) === "female" ? "\u2640\uFE0F" : ""}{" "}
                        {p.gender as string}
                      </span>
                    )}
                    {extractMRN(p) && (
                      <span className="patient-detail">
                        <span className="patient-detail-icon">{"\uD83C\uDD94"}</span>
                        {extractMRN(p)}
                      </span>
                    )}
                  </div>
                </div>
                <span className="patient-chevron">{"\u203A"}</span>
              </div>
            );
          })}
        </div>
      )}
    </>
  );
}

// ---------------------------------------------------------------------------
// Tool Call Indicator
// ---------------------------------------------------------------------------

function ToolCallIndicator({ activeCall }: { activeCall: string | null }) {
  if (!activeCall) return null;
  return (
    <div className="tool-indicator">
      <div className="spinner spinner-sm" />
      <span>Calling {activeCall}...</span>
    </div>
  );
}

// ---------------------------------------------------------------------------
// Chart View
// ---------------------------------------------------------------------------

interface ChartViewProps {
  app: App;
  patient: PatientInfo;
  cards: Record<CardKey, CardData>;
  onBack: () => void;
  onRefresh: (key: CardKey) => void;
  activeToolCall: string | null;
  setActiveToolCall: (call: string | null) => void;
}

function ChartView({ app, patient, cards, onBack, onRefresh, activeToolCall, setActiveToolCall }: ChartViewProps) {
  const cardsLoaded = !Object.values(cards).some((c) => c.loading);

  return (
    <>
      <button className="btn btn-back" onClick={onBack}>
        {"\u2190"} Back to search
      </button>

      <div className="patient-banner">
        <div className="banner-avatar">{patient.initials}</div>
        <div className="banner-content">
          <div className="banner-name">{patient.name}</div>
          <div className="banner-details">
            {patient.dob && (
              <span className="banner-field">
                <span className="banner-field-icon">{"\uD83D\uDCC5"}</span>
                <span className="banner-label">DOB:</span> {patient.dob}
              </span>
            )}
            {patient.gender && (
              <span className="banner-field">
                <span className="banner-field-icon">
                  {patient.gender === "male" ? "\u2642\uFE0F" : patient.gender === "female" ? "\u2640\uFE0F" : "\u26A5\uFE0F"}
                </span>
                <span className="banner-label">Gender:</span> {patient.gender}
              </span>
            )}
            {patient.mrn && (
              <span className="banner-field">
                <span className="banner-field-icon">{"\uD83C\uDD94"}</span>
                <span className="banner-label">MRN:</span> {patient.mrn}
              </span>
            )}
          </div>
        </div>
      </div>

      <ToolCallIndicator activeCall={activeToolCall} />

      <VitalsTrendCard app={app} patientId={patient.id} setActiveToolCall={setActiveToolCall} cardsLoaded={cardsLoaded} />

      <div className="chart-grid">
        {(Object.keys(CARD_META) as CardKey[]).map((key) => (
          <ChartCard
            key={key}
            app={app}
            cardKey={key}
            data={cards[key]}
            renderItem={CARD_RENDERERS[key]}
            onRefresh={() => onRefresh(key)}
            setActiveToolCall={setActiveToolCall}
          />
        ))}
      </div>
    </>
  );
}

// ---------------------------------------------------------------------------
// Chart Card
// ---------------------------------------------------------------------------

interface ChartCardProps {
  app: App;
  cardKey: CardKey;
  data: CardData;
  renderItem: (item: unknown, index: number) => React.ReactNode;
  onRefresh: () => void;
  setActiveToolCall: (call: string | null) => void;
}

function ChartCard({ app, cardKey, data, renderItem, onRefresh, setActiveToolCall }: ChartCardProps) {
  const meta = CARD_META[cardKey];
  const [expandedIndex, setExpandedIndex] = useState<number | null>(null);
  const [expandedData, setExpandedData] = useState<ExpandedState>({
    loading: false,
    error: null,
    resource: null,
  });

  // Reset expanded state when card data reloads
  useEffect(() => {
    if (data.loading) {
      setExpandedIndex(null);
      setExpandedData({ loading: false, error: null, resource: null });
    }
  }, [data.loading]);

  const handleRowClick = useCallback(
    async (item: unknown, index: number) => {
      if (expandedIndex === index) {
        setExpandedIndex(null);
        setExpandedData({ loading: false, error: null, resource: null });
        return;
      }

      const r = item as Record<string, unknown>;
      const resourceType = r.resourceType as string;
      const id = r.id as string;
      if (!resourceType || !id) return;

      setExpandedIndex(index);
      setExpandedData({ loading: true, error: null, resource: null });
      setActiveToolCall(`fhir_read ${resourceType}/${id}`);

      try {
        const result = await app.callServerTool({
          name: "fhir_read",
          arguments: { resourceType, id },
        });
        if (result.isError) {
          const msg = extractResultData(result);
          setExpandedData({
            loading: false,
            error: typeof msg === "string" ? msg : JSON.stringify(msg),
            resource: null,
          });
        } else {
          const resource = extractResultData(result) as Record<string, unknown>;
          setExpandedData({ loading: false, error: null, resource });
        }
      } catch (e) {
        setExpandedData({
          loading: false,
          error: e instanceof Error ? e.message : String(e),
          resource: null,
        });
      } finally {
        setActiveToolCall(null);
      }
    },
    [app, expandedIndex, setActiveToolCall]
  );

  return (
    <div className={`card ${cardKey}`}>
      <div className="card-header">
        <div className={`card-icon ${cardKey}`}>{meta.icon}</div>
        <span className={`card-title ${cardKey}`}>{meta.label}</span>
        {!data.loading && !data.error && (
          <span className="card-count">{data.items.length}</span>
        )}
        <button
          className="btn-refresh"
          onClick={(e) => {
            e.stopPropagation();
            onRefresh();
          }}
          title="Refresh"
          disabled={data.loading}
        >
          {"\u21BB"}
        </button>
      </div>
      {data.loading && (
        <div className="card-loading">
          <div className="spinner" />
          Loading...
        </div>
      )}
      {data.error && (
        <div className="card-error">
          {"\u26A0\uFE0F"} {data.error}
        </div>
      )}
      {!data.loading && !data.error && data.items.length === 0 && (
        <div className="card-empty">
          <span className="card-empty-icon">{meta.emptyIcon}</span>
          None found
        </div>
      )}
      {!data.loading && !data.error && data.items.length > 0 && (
        <div className="card-content">
          {data.items.map((item, i) => (
            <div key={i}>
              <div
                className="row-clickable"
                onClick={() => handleRowClick(item, i)}
              >
                {renderItem(item, i)}
              </div>
              {expandedIndex === i && (
                <div className="expanded-panel">
                  {expandedData.loading && (
                    <div className="expanded-loading">
                      <div className="spinner spinner-sm" />
                      Loading details...
                    </div>
                  )}
                  {expandedData.error && (
                    <div className="expanded-error">{expandedData.error}</div>
                  )}
                  {expandedData.resource && !expandedData.loading && (
                    <ExpandedDetail resource={expandedData.resource} />
                  )}
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

// ---------------------------------------------------------------------------
// Expanded Detail
// ---------------------------------------------------------------------------

function ExpandedDetail({ resource }: { resource: Record<string, unknown> }) {
  const resourceType = resource.resourceType as string;
  const fields: { label: string; value: string }[] = [];

  if (resourceType === "Condition") {
    const vs = resource.verificationStatus as
      | { coding?: { code?: string; display?: string }[]; text?: string }
      | undefined;
    if (vs) {
      const val = vs.text || vs.coding?.[0]?.display || vs.coding?.[0]?.code || "";
      if (val) fields.push({ label: "Verification", value: val });
    }

    const severity = extractCodeDisplay(resource.severity);
    if (severity) fields.push({ label: "Severity", value: severity });

    const bodySites = resource.bodySite as unknown[] | undefined;
    if (bodySites?.length) {
      fields.push({
        label: "Body Site",
        value: bodySites.map((bs) => extractCodeDisplay(bs)).filter(Boolean).join(", "),
      });
    }

    const onset = formatDate(resource.onsetDateTime);
    if (onset) fields.push({ label: "Onset", value: onset });

    const recorded = formatDate(resource.recordedDate);
    if (recorded) fields.push({ label: "Recorded", value: recorded });

    const notes = resource.note as { text?: string }[] | undefined;
    if (notes?.length) {
      fields.push({ label: "Notes", value: notes.map((n) => n.text).filter(Boolean).join("; ") });
    }
  } else if (resourceType === "MedicationRequest") {
    const dosages = resource.dosageInstruction as
      | { text?: string; route?: unknown; timing?: { code?: unknown } }[]
      | undefined;
    if (dosages?.[0]) {
      const d = dosages[0];
      if (d.text) fields.push({ label: "Dosage", value: d.text });
      const route = extractCodeDisplay(d.route);
      if (route) fields.push({ label: "Route", value: route });
      const freq = extractCodeDisplay(d.timing?.code);
      if (freq) fields.push({ label: "Frequency", value: freq });
    }

    const requester = resource.requester as
      | { display?: string; reference?: string }
      | undefined;
    if (requester) {
      const val = requester.display || requester.reference || "";
      if (val) fields.push({ label: "Requester", value: val });
    }

    const intent = resource.intent as string | undefined;
    if (intent) fields.push({ label: "Intent", value: intent });
  } else if (resourceType === "Observation") {
    const interp = resource.interpretation as unknown[] | undefined;
    if (interp?.length) {
      const val = interp.map((i) => extractCodeDisplay(i)).filter(Boolean).join(", ");
      if (val) fields.push({ label: "Interpretation", value: val });
    }

    const refRanges = resource.referenceRange as
      | { text?: string; low?: { value?: number; unit?: string }; high?: { value?: number; unit?: string } }[]
      | undefined;
    if (refRanges?.length) {
      const rr = refRanges[0];
      let rangeStr = rr.text || "";
      if (!rangeStr && (rr.low || rr.high)) {
        const low = rr.low?.value !== undefined ? `${rr.low.value}` : "";
        const high = rr.high?.value !== undefined ? `${rr.high.value}` : "";
        const unit = rr.low?.unit || rr.high?.unit || "";
        rangeStr = `${low} - ${high}${unit ? ` ${unit}` : ""}`;
      }
      if (rangeStr) fields.push({ label: "Reference Range", value: rangeStr });
    }

    const method = extractCodeDisplay(resource.method);
    if (method) fields.push({ label: "Method", value: method });

    const notes = resource.note as { text?: string }[] | undefined;
    if (notes?.length) {
      fields.push({ label: "Notes", value: notes.map((n) => n.text).filter(Boolean).join("; ") });
    }
  }

  // Fallback for unrecognized or sparse resources
  if (fields.length === 0) {
    if (resource.status) fields.push({ label: "Status", value: String(resource.status) });
    const code = extractCodeDisplay(resource.code);
    if (code) fields.push({ label: "Code", value: code });
    if (fields.length === 0) {
      fields.push({ label: "Resource", value: `${resourceType}/${resource.id}` });
    }
  }

  return (
    <div className="detail-grid">
      {fields.map((f, i) => (
        <div key={i} className="detail-row">
          <span className="detail-label">{f.label}</span>
          <span className="detail-value">{f.value}</span>
        </div>
      ))}
    </div>
  );
}

// ---------------------------------------------------------------------------
// Vitals Trend Components
// ---------------------------------------------------------------------------

function VitalsTrendCard({
  app,
  patientId,
  setActiveToolCall,
  cardsLoaded,
}: {
  app: App;
  patientId: string;
  setActiveToolCall: (call: string | null) => void;
  cardsLoaded: boolean;
}) {
  const [range, setRange] = useState<TrendRange>("3m");
  const [trendData, setTrendData] = useState<TrendData>({
    loading: true,
    error: null,
    bp: [],
    weight: [],
  });

  const fetchTrends = useCallback(
    async (r: TrendRange) => {
      setTrendData({ loading: true, error: null, bp: [], weight: [] });
      setActiveToolCall("fhir_search vital-signs trends");
      try {
        const startDate = getStartDate(r);
        const result = await app.callServerTool({
          name: "fhir_search",
          arguments: {
            resourceType: "Observation",
            queryParams: `patient=Patient/${patientId}&category=vital-signs&date=ge${startDate}&_sort=-date&_count=50`,
          },
        });
        if (result.isError) {
          const msg = extractResultData(result);
          setTrendData({
            loading: false,
            error: typeof msg === "string" ? msg : JSON.stringify(msg),
            bp: [],
            weight: [],
          });
        } else {
          const data = extractResultData(result);
          const entries = extractBundleEntries(data);
          setTrendData({
            loading: false,
            error: null,
            bp: extractBPPoints(entries),
            weight: extractWeightPoints(entries),
          });
        }
      } catch (e) {
        setTrendData({
          loading: false,
          error: e instanceof Error ? e.message : String(e),
          bp: [],
          weight: [],
        });
      } finally {
        setActiveToolCall(null);
      }
    },
    [app, patientId, setActiveToolCall],
  );

  useEffect(() => {
    if (!cardsLoaded) return;
    fetchTrends(range);
  }, [range, fetchTrends, cardsLoaded]);

  return (
    <div className="card trend-card">
      <div className="card-header">
        <div className="card-icon vitals">{"\uD83D\uDCC8"}</div>
        <span className="card-title vitals">Vitals Trends</span>
        <div className="trend-range-picker">
          {TREND_RANGES.map((tr) => (
            <button
              key={tr.key}
              className={`trend-range-btn${range === tr.key ? " active" : ""}`}
              onClick={() => setRange(tr.key)}
              disabled={trendData.loading}
            >
              {tr.label}
            </button>
          ))}
        </div>
        <button
          className="btn-refresh"
          onClick={() => fetchTrends(range)}
          title="Refresh"
          disabled={trendData.loading}
        >
          {"\u21BB"}
        </button>
      </div>
      {trendData.loading && (
        <div className="card-loading">
          <div className="spinner" />
          Loading trends...
        </div>
      )}
      {trendData.error && (
        <div className="card-error">
          {"\u26A0\uFE0F"} {trendData.error}
        </div>
      )}
      {!trendData.loading && !trendData.error && (
        <div className="trend-charts">
          <TrendChart title="Blood Pressure" unit="mmHg" empty={trendData.bp.length === 0} emptyMessage="No blood pressure data">
            {trendData.bp.length === 1 ? (
              <div className="trend-single-point">
                <span>{trendData.bp[0].date}</span>
                <span>{trendData.bp[0].systolic}/{trendData.bp[0].diastolic} mmHg</span>
              </div>
            ) : (
              <BPChart points={trendData.bp} />
            )}
          </TrendChart>
          <TrendChart title="Weight" unit="" empty={trendData.weight.length === 0} emptyMessage="No weight data">
            {trendData.weight.length === 1 ? (
              <div className="trend-single-point">
                <span>{trendData.weight[0].date}</span>
                <span>{trendData.weight[0].value} {trendData.weight[0].unit}</span>
              </div>
            ) : (
              <WeightChart points={trendData.weight} />
            )}
          </TrendChart>
        </div>
      )}
    </div>
  );
}

function TrendChart({
  title,
  unit,
  empty,
  emptyMessage,
  children,
}: {
  title: string;
  unit: string;
  empty: boolean;
  emptyMessage: string;
  children: React.ReactNode;
}) {
  return (
    <div className="trend-chart-panel">
      <div className="trend-chart-title">
        {title}
        {unit && <span className="trend-chart-unit">({unit})</span>}
      </div>
      {empty ? (
        <div className="trend-chart-empty">{emptyMessage}</div>
      ) : (
        children
      )}
    </div>
  );
}

function BPChart({ points }: { points: BPPoint[] }) {
  const W = 320;
  const H = 180;
  const PAD = { top: 16, right: 12, bottom: 28, left: 36 };

  const allVals = points.flatMap((p) => [p.systolic, p.diastolic]);
  const yMin = Math.floor(Math.min(...allVals) / 10) * 10 - 10;
  const yMax = Math.ceil(Math.max(...allVals) / 10) * 10 + 10;
  const ts = points.map((p) => p.timestamp);
  const sysPoints = computePolylinePoints(points.map((p) => p.systolic), ts, W, H, PAD, yMin, yMax);
  const diaPoints = computePolylinePoints(points.map((p) => p.diastolic), ts, W, H, PAD, yMin, yMax);

  const plotH = H - PAD.top - PAD.bottom;
  const yRange = yMax - yMin || 1;
  const gridLines = [];
  const step = yRange <= 60 ? 10 : 20;
  for (let v = Math.ceil(yMin / step) * step; v <= yMax; v += step) {
    const y = PAD.top + plotH - ((v - yMin) / yRange) * plotH;
    gridLines.push({ y, label: String(v) });
  }

  const tMin = ts[0];
  const tMax = ts[ts.length - 1];
  const tRange = tMax - tMin || 1;
  const plotW = W - PAD.left - PAD.right;
  const dateLabels: { x: number; label: string }[] = [];
  const labelCount = Math.min(points.length, 5);
  const labelStep = Math.max(1, Math.floor((points.length - 1) / (labelCount - 1)));
  for (let i = 0; i < points.length; i += labelStep) {
    const x = PAD.left + ((ts[i] - tMin) / tRange) * plotW;
    dateLabels.push({ x, label: points[i].date.slice(5) });
  }
  if (points.length > 1) {
    const lastI = points.length - 1;
    const x = PAD.left + ((ts[lastI] - tMin) / tRange) * plotW;
    if (!dateLabels.find((d) => Math.abs(d.x - x) < 30)) {
      dateLabels.push({ x, label: points[lastI].date.slice(5) });
    }
  }

  const dotCoords = (vals: number[]) =>
    vals.map((v, i) => ({
      cx: PAD.left + ((ts[i] - tMin) / tRange) * plotW,
      cy: PAD.top + plotH - ((v - yMin) / yRange) * plotH,
    }));

  const sysDots = dotCoords(points.map((p) => p.systolic));
  const diaDots = dotCoords(points.map((p) => p.diastolic));

  return (
    <div className="trend-svg-container">
      <svg className="trend-svg" viewBox={`0 0 ${W} ${H}`} preserveAspectRatio="xMidYMid meet">
        {gridLines.map((g, i) => (
          <g key={i}>
            <line x1={PAD.left} y1={g.y} x2={W - PAD.right} y2={g.y} stroke="var(--color-border-primary)" strokeWidth="0.5" />
            <text x={PAD.left - 4} y={g.y + 3} textAnchor="end" fontSize="8" fill="var(--color-text-secondary)">{g.label}</text>
          </g>
        ))}
        {dateLabels.map((d, i) => (
          <text key={i} x={d.x} y={H - 6} textAnchor="middle" fontSize="7" fill="var(--color-text-secondary)">{d.label}</text>
        ))}
        <polyline points={sysPoints} fill="none" stroke="#ef4444" strokeWidth="1.5" strokeLinejoin="round" />
        <polyline points={diaPoints} fill="none" stroke="#3b82f6" strokeWidth="1.5" strokeLinejoin="round" />
        {sysDots.map((d, i) => (
          <circle key={`s${i}`} cx={d.cx} cy={d.cy} r="2.5" fill="#ef4444" />
        ))}
        {diaDots.map((d, i) => (
          <circle key={`d${i}`} cx={d.cx} cy={d.cy} r="2.5" fill="#3b82f6" />
        ))}
      </svg>
      <div className="trend-legend">
        <span className="trend-legend-item"><span className="trend-legend-swatch" style={{ background: "#ef4444" }} />Systolic</span>
        <span className="trend-legend-item"><span className="trend-legend-swatch" style={{ background: "#3b82f6" }} />Diastolic</span>
      </div>
    </div>
  );
}

function WeightChart({ points }: { points: WeightPoint[] }) {
  const W = 320;
  const H = 180;
  const PAD = { top: 16, right: 12, bottom: 28, left: 36 };

  const vals = points.map((p) => p.value);
  const rawMin = Math.min(...vals);
  const rawMax = Math.max(...vals);
  const margin = (rawMax - rawMin) * 0.15 || 2;
  const yMin = Math.floor((rawMin - margin) * 10) / 10;
  const yMax = Math.ceil((rawMax + margin) * 10) / 10;
  const ts = points.map((p) => p.timestamp);
  const linePoints = computePolylinePoints(vals, ts, W, H, PAD, yMin, yMax);

  const plotH = H - PAD.top - PAD.bottom;
  const plotW = W - PAD.left - PAD.right;
  const yRange = yMax - yMin || 1;

  // Area fill: line points + close to bottom
  const tMin = ts[0];
  const tRange = (ts[ts.length - 1] - tMin) || 1;
  const firstX = PAD.left;
  const lastX = PAD.left + plotW;
  const bottomY = PAD.top + plotH;
  const areaPoints = `${firstX},${bottomY} ${linePoints} ${lastX},${bottomY}`;

  const gridLines: { y: number; label: string }[] = [];
  const range = yMax - yMin;
  const step = range <= 10 ? 2 : range <= 30 ? 5 : 10;
  for (let v = Math.ceil(yMin / step) * step; v <= yMax; v += step) {
    const y = PAD.top + plotH - ((v - yMin) / yRange) * plotH;
    gridLines.push({ y, label: v % 1 === 0 ? String(v) : v.toFixed(1) });
  }

  const dateLabels: { x: number; label: string }[] = [];
  const labelCount = Math.min(points.length, 5);
  const labelStep = Math.max(1, Math.floor((points.length - 1) / (labelCount - 1)));
  for (let i = 0; i < points.length; i += labelStep) {
    const x = PAD.left + ((ts[i] - tMin) / tRange) * plotW;
    dateLabels.push({ x, label: points[i].date.slice(5) });
  }
  if (points.length > 1) {
    const lastI = points.length - 1;
    const x = PAD.left + ((ts[lastI] - tMin) / tRange) * plotW;
    if (!dateLabels.find((d) => Math.abs(d.x - x) < 30)) {
      dateLabels.push({ x, label: points[lastI].date.slice(5) });
    }
  }

  const dots = vals.map((v, i) => ({
    cx: PAD.left + ((ts[i] - tMin) / tRange) * plotW,
    cy: PAD.top + plotH - ((v - yMin) / yRange) * plotH,
  }));

  const unit = points[0]?.unit || "kg";

  return (
    <div className="trend-svg-container">
      <svg className="trend-svg" viewBox={`0 0 ${W} ${H}`} preserveAspectRatio="xMidYMid meet">
        {gridLines.map((g, i) => (
          <g key={i}>
            <line x1={PAD.left} y1={g.y} x2={W - PAD.right} y2={g.y} stroke="var(--color-border-primary)" strokeWidth="0.5" />
            <text x={PAD.left - 4} y={g.y + 3} textAnchor="end" fontSize="8" fill="var(--color-text-secondary)">{g.label}</text>
          </g>
        ))}
        {dateLabels.map((d, i) => (
          <text key={i} x={d.x} y={H - 6} textAnchor="middle" fontSize="7" fill="var(--color-text-secondary)">{d.label}</text>
        ))}
        <polygon points={areaPoints} fill="rgba(16,185,129,0.12)" />
        <polyline points={linePoints} fill="none" stroke="#10b981" strokeWidth="1.5" strokeLinejoin="round" />
        {dots.map((d, i) => (
          <circle key={i} cx={d.cx} cy={d.cy} r="2.5" fill="#10b981" />
        ))}
      </svg>
      <div className="trend-legend">
        <span className="trend-legend-item"><span className="trend-legend-swatch" style={{ background: "#10b981" }} />{unit}</span>
      </div>
    </div>
  );
}

// ---------------------------------------------------------------------------
// Card item renderers
// ---------------------------------------------------------------------------

function renderCondition(item: unknown, index: number) {
  const r = item as Record<string, unknown>;
  const display = extractCodeDisplay(r.code);
  const cs = r.clinicalStatus as
    | { coding?: { code?: string }[] }
    | undefined;
  const status = cs?.coding?.[0]?.code || "";
  const onset = formatDate(r.onsetDateTime);

  return (
    <div key={index} className="item-row">
      <span className="item-icon">{"\u25CF"}</span>
      <span className="item-primary">{display || "Unknown condition"}</span>
      {onset && <span className="item-secondary">{onset}</span>}
      {status && <StatusBadge status={status} />}
    </div>
  );
}

function renderMedication(item: unknown, index: number) {
  const r = item as Record<string, unknown>;
  const display =
    extractCodeDisplay(r.medicationCodeableConcept) ||
    extractCodeDisplay(r.medicationReference);
  const status = (r.status as string) || "";
  const authored = formatDate(r.authoredOn);

  return (
    <div key={index} className="item-row">
      <span className="item-icon">{"\uD83D\uDC8A"}</span>
      <span className="item-primary">{display || "Unknown medication"}</span>
      {authored && <span className="item-secondary">{authored}</span>}
      {status && <StatusBadge status={status} />}
    </div>
  );
}

function renderVital(item: unknown, index: number) {
  const r = item as Record<string, unknown>;
  const display = extractCodeDisplay(r.code);
  const value = formatObservationValue(r);
  const date = formatDate(r.effectiveDateTime || r.issued);

  return (
    <div key={index} className="vital-row">
      <div className="vital-header">
        <span className="vital-name">
          <span className="vital-icon">{"\u2764"}</span>
          {display || "Vital"}
        </span>
        {date && <span className="vital-date">{date}</span>}
      </div>
      {value && <span className="vital-value">{value}</span>}
    </div>
  );
}

function renderLab(item: unknown, index: number) {
  const r = item as Record<string, unknown>;
  const display = extractCodeDisplay(r.code);
  const value = formatObservationValue(r);
  const date = formatDate(r.effectiveDateTime || r.issued);

  return (
    <div key={index} className="vital-row">
      <div className="vital-header">
        <span className="vital-name">
          <span className="vital-icon">{"\uD83E\uDDEA"}</span>
          {display || "Lab"}
        </span>
        {date && <span className="vital-date">{date}</span>}
      </div>
      {value && <span className="vital-value">{value}</span>}
    </div>
  );
}

const CARD_RENDERERS: Record<CardKey, (item: unknown, index: number) => React.ReactNode> = {
  conditions: renderCondition,
  medications: renderMedication,
  vitals: renderVital,
  labs: renderLab,
};

// ---------------------------------------------------------------------------
// Observation value formatter
// ---------------------------------------------------------------------------

function formatObservationValue(r: Record<string, unknown>): string {
  const vq = r.valueQuantity as
    | { value?: number; unit?: string }
    | undefined;
  if (vq?.value !== undefined) {
    return `${vq.value}${vq.unit ? ` ${vq.unit}` : ""}`;
  }

  if (typeof r.valueString === "string") return r.valueString;

  const vcc = r.valueCodeableConcept as
    | { text?: string; coding?: { display?: string }[] }
    | undefined;
  if (vcc) return vcc.text || vcc.coding?.[0]?.display || "";

  const components = r.component as
    | { code?: unknown; valueQuantity?: { value?: number; unit?: string } }[]
    | undefined;
  if (components && components.length > 0) {
    return components
      .map((c) => {
        const name = extractCodeDisplay(c.code);
        const val = c.valueQuantity;
        if (val?.value !== undefined) {
          return `${name ? name + ": " : ""}${val.value}${val.unit ? " " + val.unit : ""}`;
        }
        return "";
      })
      .filter(Boolean)
      .join(" / ");
  }

  return "";
}

// ---------------------------------------------------------------------------
// Badges
// ---------------------------------------------------------------------------

function StatusBadge({ status }: { status: string }) {
  const lower = status.toLowerCase();
  let cls = "badge";
  if (lower === "active" || lower === "finished" || lower === "completed") {
    cls += " badge-active";
  } else {
    cls += " badge-inactive";
  }
  return (
    <span className={cls}>
      <span className="badge-dot" />
      {status}
    </span>
  );
}

// ---------------------------------------------------------------------------
// Mount
// ---------------------------------------------------------------------------

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ChartReviewRoot />
  </StrictMode>
);
