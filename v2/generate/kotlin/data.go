package kotlin

import (
	"fmt"
	"github.com/Jumpaku/cyamli/v2/docs"
	"github.com/Jumpaku/cyamli/v2/name"
	"github.com/Jumpaku/cyamli/v2/schema"
	"github.com/samber/lo"
	"slices"
	"strconv"
	"strings"
)

type Data struct {
	Package     string
	Generator   string
	Program     ProgramData
	CommandList []CommandData
}

type ProgramData struct {
	Name    string
	Version string
}

type CommandData struct {
	schema.Command
	Program   string
	Name      name.Name
	Options   []OptionData
	Arguments []ArgumentData
}

func (d CommandData) Path() []string {
	path := []string{}
	for i := 0; i < d.Name.Len(); i++ {
		path = append(path, d.Name.Get(i).LowerCamel())
	}
	return path
}

func (d CommandData) PathLiteral() string {
	return fmt.Sprintf("%q", d.Name.Map(strings.ToLower).Join(" ", "", ""))
}
func (d CommandData) HandlerMethodName() string {
	if d.Name.Len() == 0 {
		return "execute"
	}
	return "execute_" + d.Name.UpperCamel()
}

func (d CommandData) HandlerInputType() string {
	if d.Name.Len() == 0 {
		return "Input"
	}
	return "Input_" + d.Name.UpperCamel()
}

func (d CommandData) DocText() string {
	path := strings.Split(d.Name.Map(strings.ToLower).Join(" ", "", ""), " ")
	txt := lo.Must(docs.GenerateText(d.Program, path, d.Command))
	return fmt.Sprintf("%q", txt)
}

type OptionData struct {
	Name         name.Name
	Option       string
	ShortOption  string
	Type         schema.Type
	Repeated     bool
	DefaultValue string
	Negation     bool
}

func (d OptionData) InputFieldName() string {
	return "opt_" + d.Name.UpperCamel()
}

func (d OptionData) InputFieldType() string {
	typ := primitiveType(d.Type)
	if d.Repeated {
		// Return the full type name with angle brackets directly in the template
		return "List<" + typ + ">"
	}
	return typ
}

func (d OptionData) InputFieldInit() string {
	typ := primitiveType(d.Type)
	if d.Repeated {
		return "listOf()"
	}
	switch typ {
	case "Boolean":
		switch strings.ToLower(d.DefaultValue) {
		default:
			panic(fmt.Sprintf("failed to parse %q as bool", d.DefaultValue))
		case "":
			return "false"
		case "true", "True", "1":
			return "true"
		case "false", "False", "0":
			return "false"
		}
	case "String":
		if d.DefaultValue == "" {
			return `""`
		}
		return fmt.Sprintf("%q", d.DefaultValue)
	case "Long":
		if d.DefaultValue == "" {
			return "0L"
		}
		v, err := strconv.ParseInt(d.DefaultValue, 0, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse %q as int64", d.DefaultValue))
		}
		return fmt.Sprintf("%dL", v)
	}
	return typ
}

type ArgumentData struct {
	Name     name.Name
	Type     schema.Type
	Variadic bool
}

func (d ArgumentData) InputFieldName() string {
	return "arg_" + d.Name.UpperCamel()
}

func (d ArgumentData) InputFieldType() string {
	typ := primitiveType(d.Type)
	if d.Variadic {
		// Return the full type name with angle brackets directly in the template
		return "List<" + typ + ">"
	}
	return typ
}

func primitiveType(t schema.Type) string {
	switch t {
	case schema.TypeString, schema.TypeEmpty:
		return "String"
	case schema.TypeInteger:
		return "Long"
	case schema.TypeBoolean:
		return "Boolean"
	default:
		panic("unexpected type: " + string(t))
	}
}

func ConstructData(s schema.Schema, packageName, generatorName string) Data {
	commands := s.PropagateOptions().ListCommand()
	commandList := lo.Map(commands, func(cmd schema.PathCommand, _ int) CommandData {
		options := []OptionData{}
		for option, o := range cmd.Command.Options {
			options = append(options, OptionData{
				Option:       option,
				ShortOption:  o.Short,
				Name:         name.New(option),
				Type:         o.Type,
				Repeated:     o.Repeated,
				DefaultValue: o.Default,
				Negation:     o.Negation,
			})
		}
		slices.SortFunc(options, func(a, b OptionData) int { return a.Name.Cmp(b.Name) })

		arguments := []ArgumentData{}
		for _, argument := range cmd.Command.Arguments {
			arguments = append(arguments, ArgumentData{
				Name:     name.New(argument.Name),
				Type:     argument.Type,
				Variadic: argument.Variadic,
			})
		}
		slices.SortFunc(arguments, func(a, b ArgumentData) int { return a.Name.Cmp(b.Name) })

		return CommandData{
			Program:   s.Program.Name,
			Command:   cmd.Command,
			Name:      name.Words(cmd.Path),
			Options:   options,
			Arguments: arguments,
		}
	})
	slices.SortFunc(commandList, func(a, b CommandData) int { return a.Name.Cmp(b.Name) })

	data := Data{
		Package:   packageName,
		Generator: generatorName,
		Program: ProgramData{
			Name:    s.Program.Name,
			Version: s.Program.Version,
		},
		CommandList: commandList,
	}

	return data
}
