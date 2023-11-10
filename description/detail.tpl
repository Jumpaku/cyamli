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
    {{- range $Index, $Option := .Options -}}
        {{"\n    "}}
        {{- range $Index, $Name := $Option.Names -}}
            {{- if $Index}}, {{end -}}
            {{$Name}}
            {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
            {{- else -}}=<{{$Option.Type}}>
            {{- end -}}
        {{- end -}}{{"  "}}(default={{$Option.Default}}):
        {{- range $Index, $Line := $Option.Description -}}
            {{"\n        "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}


{{- if .HasArguments -}}
{{"\n"}}Arguments:
    {{- range $Index, $Argument := .Arguments -}}
        {{"\n    "}}
        {{- if $Argument.Variadic -}}[{{$Index}}:] [<{{$Argument.Name}}:{{$Argument.Type}}>]...
        {{- else -}}[{{$Index}}]  <{{$Argument.Name}}:{{$Argument.Type}}>
        {{- end -}}
        {{- range $Index, $Line := $Argument.Description -}}
            {{"\n        "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}


{{- if .HasSubcommands -}}
{{"\n"}}Subcommands:
    {{- range $Index, $Subcommand := .Subcommands -}}
        {{"\n    "}}{{$Subcommand.Name}}:
        {{- range $Index, $Line := $Subcommand.Description -}}
            {{"\n        "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}