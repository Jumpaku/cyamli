package docs

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"html/template"
)

//go:embed docs.html.tpl
var docsHTMLTemplate string
var executorHTML = template.Must(template.New("docs.html.tpl").Parse(docsHTMLTemplate))

func GenerateHTML(program string, path []string, cmd schema.Command) (string, error) {
	data := Construct(program, path, cmd)
	buf := bytes.NewBuffer(nil)
	if err := executorHTML.Execute(buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template for html: %w", err)
	}
	return buf.String(), nil
}
