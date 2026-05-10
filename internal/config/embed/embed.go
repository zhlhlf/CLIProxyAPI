// Package embed provides embedded configuration files for the CLI Proxy API.
// This file is intentionally left empty - it exists only to allow the embed directive.
package embed

import (
	_ "embed"
)

//go:embed config.yaml
var Config []byte
