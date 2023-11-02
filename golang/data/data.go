package data

import (
	"bytes"
	"cliautor"
	"cliautor/name"
	"cliautor/schema"
	"fmt"
)

type Data struct {
	Package          string
	Generator        string
	GeneratorVersion string
	SchemaYAML       string
	Program          Program
	Commands         []Command
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
	data.SchemaYAML = fmt.Sprintf("%q", buffer.String())

	err := walkCmd(nil, s.Program.Command(), func(path []string, cmd *schema.Command) error {
		cmdData := Command{Name: path}
		for optName, opt := range cmd.Options {
			cmdData.Options = append(cmdData.Options, Option{
				Name: name.MakePath(optName),
				Type: opt.Type,
			})
		}
		for _, arg := range cmd.Arguments {
			cmdData.Arguments = append(cmdData.Arguments, Argument{
				Name: name.MakePath(arg.Name),
				Type: arg.Type,
			})
		}
		for subName := range cmd.Subcommands {
			cmdData.Subcommands = append(cmdData.Subcommands, Subcommand{
				Name: name.Path(path).Append(subName),
			})
		}

		if len(path) == 0 {
			data.Program = Program{
				Version:     s.Program.Version,
				Name:        name.MakePath(s.Program.Name),
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
