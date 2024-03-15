package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"

	_ "embed"

	"github.com/Jumpaku/cyamli/golang"
	"github.com/Jumpaku/cyamli/python3"
	"github.com/Jumpaku/cyamli/schema"
)

//go:embed cli.yaml
var schemaData []byte

func LoadSchema() *schema.Schema {
	s, _ := schema.Load(bytes.NewBuffer(schemaData))
	return s
}

var cli = NewCLI()

func main() {
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		switch {
		case inputErr != nil:
			fmt.Println(cli.DESC_Simple())
		case input.Opt_Version:
			fmt.Printf("version: %s\n", LoadSchema().Program.Version)
		case input.Opt_Help:
			fmt.Println(cli.DESC_Detail())
		}
		return nil
	}
	cli.Golang.FUNC = funcGolang
	cli.Python3.FUNC = funcPython3

	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func funcGolang(subcommand []string, input CLI_Golang_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(cli.DESC_Simple())
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
	}
	if input.Opt_Help {
		fmt.Println(cli.Golang.DESC_Detail())
		return nil
	}
	var reader io.Reader = os.Stdin
	if input.Opt_SchemaPath != "" {
		f, err := os.Open(input.Opt_SchemaPath)
		if err != nil {
			return fmt.Errorf("fail to open schema file %q: %w", input.Opt_SchemaPath, err)
		}
		defer f.Close()

		reader = f
	}

	schema, err := schema.Load(reader)
	if err != nil {
		return fmt.Errorf("fail to open schema file %q: %w", input.Opt_SchemaPath, err)
	}

	var writer io.Writer = os.Stdout
	if input.Opt_OutPath != "" {
		f, err := os.Create(input.Opt_OutPath)
		if err != nil {
			return fmt.Errorf("fail to open output file %q: %w", input.Opt_OutPath, err)
		}
		defer f.Close()

		writer = f
	}

	buf := bytes.NewBuffer(nil)
	if err := golang.Generate(input.Opt_Package, schema, buf); err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}
	fmtBytes, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("fail to format cli code: %w", err)
	}
	if _, err := writer.Write(fmtBytes); err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}

	return nil
}

func funcPython3(subcommand []string, input CLI_Python3_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(cli.DESC_Simple())
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
	}
	if input.Opt_Help {
		fmt.Println(cli.Golang.DESC_Detail())
		return nil
	}
	var reader io.Reader = os.Stdin
	if input.Opt_SchemaPath != "" {
		f, err := os.Open(input.Opt_SchemaPath)
		if err != nil {
			return fmt.Errorf("fail to open schema file %q: %w", input.Opt_SchemaPath, err)
		}
		defer f.Close()

		reader = f
	}

	schema, err := schema.Load(reader)
	if err != nil {
		return fmt.Errorf("fail to open schema file %q: %w", input.Opt_SchemaPath, err)
	}

	var writer io.Writer = os.Stdout
	if input.Opt_OutPath != "" {
		f, err := os.Create(input.Opt_OutPath)
		if err != nil {
			return fmt.Errorf("fail to open output file %q: %w", input.Opt_OutPath, err)
		}
		defer f.Close()

		writer = f
	}

	buf := bytes.NewBuffer(nil)
	if err := python3.Generate(schema, buf); err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}

	if _, err := writer.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}

	return nil
}
