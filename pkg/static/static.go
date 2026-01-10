package static

import "embed"

//go:embed public
var Public embed.FS

//go:embed store/mysql store/postgres store/sqlite store/mssql upgrade/mysql upgrade/postgres upgrade/sqlite
var Store embed.FS

//go:embed message.json template mcp
var Static embed.FS

//go:embed template
var Report embed.FS

const (
	ApiPath        = "/api/v6"
	ClientPath     = "/client"
	DocsClientPath = "/docs6/client"
	AdminPath      = "/admin"
	DocsPath       = "/docs6"

	ClientAuthRedirectURL = "/client/api/auth/callback"
	OAuthAuthRedirectURL  = "/oauth/callback"
	BrowserRowLimit       = 100

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
		"allow_headers":   "Accept,Authorization,Content-Type,X-CSRF-Token,X-Api-Key,x-payload-digest,Stripe-Signature,Mcp-Session-Id,Mcp-Protocol-Version",
		"expose_headers":  "",
		"trusted_origins": "http://localhost:5000,http://localhost:5500",
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
	"mcp": {
		"enabled": "true",
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
		"alg":            "HS256",
		"user":           "user_name",
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

const McpPage = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover" />
		<title>{{ .title }}</title>
		<link rel="icon" type="image/svg+xml" href="/static/favicon.svg">
		<link rel="stylesheet" href="/static/css/index.css" />
		<link rel="stylesheet" href="/public/css/task.css" />
	</head>
	<body>
	<div class="container">
	{{if .tools}}
		<h3>Tools</h3>
		<ul>
		{{range $scope, $names := .tools}}
			<li><b>/mcp/{{ $scope }}</b>
			<ul>
			{{range $name, $values := $names}}
				<li><b>{{ $name }}</b>: <i>{{ $values.description }}</i>
				{{ if eq $scope "all" }}
				{{ range $values.scopes }}<span style="color:rgb(var(--functional-green))">{{ . }}</span> {{ end }}
				{{ end}}
				</li>
			{{end}}
			</ul>
			</li>
		{{end}}
		</ul>
	{{end}}
	</div>
	</body>
</html>`

const AuthPage = `<!doctype html>
<html lang="en" class="h-full bg-gray-900">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{ .title }}</title>
	<link rel="icon" type="image/svg+xml" href="/static/favicon.svg">
  <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
</head>
<body class="h-full">
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <img src="/public/images/logo.svg" alt="Nervatura" class="mx-auto h-24 w-auto" />
      <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight text-white">{{ .subtitle }}</h2>
    </div>
    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
		  {{if eq .code ""}}
      <form action="/oauth/authorization" method="POST" class="space-y-6">
				<input type="hidden" name="session_id" value="{{ .session_id }}" />
        <div>
          <label for="username" class="block text-sm/6 font-medium text-gray-100">{{ .username_label }}</label>
          <div class="mt-2">
            <input id="username" type="username" name="username" value="{{ .username }}" required autofocus autocomplete="username" class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6" />
          </div>
        </div>
        <div>
          <div class="flex items-center justify-between">
            <label for="password" class="block text-sm/6 font-medium text-gray-100">{{ .password_label }}</label>
          </div>
          <div class="mt-2">
            <input id="password" type="password" name="password" required autocomplete="current-password" class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6" />
          </div>
        </div>
				{{if .error_msg}}
				<div class="mt-2">
					<div class="flex items-center justify-between">
						<label class="block font-bold text-sm/6 font-medium text-red-500">{{ .error_msg }}</label>
					</div>
				</div>
				{{end}}
        <div>
          <button type="submit" class="cursor-pointer flex w-full justify-center rounded-md bg-indigo-500 px-3 py-1.5 text-sm/6 font-semibold text-white hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">{{ .login_button }}</button>
        </div>
      </form>
			{{else}}
			<div class="space-y-6">
				<div>
					<div class="flex items-center justify-between">
						<label class="block font-bold text-sm/6 font-medium text-gray-100">{{ .code_label }}</label>
					</div>
					<div>
						<label class="block text-sm/6 font-medium text-green-500">{{ .code }}</label>
					</div>
				</div>
				<div>
					<div class="flex items-center justify-between">
						<label class="block font-bold text-sm/6 font-medium text-gray-100">{{ .state_label }}</label>
					</div>
					<div>
						<label class="block text-sm/6 font-medium text-orange-500">{{ .state }}</label>
					</div>
				</div>
			</div>
			{{end}}
    </div>
  </div>
</body>
</html>`
