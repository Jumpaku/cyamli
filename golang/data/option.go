package data

import (
	"cliautor/name"
	"strings"

	"github.com/samber/lo"
)

type Option struct {
	Name   name.Name
	GoType string
}

func (d Option) Identifier() string {
	titled := lo.Map(d.Name, func(p string, i int) string { return name.Title(p) })
	return "Opt_" + strings.Join(titled, "")
}
