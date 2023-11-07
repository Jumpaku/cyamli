package description_test

import (
	"bytes"
	"testing"

	"github.com/Jumpaku/cyamli/description"
	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"
	"github.com/Jumpaku/cyamli/test/testdata"

	"github.com/stretchr/testify/assert"
)

func TestDescribeCommand_Simple_Empty(t *testing.T) {
	executor := description.SimpleExecutor()
	emptySchema, err := schema.Load(bytes.NewBufferString(testdata.EmptyYAML))
	if err != nil {
		t.Fatalf("fail to read load schema: %+v", err)
	}

	testcases := []struct {
		path name.Path
		cmd  *schema.Command
		want string
	}{
		{
			path: name.Path{},
			cmd:  emptySchema.Program.Command(),
			want: `main:

Usage:
    $ main
`,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.path.Join(" ", "", ""), func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			data := description.CreateCommandData("main", testcase.path, testcase.cmd)

			err := description.DescribeCommand(executor, data, buffer)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testcase.want, buffer.String())
		})
	}
}

func TestDescribeCommand_Detail_Empty(t *testing.T) {
	executor := description.DetailExecutor()
	emptySchema, err := schema.Load(bytes.NewBufferString(testdata.EmptyYAML))
	if err != nil {
		t.Fatalf("fail to read load schema: %+v", err)
	}

	testcases := []struct {
		path name.Path
		cmd  *schema.Command
		want string
	}{
		{
			path: name.Path{},
			cmd:  emptySchema.Program.Command(),
			want: `main:

Usage:
    $ main
`,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.path.Join(" ", "", ""), func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			data := description.CreateCommandData("main", testcase.path, testcase.cmd)

			err := description.DescribeCommand(executor, data, buffer)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testcase.want, buffer.String())
		})
	}
}

func TestDescribeCommand_Simple_Example(t *testing.T) {
	executor := description.SimpleExecutor()
	exampleSchema, err := schema.Load(bytes.NewBufferString(testdata.ExampleYAML))
	if err != nil {
		t.Fatalf("fail to read load schema: %+v", err)
	}

	testcases := []struct {
		path name.Path
		cmd  *schema.Command
		want string
	}{
		{
			path: name.Path{},
			cmd:  exampleSchema.Program.Command(),
			want: `main:
this is an example command

Usage:
    $ main [<option>|<argument>]... [-- [<argument>]...]

Options:
    -option-a, -option-b, -option-c, -option-d, -option-e

Arguments:
    <arg_a> <arg_b> <arg_c> <arg_d> <arg_e> <arg_v>...

Subcommands:
    sub1, sub2, sub3
`,
		},
		{
			path: name.Path{"sub1"},
			cmd:  exampleSchema.Program.Subcommands["sub1"],
			want: `main sub1:
1 - this is a sub command

Usage:
    $ main sub1
`,
		},
		{
			path: name.Path{"sub2"},
			cmd:  exampleSchema.Program.Subcommands["sub2"],
			want: `main sub2:
2 - this is a sub command

Usage:
    $ main sub2
`,
		},
		{
			path: name.Path{"sub3"},
			cmd:  exampleSchema.Program.Subcommands["sub3"],
			want: `main sub3:
3 - this is a sub command

Usage:
    $ main sub3 [<option>|<argument>]... [-- [<argument>]...]

Options:
    -option-a, -option-b, -option-c, -option-d, -option-e

Arguments:
    <arg_a> <arg_b> <arg_c> <arg_d> <arg_e> <arg_v>...

Subcommands:
    subx, suby
`,
		},
		{
			path: name.Path{"sub3", "subx"},
			cmd:  exampleSchema.Program.Subcommands["sub3"].Subcommands["subx"],
			want: `main sub3 subx:

Usage:
    $ main sub3 subx
`,
		},
		{
			path: name.Path{"sub3", "suby"},
			cmd:  exampleSchema.Program.Subcommands["sub3"].Subcommands["suby"],
			want: `main sub3 suby:

Usage:
    $ main sub3 suby
`,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.path.Join(" ", "", ""), func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			data := description.CreateCommandData("main", testcase.path, testcase.cmd)

			err := description.DescribeCommand(executor, data, buffer)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testcase.want, buffer.String())
		})
	}
}

func TestDescribeCommand_Detail_Example(t *testing.T) {
	executor := description.DetailExecutor()
	exampleSchema, err := schema.Load(bytes.NewBufferString(testdata.ExampleYAML))
	if err != nil {
		t.Fatalf("fail to read load schema: %+v", err)
	}

	testcases := []struct {
		path name.Path
		cmd  *schema.Command
		want string
	}{
		{
			path: name.Path{},
			cmd:  exampleSchema.Program.Command(),
			want: `main:
this is an example command

Usage:
    $ main [<option>|<argument>]... [-- [<argument>]...]

Options:
    -option-a=<string>, -a=<string>  [default="abc"]:
        a - this is an option for root command

    -option-b=<integer>, -b=<integer>  [default=-123]:
        b - this is an option for root command

    -option-c[=<boolean>], -c[=<boolean>]  [default=true]:
        c - this is an option for root command

    -option-d=<float>, -d=<float>  [default=-123.456]:
        d - this is an option for root command

    -option-e=<string>  [default=""]:

Arguments:
    [0]  <arg_a:string>
        a - this is an argument for root command

    [1]  <arg_b:integer>
        b - this is an argument for root command

    [2]  <arg_c:boolean>
        c - this is an argument for root command

    [3]  <arg_d:float>
        d - this is an argument for root command

    [4]  <arg_e:string>

    [5:] [<arg_v:string>]...
        v - this is an argument for root command

Subcommands:
    sub1:
        1 - this is a sub command

    sub2:
        2 - this is a sub command

    sub3:
        3 - this is a sub command
`,
		},
		{
			path: name.Path{"sub1"},
			cmd:  exampleSchema.Program.Subcommands["sub1"],
			want: `main sub1:
1 - this is a sub command

Usage:
    $ main sub1
`,
		},
		{
			path: name.Path{"sub2"},
			cmd:  exampleSchema.Program.Subcommands["sub2"],
			want: `main sub2:
2 - this is a sub command

Usage:
    $ main sub2
`,
		},
		{
			path: name.Path{"sub3"},
			cmd:  exampleSchema.Program.Subcommands["sub3"],
			want: `main sub3:
3 - this is a sub command

Usage:
    $ main sub3 [<option>|<argument>]... [-- [<argument>]...]

Options:
    -option-a=<string>, -a=<string>  [default="abc"]:
        3 - a - this is an option for root command

    -option-b=<integer>, -b=<integer>  [default=-123]:
        3 - b - this is an option for root command

    -option-c[=<boolean>], -c[=<boolean>]  [default=true]:
        3 - c - this is an option for root command

    -option-d=<float>, -d=<float>  [default=-123.456]:
        3 - d - this is an option for root command

    -option-e=<string>  [default=""]:

Arguments:
    [0]  <arg_a:string>
        3 - a - this is an argument for root command

    [1]  <arg_b:integer>
        3 - b - this is an argument for root command

    [2]  <arg_c:boolean>
        3 - c - this is an argument for root command

    [3]  <arg_d:float>
        3 - d - this is an argument for root command

    [4]  <arg_e:string>

    [5:] [<arg_v:string>]...
        3 - v - this is an argument for root command

Subcommands:
    subx:

    suby:
`,
		},
		{
			path: name.Path{"sub3", "subx"},
			cmd:  exampleSchema.Program.Subcommands["sub3"].Subcommands["subx"],
			want: `main sub3 subx:

Usage:
    $ main sub3 subx
`,
		},
		{
			path: name.Path{"sub3", "suby"},
			cmd:  exampleSchema.Program.Subcommands["sub3"].Subcommands["suby"],
			want: `main sub3 suby:

Usage:
    $ main sub3 suby
`,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.path.Join(" ", "", ""), func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			data := description.CreateCommandData("main", testcase.path, testcase.cmd)

			err := description.DescribeCommand(executor, data, buffer)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testcase.want, buffer.String())
		})
	}
}
