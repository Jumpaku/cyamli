package schema

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/samber/lo"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"io"
	"slices"
	"strconv"
)

type Schema struct {
	optionPropagated bool
	Program          Program
}

// PropagateOptions creates a deep copy of the Schema.
func (s Schema) PropagateOptions() Schema {
	if s.optionPropagated {
		return s
	}
	return Schema{
		optionPropagated: true,
		Program:          s.Program.propagateOptions(),
	}
}

// Load loads a CLI schema from a JSON data.
func Load(reader io.Reader) (Schema, error) {
	decoder := yaml.NewDecoder(reader, yaml.Strict())
	schema := Schema{}
	if err := decoder.Decode(&schema.Program); err != nil {
		return Schema{}, fmt.Errorf(`fail to unmarshal yaml as schema: %w`, err)
	}
	if err := schema.Validate(); err != nil {
		return Schema{}, fmt.Errorf(`fail to validate schema: %w`, err)
	}
	return schema, nil
}

//go:embed cyamli-cli.schema.json
var cliSchemaJSON string
var cliJSONSchema *jsonschema.Schema = lo.Must1(jsonschema.CompileString("", cliSchemaJSON))

func validateByJSONSchema(b []byte) error {
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.UseNumber()
	var v any
	if err := decoder.Decode(&v); err != nil {
		return fmt.Errorf("fail to unmarshal JSON: %w", err)
	}

	if err := cliJSONSchema.Validate(v); err != nil {
		return fmt.Errorf("fail to validate schema: %w", err)
	}

	return nil
}

func (s Schema) Validate() error {
	b := lo.Must1(json.Marshal(s.Program))
	if err := validateByJSONSchema(b); err != nil {
		return fmt.Errorf(`fail to validate schema based on JSON schema: %w`, err)
	}

	if err := s.Program.validate(s.optionPropagated); err != nil {
		return fmt.Errorf(`fail to validate program: %w`, err)
	}
	return nil
}

type PathCommand struct {
	Path    []string
	Command Command
}

func (s Schema) ListCommand() (commands []PathCommand) {
	_ = s.Program.Walk([]string{}, func(cmd Command, path []string) error {
		commands = append(commands, PathCommand{Path: path, Command: cmd})
		return nil
	})

	slices.SortFunc(commands, func(a, b PathCommand) int { return slices.Compare(a.Path, b.Path) })

	return commands
}

// Type represents a type of a value that can be assigned to an option or an argument.
// One of "string", "integer", or "boolean" is available.
type Type string

const (
	TypeEmpty   Type = ""
	TypeString  Type = "string"
	TypeInteger Type = "integer"
	TypeBoolean Type = "boolean"
)

func (t Type) validate() error {
	switch t {
	default:
		return fmt.Errorf("unknown type: %s", t)
	case TypeEmpty, TypeString, TypeInteger, TypeBoolean:
		return nil
	}
}

// Program is a root command that may have a name and a version.
// It consists of commands recursively.
type Program struct {
	// Name of the program.
	// The default value is an empty string.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Version of the program.
	// The default value is an empty string.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Embedded Command fields
	Command `yaml:",inline"`
}

// propagateOptions creates a deep copy of the Program.
func (p Program) propagateOptions() Program {
	clone := Program{
		Name:    p.Name,
		Version: p.Version,
		Command: p.Command.propagateOptions(map[string]Option{}),
	}
	return clone
}

func (p Program) validate(propagated bool) error {
	if err := p.Command.validate([]string{}, propagated); err != nil {
		return fmt.Errorf("invalid command: %w", err)
	}
	return nil
}

// Command represents a root command or a subcommand of the program.
// It may have options, arguments, and subcommands recursively.
type Command struct {
	// Description of the command.
	// The default value is an empty string.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A collection of options, which is a mapping from option names to options.
	// The default value is an empty object.
	// A property name is a name of an option, which must match the regular expression `^(-[a-z][a-z0-9]*)+$` and be unique in options of the command.
	Options map[string]Option `json:"options,omitempty" yaml:"options,omitempty"`

	// A list of arguments.
	// The default value is an empty array.
	Arguments []Argument `json:"arguments,omitempty" yaml:"arguments,omitempty"`

	// A collection of subcommands, which is a mapping from subcommand names to child commands.
	// The default value is an empty object.
	// A property name is a name of a subcommand, which must match the regular expression `^[a-z][a-z0-9]*$` and be unique in subcommands of the command.
	Subcommands map[string]Command `json:"subcommands,omitempty" yaml:"subcommands,omitempty"`
}

// propagateOptions creates a deep copy of the Command.
func (cmd Command) propagateOptions(propagatedOptions map[string]Option) (propagated Command) {
	propagated = Command{
		Description: cmd.Description,
	}

	propagatedOptionsClone := make(map[string]Option)
	for k, v := range propagatedOptions {
		propagatedOptionsClone[k] = v
	}

	if cmd.Options != nil || len(propagatedOptionsClone) > 0 {
		propagated.Options = make(map[string]Option)
	}
	for k, v := range propagatedOptions {
		propagated.Options[k] = v
	}
	if cmd.Options != nil {
		for k, v := range cmd.Options {
			propagated.Options[k] = v
			if v.Propagates {
				propagatedOptionsClone[k] = v
			}
		}
	}

	if cmd.Arguments != nil {
		propagated.Arguments = make([]Argument, len(cmd.Arguments))
		for i, arg := range cmd.Arguments {
			propagated.Arguments[i] = arg
		}
	}

	if cmd.Subcommands != nil {
		propagated.Subcommands = make(map[string]Command)
		for k, v := range cmd.Subcommands {
			propagated.Subcommands[k] = v.propagateOptions(propagatedOptionsClone)
		}
	}

	return propagated
}

func (cmd Command) validate(propagatedOptions []string, propagated bool) error {
	propagatedOptions = append([]string{}, propagatedOptions...)

	// Validate options
	options := lo.SliceToMap(propagatedOptions, func(opt string) (string, bool) {
		return opt, true
	})
	for name, opt := range cmd.Options {
		if _, ok := options[name]; ok {
			return fmt.Errorf("duplicate option name %q", name)
		}
		options[name] = true

		if opt.Short != "" {
			if _, ok := options[opt.Short]; ok {
				return fmt.Errorf("duplicate short option name %q", opt.Short)
			}
			options[opt.Short] = true
		}
		if opt.Negation {
			negatedName := "-no" + name
			if _, ok := options[negatedName]; ok {
				return fmt.Errorf("duplicate negated option name %q", negatedName)
			}
			options[negatedName] = true
		}

		if err := opt.validate(); err != nil {
			return fmt.Errorf("invalid option %q: %w", name, err)
		}

		if !propagated {
			if opt.Propagates {
				// If the option propagates, add it to the propagated options
				propagatedOptions = append(propagatedOptions, name)
				if opt.Short != "" {
					propagatedOptions = append(propagatedOptions, opt.Short)
				}
				if opt.Negation {
					propagatedOptions = append(propagatedOptions, "-no"+name)
				}
			}
		}
	}

	// Validate arguments
	arguments := map[string]bool{}
	for idx, arg := range cmd.Arguments {
		if _, ok := arguments[arg.Name]; ok {
			return fmt.Errorf("duplicate argument name %q", arg.Name)
		}
		arguments[arg.Name] = true

		if idx < len(cmd.Arguments)-1 && arg.Variadic {
			return fmt.Errorf("argument %q cannot be variadic because it is not the last argument", arg.Name)
		}
		if err := arg.validate(); err != nil {
			return fmt.Errorf("invalid argument %q: %w", arg.Name, err)
		}

	}

	// Validate subcommands
	for name, sub := range cmd.Subcommands {
		if err := sub.validate(propagatedOptions, propagated); err != nil {
			return fmt.Errorf("invalid subcommand %q: %w", name, err)
		}
	}

	return nil
}
func (cmd Command) Walk(path []string, f func(cmd Command, path []string) error) error {
	// Walk the command itself
	if err := f(cmd, path); err != nil {
		return err
	}

	// Walk subcommands recursively
	for name, subcmd := range cmd.Subcommands {
		if err := subcmd.Walk(append(append([]string{}, path...), name), f); err != nil {
			return err
		}
	}
	return nil
}

// Option represents an optional argument in command line arguments.
type Option struct {
	// Short name of the option, which must match the regular expression `^-[a-z]$` and be unique in options of the command which the option belongs to.
	// If short is not specified then short name for this option is not available.
	Short string `json:"short,omitempty" yaml:"short,omitempty"`

	// Description of the option.
	// The default value is an empty string.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Type of the value that is assignable to this option.
	// The default value is "string".
	Type Type `json:"type,omitempty" yaml:"type,omitempty"`

	// Whether the option of typed boolean has a negated version.
	// If true then the option can be specified with a negation prefix `-no` in the command line arguments.
	// The default value is false.
	Negation bool `json:"negation,omitempty" yaml:"negation,omitempty"`

	// String value representing the default value of the option.
	// It must be a string that can be parsed as a value of the option type.
	// If not specified, the following values corresponding to the option type.
	// - boolean: "false"
	// - string: ""
	// - integer: "0"
	Default string `json:"default,omitempty" yaml:"default,omitempty"`

	// Whether the option can be specified multiple times.
	// If true then the option can be specified multiple times in the command line arguments.
	// The default value is false.
	Repeated bool `json:"repeated,omitempty" yaml:"repeated,omitempty"`

	// Whether the option propagates to subcommands.
	// If true then the option is available in all subcommands of the command which the option belongs to.
	// The default value is false.
	Propagates bool `json:"propagates,omitempty" yaml:"propagates,omitempty"`
}

func (o Option) validate() error {
	if o.Default != "" {
		if o.Repeated {
			return fmt.Errorf("default value cannot be specified when repeated is true")
		}
		switch o.Type {
		case TypeInteger:
			if _, err := strconv.ParseInt(o.Default, 0, 64); err != nil {
				return fmt.Errorf(`fail to parse default value as integer: %s: %w`, o.Default, err)
			}
		case TypeBoolean:
			if _, err := strconv.ParseBool(o.Default); err != nil {
				return fmt.Errorf(`fail to parse default value as boolean: %s: %w`, o.Default, err)
			}
		}
	}
	if o.Type != TypeBoolean {
		if o.Negation {
			return fmt.Errorf("negation is only available for boolean type options")
		}
	}
	if err := o.Type.validate(); err != nil {
		return fmt.Errorf("invalid type %q: %w", o.Type, err)
	}
	return nil
}

// Argument represents a positional required argument in command line arguments.
type Argument struct {
	// Name of the argument, which must match the regular expression `^[a-z][a-z0-9]*$` and be unique in arguments of the command which the argument belongs to.
	// This property is required.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Description of the argument.
	// The default value is an empty string.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Type of the value that is assignable to the argument.
	// The default value is "string".
	Type Type `json:"type,omitempty" yaml:"type,omitempty"`

	// Whether the argument is variadic (i.e. can have zero or more values).
	// It can be true only if this argument is the last argument in the arguments of the belonging command.
	// The default value is false.
	Variadic bool `json:"variadic,omitempty" yaml:"variadic,omitempty"`
}

func (arg Argument) validate() error {
	if err := arg.Type.validate(); err != nil {
		return fmt.Errorf("invalid type %q: %w", arg.Type, err)
	}
	return nil
}
