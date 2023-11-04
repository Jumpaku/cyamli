package golang

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"

	"github.com/Jumpaku/cliautor/golang/data"
	"github.com/Jumpaku/cliautor/schema"
)

//go:embed cli.gen.go.tpl
var cliGoTemplate string
var executor = template.Must(template.New("cli.go.tpl").Parse(cliGoTemplate))

func Generate(packageName string, schema *schema.Schema, out io.Writer) error {
	d, err := data.Construct(packageName, schema)
	if err != nil {
		return fmt.Errorf("fail to create CLI data from schema: %w", err)
	}
	err = executor.Execute(out, d)
	if err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}
	return nil
}
