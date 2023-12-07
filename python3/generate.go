package python3

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"

	"github.com/Jumpaku/cyamli/python3/data"
	"github.com/Jumpaku/cyamli/schema"
)

//go:embed cli_gen.py.tpl
var cliPyTemplate string
var executor = template.Must(template.New("cli.py.tpl").Parse(cliPyTemplate))

func Generate(schema *schema.Schema, out io.Writer) error {
	d, err := data.Construct(schema)
	if err != nil {
		return fmt.Errorf("fail to create CLI data from schema: %w", err)
	}
	err = executor.Execute(out, d)
	if err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}
	return nil
}
