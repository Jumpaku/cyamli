package main

import (
	"cliautor/golang"
	"cliautor/schema"
	"log"
	"os"
)

func main() {
	schema, err := schema.Load(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	err = golang.Generate("main", schema, os.Stdout)
	if err != nil {
		log.Panic(err)
	}
}
