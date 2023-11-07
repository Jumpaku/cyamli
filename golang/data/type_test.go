package data

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/schema"

	"github.com/stretchr/testify/assert"
)

func TestGoType(t *testing.T) {
	testcases := []struct {
		in_type     schema.Type
		in_variadic bool
		want        string
	}{
		{
			in_type:     schema.TypeBoolean,
			in_variadic: true,
			want:        "[]bool",
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: false,
			want:        "bool",
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: true,
			want:        "[]int64",
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: false,
			want:        "int64",
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: true,
			want:        "[]float64",
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: false,
			want:        "float64",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: true,
			want:        "[]string",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: false,
			want:        "string",
		},
		{
			in_type:     schema.TypeUnspecified,
			in_variadic: true,
			want:        "[]string",
		},
		{
			in_type:     schema.TypeUnspecified,
			in_variadic: false,
			want:        "string",
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("GoType(%v, %v)", testcase.in_type, testcase.in_variadic), func(t *testing.T) {
			got := GoType(testcase.in_type, testcase.in_variadic)
			assert.Equal(t, testcase.want, got)
		})
	}
}
