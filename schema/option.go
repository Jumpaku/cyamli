package schema

import (
	"fmt"
	"regexp"
	"strconv"
)

type Option struct {
	Short       string `yaml:"short"`
	Description string `yaml:"description"`
	Type        Type   `yaml:"type"`
	Default     string `yaml:"default"`
}

func (o *Option) Validate() error {
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid type: %w", err)
	}
	if err := validateOptionShortName(o.Short); err != nil {
		return fmt.Errorf("invalid short name: %w", err)
	}
	if o.Default != "" {
		switch o.Type {
		case TypeBoolean:
			if _, err := strconv.ParseBool(o.Default); err != nil {
				return fmt.Errorf(`fail to parse default value as boolean: %s: %w`, o.Default, err)
			}
		case TypeFloat:
			if _, err := strconv.ParseFloat(o.Default, 64); err != nil {
				return fmt.Errorf(`fail to parse default value as float: %s: %w`, o.Default, err)
			}
		case TypeInteger:
			if _, err := strconv.ParseInt(o.Default, 10, 64); err != nil {
				return fmt.Errorf(`fail to parse default value as integer: %s: %w`, o.Default, err)
			}
		case TypeString, TypeUnspecified:
			// default value of option type is TypeString
			return nil
		}
	}
	return nil
}

func validateOptionShortName(name string) error {
	r := regexp.MustCompile(`(^$)|(^-[a-z]$)`)
	match := r.MatchString(name)
	if !match {
		return fmt.Errorf("option short name %s must match %v", name, r.String())
	}

	return nil
}
