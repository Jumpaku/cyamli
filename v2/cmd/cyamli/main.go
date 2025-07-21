package main

import (
	"bytes"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/generate/cpp"
	"github.com/Jumpaku/cyamli/v2/generate/csharp"
	"github.com/Jumpaku/cyamli/v2/generate/dart3"
	"github.com/Jumpaku/cyamli/v2/generate/docs"
	"github.com/Jumpaku/cyamli/v2/generate/golang"
	"github.com/Jumpaku/cyamli/v2/generate/kotlin"
	"github.com/Jumpaku/cyamli/v2/generate/php"
	"github.com/Jumpaku/cyamli/v2/generate/python3"
	"github.com/Jumpaku/cyamli/v2/generate/typescript"
	"github.com/Jumpaku/cyamli/v2/schema"
	"github.com/samber/lo"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//go:generate go run github.com/Jumpaku/cyamli/v2/cmd/cyamli generate golang -schema-path=cli.cyamli.yaml -out-path=cli.gen.go
func main() {
	if err := Run(cli{}, os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type cli struct {
}

var _ CLIHandler = cli{}

func (c cli) Run(input Input) error {
	help(true, input.Subcommand)

	return nil
}

func (c cli) Run_Version(Input_Version) error {
	fmt.Println(GetVersion())
	return nil
}

func (c cli) Run_Generate(input Input_Generate) error {
	help(true, input.Subcommand)

	return nil
}

func (c cli) Run_GenerateGolang(input Input_GenerateGolang) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := golang.Generate(s, input.Opt_Package, generator(), w); err != nil {
		return fmt.Errorf("fail to generate Golang code: %w", err)
	}

	return nil
}

func (c cli) Run_GenerateDart3(input Input_GenerateDart3) (err error) {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := dart3.Generate(s, generator(), w); err != nil {
		return fmt.Errorf("fail to generate Dart3 code: %w", err)
	}

	return nil
}

func (c cli) Run_GenerateKotlin(input Input_GenerateKotlin) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := kotlin.Generate(s, input.Opt_Package, generator(), w); err != nil {
		return fmt.Errorf("fail to generate Kotlin code: %w", err)
	}

	return nil
}

func (c cli) Run_GeneratePython3(input Input_GeneratePython3) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := python3.Generate(s, generator(), w); err != nil {
		return fmt.Errorf("fail to generate Python3 code: %w", err)
	}

	return nil
}

func (c cli) Run_GenerateTypescript(input Input_GenerateTypescript) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := typescript.Generate(s, generator(), w); err != nil {
		return fmt.Errorf("fail to generate Typescript code: %w", err)
	}

	return nil
}

func (c cli) Run_GenerateCpp(input Input_GenerateCpp) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	if input.Opt_IncludeHeader == "" {
		return fmt.Errorf("include header path is required")
	}
	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}
	{
		w, err := outputWriter(input.Opt_OutHeaderPath)
		if err != nil {
			return fmt.Errorf("fail to create output writer: %w", err)
		}
		defer w.Close()

		if err := cpp.GenerateH(s, input.Opt_Namespace, generator(), w); err != nil {
			return fmt.Errorf("fail to generate C++ header file: %w", err)
		}
	}
	{
		w, err := outputWriter(input.Opt_OutSourcePath)
		if err != nil {
			return fmt.Errorf("fail to create output writer: %w", err)
		}
		defer w.Close()

		if err := cpp.GenerateCpp(s, input.Opt_IncludeHeader, input.Opt_Namespace, generator(), w); err != nil {
			return fmt.Errorf("fail to generate C++ source file: %w", err)
		}
	}

	return nil
}

func (c cli) Run_GenerateCsharp(input Input_GenerateCsharp) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := csharp.Generate(s, input.Opt_Namespace, generator(), w); err != nil {
		return fmt.Errorf("fail to generate Golang code: %w", err)
	}

	return nil
}

type phpNamedWriter struct {
	OutDir string
}

func (w phpNamedWriter) Write(name string, b []byte) (int, error) {
	if w.OutDir != "" {
		if err := os.MkdirAll(w.OutDir, 0755); err != nil {
			return 0, fmt.Errorf("fail to create output directory %q: %w", w.OutDir, err)
		}
	}

	o, err := outputWriter(lo.Ternary(w.OutDir == "", "", filepath.Join(w.OutDir, name)))
	if err != nil {
		return 0, fmt.Errorf("fail to create file %q: %w", name, err)
	}
	defer o.Close()

	return o.Write(b)
}

func (c cli) Run_GeneratePhp(input Input_GeneratePhp) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}
	if err := php.Generate(s, input.Opt_Namespace, generator(), phpNamedWriter{OutDir: input.Opt_OutDir}); err != nil {
		return fmt.Errorf("fail to generate PHP code: %w", err)
	}
	return nil
}

func (c cli) Run_GenerateDocs(input Input_GenerateDocs) error {
	exitIfErrorMessage(input.ErrorMessage, input.Subcommand)
	help(input.Opt_Help, input.Subcommand)

	docFormat := docs.DocFormat(input.Opt_Format)
	switch docFormat {
	default:
		exitIfErrorMessage("unsupported doc format: "+input.Opt_Format, input.Subcommand)
	case docs.DocFormatHTML, docs.DocFormatMarkdown, docs.DocFormatText:
	}

	s, err := loadSchema(input.Opt_SchemaPath)
	if err != nil {
		return fmt.Errorf("fail to load schema: %w", err)
	}

	w, err := outputWriter(input.Opt_OutPath)
	if err != nil {
		return fmt.Errorf("fail to create output writer: %w", err)
	}
	defer w.Close()

	if err := docs.Generate(s, docFormat, w); err != nil {
		return fmt.Errorf("fail to generate Documentation code: %w", err)
	}

	return nil
}

func exitIfErrorMessage(errorMessage string, subcommand []string) {
	if errorMessage != "" {
		_, _ = fmt.Fprintf(os.Stderr, `error: cyamli %s
    %s
To see help Run
  $ cyamli %s -help`, strings.Join(subcommand, " "), errorMessage, strings.Join(subcommand, " "))
		os.Exit(1)
	}

}

func help(helpOption bool, subcommand []string) {
	if helpOption {
		fmt.Println(GetDoc(subcommand))
		os.Exit(0)
	}
}

func loadSchema(schemaPath string) (s schema.Schema, err error) {
	var buf []byte
	if schemaPath == "" {
		buf, err = io.ReadAll(os.Stdin)
	} else {
		buf, err = os.ReadFile(schemaPath)
		if err != nil {
			return schema.Schema{}, fmt.Errorf("fail to read schema file: %w", err)
		}
	}
	s, err = schema.Load(bytes.NewBuffer(buf))
	if err != nil {
		return schema.Schema{}, fmt.Errorf("fail to load schema: %w", err)
	}
	return s, nil
}

func outputWriter(outPath string) (io.WriteCloser, error) {
	if outPath == "" {
		return os.Stdout, nil
	}
	f, err := os.Create(outPath)
	if err != nil {
		return nil, fmt.Errorf("fail to create output file: %w", err)
	}
	return f, nil

}

func generator() string {
	return "cyamli"
}
