package data

import (
	"cliautor/name"
	"fmt"
)

type Program struct {
	Name        name.Path
	Version     string
	Description string
	Options     []Option
	Arguments   []Argument
	Subcommands []Subcommand
}

func (d Program) CLIStructName() string {
	return "CLI"
}

func (d Program) CLIInputStructName() string {
	return "CLI_Input"
}

func (d Program) DescriptionLiteral() string {
	return fmt.Sprintf("%q", d.Description)
}
