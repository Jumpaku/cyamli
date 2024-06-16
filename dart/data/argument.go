package data

import (
	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
)

type Argument struct {
	Name     name.Path
	Type     schema.Type
	Variadic bool
}

func (d Argument) InputFieldName() string {
	return d.Name.Map(name.Title).Join("", "arg", "")
}

func (d Argument) InputFieldType() string {
	return DartType(d.Type, d.Variadic)
}
