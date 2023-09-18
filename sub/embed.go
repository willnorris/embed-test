package embedtest

import "embed"

//go:embed sub-linked.*
var EmbeddedFS embed.FS
