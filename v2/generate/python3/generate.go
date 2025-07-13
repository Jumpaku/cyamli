package python3

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"io"
	"text/template"
)

//go:embed cli_gen.py.tpl
var cliGenPyTemplate string
var executor = template.Must(template.New("cli_gen.py.tpl").Parse(cliGenPyTemplate))

func Generate(schema schema.Schema, generator string, out io.Writer) error {
	d := ConstructData(schema, "", generator)

	buf := bytes.NewBuffer(nil)
	if err := executor.Execute(buf, d); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write generated code: %w", err)
	}
	return nil
}

//go:embed test_cli_gen.py.tpl
var testCliGenPyTemplate string
var executorTest = template.Must(template.New("test_cli_gen.py.tpl").Parse(testCliGenPyTemplate))

func GenerateTest(schema schema.Schema, moduleName, generator string, out io.Writer) error {
	d := ConstructData(schema, moduleName, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorTest.Execute(buf, d); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write generated code: %w", err)
	}
	return nil
}
