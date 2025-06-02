{{/* Go Template */}}
        <section>
            <h2>{{ .Path }}</h2>

            {{ if .DescriptionLines }}
            <h3>Description</h3>
            <p>
                {{ range $Index, $Line := .DescriptionLines }}{{$Line}}
                {{ end }}
            </p>
            {{ end }}

            <h3>Syntax</h3>
            <p>
                <code>{{.Syntax}}</code>
            </p>

            {{ if .Options }}
            <h3>Options</h3>
            <ul>
                {{ range $Index, $Option := .Options }}
                    <li>
                        <h4>
                            {{ range $Index, $Name := $Option.Options -}}
                                {{ if $Index }}, {{ end }}
                                <code>
                                    {{$Name}}
                                    {{- if (eq $Option.Type "boolean") -}}[=<{{$Option.Type}}>]
                                    {{- else -}}=<{{$Option.Type}}>
                                    {{- end -}}
                                </code>
                            {{- end }} (default=<code>{{$Option.Default}}</code>){{ if $Option.Negation -}}, <br><code>{{$Option.NegatedOption}}[=<{{$Option.Type}}>]</code>{{ end }}:
                        </h4>
                        <p>
                            {{ range $Index, $Line := $Option.DescriptionLines }}{{$Line}}
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
                            {{ range $Index, $Line := $Argument.DescriptionLines }}{{$Line}}
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
                            {{ range $Index, $Line := $Subcommand.DescriptionLines }}{{$Line}}
                            {{ end }}
                        </p>
                    </li>
                {{ end }}
            </ul>
            {{ end }}
        </section>