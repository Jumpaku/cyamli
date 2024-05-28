package main

import (
	"flag"
	"github.com/goccy/go-yaml"
	"io"
	"log"
	"os"
)

var (
	in  = flag.String("in", "", "input file path")
	out = flag.String("out", "", "output file path")
)

func main() {
	reader := os.Stdin
	if *in != "" {
		f, err := os.Open(*in)
		if err != nil {
			log.Panicf("%#v", err)
		}
		defer f.Close()
		reader = f

	}

	writer := os.Stdout
	if *out != "" {
		f, err := os.Create(*out)
		if err != nil {
			log.Panicf("%#v", err)
		}
		defer f.Close()
		writer = f
	}

	y, err := io.ReadAll(reader)
	if err != nil {
		log.Panicf("%#v", err)
	}

	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		log.Panicf("%#v", err)
	}

	if _, err := writer.Write(j); err != nil {
		log.Panicf("%#v", err)
	}
}
