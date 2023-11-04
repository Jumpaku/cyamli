package schema_test

import (
	"testing"

	"github.com/Jumpaku/cliautor/schema"

	"github.com/stretchr/testify/assert"
)

func TestArgument_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		sut       *schema.Argument
		shouldErr bool
	}{
		{
			name: "valid",
			sut: &schema.Argument{
				Name:        "a_r_g",
				Description: "desc",
				Type:        schema.TypeString,
				Variadic:    true,
			},
		},
		{
			name: "empty name",
			sut: &schema.Argument{
				Name:        "",
				Description: "desc",
				Type:        schema.TypeString,
				Variadic:    true,
			},
			shouldErr: true,
		},
		{
			name: "upper chain case name",
			sut: &schema.Argument{
				Name:        "A-R-G",
				Description: "desc",
				Type:        schema.TypeString,
				Variadic:    true,
			},
			shouldErr: true,
		},
		{
			name: "empty type is valid",
			sut: &schema.Argument{
				Name:        "a_r_g",
				Description: "desc",
				Variadic:    true,
			},
		},
		{
			name: "unknown type",
			sut: &schema.Argument{
				Name:        "a_r_g",
				Description: "desc",
				Type:        schema.Type("Unknown"),
				Variadic:    true,
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
