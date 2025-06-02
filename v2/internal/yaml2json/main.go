package main

import (
	"encoding/json"
	"fmt"
	"github.com/goccy/go-yaml"
	"os"
)

func main() {
	var v any
	if err := yaml.NewDecoder(os.Stdin).Decode(&v); err != nil {
		panic(fmt.Sprintf("failed to decode YAML: %+v", err))
	}
	if err := json.NewEncoder(os.Stdout).Encode(v); err != nil {
		panic(fmt.Sprintf("failed to encode JSON: %+v", err))
	}
}
