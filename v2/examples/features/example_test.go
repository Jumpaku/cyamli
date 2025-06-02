package features_test

import "os"

//go:generate cyamli generate golang -schema-path=cli.cyamli.yaml -out-path=cli.gen_test.go -package=features_test
func Example_features() {
	if err := Run(cli{}, os.Args); err != nil {
		panic(err)
	}
	// Output:
}

type cli struct{}

func (c cli) Run(input Input) error {
	return nil
}

func (c cli) Run_Sub1(input Input_Sub1) error {
	return nil
}

func (c cli) Run_Sub1Sub2(input Input_Sub1Sub2) error {
	return nil
}

func (c cli) Run_Sub1Sub2Sub3(input Input_Sub1Sub2Sub3) error {
	return nil
}
