package docs

import (
	"bufio"
	"cmp"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/name"
	"github.com/Jumpaku/cyamli/v2/schema"
	"slices"
	"strings"
)

type CommandData struct {
	Program     string
	Name        name.Name
	Description string
	Options     []OptionData
	Arguments   []ArgumentData
	Subcommands []SubcommandData
}

func (d CommandData) Path() string {
	path := d.Program
	if d.Name.Len() > 0 {
		path += " "
	}
	path += d.Name.Join(" ", "", "")
	return path
}

func (d CommandData) Syntax() string {
	path := d.Program
	if d.Name.Len() > 0 {
		path += " "
	}
	path += d.Name.Join(" ", "", "")
	hasOptions := len(d.Options) > 0
	hasArguments := len(d.Arguments) > 0
	switch {
	case hasOptions && hasArguments:
		return path + ` [<option>|<argument>]... [-- [<argument>]...]`
	case hasOptions:
		return path + ` [<option>]...`
	case hasArguments:
		return path + ` [<argument>]... [-- [<argument>]...]`
	default:
		return path
	}
}

func (d CommandData) DescriptionLines() []string {
	return splitLines(d.Description)
}

type OptionData struct {
	Option      string
	ShortOption string
	Description string
	Type        schema.Type
	Repeated    bool
	Default     string
	Negation    bool
}

func (d OptionData) DescriptionLines() []string {
	return splitLines(d.Description)
}

func (d OptionData) Options() (options []string) {
	options = append(options, d.Option)
	if d.ShortOption != "" {
		options = append(options, d.ShortOption)
	}
	return options
}

func (d OptionData) NegatedOption() (option string) {
	if d.Negation {
		return "-no" + d.Option
	}
	panic(fmt.Sprintf("Option %q cannot not be negated", d.Option))
}

type ArgumentData struct {
	Position    int
	Name        name.Name
	Description string
	Type        schema.Type
	Variadic    bool
}

func (d ArgumentData) DescriptionLines() []string {
	return splitLines(d.Description)
}

type SubcommandData struct {
	Name        string
	Description string
}

func (d SubcommandData) DescriptionLines() []string {
	return splitLines(d.Description)
}

func Construct(program string, path []string, cmd schema.Command) CommandData {
	data := CommandData{
		Program:     program,
		Name:        name.Words(path),
		Description: cmd.Description,
	}
	if data.Program == "" {
		data.Program = "<program>"
	}

	{
		for optionName, option := range cmd.Options {
			optionData := OptionData{
				Option:      optionName,
				ShortOption: option.Short,
				Description: option.Description,
				Type:        option.Type,
				Repeated:    option.Repeated,
				Default:     option.Default,
				Negation:    option.Negation,
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
			case schema.TypeString, schema.TypeEmpty:
				optionData.Type = schema.TypeString
				optionData.Default = fmt.Sprintf("%q", option.Default)
			}

			data.Options = append(data.Options, optionData)
		}
		slices.SortFunc(data.Options, func(a, b OptionData) int { return cmp.Compare(a.Option, b.Option) })
	}
	{
		for idx, argument := range cmd.Arguments {
			argumentData := ArgumentData{
				Position:    idx + 1,
				Name:        name.New(argument.Name),
				Description: argument.Description,
				Type:        argument.Type,
				Variadic:    argument.Variadic,
			}

			if argumentData.Type == schema.TypeEmpty {
				argumentData.Type = schema.TypeString
			}

			data.Arguments = append(data.Arguments, argumentData)
		}
	}
	{
		for subcommandName, subcommand := range cmd.Subcommands {
			data.Subcommands = append(data.Subcommands, SubcommandData{
				Name:        subcommandName,
				Description: subcommand.Description,
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
