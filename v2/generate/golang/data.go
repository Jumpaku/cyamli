package golang

import (
	"fmt"
	"github.com/Jumpaku/cyamli/v2/docs"
	"github.com/Jumpaku/cyamli/v2/name"
	"github.com/Jumpaku/cyamli/v2/schema"
	"github.com/emirpasic/gods/trees/redblacktree"
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
	Repeated     bool
	DefaultValue string
}

func (d OptionData) InputFieldName() string {
	return "Opt_" + d.Name.UpperCamel()
}

func (d OptionData) InputFieldType() string {
	typ := primitiveType(d.Type)
	if d.Repeated {
		return "[]" + typ
	}
	return typ
}

func (d OptionData) InputFieldInit() string {
	typ := primitiveType(d.Type)
	if d.Repeated {
		return "[]" + typ + "{}"
	}
	switch typ {
	case "bool":
		if d.DefaultValue == "" {
			return "false"
		}
		v, err := strconv.ParseBool(d.DefaultValue)
		if err != nil {
			panic(fmt.Sprintf("failed to parse %q as bool", d.DefaultValue))
		}
		return fmt.Sprintf("%t", v)
	case "string":
		if d.DefaultValue == "" {
			return `""`
		}
		return fmt.Sprintf("%q", d.DefaultValue)
	case "int64":
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
	typ := primitiveType(d.Type)
	if d.Variadic {
		return "[]" + typ
	}
	return typ
}

func primitiveType(t schema.Type) string {
	switch t {
	case schema.TypeString, schema.TypeEmpty:
		return "string"
	case schema.TypeInteger:
		return "int64"
	case schema.TypeBoolean:
		return "bool"
	default:
		panic("unexpected type: " + string(t))
	}
}

func ConstructData(s schema.Schema, packageName, generatorName string) Data {
	commands := listCommands(s)
	commandList := lo.Map(commands, func(cmd pathCommand, _ int) CommandData {
		options := []OptionData{}
		for option, o := range cmd.Command.Options {
			options = append(options, OptionData{
				Option:       option,
				ShortOption:  o.Short,
				Name:         name.New(option),
				Type:         o.Type,
				Repeated:     o.Repeated,
				DefaultValue: o.Default,
			})
		}
		for option, o := range cmd.PropagatedOptions {
			options = append(options, OptionData{
				Option:       option,
				ShortOption:  o.Short,
				Name:         name.New(option),
				Type:         o.Type,
				Repeated:     o.Repeated,
				DefaultValue: o.Default,
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

type pathCommand struct {
	Path              []string
	Command           schema.Command
	PropagatedOptions map[string]schema.Option
}

func listCommands(s schema.Schema) (commands []pathCommand) {
	tree := redblacktree.NewWith(func(a, b interface{}) int {
		return slices.Compare(a.([]string), b.([]string))
	})

	for _, cmd := range s.ListCommand() {
		tree.Put(cmd.Path, &pathCommand{
			Path:              cmd.Path,
			Command:           cmd.Command,
			PropagatedOptions: make(map[string]schema.Option),
		})
	}

	itr := tree.Iterator()
	for itr.Next() {
		path := itr.Key().([]string)
		cmd := itr.Value().(*pathCommand)
		if len(path) > 0 {
			parentCmd, _ := tree.Get(path[:len(path)-1])
			for name, option := range parentCmd.(*pathCommand).Command.Options {
				if option.Propagates {
					cmd.PropagatedOptions[name] = option
				}
			}
			for name, option := range parentCmd.(*pathCommand).PropagatedOptions {
				cmd.PropagatedOptions[name] = option
			}
		}

		commands = append(commands, *cmd)
	}

	return commands
}
