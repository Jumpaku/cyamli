{{- /* Go Template */ -}}
{{.Program.Name}}{{ if .Program.Version }}{{" "}}({{.Program.Version}}){{ end }}{{"\n"}}

{{ range .CommandList }}
    {{.DocText}}
{{ end }}
