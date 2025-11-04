package static

import "embed"

//go:embed public
var Public embed.FS

//go:embed store/mysql store/postgres store/sqlite upgrade/mysql upgrade/postgres upgrade/sqlite
var Store embed.FS

//go:embed message.json
var Static embed.FS

//go:embed template
var Report embed.FS

const (
	ApiPath        = "/api/v6"
	ClientPath     = "/client"
	DocsClientPath = "/docs6/client"
	AdminPath      = "/admin"
	DocsPath       = "/docs6"

	AuthRedirectURL = "http://%s/api/auth/callback"
	BrowserRowLimit = 100

	DefaultModule      = "search"
	DefaultSearchView  = "transitem_simple"
	DefaultLang        = "en"
	DefaultTheme       = "light"
	DefaultOrientation = "P"
	DefaultPaperSize   = "a4"
	DefaultPagination  = "10"
	DefaultHistory     = 5
	DefaultExportSep   = ","
)

var ClientLang []string = []string{"en,English"}
var ClientTheme []string = []string{"light", "dark"}

var DefaultConfig map[string]map[string]string = map[string]map[string]string{
	"cors": {
		"allow_origins":   "*",
		"allow_methods":   "GET,POST,PUT,DELETE,OPTIONS",
		"allow_headers":   "Accept,Authorization,Content-Type,X-CSRF-Token,X-Api-Key,x-payload-digest,Stripe-Signature",
		"expose_headers":  "",
		"trusted_origins": "localhost:5000",
	},
	"http": {
		"http_enabled":  "true",
		"port":          "5000",
		"tls_enabled":   "false",
		"read_timeout":  "30",
		"write_timeout": "30",
		"home":          "/client",
	},
	"grpc": {
		"enabled":     "true",
		"port":        "9200",
		"tls_enabled": "false",
	},
	"report": {
		"font_family": "",
		"font_dir":    "",
		"dir":         "",
	},
	"token": {
		"iss":            "nervatura",
		"exp":            "6",
		"public_key_url": "",
	},
	"smtp": {
		"host":            "",
		"port":            "465",
		"tls_min_version": "0",
		"user":            "",
	},
	"dev": {
		"hide_error": "false",
		"update_log": "true",
	},
	"session": {
		"table":    "session",
		"file_dir": "",
		"exp":      "1",
		"method":   "0",
	},
	"sql": {
		"max_open_conns":    "10",
		"max_idle_conns":    "3",
		"conn_max_lifetime": "15",
	},
	"connection": {
		"auth_callback": "http://%s/api/auth/callback",
		"tls_cert_file": "",
		"tls_key_file":  "",
		"default_alias": "demo",
		"default_admin": "USR0000000000N1",
	},
	"update_log": {
		// "table_name": "INSERT,UPDATE,DELETE",
	},
}

const ErrorTemplate = `<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover" />
			<title>{{ .title }}</title>
			<link rel="icon" type="image/svg+xml" href="/static/favicon.svg">
			<link rel="stylesheet" href="/public/css/app.css" />
			<link rel="stylesheet" href="/static/css/index.css" />
			<link rel="preconnect" href="https://rsms.me/">
			<link rel="stylesheet" href="https://rsms.me/inter/inter.css">
		</head>
		<body>
		  <div class="login-modal">
			<div class="middle"><div class="dialog">
			<div class="row title">
	    <div class="cell title-cell login-title-cell" >{{ .error_title }}</div>
	    </div>
			<div class="row full section" >
			<div class="cell container center bold" style="color: red;" >{{ .error_msg }}</div>
			</div>
			<div class="row full section buttons" >
			<div class="cell container center" >
			<a href="/" link-type="primary" >{{ .login }}</a>
			</div>
			</div>
			</div></div></div>
		</body>
	</html>`

const TaskPage = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover" />
		<title>{{ .title }}</title>
		<link rel="icon" type="image/svg+xml" href="/static/favicon.svg">
		<link rel="stylesheet" href="/static/css/index.css" />
		<link rel="stylesheet" href="/public/css/task.css" />
	</head>
	<body><div class="task-container row full mobile" theme="dark" style="margin:auto;">
	{{if .env_result}}
		<div class="container section">
		{{range .env_result}}
			<div class="row full border-top" >
				<div class="cell mobile small">
					<div class="cell padding-normal bold" style="white-space:nowrap;vertical-align: top;" >{{ .envkey }}</div>
				<div class="cell mobile padding-normal" >
				{{if .envvalue}}<span style="color:rgb(var(--functional-green));white-space:wrap;">{{ .envvalue }}</span>
				{{else}}<span style="color:rgb(var(--functional-red));">X</span>{{end}}
				</div>
			</div>
		</div>
		{{end}}
	</div>
	{{end}}
	</div></body>
</html>`
