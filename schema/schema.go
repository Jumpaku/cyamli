package schema

import (
	"fmt"
	"io"

	"github.com/Jumpaku/cliautor/name"

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

func (s *Schema) Walk(f func(path name.Path, cmd *Command) error) error {
	return walkImpl(nil, s.Program.Command(), f)
}

func walkImpl(path name.Path, cmd *Command, f func(path name.Path, cmd *Command) error) error {
	if err := f(path, cmd); err != nil {
		return err
	}
	for cmdName, cmd := range cmd.Subcommands {
		if err := walkImpl(name.Path(path).Append(cmdName), cmd, f); err != nil {
			return err
		}
	}
	return nil
}
