package data

import (
	"github.com/Jumpaku/cyamli/schema"

	"github.com/Jumpaku/go-assert"
)

func GoType(t schema.Type, variadic bool) string {
	if variadic {
		switch t {
		default:
			return assert.Unexpected1[string]("unexpected type: %s", t)
		case schema.TypeBoolean:
			return "[]bool"
		case schema.TypeFloat:
			return "[]float64"
		case schema.TypeInteger:
			return "[]int64"
		case schema.TypeString, schema.TypeUnspecified:
			return "[]string"
		}
	} else {
		switch t {
		default:
			return assert.Unexpected1[string]("unexpected type: %s", t)
		case schema.TypeBoolean:
			return "bool"
		case schema.TypeFloat:
			return "float64"
		case schema.TypeInteger:
			return "int64"
		case schema.TypeString, schema.TypeUnspecified:
			return "string"
		}
	}
}
