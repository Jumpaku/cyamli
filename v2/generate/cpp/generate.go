package cpp

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"io"
	"text/template"
)

//go:embed cli.gen.cpp.tpl
var cliGenCppTemplate string
var executorCpp = template.Must(template.New("cli.gen.cpp.tpl").Parse(cliGenCppTemplate))

func GenerateCpp(schema schema.Schema, headerFile, namespace, generator string, out io.Writer) error {
	d := ConstructData(schema, headerFile, namespace, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorCpp.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}

//go:embed cli.gen.h.tpl
var cliGenHTemplate string
var executorH = template.Must(template.New("cli.gen.h.tpl").Parse(cliGenHTemplate))

func GenerateH(schema schema.Schema, namespace, generator string, out io.Writer) error {
	d := ConstructData(schema, "", namespace, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorH.Execute(buf, d); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to write generated code: %w", err)
	}
	return nil
}

/*
//go:embed cli.gen_test.go.tpl
var cliGenTestGoTemplate string
var executorTest = template.Must(template.New("cli.gen_test.go.tpl").Parse(cliGenTestGoTemplate))

func GenerateTest(schema schema.Schema, moduleName, packageName, generator string, out io.Writer) error {
	d := ConstructData(schema, moduleName, packageName, generator)

	buf := bytes.NewBuffer(nil)
	if err := executorTest.Execute(buf, d); err != nil {
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
*/
