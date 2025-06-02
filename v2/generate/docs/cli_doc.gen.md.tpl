{{- /* Go Template */ -}}
# {{.Program.Name}}{{ if .Program.Version }}{{" "}}({{.Program.Version}}){{ end }}{{"\n\n"}}

{{- range .CommandList}}
{{.DocTextMarkdown}}
{{ end }}
