package schema

import (
	"fmt"
	"io"

	"github.com/goccy/go-yaml"
)

type Schema struct {
	Program Program
}

func Load(reader io.Reader) (*Schema, error) {
	schema := Schema{}
	decoder := yaml.NewDecoder(reader, yaml.Strict())
	err := decoder.Decode(&schema.Program)
	if err != nil {
		return nil, fmt.Errorf(`fail to unmarshal yaml as schema: %w`, err)
	}
	if err := schema.Validate(); err != nil {
		return nil, fmt.Errorf(`fail to validate schema: %w`, err)
	}
	return &schema, nil
}

func (s *Schema) Save(writer io.Writer) error {
	err := yaml.NewEncoder(writer).Encode(s.Program)
	if err != nil {
		return fmt.Errorf("fail to marshal schema into yaml: %w", err)
	}
	return nil
}

func (s *Schema) Validate() error {
	return s.Program.Validate()
}
