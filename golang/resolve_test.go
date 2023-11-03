package golang_test

import (
	"cliautor/golang"
	"cliautor/schema"
	"cliautor/test"
	_ "embed"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGoValue(t *testing.T) {
	testcases := []struct {
		in_type     schema.Type
		in_variadic bool
		in_str      []string
		shouldErr   bool
		want        any
	}{
		// Boolean
		{
			in_type:     schema.TypeBoolean,
			in_variadic: true,
			in_str:      []string{},
			want:        []bool{},
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: true,
			in_str:      []string{"1", "0", "t", "f", "true", "false"},
			want:        []bool{true, false, true, false, true, false},
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: false,
			in_str:      []string{"true"},
			want:        true,
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: false,
			in_str:      []string{"1", "0", "t", "f", "true", "false"},
			shouldErr:   true,
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: false,
			in_str:      []string{"not_bool"},
			shouldErr:   true,
		},

		// Integer
		{
			in_type:     schema.TypeInteger,
			in_variadic: true,
			in_str:      []string{},
			want:        []int64{},
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: true,
			in_str:      []string{"-123", "+456", "0b0101", "-0b1010", "017", "-0o17", "0xff", "-0xff"},
			want:        []int64{-123, 456, 5, -10, 15, -15, 255, -255},
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: false,
			in_str:      []string{"0"},
			want:        int64(0),
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: false,
			in_str:      []string{"-123", "+456", "0b0101", "-0b1010", "017", "-0o17", "0xff", "-0xff"},
			shouldErr:   true,
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: false,
			in_str:      []string{"not_int"},
			shouldErr:   true,
		},

		// Float
		{
			in_type:     schema.TypeFloat,
			in_variadic: true,
			in_str:      []string{},
			want:        []float64{},
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: true,
			in_str:      []string{"-123.456", "+456.789", "inf", "+inf", "Infinity", "+Infinity", "-inf", "-Infinity"},
			want:        []float64{-123.456, 456.789, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(-1), math.Inf(-1)},
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: false,
			in_str:      []string{"0"},
			want:        float64(0),
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: false,
			in_str:      []string{"-123.456", "+456.789", "inf", "+inf", "Infinity", "+Infinity", "-inf", "-Infinity"},
			shouldErr:   true,
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: false,
			in_str:      []string{"not_float"},
			shouldErr:   true,
		},

		// String
		{
			in_type:     schema.TypeString,
			in_variadic: true,
			in_str:      []string{},
			want:        []string{},
		},
		{
			in_type:     schema.TypeString,
			in_variadic: true,
			in_str:      []string{"abc", "def", "ghi"},
			want:        []string{"abc", "def", "ghi"},
		},
		{
			in_type:     schema.TypeString,
			in_variadic: false,
			in_str:      []string{"abcdefghi"},
			want:        "abcdefghi",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: false,
			in_str:      []string{"abc", "def", "ghi"},
			shouldErr:   true,
		},

		// Unspecified
		{
			in_variadic: true,
			in_str:      []string{},
			want:        []string{},
		},
		{
			in_variadic: true,
			in_str:      []string{"abc", "def", "ghi"},
			want:        []string{"abc", "def", "ghi"},
		},
		{
			in_variadic: false,
			in_str:      []string{"abcdefghi"},
			want:        "abcdefghi",
		},
		{
			in_variadic: false,
			in_str:      []string{"abc", "def", "ghi"},
			shouldErr:   true,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%#v", testcase.in_str), func(t *testing.T) {
			got, err := golang.ParseGoValue(testcase.in_type, testcase.in_variadic, testcase.in_str...)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, testcase.want, got)
			}
		})
	}
}

func newExampleSchema() *schema.Schema {
	return &schema.Schema{
		Program: schema.Program{
			Name:        "example",
			Version:     "v1.0.0",
			Description: "this is an example command",
			Options: map[string]*schema.Option{
				"-option-a": {
					Short:       "-a",
					Description: "a - this is an option for root command",
					Type:        schema.TypeString,
					Default:     "abc",
				},
				"-option-b": {
					Short:       "-b",
					Description: "b - this is an option for root command",
					Type:        schema.TypeInteger,
					Default:     "-123",
				},
				"-option-c": {
					Short:       "-c",
					Description: "c - this is an option for root command",
					Type:        schema.TypeBoolean,
					Default:     "true",
				},
				"-option-d": {
					Short:       "-d",
					Description: "d - this is an option for root command",
					Type:        schema.TypeFloat,
					Default:     "-123.456",
				},
				"-option-e": {},
			},
			Arguments: []*schema.Argument{
				{
					Name:        "arg_a",
					Description: "a - this is an argument for root command",
					Type:        schema.TypeString,
				},
				{
					Name:        "arg_b",
					Description: "b - this is an argument for root command",
					Type:        schema.TypeInteger,
				},
				{
					Name:        "arg_c",
					Description: "c - this is an argument for root command",
					Type:        schema.TypeBoolean,
				},
				{
					Name:        "arg_d",
					Description: "d - this is an argument for root command",
					Type:        schema.TypeFloat,
				},
				{
					Name: "arg_e",
				},
				{
					Name:        "arg_v",
					Description: "v - this is an argument for root command",
					Variadic:    true,
				},
			},
			Subcommands: map[string]*schema.Command{
				"sub1": {
					Description: "1 - this is a sub command",
					Options:     map[string]*schema.Option{},
					Arguments:   []*schema.Argument{},
					Subcommands: map[string]*schema.Command{},
				},
				"sub2": {
					Description: "2 - this is a sub command",
					Options:     map[string]*schema.Option{},
					Arguments:   []*schema.Argument{},
					Subcommands: map[string]*schema.Command{},
				},
				"sub3": {
					Description: "3 - this is a sub command",
					Options: map[string]*schema.Option{
						"-option-a": {
							Short:       "-a",
							Description: "3 - a - this is an option for root command",
							Type:        schema.TypeString,
							Default:     "abc",
						},
						"-option-b": {
							Short:       "-b",
							Description: "3 - b - this is an option for root command",
							Type:        schema.TypeInteger,
							Default:     "-123",
						},
						"-option-c": {
							Short:       "-c",
							Description: "3 - c - this is an option for root command",
							Type:        schema.TypeBoolean,
							Default:     "true",
						},
						"-option-d": {
							Short:       "-d",
							Description: "3 - d - this is an option for root command",
							Type:        schema.TypeFloat,
							Default:     "-123.456",
						},
						"-option-e": {},
					},
					Arguments: []*schema.Argument{
						{
							Name:        "arg_a",
							Description: "3 - a - this is an argument for root command",
							Type:        schema.TypeString,
						},
						{
							Name:        "arg_b",
							Description: "3 - b - this is an argument for root command",
							Type:        schema.TypeInteger,
						},
						{
							Name:        "arg_c",
							Description: "3 - c - this is an argument for root command",
							Type:        schema.TypeBoolean,
						},
						{
							Name:        "arg_d",
							Description: "3 - d - this is an argument for root command",
							Type:        schema.TypeFloat,
						},
						{
							Name: "arg_e",
						},
						{
							Name:        "arg_v",
							Description: "3 - v - this is an argument for root command",
							Variadic:    true,
						},
					},
					Subcommands: map[string]*schema.Command{
						"subx": {},
						"suby": {},
					},
				},
			},
		},
	}
}
func TestResolveSubcommand(t *testing.T) {
	program := newExampleSchema().Program
	testcases := []struct {
		args      []string
		want_cmd  *schema.Command
		want_str  string
		want_args []string
	}{
		{
			args:      []string{"program"},
			want_cmd:  program.Command(),
			want_str:  "",
			want_args: nil,
		},
		{
			args:      []string{"program", "-opt-x", "arg_x"},
			want_cmd:  program.Command(),
			want_str:  "",
			want_args: []string{"-opt-x", "arg_x"},
		},
		{
			args:      []string{"program", "sub1"},
			want_cmd:  program.Subcommands["sub1"],
			want_str:  "sub1",
			want_args: nil,
		},
		{
			args:      []string{"program", "sub1", "-opt-x", "arg_x"},
			want_cmd:  program.Subcommands["sub1"],
			want_str:  "sub1",
			want_args: []string{"-opt-x", "arg_x"},
		},
		{
			args:      []string{"program", "--", "sub1", "-opt-x", "arg_x"},
			want_cmd:  program.Command(),
			want_str:  "",
			want_args: []string{"--", "sub1", "-opt-x", "arg_x"},
		},
		{
			args:      []string{"program", "sub2"},
			want_cmd:  program.Subcommands["sub2"],
			want_str:  "sub2",
			want_args: nil,
		},
		{
			args:      []string{"program", "sub2", "-opt-x", "arg_x"},
			want_cmd:  program.Subcommands["sub2"],
			want_str:  "sub2",
			want_args: []string{"-opt-x", "arg_x"},
		},
		{
			args:      []string{"program", "sub3"},
			want_cmd:  program.Subcommands["sub3"],
			want_str:  "sub3",
			want_args: nil,
		},
		{
			args:      []string{"program", "sub3", "-opt-x", "arg_x"},
			want_cmd:  program.Subcommands["sub3"],
			want_str:  "sub3",
			want_args: []string{"-opt-x", "arg_x"},
		},
		{
			args:      []string{"program", "sub3", "subx"},
			want_cmd:  program.Subcommands["sub3"].Subcommands["subx"],
			want_str:  "sub3 subx",
			want_args: nil,
		},
		{
			args:      []string{"program", "sub3", "subx", "-opt-x", "arg_x"},
			want_cmd:  program.Subcommands["sub3"].Subcommands["subx"],
			want_str:  "sub3 subx",
			want_args: []string{"-opt-x", "arg_x"},
		},
		{
			args:      []string{"program", "sub3", "suby"},
			want_cmd:  program.Subcommands["sub3"].Subcommands["suby"],
			want_str:  "sub3 suby",
			want_args: nil,
		},
		{
			args:      []string{"program", "sub3", "--", "subx"},
			want_cmd:  program.Subcommands["sub3"],
			want_str:  "sub3",
			want_args: []string{"--", "subx"},
		},
		{
			args:      []string{"program", "sub3", "--", "suby"},
			want_cmd:  program.Subcommands["sub3"],
			want_str:  "sub3",
			want_args: []string{"--", "suby"},
		},
		{
			args:      []string{"program", "--", "sub3", "subx"},
			want_cmd:  program.Command(),
			want_str:  "",
			want_args: []string{"--", "sub3", "subx"},
		},
		{
			args:      []string{"program", "--", "sub3", "suby"},
			want_cmd:  program.Command(),
			want_str:  "",
			want_args: []string{"--", "sub3", "suby"},
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%#v", testcase.args), func(t *testing.T) {
			got_cmd, got_str, got_args := golang.ResolveSubcommand(&schema.Schema{Program: program}, testcase.args)

			test.AssertMatchCommand(t, testcase.want_cmd, got_cmd)
			assert.Equal(t, testcase.want_str, got_str)
			assert.ElementsMatch(t, testcase.want_args, got_args)
		})
	}
}

func TestResolveOptions(t *testing.T) {
	type Input struct {
		Opt_OptionA string
		Opt_OptionB int64
		Opt_OptionC bool
		Opt_OptionD float64
		Opt_OptionE string

		Arg_ArgA string
		Arg_ArgB int64
		Arg_ArgC bool
		Arg_ArgD float64
		Arg_ArgE string
		Arg_ArgV []string
	}

	program := newExampleSchema().Program

	testcases := []struct {
		cmd       *schema.Command
		restArgs  []string
		shouldErr bool
		want      Input
	}{
		{
			restArgs: []string{"abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "abc",
				Opt_OptionB: -123,
				Opt_OptionC: true,
				Opt_OptionD: -123.456,
				Opt_OptionE: "",
			},
		},
		{
			restArgs: []string{"-option-a=def", "-option-b=456", "-option-c=false", "-option-d=456.789", "-option-e=str", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "def",
				Opt_OptionB: 456,
				Opt_OptionC: false,
				Opt_OptionD: 456.789,
				Opt_OptionE: "str",
			},
		},
		{
			restArgs: []string{"abc", "123", "true", "123.456", "xyz", "-option-a=def", "-option-b=456", "-option-c=false", "-option-d=456.789", "-option-e=str"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "def",
				Opt_OptionB: 456,
				Opt_OptionC: false,
				Opt_OptionD: 456.789,
				Opt_OptionE: "str",
			},
		},
		{
			restArgs: []string{"-a=def", "-b=456", "-c=false", "-d=456.789", "-option-e=str", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "def",
				Opt_OptionB: 456,
				Opt_OptionC: false,
				Opt_OptionD: 456.789,
				Opt_OptionE: "str",
			},
		},
		{
			restArgs: []string{"-option-c", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "abc",
				Opt_OptionB: -123,
				Opt_OptionC: true,
				Opt_OptionD: -123.456,
				Opt_OptionE: "",
			},
		},
		{
			restArgs:  []string{"-option-a", "abc", "123", "true", "123.456", "xyz"},
			cmd:       program.Command(),
			shouldErr: true,
		},
		{
			restArgs:  []string{"-option-x", "abc", "123", "true", "123.456", "xyz"},
			cmd:       program.Command(),
			shouldErr: true,
		},
		{
			restArgs:  []string{"-option-b=not_int", "abc", "123", "true", "123.456", "xyz"},
			cmd:       program.Command(),
			shouldErr: true,
		},
		{
			restArgs: []string{"--", "-option-a=def", "-option-b=456", "-option-c=false", "-option-d=456.789", "-option-e=str", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "abc",
				Opt_OptionB: -123,
				Opt_OptionC: true,
				Opt_OptionD: -123.456,
				Opt_OptionE: "",
			},
		},
		{
			restArgs: []string{"--", "-option-x", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Opt_OptionA: "abc",
				Opt_OptionB: -123,
				Opt_OptionC: true,
				Opt_OptionD: -123.456,
				Opt_OptionE: "",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%#v", testcase.restArgs), func(t *testing.T) {
			input := Input{
				Opt_OptionA: "abc",
				Opt_OptionB: -123,
				Opt_OptionC: true,
				Opt_OptionD: -123.456,
			}
			err := golang.ResolveOptions(testcase.cmd.Options, testcase.restArgs, &input)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, testcase.want, input)
			}
		})
	}
}

func TestResolveArguments(t *testing.T) {
	type Input struct {
		Opt_OptionA string
		Opt_OptionB int64
		Opt_OptionC bool
		Opt_OptionD float64
		Opt_OptionE string

		Arg_ArgA string
		Arg_ArgB int64
		Arg_ArgC bool
		Arg_ArgD float64
		Arg_ArgE string
		Arg_ArgV []string
	}

	program := newExampleSchema().Program

	testcases := []struct {
		cmd       *schema.Command
		restArgs  []string
		shouldErr bool
		want      Input
	}{
		{
			restArgs: []string{"abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{},
			},
		},
		{
			restArgs: []string{"-option-a=def", "-option-b=456", "-option-c=false", "-option-d=456.789", "-option-e=str", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{},
			},
		},
		{
			restArgs: []string{"abc", "123", "true", "123.456", "xyz", "-option-a=def", "-option-b=456", "-option-c=false", "-option-d=456.789", "-option-e=str"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{},
			},
		},
		{
			restArgs: []string{"-a=def", "-b=456", "-c=false", "-d=456.789", "-option-e=str", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{},
			},
		},
		{
			restArgs: []string{"-option-c", "abc", "123", "true", "123.456", "xyz"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{},
			},
		},
		{
			restArgs: []string{"-a=def", "-b=456", "-c=false", "-d=456.789", "-option-e=str", "abc", "123", "true", "123.456", "xyz", "v0", "v1", "v2", "v3"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{"v0", "v1", "v2", "v3"},
			},
		},
		{
			restArgs: []string{"abc", "123", "true", "123.456", "xyz", "v0", "v1", "-a=def", "-b=456", "-c=false", "-d=456.789", "-option-e=str", "v2", "v3"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{"v0", "v1", "v2", "v3"},
			},
		},
		{
			restArgs: []string{"--", "abc", "123", "true", "123.456", "xyz", "v0", "v1", "-a=def", "-b=456", "-c=false", "-d=456.789", "-option-e=str", "v2", "v3"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{"v0", "v1", "-a=def", "-b=456", "-c=false", "-d=456.789", "-option-e=str", "v2", "v3"},
			},
		},
		{
			restArgs: []string{"-a=def", "-b=456", "abc", "-c=false", "-d=456.789", "--", "-123", "true", "-123.456", "xyz", "v0", "v1", "-option-e=str", "v2", "v3"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: -123,
				Arg_ArgC: true,
				Arg_ArgD: -123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{"v0", "v1", "-option-e=str", "v2", "v3"},
			},
		},
		{
			restArgs: []string{"-a=def", "-b=456", "abc", "-c=false", "-d=456.789", "123", "true", "123.456", "xyz", "v0", "v1", "-option-e=str", "v2", "v3", "--"},
			cmd:      program.Command(),
			want: Input{
				Arg_ArgA: "abc",
				Arg_ArgB: 123,
				Arg_ArgC: true,
				Arg_ArgD: 123.456,
				Arg_ArgE: "xyz",
				Arg_ArgV: []string{"v0", "v1", "v2", "v3"},
			},
		},
		{
			restArgs:  []string{"-a=def", "-b=456", "abc", "-c=false", "-d=456.789", "not_int", "not_bool", "not_float", "xyz", "v0", "v1", "-option-e=str", "v2", "v3"},
			cmd:       program.Command(),
			shouldErr: true,
		},
		{
			restArgs:  []string{"abc", "123", "true", "123.456"},
			cmd:       program.Command(),
			shouldErr: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%#v", testcase.restArgs), func(t *testing.T) {
			input := Input{}
			err := golang.ResolveArguments(testcase.cmd.Arguments, testcase.restArgs, &input)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, testcase.want, input)
			}
		})
	}
}
