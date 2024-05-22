package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		spew.Dump(inputErr)
		spew.Dump(input)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.List.FUNC = func(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
		spew.Dump(inputErr)
		spew.Dump(input)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.Fetch.FUNC = func(subcommand []string, input CLI_Fetch_Input, inputErr error) (err error) {
		spew.Dump(inputErr)
		spew.Dump(input)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		log.Panicln(err)
	}
}
