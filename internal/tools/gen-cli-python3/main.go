package main

import (
	"log"
	"os"

	"github.com/Jumpaku/cyamli/python3"
	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	if err := python3.Generate(s, os.Stdout); err != nil {
		log.Panicln(err)
	}
}
