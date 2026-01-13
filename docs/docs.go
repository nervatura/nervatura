package docs

import "embed"

//go:embed index.html 404.html open css grpc cli images model install start examples screenshots mcp client editor favicon.svg upgrade
var Docs embed.FS
