package main

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"log"
	"os"
)

func main() {
	{
		y, err := os.ReadFile("schema/cli.schema.yaml")
		if err != nil {
			log.Panicf("%#v", err)
		}
		j, err := yaml.YAMLToJSON(y)
		if err != nil {
			log.Panicf("%#v", err)
		}
		err = os.WriteFile("schema/cli.schema.json", j, os.ModePerm)
		if err != nil {
			log.Panicf("%#v", err)
		}
	}
	sch, err := jsonschema.Compile("schema/cli.schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}
	{
		y, err := os.ReadFile("examples/cmd/example/cli.yaml")
		if err != nil {
			log.Panicf("%#v", err)
		}
		j, err := yaml.YAMLToJSON(y)
		if err != nil {
			log.Panicf("%#v", err)
		}
		var v any
		err = json.Unmarshal(j, &v)
		if err != nil {
			log.Panicf("%#v", err)
		}
		err = sch.Validate(v)
		if err != nil {
			log.Panicf("%#v", err)
		}
	}
}
