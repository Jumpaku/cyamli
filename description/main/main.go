package main

import (
	"os"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = s.Walk(func(path name.Path, cmd *schema.Command) error {
		return description.DescribeCommand(
			description.DetailExecutor(),
			description.CreateCommandData("main", "", path, cmd),
			os.Stdout)
	})
	if err != nil {
		panic(err)
	}
}
