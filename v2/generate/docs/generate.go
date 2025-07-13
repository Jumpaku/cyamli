package docs

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"io"
	"text/template"
)

//go:embed cli_doc.gen.html.tpl
var cliDocGenHTMLTemplate string
var executorHTML = template.Must(template.New("cli_doc.gen.html.tpl").Parse(cliDocGenHTMLTemplate))

//go:embed cli_doc.gen.md.tpl
var cliDocGenMDTemplate string
var executorMD = template.Must(template.New("cli_doc.gen.md.tpl").Parse(cliDocGenMDTemplate))

//go:embed cli_doc.gen.txt.tpl
var cliDocGenTXTTemplate string
var executorTXT = template.Must(template.New("cli_doc.gen.txt.tpl").Parse(cliDocGenTXTTemplate))

type DocFormat string

const (
	DocFormatHTML     DocFormat = "html"
	DocFormatMarkdown DocFormat = "markdown"
	DocFormatText     DocFormat = "text"
)

func Generate(schema schema.Schema, format DocFormat, out io.Writer) error {
	d := ConstructData(schema)

	buf := bytes.NewBuffer(nil)
	switch format {
	default:
		return fmt.Errorf("unsupported doc format: %s", format)
	case DocFormatHTML:
		if err := executorHTML.Execute(buf, d); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	case DocFormatMarkdown:
		if err := executorMD.Execute(buf, d); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	case DocFormatText:
		if err := executorTXT.Execute(buf, d); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	}
	if _, err := out.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write generated code: %w", err)
	}
	return nil
}
