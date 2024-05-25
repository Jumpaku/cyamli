package main

import (
	"bytes"
	"github.com/Jumpaku/cyamli/info"
	"go/format"
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

	buf := bytes.NewBuffer(nil)

	err = golang.Generate(s, info.Name, info.Version, "cyamli", buf)
	if err != nil {
		log.Panicln(err)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Panicln(err)
	}

	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Panicln(err)
	}
}
