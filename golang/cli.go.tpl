package {{.Package}}

import (
	"fmt"

	"cliautor"
	cliautor_golang "cliautor/golang"
)

var schema, _ = cliautor.Load([]byte({{.SchemaYAML}}))

type Func[Input any] func(Input)

{{/* root command */}}
{{with .Program}}

type ProgramInput struct {
{{ range $Index, $Option := .Options }}
	{{$Option.Identifier}} {{$Option.GoType}}
{{ end }}
{{ range $Index, $Arg := .Args }}	{{$Arg.Identifier}} {{$Arg.GoType}}{{ end }}
}

type Program struct {
	Func Func[ProgramInput]
{{ range $Index, $Subcommand := .Subcommands }}	{{$Subcommand.Identifier}} {{$Subcommand.FullIdentifier}}{{ end }}
}

{{end}}


{{/* subcommands */}}
{{range .Commands}}

type {{.FullIdentifier}}Input struct {
{{range $Index, $Option := .Options}}
	{{$Option.Identifier}} {{$Option.GoType}}
{{end}}
{{range $Index, $Arg := .Args}}
	{{$Arg.Identifier}} {{$Arg.GoType}}
{{end}}
}

type {{.FullIdentifier}} struct {
	Func Func[{{.FullIdentifier}}Input]

{{range $Index, $Subcommand := .Subcommands}}
	{{$Subcommand.Identifier}} {{$Subcommand.FullIdentifier}}
{{end}}
}

{{end}}


{{/* entrypoint */}}
func Run(program Program, args []string) error {
	sub, path, rest := cliautor_golang.InterpretSubcommand(schema, args)
	switch cliautor_golang.PathFullIdentifier(path) { 
	case `Cmd`:
		input, err := cliautor_golang.InterpretInput[ProgramInput](sub, rest)
		if err != nil {
			return fmt.Errorf("fail to interpret input: %w", err)
		}
		program.Func(input)
{{range .Commands}}	case `{{.FullIdentifier}}`:
		input, err := cliautor_golang.InterpretInput[{{.FullIdentifier}}Input](sub, rest)
		if err != nil {
			return fmt.Errorf("fail to interpret input: %w", err)
		}
		program.{{.FuncIdentifier}}(input)
{{end}}
	}
	return nil
}