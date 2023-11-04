package testdata

import _ "embed"

//go:embed empty.yaml
var EmptyYAML string

//go:embed example.yaml
var ExampleYAML string
