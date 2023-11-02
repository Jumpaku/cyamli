package data

type Program struct {
	Name        string
	Version     string
	Options     []Option
	Arguments   []Argument
	Subcommands []Subcommand
}
