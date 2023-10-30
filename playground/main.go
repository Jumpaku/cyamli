package main

import (
	"cliautor"
	"cliautor/golang"
	"io"
	"log"
	"os"
)

func main() {
	exampleYAML, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	schema, err := cliautor.Load(exampleYAML)
	if err != nil {
		log.Panic(err)
	}
	golang.Generate(schema, os.Stdout)

}
