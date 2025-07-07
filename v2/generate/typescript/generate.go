package typescript

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"io"
	"text/template"
)

//go:embed cli.gen.ts.tpl
var cliGenTsTemplate string
var executor = template.Must(template.New("cli.gen.ts.tpl").Parse(cliGenTsTemplate))

func Generate(schema schema.Schema, generator string, out io.Writer) error {
	d := ConstructData(schema, "", generator)

	buf := bytes.NewBuffer(nil)
	if err := executor.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}

//go:embed cli.gen.test.ts.tpl
var cliGenTestTsTemplate string
var executorTest = template.Must(template.New("cli.gen.test.ts.tpl").Parse(cliGenTestTsTemplate))

func GenerateTest(schema schema.Schema, moduleFile, generator string, out io.Writer) error {
	d := ConstructData(schema, moduleFile, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorTest.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}
