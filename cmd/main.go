package main

import (
	"cliautor"
	_ "embed"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

//go:embed example.yaml
var example []byte

func main() {
	cmd, err := cliautor.Load(example)
	fmt.Printf("%+v", err)
	spew.Dump(cmd)
}
