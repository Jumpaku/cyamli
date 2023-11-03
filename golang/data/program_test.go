package data_test

import (
	"cliautor/golang/data"
	"cliautor/name"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgram_CLIStructName(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name:        name.Path{"program"},
				Description: "program description",
			},
			want: `CLI`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIStructName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestProgram_CLIInputStructName(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name:        name.Path{"program"},
				Description: "program description",
			},
			want: `CLI_Input`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIInputStructName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestProgram_DescriptionLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name:        name.Path{"program"},
				Description: "program description",
			},
			want: `"program description"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.DescriptionLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestProgram_NameLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Program
		want string
	}{
		{
			sut: data.Program{
				Name:        name.Path{"program"},
				Description: "program description",
			},
			want: `""`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.NameLiteral()
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
				Name:        name.Path{"program"},
				Description: "program description",
			},
			want: `Func`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.CLIFuncMethodChain()
			assert.Equal(t, testcase.want, got)
		})
	}
}
