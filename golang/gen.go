package golang

import (
	"cliautor/golang/data"
	"cliautor/schema"
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed cli.go.tpl
var cliGoTemplate string
var executor = template.Must(template.New("cli.go.tpl").Parse(cliGoTemplate))

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
