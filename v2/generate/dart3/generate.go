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

func Generate(schema schema.Schema, packageName, generator string, out io.Writer) error {
	d := ConstructData(schema, packageName, generator)

	buf := bytes.NewBuffer(nil)
	if err := executor.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}
