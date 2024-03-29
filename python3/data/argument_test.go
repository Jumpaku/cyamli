package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/python3/data"
	"github.com/Jumpaku/cyamli/schema"

	"github.com/stretchr/testify/assert"
)

func TestArgument_InputFieldName(t *testing.T) {
	testcases := []struct {
		sut  data.Argument
		want string
	}{
		{
			sut: data.Argument{
				Name:     name.Path{"arg", "name", "123"},
				Type:     schema.TypeInteger,
				Variadic: true,
			},
			want: `arg_arg_name_123`,
		},
		{
			sut: data.Argument{
				Name:     name.Path{"argname"},
				Type:     schema.TypeInteger,
				Variadic: true,
			},
			want: `arg_argname`,
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
				Type:     schema.TypeInteger,
				Variadic: true,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `tuple[int,...]`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeInteger,
				Variadic: false,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `int`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeBoolean,
				Variadic: true,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `tuple[bool,...]`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeBoolean,
				Variadic: false,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `bool`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeFloat,
				Variadic: true,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `tuple[float,...]`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeFloat,
				Variadic: false,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `float`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeString,
				Variadic: true,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `tuple[str,...]`,
		},
		{
			sut: data.Argument{
				Type:     schema.TypeString,
				Variadic: false,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `str`,
		},
		{
			sut: data.Argument{
				Variadic: true,
				Name:     name.Path{"arg", "name", "123"},
			},
			want: `tuple[str,...]`,
		},
		{
			sut: data.Argument{
				Name: name.Path{"arg", "name", "123"},
			},
			want: `str`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: type=%#v,variadic=%#v", number, testcase.sut.Type, testcase.sut.Variadic), func(t *testing.T) {
			got := testcase.sut.InputFieldType()
			assert.Equal(t, testcase.want, got)
		})
	}
}
