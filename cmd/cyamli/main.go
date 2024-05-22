package main

import (
	"bytes"
	"fmt"
	"github.com/Jumpaku/cyamli/docs"
	"github.com/Jumpaku/cyamli/name"
	"go/format"
	"io"
	"os"
	"slices"

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

//go:generate go run "github.com/Jumpaku/cyamli/cmd/cyamli@v0.0.15" generate golang -package=main -schema-path=cli.yaml -out-path=cli.gen.go

//go:generate go run "../../internal/cmd/gen-docs" -- cli.yaml docs.gen.go main

func main() {
	var cli = NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		switch {
		case inputErr != nil:
			fmt.Println(GetDoc(subcommand))
		case input.Opt_Version:
			fmt.Printf("version: %s\n", LoadSchema().Program.Version)
		case input.Opt_Help:
			fmt.Println(GetDoc(subcommand))
		}
		return nil
	}
	cli.List.FUNC = funcList
	cli.Generate.Golang.FUNC = funcGenerateGolang
	cli.Generate.Python3.FUNC = funcGeneratePython3
	cli.Generate.Docs.FUNC = funcGenerateDocs

	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func funcList(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(GetDoc(subcommand))
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
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

	s, err := schema.Load(reader)
	if err != nil {
		return fmt.Errorf("fail to open schema file %q: %w", input.Opt_SchemaPath, err)
	}

	subcommands := []name.Path{}
	_ = s.Walk(func(path name.Path, cmd *schema.Command) error {
		subcommands = append(subcommands, path)
		return nil
	})
	slices.SortFunc(subcommands, func(a, b name.Path) int { return slices.Compare(a, b) })

	for _, path := range subcommands {
		program := s.Program.Name
		if program == "" {
			program = "<program>"
		}

		_, err := fmt.Fprintln(os.Stdout, append(name.Path{program}, path...).Join(" ", "", ""))
		if err != nil {
			return fmt.Errorf("fail to print %q: %w", input.Opt_SchemaPath, err)
		}
	}

	return nil
}

func funcGenerateGolang(subcommand []string, input CLI_GenerateGolang_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(GetDoc(subcommand))
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
	}
	if input.Opt_Help {
		fmt.Println(GetDoc(subcommand))
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

func funcGeneratePython3(subcommand []string, input CLI_GeneratePython3_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(GetDoc(subcommand))
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
	}
	if input.Opt_Help {
		fmt.Println(GetDoc(subcommand))
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

func funcGenerateDocs(subcommand []string, input CLI_GenerateDocs_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(GetDoc(subcommand))
		return fmt.Errorf("fail to resolve command line arguments: %w", inputErr)
	}
	if input.Opt_Help {
		fmt.Println(GetDoc(subcommand))
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

	args := docs.GenerateArgs{
		Format:     docs.DocsFormat(input.Opt_Format),
		All:        input.Opt_All,
		Subcommand: input.Arg_Subcommands,
	}
	if err := docs.Generate(schema, args, writer); err != nil {
		return fmt.Errorf("fail to generate cli %q: %w", input.Opt_SchemaPath, err)
	}

	return nil
}
