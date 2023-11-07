package main

import (
	"fmt"
	"os"

	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	// Create the CLI object
	cli := NewCLI()
	// Overwrite behaviors
	cli.Func = showHelp
	cli.Sub_Hello.Func = sayHello
	// Run with command line arguments
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func showHelp(cmd *schema.Command, subcommand []string, input CLI_Input) (err error) {
	if input.Opt_Help {
		fmt.Println("This is an example program.")
	} else {
		fmt.Println("Do nothing.")
	}
	return nil
}
func sayHello(cmd *schema.Command, subcommand []string, input CLI_Hello_Input) (err error) {
	hello := "Hello"
	if input.Opt_TargetName != "" {
		hello += ", " + input.Opt_TargetName
	}
	hello += "! My name is " + input.Arg_Greeter + "!"
	fmt.Println(hello)
	return nil
}
