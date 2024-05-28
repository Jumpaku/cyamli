package schema

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/Jumpaku/cyamli/name"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"io"
)

type Schema struct {
	Program Program
}

func Validate(reader io.Reader) error {
	b, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf(`fail to read schema: %w`, err)
	}

	if err := validateByJSONSchema(b); err != nil {
		return fmt.Errorf(`fail to validate schema based on JSON schema: %w`, err)
	}

	_, err = decodeSchemaYAML(bytes.NewBuffer(b))
	return err
}

func Load(reader io.Reader) (*Schema, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf(`fail to read schema: %w`, err)
	}

	if err := validateByJSONSchema(b); err != nil {
		return nil, fmt.Errorf(`fail to validate schema based on JSON schema: %w`, err)
	}

	return decodeSchemaYAML(bytes.NewBuffer(b))
}

func decodeSchemaYAML(reader io.Reader) (*Schema, error) {
	decoder := yaml.NewDecoder(reader, yaml.Strict())
	schema := Schema{}
	if err := decoder.Decode(&schema.Program); err != nil {
		return nil, fmt.Errorf(`fail to unmarshal yaml as schema: %w`, err)
	}

	if err := schema.Validate(); err != nil {
		return nil, fmt.Errorf(`fail to validate schema: %w`, err)
	}

	return &schema, nil
}

//go:generate go run ../internal/cmd/yaml-to-json -in=schema/cli.schema.yaml -out=schema/cli.schema.json
//go:embed cli.schema.json
var cliSchemaJSON string

func validateByJSONSchema(b []byte) error {
	schema, err := jsonschema.CompileString("", cliSchemaJSON)
	if err != nil {
		return fmt.Errorf("fail to compile schema: %w", err)
	}

	jb, err := yaml.YAMLToJSON(b)
	if err != nil {
		return fmt.Errorf("fail to convert YAML to JSON: %w", err)
	}

	var v any
	if err := json.Unmarshal(jb, &v); err != nil {
		return fmt.Errorf("fail to unmarshal JSON: %w", err)
	}

	if err := schema.Validate(v); err != nil {
		return fmt.Errorf("fail to validate schema: %w", err)
	}

	return nil
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

func (s *Schema) Find(subcommandPath []string) *Command {
	cmd := s.Program.Command()
	for _, subcommand := range subcommandPath {
		found, ok := cmd.Subcommands[subcommand]
		if !ok {
			return nil
		}

		cmd = found
	}
	return cmd
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
