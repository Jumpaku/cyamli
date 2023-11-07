package schema_test

import (
	"testing"

	"github.com/Jumpaku/cyamli/schema"

	"github.com/stretchr/testify/assert"
)

func TestCommand_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		sut       *schema.Command
		shouldErr bool
	}{
		{
			name: "valid",
			sut: &schema.Command{
				Options: map[string]*schema.Option{
					"-o0": {},
					"-o1": {},
				},
				Arguments: []*schema.Argument{
					{Name: "arg0"},
					{Name: "arg1", Variadic: true},
				},
				Subcommands: map[string]*schema.Command{
					"sub0": {},
					"sub1": {},
				},
			},
		},
		{
			name: "empty is valid",
			sut:  &schema.Command{},
		},
		{
			name: "invalid option name",
			sut: &schema.Command{
				Options: map[string]*schema.Option{"": {}},
			},
			shouldErr: true,
		},
		{
			name: "duplicated option name",
			sut: &schema.Command{
				Options: map[string]*schema.Option{"-s": {}, "-o": {Short: "-s"}},
			},
			shouldErr: true,
		},
		{
			name: "empty options",
			sut: &schema.Command{
				Options: map[string]*schema.Option{"-o1": {}, "-o2": {}},
			},
			shouldErr: false,
		},
		{
			name: "duplicated option name",
			sut: &schema.Command{
				Options: map[string]*schema.Option{"-o": {Short: "-s"}, "-s": {}},
			},
			shouldErr: true,
		},
		{
			name: "duplicated option short name",
			sut: &schema.Command{
				Options: map[string]*schema.Option{"-o1": {Short: "-s"}, "-o2": {Short: "-s"}},
			},
			shouldErr: true,
		},
		{
			name: "invalid argument name",
			sut: &schema.Command{
				Arguments: []*schema.Argument{{}},
			},
			shouldErr: true,
		},
		{
			name: "duplicated argument name",
			sut: &schema.Command{
				Arguments: []*schema.Argument{{Name: "arg"}, {Name: "arg"}},
			},
			shouldErr: true,
		},
		{
			name: "invalid subcommand name",
			sut: &schema.Command{
				Subcommands: map[string]*schema.Command{"": {}},
			},
			shouldErr: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.sut.Validate()
			assert.Equalf(t, testcase.shouldErr, err != nil, "%+v", err)
		})
	}
}
