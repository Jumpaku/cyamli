package schema

import (
	"fmt"
	"regexp"
)

type Argument struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        Type   `yaml:"type"`
	Variadic    bool   `yaml:"variadic"`
}

func (a *Argument) Validate() error {
	if err := validateArgumentName(a.Name); err != nil {
		return fmt.Errorf("invalid argument name: %q: %w", a.Name, err)
	}
	if err := a.Type.Validate(); err != nil {
		return fmt.Errorf("invalid type: %w", err)
	}
	return nil
}

func validateArgumentName(name string) error {
	r := regexp.MustCompile(`^[a-z][a-z0-9]*(_[a-z0-9]+)*$`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("arg name %q must match %v", name, r.String())
	}

	return nil
}
