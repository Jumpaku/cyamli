package schema_test

import (
	"bytes"
	"cliautor/schema"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/empty.yaml
var emptyYAML string

//go:embed testdata/example.yaml
var exampleYAML string

func TestSchema_Load(t *testing.T) {
	testcases := []struct {
		name      string
		in        string
		want      *schema.Schema
		shouldErr bool
	}{
		{

			name: "empty",
			in:   emptyYAML,
			want: &schema.Schema{
				Program: schema.Program{},
			},
		},
		{
			name: "example",
			in:   exampleYAML,
			want: &schema.Schema{
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
								"sub1": {},
								"sub2": {},
							},
						},
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got, err := schema.Load(bytes.NewBufferString(testcase.in))
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				assertMatchSchema(t, testcase.want, got)
			}
		})
	}
}

func assertMatchSchema(t *testing.T, want, got *schema.Schema) {
	wp, gp := want.Program, got.Program
	assert.Equal(t, wp.Name, gp.Name)
	assert.Equal(t, wp.Version, gp.Version)
	assertMatchCommand(t, wp.Command(), gp.Command())
}

func assertMatchCommand(t *testing.T, want, got *schema.Command) {
	assert.Equal(t, want.Description, got.Description)

	wOpt, gOpt := want.Options, got.Options
	for name, w := range wOpt {
		g, ok := gOpt[name]
		assert.Conditionf(t, func() (success bool) { return ok }, "option %q not found", name)
		assert.Equal(t, w.Short, g.Short)
		assert.Equal(t, w.Default, g.Default)
		assert.Equal(t, w.Description, g.Description)
		assert.Equal(t, w.Type, g.Type)
	}

	wArg, gArg := want.Arguments, got.Arguments
	for idx, w := range wArg {
		assert.Conditionf(t, func() (success bool) {
			return idx < len(gArg)
		}, "argument %q not found at %d", w.Name, idx)
		g := gArg[idx]
		assert.Equal(t, w.Name, g.Name)
		assert.Equal(t, w.Variadic, g.Variadic)
		assert.Equal(t, w.Description, g.Description)
		assert.Equal(t, w.Type, g.Type)
	}

	wSub, gSub := want.Subcommands, got.Subcommands
	for name, w := range wSub {
		g, ok := gSub[name]
		assert.Conditionf(t, func() (success bool) {
			return ok
		}, "subcommand %q not found", name)
		assertMatchCommand(t, w, g)
	}
}
