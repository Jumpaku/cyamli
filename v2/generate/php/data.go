package php

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
	Namespace   string
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

	Namespace string
	Generator string

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
		return "Run"
	}
	return "Run_" + d.Name.UpperCamel()
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
	Negation     bool
	Repeated     bool
	DefaultValue string
}

func (d OptionData) InputFieldName() string {
	return "Opt_" + d.Name.UpperCamel()
}

func (d OptionData) InputFieldType() string {
	if d.Repeated {
		return "array"
	}
	return primitiveType(d.Type)
}

func (d OptionData) InputFieldDocType() string {
	typ := primitiveType(d.Type)
	if d.Repeated {
		return typ + "[]"
	}
	return typ
}

func (d OptionData) InputFieldInit() string {
	typ := primitiveType(d.Type)
	if d.Repeated {
		return "[]"
	}
	switch typ {
	case "bool":
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
	case "string":
		if d.DefaultValue == "" {
			return `""`
		}
		return fmt.Sprintf("%q", d.DefaultValue)
	case "int":
		if d.DefaultValue == "" {
			return "0"
		}
		v, err := strconv.ParseInt(d.DefaultValue, 0, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse %q as int64", d.DefaultValue))
		}
		return fmt.Sprintf("%d", v)
	}
	return typ
}

type ArgumentData struct {
	Name     name.Name
	Type     schema.Type
	Variadic bool
}

func (d ArgumentData) InputFieldName() string {
	return "Arg_" + d.Name.UpperCamel()
}

func (d ArgumentData) InputFieldType() string {
	if d.Variadic {
		return "array"
	}
	return primitiveType(d.Type)
}

func (d ArgumentData) InputFieldDocType() string {
	typ := primitiveType(d.Type)
	if d.Variadic {
		return typ + "[]"
	}
	return typ
}

func primitiveType(t schema.Type) string {
	switch t {
	case schema.TypeString, schema.TypeEmpty:
		return "string"
	case schema.TypeInteger:
		return "int"
	case schema.TypeBoolean:
		return "bool"
	default:
		panic("unexpected type: " + string(t))
	}
}

func ConstructData(s schema.Schema, namespace, generatorName string) Data {
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
			Command:   cmd.Command,
			Namespace: namespace,
			Generator: generatorName,
			Program:   s.Program.Name,
			Name:      name.Words(cmd.Path),
			Options:   options,
			Arguments: arguments,
		}
	})
	slices.SortFunc(commandList, func(a, b CommandData) int { return a.Name.Cmp(b.Name) })

	data := Data{
		Namespace: namespace,
		Generator: generatorName,
		Program: ProgramData{
			Name:    s.Program.Name,
			Version: s.Program.Version,
		},
		CommandList: commandList,
	}

	return data
}
