// Code generated by github.com/Jumpaku/cyamli v1.1.5, DO NOT EDIT.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Func[Input any] func(subcommand []string, input Input, inputErr error) (err error)

type CLI struct {
	Fetch CLI_Fetch

	List CLI_List

	FUNC Func[CLI_Input]
}
type CLI_Input struct {
}

func resolve_CLI_Input(input *CLI_Input, restArgs []string) error {
	*input = CLI_Input{}

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

type CLI_Fetch struct {
	FUNC Func[CLI_Fetch_Input]
}
type CLI_Fetch_Input struct {
	Opt_Config string

	Opt_Verbose bool

	Arg_Tables []string
}

func resolve_CLI_Fetch_Input(input *CLI_Fetch_Input, restArgs []string) error {
	*input = CLI_Fetch_Input{

		Opt_Config: "",

		Opt_Verbose: false,
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

		case "-config", "-c":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_Config, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-verbose", "-v":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Verbose, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_Tables, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
	}

	return nil
}

type CLI_List struct {
	FUNC Func[CLI_List_Input]
}
type CLI_List_Input struct {
	Opt_Config string
}

func resolve_CLI_List_Input(input *CLI_List_Input, restArgs []string) error {
	*input = CLI_List_Input{

		Opt_Config: "",
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

		case "-config", "-c":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_Config, lit); err != nil {
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

	case "fetch":
		funcMethod := cli.Fetch.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Fetch.FUNC not assigned", "fetch")
		}
		var input CLI_Fetch_Input
		err := resolve_CLI_Fetch_Input(&input, restArgs)
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
		"":      true,
		"fetch": true, "list": true,
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

func GetDoc(subcommands []string) string {
	switch strings.Join(subcommands, " ") {
	default:
		panic(fmt.Sprintf(`invalid subcommands: %v`, subcommands))

	case "":
		return "demo\n\ndemo\n\n    Description:\n        demo app to get table information from databases\n\n    Syntax:\n        $ demo \n\n    Subcommands:\n        fetch:\n            show information of tables\n\n        list:\n            list tables\n\n\n"

	case "fetch":
		return "demo\n\ndemo fetch\n\n    Description:\n        show information of tables\n\n    Syntax:\n        $ demo fetch [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -config=<string>, -c=<string>  (default=\"\"):\n            path to config file\n\n        -verbose[=<boolean>], -v[=<boolean>]  (default=false):\n            shows detailed log\n\n    Arguments:\n        1. [<tables:string>]...\n            names of tables to be described\n\n\n"

	case "list":
		return "demo\n\ndemo list\n\n    Description:\n        list tables\n\n    Syntax:\n        $ demo list [<option>]...\n\n    Options:\n        -config=<string>, -c=<string>  (default=\"\"):\n            path to config file\n\n\n"

	}
}
