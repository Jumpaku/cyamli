package schema_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/Jumpaku/cyamli/schema"
	"github.com/Jumpaku/cyamli/test"
	"github.com/Jumpaku/cyamli/test/testdata"

	"github.com/stretchr/testify/assert"
)

func TestSchema_Load(t *testing.T) {
	testcases := []struct {
		name      string
		in        string
		want      *schema.Schema
		shouldErr bool
	}{
		{
			name: "empty",
			in:   testdata.EmptyYAML,
			want: &schema.Schema{
				Program: schema.Program{},
			},
		},
		{
			name: "example",
			in:   testdata.ExampleYAML,
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
								"subx": {},
								"suby": {},
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
				test.AssertMatchSchema(t, testcase.want, got)
			}
		})
	}
}

func TestSchema_Save(t *testing.T) {
	testcases := []struct {
		name      string
		original  *schema.Schema
		shouldErr bool
	}{
		{

			name: "empty",
			original: &schema.Schema{
				Program: schema.Program{},
			},
		},
		{
			name: "example",
			original: &schema.Schema{
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
			buffer := bytes.NewBuffer(nil)
			err := testcase.original.Save(buffer)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				loaded, err := schema.Load(buffer)
				if err != nil {
					t.Fatalf("fail to load schema: %+v", err)
				}
				test.AssertMatchSchema(t, testcase.original, loaded)
			}
		})
	}
}
