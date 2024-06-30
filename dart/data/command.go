package data

import (
	"bytes"
	"fmt"
	"github.com/Jumpaku/cyamli/docs"
	"strings"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
	"github.com/Jumpaku/go-assert"
)

type Command struct {
	schema        *schema.Schema
	schemaCommand *schema.Command
	Name          name.Path
	Options       []Option
	Arguments     []Argument
	Subcommands   []Subcommand
}

func (d Command) FullPathLiteral() string {
	return fmt.Sprintf("%q", d.Name.Join(" ", "", ""))
}

func (d Command) CLIClassName() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "")
}

func (d Command) CLIInputRecordName() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "_Input")
}

func (d Command) CLIFuncMethodChain() string {
	return d.Name.Map(strings.ToLower).Join(".", "", ".FUNC")
}

func (d Command) DocText() string {
	buf := bytes.NewBuffer(nil)
	err := docs.Generate(d.schema, docs.GenerateArgs{
		Format:     docs.DocsFormatText,
		All:        false,
		Subcommand: d.Name,
	}, buf)
	assert.State(err == nil, "fail to generate simple description: %w", err)

	docText := fmt.Sprintf("%q", buf.String())
	return strings.ReplaceAll(docText, `$`, `\$`)
}
