package templates

import "embed"

//go:embed *.gohtml admin/*.gohtml
var FS embed.FS
