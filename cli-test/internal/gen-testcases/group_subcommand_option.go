package main

var groupSubcommandOption = group{
	Group:     "subcommand_options",
	Languages: languages,
	Testcases: []testcase{
		{
			Message: `subcommand should handle optional arguments of type integer`,
			Want:    `sub_123_false_`,
			Args:    []string{`sub -opt-integer=123`},
		},
		{
			Message: `subcommand should handle optional arguments of type boolean`,
			Want:    `sub_0_true_`,
			Args:    []string{`sub -opt-boolean=true`},
		},
		{
			Message: `subcommand should handle optional arguments of type boolean`,
			Want:    `sub_0_false_abc`,
			Args:    []string{`sub -opt-string=abc`},
		},
		{
			Message: `subcommand should handle optional arguments of type integer`,
			Want:    `sub_123_false_`,
			Args:    []string{`sub -i=123`},
		},
		{
			Message: `subcommand should handle short optional arguments of type boolean`,
			Want:    `sub_0_true_`,
			Args:    []string{`sub -b=true`},
		},
		{
			Message: `subcommand should handle short optional arguments of type boolean`,
			Want:    `sub_0_false_abc`,
			Args:    []string{`sub -s=abc`},
		},
		{
			Message: `subcommand should handle default optional argument values`,
			Want:    `sub_0_false_`,
			Args:    []string{`sub`},
		},
		{
			Message: `subcommand should handle optional arguments of type boolean without value`,
			Want:    `sub_0_true_`,
			Args:    []string{`sub -opt-boolean`},
		},
		{
			Message: `subcommand should handle short optional arguments of type boolean without value`,
			Want:    `sub_0_true_`,
			Args:    []string{`sub -b`},
		},
	},
}
