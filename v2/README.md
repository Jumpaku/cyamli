# cyamli

A command line tool to generate interfaces for command line tools from YAML-based CLI schemas.

## Overview

Console app development involves defining and parsing command line interfaces (CLIs) such as command line arguments, which consist of subcommands, options, and positional arguments.

`cyamli` is a schema-based code generator that generates APIs (Application Programming Interfaces, such as types and functions) to handle typed CLIs.
The schema of a typed CLI can be written in YAML according to the JSON Schema which is available at https://github.com/Jumpaku/cyamli/blob/main/docs/cyamli-cli.schema.json .

## Motivation

- Introducing the schema-based approach utilizing language-independent and consistent sources.
- Promoting typed CLIs for the benefits of static checking and code completion.
- Reducing boilerplate by generating the necessary code automatically.


## Installation

### Using Go

```shell
go install github.com/Jumpaku/cyamli/v2/cmd/cyamli@latest
```


### Using Docker

```shell
docker run -i -v $(pwd):/workspace ghcr.io/jumpaku/cyamli:latest cyamli
```


### Downloading executable binary files

https://github.com/Jumpaku/cyamli/releases

Note that the downloaded executable binary file may require a security confirmation before it can be run.


### Building from source

```shell
git clone https://github.com/Jumpaku/cyamli.git
cd cyamli/v2
go install ./cmd/cyamli
```


## Usage with an example

Assume a situation where you need to develop a console app in Go to fetch information from a database.

Usage of `cyamli` is as follows:

1. Define a CLI as a YAML file.
2. Generate the API to parse the CLI in Go.
3. Implement the interface in the generated API.

### Define a CLI as a YAML file

The following YAML file, `cli.yaml`, defines a CLI for the example console app.

```yaml
name: demo
options:
  -config:
    description: path to the config file, which can be available in subcommands.
    propagates: true
subcommands:
  tables:
    description: list tables from the database.
  schema:
    description: fetch schema of the specified table from the database.
    arguments:
      - name: table
  data:
    description: dump data from the specified table from the database.
    options:
      -where:
        description: filter data by the condition.
    arguments:
      - name: table
```


### Generate API to parse the CLI in Go

The following command reads a schema from `cli.yaml` and writes the Go API into `cli.gen.go`.

```shell
cyamli generate golang -schema-path=cli.yaml -out-path=cli.gen.go
```

`cli.gen.go` includes the following API:

```go
// Run is an entry point to handle the CLI.
func Run(handler CLIHandler, args []string) error

// CLIHandler is an interface that defines handlers for subcommands.
type CLIHandler interface {
	// ...
	// Run_Data is called for the data subcommand.
	Run_Data(input Input_Data) error
	// ...
}

//...
// Input_Data is parsed from command line arguments for the data subcommand.
type Input_Data struct {
	Opt_Config string
	Opt_Where  string
	Arg_Table  string
	// ...
}
//...

// GetVersion returns the version of the program.
func GetVersion() string
// GetProgram returns the name of the program.
func GetProgram() string
// GetDoc returns the documentation for a subcommand.
func GetDoc(subcommandPath []string) string
```

### Implement the interface in the generated API

The following code snippet demonstrates an implementation for the example console app.

```go
package main

import "os"

func main() {
	// Pass a CLI handler instance to the entry point.
	if err := Run(cli{}, os.Args); err != nil {
		panic(err)
	}
}

// cli implements CLIHandler interface.
type cli struct{}
// ...
// Run_Data is called for the data subcommand.
func (c cli) Run_Data(input Input_Data) error {
	// do something with input.Arg_Table, input.Opt_Config, and input.Opt_Where
	return nil
}
// ...
```

The example console app can be executed as follows:

```shell
go run main.go cli.gen.go data -config=config.yaml -where="age > 20" Users
```

## Details

### Supported features

* [x] Subcommands
* [x] Options
  * [x] Short options
  * [x] Primitive types
    * [x] boolean
      * [x] negated option with `-no` prefix
    * [x] integer
    * [x] string
  * [x] Default values
  * [x] Repeated options
  * [x] Option propagation to descendant commands
* [x] Arguments
  * [x] Primitive types
    * [x] boolean
    * [x] integer
    * [x] string
  * [x] Variadic arguments
* [x] Help documentation

### Supported programming languages

The following programming languages are currently supported:

* Go
* Python3
* Dart3
* Kotlin
* TypeScript
* C#
* C++ 11
* PHP 7.4
* Documentation in text, HTML, and Markdown


### Handling command line arguments

Command line arguments according to the following syntax can be handled by the generated API.

```
<program> <subcommand> [<option>|<argument>]... [-- [<argument>]...]
```

- `<program>` is the path to your executable file.
- `<subcommand>` is a sequence of tokens, which represents a path in the command tree illustrated in a defined CLI schema.
	- Each element of `<subcommand>` must match the regular expression `^[a-z][a-z0-9]*$`.
    - `<subcommand>` may be empty, which means the execution of the root command.
- `<option>` represents an option, which is a token in form of `<option_name>[=<option_value>]`.
    - `<option_name>` must match the regular expression `^(-[a-z][a-z0-9]*)+$` (or `^-[a-z]$` in short version).
    - `=<option_value>` can be omitted if the type of the option is boolean.
      - `<option_value>` must be a string that can be parsed as a value of the type of the option.
- `<argument>` represents an argument, which must be a string that can be parsed as a value of the type of the argument.
    - Tokens after `--` are handled as arguments even if prefixed by `-`.


### Usage of cyamli command

The documentation for `cyamli` command is provided at https://github.com/Jumpaku/cyamli/blob/main/docs/cyamli-docs.md .
