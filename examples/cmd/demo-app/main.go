package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		fmt.Printf("%+v\n", inputErr)
		fmt.Printf("%+v\n", input)
		return nil
	}
	cli.List.FUNC = func(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
		fmt.Printf("%+v\n", inputErr)
		fmt.Printf("%+v\n", input)
		return nil
	}
	cli.Fetch.FUNC = func(subcommand []string, input CLI_Fetch_Input, inputErr error) (err error) {
		fmt.Printf("%+v\n", inputErr)
		fmt.Printf("%+v\n", input)
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		log.Panicln(err)
	}
}
