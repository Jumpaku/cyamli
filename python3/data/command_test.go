package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/python3/data"

	"github.com/stretchr/testify/assert"
)

func TestCommand_CLIClassName(t *testing.T) {
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
			got := testcase.sut.CLIClassName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestCommand_CLIInputClassName(t *testing.T) {
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
			got := testcase.sut.CLIInputClassName()
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
			want: `cmd.name.abc.FUNC`,
		},
		{
			sut: data.Command{
				Name: name.Path{"cmdname"},
			},
			want: `cmdname.FUNC`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIFuncMethodChain()
			assert.Equal(t, testcase.want, got)
		})
	}
}
