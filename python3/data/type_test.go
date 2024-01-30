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
			want:        "tuple[bool,...]",
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: false,
			want:        "bool",
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: true,
			want:        "tuple[int,...]",
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: false,
			want:        "int",
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: true,
			want:        "tuple[float,...]",
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: false,
			want:        "float",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: true,
			want:        "tuple[str,...]",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: false,
			want:        "str",
		},
		{
			in_type:     schema.TypeUnspecified,
			in_variadic: true,
			want:        "tuple[str,...]",
		},
		{
			in_type:     schema.TypeUnspecified,
			in_variadic: false,
			want:        "str",
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Python3Type(%v, %v)", testcase.in_type, testcase.in_variadic), func(t *testing.T) {
			got := Python3Type(testcase.in_type, testcase.in_variadic)
			assert.Equal(t, testcase.want, got)
		})
	}
}
