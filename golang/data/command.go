package data

import "cliautor/name"

type Command struct {
	Name        name.Name
	Options     []Option
	Arguments   []Argument
	Subcommands []Subcommand
}
