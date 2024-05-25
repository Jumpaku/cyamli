package data

import (
	"bufio"
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
)

type Data struct {
	Program  string
	Version  string
	Commands []CommandData
}

type CommandData struct {
	Path        name.Path
	Syntax      string
	Description []string
	Options     []OptionData
	Arguments   []ArgumentData
	Subcommands []SubcommandData
}

type OptionData struct {
	Names       []string
	Description []string
	Type        schema.Type
	Default     string
}
type ArgumentData struct {
	Name        string
	Description []string
	Type        schema.Type
	Variadic    bool
}
type SubcommandData struct {
	Name        string
	Description []string
}

func Construct(s *schema.Schema) Data {
	data := Data{
		Program: s.Program.Name,
		Version: s.Program.Version,
	}
	if data.Program == "" {
		data.Program = "<program>"
	}
	_ = s.Walk(func(path name.Path, cmd *schema.Command) error {
		data.Commands = append(data.Commands, createCommandData(path, cmd))
		return nil
	})
	slices.SortFunc(data.Commands, func(a, b CommandData) int {
		return cmp.Compare(a.Path.Join(" ", "", ""), b.Path.Join(" ", "", ""))
	})
	return data
}

func createCommandData(subcommand name.Path, cmd *schema.Command) CommandData {
	data := CommandData{Path: subcommand, Description: splitLines(cmd.Description)}
	{
		hasOptions := len(cmd.Options) > 0
		hasArguments := len(cmd.Arguments) > 0
		syntax := subcommand.Join(" ", "", "")
		switch {
		case hasOptions && hasArguments:
			syntax += ` [<option>|<argument>]... [-- [<argument>]...]`
		case hasOptions:
			syntax += ` [<option>]...`
		case hasArguments:
			syntax += ` [<argument>]... [-- [<argument>]...]`
		}
		data.Syntax = syntax
	}
	{
		for optionName, option := range cmd.Options {
			optionData := OptionData{
				Names:       []string{optionName},
				Description: splitLines(option.Description),
				Type:        option.Type,
				Default:     option.Default,
			}

			if option.Short != "" {
				optionData.Names = append(optionData.Names, option.Short)
			}

			switch option.Type {
			case schema.TypeBoolean:
				if optionData.Default == "" {
					optionData.Default = "false"
				}
			case schema.TypeInteger:
				if optionData.Default == "" {
					optionData.Default = "0"
				}
			case schema.TypeFloat:
				if optionData.Default == "" {
					optionData.Default = "0.0"
				}
			case schema.TypeString, schema.TypeUnspecified:
				optionData.Type = schema.TypeString
				optionData.Default = fmt.Sprintf("%q", option.Default)
			}

			data.Options = append(data.Options, optionData)
		}
		slices.SortFunc(data.Options, func(a, b OptionData) int { return cmp.Compare(a.Names[0], b.Names[0]) })
	}
	{
		for _, argument := range cmd.Arguments {
			argumentData := ArgumentData{
				Name:        argument.Name,
				Description: splitLines(argument.Description),
				Type:        argument.Type,
				Variadic:    argument.Variadic,
			}

			if argumentData.Type == schema.TypeUnspecified {
				argumentData.Type = schema.TypeString
			}

			data.Arguments = append(data.Arguments, argumentData)
		}
	}
	{
		for subcommandName, subcommand := range cmd.Subcommands {
			data.Subcommands = append(data.Subcommands, SubcommandData{
				Name:        subcommandName,
				Description: splitLines(subcommand.Description),
			})
		}
		slices.SortFunc(data.Subcommands, func(a, b SubcommandData) int { return cmp.Compare(a.Name, b.Name) })
	}
	return data
}

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
