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
)

type Schema struct {
	Program Program
}

// Load loads a CLI schema from a JSON data.
func Load(reader io.Reader) (Schema, error) {
	decoder := yaml.NewDecoder(reader, yaml.Strict())
	schema := Schema{}
	if err := decoder.Decode(&schema.Program); err != nil {
		return Schema{}, fmt.Errorf(`fail to unmarshal yaml as schema: %w`, err)
	}

	return schema, nil
}

//go:embed cli.schema.json
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

	if err := s.Program.validate(); err != nil {
		return fmt.Errorf(`fail to validate program: %w`, err)
	}
	return nil
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
	Name string `json:"name,omitempty"`

	// Version of the program.
	// The default value is an empty string.
	Version string `json:"version,omitempty"`

	// Embedded Command fields
	Command
}

func (p Program) validate() error {
	if err := p.Command.validate([]string{}); err != nil {
		return fmt.Errorf("invalid command: %w", err)
	}
	return nil
}

// Command represents a root command or a subcommand of the program.
// It may have options, arguments, and subcommands recursively.
type Command struct {
	// Description of the command.
	// The default value is an empty string.
	Description string `json:"description,omitempty"`

	// A collection of options, which is a mapping from option names to options.
	// The default value is an empty object.
	// A property name is a name of an option, which must match the regular expression `^(-[a-z][a-z0-9]*)+$` and be unique in options of the command.
	Options map[string]Option `json:"options,omitempty"`

	// A list of arguments.
	// The default value is an empty array.
	Arguments []Argument `json:"arguments,omitempty"`

	// A collection of subcommands, which is a mapping from subcommand names to child commands.
	// The default value is an empty object.
	// A property name is a name of a subcommand, which must match the regular expression `^[a-z][a-z0-9]*$` and be unique in subcommands of the command.
	Subcommands map[string]Command `json:"subcommands,omitempty"`
}

func (cmd Command) validate(propagatedOptions []string) error {
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

		if err := opt.Validate(); err != nil {
			return fmt.Errorf("invalid option %q: %w", name, err)
		}

		if opt.Propagates {
			// If the option propagates, add it to the propagated options
			propagatedOptions = append(propagatedOptions, name)
			if opt.Short != "" {
				propagatedOptions = append(propagatedOptions, opt.Short)
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
	for name, subcmd := range cmd.Subcommands {
		if err := subcmd.validate(propagatedOptions); err != nil {
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
	Short string `json:"short,omitempty"`

	// Description of the option.
	// The default value is an empty string.
	Description string `json:"description,omitempty"`

	// Type of the value that is assignable to this option.
	// The default value is "string".
	Type Type `json:"type,omitempty"`

	// String value representing the default value of the option.
	// It must be a string that can be parsed as a value of the option type.
	// If not specified, the following values corresponding to the option type.
	// - boolean: "false"
	// - string: ""
	// - integer: "0"
	Default string `json:"default,omitempty"`

	// Whether the option propagates to subcommands.
	// If true then the option is available in all subcommands of the command which the option belongs to.
	// The default value is false.
	Propagates bool `json:"propagates,omitempty"`
}

func (o Option) Validate() error {
	if err := o.Type.validate(); err != nil {
		return fmt.Errorf("invalid type %q: %w", o.Type, err)
	}
	return nil
}

// Argument represents a positional required argument in command line arguments.
type Argument struct {
	// Name of the argument, which must match the regular expression `^[a-z][a-z0-9]*$` and be unique in arguments of the command which the argument belongs to.
	// This property is required.
	Name string `json:"name"`

	// Description of the argument.
	// The default value is an empty string.
	Description string `json:"description,omitempty"`

	// Type of the value that is assignable to the argument.
	// The default value is "string".
	Type Type `json:"type,omitempty"`

	// Whether the argument is variadic (i.e. can have zero or more values).
	// It can be true only if this argument is the last argument in the arguments of the belonging command.
	// The default value is false.
	Variadic bool `json:"variadic,omitempty"`
}

func (arg Argument) validate() error {
	if err := arg.Type.validate(); err != nil {
		return fmt.Errorf("invalid type %q: %w", arg.Type, err)
	}
	return nil
}
