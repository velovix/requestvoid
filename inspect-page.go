package main

import "html/template"

var inspectPage = `
<html>
	<head>
		<title>RequestVoid</title>

		<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
	</head>

	<body>
		{{ range . }}
			<strong>Received</strong>: {{ printf "%s" .TimeSinceReceived }}
			<br />
			<pre style="white-space:normal;">
{{ .Body }}
			</pre>
			<hr>
		{{ end }}
	</body>
</html>
`
var inspectPageTemplate *template.Template

func init() {
	inspectPageTemplate = template.Must(template.New("").Parse(inspectPage))
}
