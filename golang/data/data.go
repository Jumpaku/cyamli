package data

import (
	"cliautor/name"
	"cliautor/schema"
	"fmt"

	"github.com/Jumpaku/go-assert"
	"gopkg.in/yaml.v3"
)

type CLIData struct {
	Package          string
	Generator        string
	GeneratorVersion string
	SchemaYAML       string
	Program          Program
	Commands         []Command
}

func Construct(s *schema.Schema) (CLIData, error) {
	data := CLIData{
		Package:          "cmd",
		Generator:        "cliautor",
		GeneratorVersion: "0.0.0",
	}

	schemaYAMLBytes, err := yaml.Marshal(s)
	if err != nil {
		return CLIData{}, fmt.Errorf("fail to create CLI data: %w", err)
	}
	data.SchemaYAML = fmt.Sprintf("%q", string(schemaYAMLBytes))

	err = walkCmd(nil, s.Program.Command(), func(path []string, cmd *schema.Command) error {
		cmdData := Command{Name: path}
		for optName, opt := range cmd.Options {
			cmdData.Options = append(cmdData.Options, Option{
				Name:   name.MakeName(optName),
				GoType: goType(opt.Type),
			})
		}
		for _, arg := range cmd.Arguments {
			cmdData.Arguments = append(cmdData.Arguments, Argument{
				Name:   name.MakeName(arg.Name),
				GoType: goType(arg.Type),
			})
		}
		for subName := range cmd.Subcommands {
			cmdData.Subcommands = append(cmdData.Subcommands, Subcommand{
				Name: name.Name(path).Append(subName),
			})
		}

		if len(path) == 0 {
			data.Program = Program{
				Version:     s.Program.Version,
				Name:        s.Program.Name,
				Options:     cmdData.Options,
				Arguments:   cmdData.Arguments,
				Subcommands: cmdData.Subcommands,
			}
		} else {
			data.Commands = append(data.Commands, cmdData)
		}

		return nil
	})

	if err != nil {
		return CLIData{}, fmt.Errorf("fail to create CLI data: %w", err)
	}

	return data, nil
}

func walkCmd(path []string, cmd *schema.Command, f func(path []string, cmd *schema.Command) error) error {
	if err := f(path, cmd); err != nil {
		return err
	}
	for name, cmd := range cmd.Subcommands {
		if err := walkCmd(append(append([]string{}, path...), name), cmd, f); err != nil {
			return err
		}
	}
	return nil
}

func goType(t schema.Type) string {
	switch t {
	default:
		return assert.Unexpected1[string]("unexpected type: %s", t)
	case schema.TypeBoolean:
		return "bool"
	case schema.TypeFloat:
		return "float64"
	case schema.TypeInteger:
		return "int64"
	case schema.TypeString, schema.TypeUnspecified:
		return "string"
	}
}
