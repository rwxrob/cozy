package cozy

import _ "embed"

// Keep all documentation in this file as embeds in order to easily
// support compiles to other target languages by simply changing the
// language identifier before compilation.

//go:embed text/en/cozy.md
var _cozy string
