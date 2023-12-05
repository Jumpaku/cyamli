package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/golang/data"
	"github.com/Jumpaku/cyamli/name"

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
			want: `Name`,
		},
		{
			sut: data.Subcommand{
				Name: name.Path{"subcmdname"},
			},
			want: `Subcmdname`,
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
