package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	cli := NewCLI()
	cli.FUNC = HelpFunc[CLI_Input]
	cli.Sub1.FUNC = HelpFunc[CLI_Sub1_Input]
	cli.Sub2.FUNC = HelpFunc[CLI_Sub2_Input]
	cli.Sub3.FUNC = HelpFunc[CLI_Sub3_Input]
	cli.Sub3.Suba.FUNC = HelpFunc[CLI_Sub3Suba_Input]
	cli.Sub3.Subb.FUNC = HelpFunc[CLI_Sub3Subb_Input]
	cli.Sub3.Subc.FUNC = HelpFunc[CLI_Sub3Subc_Input]
	cli.Sub3.Subd.FUNC = HelpFunc[CLI_Sub3Subd_Input]
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func HelpFunc[Input any](subcommand []string, input Input, inputErr error) error {
	spew.Dump(inputErr)
	spew.Dump(input)
	fmt.Println(GetDoc(subcommand))
	return nil
}
