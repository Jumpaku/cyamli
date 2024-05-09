// Code generated by cyamli v0.0.15, DO NOT EDIT.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Func[Input any] func(subcommand []string, input Input, inputErr error) (err error)

type CLI struct {
	Generate CLI_Generate

	List CLI_List

	FUNC Func[CLI_Input]
}

func (CLI) DESC_Simple() string {
	return "cyamli (v0.0.15):\nA command line tool to generate CLI for your app from YAML-based schema.\n\nUsage:\n    $ cyamli [<option>]...\n\nOptions:\n    -help, -version\n\nSubcommands:\n    generate, list\n\n"
}
func (CLI) DESC_Detail() string {
	return "cyamli (v0.0.15):\nA command line tool to generate CLI for your app from YAML-based schema.\n\nUsage:\n    $ cyamli [<option>]...\n\n\nOptions:\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        shows description of this app\n\n    -version[=<boolean>], -v[=<boolean>]  (default=false):\n        shows version of this app\n\n\nSubcommands:\n    generate:\n\n    list:\n        shows subcommands\n\n"
}

type CLI_Input struct {
	Opt_Help bool

	Opt_Version bool
}

func resolve_CLI_Input(input *CLI_Input, restArgs []string) error {
	*input = CLI_Input{

		Opt_Help: false,

		Opt_Version: false,
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-version", "-v":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Version, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	return nil
}

type CLI_Generate struct {
	Docs CLI_GenerateDocs

	Golang CLI_GenerateGolang

	Python3 CLI_GeneratePython3

	FUNC Func[CLI_Generate_Input]
}

func (CLI_Generate) DESC_Simple() string {
	return "\nUsage:\n    $ <program> generate\n\nSubcommands:\n    docs, golang, python3\n\n"
}
func (CLI_Generate) DESC_Detail() string {
	return "\nUsage:\n    $ <program> generate\n\n\nSubcommands:\n    docs:\n        generates documentation for your CLI app.\n\n    golang:\n        generates CLI for your app written in Go.\n\n    python3:\n        generates CLI for your app written in Python3.\n\n"
}

type CLI_Generate_Input struct {
}

func resolve_CLI_Generate_Input(input *CLI_Generate_Input, restArgs []string) error {
	*input = CLI_Generate_Input{}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		}
	}

	return nil
}

type CLI_GenerateDocs struct {
	FUNC Func[CLI_GenerateDocs_Input]
}

func (CLI_GenerateDocs) DESC_Simple() string {
	return "generates documentation for your CLI app.\n\nUsage:\n    $ <program> generate docs [<option>|<argument>]... [-- [<argument>]...]\n\nOptions:\n    -all, -format, -help, -out-path, -schema-path\n\nArguments:\n    <subcommands>...\n\n"
}
func (CLI_GenerateDocs) DESC_Detail() string {
	return "generates documentation for your CLI app.\n\nUsage:\n    $ <program> generate docs [<option>|<argument>]... [-- [<argument>]...]\n\n\nOptions:\n    -all[=<boolean>], -a[=<boolean>]  (default=false):\n        if specified then outputs documentation for all subcommands, otherwise in text format.\n\n    -format=<string>, -f=<string>  (default=\"text\"):\n        specifies output format of the documentation in text or markdown\n\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        shows description of docs subcommand\n\n    -out-path=<string>  (default=\"\"):\n        if specified then creates a file at the path and writes generated documentation, otherwise outputs to stdout.\n\n    -schema-path=<string>  (default=\"\"):\n        if specified then reads schema file from the path, otherwise reads from stdin.\n\n\nArguments:\n    [0:] [<subcommands:string>]...\n        selects subcommand for which the documentation is output.\n\n"
}

type CLI_GenerateDocs_Input struct {
	Opt_All bool

	Opt_Format string

	Opt_Help bool

	Opt_OutPath string

	Opt_SchemaPath string

	Arg_Subcommands []string
}

func resolve_CLI_GenerateDocs_Input(input *CLI_GenerateDocs_Input, restArgs []string) error {
	*input = CLI_GenerateDocs_Input{

		Opt_All: false,

		Opt_Format: "text",

		Opt_Help: false,

		Opt_OutPath: "",

		Opt_SchemaPath: "",
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-all", "-a":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_All, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-format", "-f":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_Format, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-out-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OutPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-schema-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_SchemaPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_Subcommands, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
	}

	return nil
}

type CLI_GenerateGolang struct {
	FUNC Func[CLI_GenerateGolang_Input]
}

func (CLI_GenerateGolang) DESC_Simple() string {
	return "generates CLI for your app written in Go.\n\nUsage:\n    $ <program> generate golang [<option>]...\n\nOptions:\n    -help, -out-path, -package, -schema-path\n\n"
}
func (CLI_GenerateGolang) DESC_Detail() string {
	return "generates CLI for your app written in Go.\n\nUsage:\n    $ <program> generate golang [<option>]...\n\n\nOptions:\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        shows description of golang subcommand\n\n    -out-path=<string>  (default=\"\"):\n        if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.\n\n    -package=<string>  (default=\"main\"):\n        package name where the generated file will be placed.\n\n    -schema-path=<string>  (default=\"\"):\n        if specified then reads schema file from the path, otherwise reads from stdin.\n\n"
}

type CLI_GenerateGolang_Input struct {
	Opt_Help bool

	Opt_OutPath string

	Opt_Package string

	Opt_SchemaPath string
}

func resolve_CLI_GenerateGolang_Input(input *CLI_GenerateGolang_Input, restArgs []string) error {
	*input = CLI_GenerateGolang_Input{

		Opt_Help: false,

		Opt_OutPath: "",

		Opt_Package: "main",

		Opt_SchemaPath: "",
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-out-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OutPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-package":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_Package, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-schema-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_SchemaPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	return nil
}

type CLI_GeneratePython3 struct {
	FUNC Func[CLI_GeneratePython3_Input]
}

func (CLI_GeneratePython3) DESC_Simple() string {
	return "generates CLI for your app written in Python3.\n\nUsage:\n    $ <program> generate python3 [<option>]...\n\nOptions:\n    -help, -out-path, -schema-path\n\n"
}
func (CLI_GeneratePython3) DESC_Detail() string {
	return "generates CLI for your app written in Python3.\n\nUsage:\n    $ <program> generate python3 [<option>]...\n\n\nOptions:\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        shows description of python3 subcommand\n\n    -out-path=<string>  (default=\"\"):\n        if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.\n\n    -schema-path=<string>  (default=\"\"):\n        if specified then reads schema file from the path, otherwise reads from stdin.\n\n"
}

type CLI_GeneratePython3_Input struct {
	Opt_Help bool

	Opt_OutPath string

	Opt_SchemaPath string
}

func resolve_CLI_GeneratePython3_Input(input *CLI_GeneratePython3_Input, restArgs []string) error {
	*input = CLI_GeneratePython3_Input{

		Opt_Help: false,

		Opt_OutPath: "",

		Opt_SchemaPath: "",
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-out-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OutPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-schema-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_SchemaPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	return nil
}

type CLI_List struct {
	FUNC Func[CLI_List_Input]
}

func (CLI_List) DESC_Simple() string {
	return "shows subcommands\n\nUsage:\n    $ <program> list [<option>]...\n\nOptions:\n    -schema-path\n\n"
}
func (CLI_List) DESC_Detail() string {
	return "shows subcommands\n\nUsage:\n    $ <program> list [<option>]...\n\n\nOptions:\n    -schema-path=<string>  (default=\"\"):\n        if specified then reads schema file from the path, otherwise reads from stdin.\n\n"
}

type CLI_List_Input struct {
	Opt_SchemaPath string
}

func resolve_CLI_List_Input(input *CLI_List_Input, restArgs []string) error {
	*input = CLI_List_Input{

		Opt_SchemaPath: "",
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-schema-path":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_SchemaPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	return nil
}

func NewCLI() CLI {
	return CLI{}
}

func Run(cli CLI, args []string) error {
	subcommandPath, restArgs := resolveSubcommand(args)
	switch strings.Join(subcommandPath, " ") {

	case "":
		funcMethod := cli.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.FUNC not assigned", "")
		}
		var input CLI_Input
		err := resolve_CLI_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "generate":
		funcMethod := cli.Generate.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Generate.FUNC not assigned", "generate")
		}
		var input CLI_Generate_Input
		err := resolve_CLI_Generate_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "generate docs":
		funcMethod := cli.Generate.Docs.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Generate.Docs.FUNC not assigned", "generate docs")
		}
		var input CLI_GenerateDocs_Input
		err := resolve_CLI_GenerateDocs_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "generate golang":
		funcMethod := cli.Generate.Golang.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Generate.Golang.FUNC not assigned", "generate golang")
		}
		var input CLI_GenerateGolang_Input
		err := resolve_CLI_GenerateGolang_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "generate python3":
		funcMethod := cli.Generate.Python3.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Generate.Python3.FUNC not assigned", "generate python3")
		}
		var input CLI_GeneratePython3_Input
		err := resolve_CLI_GeneratePython3_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "list":
		funcMethod := cli.List.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.List.FUNC not assigned", "list")
		}
		var input CLI_List_Input
		err := resolve_CLI_List_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	}
	return nil
}

func resolveSubcommand(args []string) (subcommandPath []string, restArgs []string) {
	if len(args) == 0 {
		panic("command line arguments are too few")
	}
	subcommandSet := map[string]bool{
		"":         true,
		"generate": true, "generate docs": true, "generate golang": true, "generate python3": true, "list": true,
	}

	for _, arg := range args[1:] {
		if arg == "--" {
			break
		}
		pathLiteral := strings.Join(append(append([]string{}, subcommandPath...), arg), " ")
		if !subcommandSet[pathLiteral] {
			break
		}
		subcommandPath = append(subcommandPath, arg)
	}

	return subcommandPath, args[1+len(subcommandPath):]
}

func parseValue(dstPtr any, strValue ...string) error {
	switch dstPtr := dstPtr.(type) {
	case *[]bool:
		val := make([]bool, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []bool: %w", str, err)
			}
		}
		*dstPtr = val
	case *[]float64:
		val := make([]float64, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []float64: %w", str, err)
			}
		}
		*dstPtr = val
	case *[]int64:
		val := make([]int64, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []int64: %w", str, err)
			}
		}
		*dstPtr = val
	case *[]string:
		val := make([]string, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []string: %w", str, err)
			}
		}
		*dstPtr = val
	case *bool:
		val, err := strconv.ParseBool(strValue[0])
		if err != nil {
			return fmt.Errorf("fail to parse %q as bool: %w", strValue[0], err)
		}
		*dstPtr = val
	case *float64:
		val, err := strconv.ParseFloat(strValue[0], 64)
		if err != nil {
			return fmt.Errorf("fail to parse %q as float64: %w", strValue[0], err)
		}
		*dstPtr = val
	case *int64:
		val, err := strconv.ParseInt(strValue[0], 0, 64)
		if err != nil {
			return fmt.Errorf("fail to parse %q as int64: %w", strValue[0], err)
		}
		*dstPtr = val
	case *string:
		*dstPtr = strValue[0]
	}

	return nil
}

func consumeVariables(...any) {}
