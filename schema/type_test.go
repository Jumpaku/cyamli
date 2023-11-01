package schema_test

import (
	"cliautor/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		sut       schema.Type
		shouldErr bool
	}{
		{
			name: "unspecified",
			sut:  schema.TypeUnspecified,
		},
		{
			name: "string",
			sut:  schema.TypeString,
		},
		{
			name: "integer",
			sut:  schema.TypeInteger,
		},
		{
			name: "float",
			sut:  schema.TypeFloat,
		},
		{
			name: "boolean",
			sut:  schema.TypeBoolean,
		},
		{
			name:      "unknown",
			sut:       schema.Type("unknown"),
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
