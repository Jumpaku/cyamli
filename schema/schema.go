package schema

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Schema struct {
	Program Program
}

func Load(b []byte) (*Schema, error) {
	schema := Schema{}
	err := yaml.Unmarshal(b, &schema.Program)
	if err != nil {
		return nil, fmt.Errorf(`fail to parse yaml as command structure: %w`, err)
	}
	if err := (&schema).Validate(); err != nil {
		return nil, fmt.Errorf(`fail to validate command: %w`, err)
	}
	return &schema, nil
}

func (s *Schema) Validate() error {
	return s.Program.Validate()
}
