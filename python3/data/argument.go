package data

import (
	"strings"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
)

type Argument struct {
	Name     name.Path
	Type     schema.Type
	Variadic bool
}

func (d Argument) InputFieldName() string {
	return d.Name.Map(strings.ToLower).Join("_", "arg_", "")
}

func (d Argument) InputFieldType() string {
	return Python3Type(d.Type, d.Variadic)
}
