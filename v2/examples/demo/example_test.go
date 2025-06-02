package demo_test

import "os"

//go:generate cyamli generate golang -schema-path=cli.cyamli.yaml -out-path=cli.gen_test.go -package=demo_test
func Example_demo() {
	if err := Run(cli{}, os.Args); err != nil {
		panic(err)
	}
	// Output:
}

type cli struct{}

func (c cli) Run(input Input) error {
	return nil
}

func (c cli) Run_Data(input Input_Data) error {
	return nil
}

func (c cli) Run_Schema(input Input_Schema) error {
	return nil
}

func (c cli) Run_Tables(input Input_Tables) error {
	return nil
}
