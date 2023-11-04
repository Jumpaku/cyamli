package data_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/Jumpaku/cliautor"
	"github.com/Jumpaku/cliautor/golang/data"
	"github.com/Jumpaku/cliautor/name"
	"github.com/Jumpaku/cliautor/schema"
	"github.com/Jumpaku/cliautor/test"
	"github.com/Jumpaku/cliautor/test/testdata"

	"github.com/stretchr/testify/assert"
)

func TestData_Construct(t *testing.T) {
	testcases := []struct {
		packageName string
		schemaYAML  string
		want        data.Data
	}{
		{
			packageName: "empty",
			schemaYAML:  testdata.EmptyYAML,
			want: data.Data{
				Package:          "empty",
				Generator:        cliautor.Name,
				GeneratorVersion: cliautor.Version,
				SchemaYAML:       testdata.EmptyYAML,
				Program:          data.Program{},
				Commands:         []data.Command{},
			},
		},
		{
			packageName: "example",
			schemaYAML:  testdata.ExampleYAML,
			want: data.Data{
				Package:          "example",
				Generator:        cliautor.Name,
				GeneratorVersion: cliautor.Version,
				SchemaYAML:       testdata.ExampleYAML,
				Program: data.Program{
					Name:    name.Path{"example"},
					Version: "v1.0.0",
					Options: []data.Option{
						{
							Name:    name.Path{"option", "a"},
							Short:   name.Path{"a"},
							Type:    schema.TypeString,
							Default: "abc",
						},
						{
							Name:    name.Path{"option", "b"},
							Short:   name.Path{"b"},
							Type:    schema.TypeInteger,
							Default: "-123",
						},
						{
							Name:    name.Path{"option", "c"},
							Short:   name.Path{"c"},
							Type:    schema.TypeBoolean,
							Default: "true",
						},
						{
							Name:    name.Path{"option", "d"},
							Short:   name.Path{"d"},
							Type:    schema.TypeFloat,
							Default: "-123.456",
						},
						{
							Name: name.Path{"option", "e"},
						},
					},
					Arguments: []data.Argument{
						{
							Name: name.Path{"arg", "a"},
							Type: schema.TypeString,
						},
						{
							Name: name.Path{"arg", "b"},
							Type: schema.TypeInteger,
						},
						{
							Name: name.Path{"arg", "c"},
							Type: schema.TypeBoolean,
						},
						{
							Name: name.Path{"arg", "d"},
							Type: schema.TypeFloat,
						},
						{
							Name: name.Path{"arg", "e"},
						},
						{
							Name:     name.Path{"arg", "v"},
							Variadic: true,
						},
					},
					Subcommands: []data.Subcommand{
						{Name: name.Path{"sub1"}},
						{Name: name.Path{"sub2"}},
						{Name: name.Path{"sub3"}},
					},
				},
				Commands: []data.Command{
					{
						Name: name.Path{"sub1"},
					},
					{
						Name: name.Path{"sub2"},
					},
					{
						Name: name.Path{"sub3"},
						Options: []data.Option{
							{
								Name:    name.Path{"option", "a"},
								Short:   name.Path{"a"},
								Type:    schema.TypeString,
								Default: "abc",
							},
							{
								Name:    name.Path{"option", "b"},
								Short:   name.Path{"b"},
								Type:    schema.TypeInteger,
								Default: "-123",
							},
							{
								Name:    name.Path{"option", "c"},
								Short:   name.Path{"c"},
								Type:    schema.TypeBoolean,
								Default: "true",
							},
							{
								Name:    name.Path{"option", "d"},
								Short:   name.Path{"d"},
								Type:    schema.TypeFloat,
								Default: "-123.456",
							},
							{
								Name: name.Path{"option", "e"},
							},
						},
						Arguments: []data.Argument{
							{
								Name: name.Path{"arg", "a"},
								Type: schema.TypeString,
							},
							{
								Name: name.Path{"arg", "b"},
								Type: schema.TypeInteger,
							},
							{
								Name: name.Path{"arg", "c"},
								Type: schema.TypeBoolean,
							},
							{
								Name: name.Path{"arg", "d"},
								Type: schema.TypeFloat,
							},
							{
								Name: name.Path{"arg", "e"},
							},
							{
								Name:     name.Path{"arg", "v"},
								Variadic: true,
							},
						},
						Subcommands: []data.Subcommand{
							{Name: name.Path{"sub3", "subx"}},
							{Name: name.Path{"sub3", "suby"}},
						},
					},
					{
						Name: name.Path{"sub3", "subx"},
					},
					{
						Name: name.Path{"sub3", "suby"},
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.packageName, func(t *testing.T) {
			originalSchema, err := schema.Load(bytes.NewBufferString(testcase.schemaYAML))
			if err != nil {
				t.Fatalf("fail to read load schema: %+v", err)
			}
			got, err := data.Construct(testcase.packageName, originalSchema)
			if err != nil {
				t.Fatalf("fail to read construct template data: %+v", err)
			}

			assert.Equal(t, testcase.want.Package, got.Package)
			assert.Equal(t, testcase.want.Generator, got.Generator)
			assert.Equal(t, testcase.want.GeneratorVersion, got.GeneratorVersion)
			{
				gotSchema, err := schema.Load(bytes.NewBufferString(got.SchemaYAML))
				if err != nil {
					t.Fatalf("fail to load schema: %+v", err)
				}
				test.AssertMatchSchema(t, originalSchema, gotSchema)
			}
			{
				want, got := testcase.want.Program, got.Program
				assertMatchProgram(t, want, got)
			}
			{
				want, got := testcase.want.Commands, got.Commands
				assert.Equal(t, len(want), len(got))
				for idx, want := range want {
					got := got[idx]
					assertMatchCommand(t, want, got)
				}
			}
		})
	}
}

func assertMatchProgram(t *testing.T, want, got data.Program) {
	t.Helper()
	assert.ElementsMatch(t, want.Name, got.Name)
	assert.Equal(t, want.Version, got.Version)
	{
		assert.Equal(t, len(want.Options), len(got.Options))
		for idx, want := range want.Options {
			got := got.Options[idx]
			assertMatchOption(t, want, got)
		}
	}
	{
		assert.Equal(t, len(want.Arguments), len(got.Arguments))
		for idx, want := range want.Arguments {
			got := got.Arguments[idx]
			assertMatchArgument(t, want, got)
		}
	}
	{
		assert.Equal(t, len(want.Subcommands), len(got.Subcommands))
		for idx, want := range want.Subcommands {
			got := got.Subcommands[idx]
			assertMatchSubcommand(t, want, got)
		}
	}

}

func assertMatchCommand(t *testing.T, want, got data.Command) {
	t.Helper()
	assert.ElementsMatch(t, want.Name, got.Name)
	{
		assert.Equal(t, len(want.Options), len(got.Options))
		for idx, want := range want.Options {
			got := got.Options[idx]
			assertMatchOption(t, want, got)
		}
	}
	{
		assert.Equal(t, len(want.Arguments), len(got.Arguments))
		for idx, want := range want.Arguments {
			got := got.Arguments[idx]
			assertMatchArgument(t, want, got)
		}
	}
	{
		assert.Equal(t, len(want.Subcommands), len(got.Subcommands))
		for idx, want := range want.Subcommands {
			got := got.Subcommands[idx]
			assertMatchSubcommand(t, want, got)
		}
	}

}

func assertMatchOption(t *testing.T, want, got data.Option) {
	t.Helper()
	assert.ElementsMatch(t, want.Name, got.Name)
	assert.ElementsMatch(t, want.Short, got.Short)
	assert.Equal(t, want.Default, got.Default)
	assert.Equal(t, want.Type, got.Type)
}

func assertMatchArgument(t *testing.T, want, got data.Argument) {
	t.Helper()
	assert.ElementsMatch(t, want.Name, got.Name)
	assert.Equal(t, want.Type, got.Type)
	assert.Equal(t, want.Variadic, got.Variadic)
}

func assertMatchSubcommand(t *testing.T, want, got data.Subcommand) {
	t.Helper()
	assert.ElementsMatch(t, want.Name, got.Name)
}
