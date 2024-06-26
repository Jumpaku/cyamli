package data

import (
	"github.com/Jumpaku/cyamli/name"
	"strings"
)

type Subcommand struct {
	Name name.Path
}

func (d Subcommand) SubcommandFieldName() string {
	return strings.ToLower(d.Name[len(d.Name)-1])
}

func (d Subcommand) SubcommandFieldType() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "")
}
