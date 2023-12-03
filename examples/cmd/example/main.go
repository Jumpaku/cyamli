package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/schema"
	"github.com/davecgh/go-spew/spew"
)

//go:generate go run "github.com/Jumpaku/cyamli/cmd/cyamli@latest" golang -schema-path=cli.yaml -out-path=cli.gen.go
func main() {
	schema := LoadSchema()
	cli := NewCLI()
	cli.FUNC = HelpFunc[CLI_Input](schema)
	cli.Sub1.FUNC = HelpFunc[CLI_Sub1_Input](schema)
	cli.Sub2.FUNC = HelpFunc[CLI_Sub2_Input](schema)
	cli.Sub3.FUNC = HelpFunc[CLI_Sub3_Input](schema)
	cli.Sub3.Suba.FUNC = HelpFunc[CLI_Sub3Suba_Input](schema)
	cli.Sub3.Subb.FUNC = HelpFunc[CLI_Sub3Subb_Input](schema)
	cli.Sub3.Subc.FUNC = HelpFunc[CLI_Sub3Subc_Input](schema)
	cli.Sub3.Subd.FUNC = HelpFunc[CLI_Sub3Subd_Input](schema)
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

//go:embed cli.yaml
var schemaData []byte

func LoadSchema() *schema.Schema {
	s, _ := schema.Load(bytes.NewBuffer(schemaData))
	return s
}

func HelpFunc[Input any](schema *schema.Schema) func(subcommand []string, input Input, inputErr error) error {
	return func(subcommand []string, input Input, inputErr error) error {
		if inputErr != nil {
			return inputErr
		}
		spew.Dump(input)

		cmd := schema.Find(subcommand)
		err := description.DescribeCommand(
			description.DetailExecutor(),
			description.CreateCommandData(schema.Program.Name, schema.Program.Version, subcommand, cmd),
			os.Stderr)
		if err != nil {
			return fmt.Errorf("fail to generate help")
		}
		return err
	}
}
