package data

import (
	"github.com/Jumpaku/cyamli/schema"

	"github.com/Jumpaku/go-assert"
)

func Python3Type(t schema.Type, variadic bool) string {
	if variadic {
		switch t {
		default:
			return assert.Unexpected1[string]("unexpected type: %s", t)
		case schema.TypeBoolean:
			return "tuple[bool,...]"
		case schema.TypeFloat:
			return "tuple[float,...]"
		case schema.TypeInteger:
			return "tuple[int,...]"
		case schema.TypeString, schema.TypeUnspecified:
			return "tuple[str,...]"
		}
	} else {
		switch t {
		default:
			return assert.Unexpected1[string]("unexpected type: %s", t)
		case schema.TypeBoolean:
			return "bool"
		case schema.TypeFloat:
			return "float"
		case schema.TypeInteger:
			return "int"
		case schema.TypeString, schema.TypeUnspecified:
			return "str"
		}
	}
}
