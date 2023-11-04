package schema_test

import (
	"testing"

	"github.com/Jumpaku/cliautor/schema"

	"github.com/stretchr/testify/assert"
)

func TestProgram_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		sut       *schema.Program
		shouldErr bool
	}{
		{
			name: "valid",
			sut: &schema.Program{
				Name:    "prog",
				Version: "0.0.0",
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
			sut:  &schema.Program{},
		},
		{
			name: "invalid option name",
			sut: &schema.Program{
				Options: map[string]*schema.Option{"": {}},
			},
			shouldErr: true,
		},
		{
			name: "duplicated option short name",
			sut: &schema.Program{
				Options: map[string]*schema.Option{"-o1": {Short: "-s"}, "-o2": {Short: "-s"}},
			},
			shouldErr: true,
		},
		{
			name: "invalid argument name",
			sut: &schema.Program{
				Arguments: []*schema.Argument{{}},
			},
			shouldErr: true,
		},
		{
			name: "duplicated argument name",
			sut: &schema.Program{
				Arguments: []*schema.Argument{{Name: "arg"}, {Name: "arg"}},
			},
			shouldErr: true,
		},
		{
			name: "invalid subcommand name",
			sut: &schema.Program{
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
