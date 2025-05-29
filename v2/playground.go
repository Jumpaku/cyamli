package main

import (
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/generate/golang"
	"github.com/Jumpaku/cyamli/v2/schema"
	"os"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("fail to load schema: %+v", err))
	}
	if err = golang.Generate(s, "main", "cyamli", os.Stdout); err != nil {
		panic(err)
	}
}
