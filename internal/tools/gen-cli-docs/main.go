package main

import (
	"github.com/Jumpaku/cyamli/docs"
	"log"
	"os"

	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	err = docs.Generate(s, docs.GenerateArgs{
		Format:     docs.DocsFormatMarkdown,
		All:        true,
		Subcommand: nil,
	}, os.Stdout)
	if err != nil {
		log.Panicln(err)
	}
}
