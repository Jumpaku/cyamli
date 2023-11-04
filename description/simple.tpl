{{- /* Go Template */ -}}
{{- if .Program -}}
    {{.Program}}
    {{- range .Path -}}{{" "}}{{.}}{{- end -}}:{{"\n"}}
{{- end -}}
{{- if .Description -}}
    {{- range $Index, $Line := .Description -}}
        {{- if $Index -}}{{"\n"}}{{- end -}}{{$Line}}
    {{- end -}}
    {{"\n"}}
{{- end -}}


{{"\n"}}Usage:
    $ {{.Syntax}}{{"\n"}}


{{- if .HasOptions -}}
{{"\n"}}Options:
{{"    "}}
    {{- range $Index, $Option := .Options -}}
        {{- if $Index }}, {{end -}}{{ index $Option.Names 0 }}
    {{- end -}}
{{"\n"}}
{{- end -}}


{{- if .HasArguments -}}
{{"\n"}}Arguments:
{{"    "}}
    {{- range $Index, $Argument := .Arguments -}}
        {{- if $Index }} {{end -}}
        {{- if $Argument.Variadic -}}<{{$Argument.Name}}>...
        {{- else -}}<{{$Argument.Name}}>
        {{- end -}}
    {{- end -}}
{{"\n"}}
{{- end -}}


{{- if .HasSubcommands -}}
{{"\n"}}Subcommands:
{{"    "}}
    {{- range $Index, $Subcommand := .Subcommands -}}
        {{- if $Index }}, {{end -}}{{$Subcommand.Name}}
    {{- end -}}
{{"\n"}}
{{- end -}}