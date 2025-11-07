{{- /* Go Template */ -}}
{{.Path}}{{"\n"}}

{{- if .DescriptionLines -}}
    {{"\n    "}}Description:{{"\n"}}
    {{- range $Index, $Line := .DescriptionLines -}}
        {{- if $Index -}}{{"\n"}}{{- end -}}{{"        "}}{{$Line}}
    {{- end -}}
    {{"\n"}}
{{- end -}}

{{"\n    "}}Syntax:
{{"        "}}$ {{.Syntax}}{{"\n"}}

{{- if .Options -}}
    {{"\n    "}}Options:
    {{- range $Index, $Option := .Options -}}
        {{"\n        "}}
        {{- range $Index, $Name := $Option.Options -}}
            {{- if $Index}}, {{end -}}
            {{$Name}}
            {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
            {{- else -}}=<{{$Option.Type}}>
            {{- end -}}{{- if $Option.Repeated}} ... {{end -}}
        {{- end -}}{{- if not $Option.Repeated -}}  (default={{$Option.Default}}){{end -}}
        {{- if $Option.Negation -}}
            {{",\n        "}}{{$Option.NegatedOption}}[=<{{$Option.Type}}>]{{- if $Option.Repeated}} ...{{end -}}
        {{- end -}}:
        {{- range $Index, $Line := $Option.DescriptionLines -}}
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
            {{$Argument.Position}}. [<{{$Argument.Name}}{{":"}}{{$Argument.Type}}>]...
        {{- else -}}
            {{$Argument.Position}}.  <{{$Argument.Name}}{{":"}}{{$Argument.Type}}>
        {{- end -}}
        {{- range $Index, $Line := $Argument.DescriptionLines -}}
            {{"\n            "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}

{{- if .Subcommands -}}
    {{"\n    "}}Subcommands:
    {{- range $Index, $Subcommand := .Subcommands -}}
        {{"\n        "}}{{$Subcommand.Name}}:
        {{- range $Index, $Line := $Subcommand.DescriptionLines -}}
            {{"\n            "}}{{$Line}}
        {{- end -}}
        {{"\n"}}
    {{- end -}}
{{- end -}}
{{"\n\n"}}