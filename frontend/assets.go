package frontend

import "embed"

//go:embed all:dist
var WebAssets embed.FS
