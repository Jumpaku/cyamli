package data

import (
	"bytes"
	"fmt"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
	"github.com/Jumpaku/go-assert"
)

type Program struct {
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

func (d Program) CLIInputClassName() string {
	return "CLI_Input"
}

func (d Program) SimpleDescriptionLiteral() string {
	cmdData := description.CreateCommandData(d.schemaProgram.Name, d.schemaProgram.Version, name.Path{}, d.schemaProgram.Command())
	buf := bytes.NewBuffer(nil)
	err := description.DescribeCommand(description.SimpleExecutor(), cmdData, buf)
	assert.State(err == nil, "fail to generate simple description: %w", err)

	return fmt.Sprintf("%q", buf.String())
}

func (d Program) DetailDescriptionLiteral() string {
	cmdData := description.CreateCommandData(d.schemaProgram.Name, d.schemaProgram.Version, name.Path{}, d.schemaProgram.Command())
	buf := bytes.NewBuffer(nil)
	err := description.DescribeCommand(description.DetailExecutor(), cmdData, buf)
	assert.State(err == nil, "fail to generate simple description: %w", err)

	return fmt.Sprintf("%q", buf.String())
}
