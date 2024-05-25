package data

import (
	"fmt"
	"slices"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
)

type Data struct {
	Package          string
	Generator        string
	GeneratorVersion string
	SchemaYAML       string
	Program          Program
	Commands         []Command
}

func Construct(s *schema.Schema, generatorName, generatorVersion string, packageName string) (Data, error) {
	data := Data{
		Package:          packageName,
		Generator:        generatorName,
		GeneratorVersion: generatorVersion,
	}

	err := s.Walk(func(path name.Path, cmd *schema.Command) error {
		cmdData := Command{schema: s, schemaCommand: cmd, Name: path}

		for optName, opt := range cmd.Options {
			cmdData.Options = append(cmdData.Options, Option{
				Name:    name.MakePath(optName),
				Short:   name.MakePath(opt.Short),
				Default: opt.Default,
				Type:    opt.Type,
			})
		}
		sortOptions(cmdData.Options)

		for _, arg := range cmd.Arguments {
			cmdData.Arguments = append(cmdData.Arguments, Argument{
				Name:     name.MakePath(arg.Name),
				Type:     arg.Type,
				Variadic: arg.Variadic,
			})
		}

		for subName := range cmd.Subcommands {
			cmdData.Subcommands = append(cmdData.Subcommands, Subcommand{
				Name: path.Append(subName),
			})
		}
		sortSubcommands(cmdData.Subcommands)

		if len(path) == 0 {
			var programName name.Path
			if s.Program.Name != "" {
				programName = programName.Append(s.Program.Name)
			}
			data.Program = Program{
				schema:        s,
				schemaProgram: &s.Program,
				Name:          programName,
				Version:       s.Program.Version,
				Options:       cmdData.Options,
				Arguments:     cmdData.Arguments,
				Subcommands:   cmdData.Subcommands,
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

func sortOptions(options []Option) []Option {
	slices.SortFunc(options, func(a, b Option) int {
		return slices.Compare(a.Name, b.Name)
	})
	return options
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
