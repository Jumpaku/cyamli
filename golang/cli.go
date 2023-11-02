package golang

import "cliautor/name"

type CLIBase struct {
}

func (cli CLIBase) ResolveSubcommand(args []string) name.Path {
	return nil
}

func (cli CLIBase) ResolveOptions(args []string) map[string]any {
	return nil
}

func (cli CLIBase) ResolveArguments(args []string) map[string]any {
	return nil
}

func Run[Input any](subcommand []string, options map[string]any, arguments map[string]any) {

}
