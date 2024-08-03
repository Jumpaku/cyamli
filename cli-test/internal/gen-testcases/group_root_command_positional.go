package main

var groupRootCommandPositional = group{
	Group:     "root_command_positional",
	Languages: languages,
	Testcases: []testcase{
		{
			Message: `root command should handle positional arguments with zero variadic arguments`,
			Want:    `_123_true_abc_`,
			Args:    []string{`123`, `true`, `abc`},
		},
		{
			Message: `root command should handle positional arguments with multiple variadic arguments`,
			Want:    `_123_true_abc_x,y,z,w`,
			Args:    []string{`123`, `true`, `abc`, `x`, `y`, `z`, `w`},
		},
	},
}
