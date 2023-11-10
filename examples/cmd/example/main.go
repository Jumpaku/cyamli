package main

import (
	"os"

	"github.com/Jumpaku/cyamli/golang"
)

func main() {
	schema := LoadSchema()
	cli := NewCLI()
	cli.Sub_Sub1.Func = golang.HelpFunc[CLI_Sub1_Input](schema)
	cli.Sub_Sub2.Func = golang.HelpFunc[CLI_Sub2_Input](schema)
	cli.Sub_Sub3.Func = golang.HelpFunc[CLI_Sub3_Input](schema)
	cli.Sub_Sub3.Sub_Subx.Func = golang.HelpFunc[CLI_Sub3Subx_Input](schema)
	cli.Sub_Sub3.Sub_Suby.Func = golang.HelpFunc[CLI_Sub3Suby_Input](schema)
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
