package main

import (
	"cliautor/schema"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	schema, err := schema.Load(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	spew.Dump(schema)
}
