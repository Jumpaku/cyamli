# cyamli (v2.0.0-alpha.5)


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

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  

### Subcommands

* cpp:  
  generates CLI for your app written in C++ 11.  

* csharp:  
  generates CLI for your app written in C#.  

* dart3:  
  generates CLI for your app written in Dart.  

* docs:  
  generates documentation for your CLI app.  

* golang:  
  generates CLI for your app written in Go.  

* kotlin:  
  generates CLI for your app written in Kotlin.  

* php:  
  generates CLI for your app written in PHP 7.4.  

* python3:  
  generates CLI for your app written in Python3.  

* typescript:  
  generates CLI for your app written in TypeScript.  




## cyamli generate cpp

### Description

generates CLI for your app written in C++ 11.

### Syntax

```shell
cyamli generate cpp [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-namespace=<string>`  (default=`""`):  
  namespace where the generated file will be placed.  

* `-out-header-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-out-source-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




## cyamli generate csharp

### Description

generates CLI for your app written in C#.

### Syntax

```shell
cyamli generate csharp [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-namespace=<string>`  (default=`""`):  
  namespace where the generated file will be placed.  

* `-out-path=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

* `-schema-path=<string>`  (default=`""`):  
  if specified then reads schema file from the path, otherwise reads from stdin.  




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




## cyamli generate php

### Description

generates CLI for your app written in PHP 7.4.

### Syntax

```shell
cyamli generate php [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  shows description of this app.  

* `-namespace=<string>`  (default=`""`):  
  namespace where the generated file will be placed.  

* `-out-dir=<string>`  (default=`""`):  
  if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.  

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




