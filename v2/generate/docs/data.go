package docs

import (
	"fmt"
	"github.com/Jumpaku/cyamli/v2/docs"
	"github.com/Jumpaku/cyamli/v2/schema"
	"github.com/samber/lo"
	"slices"
)

type Data struct {
	Program     ProgramData
	CommandList []CommandData
}

type ProgramData struct {
	Name    string
	Version string
}

type CommandData struct {
	schema.Command
	Program string
	Path    []string
}

func (d CommandData) DocTextMarkdown() string {
	docTextMarkdown, err := docs.GenerateMarkdown(d.Program, d.Path, d.Command)
	if err != nil {
		panic(fmt.Sprintf("failed to generate markdown doc text for command %#v: %v", d.Path, err))
	}
	return docTextMarkdown
}

func (d CommandData) DocTextHTML() string {
	docTextHTML, err := docs.GenerateHTML(d.Program, d.Path, d.Command)
	if err != nil {
		panic(fmt.Sprintf("failed to generate HTML doc text for command %#v: %v", d.Path, err))
	}
	return docTextHTML
}

func (d CommandData) DocText() string {
	docText, err := docs.GenerateText(d.Program, d.Path, d.Command)
	if err != nil {
		panic(fmt.Sprintf("failed to generate doc text for command %#v: %v", d.Path, err))
	}
	return docText
}

func ConstructData(s schema.Schema) Data {
	commands := s.PropagateOptions().ListCommand()
	commandList := lo.Map(commands, func(cmd schema.PathCommand, _ int) CommandData {
		return CommandData{
			Program: s.Program.Name,
			Command: cmd.Command,
			Path:    cmd.Path,
		}
	})
	slices.SortFunc(commandList, func(a, b CommandData) int { return slices.Compare(a.Path, b.Path) })

	data := Data{
		Program: ProgramData{
			Name:    s.Program.Name,
			Version: s.Program.Version,
		},
		CommandList: commandList,
	}

	return data
}
