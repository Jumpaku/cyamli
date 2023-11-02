package data

import (
	"cliautor/name"
	"strings"

	"github.com/samber/lo"
)

type Subcommand struct {
	Name name.Name
}

func (d Subcommand) Identifier() string {
	if len(d.Name) == 0 {
		return "Cmd"
	}
	return "Cmd_" + name.Title(strings.Join(name.MakeName(d.Name[len(d.Name)-1]), ""))
}

func (d Subcommand) FullIdentifier() string {
	if len(d.Name) == 0 {
		return "Cmd"
	}
	titled := lo.Map(d.Name, func(p string, i int) string { return name.Title(strings.Join(name.MakeName(p), "")) })
	return "Cmd_" + strings.Join(titled, "")
}

func (d Subcommand) FuncIdentifier() string {
	if len(d.Name) == 0 {
		return "Func"
	}
	titled := lo.Map(d.Name, func(p string, i int) string { return "Cmd_" + name.Title(strings.Join(name.MakeName(p), "")) })
	return strings.Join(titled, ".") + ".Func"
}
