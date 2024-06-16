package dart

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"

	"github.com/Jumpaku/cyamli/dart/data"
	"github.com/Jumpaku/cyamli/schema"
)

//go:embed cli.g.dart.tpl
var cliDartTemplate string
var executor = template.Must(template.New("cli.go.tpl").Parse(cliDartTemplate))

func Generate(schema *schema.Schema, generatorName, generatorVersion string, out io.Writer) error {
	d, err := data.Construct(schema, generatorName, generatorVersion)
	if err != nil {
		return fmt.Errorf("fail to create CLI data from schema: %w", err)
	}

	err = executor.Execute(out, d)
	if err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	return nil
}
