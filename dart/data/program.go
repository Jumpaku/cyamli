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

type Program struct {
	schema        *schema.Schema
	schemaProgram *schema.Program
	Name          name.Path
	Version       string
	Options       []Option
	Arguments     []Argument
	Subcommands   []Subcommand
}

func (d Program) FullPathLiteral() string {
	return `""`
}

func (d Program) CLIClassName() string {
	return "CLI"
}

func (d Program) CLIFuncMethodChain() string {
	return "FUNC"
}

func (d Program) CLIInputRecordName() string {
	return "CLI_Input"
}

func (d Program) DocText() string {
	buf := bytes.NewBuffer(nil)
	err := docs.Generate(d.schema, docs.GenerateArgs{
		Format:     docs.DocsFormatText,
		All:        false,
		Subcommand: name.Path{},
	}, buf)
	assert.State(err == nil, "fail to generate simple description: %w", err)

	docText := fmt.Sprintf("%q", buf.String())
	return strings.ReplaceAll(docText, `$`, `\$`)
}
