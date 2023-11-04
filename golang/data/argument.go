package data

import (
	"github.com/Jumpaku/cliautor/name"
	"github.com/Jumpaku/cliautor/schema"
)

type Argument struct {
	Name     name.Path
	Type     schema.Type
	Variadic bool
}

func (d Argument) InputFieldName() string {
	return d.Name.Map(name.Title).Join("", "Arg_", "")
}

func (d Argument) InputFieldType() string {
	return GoType(d.Type, d.Variadic)
}
