# cyamli (v1.1.7)

## cyamli

### Description

A command line tool to generate CLI for your app from YAML-based schema.

### Syntax

```shell
cyamli  [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-version[=<boolean>]`, `-v[=<boolean>]`  (default=`false`):  
  shows version of this app.  

### Subcommands

* generate:  
  holds subcommands to generate CLI code.  

* list:  
  shows subcommands  

* validate:  
  validates CLI schema.  


## cyamli generate

### Description

holds subcommands to generate CLI code.

### Syntax

```shell
cyamli generate [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of generate subcommand.  

### Subcommands

* dart:  
  generates CLI for your app written in Dart.  

* docs:  
  generates documentation for your CLI app.  

* golang:  
  generates CLI for your app written in Go.  

* python3:  
  generates CLI for your app written in Python3.  


## cyamli generate dart

### Description

generates CLI for your app written in Dart.

### Syntax

```shell
cyamli generate dart [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of dart subcommand.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  


## cyamli generate docs

### Description

generates documentation for your CLI app.

### Syntax

```shell
cyamli generate docs [<option>|<argument>]... [-- [<argument>]...]
```

### Options

* `-all[=<boolean>]`, `-a[=<boolean>]`  (default=`false`):  
  if specified then outputs documentation for all subcommands, otherwise in text format.  

* `-format=<string>`, `-f=<string>`  (default=`"text"`):  
  specifies output format of the documentation in text or markdown.  

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of docs subcommand.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated documentation, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  

### Arguments

1. `[<subcommands:string>]...`  
  selects subcommand for which the documentation is output.  


## cyamli generate golang

### Description

generates CLI for your app written in Go.

### Syntax

```shell
cyamli generate golang [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of golang subcommand.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-package=<string>`  (default=`"main"`):  
  package name where the generated file will be placed.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  


## cyamli generate python3

### Description

generates CLI for your app written in Python3.

### Syntax

```shell
cyamli generate python3 [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of python3 subcommand.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  


## cyamli list

### Description

shows subcommands

### Syntax

```shell
cyamli list [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of list subcommand.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  


## cyamli validate

### Description

validates CLI schema.

### Syntax

```shell
cyamli validate [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of validates subcommand.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  


