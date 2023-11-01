package schema_test

import (
	"cliautor/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		sut       *schema.Option
		shouldErr bool
	}{
		{
			name: "valid",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeString,
				Default:     "abc",
			},
		},
		{
			name: "empty short name is valid",
			sut: &schema.Option{
				Short:       "",
				Description: "desc",
				Type:        schema.TypeString,
				Default:     "abc",
			},
		},
		{
			name: "empty type is valid",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Default:     "abc",
			},
		},
		{
			name: "too long short name",
			sut: &schema.Option{
				Short:       "-abc",
				Description: "desc",
				Type:        schema.TypeString,
				Default:     "abc",
			},
			shouldErr: true,
		},
		{
			name: "invalid prefixed short name",
			sut: &schema.Option{
				Short:       "+e",
				Description: "desc",
				Type:        schema.TypeString,
				Default:     "abc",
			},
			shouldErr: true,
		},
		{
			name: "unknown type",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.Type("Unknown"),
				Default:     "abc",
			},
			shouldErr: true,
		},
		{
			name: "string without default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeString,
			},
		},
		{
			name: "integer type",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeInteger,
				Default:     "-123",
			},
		},
		{
			name: "integer without default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeInteger,
			},
		},
		{
			name: "invalid integer default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeInteger,
				Default:     "abc",
			},
			shouldErr: true,
		},
		{
			name: "boolean type",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeBoolean,
				Default:     "true",
			},
		},
		{
			name: "boolean without default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeBoolean,
			},
		},
		{
			name: "invalid boolean default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeBoolean,
				Default:     "abc",
			},
			shouldErr: true,
		},
		{
			name: "float type",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeFloat,
				Default:     "-1.234",
			},
		},
		{
			name: "float without default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeFloat,
			},
		},
		{
			name: "invalid float default",
			sut: &schema.Option{
				Short:       "-o",
				Description: "desc",
				Type:        schema.TypeFloat,
				Default:     "abc",
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
