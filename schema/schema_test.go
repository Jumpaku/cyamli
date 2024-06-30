package schema_test

import (
	"bytes"
	_ "embed"
	"github.com/Jumpaku/cyamli/info"
	"github.com/Jumpaku/cyamli/schema"
	"github.com/kortschak/utter"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

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
		assert.Truef(t, ok, "option %q not found", name)
		assert.Equal(t, w.Short, g.Short)
		assert.Equal(t, w.Default, g.Default)
		assert.Equal(t, w.Description, g.Description)
		assert.Equal(t, w.Type, g.Type)
	}

	wArg, gArg := want.Arguments, got.Arguments
	for idx, w := range wArg {
		assert.Truef(t, idx < len(gArg), "argument %q not found at %d", w.Name, idx)
		g := gArg[idx]
		assert.Equal(t, w.Name, g.Name)
		assert.Equal(t, w.Variadic, g.Variadic)
		assert.Equal(t, w.Description, g.Description)
		assert.Equal(t, w.Type, g.Type)
	}

	wSub, gSub := want.Subcommands, got.Subcommands
	for name, w := range wSub {
		g, ok := gSub[name]
		assert.Truef(t, ok, "subcommand %q not found", name)
		assertMatchCommand(t, w, g)
	}
}

func mustRead(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var cyamliYAML = mustRead("testdata/cyamli.yaml")

var demoAppYAML = mustRead("testdata/demo-app.yaml")

var cyamliSchema = &schema.Schema{
	Program: schema.Program{
		Name:        string("cyamli"),
		Version:     string(info.Version),
		Description: string("A command line tool to generate CLI for your app from YAML-based schema."),
		Options: map[string]*schema.Option{
			string("-help"): &schema.Option{
				Short:       string("-h"),
				Description: string("shows description of this app."),
				Type:        schema.Type("boolean"),
				Default:     string(""),
			},
			string("-version"): &schema.Option{
				Short:       string("-v"),
				Description: string("shows version of this app."),
				Type:        schema.Type("boolean"),
				Default:     string(""),
			},
		},
		Arguments: []*schema.Argument(nil),
		Subcommands: map[string]*schema.Command{
			string("list"): &schema.Command{
				Description: string("shows subcommands"),
				Options: map[string]*schema.Option{
					string("-schema-path"): &schema.Option{
						Short:       string(""),
						Description: string("if specified then reads schema file from the path, otherwise reads from stdin."),
						Type:        schema.Type(""),
						Default:     string(""),
					},
				},
				Arguments:   []*schema.Argument(nil),
				Subcommands: map[string]*schema.Command(nil),
			},
			string("generate"): &schema.Command{
				Description: string("holds subcommands to generate CLI code."),
				Options:     map[string]*schema.Option(nil),
				Arguments:   []*schema.Argument(nil),
				Subcommands: map[string]*schema.Command{
					string("golang"): &schema.Command{
						Description: string("generates CLI for your app written in Go."),
						Options: map[string]*schema.Option{
							string("-package"): &schema.Option{
								Short:       string(""),
								Description: string("package name where the generated file will be placed."),
								Type:        schema.Type(""),
								Default:     string("main"),
							},
							string("-schema-path"): &schema.Option{
								Short:       string(""),
								Description: string("if specified then reads schema file from the path, otherwise reads from stdin."),
								Type:        schema.Type(""),
								Default:     string(""),
							},
							string("-out-path"): &schema.Option{
								Short:       string(""),
								Description: string("if specified then creates a file at the path and writes generated code, otherwise outputs to stdout."),
								Type:        schema.Type(""),
								Default:     string(""),
							},
							string("-help"): &schema.Option{
								Short:       string("-h"),
								Description: string("shows description of golang subcommand."),
								Type:        schema.Type("boolean"),
								Default:     string(""),
							},
						},
						Arguments:   []*schema.Argument(nil),
						Subcommands: map[string]*schema.Command(nil),
					},
					string("python3"): &schema.Command{
						Description: string("generates CLI for your app written in Python3."),
						Options: map[string]*schema.Option{
							string("-out-path"): &schema.Option{
								Short:       string(""),
								Description: string("if specified then creates a file at the path and writes generated code, otherwise outputs to stdout."),
								Type:        schema.Type(""),
								Default:     string(""),
							},
							string("-help"): &schema.Option{
								Short:       string("-h"),
								Description: string("shows description of python3 subcommand."),
								Type:        schema.Type("boolean"),
								Default:     string(""),
							},
							string("-schema-path"): &schema.Option{
								Short:       string(""),
								Description: string("if specified then reads schema file from the path, otherwise reads from stdin."),
								Type:        schema.Type(""),
								Default:     string(""),
							},
						},
						Arguments:   []*schema.Argument(nil),
						Subcommands: map[string]*schema.Command(nil),
					},
					string("docs"): &schema.Command{
						Description: string("generates documentation for your CLI app."),
						Options: map[string]*schema.Option{
							string("-all"): &schema.Option{
								Short:       string("-a"),
								Description: string("if specified then outputs documentation for all subcommands, otherwise in text format."),
								Type:        schema.Type("boolean"),
								Default:     string(""),
							},
							string("-schema-path"): &schema.Option{
								Short:       string(""),
								Description: string("if specified then reads schema file from the path, otherwise reads from stdin."),
								Type:        schema.Type(""),
								Default:     string(""),
							},
							string("-out-path"): &schema.Option{
								Short:       string(""),
								Description: string("if specified then creates a file at the path and writes generated documentation, otherwise outputs to stdout."),
								Type:        schema.Type(""),
								Default:     string(""),
							},
							string("-help"): &schema.Option{
								Short:       string("-h"),
								Description: string("shows description of docs subcommand."),
								Type:        schema.Type("boolean"),
								Default:     string(""),
							},
							string("-format"): &schema.Option{
								Short:       string("-f"),
								Description: string("specifies output format of the documentation in text or markdown."),
								Type:        schema.Type(""),
								Default:     string("text"),
							},
						},
						Arguments: []*schema.Argument{
							&schema.Argument{
								Name:        string("subcommands"),
								Description: string("selects subcommand for which the documentation is output."),
								Type:        schema.Type(""),
								Variadic:    bool(true),
							},
						},
						Subcommands: map[string]*schema.Command(nil),
					},
				},
			},
		},
	},
}

var demoAppSchema = &schema.Schema{
	Program: schema.Program{
		Name:        string("demo"),
		Version:     string(""),
		Description: string("demo app to get table information from databases"),
		Options:     map[string]*schema.Option(nil),
		Arguments:   []*schema.Argument(nil),
		Subcommands: map[string]*schema.Command{
			string("list"): &schema.Command{
				Description: string("list tables"),
				Options: map[string]*schema.Option{
					string("-config"): &schema.Option{
						Short:       string("-c"),
						Description: string("path to config file"),
						Type:        schema.Type(""),
						Default:     string(""),
					},
				},
				Arguments:   []*schema.Argument(nil),
				Subcommands: map[string]*schema.Command(nil),
			},
			string("fetch"): &schema.Command{
				Description: string("show information of tables"),
				Options: map[string]*schema.Option{
					string("-config"): &schema.Option{
						Short:       string("-c"),
						Description: string("path to config file"),
						Type:        schema.Type(""),
						Default:     string(""),
					},
					string("-verbose"): &schema.Option{
						Short:       string("-v"),
						Description: string("shows detailed log"),
						Type:        schema.Type("boolean"),
						Default:     string(""),
					},
				},
				Arguments: []*schema.Argument{
					&schema.Argument{
						Name:        string("tables"),
						Description: string("names of tables to be described"),
						Type:        schema.Type(""),
						Variadic:    bool(true),
					},
				},
				Subcommands: map[string]*schema.Command(nil),
			},
		},
	},
}

func TestSchema_Load(t *testing.T) {
	testcases := []struct {
		name      string
		in        string
		want      *schema.Schema
		shouldErr bool
	}{
		{
			name: "cyamli",
			in:   cyamliYAML,
			want: cyamliSchema,
		},
		{
			name: "demo-app",
			in:   demoAppYAML,
			want: demoAppSchema,
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

func TestSchema_Save(t *testing.T) {
	testcases := []struct {
		name      string
		sut       *schema.Schema
		want      string
		shouldErr bool
	}{
		{
			name: "cyamli",
			sut:  cyamliSchema,
			want: cyamliYAML,
		},
		{
			name: "demo-app",
			sut:  demoAppSchema,
			want: demoAppYAML,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			err := testcase.sut.Save(buffer)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				loaded, err := schema.Load(buffer)
				if err != nil {
					t.Fatalf("fail to load schema: %+v", err)
				}
				assertMatchSchema(t, testcase.sut, loaded)
			}
		})
	}
}

func TestSchema_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		in        io.Reader
		shouldErr bool
	}{
		{
			name:      "cyamli",
			in:        bytes.NewBufferString(cyamliYAML),
			shouldErr: false,
		},
		{
			name:      "demo-app",
			in:        bytes.NewBufferString(demoAppYAML),
			shouldErr: false,
		},
		{
			name:      "unknown property",
			in:        bytes.NewBufferString(`unknown: null`),
			shouldErr: true,
		},
		// option
		{
			name:      "valid option",
			in:        bytes.NewBufferString(`options: { -o: { short: -s, description: "", type: string, default: "" } }`),
			shouldErr: false,
		},
		{
			name:      "invalid option name",
			in:        bytes.NewBufferString(`options: { invalid: { short: -s, description: "", type: string, default: "" } }`),
			shouldErr: true,
		},
		{
			name:      "invalid option short name",
			in:        bytes.NewBufferString(`options: { -o: { short: invalid, description: "", type: string, default: "" } }`),
			shouldErr: true,
		},
		{
			name:      "invalid option property",
			in:        bytes.NewBufferString(`options: { -o: { invalid: null } }`),
			shouldErr: true,
		},
		{
			name:      "invalid option type",
			in:        bytes.NewBufferString(`options: { -o: { short: -s, description: "", type: invalid, default: "" } }`),
			shouldErr: true,
		},
		// argument
		{
			name:      "valid argument",
			in:        bytes.NewBufferString(`arguments: [ { name: a, description: "", type: string, variadic: false } ]`),
			shouldErr: false,
		},
		{
			name:      "invalid argument name",
			in:        bytes.NewBufferString(`arguments: [ { name: INVALID, description: "", type: string, variadic: false } ]`),
			shouldErr: true,
		},
		{
			name:      "invalid argument without name",
			in:        bytes.NewBufferString(`arguments: [ { description: "", type: string, variadic: false } ]`),
			shouldErr: true,
		},
		{
			name:      "invalid argument property",
			in:        bytes.NewBufferString(`arguments: [ { name: a, invalid: null } ]`),
			shouldErr: true,
		},
		{
			name:      "invalid argument type",
			in:        bytes.NewBufferString(`arguments: [ { name: a, description: "", type: invalid, variadic: false } ]`),
			shouldErr: true,
		},
		// subcommand
		{
			name:      "valid subcommand",
			in:        bytes.NewBufferString(`subcommands: { s: { description: "", options: {}, arguments: [], subcommands: {} } }`),
			shouldErr: false,
		},
		{
			name:      "invalid subcommand name",
			in:        bytes.NewBufferString(`subcommands: [ S: { description: "", options: {}, arguments: [], subcommands: {} } ]`),
			shouldErr: true,
		},
		{
			name:      "invalid subcommand property",
			in:        bytes.NewBufferString(`subcommands: { s: { invalid: null, description: "", options: {}, arguments: [], subcommands: {} } }`),
			shouldErr: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			err := schema.Validate(testcase.in)
			if testcase.shouldErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSchema_Dump(t *testing.T) {
	t.Skip()
	s, err := schema.Load(bytes.NewBufferString(demoAppYAML))
	if err != nil {
		t.Fatalf("fail to load schema: %+v", err)
	}
	utter.Dump(s)
}
