package dart3

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"io"
	"text/template"
)

//go:embed cli.gen.dart.tpl
var cliGenDartTemplate string
var executor = template.Must(template.New("cli.gen.dart.tpl").Parse(cliGenDartTemplate))

//go:embed cli.gen_test.dart.tpl
var cliGenTestDartTemplate string
var executorTest = template.Must(template.New("cli.gen_test.dart.tpl").Parse(cliGenTestDartTemplate))

func Generate(schema schema.Schema, generator string, out io.Writer) error {
	d := ConstructData(schema, generator)

	buf := bytes.NewBuffer(nil)
	if err := executor.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}

func GenerateTest(schema schema.Schema, generator string, out io.Writer) error {
	d := ConstructData(schema, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorTest.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}
