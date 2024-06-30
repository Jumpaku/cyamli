# cyamli

A command line tool to generate interfaces for command line tools from YAML-based CLI schemas.

<!-- TOC -->

- [Overview](#overview)
- [Motivation](#motivation)
- [Installation](#installation)
    - [Using Go](#using-go)
    - [Using Docker](#using-docker)
    - [Downloading executable binary files](#downloading-executable-binary-files)
    - [Building from source](#building-from-source)
- [Usage](#usage)
    - [Usage with an example in Go](#usage-with-an-example-in-go)
    - [Usage with an example in Python3](#usage-with-an-example-in-python3)
    - [Usage with an example in Dart](#usage-with-an-example-in-dart)
- [Details](#details)
    - [Supported programming languages](#supported-programming-languages)
    - [Handling command line arguments](#handling-command-line-arguments)
    - [Documentation](#documentation)


## Overview

Developing console apps involves defining and parsing command line interfaces (CLIs) such as command line arguments, which consist of subcommands, options, and positional arguments.

`cyamli` is a schema-based code generator that generates APIs (Application Programming Interfaces, such as types and functions) to handle typed CLIs.
The schema of a typed CLI can be written in YAML according to the JSON schema at https://github.com/Jumpaku/cyamli/blob/main/schema/cli.schema.yaml ( JSON version is also available at https://github.com/Jumpaku/cyamli/blob/main/schema/cli.schema.yaml ).


## Motivation

- Schema-based approach leveraging standardized and consistent sources.
- Promoting typed CLIs for the benefits of static checking and code completion.
- Reducing boilerplate by automatically generating the necessary code.


## Installation

### Using Go

```shell
go install github.com/Jumpaku/cyamli@latest
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
cd cyamli
go install .
```


## Usage

Usage of `cyamli` is as follows:

1. Define a CLI as a YAML file.
2. Generate the API to parse the CLI in Go.
3. Assign functions to the generated API.

The detailed usages are described in the following sections.


### Usage with an example in Go

<details>
<summary>Click to expand</summary>

Assume a situation where you need to develop a console app in Go to fetch information from a database.


#### Define a CLI as a YAML file

The following YAML file, `cli.yaml`, defines a CLI for the example console app.

```yaml
name: demo
description: demo app to get table information from databases
subcommands:
  list:
    description: list tables
    options:
      -config:
        description: path to config file
        short: -c
  fetch:
    description: show information of tables
    options:
      -config:
        description: path to config file
        short: -c
      -verbose:
        description: show detailed contents for specified tables
        short: -v
        type: boolean
    arguments:
      - name: tables
        variadic: true
        description: names of tables to be described
```


#### Generate API to parse the CLI in Go

The following command reads a schema from `cli.yaml` and writes the Go API into `cli.gen.go`.

```shell
cyamli generate golang -schema-path=cli.yaml -out-path=cli.gen.go
```

`cli.gen.go` includes the following API:

```go
// CLI represents a root command.
type CLI struct
// NewCLI returns a CLI object.
func NewCLI() CLI
// Run parses command line arguments args and calls a corresponding function assigned in cli.
func Run(cli CLI, args []string) error
// GetDoc returns a help message corresponding to subcommand.
func GetDoc(subcommand []string) string
```


#### Assign functions to the generated API.

`NewCLI()` returns an object `cli` which represents a root command, and its descendant objects represent subcommands.
Each of them has a `FUNC` field.
A function assigned to this field will be called by `Run(cli, os.Args)`.

The following code snippet demonstrates an implementation for the example console app.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.List.FUNC = func(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.Fetch.FUNC = func(subcommand []string, input CLI_Fetch_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
```

The example console app can be executed as follows:

```shell
go run main.go cli.gen.go list -c=config.yaml
go run main.go cli.gen.go fetch -c=config.yaml -v table1 table2
```

</details>


### Usage with an example in Python3

<details>
<summary>Click to expand</summary>

<!--
Assume a situation where you need to develop a console app in Go to fetch information from a database.


#### Define a CLI as a YAML file

The following YAML file, `cli.yaml`, defines a CLI for the example console app.

```yaml
name: demo
description: demo app to get table information from databases
subcommands:
  list:
    description: list tables
    options:
      -config:
        description: path to config file
        short: -c
  fetch:
    description: show information of tables
    options:
      -config:
        description: path to config file
        short: -c
      -verbose:
        description: show detailed contents for specified tables
        short: -v
        type: boolean
    arguments:
      - name: tables
        variadic: true
        description: names of tables to be described
```


#### Generate API to parse the CLI in Go

The following command reads a schema from `cli.yaml` and writes the Go API into `cli.gen.go`.

```shell
cyamli generate golang -schema-path=cli.yaml -out-path=cli.gen.go
```

`cli.gen.go` includes the following API:

```go
// CLI represents a root command.
type CLI struct
// NewCLI returns a CLI object.
func NewCLI() CLI
// Run parses command line arguments args and calls a corresponding function assigned in cli.
func Run(cli CLI, args []string) error
// GetDoc returns a help message corresponding to subcommand.
func GetDoc(subcommand []string) string
```


#### Assign functions to the generated API.


`NewCLI()` returns an object `cli` which represents a root command, and its descendant objects represent subcommands.
Each of them has a `FUNC` field.
A function assigned to this field will be called by `Run(cli, os.Args)`.

The following code snippet demonstrates an implementation for the example console app.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.List.FUNC = func(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.Fetch.FUNC = func(subcommand []string, input CLI_Fetch_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
```

The example console app can be executed as follows:

```shell
go run main.go cli.gen.go list -c=config.yaml
go run main.go cli.gen.go fetch -c=config.yaml -v table1 table2
```
-->

</details>


### Usage with an example in Dart

<details>
<summary>Click to expand</summary>

<!--
Assume a situation where you need to develop a console app in Go to fetch information from a database.

#### Define a CLI as a YAML file

The following YAML file, `cli.yaml`, defines a CLI for the example console app.

```yaml
name: demo
description: demo app to get table information from databases
subcommands:
  list:
    description: list tables
    options:
      -config:
        description: path to config file
        short: -c
  fetch:
    description: show information of tables
    options:
      -config:
        description: path to config file
        short: -c
      -verbose:
        description: show detailed contents for specified tables
        short: -v
        type: boolean
    arguments:
      - name: tables
        variadic: true
        description: names of tables to be described
```


#### Generate API to parse the CLI in Go

The following command reads a schema from `cli.yaml` and writes the Go API into `cli.gen.go`.

```shell
cyamli generate golang -schema-path=cli.yaml -out-path=cli.gen.go
```

`cli.gen.go` includes the following API:

```go
// CLI represents a root command.
type CLI struct
// NewCLI returns a CLI object.
func NewCLI() CLI
// Run parses command line arguments args and calls a corresponding function assigned in cli.
func Run(cli CLI, args []string) error
// GetDoc returns a help message corresponding to subcommand.
func GetDoc(subcommand []string) string
```


#### Assign functions to the generated API.

`NewCLI()` returns an object `cli` which represents a root command, and its descendant objects represent subcommands.
Each of them has a `FUNC` field.
A function assigned to this field will be called by `Run(cli, os.Args)`.

The following code snippet demonstrates an implementation for the example console app.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.List.FUNC = func(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	cli.Fetch.FUNC = func(subcommand []string, input CLI_Fetch_Input, inputErr error) (err error) {
		fmt.Println(input, inputErr)
		fmt.Println(GetDoc(subcommand))
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
```

The example console app can be executed as follows:

```shell
go run main.go cli.gen.go list -c=config.yaml
go run main.go cli.gen.go fetch -c=config.yaml -v table1 table2
```
-->

</details>


## Details

### Supported programming languages

The following programming languages are currently supported:

* Go
* Python3
* Dart3
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
    - `<option_value>` must be a string that can be parsed as a value of the type of the option.
    - `=<option_value>` can be omitted if the type of the option is boolean.
- `<argument>` represents an argument, which must be a string that can be parsed as a value of the type of the argument.
    - Tokens after `--` are handled as arguments even if prefixed by `-`.


### Documentation

The documentation for `cyamli` command is provided at https://github.com/Jumpaku/cyamli/blob/main/cyamli-docs.md .
