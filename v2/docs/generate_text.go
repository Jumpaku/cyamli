package docs

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/schema"
	"text/template"
)

//go:embed docs.txt.tpl
var docsTextTemplate string
var executorText = template.Must(template.New("docs.txt.tpl").Parse(docsTextTemplate))

func GenerateText(program string, path []string, cmd schema.Command) (string, error) {
	data := Construct(program, path, cmd)
	buf := bytes.NewBuffer(nil)
	if err := executorText.Execute(buf, data); err != nil {
		return "", fmt.Errorf("fail to execute template for text: %w", err)
	}
	return buf.String(), nil
}
