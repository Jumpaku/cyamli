package docs

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/docs/data"
	"github.com/Jumpaku/cyamli/name"
	"github.com/samber/lo"
	"github.com/yosssi/gohtml"
	template2 "html/template"
	"io"
	"slices"
	"text/template"

	"github.com/Jumpaku/cyamli/schema"
)

var funcs = template.FuncMap{
	"func_add": func(a, b int) int { return a + b },
	"func_sub": func(a, b int) int { return a - b },
}

//go:embed docs.txt.tpl
var docsTextTemplate string
var executorText = func() *template.Template {
	e := template.New("docs.txt.tpl")
	e.Funcs(funcs)
	return template.Must(e.Parse(docsTextTemplate))
}()

//go:embed docs.md.tpl
var docsMarkdownTemplate string
var executorMarkdown = func() *template.Template {
	e := template.New("docs.md.tpl")
	e.Funcs(funcs)
	return template.Must(e.Parse(docsMarkdownTemplate))
}()

//go:embed docs.html.tpl
var docsHTMLTemplate string
var executorHTML = func() *template2.Template {
	e := template2.New("docs.html.tpl")
	e.Funcs(funcs)
	return template2.Must(e.Parse(docsHTMLTemplate))
}()

type DocsFormat string

const (
	DocsFormatUnspecified DocsFormat = ""
	DocsFormatText        DocsFormat = "text"
	DocsFormatMarkdown    DocsFormat = "markdown"
	DocsFormatHTML        DocsFormat = "html"
)

type GenerateArgs struct {
	Format     DocsFormat
	All        bool
	Subcommand name.Path
}

func Generate(schema *schema.Schema, args GenerateArgs, out io.Writer) error {
	d := data.Construct(schema)

	if !args.All {
		d.Commands = lo.Filter(d.Commands, func(item data.CommandData, _ int) bool {
			return slices.Equal(item.Path, args.Subcommand)
		})
	}
	if len(d.Commands) == 0 {
		return fmt.Errorf(`selected subcommand not found: %q`, args.Subcommand.Join(" ", "", ""))
	}

	switch args.Format {
	default:
		return fmt.Errorf(`unsupported format: %q`, args.Format)
	case DocsFormatText, DocsFormatUnspecified:
		if err := executorText.Execute(out, d); err != nil {
			return fmt.Errorf("fail to execute template for text: %w", err)
		}
	case DocsFormatMarkdown:
		if err := executorMarkdown.Execute(out, d); err != nil {
			return fmt.Errorf("fail to execute template for markdown: %w", err)
		}
	case DocsFormatHTML:
		buf := bytes.NewBuffer(nil)
		if err := executorHTML.Execute(buf, d); err != nil {
			return fmt.Errorf("fail to execute template for HTML: %w", err)
		}

		b := gohtml.FormatBytes(buf.Bytes())

		if _, err := io.Copy(out, bytes.NewBuffer(b)); err != nil {
			return fmt.Errorf("fail to format generated HTML: %w", err)
		}
	}

	return nil
}
