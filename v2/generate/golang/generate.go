package golang

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"go/format"
	"io"
	"text/template"
)

//go:embed cli.gen.go.tpl
var cliGenGoTemplate string
var executor = template.Must(template.New("cli.gen.go.tpl").Parse(cliGenGoTemplate))

func Generate(schema schema.Schema, packageName, generator string, out io.Writer) error {
	d := ConstructData(schema, "", packageName, generator)

	buf := bytes.NewBuffer(nil)
	if err := executor.Execute(buf, d); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
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

//go:embed cli.gen_test.go.tpl
var cliGenTestGoTemplate string
var executorTest = template.Must(template.New("cli.gen_test.go.tpl").Parse(cliGenTestGoTemplate))

func GenerateTest(schema schema.Schema, moduleName, packageName, generator string, out io.Writer) error {
	d := ConstructData(schema, moduleName, packageName, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorTest.Execute(buf, d); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
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
