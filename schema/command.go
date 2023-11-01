package schema

import (
	"errors"
	"fmt"
	"regexp"
)

type Command struct {
	Description string              `yaml:"description"`
	Options     map[string]*Option  `yaml:"options"`
	Arguments   []*Argument         `yaml:"arguments"`
	Subcommands map[string]*Command `yaml:"subcommands"`
}

func (c *Command) Validate() error {
	optShortNames := map[string]bool{}
	for name, opt := range c.Options {
		if err := validateOptionName(name); err != nil {
			return fmt.Errorf("invalid option name: %s: %w", name, err)
		}
		if opt.Short != "" && optShortNames[opt.Short] {
			return fmt.Errorf("option short name %q duplicated", opt.Short)
		} else {
			optShortNames[opt.Short] = true
		}
		if err := opt.Validate(); err != nil {
			return fmt.Errorf("invalid option: %s: %w", name, err)
		}
	}

	argNames := map[string]bool{}
	for idx, arg := range c.Arguments {
		if argNames[arg.Name] {
			return fmt.Errorf("argument name %q duplicated", arg.Name)
		} else {
			argNames[arg.Name] = true
		}
		if err := arg.Validate(); err != nil {
			return fmt.Errorf("invalid argument: %w", err)
		}
		if idx < len(c.Arguments)-1 && arg.Variadic {
			return errors.New("variadic argument must be placed at last")
		}
	}

	for name, sub := range c.Subcommands {
		if err := validateSubcommandName(name); err != nil {
			return fmt.Errorf("invalid subcommand name: %s: %w", name, err)
		}
		if err := sub.Validate(); err != nil {
			return fmt.Errorf("invalid subcommand: %s: %w", name, err)
		}
	}
	return nil
}

func validateOptionName(name string) error {
	r := regexp.MustCompile(`^(-[a-z][a-z0-9]*)+$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("option name %s must match %v", name, r.String())
	}

	return nil
}

func validateSubcommandName(name string) error {
	r := regexp.MustCompile(`^[a-z][a-z0-9]*$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("subcommand name %s must match %v", name, r.String())
	}

	return nil
}
