package golang

import (
	"cliautor"
	"fmt"
	"reflect"
	"strings"

	"github.com/samber/lo"
)

type InterpretResult[Input any] struct {
	SubcommandPath []string
	Input          Input
}

func InterpretSubcommand(s *cliautor.Schema, args []string) (subcommand *cliautor.Command, subcommandPath []string, restArgs []string) {
	cmd := s.Program.Command()

	// Extract subcommand path
	for _, arg := range args {
		var ok bool
		cmd, ok = cmd.Subcommands[arg]
		if !ok {
			break
		}
		subcommandPath = append(subcommandPath, arg)
	}

	return cmd, subcommandPath, args[len(subcommandPath):]
}
func InterpretInput[Input any](cmd *cliautor.Command, restArgs []string) (result Input, err error) {
	// initialize Options with default values
	inputRV := reflect.Indirect(reflect.ValueOf(&result))
	for optName, opt := range cmd.Options {
		field := inputRV.FieldByName(InputOptFieldName(optName))
		if opt.Default != "" {
			v, err := parseGoValue(opt.Type, false, opt.Default)
			if err != nil {
				return result, fmt.Errorf("fail to set %q as default value for option %q", opt.Default, optName)
			}
			field.Set(reflect.ValueOf(v))
		}
	}

	// interpret remaining command line arguments
	optionArgs := []string{}
	positionalArgs := []string{}
	for _, arg := range restArgs {
		if strings.HasPrefix(arg, "-") {
			optionArgs = append(optionArgs, arg)
		} else {
			positionalArgs = append(positionalArgs, arg)
		}
	}

	for _, arg := range optionArgs {
		resolved, err := resolveOpt(cmd, arg)
		if err != nil {
			return result, fmt.Errorf("fail to interpret command line arguments: %w", err)
		}

		field := inputRV.FieldByName(InputOptFieldName(resolved.Name))
		if err != nil {
			return result, fmt.Errorf("fail to set value %v to option %q", resolved.Value, resolved.Name)
		}

		field.Set(reflect.ValueOf(resolved.Value))
	}
	for idx := range positionalArgs {
		resolved, err := resolveArg(cmd, positionalArgs, idx)
		if err != nil {
			return result, fmt.Errorf("fail to interpret command line arguments: %w", err)
		}

		field := inputRV.FieldByName(InputArgFieldName(resolved.Name))
		if err != nil {
			return result, fmt.Errorf("fail to set value %v to positional argument %q", resolved.Value, resolved.Name)
		}

		field.Set(reflect.ValueOf(resolved.Value))
	}

	return result, nil
}

type InterpretedValue struct {
	Name  string
	Value any
}

func resolveOpt(cmd *cliautor.Command, arg string) (val InterpretedValue, err error) {
	for optName, opt := range cmd.Options {
		name, lit, cut := strings.Cut(arg, "=")
		if optName == name || opt.Short == name {
			val := InterpretedValue{Name: optName}
			if !cut {
				if opt.Type == cliautor.TypeBoolean {
					lit = "false"
				} else {
					return val, fmt.Errorf("fail to set value for option %q", optName)
				}
			}
			val.Value, err = parseGoValue(opt.Type, false, lit)
			if err != nil {
				return val, fmt.Errorf("fail to set %q to option %q", lit, optName)
			}
			return val, nil
		}
	}
	return val, fmt.Errorf("fail to resolve specified option %q", arg)
}

func resolveArg(cmd *cliautor.Command, args []string, at int) (val InterpretedValue, err error) {
	if at >= len(cmd.Arguments) {
		return InterpretedValue{}, fmt.Errorf("too many positional arguments")
	}
	val = InterpretedValue{Name: cmd.Arguments[at].Name}
	val.Value, err = parseGoValue(cmd.Arguments[at].Type, cmd.Arguments[at].Variadic, args[at], args[at:]...)
	if err != nil {
		return val, fmt.Errorf("fail to resolve positional arguments")
	}

	return val, nil
}

func parseGoValue(typ cliautor.Type, variadic bool, str string, variadicStr ...string) (any, error) {
	if variadic {
		switch typ {
		case cliautor.TypeBoolean:

		case cliautor.TypeFloat:
		case cliautor.TypeInteger:
		default:
		}
	} else {
		switch typ {
		case cliautor.TypeBoolean:
		case cliautor.TypeFloat:
		case cliautor.TypeInteger:
		default:
		}
	}
	return nil, nil
}

func InputOptFieldName(string) string { return "" }
func InputArgFieldName(string) string { return "" }

func PathFullIdentifier(path []string) string {
	if len(path) == 0 {
		return "Cmd"
	}
	titled := lo.Map(path, func(p string, i int) string {
		return Title(strings.Join(cliautor.MakeName(p), ""))
	})
	return "Cmd_" + strings.Join(titled, "")
}
