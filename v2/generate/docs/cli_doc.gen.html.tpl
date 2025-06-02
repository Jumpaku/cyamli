{{- /* Go Template */ -}}

<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Program.Name}}{{ if .Program.Version }}{{" "}}({{.Program.Version}}){{ end }}</title>
</head>
<body>
<main>
    <h1>{{.Program.Name}}{{ if .Program.Version }}{{" "}}({{.Program.Version}}){{ end }}</h1>
    {{ range .CommandList }}
        {{.DocTextHTML}}
    {{ end }}
</main>
</body>
</html>
