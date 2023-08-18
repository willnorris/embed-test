package embedtest

import "embed"

//go:embed sub-linked.*
var embeddedFS embed.FS
