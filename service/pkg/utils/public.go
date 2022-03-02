//go:build http || all
// +build http all

package utils

import "embed"

//go:embed static/client static/css
var Public embed.FS
