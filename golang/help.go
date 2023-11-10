package golang

import (
	"fmt"
	"os"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/schema"
)

func HelpFunc[Input any](schema *schema.Schema) func(subcommand []string, input Input, inputErr error) error {
	return func(subcommand []string, input Input, inputErr error) error {
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
