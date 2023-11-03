package data_test

import (
	"cliautor/golang/data"
	"cliautor/name"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand_CLIStructName(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name:        name.Path{"cmd", "name", "abc"},
				Description: "command description",
			},
			want: `CLI_CmdNameAbc`,
		},
		{
			sut: data.Command{
				Name:        name.Path{"cmdname"},
				Description: "command description",
			},
			want: `CLI_Cmdname`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIStructName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestCommand_CLIInputStructName(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name:        name.Path{"cmd", "name", "abc"},
				Description: "command description",
			},
			want: `CLI_CmdNameAbc_Input`,
		},
		{
			sut: data.Command{
				Name:        name.Path{"cmdname"},
				Description: "command description",
			},
			want: `CLI_Cmdname_Input`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIInputStructName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestCommand_DescriptionLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name:        name.Path{"cmd", "name", "abc"},
				Description: "command description",
			},
			want: `"command description"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.DescriptionLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestCommand_NameLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name:        name.Path{"cmd", "name", "abc"},
				Description: "command description",
			},
			want: `"cmd name abc"`,
		},
		{
			sut: data.Command{
				Name:        name.Path{"cmdname"},
				Description: "command description",
			},
			want: `"cmdname"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.NameLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestCommand_FuncMethodChain(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name:        name.Path{"cmd", "name", "abc"},
				Description: "command description",
			},
			want: `Sub_Cmd.Sub_Name.Sub_Abc.Func`,
		},
		{
			sut: data.Command{
				Name:        name.Path{"cmdname"},
				Description: "command description",
			},
			want: `Sub_Cmdname.Func`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIFuncMethodChain()
			assert.Equal(t, testcase.want, got)
		})
	}
}
