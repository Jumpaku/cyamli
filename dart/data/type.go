package data

import (
	"github.com/Jumpaku/cyamli/schema"

	"github.com/Jumpaku/go-assert"
)

func DartType(t schema.Type, variadic bool) string {
	if variadic {
		switch t {
		default:
			return assert.Unexpected1[string]("unexpected type: %s", t)
		case schema.TypeBoolean:
			return "List<bool>"
		case schema.TypeFloat:
			return "List<double>"
		case schema.TypeInteger:
			return "List<int>"
		case schema.TypeString, schema.TypeUnspecified:
			return "List<String>"
		}
	} else {
		switch t {
		default:
			return assert.Unexpected1[string]("unexpected type: %s", t)
		case schema.TypeBoolean:
			return "bool"
		case schema.TypeFloat:
			return "double"
		case schema.TypeInteger:
			return "int"
		case schema.TypeString, schema.TypeUnspecified:
			return "String"
		}
	}
}
