package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cliautor/golang/data"
	"github.com/Jumpaku/cliautor/name"

	"github.com/stretchr/testify/assert"
)

func TestSubcommand_SubcommandFieldName(t *testing.T) {
	testcases := []struct {
		sut  data.Subcommand
		want string
	}{
		{
			sut: data.Subcommand{
				Name: name.Path{"sub", "cmd", "name"},
			},
			want: `Sub_Name`,
		},
		{
			sut: data.Subcommand{
				Name: name.Path{"subcmdname"},
			},
			want: `Sub_Subcmdname`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.SubcommandFieldName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestSubcommand_SubcommandFieldType(t *testing.T) {
	testcases := []struct {
		sut  data.Subcommand
		want string
	}{
		{
			sut: data.Subcommand{
				Name: name.Path{"sub", "cmd", "name"},
			},
			want: `CLI_SubCmdName`,
		},
		{
			sut: data.Subcommand{
				Name: name.Path{"subcmdname"},
			},
			want: `CLI_Subcmdname`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.SubcommandFieldType()
			assert.Equal(t, testcase.want, got)
		})
	}
}
