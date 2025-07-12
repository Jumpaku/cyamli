package php

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"text/template"
)

type NamedWriter interface {
	Write(name string, b []byte) (int, error)
}

//go:embed Cyamli.php.tpl
var cyamliPhpTemplate string
var executorCyamli = template.Must(template.New("Cyamli.php.tpl").Parse(cyamliPhpTemplate))

//go:embed CLIHandler.php.tpl
var cliHandlerPhpTemplate string
var executorCliHandler = template.Must(template.New("CliHandler.php.tpl").Parse(cliHandlerPhpTemplate))

//go:embed Input.php.tpl
var inputPhpTemplate string
var executorInput = template.Must(template.New("Input.php.tpl").Parse(inputPhpTemplate))

func Generate(schema schema.Schema, namespace, generator string, out NamedWriter) error {
	d := ConstructData(schema, namespace, generator)

	{
		buf := bytes.NewBuffer(nil)
		if err := executorCyamli.Execute(buf, d); err != nil {
			return fmt.Errorf("fail to execute template: %w", err)
		}

		if _, err := out.Write("Cyamli.php", buf.Bytes()); err != nil {
			return fmt.Errorf("fail to write generated code: %w", err)
		}
	}
	{
		buf := bytes.NewBuffer(nil)
		if err := executorCliHandler.Execute(buf, d); err != nil {
			return fmt.Errorf("fail to execute template: %w", err)
		}

		if _, err := out.Write("CliHandler.php", buf.Bytes()); err != nil {
			return fmt.Errorf("fail to write generated code: %w", err)
		}
	}
	for _, d := range d.CommandList {
		buf := bytes.NewBuffer(nil)
		if err := executorInput.Execute(buf, d); err != nil {
			return fmt.Errorf("fail to execute template: %w", err)
		}

		if _, err := out.Write(d.HandlerInputType()+".php", buf.Bytes()); err != nil {
			return fmt.Errorf("fail to write generated code: %w", err)
		}
	}
	return nil
}

//go:embed CLIHandlerMock.php.tpl
var cliHandlerMockPhpTemplate string
var executorMock = template.Must(template.New("CLIHandler.php.tpl").Parse(cliHandlerMockPhpTemplate))

//go:embed RunTest.php.tpl
var runTestPhpTemplate string
var executorTest = template.Must(template.New("RunTest.php.tpl").Parse(runTestPhpTemplate))

func GenerateTest(schema schema.Schema, namespace, generator string, out NamedWriter) error {
	d := ConstructData(schema, namespace, generator)
	{
		buf := bytes.NewBuffer(nil)
		if err := executorMock.Execute(buf, d); err != nil {
			return fmt.Errorf("fail to execute template: %w", err)
		}

		if _, err := out.Write("CLIHandlerMock.php", buf.Bytes()); err != nil {
			return fmt.Errorf("fail to write generated code: %w", err)
		}
	}
	for _, d := range d.CommandList {
		buf := bytes.NewBuffer(nil)
		if err := executorTest.Execute(buf, d); err != nil {
			return fmt.Errorf("fail to execute template: %w", err)
		}

		if _, err := out.Write(d.HandlerMethodName()+"_Test.php", buf.Bytes()); err != nil {
			return fmt.Errorf("fail to write generated code: %w", err)
		}
	}

	return nil
}
