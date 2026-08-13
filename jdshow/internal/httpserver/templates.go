package httpserver

import "embed"

//go:embed templates/*.html
var templateFiles embed.FS
