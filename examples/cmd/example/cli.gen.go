// Code generated by github.com/Jumpaku/cyamli v1.1.6, DO NOT EDIT.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Func[Input any] func(subcommand []string, input Input, inputErr error) (err error)

type CLI struct {
	Sub1 CLI_Sub1

	Sub2 CLI_Sub2

	Sub3 CLI_Sub3

	FUNC Func[CLI_Input]
}
type CLI_Input struct {
	Opt_OptionA string

	Opt_OptionB int64

	Opt_OptionC bool

	Opt_OptionD float64

	Opt_OptionE string

	Arg_ArgA string

	Arg_ArgB int64

	Arg_ArgC bool

	Arg_ArgD float64

	Arg_ArgE string

	Arg_ArgV []string
}

func resolve_CLI_Input(input *CLI_Input, restArgs []string) error {
	*input = CLI_Input{

		Opt_OptionA: "abc",

		Opt_OptionB: int64(-123),

		Opt_OptionC: true,

		Opt_OptionD: float64(-123.456),

		Opt_OptionE: "",
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

		case "-option-a", "-a":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionA, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-b", "-b":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionB, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-c", "-c":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_OptionC, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-d", "-d":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionD, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-e":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionE, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	if len(arguments) <= 0 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgA, arguments[0]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 0)
	}

	if len(arguments) <= 1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgB, arguments[1]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 1)
	}

	if len(arguments) <= 2 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgC, arguments[2]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 2)
	}

	if len(arguments) <= 3 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgD, arguments[3]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 3)
	}

	if len(arguments) <= 4 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgE, arguments[4]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 4)
	}

	if len(arguments) <= 5-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgV, arguments[5:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[5:], " "), 5)
	}

	return nil
}

type CLI_Sub1 struct {
	FUNC Func[CLI_Sub1_Input]
}
type CLI_Sub1_Input struct {
}

func resolve_CLI_Sub1_Input(input *CLI_Sub1_Input, restArgs []string) error {
	*input = CLI_Sub1_Input{}

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

type CLI_Sub2 struct {
	FUNC Func[CLI_Sub2_Input]
}
type CLI_Sub2_Input struct {
}

func resolve_CLI_Sub2_Input(input *CLI_Sub2_Input, restArgs []string) error {
	*input = CLI_Sub2_Input{}

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

type CLI_Sub3 struct {
	Suba CLI_Sub3Suba

	Subb CLI_Sub3Subb

	Subc CLI_Sub3Subc

	Subd CLI_Sub3Subd

	FUNC Func[CLI_Sub3_Input]
}
type CLI_Sub3_Input struct {
	Opt_OptionA string

	Opt_OptionB int64

	Opt_OptionC bool

	Opt_OptionD float64

	Opt_OptionE string

	Arg_ArgA string

	Arg_ArgB int64

	Arg_ArgC bool

	Arg_ArgD float64

	Arg_ArgE string

	Arg_ArgV []string
}

func resolve_CLI_Sub3_Input(input *CLI_Sub3_Input, restArgs []string) error {
	*input = CLI_Sub3_Input{

		Opt_OptionA: "abc",

		Opt_OptionB: int64(-123),

		Opt_OptionC: true,

		Opt_OptionD: float64(-123.456),

		Opt_OptionE: "",
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

		case "-option-a", "-a":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionA, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-b", "-b":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionB, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-c", "-c":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_OptionC, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-d", "-d":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionD, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-option-e":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OptionE, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	if len(arguments) <= 0 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgA, arguments[0]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 0)
	}

	if len(arguments) <= 1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgB, arguments[1]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 1)
	}

	if len(arguments) <= 2 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgC, arguments[2]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 2)
	}

	if len(arguments) <= 3 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgD, arguments[3]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 3)
	}

	if len(arguments) <= 4 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgE, arguments[4]); err != nil {
		return fmt.Errorf("value is not assignable to argument at [%d]", 4)
	}

	if len(arguments) <= 5-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgV, arguments[5:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[5:], " "), 5)
	}

	return nil
}

type CLI_Sub3Suba struct {
	FUNC Func[CLI_Sub3Suba_Input]
}
type CLI_Sub3Suba_Input struct {
	Arg_ArgV []bool
}

func resolve_CLI_Sub3Suba_Input(input *CLI_Sub3Suba_Input, restArgs []string) error {
	*input = CLI_Sub3Suba_Input{}

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

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgV, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
	}

	return nil
}

type CLI_Sub3Subb struct {
	FUNC Func[CLI_Sub3Subb_Input]
}
type CLI_Sub3Subb_Input struct {
	Arg_ArgV []int64
}

func resolve_CLI_Sub3Subb_Input(input *CLI_Sub3Subb_Input, restArgs []string) error {
	*input = CLI_Sub3Subb_Input{}

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

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgV, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
	}

	return nil
}

type CLI_Sub3Subc struct {
	FUNC Func[CLI_Sub3Subc_Input]
}
type CLI_Sub3Subc_Input struct {
	Arg_ArgV []float64
}

func resolve_CLI_Sub3Subc_Input(input *CLI_Sub3Subc_Input, restArgs []string) error {
	*input = CLI_Sub3Subc_Input{}

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

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgV, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
	}

	return nil
}

type CLI_Sub3Subd struct {
	FUNC Func[CLI_Sub3Subd_Input]
}
type CLI_Sub3Subd_Input struct {
	Arg_ArgV []string
}

func resolve_CLI_Sub3Subd_Input(input *CLI_Sub3Subd_Input, restArgs []string) error {
	*input = CLI_Sub3Subd_Input{}

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

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_ArgV, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
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

	case "sub1":
		funcMethod := cli.Sub1.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub1.FUNC not assigned", "sub1")
		}
		var input CLI_Sub1_Input
		err := resolve_CLI_Sub1_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "sub2":
		funcMethod := cli.Sub2.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub2.FUNC not assigned", "sub2")
		}
		var input CLI_Sub2_Input
		err := resolve_CLI_Sub2_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "sub3":
		funcMethod := cli.Sub3.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub3.FUNC not assigned", "sub3")
		}
		var input CLI_Sub3_Input
		err := resolve_CLI_Sub3_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "sub3 suba":
		funcMethod := cli.Sub3.Suba.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub3.Suba.FUNC not assigned", "sub3 suba")
		}
		var input CLI_Sub3Suba_Input
		err := resolve_CLI_Sub3Suba_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "sub3 subb":
		funcMethod := cli.Sub3.Subb.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub3.Subb.FUNC not assigned", "sub3 subb")
		}
		var input CLI_Sub3Subb_Input
		err := resolve_CLI_Sub3Subb_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "sub3 subc":
		funcMethod := cli.Sub3.Subc.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub3.Subc.FUNC not assigned", "sub3 subc")
		}
		var input CLI_Sub3Subc_Input
		err := resolve_CLI_Sub3Subc_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "sub3 subd":
		funcMethod := cli.Sub3.Subd.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Sub3.Subd.FUNC not assigned", "sub3 subd")
		}
		var input CLI_Sub3Subd_Input
		err := resolve_CLI_Sub3Subd_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	}
	return nil
}

func resolveSubcommand(args []string) (subcommandPath []string, restArgs []string) {
	if len(args) == 0 {
		panic("command line arguments are too few")
	}
	subcommandSet := map[string]bool{
		"":     true,
		"sub1": true, "sub2": true, "sub3": true, "sub3 suba": true, "sub3 subb": true, "sub3 subc": true, "sub3 subd": true,
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
		return "example (v1.0.0)\n\nexample\n\n    Description:\n        this is an example command\n\n    Syntax:\n        $ example  [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -option-a=<string>, -a=<string>  (default=\"abc\"):\n            a - this is an option for root command\n\n        -option-b=<integer>, -b=<integer>  (default=-123):\n            b - this is an option for root command\n\n        -option-c[=<boolean>], -c[=<boolean>]  (default=true):\n            c - this is an option for root command\n\n        -option-d=<float>, -d=<float>  (default=-123.456):\n            d - this is an option for root command\n\n        -option-e=<string>  (default=\"\"):\n\n    Arguments:\n        1.  <arg_a:string>\n            a - this is an argument for root command\n\n        2.  <arg_b:integer>\n            b - this is an argument for root command\n\n        3.  <arg_c:boolean>\n            c - this is an argument for root command\n\n        4.  <arg_d:float>\n            d - this is an argument for root command\n\n        5.  <arg_e:string>\n\n        6. [<arg_v:string>]...\n            v - this is an argument for root command\n\n    Subcommands:\n        sub1:\n            1 - this is a sub command\n\n        sub2:\n            2 - this is a sub command\n\n        sub3:\n            3 - this is a sub command\n\n\n"

	case "sub1":
		return "example (v1.0.0)\n\nexample sub1\n\n    Description:\n        1 - this is a sub command\n\n    Syntax:\n        $ example sub1\n\n\n"

	case "sub2":
		return "example (v1.0.0)\n\nexample sub2\n\n    Description:\n        2 - this is a sub command\n\n    Syntax:\n        $ example sub2\n\n\n"

	case "sub3":
		return "example (v1.0.0)\n\nexample sub3\n\n    Description:\n        3 - this is a sub command\n\n    Syntax:\n        $ example sub3 [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -option-a=<string>, -a=<string>  (default=\"abc\"):\n            3 - a - this is an option for root command\n\n        -option-b=<integer>, -b=<integer>  (default=-123):\n            3 - b - this is an option for root command\n\n        -option-c[=<boolean>], -c[=<boolean>]  (default=true):\n            3 - c - this is an option for root command\n\n        -option-d=<float>, -d=<float>  (default=-123.456):\n            3 - d - this is an option for root command\n\n        -option-e=<string>  (default=\"\"):\n\n    Arguments:\n        1.  <arg_a:string>\n            3 - a - this is an argument for root command\n\n        2.  <arg_b:integer>\n            3 - b - this is an argument for root command\n\n        3.  <arg_c:boolean>\n            3 - c - this is an argument for root command\n\n        4.  <arg_d:float>\n            3 - d - this is an argument for root command\n\n        5.  <arg_e:string>\n\n        6. [<arg_v:string>]...\n            3 - v - this is an argument for root command\n\n    Subcommands:\n        suba:\n\n        subb:\n\n        subc:\n\n        subd:\n\n\n"

	case "sub3 suba":
		return "example (v1.0.0)\n\nexample sub3 suba\n\n    Syntax:\n        $ example sub3 suba [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:boolean>]...\n\n\n"

	case "sub3 subb":
		return "example (v1.0.0)\n\nexample sub3 subb\n\n    Syntax:\n        $ example sub3 subb [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:integer>]...\n\n\n"

	case "sub3 subc":
		return "example (v1.0.0)\n\nexample sub3 subc\n\n    Syntax:\n        $ example sub3 subc [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:float>]...\n\n\n"

	case "sub3 subd":
		return "example (v1.0.0)\n\nexample sub3 subd\n\n    Syntax:\n        $ example sub3 subd [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:string>]...\n\n\n"

	}
}
