package main

import (
	"github.com/goccy/go-yaml"
	"io"
	"log"
	"os"
)

func main() {
	y, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Panicf("%#v", err)
	}

	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		log.Panicf("%#v", err)
	}

	if _, err := os.Stdout.Write(j); err != nil {
		log.Panicf("%#v", err)
	}
}
