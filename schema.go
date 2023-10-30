package cliautor

import (
	"errors"
	"fmt"
	"regexp"
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
	Variadic    bool   `yaml:"variadic"`
}
type Cmd struct {
	Description string          `yaml:"description"`
	Options     map[string]*Opt `yaml:"options"`
	Args        []*Arg          `yaml:"args"`
	Subcommands map[string]*Cmd `yaml:"subcommands"`
}
type Program struct {
	Name        string          `yaml:"name"`
	Version     string          `yaml:"version"`
	Description string          `yaml:"description"`
	Options     map[string]*Opt `yaml:"options"`
	Args        []*Arg          `yaml:"args"`
	Subcommands map[string]*Cmd `yaml:"subcommands"`
}

func (program Program) Cmd() *Cmd {
	return &Cmd{
		Description: program.Description,
		Options:     program.Options,
		Args:        program.Args,
		Subcommands: program.Subcommands,
	}
}

func Load(b []byte) (*Schema, error) {
	schema := Schema{}
	err := yaml.Unmarshal(b, &schema.Program)
	if err != nil {
		return nil, fmt.Errorf(`fail to parse yaml as command structure: %w`, err)
	}
	if err := (&schema).Validate(); err != nil {
		return nil, fmt.Errorf(`fail to validate command: %w`, err)
	}
	return &schema, nil
}

type Schema struct {
	Program Program
}

func (s *Schema) Validate() error {
	return validateProgram(&s.Program)
}

func validateProgram(program *Program) error {
	return validateCmd(program.Cmd())
}

func validateCmd(cmd *Cmd) error {
	optionNames := map[string]string{}
	for name, opt := range cmd.Options {
		if err := validateOptName(name); err != nil {
			return fmt.Errorf("invalid option name: %s: %w", name, err)
		}
		if conflict, exists := optionNames[strings.Join(Words(name), " ")]; exists {
			return fmt.Errorf("option %s conflicts with %s", name, conflict)
		}
		optionNames[strings.Join(Words(name), " ")] = name

		if err := validateOpt(opt); err != nil {
			return fmt.Errorf("invalid option: %s: %w", name, err)
		}
	}

	argNames := map[string]string{}
	for idx, arg := range cmd.Args {
		if err := validateArg(arg); err != nil {
			return fmt.Errorf("invalid arg: %w", err)
		}
		if conflict, exists := argNames[strings.Join(Words(arg.Name), " ")]; exists {
			return fmt.Errorf("arg %s conflicts with %s", arg.Name, conflict)
		}
		argNames[strings.Join(Words(arg.Name), " ")] = arg.Name

		if idx < len(cmd.Args)-1 && arg.Variadic {
			return errors.New("variadic arg must be placed at last")
		}
	}

	subNames := map[string]string{}
	for name, sub := range cmd.Subcommands {
		if name == "" {
			return errors.New("name of subcommand must be not empty")
		}
		if err := validateSubcommandName(name); err != nil {
			return fmt.Errorf("invalid subcommand name: %s: %w", name, err)
		}
		if conflict, exists := subNames[strings.Join(Words(name), " ")]; exists {
			return fmt.Errorf("subcommand %s conflicts with %s", name, conflict)
		}
		subNames[strings.Join(Words(name), " ")] = name

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
	if err := validateOptShortName(opt.Short); err != nil {
		return fmt.Errorf("invalid short name: %w", err)
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
	if arg.Name == "" {
		return errors.New("name of arg is required")
	}
	if err := validateArgName(arg.Name); err != nil {
		return fmt.Errorf("invalid arg name: %s: %w", arg.Name, err)
	}
	if err := validateType(arg.Type); err != nil {
		return fmt.Errorf("invalid type: %w", err)
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

func validateSubcommandName(name string) error {
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-_]*$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("subcommand name %s must match %v", name, r.String())
	}

	return nil
}

func validateOptName(name string) error {
	r := regexp.MustCompile(`^(-[a-zA-Z][a-zA-Z0-9]*)+$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("option name %s must match %v", name, r.String())
	}

	return nil
}

func validateOptShortName(name string) error {
	r := regexp.MustCompile(`^-[a-zA-Z]$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("option short name %s must match %v", name, r.String())
	}

	return nil
}

func validateArgName(name string) error {
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("arg name %s must match %v", name, r.String())
	}

	return nil
}
