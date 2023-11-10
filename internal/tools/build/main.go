package main

import (
	"log"
	"os"

	"github.com/Jumpaku/cyamli/golang"
	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	err = golang.Generate("main", s, os.Stdout)
	if err != nil {
		log.Panicln(err)
	}
}
