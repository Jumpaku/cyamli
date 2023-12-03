package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/golang/data"
	"github.com/Jumpaku/cyamli/name"

	"github.com/stretchr/testify/assert"
)

func TestCommand_CLIStructName(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name: name.Path{"cmd", "name", "abc"},
			},
			want: `CLI_CmdNameAbc`,
		},
		{
			sut: data.Command{
				Name: name.Path{"cmdname"},
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
				Name: name.Path{"cmd", "name", "abc"},
			},
			want: `CLI_CmdNameAbc_Input`,
		},
		{
			sut: data.Command{
				Name: name.Path{"cmdname"},
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

func TestCommand_FullPathLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Command
		want string
	}{
		{
			sut: data.Command{
				Name: name.Path{"cmd", "name", "abc"},
			},
			want: `"cmd name abc"`,
		},
		{
			sut: data.Command{
				Name: name.Path{"cmdname"},
			},
			want: `"cmdname"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.FullPathLiteral()
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
				Name: name.Path{"cmd", "name", "abc"},
			},
			want: `Cmd.Name.Abc.FUNC`,
		},
		{
			sut: data.Command{
				Name: name.Path{"cmdname"},
			},
			want: `Cmdname.FUNC`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIFuncMethodChain()
			assert.Equal(t, testcase.want, got)
		})
	}
}
