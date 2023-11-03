package main

import "os"

func main() {
	cli := NewCLI()
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
