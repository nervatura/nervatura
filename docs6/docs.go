package docs

import "embed"

//go:embed index.html open css grpc cli images model install create examples screenshots mcp client favicon.svg upgrade
var Docs embed.FS
