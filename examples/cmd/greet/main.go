package main

import (
	"fmt"
	"os"
)

//go:generate go run "github.com/Jumpaku/cyamli/cmd/cyamli@latest" golang -schema-path=cli.yaml -out-path=cli.gen.go

// Create the CLI object
var cli = NewCLI()

func main() {
	// Overwrite behaviors
	cli.FUNC = showHelp
	cli.Hello.FUNC = sayHello
	// Run with command line arguments
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func showHelp(subcommand []string, input CLI_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(cli.DESC_Simple())
		panic(inputErr)
	}
	if input.Opt_Help {
		fmt.Println(cli.DESC_Detail())
	}
	return nil
}
func sayHello(subcommand []string, input CLI_Hello_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(cli.Hello.DESC_Simple())
		return inputErr
	}
	if input.Opt_TargetName != "" {
		fmt.Printf("Hello, %s! My name is %s!\n", input.Opt_TargetName, input.Arg_Greeter)
	} else {
		fmt.Printf("Hello! My name is %s!\n", input.Arg_Greeter)
	}
	return nil
}
