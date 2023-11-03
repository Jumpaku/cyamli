package main

import (
	"cliautor/golang"
	"cliautor/schema"
	"fmt"
	"io"
	"os"
)

func main() {
	cli := NewCLI()
	cli.Sub_Golang.Func = funcGolang

	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func funcGolang(subcommand []string, input CLI_Golang_Input) (err error) {
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

	err = golang.Generate(input.Opt_Package, schema, writer)
	if err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}

	return nil
}
