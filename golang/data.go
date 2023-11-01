package golang

import (
	"cliautor"
	"fmt"
	"strings"

	"github.com/Jumpaku/go-assert"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

type CLIData struct {
	Package          string
	Generator        string
	GeneratorVersion string
	SchemaYAML       string
	Program          CLIProgramData
	Commands         []CLICmdData
}

type CLIProgramData struct {
	Name        string
	Version     string
	Options     []CLIOptData
	Args        []CLIArgData
	Subcommands []CLISubCmdData
}

type CLICmdData struct {
	CLISubCmdData
	Options     []CLIOptData
	Args        []CLIArgData
	Subcommands []CLISubCmdData
}

type CLISubCmdData struct {
	Path []string
}

func (d CLISubCmdData) Identifier() string {
	if len(d.Path) == 0 {
		return "Cmd"
	}
	name := d.Path[len(d.Path)-1]
	return "Cmd_" + Title(strings.Join(cliautor.MakeName(name), ""))
}

func (d CLISubCmdData) FullIdentifier() string {
	if len(d.Path) == 0 {
		return "Cmd"
	}
	titled := lo.Map(d.Path, func(p string, i int) string { return Title(strings.Join(cliautor.MakeName(p), "")) })
	return "Cmd_" + strings.Join(titled, "")
}

func (d CLISubCmdData) FuncIdentifier() string {
	if len(d.Path) == 0 {
		return "Func"
	}
	titled := lo.Map(d.Path, func(p string, i int) string { return "Cmd_" + Title(strings.Join(cliautor.MakeName(p), "")) })
	return strings.Join(titled, ".") + ".Func"
}

type CLIOptData struct {
	Name   string
	GoType string
}

func (d CLIOptData) Identifier() string {
	titled := lo.Map(cliautor.MakeName(d.Name), func(p string, i int) string { return Title(p) })
	return "Opt_" + strings.Join(titled, "")
}

type CLIArgData struct {
	Name   string
	GoType string
}

func (d CLIArgData) Identifier() string {
	return "Arg_" + d.Name
}

func createCLIData(schema *cliautor.Schema) (CLIData, error) {
	data := CLIData{
		Package:          "cmd",
		Generator:        "cliautor",
		GeneratorVersion: "0.0.0",
	}

	schemaYAMLBytes, err := yaml.Marshal(schema)
	if err != nil {
		return CLIData{}, fmt.Errorf("fail to create CLI data: %w", err)
	}
	data.SchemaYAML = fmt.Sprintf("%q", string(schemaYAMLBytes))

	err = walkCmd(nil, schema.Program.Command(), func(path []string, cmd *cliautor.Command) error {
		cmdData := CLICmdData{
			CLISubCmdData: CLISubCmdData{Path: path},
		}
		for name, opt := range cmd.Options {
			cmdData.Options = append(cmdData.Options, CLIOptData{
				Name:   name,
				GoType: goType(opt.Type),
			})
		}
		for _, arg := range cmd.Arguments {
			cmdData.Args = append(cmdData.Args, CLIArgData{
				Name:   arg.Name,
				GoType: goType(arg.Type),
			})
		}
		for name := range cmd.Subcommands {
			cmdData.Subcommands = append(cmdData.Subcommands, CLISubCmdData{
				Path: append(append([]string{}, path...), name),
			})
		}

		if len(path) == 0 {
			data.Program = CLIProgramData{
				Version:     schema.Program.Version,
				Name:        schema.Program.Name,
				Options:     cmdData.Options,
				Args:        cmdData.Args,
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

func walkCmd(path []string, cmd *cliautor.Command, f func(path []string, cmd *cliautor.Command) error) error {
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

func goType(t cliautor.Type) string {
	switch t {
	default:
		return assert.Unexpected1[string]("unexpected type: %s", t)
	case cliautor.TypeBoolean:
		return "bool"
	case cliautor.TypeFloat:
		return "float64"
	case cliautor.TypeInteger:
		return "int64"
	case cliautor.TypeString, cliautor.TypeUnspecified:
		return "string"
	}
}
