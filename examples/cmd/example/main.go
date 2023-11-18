package main

import (
	"os"

	"github.com/Jumpaku/cyamli/golang"
)

//go:generate go run "github.com/Jumpaku/cyamli/cmd/cyamli@latest" golang -schema-path=cli.yaml -out-path=cli.gen.go
func main() {
	schema := LoadSchema()
	cli := NewCLI()
	cli.FUNC = golang.HelpFunc[CLI_Input](schema)
	cli.Sub1.FUNC = golang.HelpFunc[CLI_Sub1_Input](schema)
	cli.Sub2.FUNC = golang.HelpFunc[CLI_Sub2_Input](schema)
	cli.Sub3.FUNC = golang.HelpFunc[CLI_Sub3_Input](schema)
	cli.Sub3.Subx.FUNC = golang.HelpFunc[CLI_Sub3Subx_Input](schema)
	cli.Sub3.Suby.FUNC = golang.HelpFunc[CLI_Sub3Suby_Input](schema)
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
