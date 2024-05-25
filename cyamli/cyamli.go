package cyamli

import (
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/docs"
	"github.com/Jumpaku/cyamli/golang"
	"github.com/Jumpaku/cyamli/info"
	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/python3"
	"github.com/Jumpaku/cyamli/schema"
	"io"
	"os"
	"slices"
	"strings"
)

func mustPrintf(w io.Writer, format string, args ...any) {
	if _, err := fmt.Fprintf(w, format, args...); err != nil {
		panic(err)
	}
}

func panicIfErrorf(err error, format string, args ...any) {
	if err != nil {
		panic(fmt.Errorf(format+": %w", append(args, err)...))
	}
}

func Execute(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (code int) {
	defer func() {
		if err := recover(); err != nil {
			code = 1
			mustPrintf(stderr, "%+v\n", err)
		}
	}()

	var cli = NewCLI()
	impl := implementation{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	cli.FUNC = impl.root
	cli.List.FUNC = impl.list
	cli.Generate.FUNC = impl.generate
	cli.Generate.Golang.FUNC = impl.generateGolang
	cli.Generate.Python3.FUNC = impl.generatePython3
	cli.Generate.Docs.FUNC = impl.generateDocs

	if err := Run(cli, args); err != nil {
		panic(err)
	}

	return 0
}

type implementation struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (i implementation) root(subcommand []string, input CLI_Input, inputErr error) (err error) {
	panicIfErrorf(inputErr, "failed to parse arguments ('cyamli -help' to see help)")

	switch {
	case input.Opt_Version:
		mustPrintf(i.Stdout, "version: %s\n", info.Version)
	default:
		mustPrintf(i.Stdout, GetDoc(subcommand))
	}
	return nil
}

func (i implementation) list(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
	panicIfErrorf(inputErr, "failed to parse arguments ('cyamli %s -help' to see help)", strings.Join(subcommand, " "))

	if input.Opt_Help {
		mustPrintf(i.Stdout, GetDoc(subcommand))
		return nil
	}

	var reader = i.Stdin
	schemaSource := "(stdin)"
	if input.Opt_SchemaPath != "" {
		schemaSource = input.Opt_SchemaPath
		f, err := os.Open(schemaSource)
		panicIfErrorf(err, "fail to open schema file %q", schemaSource)
		defer f.Close()
		reader = f
	}

	s, err := schema.Load(reader)
	panicIfErrorf(err, "fail to read schema from %q", schemaSource)

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

		mustPrintf(i.Stdout, append(name.Path{program}, path...).Join(" ", "", "\n"))
	}

	return nil
}

func (i implementation) generate(subcommand []string, input CLI_Generate_Input, inputErr error) (err error) {
	panicIfErrorf(inputErr, "failed to parse arguments ('cyamli -help' to see help)")

	mustPrintf(i.Stdout, GetDoc(subcommand))

	return nil
}

func (i implementation) generateGolang(subcommand []string, input CLI_GenerateGolang_Input, inputErr error) (err error) {
	panicIfErrorf(inputErr, "failed to parse arguments ('cyamli %s -help' to see help)", strings.Join(subcommand, " "))

	if input.Opt_Help {
		mustPrintf(i.Stdout, GetDoc(subcommand))
		return nil
	}
	var reader = i.Stdin
	schemaSource := "(stdin)"
	if input.Opt_SchemaPath != "" {
		schemaSource = input.Opt_SchemaPath
		f, err := os.Open(schemaSource)
		panicIfErrorf(err, "fail to open schema file %q", schemaSource)
		defer f.Close()
		reader = f
	}

	schema, err := schema.Load(reader)
	panicIfErrorf(err, "fail to read schema file %q", schemaSource)

	var writer = i.Stdout
	if input.Opt_OutPath != "" {
		f, err := os.Create(input.Opt_OutPath)
		panicIfErrorf(err, "fail to open output file %q", input.Opt_OutPath)
		defer f.Close()
		writer = f
	}

	err = golang.Generate(schema, info.Name, info.Version, input.Opt_Package, writer)
	panicIfErrorf(err, "fail to generate cli in golang")

	return nil
}

func (i implementation) generatePython3(subcommand []string, input CLI_GeneratePython3_Input, inputErr error) (err error) {
	panicIfErrorf(inputErr, "failed to parse arguments ('cyamli %s -help' to see help)", strings.Join(subcommand, " "))

	if input.Opt_Help {
		mustPrintf(i.Stdout, GetDoc(subcommand))
		return nil
	}

	var reader = i.Stdin
	if input.Opt_SchemaPath != "" {
		f, err := os.Open(input.Opt_SchemaPath)
		panicIfErrorf(err, "fail to open schema file %q", input.Opt_SchemaPath)
		defer f.Close()
		reader = f
	}

	schema, err := schema.Load(reader)
	panicIfErrorf(err, "fail to read schema file %q", input.Opt_SchemaPath)

	var writer = i.Stdout
	if input.Opt_OutPath != "" {
		f, err := os.Create(input.Opt_OutPath)
		panicIfErrorf(err, "fail to open output file %q", input.Opt_OutPath)
		defer f.Close()
		writer = f
	}

	err = python3.Generate(schema, info.Name, info.Version, writer)
	panicIfErrorf(err, "fail to generate cli in python3")

	return nil
}

func (i implementation) generateDocs(subcommand []string, input CLI_GenerateDocs_Input, inputErr error) (err error) {
	panicIfErrorf(inputErr, "failed to parse arguments ('cyamli %s -help' to see help)", strings.Join(subcommand, " "))

	if input.Opt_Help {
		mustPrintf(i.Stdout, GetDoc(subcommand))
		return nil
	}

	var reader = i.Stdin
	if input.Opt_SchemaPath != "" {
		f, err := os.Open(input.Opt_SchemaPath)
		panicIfErrorf(err, "fail to open schema file %q", input.Opt_SchemaPath)
		defer f.Close()
		reader = f
	}

	schema, err := schema.Load(reader)
	panicIfErrorf(err, "fail to read schema file %q", input.Opt_SchemaPath)

	var writer = i.Stdout
	if input.Opt_OutPath != "" {
		f, err := os.Create(input.Opt_OutPath)
		panicIfErrorf(err, "fail to open output file %q", input.Opt_OutPath)
		defer f.Close()
		writer = f
	}

	args := docs.GenerateArgs{
		Format:     docs.DocsFormat(input.Opt_Format),
		All:        input.Opt_All,
		Subcommand: input.Arg_Subcommands,
	}
	err = docs.Generate(schema, args, writer)
	panicIfErrorf(err, "fail to generate cli in python3")

	return nil
}
