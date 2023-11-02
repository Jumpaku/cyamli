package data_test

import (
	"cliautor/golang/data"
	"cliautor/name"
	"cliautor/schema"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgument_InputFieldName(t *testing.T) {
	testcases := []struct {
		sut  data.Argument
		want string
	}{
		{
			sut: data.Argument{
				Name:        name.Path{"arg", "name", "123"},
				Type:        schema.TypeInteger,
				Variadic:    true,
				Description: "argument description",
			},
			want: `Arg_ArgName123`,
		},
		{
			sut: data.Argument{
				Name:        name.Path{"argname"},
				Type:        schema.TypeInteger,
				Variadic:    true,
				Description: "argument description",
			},
			want: `Arg_Argname`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.InputFieldName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestArgument_InputFieldType(t *testing.T) {
	testcases := []struct {
		sut  data.Argument
		want string
	}{
		{
			sut: data.Argument{
				Type:        schema.TypeInteger,
				Variadic:    true,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `[]int64`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeInteger,
				Variadic:    false,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `int64`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeBoolean,
				Variadic:    true,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `[]bool`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeBoolean,
				Variadic:    false,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `bool`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeFloat,
				Variadic:    true,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `[]float64`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeFloat,
				Variadic:    false,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `float64`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeString,
				Variadic:    true,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `[]string`,
		},
		{
			sut: data.Argument{
				Type:        schema.TypeString,
				Variadic:    false,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `string`,
		},
		{
			sut: data.Argument{
				Variadic:    true,
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `[]string`,
		},
		{
			sut: data.Argument{
				Name:        name.Path{"arg", "name", "123"},
				Description: "argument description",
			},
			want: `string`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: type=%#v,variadic=%#v", number, testcase.sut.Type, testcase.sut.Variadic), func(t *testing.T) {
			got := testcase.sut.InputFieldType()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestArgument_DescriptionLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Argument
		want string
	}{
		{
			sut: data.Argument{
				Name:        name.Path{"arg", "name", "123"},
				Type:        schema.TypeInteger,
				Variadic:    true,
				Description: "argument description",
			},
			want: `"argument description"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Description), func(t *testing.T) {
			got := testcase.sut.DescriptionLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}
