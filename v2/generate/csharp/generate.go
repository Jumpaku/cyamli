package csharp

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"io"
	"text/template"
)

//go:embed cli.gen.cs.tpl
var cliGenGoTemplate string
var executor = template.Must(template.New("cli.gen.cs.tpl").Parse(cliGenGoTemplate))

func Generate(schema schema.Schema, namespace, generator string, out io.Writer) error {
	d := ConstructData(schema, namespace, generator)

	buf := bytes.NewBuffer(nil)
	if err := executor.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}

//go:embed cli_test.gen.cs.tpl
var cliTestGenGoTemplate string
var executorTest = template.Must(template.New("cli_test.gen.cs.tpl").Parse(cliTestGenGoTemplate))

func GenerateTest(schema schema.Schema, namespace, generator string, out io.Writer) error {
	d := ConstructData(schema, namespace, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorTest.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}
