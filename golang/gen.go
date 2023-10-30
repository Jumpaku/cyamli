package golang

import (
	"cliautor"
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed cli.go.tpl
var cliGoTemplate string
var executor = template.Must(template.New("cli.go.tpl").Parse(cliGoTemplate))

func Generate(schema *cliautor.Schema, out io.Writer) error {
	data, err := createCLIData(schema)
	if err != nil {
		return fmt.Errorf("fail to create CLI data from schema: %w", err)
	}
	err = executor.Execute(out, data)
	if err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}
	return nil
}
