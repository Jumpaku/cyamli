# Examples of CLI application using cyamli

## example

```sh
# Go to this directory.
cd examples

# Customize example/cli.yaml.

# Generate the CLI types.
go generate ./example/main.go

# Run the CLI application.
go run ./cmd/example
```

## greet

```sh
# Go to this directory.
cd greet 

# Customize greet/cli.yaml.

# Generate the CLI types.
go generate ./example/main.go

# Run the CLI application.
go run ./cmd/greet
```