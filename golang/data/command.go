package data

import (
	"bytes"
	"fmt"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
	"github.com/Jumpaku/go-assert"
)

type Command struct {
	schemaCommand *schema.Command
	Name          name.Path
	Options       []Option
	Arguments     []Argument
	Subcommands   []Subcommand
}

func (d Command) FullPathLiteral() string {
	return fmt.Sprintf("%q", d.Name.Join(" ", "", ""))
}

func (d Command) CLIStructName() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "")
}

func (d Command) CLIInputStructName() string {
	return d.Name.Map(name.Title).Join("", "CLI_", "_Input")
}

func (d Command) CLIFuncMethodChain() string {
	return d.Name.Map(name.Title).Join(".", "", ".FUNC")
}

func (d Command) SimpleDescriptionLiteral() string {
	cmdData := description.CreateCommandData("", "", d.Name, d.schemaCommand)
	buf := bytes.NewBuffer(nil)
	err := description.DescribeCommand(description.SimpleExecutor(), cmdData, buf)
	assert.State(err == nil, "fail to generate simple description: %w", err)

	return fmt.Sprintf("%q", buf.String())
}

func (d Command) DetailDescriptionLiteral() string {
	cmdData := description.CreateCommandData("", "", d.Name, d.schemaCommand)
	buf := bytes.NewBuffer(nil)
	err := description.DescribeCommand(description.DetailExecutor(), cmdData, buf)
	assert.State(err == nil, "fail to generate simple description: %w", err)

	return fmt.Sprintf("%q", buf.String())
}
