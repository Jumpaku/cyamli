package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		if err != nil {
			return err
		}
		fmt.Printf("%s_%d_%t_%s\n", strings.Join(subcommand, "-"), input.Opt_OptInteger, input.Opt_OptBoolean, input.Opt_OptString)
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
