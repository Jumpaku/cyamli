{{- /* Go Template */ -}}
## {{.Path}}{{"\n\n"}}


{{- if .Description -}}
### Description{{"\n\n"}}
{{- range $Index, $Line := .DescriptionLines -}}
    {{- if $Index -}}{{"\n"}}{{- end -}}{{$Line}}
{{- end -}}
{{"\n\n"}}
{{- end -}}


### Syntax

```shell
{{.Syntax}}
```
{{"\n"}}


{{- if .Options -}}
### Options{{"\n"}}
{{- range $Index, $Option := .Options -}}
    {{"\n"}}*{{" "}}
    {{- range $Index, $Name := $Option.Options -}}
        {{- if $Index}}, {{end -}}
        `{{$Name}}
        {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
        {{- else -}}=<{{$Option.Type}}>
        {{- end -}}`
    {{- end -}}
    {{"  "}}(default=`{{$Option.Default}}`)
    {{- if $Option.Negation -}},{{"  \n  "}}`{{$Option.NegatedOption}}[=<{{$Option.Type}}>]`{{- end -}}:{{"  \n"}}
    {{- range $Index, $Line := $Option.DescriptionLines -}}
        {{"  "}}{{$Line}}{{"  \n"}}
    {{- end -}}
{{- end -}}
{{"\n"}}
{{- end -}}


{{- if .Arguments -}}
### Arguments{{"\n"}}
{{- range $Index, $Argument := .Arguments -}}
    {{"\n"}}{{$Index}}.{{" "}}
    {{- if $Argument.Variadic -}}
        `[<{{$Argument.Name.LowerSnake}}{{":"}}{{$Argument.Type}}>]...`
    {{- else -}}
        `<{{$Argument.Name.LowerSnake}}{{":"}}{{$Argument.Type}}>`
    {{- end -}}
    {{"  \n"}}
    {{- range $Index, $Line := $Argument.DescriptionLines -}}
        {{"  "}}{{$Line}}{{"  \n"}}
    {{- end -}}
{{- end -}}
{{"\n"}}
{{- end -}}


{{- if .Subcommands -}}
### Subcommands{{"\n"}}
{{- range $Index, $Subcommand := .Subcommands -}}
    {{"\n"}}* {{$Subcommand.Name}}:{{"  \n"}}
    {{- range $Index, $Line := $Subcommand.DescriptionLines -}}
        {{"  "}}{{$Line}}{{"  \n"}}
    {{- end -}}
{{- end -}}
{{"\n"}}
{{- end -}}

{{"\n"}}