{{/* Go Template */}}
{{- $Program := .Program -}}
{{- $Version := .Version -}}

<html lang="">
<head>
    <meta charset="UTF-8" >
    <title>{{$Program}}{{ if $Version }}{{" "}}({{$Version}}){{ end }}</title>
</head>
<body>
<main>
    <h1>{{$Program}}{{ if $Version }}{{" "}}({{$Version}}){{ end }}</h1>
    {{ range .Commands }}
        <section>
            <h2> {{$Program}}{{ range .Path }}{{" "}}{{.}}{{ end }}</h2>

            {{ if .Description }}
                <h3>Description</h3>
                <p>
                    {{ range $Index, $Line := .Description }}{{$Line}}
                    {{ end }}
                </p>
            {{ end }}

            <h3>Syntax</h3>
            <p>
                <code>{{$Program}} {{.Syntax}}</code>
            </p>


            {{ if .Options }}
                <h3>Options</h3>
                <ul>
                    {{ range $Index, $Option := .Options }}
                        <li>
                            <h4>
                                {{ range $Index, $Name := $Option.Names -}}
                                    {{ if $Index }}, {{ end }}
                                    <code>
                                        {{$Name}}
                                        {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
                                        {{- else -}}=<{{$Option.Type}}>
                                        {{- end -}}
                                    </code>
                                {{- end }} (default=<code>{{$Option.Default}}</code>):
                            </h4>
                            <p>
                                {{ range $Index, $Line := $Option.Description }}{{$Line}}
                                {{ end }}
                            </p>
                        </li>
                    {{ end }}
                </ul>
            {{ end }}


            {{ if .Arguments }}
                <h3>Arguments</h3>
                <ol>
                    {{ range $Index, $Argument := .Arguments }}
                        <li>
                            <h4>
                                {{ if $Argument.Variadic }}
                                    <code>[<{{$Argument.Name}}{{":"}}{{$Argument.Type}}>]...</code>
                                {{ else }}
                                    <code><{{$Argument.Name}}{{":"}}{{$Argument.Type}}></code>
                                {{ end }}
                            </h4>
                            <p>
                                {{ range $Index, $Line := $Argument.Description }}{{$Line}}
                                {{ end }}
                            </p>
                        </li>
                    {{ end }}
                </ol>
            {{ end }}


            {{ if .Subcommands }}
                <h3>Subcommands</h3>
                <ul>
                    {{ range $Index, $Subcommand := .Subcommands }}
                        <li>
                            <h4>{{$Subcommand.Name}}</h4>
                            <p>
                                {{ range $Index, $Line := $Subcommand.Description }}{{$Line}}
                                {{ end }}
                            </p>
                        </li>
                    {{ end }}
                </ul>
            {{ end }}
        </section>
    {{ end }}
</main>
</body>
</html>
