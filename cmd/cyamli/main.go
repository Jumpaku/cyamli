package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/golang"
	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		if input.Opt_Version {
			fmt.Printf("version: %s\n", LoadSchema().Program.Version)
			return nil
		}
		showDetailDescription(subcommand, os.Stdout)
		return nil
	}
	cli.Golang.FUNC = funcGolang

	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func showSimpleDescription(subcommand []string, writer io.Writer, inputErr error) {
	schema := LoadSchema()
	_ = description.DescribeCommand(
		description.SimpleExecutor(),
		description.CreateCommandData(schema.Program.Name, schema.Program.Version, subcommand, schema.Find(subcommand)),
		writer,
	)
}

func showDetailDescription(subcommand []string, writer io.Writer) {
	schema := LoadSchema()
	_ = description.DescribeCommand(
		description.DetailExecutor(),
		description.CreateCommandData(schema.Program.Name, schema.Program.Version, subcommand, schema.Find(subcommand)),
		writer,
	)
}

func funcGolang(subcommand []string, input CLI_Golang_Input, inputErr error) (err error) {
	if inputErr != nil {
		showSimpleDescription(subcommand, os.Stderr, inputErr)
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
	}
	if input.Opt_Help {
		showDetailDescription(subcommand, os.Stdout)
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

	err = golang.Generate(input.Opt_Package, schema, writer)
	if err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}

	return nil
}
