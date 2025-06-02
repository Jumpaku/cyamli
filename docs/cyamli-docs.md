# cyamli (v2.0.0-alpha.2)


## cyamli

### Description

A command line tool to generate CLI for your app from YAML-based schema.

### Syntax

```shell
cyamli [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

### Subcommands

* generate:  
  holds subcommands to generate CLI code.  

* version:  
  shows version of this app.  




## cyamli generate

### Description

holds subcommands to generate CLI code.

### Syntax

```shell
cyamli generate [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  

### Subcommands

* dart3:  
  generates CLI for your app written in Dart.  

* docs:  
  generates documentation for your CLI app.  

* golang:  
  generates CLI for your app written in Go.  

* kotlin:  
  generates CLI for your app written in Kotlin.  

* python3:  
  generates CLI for your app written in Python3.  

* typescript:  
  generates CLI for your app written in TypeScript.  




## cyamli generate dart3

### Description

generates CLI for your app written in Dart.

### Syntax

```shell
cyamli generate dart3 [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




## cyamli generate docs

### Description

generates documentation for your CLI app.

### Syntax

```shell
cyamli generate docs [<option>]...
```

### Options

* `-format=<string>`, `-f=<string>`  (default=`"text"`):  
  specifies output format of the documentation in text or markdown.  

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




## cyamli generate golang

### Description

generates CLI for your app written in Go.

### Syntax

```shell
cyamli generate golang [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-package=<string>`  (default=`"main"`):  
  package name where the generated file will be placed.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




## cyamli generate kotlin

### Description

generates CLI for your app written in Kotlin.

### Syntax

```shell
cyamli generate kotlin [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-package=<string>`  (default=`""`):  
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
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




## cyamli generate typescript

### Description

generates CLI for your app written in TypeScript.

### Syntax

```shell
cyamli generate typescript [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




## cyamli version

### Description

shows version of this app.

### Syntax

```shell
cyamli version [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  




