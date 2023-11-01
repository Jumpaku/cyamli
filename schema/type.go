package schema

import "errors"

type Type string

const (
	TypeUnspecified Type = ""
	TypeString      Type = "string"
	TypeInteger     Type = "integer"
	TypeFloat       Type = "float"
	TypeBoolean     Type = "boolean"
)

func (typ Type) Validate() error {
	switch typ {
	default:
		return errors.New(`unknown type: ` + string(typ))
	case TypeBoolean, TypeFloat, TypeInteger, TypeString, TypeUnspecified:
		return nil
	}
}
