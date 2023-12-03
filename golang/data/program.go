package data

import (
	"github.com/Jumpaku/cyamli/name"
)

type Program struct {
	Name        name.Path
	Version     string
	Options     []Option
	Arguments   []Argument
	Subcommands []Subcommand
}

func (d Program) FullPathLiteral() string {
	return `""`
}

func (d Program) CLIStructName() string {
	return "CLI"
}

func (d Program) CLIFuncMethodChain() string {
	return "FUNC"
}

func (d Program) CLIInputStructName() string {
	return "CLI_Input"
}
