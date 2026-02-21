package apps

import "embed"

// AppBundles embeds all built MCP App HTML bundles.
// These are single-file HTML bundles produced by Vite with all JS/CSS inlined.
// Uses dist/* so go:embed works on clean checkout when only .gitkeep.html exists.
//
//go:embed dist/*
var AppBundles embed.FS
