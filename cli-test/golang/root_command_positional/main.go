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
		fmt.Printf("%s_%d_%t_%s_%s\n", strings.Join(subcommand, "-"), input.Arg_ArgInteger, input.Arg_ArgBoolean, input.Arg_ArgString, strings.Join(input.Arg_ArgVariadic, ","))
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
