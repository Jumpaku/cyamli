package description

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed detail.tpl
var detailTemplate string

func DetailExecutor() *template.Template {
	return template.Must(template.New("detail.tpl").Parse(detailTemplate))
}

//go:embed simple.tpl
var simpleTemplate string

func SimpleExecutor() *template.Template {
	return template.Must(template.New("simple.tpl").Parse(simpleTemplate))
}

type Executor interface {
	Execute(wr io.Writer, data any) error
}

func DescribeCommand(exec Executor, data CommandData, writer io.Writer) error {
	if err := exec.Execute(writer, data); err != nil {
		return fmt.Errorf("fail to execute template: %w", err)
	}

	return nil
}
