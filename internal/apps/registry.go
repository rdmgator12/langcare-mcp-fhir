package apps

import (
	"fmt"
	"io/fs"
)

// AppConfig defines an MCP App with its own dedicated tool.
type AppConfig struct {
	Name        string // e.g., "fhir-explorer"
	ToolName    string // Dedicated tool name, e.g., "fhir_explorer"
	ToolDesc    string // Tool description for LLM
	ResourceURI string // e.g., "ui://fhir-explorer/app.html"
	Description string // Human-readable resource description
	HTMLFile    string // Filename in dist/, e.g., "fhir-explorer.html"
}

// DefaultApps is the list of built-in MCP Apps.
// Each app has its own dedicated tool — the MCP Apps model is one tool → one UI resource.
var DefaultApps = []AppConfig{
	{
		Name:        "fhir-explorer",
		ToolName:    "fhir_explorer",
		ToolDesc:    "Interactive FHIR resource browser. Opens a UI to search, read, create, and update any FHIR R4 resource type. Use this when the user wants to explore or browse FHIR data interactively.",
		ResourceURI: "ui://fhir-explorer/app.html",
		Description: "Interactive FHIR resource browser — search, read, and inspect any FHIR R4 resource type.",
		HTMLFile:    "fhir-explorer.html",
	},
	{
		Name:        "patient-chart-review",
		ToolName:    "patient_chart_review",
		ToolDesc:    "Patient chart review dashboard. Opens a UI showing demographics, active conditions, medications, allergies, vitals, labs, and recent encounters. Use this when the user wants to review a patient's medical record.",
		ResourceURI: "ui://patient-chart-review/app.html",
		Description: "Patient chart review — demographics, conditions, medications, allergies, vitals, and encounters.",
		HTMLFile:    "patient-chart-review.html",
	},
}

// LoadAppHTML reads the embedded HTML bundle for a given app.
func LoadAppHTML(appName string) (string, error) {
	for _, app := range DefaultApps {
		if app.Name == appName {
			data, err := fs.ReadFile(AppBundles, "dist/"+app.HTMLFile)
			if err != nil {
				return "", fmt.Errorf("failed to read app bundle %s: %w", app.HTMLFile, err)
			}
			return string(data), nil
		}
	}
	return "", fmt.Errorf("app not found: %s", appName)
}
