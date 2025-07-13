package docs

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"text/template"
)

//go:embed docs.md.tpl
var docsMarkdownTemplate string
var executorMarkdown = template.Must(template.New("docs.md.tpl").Parse(docsMarkdownTemplate))

func GenerateMarkdown(program string, path []string, cmd schema.Command) (string, error) {
	data := Construct(program, path, cmd)
	buf := bytes.NewBuffer(nil)
	if err := executorMarkdown.Execute(buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template for markdown: %w", err)
	}
	return buf.String(), nil
}
