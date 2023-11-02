package data

import (
	"bytes"
	"cliautor"
	"cliautor/name"
	"cliautor/schema"
	"fmt"
	"slices"
)

type Data struct {
	Package          string
	Generator        string
	GeneratorVersion string
	SchemaYAML       string
	Program          Program
	Commands         []Command
}

func (d Data) SchemaYAMLLiteral() string {
	return fmt.Sprintf("%q", d.SchemaYAML)
}

func Construct(packageName string, s *schema.Schema) (Data, error) {
	data := Data{
		Package:          packageName,
		Generator:        cliautor.Name,
		GeneratorVersion: cliautor.Version,
	}

	buffer := bytes.NewBuffer(nil)
	if err := s.Save(buffer); err != nil {
		return Data{}, fmt.Errorf("fail to create CLI data: %w", err)
	}
	data.SchemaYAML = buffer.String()

	err := walkCmd(nil, s.Program.Command(), func(path []string, cmd *schema.Command) error {
		cmdData := Command{
			Name:        path,
			Description: cmd.Description,
		}
		for optName, opt := range cmd.Options {
			cmdData.Options = append(cmdData.Options, Option{
				Name:        name.MakePath(optName),
				Short:       name.MakePath(opt.Short),
				Description: opt.Description,
				Default:     opt.Default,
				Type:        opt.Type,
			})
		}
		sortOptions(cmdData.Options)

		for _, arg := range cmd.Arguments {
			cmdData.Arguments = append(cmdData.Arguments, Argument{
				Name:        name.MakePath(arg.Name),
				Description: arg.Description,
				Type:        arg.Type,
				Variadic:    arg.Variadic,
			})
		}
		sortArguments(cmdData.Arguments)

		for subName := range cmd.Subcommands {
			cmdData.Subcommands = append(cmdData.Subcommands, Subcommand{
				Name: name.Path(path).Append(subName),
			})
		}
		sortSubcommands(cmdData.Subcommands)

		if len(path) == 0 {
			var programName name.Path
			if s.Program.Name != "" {
				programName = programName.Append(s.Program.Name)
			}
			data.Program = Program{
				Name:        programName,
				Version:     s.Program.Version,
				Description: cmd.Description,
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
		return Data{}, fmt.Errorf("fail to create CLI data: %w", err)
	}

	sortCommands(data.Commands)

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

func sortOptions(options []Option) []Option {
	slices.SortFunc(options, func(a, b Option) int {
		return slices.Compare(a.Name, b.Name)
	})
	return options
}

func sortArguments(arguments []Argument) []Argument {
	slices.SortFunc(arguments, func(a, b Argument) int {
		return slices.Compare(a.Name, b.Name)
	})
	return arguments
}

func sortCommands(commands []Command) []Command {
	slices.SortFunc(commands, func(a, b Command) int {
		return slices.Compare(a.Name, b.Name)
	})
	return commands
}

func sortSubcommands(subcommands []Subcommand) []Subcommand {
	slices.SortFunc(subcommands, func(a, b Subcommand) int {
		return slices.Compare(a.Name, b.Name)
	})
	return subcommands
}
