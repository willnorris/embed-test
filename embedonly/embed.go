package embedonly

import "embed"

//go:embed *.go
var _ embed.FS
