package main

var groupSubcommandPositional = group{
	Group:     "subcommand_positional",
	Languages: languages,
	Testcases: []testcase{
		{
			Message: `subcommand should handle positional arguments with zero variadic arguments`,
			Want:    `sub_123_true_abc_`,
			Args:    []string{`sub`, `123`, `true`, `abc`},
		},
		{
			Message: `subcommand should handle positional arguments with multiple variadic arguments`,
			Want:    `sub_123_true_abc_x,y,z,w`,
			Args:    []string{`sub`, `123`, `true`, `abc`, `x`, `y`, `z`, `w`},
		},
	},
}
