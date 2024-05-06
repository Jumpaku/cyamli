{{- /* Go Template */ -}}

{{- $Program := .Program -}}
{{- $Version := .Version -}}

# {{$Program}}{{- if $Version -}}{{" "}}({{$Version}}){{- end -}}{{"\n\n"}}


{{- range .Commands -}}


## {{$Program}}{{- range .Path -}}{{" "}}{{.}}{{- end -}}{{"\n\n"}}

{{- if .Description -}}
    ### Description{{"\n\n"}}
    {{- range $Index, $Line := .Description -}}
        {{- if $Index -}}{{"\n"}}{{- end -}}{{$Line}}
    {{- end -}}
    {{"\n\n"}}
{{- end -}}

### Syntax

```shell
{{$Program}} {{.Syntax}}
```
{{"\n"}}


{{- if .Options -}}
    ### Options{{"\n"}}
    {{- range $Index, $Option := .Options -}}
        {{"\n"}}*{{" "}}
        {{- range $Index, $Name := $Option.Names -}}
            {{- if $Index}}, {{end -}}
            `{{$Name}}
            {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
            {{- else -}}=<{{$Option.Type}}>
            {{- end -}}`
        {{- end -}}
        {{"  "}}(default=`{{$Option.Default}}`):{{"  \n"}}
        {{- range $Index, $Line := $Option.Description -}}
            {{"  "}}{{$Line}}{{"  \n"}}
        {{- end -}}
    {{- end -}}
    {{"\n"}}
{{- end -}}


{{- if .Arguments -}}
    ### Arguments{{"\n"}}
    {{- range $Index, $Argument := .Arguments -}}
        {{"\n"}}{{func_add $Index 1}}.{{" "}}
        {{- if $Argument.Variadic -}}
            `[<{{$Argument.Name}}{{":"}}{{$Argument.Type}}>]...`
        {{- else -}}
            `<{{$Argument.Name}}{{":"}}{{$Argument.Type}}>`
        {{- end -}}
        {{"  \n"}}
        {{- range $Index, $Line := $Argument.Description -}}
            {{"  "}}{{$Line}}{{"  \n"}}
        {{- end -}}
    {{- end -}}
    {{"\n"}}
{{- end -}}


{{- if .Subcommands -}}
    ### Subcommands{{"\n"}}
    {{- range $Index, $Subcommand := .Subcommands -}}
        {{"\n"}}* {{$Subcommand.Name}}:{{"  \n"}}
        {{- range $Index, $Line := $Subcommand.Description -}}
            {{"  "}}{{$Line}}{{"  \n"}}
        {{- end -}}
    {{- end -}}
    {{"\n"}}
{{- end -}}

{{"\n"}}
{{- end -}}