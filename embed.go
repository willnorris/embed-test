package embedtest

import "embed"

//go:embed root-linked.* sub/sub-rootlinked.*
var embeddedFS embed.FS
