package data

import "cliautor/name"

type Argument struct {
	Name   name.Name
	GoType string
}

func (d Argument) Identifier() string {
	return "Arg_" + d.Name.Join("", "", "")
}
