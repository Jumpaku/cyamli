package docs

import (
	"bytes"
	_ "embed"
	"encoding/xml"
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
		return "", fmt.Errorf("fail to execute template for text: %w", err)
	}

	var v any
	if err := xml.Unmarshal(buf.Bytes(), &v); err != nil {
		return "", fmt.Errorf("fail to unmarshal generated HTML: %w", err)
	}
	b, err := xml.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("fail to marshal generated HTML: %w", err)
	}

	return string(b), nil
}
