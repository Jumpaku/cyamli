package schema

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

// Helper functions for building test schemas
func buildBasicSchema() Schema {
	return Schema{
		Program: Program{
			Name:    "test-program",
			Version: "1.0.0",
			Command: Command{
				Description: "test program",
				Options:     map[string]Option{},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
		},
	}
}

func buildSchemaWithOption(name string, opt Option) Schema {
	schema := buildBasicSchema()
	schema.Program.Command.Options = map[string]Option{
		name: opt,
	}
	return schema
}

func buildSchemaWithArgument(arg Argument) Schema {
	schema := buildBasicSchema()
	schema.Program.Command.Arguments = []Argument{arg}
	return schema
}

func buildSchemaWithSubcommand(name string, cmd Command) Schema {
	schema := buildBasicSchema()
	schema.Program.Command.Subcommands = map[string]Command{
		name: cmd,
	}
	return schema
}

func TestType_validate(t *testing.T) {
	tests := []struct {
		name    string
		t       Type
		wantErr bool
	}{
		{
			name:    "empty type is valid",
			t:       TypeEmpty,
			wantErr: false,
		},
		{
			name:    "string type is valid",
			t:       TypeString,
			wantErr: false,
		},
		{
			name:    "integer type is valid",
			t:       TypeInteger,
			wantErr: false,
		},
		{
			name:    "boolean type is valid",
			t:       TypeBoolean,
			wantErr: false,
		},
		{
			name:    "unknown type is invalid",
			t:       Type("unknown"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.t.validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestArgument_validate(t *testing.T) {
	tests := []struct {
		name    string
		arg     Argument
		wantErr bool
	}{
		{
			name: "valid argument with string type",
			arg: Argument{
				Name:        "filename",
				Description: "input filename",
				Type:        TypeString,
				Variadic:    false,
			},
			wantErr: false,
		},
		{
			name: "valid argument with integer type",
			arg: Argument{
				Name:        "count",
				Description: "number of items",
				Type:        TypeInteger,
				Variadic:    false,
			},
			wantErr: false,
		},
		{
			name: "valid argument with boolean type",
			arg: Argument{
				Name:        "flag",
				Description: "boolean flag",
				Type:        TypeBoolean,
				Variadic:    false,
			},
			wantErr: false,
		},
		{
			name: "valid argument with empty type (defaults to string)",
			arg: Argument{
				Name:        "default",
				Description: "uses default type",
				Type:        TypeEmpty,
				Variadic:    false,
			},
			wantErr: false,
		},
		{
			name: "valid variadic argument",
			arg: Argument{
				Name:        "files",
				Description: "list of files",
				Type:        TypeString,
				Variadic:    true,
			},
			wantErr: false,
		},
		{
			name: "invalid argument with unknown type",
			arg: Argument{
				Name:        "invalid",
				Description: "has invalid type",
				Type:        Type("unknown"),
				Variadic:    false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.arg.validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestOption_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opt     Option
		wantErr bool
	}{
		{
			name: "valid option with string type",
			opt: Option{
				Short:       "-f",
				Description: "filename option",
				Type:        TypeString,
				Default:     "default.txt",
				Propagates:  false,
			},
			wantErr: false,
		},
		{
			name: "valid option with integer type",
			opt: Option{
				Short:       "-c",
				Description: "count option",
				Type:        TypeInteger,
				Default:     "10",
				Propagates:  false,
			},
			wantErr: false,
		},
		{
			name: "valid option with boolean type",
			opt: Option{
				Short:       "-b",
				Description: "boolean option",
				Type:        TypeBoolean,
				Default:     "true",
				Propagates:  false,
			},
			wantErr: false,
		},
		{
			name: "valid option with empty type (defaults to string)",
			opt: Option{
				Short:       "-d",
				Description: "default type option",
				Type:        TypeEmpty,
				Default:     "",
				Propagates:  false,
			},
			wantErr: false,
		},
		{
			name: "valid option with propagation",
			opt: Option{
				Short:       "-p",
				Description: "propagating option",
				Type:        TypeString,
				Default:     "",
				Propagates:  true,
			},
			wantErr: false,
		},
		{
			name: "valid option without short name",
			opt: Option{
				Short:       "",
				Description: "long-only option",
				Type:        TypeString,
				Default:     "",
				Propagates:  false,
			},
			wantErr: false,
		},
		{
			name: "invalid option with unknown type",
			opt: Option{
				Short:       "-i",
				Description: "invalid type option",
				Type:        Type("unknown"),
				Default:     "",
				Propagates:  false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opt.validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSchema_Load(t *testing.T) {
	tests := []struct {
		name    string
		yaml    string
		wantErr bool
	}{
		{
			name: "valid schema",
			yaml: `
name: test-program
version: 1.0.0
description: test program
options:
  -option1:
    short: -o
    description: option 1
    type: string
arguments:
  - name: arg1
    description: argument 1
    type: string
subcommands:
  sub1:
    description: subcommand 1
    options:
      -suboption1:
        short: -s
        description: suboption 1
        type: integer
`,
			wantErr: false,
		},
		{
			name: "invalid schema - malformed yaml",
			yaml: `
name: test-program
version: 1.0.0
description: test program
options:
  -option1:
    short: -o
    description: option 1
    type: string
  invalid yaml
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a reader from the YAML string
			reader := bytes.NewBufferString(tt.yaml)
			schema, err := Load(reader)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, schema)
				require.Equal(t, "test-program", schema.Program.Name)
				require.Equal(t, "1.0.0", schema.Program.Version)
				require.Equal(t, "test program", schema.Program.Description)
			}
		})
	}
}

func TestSchema_Validate(t *testing.T) {
	tests := []struct {
		name    string
		schema  Schema
		wantErr bool
	}{
		{
			name: "valid schema",
			schema: Schema{
				Program: Program{
					Name:    "test-program",
					Version: "1.0.0",
					Command: Command{
						Description: "test program",
						Options: map[string]Option{
							"-option1": {
								Short:       "-o",
								Description: "option 1",
								Type:        TypeString,
							},
						},
						Arguments: []Argument{
							{
								Name:        "arg1",
								Description: "argument 1",
								Type:        TypeString,
							},
						},
						Subcommands: map[string]Command{
							"sub1": {
								Description: "subcommand 1",
								Options:     map[string]Option{},
								Arguments:   []Argument{},
								Subcommands: map[string]Command{},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid schema - invalid option type",
			schema: Schema{
				Program: Program{
					Name:    "test-program",
					Version: "1.0.0",
					Command: Command{
						Description: "test program",
						Options: map[string]Option{
							"-option1": {
								Short:       "-o",
								Description: "option 1",
								Type:        Type("unknown"),
							},
						},
						Arguments:   []Argument{},
						Subcommands: map[string]Command{},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.schema.Validate()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProgram_validate(t *testing.T) {
	tests := []struct {
		name    string
		program Program
		wantErr bool
	}{
		{
			name: "valid empty program",
			program: Program{
				Name:    "test-program",
				Version: "1.0.0",
				Command: Command{
					Description: "test program",
					Options:     map[string]Option{},
					Arguments:   []Argument{},
					Subcommands: map[string]Command{},
				},
			},
			wantErr: false,
		},
		{
			name: "valid program with options",
			program: Program{
				Name:    "test-program",
				Version: "1.0.0",
				Command: Command{
					Description: "test program with options",
					Options: map[string]Option{
						"-option1": {
							Short:       "-o",
							Description: "option 1",
							Type:        TypeString,
						},
					},
					Arguments:   []Argument{},
					Subcommands: map[string]Command{},
				},
			},
			wantErr: false,
		},
		{
			name: "valid program with arguments",
			program: Program{
				Name:    "test-program",
				Version: "1.0.0",
				Command: Command{
					Description: "test program with arguments",
					Options:     map[string]Option{},
					Arguments: []Argument{
						{
							Name:        "arg1",
							Description: "argument 1",
							Type:        TypeString,
						},
					},
					Subcommands: map[string]Command{},
				},
			},
			wantErr: false,
		},
		{
			name: "valid program with subcommands",
			program: Program{
				Name:    "test-program",
				Version: "1.0.0",
				Command: Command{
					Description: "test program with subcommands",
					Options:     map[string]Option{},
					Arguments:   []Argument{},
					Subcommands: map[string]Command{
						"sub1": {
							Description: "subcommand 1",
							Options:     map[string]Option{},
							Arguments:   []Argument{},
							Subcommands: map[string]Command{},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid program with invalid command",
			program: Program{
				Name:    "test-program",
				Version: "1.0.0",
				Command: Command{
					Description: "test program with invalid command",
					Options: map[string]Option{
						"-option1": {
							Short:       "-o",
							Description: "option 1",
							Type:        Type("unknown"),
						},
					},
					Arguments:   []Argument{},
					Subcommands: map[string]Command{},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.program.validate(false)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCommand_validate(t *testing.T) {
	tests := []struct {
		name    string
		cmd     Command
		wantErr bool
	}{
		{
			name: "valid empty command",
			cmd: Command{
				Description: "empty command",
				Options:     map[string]Option{},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			wantErr: false,
		},
		{
			name: "valid command with options",
			cmd: Command{
				Description: "command with options",
				Options: map[string]Option{
					"-option1": {
						Short:       "-o",
						Description: "option 1",
						Type:        TypeString,
					},
					"-option2": {
						Short:       "-p",
						Description: "option 2",
						Type:        TypeInteger,
					},
				},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			wantErr: false,
		},
		{
			name: "valid command with arguments",
			cmd: Command{
				Description: "command with arguments",
				Options:     map[string]Option{},
				Arguments: []Argument{
					{
						Name:        "arg1",
						Description: "argument 1",
						Type:        TypeString,
					},
					{
						Name:        "arg2",
						Description: "argument 2",
						Type:        TypeInteger,
					},
				},
				Subcommands: map[string]Command{},
			},
			wantErr: false,
		},
		{
			name: "valid command with subcommands",
			cmd: Command{
				Description: "command with subcommands",
				Options:     map[string]Option{},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{
					"sub1": {
						Description: "subcommand 1",
						Options:     map[string]Option{},
						Arguments:   []Argument{},
						Subcommands: map[string]Command{},
					},
					"sub2": {
						Description: "subcommand 2",
						Options:     map[string]Option{},
						Arguments:   []Argument{},
						Subcommands: map[string]Command{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid command with duplicate short option names",
			cmd: Command{
				Description: "command with duplicate short option names",
				Options: map[string]Option{
					"-option1": {
						Short:       "-o",
						Description: "option 1",
						Type:        TypeString,
					},
					"-option2": {
						Short:       "-o", // Duplicate short name
						Description: "option 2",
						Type:        TypeInteger,
					},
				},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			wantErr: true,
		},
		{
			name: "invalid command with invalid option type",
			cmd: Command{
				Description: "command with invalid option type",
				Options: map[string]Option{
					"-option1": {
						Short:       "-o",
						Description: "option 1",
						Type:        Type("unknown"),
					},
				},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			wantErr: true,
		},
		{
			name: "invalid command with duplicate argument names",
			cmd: Command{
				Description: "command with duplicate argument names",
				Options:     map[string]Option{},
				Arguments: []Argument{
					{
						Name:        "arg1",
						Description: "argument 1",
						Type:        TypeString,
					},
					{
						Name:        "arg1", // Duplicate name
						Description: "argument 2",
						Type:        TypeInteger,
					},
				},
				Subcommands: map[string]Command{},
			},
			wantErr: true,
		},
		{
			name: "invalid command with invalid argument type",
			cmd: Command{
				Description: "command with invalid argument type",
				Options:     map[string]Option{},
				Arguments: []Argument{
					{
						Name:        "arg1",
						Description: "argument 1",
						Type:        Type("unknown"),
					},
				},
				Subcommands: map[string]Command{},
			},
			wantErr: true,
		},
		{
			name: "invalid command with non-last variadic argument",
			cmd: Command{
				Description: "command with non-last variadic argument",
				Options:     map[string]Option{},
				Arguments: []Argument{
					{
						Name:        "arg1",
						Description: "argument 1",
						Type:        TypeString,
						Variadic:    true, // Should be last
					},
					{
						Name:        "arg2",
						Description: "argument 2",
						Type:        TypeInteger,
					},
				},
				Subcommands: map[string]Command{},
			},
			wantErr: true,
		},
		{
			name: "invalid command with invalid subcommand",
			cmd: Command{
				Description: "command with invalid subcommand",
				Options:     map[string]Option{},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{
					"sub1": {
						Description: "subcommand 1",
						Options: map[string]Option{
							"-option1": {
								Short:       "-o",
								Description: "option 1",
								Type:        Type("unknown"),
							},
						},
						Arguments:   []Argument{},
						Subcommands: map[string]Command{},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cmd.validate([]string{}, false)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCommand_validate_DuplicateOptionName(t *testing.T) {
	tests := []struct {
		name              string
		cmd               Command
		propagatedOptions []string
		wantErr           bool
	}{
		{
			name: "valid command with no duplicate option names",
			cmd: Command{
				Description: "command with no duplicate option names",
				Options: map[string]Option{
					"-option1": {
						Short:       "-o",
						Description: "option 1",
						Type:        TypeString,
					},
				},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			propagatedOptions: []string{"-option2"},
			wantErr:           false,
		},
		{
			name: "invalid command with duplicate option name from propagated options",
			cmd: Command{
				Description: "command with duplicate option name from propagated options",
				Options: map[string]Option{
					"-option1": {
						Short:       "-o",
						Description: "option 1",
						Type:        TypeString,
					},
				},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			propagatedOptions: []string{"-option1"},
			wantErr:           true,
		},
		{
			name: "invalid command with duplicate short option name from propagated options",
			cmd: Command{
				Description: "command with duplicate short option name from propagated options",
				Options: map[string]Option{
					"-option1": {
						Short:       "-o",
						Description: "option 1",
						Type:        TypeString,
					},
				},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			},
			propagatedOptions: []string{"-o"},
			wantErr:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cmd.validate(tt.propagatedOptions, false)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSchema_OptionNameValidation(t *testing.T) {
	tests := []struct {
		name        string
		optionName  string
		description string
		wantErr     bool
	}{
		{
			name:        "valid option name with single dash segment",
			optionName:  "-option1",
			description: "option 1",
			wantErr:     false,
		},
		{
			name:        "valid option name with multiple dash segments",
			optionName:  "-option-with-dashes",
			description: "option with dashes",
			wantErr:     false,
		},
		{
			name:        "valid option name with numbers",
			optionName:  "-option123",
			description: "option with numbers",
			wantErr:     false,
		},
		{
			name:        "invalid option name without dash prefix",
			optionName:  "option1", // Missing dash prefix
			description: "option 1",
			wantErr:     true,
		},
		{
			name:        "invalid option name with uppercase letters",
			optionName:  "-Option1", // Contains uppercase letter
			description: "option 1",
			wantErr:     true,
		},
		{
			name:        "invalid option name starting with number after dash",
			optionName:  "-1option", // Starts with number after dash
			description: "option 1",
			wantErr:     true,
		},
		{
			name:        "invalid option name with special characters",
			optionName:  "-option_1", // Contains underscore
			description: "option 1",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := Option{
				Description: tt.description,
				Type:        TypeString,
			}
			schema := buildSchemaWithOption(tt.optionName, opt)

			err := schema.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSchema_SubcommandNameValidation(t *testing.T) {
	tests := []struct {
		name        string
		subcmdName  string
		description string
		wantErr     bool
	}{
		{
			name:        "valid subcommand name with letters only",
			subcmdName:  "subcommand",
			description: "subcommand",
			wantErr:     false,
		},
		{
			name:        "valid subcommand name with letters and numbers",
			subcmdName:  "subcommand123",
			description: "subcommand with numbers",
			wantErr:     false,
		},
		{
			name:        "invalid subcommand name starting with number",
			subcmdName:  "1subcommand", // Starts with number
			description: "subcommand starting with number",
			wantErr:     true,
		},
		{
			name:        "invalid subcommand name with uppercase letters",
			subcmdName:  "Subcommand", // Contains uppercase letter
			description: "subcommand with uppercase",
			wantErr:     true,
		},
		{
			name:        "invalid subcommand name with special characters",
			subcmdName:  "sub-command", // Contains dash
			description: "subcommand with dash",
			wantErr:     true,
		},
		{
			name:        "invalid subcommand name with underscore",
			subcmdName:  "sub_command", // Contains underscore
			description: "subcommand with underscore",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Command{
				Description: tt.description,
				Options:     map[string]Option{},
				Arguments:   []Argument{},
				Subcommands: map[string]Command{},
			}
			schema := buildSchemaWithSubcommand(tt.subcmdName, cmd)

			err := schema.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSchema_ShortOptionNameValidation(t *testing.T) {
	tests := []struct {
		name        string
		shortName   string
		description string
		wantErr     bool
	}{
		{
			name:        "valid short option name",
			shortName:   "-a",
			description: "option with valid short name",
			wantErr:     false,
		},
		{
			name:        "valid empty short option name",
			shortName:   "", // Empty is valid
			description: "option with empty short name",
			wantErr:     false,
		},
		{
			name:        "invalid short option name without dash",
			shortName:   "a", // Missing dash
			description: "option with invalid short name",
			wantErr:     true,
		},
		{
			name:        "invalid short option name with uppercase letter",
			shortName:   "-A", // Uppercase letter
			description: "option with invalid short name",
			wantErr:     true,
		},
		{
			name:        "invalid short option name with multiple characters",
			shortName:   "-ab", // Multiple characters
			description: "option with invalid short name",
			wantErr:     true,
		},
		{
			name:        "invalid short option name with number",
			shortName:   "-1", // Number instead of letter
			description: "option with invalid short name",
			wantErr:     true,
		},
		{
			name:        "invalid short option name with special character",
			shortName:   "-_", // Special character
			description: "option with invalid short name",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := Option{
				Short:       tt.shortName,
				Description: tt.description,
				Type:        TypeString,
			}
			schema := buildSchemaWithOption("-option1", opt)

			err := schema.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSchema_ArgumentNameValidation(t *testing.T) {
	tests := []struct {
		name        string
		argName     string
		description string
		wantErr     bool
	}{
		{
			name:        "valid argument name with letters only",
			argName:     "argument",
			description: "argument with valid name",
			wantErr:     false,
		},
		{
			name:        "valid argument name with letters and numbers",
			argName:     "argument123",
			description: "argument with numbers",
			wantErr:     false,
		},
		{
			name:        "valid argument name with underscore",
			argName:     "argument_name",
			description: "argument with underscore",
			wantErr:     false,
		},
		{
			name:        "valid argument name with multiple underscores",
			argName:     "argument_with_multiple_underscores",
			description: "argument with multiple underscores",
			wantErr:     false,
		},
		{
			name:        "valid argument name with numbers after underscore",
			argName:     "argument_123",
			description: "argument with numbers after underscore",
			wantErr:     false,
		},
		{
			name:        "invalid argument name starting with number",
			argName:     "1argument", // Starts with number
			description: "argument starting with number",
			wantErr:     true,
		},
		{
			name:        "invalid argument name with uppercase letters",
			argName:     "Argument", // Contains uppercase letter
			description: "argument with uppercase",
			wantErr:     true,
		},
		{
			name:        "invalid argument name with dash",
			argName:     "argument-name", // Contains dash
			description: "argument with dash",
			wantErr:     true,
		},
		{
			name:        "invalid argument name with special characters",
			argName:     "argument@name", // Contains special character
			description: "argument with special character",
			wantErr:     true,
		},
		{
			name:        "invalid argument name with underscore at start",
			argName:     "_argument", // Starts with underscore
			description: "argument starting with underscore",
			wantErr:     true,
		},
		{
			name:        "invalid argument name with consecutive underscores",
			argName:     "argument__name", // Contains consecutive underscores
			description: "argument with consecutive underscores",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := Argument{
				Name:        tt.argName,
				Description: tt.description,
				Type:        TypeString,
			}
			schema := buildSchemaWithArgument(arg)

			err := schema.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSchema_ArgumentNameRequired(t *testing.T) {
	tests := []struct {
		name        string
		argName     string
		description string
		wantErr     bool
	}{
		{
			name:        "valid argument with name",
			argName:     "argument",
			description: "argument with name",
			wantErr:     false,
		},
		{
			name:        "invalid argument without name",
			argName:     "", // Empty name
			description: "argument without name",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := Argument{
				Name:        tt.argName,
				Description: tt.description,
				Type:        TypeString,
			}
			schema := buildSchemaWithArgument(arg)

			err := schema.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
