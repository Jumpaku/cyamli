package main

var groupRootCommandOption = group{
	Group:     "root_command_options",
	Languages: languages,
	Testcases: []testcase{
		{
			Message: `root command should handle optional arguments of type integer`,
			Want:    `_123_false_`,
			Args:    []string{`-opt-integer=123`},
		},
		{
			Message: `root command should handle optional arguments of type boolean`,
			Want:    `_0_true_`,
			Args:    []string{`-opt-boolean=true`},
		},
		{
			Message: `root command should handle optional arguments of type boolean`,
			Want:    `_0_false_abc`,
			Args:    []string{`-opt-string=abc`},
		},
		{
			Message: `root command should handle optional arguments of type integer`,
			Want:    `_123_false_`,
			Args:    []string{` -i=123`},
		},
		{
			Message: `root command should handle short optional arguments of type boolean`,
			Want:    `_0_true_`,
			Args:    []string{`-b=true`},
		},
		{
			Message: `root command should handle short optional arguments of type boolean`,
			Want:    `_0_false_abc`,
			Args:    []string{`-s=abc`},
		},
		{
			Message: `root command should handle default optional argument values`,
			Want:    `_0_false_`,
			Args:    []string{},
		},
		{
			Message: `root command should handle optional arguments of type boolean without value`,
			Want:    `_0_true_`,
			Args:    []string{`-opt-boolean`},
		},
		{
			Message: `root command should handle short optional arguments of type boolean without value`,
			Want:    `_0_true_`,
			Args:    []string{`-b`},
		},
	},
}
