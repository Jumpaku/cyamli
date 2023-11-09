package schema

import (
	"fmt"
	"regexp"
)

type Program struct {
	Name        string              `yaml:"name"`
	Version     string              `yaml:"version"`
	Description string              `yaml:"description"`
	Options     map[string]*Option  `yaml:"options"`
	Arguments   []*Argument         `yaml:"arguments"`
	Subcommands map[string]*Command `yaml:"subcommands"`
}

func (p *Program) Validate() error {
	if err := p.Command().Validate(); err != nil {
		return fmt.Errorf("invalid program: %w", err)
	}
	return nil
}

func validateProgramName(name string) error {
	r := regexp.MustCompile(`(^$)|^([a-z][a-zA-Z0-9_-]*)$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("program name %q must match %v", name, r.String())
	}

	return nil
}

func (program *Program) Command() *Command {
	return &Command{
		Description: program.Description,
		Options:     program.Options,
		Arguments:   program.Arguments,
		Subcommands: program.Subcommands,
	}
}
