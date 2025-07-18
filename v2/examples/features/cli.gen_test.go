// Code generated by cyamli, DO NOT EDIT.
package features_test

import (
	"fmt"
	"strconv"
	"strings"
)

type CLIHandler interface {
	Run(input Input) error
	Run_Sub1(input Input_Sub1) error
	Run_Sub1Sub2(input Input_Sub1Sub2) error
	Run_Sub1Sub2Sub3(input Input_Sub1Sub2Sub3) error
}

func Run(handler CLIHandler, args []string) error {
	subcommandPath, options, arguments := resolveArgs(args)
	switch strings.Join(subcommandPath, " ") {
	case "":
		var input Input
		input.resolveInput(subcommandPath, options, arguments)
		return handler.Run(input)

	case "sub1":
		var input Input_Sub1
		input.resolveInput(subcommandPath, options, arguments)
		return handler.Run_Sub1(input)

	case "sub1 sub2":
		var input Input_Sub1Sub2
		input.resolveInput(subcommandPath, options, arguments)
		return handler.Run_Sub1Sub2(input)

	case "sub1 sub2 sub3":
		var input Input_Sub1Sub2Sub3
		input.resolveInput(subcommandPath, options, arguments)
		return handler.Run_Sub1Sub2Sub3(input)
	}
	return nil
}

type Input struct {
	Opt_NegationOption    bool
	Opt_Option            int64
	Opt_PropagationOption string
	Opt_RepeatableOption  string
	Arg_FirstArg          bool
	Arg_SecondArg         int64
	Arg_ThirdArg          []string
	Subcommand            []string
	Options               []string
	Arguments             []string

	ErrorMessage string
}

func (input *Input) resolveInput(subcommand, options, arguments []string) {
	*input = Input{Opt_NegationOption: false,
		Opt_Option:            123,
		Opt_PropagationOption: "",
		Opt_RepeatableOption:  "",
		Subcommand:            subcommand,
		Options:               options,
		Arguments:             arguments,
	}

	for _, arg := range input.Options {
		optName, lit, cut := strings.Cut(arg, "=")
		func(...any) {}(optName, lit, cut)

		switch optName {
		case "-negation-option":
			if !cut {
				lit = "true"
			}
			if v, err := parseValue("bool", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_NegationOption = v.(bool)
			}
		case "-no-negation-option":
			if !cut {
				lit = "true"
			}
			if v, err := parseValue("bool", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_NegationOption = !v.(bool)
			}

		case "-option", "-o":
			if !cut {
				input.ErrorMessage = fmt.Sprintf("value is not specified to option %q", optName)
				return
			}
			if v, err := parseValue("int64", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_Option = v.(int64)
			}

		case "-propagation-option":
			if !cut {
				input.ErrorMessage = fmt.Sprintf("value is not specified to option %q", optName)
				return
			}
			if v, err := parseValue("string", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_PropagationOption = v.(string)
			}

		case "-repeatable-option":
			if !cut {
				input.ErrorMessage = fmt.Sprintf("value is not specified to option %q", optName)
				return
			}
			if v, err := parseValue("string", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_RepeatableOption = v.(string)
			}

		default:
			input.ErrorMessage = fmt.Sprintf("unknown option %q", optName)
			return
		}
	}

	expectedArgs := 3
	func(...any) {}(expectedArgs)
	if len(input.Arguments) <= 0 {
		input.ErrorMessage = fmt.Sprintf("too few arguments: required %d, got %d", expectedArgs, len(input.Arguments))
		return
	}
	if v, err := parseValue("bool", input.Arguments[0:]...); err != nil {
		input.ErrorMessage = fmt.Sprintf("value %q is not assignable to argument at [%d]", input.Arguments[0], 0)
		return
	} else {
		input.Arg_FirstArg = v.(bool)
	}

	if len(input.Arguments) <= 1 {
		input.ErrorMessage = fmt.Sprintf("too few arguments: required %d, got %d", expectedArgs, len(input.Arguments))
		return
	}
	if v, err := parseValue("int64", input.Arguments[1:]...); err != nil {
		input.ErrorMessage = fmt.Sprintf("value %q is not assignable to argument at [%d]", input.Arguments[1], 1)
		return
	} else {
		input.Arg_SecondArg = v.(int64)
	}

	if len(input.Arguments) < 2 {
		input.ErrorMessage = fmt.Sprintf("too few arguments: required at least %d, got %d", expectedArgs-1, len(input.Arguments))
		return
	}

	if v, err := parseValue("[]string", input.Arguments[2:]...); err != nil {
		input.ErrorMessage = fmt.Sprintf("values [%s] are not assignable to arguments at [%d:]", strings.Join(input.Arguments[2:], " "), 2)
		return
	} else {
		input.Arg_ThirdArg = v.([]string)
	}
}

type Input_Sub1 struct {
	Opt_PropagationOption string
	Subcommand            []string
	Options               []string
	Arguments             []string

	ErrorMessage string
}

func (input *Input_Sub1) resolveInput(subcommand, options, arguments []string) {
	*input = Input_Sub1{Opt_PropagationOption: "",
		Subcommand: subcommand,
		Options:    options,
		Arguments:  arguments,
	}

	for _, arg := range input.Options {
		optName, lit, cut := strings.Cut(arg, "=")
		func(...any) {}(optName, lit, cut)

		switch optName {
		case "-propagation-option":
			if !cut {
				input.ErrorMessage = fmt.Sprintf("value is not specified to option %q", optName)
				return
			}
			if v, err := parseValue("string", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_PropagationOption = v.(string)
			}

		default:
			input.ErrorMessage = fmt.Sprintf("unknown option %q", optName)
			return
		}
	}

	expectedArgs := 0
	func(...any) {}(expectedArgs)
}

type Input_Sub1Sub2 struct {
	Opt_PropagationOption string
	Subcommand            []string
	Options               []string
	Arguments             []string

	ErrorMessage string
}

func (input *Input_Sub1Sub2) resolveInput(subcommand, options, arguments []string) {
	*input = Input_Sub1Sub2{Opt_PropagationOption: "",
		Subcommand: subcommand,
		Options:    options,
		Arguments:  arguments,
	}

	for _, arg := range input.Options {
		optName, lit, cut := strings.Cut(arg, "=")
		func(...any) {}(optName, lit, cut)

		switch optName {
		case "-propagation-option":
			if !cut {
				input.ErrorMessage = fmt.Sprintf("value is not specified to option %q", optName)
				return
			}
			if v, err := parseValue("string", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_PropagationOption = v.(string)
			}

		default:
			input.ErrorMessage = fmt.Sprintf("unknown option %q", optName)
			return
		}
	}

	expectedArgs := 0
	func(...any) {}(expectedArgs)
}

type Input_Sub1Sub2Sub3 struct {
	Opt_PropagationOption string
	Subcommand            []string
	Options               []string
	Arguments             []string

	ErrorMessage string
}

func (input *Input_Sub1Sub2Sub3) resolveInput(subcommand, options, arguments []string) {
	*input = Input_Sub1Sub2Sub3{Opt_PropagationOption: "",
		Subcommand: subcommand,
		Options:    options,
		Arguments:  arguments,
	}

	for _, arg := range input.Options {
		optName, lit, cut := strings.Cut(arg, "=")
		func(...any) {}(optName, lit, cut)

		switch optName {
		case "-propagation-option":
			if !cut {
				input.ErrorMessage = fmt.Sprintf("value is not specified to option %q", optName)
				return
			}
			if v, err := parseValue("string", lit); err != nil {
				input.ErrorMessage = fmt.Sprintf("value %q is not assignable to option %q", lit, optName)
				return
			} else {
				input.Opt_PropagationOption = v.(string)
			}

		default:
			input.ErrorMessage = fmt.Sprintf("unknown option %q", optName)
			return
		}
	}

	expectedArgs := 0
	func(...any) {}(expectedArgs)
}
func resolveArgs(args []string) (subcommandPath []string, options []string, arguments []string) {
	if len(args) == 0 {
		panic("command line arguments are too few")
	}
	subcommandSet := map[string]bool{
		"": true, "sub1": true, "sub1 sub2": true, "sub1 sub2 sub3": true,
	}

	subcommandPath, options, arguments = []string{}, []string{}, []string{}
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

	restArgs := args[1+len(subcommandPath):]
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if strings.HasPrefix(arg, "-") {
			options = append(options, arg)
		} else {
			arguments = append(arguments, arg)
		}
	}

	return subcommandPath, options, arguments
}

func parseValue(typ string, strValue ...string) (dst any, err error) {
	switch typ {
	case "[]bool":
		val := make([]bool, len(strValue))
		for idx, str := range strValue {
			var v any
			if v, err = parseValue("bool", str); err != nil {
				return nil, fmt.Errorf("fail to parse %#v as []bool: %w", str, err)
			}
			val[idx] = v.(bool)
		}
		return val, nil
	case "[]int64":
		val := make([]int64, len(strValue))
		for idx, str := range strValue {
			var v any
			if v, err = parseValue("int64", str); err != nil {
				return nil, fmt.Errorf("fail to parse %#v as []int64: %w", str, err)
			}
			val[idx] = v.(int64)
		}
		return val, nil
	case "[]string":
		val := make([]string, len(strValue))
		for idx, str := range strValue {
			var v any
			if v, err = parseValue("string", str); err != nil {
				return nil, fmt.Errorf("fail to parse %#v as []string: %w", str, err)
			}
			val[idx] = v.(string)
		}
		return val, nil
	case "bool":
		switch strings.ToLower(strValue[0]) {
		default:
			return nil, fmt.Errorf("fail to parse %q as bool: unknown value", strValue[0])
		case "true", "1", "t":
			return true, nil
		case "false", "0", "f":
			return false, nil
		}
	case "int64":
		val, err := strconv.ParseInt(strValue[0], 0, 64)
		if err != nil {
			return nil, fmt.Errorf("fail to parse %q as int64: %w", strValue[0], err)
		}
		return val, nil
	case "string":
		return strValue[0], nil
	}

	return nil, fmt.Errorf("unknown type %q", typ)
}

func GetVersion() string {
	return "1.2.3"
}
func GetProgram() string {
	return "features"
}
func GetDoc(subcommands []string) string {
	switch strings.Join(subcommands, " ") {
	case "":
		return "features \n\n    Description:\n        This is root command, which is a command with name and version.\n\n    Syntax:\n        $ features  [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -negation-option[=<boolean>]  (default=false),\n        -no-negation-option[=<boolean>]:\n            this option's negated version `-no-negation-option` can be available.\n\n        -option=<integer>, -o=<integer>  (default=123):\n            option can have:\n              a description,\n              a type of string, integer, or boolean,\n              a short name,\n              and a default value.\n\n        -propagation-option=<string>  (default=\"\"):\n            this option is available with the descendant commands.\n\n        -repeatable-option=<string>  (default=\"\"):\n            this option can be repeated multiple times.\n\n    Arguments:\n        1.  <first_arg:boolean>\n            first argument with type boolean\n\n        2.  <second_arg:integer>\n            second argument with type boolean\n\n        3. [<third_arg:string>]...\n            third argument, which can take multiple values.\n\n    Subcommands:\n        sub1:\n            this is a child command.\n\n\n"

	case "sub1":
		return "features sub1\n\n    Description:\n        this is a child command.\n\n    Syntax:\n        $ features sub1 [<option>]...\n\n    Options:\n        -propagation-option=<string>  (default=\"\"):\n            this option is available with the descendant commands.\n\n    Subcommands:\n        sub2:\n            this is a grandchild command.\n\n\n"

	case "sub1 sub2":
		return "features sub1 sub2\n\n    Description:\n        this is a grandchild command.\n\n    Syntax:\n        $ features sub1 sub2 [<option>]...\n\n    Options:\n        -propagation-option=<string>  (default=\"\"):\n            this option is available with the descendant commands.\n\n    Subcommands:\n        sub3:\n            this is a great-grandchild command.\n\n\n"

	case "sub1 sub2 sub3":
		return "features sub1 sub2 sub3\n\n    Description:\n        this is a great-grandchild command.\n\n    Syntax:\n        $ features sub1 sub2 sub3 [<option>]...\n\n    Options:\n        -propagation-option=<string>  (default=\"\"):\n            this option is available with the descendant commands.\n\n\n"
	default:
		panic(fmt.Sprintf(`invalid subcommands: %v`, subcommands))
	}
}
