package golang

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"text/template"

	"github.com/Jumpaku/cyamli/golang/data"
	"github.com/Jumpaku/cyamli/schema"
)

//go:embed cli.gen.go.tpl
var cliGoTemplate string
var executor = template.Must(template.New("cli.go.tpl").Parse(cliGoTemplate))

func Generate(schema *schema.Schema, generatorName, generatorVersion string, packageName string, out io.Writer) error {
	d, err := data.Construct(schema, generatorName, generatorVersion, packageName)
	if err != nil {
		return fmt.Errorf("fail to create CLI data from schema: %w", err)
	}

	buf := bytes.NewBuffer(nil)

	err = executor.Execute(buf, d)
	if err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("fail to format generated code: %w", err)
	}

	if _, err := out.Write(b); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}
