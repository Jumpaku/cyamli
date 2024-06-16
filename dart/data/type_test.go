package data

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cyamli/schema"

	"github.com/stretchr/testify/assert"
)

func TestDartType(t *testing.T) {
	testcases := []struct {
		in_type     schema.Type
		in_variadic bool
		want        string
	}{
		{
			in_type:     schema.TypeBoolean,
			in_variadic: true,
			want:        "List<bool>",
		},
		{
			in_type:     schema.TypeBoolean,
			in_variadic: false,
			want:        "bool",
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: true,
			want:        "List<int>",
		},
		{
			in_type:     schema.TypeInteger,
			in_variadic: false,
			want:        "int",
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: true,
			want:        "List<double>",
		},
		{
			in_type:     schema.TypeFloat,
			in_variadic: false,
			want:        "double",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: true,
			want:        "List<String>",
		},
		{
			in_type:     schema.TypeString,
			in_variadic: false,
			want:        "String",
		},
		{
			in_type:     schema.TypeUnspecified,
			in_variadic: true,
			want:        "List<String>",
		},
		{
			in_type:     schema.TypeUnspecified,
			in_variadic: false,
			want:        "String",
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("DartType(%v, %v)", testcase.in_type, testcase.in_variadic), func(t *testing.T) {
			got := DartType(testcase.in_type, testcase.in_variadic)
			assert.Equal(t, testcase.want, got)
		})
	}
}
