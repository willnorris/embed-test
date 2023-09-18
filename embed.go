package embedtest

import "embed"

//- go:embed root-linked.* sub/sub-rootlinked.*

//go:embed root-linked.*
var embeddedFS embed.FS

// - go:embed **/*.css **/*.tsx **/*.jsx
var _ embed.FS
