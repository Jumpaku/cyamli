package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/python3/data"

	"github.com/stretchr/testify/assert"
)

func TestProgram_CLIClassName(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name: name.Path{"program"},
			},
			want: `CLI`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIClassName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestProgram_CLIInputClassName(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name: name.Path{"program"},
			},
			want: `CLI_Input`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIInputClassName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestProgram_FullPathLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name: name.Path{"program"},
			},
			want: `""`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.FullPathLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestProgram_FuncMethodChain(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name: name.Path{"program"},
			},
			want: `FUNC`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIFuncMethodChain()
			assert.Equal(t, testcase.want, got)
		})
	}
}
