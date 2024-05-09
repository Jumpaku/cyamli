{{- /* Go Template */ -}}
{{- $Program := .Program -}}
{{- $Version := .Version -}}

{{- $Program -}}{{- if $Version -}}{{" "}}({{$Version}}){{- end -}}{{"\n\n"}}


{{- range .Commands -}}


{{$Program}}{{- range .Path -}}{{" "}}{{.}}{{- end -}}{{"\n"}}


{{- if .Description -}}
    {{"\n    "}}Description:{{"\n"}}
    {{- range $Index, $Line := .Description -}}
        {{- if $Index -}}{{"\n"}}{{- end -}}{{"        "}}{{$Line}}
    {{- end -}}
    {{"\n"}}
{{- end -}}


{{"\n    "}}Syntax:
{{"        "}}$ {{$Program}}{{" "}}{{.Syntax}}{{"\n"}}


{{- if .Options -}}
    {{"\n    "}}Options:
    {{- range $Index, $Option := .Options -}}
        {{"\n        "}}
        {{- range $Index, $Name := $Option.Names -}}
            {{- if $Index}}, {{end -}}
            {{$Name}}
            {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
            {{- else -}}=<{{$Option.Type}}>
            {{- end -}}
        {{- end -}}{{"  "}}(default={{$Option.Default}}):
        {{- range $Index, $Line := $Option.Description -}}
            {{"\n            "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}


{{- if .Arguments -}}
    {{"\n    "}}Arguments:
    {{- range $Index, $Argument := .Arguments -}}
        {{"\n        "}}
        {{- if $Argument.Variadic -}}
            {{func_add $Index 1}}. [<{{$Argument.Name}}{{":"}}{{$Argument.Type}}>]...
        {{- else -}}
            {{func_add $Index 1}}.  <{{$Argument.Name}}{{":"}}{{$Argument.Type}}>
        {{- end -}}
        {{- range $Index, $Line := $Argument.Description -}}
            {{"\n            "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}


{{- if .Subcommands -}}
    {{"\n    "}}Subcommands:
    {{- range $Index, $Subcommand := .Subcommands -}}
        {{"\n        "}}{{$Subcommand.Name}}:
        {{- range $Index, $Line := $Subcommand.Description -}}
            {{"\n            "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}
{{"\n\n"}}


{{- end -}}
