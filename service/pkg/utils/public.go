//go:build http || all
// +build http all

package utils

import "embed"

const ClientMsg = "static/locales/client.json"

//go:embed static/client static/css static/js
var Public embed.FS

//go:embed static/locales
var Locales embed.FS
