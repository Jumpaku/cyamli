package main

import (
	"github.com/Jumpaku/cyamli/cyamli"
	"os"
)

func main() {
	os.Exit(cyamli.Execute(os.Args, os.Stdin, os.Stdout, os.Stderr))
}
