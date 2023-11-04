package main

import (
	"os"

	"github.com/Jumpaku/cliautor/description"
	"github.com/Jumpaku/cliautor/name"
	"github.com/Jumpaku/cliautor/schema"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = s.Walk(func(path name.Path, cmd *schema.Command) error {
		return description.DescribeCommand(
			description.DetailExecutor(),
			description.CreateCommandData("main", path, cmd),
			os.Stdout)
	})
	if err != nil {
		panic(err)
	}
}
