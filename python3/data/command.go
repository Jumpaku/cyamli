package data

import (
	"fmt"
	"strings"

	"github.com/Jumpaku/cyamli/name"
)

type Command struct {
	Name        name.Path
	Options     []Option
	Arguments   []Argument
	Subcommands []Subcommand
}

func (d Command) FullPathLiteral() string {
	return fmt.Sprintf("%q", d.Name.Join(" ", "", ""))
}

func (d Command) CLIClassName() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "")
}

func (d Command) CLIInputClassName() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "_Input")
}

func (d Command) CLIFuncMethodChain() string {
	return d.Name.Map(strings.ToLower).Join(".", "", ".FUNC")
}
