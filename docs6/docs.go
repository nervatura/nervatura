package docs

import "embed"

//go:embed index.html open css grpc cli images model install create examples screenshots mcp web favicon.svg upgrade
var Docs embed.FS
