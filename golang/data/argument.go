package data

import (
	"cliautor/name"
	"cliautor/schema"
	"fmt"
)

type Argument struct {
	Name        name.Path
	Type        schema.Type
	Variadic    bool
	Description string
}

func (d Argument) InputFieldName() string {
	return d.Name.Map(name.Title).Join("", "Arg_", "")
}

func (d Argument) InputFieldType() string {
	return GoType(d.Type, d.Variadic)
}

func (d Argument) DescriptionLiteral() string {
	return fmt.Sprintf("%q", d.Description)
}
