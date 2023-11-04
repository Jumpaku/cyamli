package golang

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Jumpaku/cliautor/schema"

	"github.com/Jumpaku/cliautor/description"
	"github.com/Jumpaku/cliautor/golang/data"
	"github.com/Jumpaku/cliautor/name"

	"github.com/Jumpaku/go-assert"
)

func NewDefaultFunc[Input any](programName string) func(cmd *schema.Command, subcommand []string, input Input) (err error) {
	return func(cmd *schema.Command, subcommand []string, input Input) (err error) {
		fmt.Printf("subcommand: %q, input: %#v\n\n", strings.Join(subcommand, " "), input)

		descData := description.CreateCommandData(programName, subcommand, cmd)
		err = description.DescribeCommand(description.DetailExecutor(), descData, os.Stdout)
		if err != nil {
			return fmt.Errorf("fail to create command description: %w", err)
		}
		return nil
	}
}

func ResolveSubcommand(s *schema.Schema, args []string) (cmd *schema.Command, subcommand []string, restArgs []string) {
	assert.Params(len(args) >= 1, "first element of args must be a path of the executable file")
	cmd = s.Program.Command()

	// Extract subcommand path
	for _, arg := range args[1:] {
		if arg == "--" {
			break
		}
		found, ok := cmd.Subcommands[arg]
		if !ok {
			break
		}
		cmd = found
		subcommand = append(subcommand, arg)
	}

	return cmd, subcommand, args[1:][len(subcommand):]
}

func ResolveInput(schemaCommand *schema.Command, restArgs []string, inputPtr any) error {
	if err := ResolveOptions(schemaCommand.Options, restArgs, inputPtr); err != nil {
		return err
	}
	if err := ResolveArguments(schemaCommand.Arguments, restArgs, inputPtr); err != nil {
		return err
	}
	return nil
}

func ResolveOptions(schemaOptions map[string]*schema.Option, restArgs []string, inputPtr any) error {
	ptrRV := reflect.ValueOf(inputPtr)
	if ptrRV.Kind() != reflect.Pointer {
		return fmt.Errorf("inputPtr must be a not nil pointer to struct: %T", inputPtr)
	}

	structRV := reflect.Indirect(ptrRV)
	if structRV.Kind() != reflect.Struct {
		return fmt.Errorf("inputPtr must be a not nil pointer to struct: %T", inputPtr)
	}

	options := map[string]*schema.Option{}
	for optName, schemaOption := range schemaOptions {
		options[optName] = schemaOption
		if schemaOption.Short != "" {
			options[schemaOption.Short] = schemaOption
		}
	}

	nameMap := map[string]string{}
	for optName, schemaOption := range schemaOptions {
		nameMap[optName] = optName
		if schemaOption.Short != "" {
			nameMap[schemaOption.Short] = optName
		}
	}

	for _, arg := range restArgs {
		if arg == "--" {
			return nil
		}
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		option, ok := options[optName]
		if !ok {
			return fmt.Errorf("unknown option: %s", arg)
		}
		if !cut {
			if option.Type != schema.TypeBoolean {
				return fmt.Errorf("value must be assigned if option type is not boolean: %q", arg)
			}
			lit = "true"
		}

		fieldName := data.Option{Name: name.MakePath(nameMap[optName])}.InputFieldName()

		value, err := ParseGoValue(option.Type, false, lit)
		if err != nil {
			return fmt.Errorf("fail to parse option value: %q: %w", arg, err)
		}

		structRV.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}
	return nil
}

func ResolveArguments(arguments []*schema.Argument, restArgs []string, inputPtr any) error {
	ptrRV := reflect.ValueOf(inputPtr)
	if ptrRV.Kind() != reflect.Pointer {
		return fmt.Errorf("inputPtr must be a not nil pointer to struct: %T", inputPtr)
	}

	structRV := reflect.Indirect(ptrRV)
	if structRV.Kind() != reflect.Struct {
		return fmt.Errorf("inputPtr must be a not nil pointer to struct: %T", inputPtr)
	}

	var args []string
	for idx, arg := range restArgs {
		if arg == "--" {
			args = append(args, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		}
	}

	if len(arguments) == 0 {
		if len(args) > 0 {
			return fmt.Errorf("too many arguments")
		}
		return nil
	}
	if arguments[len(arguments)-1].Variadic && len(args) < len(arguments)-1 {
		return fmt.Errorf("too few arguments")
	}
	if !arguments[len(arguments)-1].Variadic && len(args) < len(arguments) {
		return fmt.Errorf("too few arguments")
	}
	if !arguments[len(arguments)-1].Variadic && len(args) > len(arguments) {
		return fmt.Errorf("too many arguments")
	}

	for idx, argument := range arguments {
		var value any
		var err error
		if argument.Variadic {
			value, err = ParseGoValue(argument.Type, true, args[idx:]...)
			if err != nil {
				return fmt.Errorf("fail to parse argument values: %s=%#v,position=[%d,%d): %w", argument.Name, args[idx:], idx, len(args), err)
			}
		} else {
			value, err = ParseGoValue(argument.Type, false, args[idx])
			if err != nil {
				return fmt.Errorf("fail to parse argument value: %s=%#v,position=[%d]: %w", argument.Name, args[idx], idx, err)
			}
		}

		fieldName := data.Argument{Name: name.MakePath(argument.Name)}.InputFieldName()

		structRV.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}
	return nil
}

func ParseGoValue(typ schema.Type, variadic bool, str ...string) (any, error) {
	if variadic {
		switch typ {
		default:
			return assert.Unexpected2[any, error]("unexpected type: %s", typ)
		case schema.TypeBoolean:
			val := []bool{}
			for _, str := range str {
				v, err := ParseGoValue(typ, false, str)
				if err != nil {
					return nil, fmt.Errorf("fail to parse %#v as []bool: %w", str, err)
				}
				val = append(val, v.(bool))
			}
			return val, nil
		case schema.TypeFloat:
			val := []float64{}
			for _, str := range str {
				v, err := ParseGoValue(typ, false, str)
				if err != nil {
					return nil, fmt.Errorf("fail to parse %#v as []bool: %w", str, err)
				}
				val = append(val, v.(float64))
			}
			return val, nil
		case schema.TypeInteger:
			val := []int64{}
			for _, str := range str {
				v, err := ParseGoValue(typ, false, str)
				if err != nil {
					return nil, fmt.Errorf("fail to parse %#v as []bool: %w", str, err)
				}
				val = append(val, v.(int64))
			}
			return val, nil
		case schema.TypeString, schema.TypeUnspecified:
			if str == nil {
				return []string{}, nil
			}
			return str, nil
		}
	} else {
		if len(str) != 1 {
			return nil, fmt.Errorf("single string must be provided if not variadic")
		}
		str := str[0]
		switch typ {
		default:
			return assert.Unexpected2[any, error]("unexpected type: %s", typ)
		case schema.TypeBoolean:
			val, err := strconv.ParseBool(str)
			if err != nil {
				return nil, fmt.Errorf("fail to parse %q as bool: %w", str, err)
			}
			return val, nil
		case schema.TypeFloat:
			val, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return nil, fmt.Errorf("fail to parse %q as float64: %w", str, err)
			}
			return val, nil
		case schema.TypeInteger:
			val, err := strconv.ParseInt(str, 0, 64)
			if err != nil {
				return nil, fmt.Errorf("fail to parse %q as int64: %w", str, err)
			}
			return val, nil
		case schema.TypeString, schema.TypeUnspecified:
			return str, nil
		}
	}
}
