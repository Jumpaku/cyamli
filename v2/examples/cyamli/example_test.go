package cyamli_test

import "os"

//go:generate cyamli generate golang -schema-path=cli.cyamli.yaml -out-path=cli.gen_test.go -package=cyamli_test
func Example_cyamli() {
	if err := Run(cli{}, os.Args); err != nil {
		panic(err)
	}
	// Output:
}

type cli struct{}

func (c cli) Run(input Input) error {
	return nil
}
func (c cli) Run_Version(input Input_Version) error {
	return nil
}
func (c cli) Run_Generate(input Input_Generate) error {
	return nil
}
func (c cli) Run_GenerateDart3(input Input_GenerateDart3) (err error) {
	return nil
}
func (c cli) Run_GenerateDocs(input Input_GenerateDocs) error {
	return nil
}
func (c cli) Run_GenerateGolang(input Input_GenerateGolang) error {
	return nil
}
func (c cli) Run_GenerateKotlin(input Input_GenerateKotlin) error {
	return nil
}
func (c cli) Run_GeneratePython3(input Input_GeneratePython3) error {
	return nil
}
func (c cli) Run_GenerateTypescript(input Input_GenerateTypescript) error {
	return nil
}
