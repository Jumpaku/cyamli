# cyamli

A command line tool to generate command line interfaces for your command line tools from the YAML-based CLI schemas.

## Highlights

This repository:
- defines the CLI schema in YAML.
- provides a typed CLI code generator from the CLI schema.

## CLI schema

The CLI schema definition is provided in [`cli-schema-definition.ts`](https://github.com/Jumpaku/cyamli/blob/main/cli-schema-definition.ts).

## CLI code generator

From a YAML file written according to the CLI schema definition, `cyamli` generates a typed code for handling command line arguments.

### Installation

`cyamli` can be installed as follows:

```sh
go install "github.com/Jumpaku/cyamli/cmd/cyamli@latest" 
```

or use go generate as follows:

```go
//go:generate "github.com/Jumpaku/cyamli/cmd/cyamli@latest" -schema-path=path/to/cli.yaml -out-path=path/to/cli.gen.go
```

### Usage

1. Generating the CLI type.
2. Overwriting the CLI object.
3. Executing program.

#### Generating the CLI type

Prepare the CLI schema for your application, for example:

```yaml
name: greet
description: this is an example program
options:
  -help:
    short: -h
    description: Show help information.
    type: boolean
subcommands:
  hello:
    description: Prints "Hello, <target name>! My name is <greeter>!"
    options:
      -target-name:
        short: -t
        description: The name of the person to be said hello.
    arguments:
      - name: greeter
        description: The name of the person who says hello.
```

Run `cyamli` as follows:

```sh
cyamli golang < path/to/cli-schema.yaml > path/to/generated/code.go
```

The above generates a Go code which includes:

```go
type CLI 
type CLI_Input 
type CLI_Hello 
type CLI_Hello_Input

func NewCLI() CLI
func Run(cli CLI, args []string) error
```

#### Overwriting the CLI object

To define the behavior of your program, you can utilize the generated types and functions as follows:

```go
// Create the CLI object
var cli = NewCLI()

func main() {
	// Overwrite behaviors
	cli.FUNC = showHelp
	cli.Hello.FUNC = sayHello
	// Run with command line arguments
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
```

Example implementations for `showHelp` and `sayHello` are as follows:

```go
func showHelp(subcommand []string, input CLI_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(cli.DESC_Simple())
		panic(inputErr)
	}
	if input.Opt_Help {
		fmt.Println(cli.DESC_Detail())
	}
	return nil
}
func sayHello(subcommand []string, input CLI_Hello_Input, inputErr error) (err error) {
	if inputErr != nil {
		fmt.Println(cli.Hello.DESC_Simple())
		return inputErr
	}
	if input.Opt_TargetName != "" {
		fmt.Printf("Hello, %s! My name is %s!\n", input.Opt_TargetName, input.Arg_Greeter)
	} else {
		fmt.Printf("Hello! My name is %s!\n", input.Arg_Greeter)
	}
	return nil
}
```

#### Executing program

Execute the `main` function as follows

```sh
go run main.go -h
# => This is an example program.
go run main.go hello Alice
# => Hello! My name is Alice!
go run main.go hello -target-name=Bob Alice
# => Hello, Bob! My name is Alice!
```

## Examples

The example CLI applications are found in https://github.com/Jumpaku/cyamli/tree/main/examples .

## Details

### Supported programming languages

Only Go is supported currently.

### Handling command line arguments

```
<program> <subcommand> [<option>|<argument>]... [-- [<argument>]...]
```

- `<program>` is the path to your executable file.
- `<subcommand>` is a sequence of tokens, which represents a path in the command tree illustrated in your CLI schema.
	- Each element of `<subcommand>` must match the regular expression `^[a-z][a-z0-9]*$`.
    - `<subcommand>` may be empty, which means the execution of the root command.
- `<option>` represents an option, which is a token in form of `<option_name>[=<option_value>]`.
    - `<option_name>` must match the regular expression `^(-[a-z][a-z0-9]*)+$` (or `^-[a-z]$` in short version).
    - `<option_value>` must be a string that can be parsed as a value of the type of the option.
    - `=<option_value>` can be omitted if the type of the option is boolean.
- `<argument>` represents an argument, which must be a string that can be parsed as a value of the type of the argument.
    - Tokens after `--` are handled as arguments even if prefixed by `-`.

