package cmd

import (
	"fmt"

	"cliautor"
	cliautor_golang "cliautor/golang"
)

var schema, _ = cliautor.Load([]byte("program:\n    name: \"\"\n    version: \"\"\n    description: this is an example command\n    options:\n        -option-a:\n            short: -a\n            description: this is an option for root command\n            type: string\n            default: aaa\n        -option-b:\n            short: -b\n            description: this is an option for root command\n            type: integer\n            default: \"1\"\n    args:\n        - name: argsV\n          description: this is an argument for root command\n          type: string\n          variadic: true\n    subcommands:\n        sub-1:\n            description: this is a sub command\n            options:\n                -option-a:\n                    short: -a\n                    description: this is an option for sub1\n                    type: float\n                    default: \"2\"\n                -option-b:\n                    short: -b\n                    description: this is an option for sub1\n                    type: boolean\n                    default: \"true\"\n            args:\n                - name: argsA\n                  description: this is an argument for sub1\n                  type: string\n                  variadic: false\n                - name: argsB\n                  description: this is an argument for sub1\n                  type: integer\n                  variadic: false\n            subcommands: {}\n"))

type Func[Input any] func(Input)

type ProgramInput struct {
	Opt_OptionA string

	Opt_OptionB int64

	Arg_argsV string
}

type Program struct {
	Func     Func[ProgramInput]
	Cmd_Sub1 Cmd_Sub1
}

type Cmd_Sub1Input struct {
	Opt_OptionA float64

	Opt_OptionB bool

	Arg_argsA string

	Arg_argsB int64
}

type Cmd_Sub1 struct {
	Func Func[Cmd_Sub1Input]
}

func Run(program Program, args []string) error {
	sub, path, rest := cliautor_golang.InterpretSubcommand(schema, args)
	switch cliautor_golang.PathFullIdentifier(path) {
	case `Cmd`:
		input, err := cliautor_golang.InterpretInput[ProgramInput](sub, rest)
		if err != nil {
			return fmt.Errorf("fail to interpret input: %w", err)
		}
		program.Func(input)
	case `Cmd_Sub1`:
		input, err := cliautor_golang.InterpretInput[Cmd_Sub1Input](sub, rest)
		if err != nil {
			return fmt.Errorf("fail to interpret input: %w", err)
		}
		program.Cmd_Sub1.Func(input)

	}
	return nil
}
