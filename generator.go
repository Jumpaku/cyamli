package cliautor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Type string

const (
	TypeUnspecified Type = ""
	TypeString      Type = "string"
	TypeInteger     Type = "integer"
	TypeFloat       Type = "float"
	TypeBoolean     Type = "boolean"
)

type Opt struct {
	Short       string `yaml:"short"`
	Description string `yaml:"description"`
	Type        Type   `yaml:"type"`
	Default     string `yaml:"default"`
}
type Arg struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        Type   `yaml:"type"`
	Default     string `yaml:"default"`
	Variadic    bool   `yaml:"variadic"`
}
type Cmd struct {
	Description string          `yaml:"description"`
	Options     map[string]*Opt `yaml:"options"`
	Args        []*Arg          `yaml:"args"`
	Subcommand  map[string]*Cmd `yaml:"subcommands"`
}

func Load(b []byte) (*Generator, error) {
	program := Generator{}
	err := yaml.Unmarshal(b, &program)
	if err != nil {
		return nil, fmt.Errorf(`fail to parse yaml as command structure: %w`, err)
	}
	if err := (&program).Validate(); err != nil {
		return nil, fmt.Errorf(`fail to validate command: %w`, err)
	}
	return &program, nil
}

type Generator struct {
	Cmd
}

func (g *Generator) Validate() error {
	return validateCmd(&g.Cmd)
}

type Subcommand []string
type Input struct {
	BooleanVal bool
	FloatVal   float64
	IntegerVal int64
	StringVal  string
}

func (g *Generator) Decompose(args []string) (Subcommand, Input, error) {
	subcommand := Subcommand{}
	cmd := &g.Cmd
	for _, arg := range args {
		var ok bool
		cmd, ok = cmd.Subcommand[arg]
		if !ok {
			break
		}
		subcommand = append(subcommand, arg)
	}

	input := Input{}
	for _, argVal := range args[len(subcommand):] {
		arg, valStr, _ := strings.Cut(argVal, "=")
		option, ok := cmd.Options[strings.TrimPrefix(arg, "-")]
		if !ok {
			break
		}
		var err error
		switch option.Type {
		case TypeBoolean:
			input.BooleanVal, err = strconv.ParseBool(valStr)
			if err != nil {
				return nil, Input{}, fmt.Errorf(`fail to parse boolean option value: %s: %w`, arg, err)
			}
		case TypeFloat:
			input.FloatVal, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, Input{}, fmt.Errorf(`fail to parse float option value: %s: %w`, arg, err)
			}
		case TypeInteger:
			input.IntegerVal, err = strconv.ParseInt(valStr, 10, 64)
			if err != nil {
				return nil, Input{}, fmt.Errorf(`fail to parse integer option value: %s: %w`, arg, err)
			}
		default:
			input.StringVal = valStr
		}
	}

	return subcommand, input, nil
}

func validateCmd(cmd *Cmd) error {
	if len(cmd.Args) > 0 && len(cmd.Subcommand) > 0 {
		return errors.New("args and subcommands can be specified only one or none of them but not both")
	}
	for name, opt := range cmd.Options {
		if err := validateOpt(opt); err != nil {
			return fmt.Errorf("invalid option: %s: %w", name, err)
		}
	}
	for idx, arg := range cmd.Args {
		if err := validateArg(arg); err != nil {
			return fmt.Errorf("invalid arg: %w", err)
		}
		if idx < len(cmd.Args)-1 && arg.Variadic {
			return errors.New("variadic arg must be placed at last")
		}
	}
	for name, sub := range cmd.Subcommand {
		if name == "" {
			return errors.New("name of subcommand must be not empty")
		}
		if err := validateCmd(sub); err != nil {
			return fmt.Errorf("invalid subcommand: %s: %w", name, err)
		}
	}
	return nil
}

func validateOpt(opt *Opt) error {
	if err := validateType(opt.Type); err != nil {
		return fmt.Errorf("invalid type: %w", err)
	}
	if opt.Default != "" {
		switch opt.Type {
		case TypeBoolean:
			if _, err := strconv.ParseBool(opt.Default); err != nil {
				return fmt.Errorf(`fail to parse default value as boolean: %s: %w`, opt.Default, err)
			}
		case TypeFloat:
			if _, err := strconv.ParseFloat(opt.Default, 64); err != nil {
				return fmt.Errorf(`fail to parse default value as float: %s: %w`, opt.Default, err)
			}
		case TypeInteger:
			if _, err := strconv.ParseInt(opt.Default, 10, 64); err != nil {
				return fmt.Errorf(`fail to parse default value as integer: %s: %w`, opt.Default, err)
			}
		case TypeString, TypeUnspecified:
			// default value of option type is TypeString
			return nil
		}
	}
	return nil
}

func validateArg(arg *Arg) error {
	if err := validateType(arg.Type); err != nil {
		return fmt.Errorf("invalid type: %w", err)
	}
	if arg.Name == "" {
		return errors.New("name of arg is required")
	}
	return nil
}

func validateType(typ Type) error {
	switch typ {
	case TypeBoolean, TypeFloat, TypeInteger, TypeString, TypeUnspecified:
		return nil
	default:
		return errors.New(`unknown type: ` + string(typ))
	}
}
