// GENERATED CODE - DO NOT MODIFY BY HAND
// generator: {{.Generator}} {{.GeneratorVersion}}

// ignore_for_file: unused_local_variable

typedef Func<Input> = Function(List<String> subcommand, Input? input, Exception? inputErr);


{{/* Root command */}}
{{with .Program}}

class {{.CLIClassName}} {
  Func<{{.CLIInputRecordName}}>? FUNC;
  {{range $Index, $Subcommand := .Subcommands}}
  final {{$Subcommand.SubcommandFieldType}} {{$Subcommand.SubcommandFieldName}} = {{$Subcommand.SubcommandFieldType}}();
  {{end}}
}

{{if (or .Options .Arguments)}}
typedef {{.CLIInputRecordName}} = ({
  {{range $Index, $Option := .Options}}
  {{$Option.InputFieldType}} {{$Option.InputFieldName}},
  {{end}}
  {{range $Index, $Argument := .Arguments}}
  {{$Argument.InputFieldType}} {{$Argument.InputFieldName}},
  {{end}}
});
{{else}}
typedef {{.CLIInputRecordName}} = ();
{{end}}

{{.CLIInputRecordName}} _resolve_{{.CLIInputRecordName}}(List<String> restArgs) {
  {{range $Index, $Option := .Options}}
  {{$Option.InputFieldType}} var_{{$Option.InputFieldName}} = {{$Option.DefaultLiteral}};
  {{end}}
  List<String> arguments = [];
  for (int idx = 0; idx < restArgs.length; ++idx) {
    final arg = restArgs[idx];
    if (arg == "--") {
      arguments.addAll(restArgs.sublist(idx + 1));
      break;
    }
    if (!arg.startsWith("-")) {
      arguments.add(arg);
      continue;
    }
    final eqPos = arg.indexOf("=");
    final cut = eqPos >= 0;
    final optName = cut?arg.substring(0,eqPos) : arg;
    var lit = cut?arg.substring(eqPos+1) : "";

    switch (optName) {
    {{range $Index, $Option := .Options}}
    case {{$Option.NameLiteral}}{{if $Option.ShortNameLiteral}} || {{$Option.ShortNameLiteral}}{{end}}:
      if (!cut) {
        {{if eq $Option.InputFieldType "bool"}}lit = "true";
        {{else}} throw Exception("value is not specified to option ${optName}");
        {{end}}
      }
      var_{{$Option.InputFieldName}} = _parseValue(var_{{$Option.InputFieldName}}.runtimeType, [lit]) as {{$Option.InputFieldType}};
    {{end}}
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  {{range $Index, $Argument := .Arguments}}

  {{if $Argument.Variadic}}
  if (arguments.length <= {{$Index}} - 1) {
    throw Exception("too few arguments");
  }
  {{$Argument.InputFieldType}} var_{{$Argument.InputFieldName}} = _parseValue({{$Argument.InputFieldType}}, arguments.sublist({{$Index}})) as {{$Argument.InputFieldType}};
  {{else}}
  if (arguments.length <= {{$Index}}) {
    throw Exception("too few arguments");
  }
  {{$Argument.InputFieldType}} var_{{$Argument.InputFieldName}} = _parseValue({{$Argument.InputFieldType}}, [arguments[{{$Index}}]]) as {{$Argument.InputFieldType}};
  {{end}}

  {{end}}

  return (
  {{range $Index, $Option := .Options}}
    {{$Option.InputFieldName}}: var_{{$Option.InputFieldName}},
  {{end}}
  {{range $Index, $Argument := .Arguments}}
    {{$Argument.InputFieldName}}: var_{{$Argument.InputFieldName}},
  {{end}}
  );
}

{{end}}

{{/* Child commands */}}
{{range .Commands}}

class {{.CLIClassName}} {
  Func<{{.CLIInputRecordName}}>? FUNC;
  {{range $Index, $Subcommand := .Subcommands}}
  final {{$Subcommand.SubcommandFieldType}} {{$Subcommand.SubcommandFieldName}} = {{$Subcommand.SubcommandFieldType}}();
  {{end}}
}

{{if (or .Options .Arguments)}}
typedef {{.CLIInputRecordName}} = ({
  {{range $Index, $Option := .Options}}
  {{$Option.InputFieldType}} {{$Option.InputFieldName}},
  {{end}}
  {{range $Index, $Argument := .Arguments}}
  {{$Argument.InputFieldType}} {{$Argument.InputFieldName}},
  {{end}}
});
{{else}}
typedef {{.CLIInputRecordName}} = ();
{{end}}

{{.CLIInputRecordName}} _resolve_{{.CLIInputRecordName}}(List<String> restArgs) {
  {{range $Index, $Option := .Options}}
  {{$Option.InputFieldType}} var_{{$Option.InputFieldName}} = {{$Option.DefaultLiteral}};
  {{end}}
  List<String> arguments = [];
  for (int idx = 0; idx < restArgs.length; ++idx) {
    final arg = restArgs[idx];
    if (arg == "--") {
      arguments.addAll(restArgs.sublist(idx + 1));
      break;
    }
    if (!arg.startsWith("-")) {
      arguments.add(arg);
      continue;
    }
    final eqPos = arg.indexOf("=");
    final cut = eqPos >= 0;
    final optName = cut?arg.substring(0,eqPos) : arg;
    var lit = cut?arg.substring(eqPos+1) : "";

    switch (optName) {
    {{range $Index, $Option := .Options}}
    case {{$Option.NameLiteral}}{{if $Option.ShortNameLiteral}} || {{$Option.ShortNameLiteral}}{{end}}:
      if (!cut) {
        {{if eq $Option.InputFieldType "bool"}}lit = "true";
        {{else}} throw Exception("value is not specified to option ${optName}");
        {{end}}
      }
      var_{{$Option.InputFieldName}} = _parseValue(var_{{$Option.InputFieldName}}.runtimeType, [lit]) as {{$Option.InputFieldType}};
    {{end}}
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  {{range $Index, $Argument := .Arguments}}

  {{if $Argument.Variadic}}
  if (arguments.length <= {{$Index}} - 1) {
    throw Exception("too few arguments");
  }
  {{$Argument.InputFieldType}} var_{{$Argument.InputFieldName}} = _parseValue({{$Argument.InputFieldType}}, arguments.sublist({{$Index}})) as {{$Argument.InputFieldType}};
  {{else}}
  if (arguments.length <= {{$Index}}) {
    throw Exception("too few arguments");
  }
  {{$Argument.InputFieldType}} var_{{$Argument.InputFieldName}} = _parseValue({{$Argument.InputFieldType}}, [arguments[{{$Index}}]]) as {{$Argument.InputFieldType}};
  {{end}}

  {{end}}

  return (
  {{range $Index, $Option := .Options}}
    {{$Option.InputFieldName}}: var_{{$Option.InputFieldName}},
  {{end}}
  {{range $Index, $Argument := .Arguments}}
    {{$Argument.InputFieldName}}: var_{{$Argument.InputFieldName}},
  {{end}}
  );
}

{{end}}

{{/* Entry point */}}
void run(CLI cli, List<String> args) {
  var (subcommandPath: subcommandPath, restArgs: restArgs) = _resolveSubcommand(args);
  switch (subcommandPath.join(" ")) {
{{with .Program}}
  case {{.FullPathLiteral}}:
    final funcMethod = cli.{{.CLIFuncMethodChain}};
    if (funcMethod == null) {
      throw Exception("'${ {{.FullPathLiteral}} }' is unsupported: cli.{{.CLIFuncMethodChain}} not assigned");
    }

    {{.CLIInputRecordName}}? input;
    Exception? err;
    try {
      input = _resolve_{{.CLIInputRecordName}}(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);
{{end}}
{{range .Commands}}
  case {{.FullPathLiteral}}:
    final funcMethod = cli.{{.CLIFuncMethodChain}};
    if (funcMethod == null) {
      throw Exception("'${ {{.FullPathLiteral}} }' is unsupported: cli.{{.CLIFuncMethodChain}} not assigned");
    }

    {{.CLIInputRecordName}}? input;
    Exception? err;
    try {
      input = _resolve_{{.CLIInputRecordName}}(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);
{{end}}
  }
}


({List<String> subcommandPath, List<String> restArgs}) _resolveSubcommand(List<String> args) {
  final subcommandSet = {
    {{with .Program}}{{.FullPathLiteral}}: true,{{end}}
    {{range .Commands}}{{.FullPathLiteral}}: true,
    {{end}}
  };

  List<String> subcommandPath = [];
  for (var arg in args) {
    if (arg == "--") {
      break;
    }
    final pathLiteral = ([]..addAll(subcommandPath)..add(arg)).join(" ");
    if (!subcommandSet.containsKey(pathLiteral)) {
      break;
    }

    subcommandPath.add(arg);
  }

  return (subcommandPath: subcommandPath, restArgs: args.sublist(subcommandPath.length));
}

dynamic _parseValue(Type t, List<String> strValue) {
  switch (t) {
  case const (List<bool>):
    return strValue.map((s)=>_parseValue(bool, [s]) as bool).toList();
  case const (List<int>):
    return strValue.map((s)=>_parseValue(int, [s]) as int).toList();
  case const (List<double>):
    return strValue.map((s)=>_parseValue(double, [s]) as double).toList();
  case const (List<String>):
    return strValue.map((s)=>_parseValue(String, [s]) as String).toList();
  case bool when strValue.length == 1:
    return switch(strValue[0]) {
      "1" || "t" || "T" || "true" || "TRUE" || "True" => true,
      "0" || "f" || "F" || "false" || "FALSE" || "False" => false,
      _ => throw Exception("invalid boolean value: ${strValue[0]}"),
    };
  case int when strValue.length == 1:
    return int.parse(strValue[0]);
  case double when strValue.length == 1:
    return double.parse(strValue[0]);
  case String when strValue.length == 1:
    return strValue[0];
  }

  throw Exception("invalid type: ${t}");
}

{{/* Documents */}}
String getDoc(List<String> subcommands) {
  switch (subcommands.join(" ")) {
{{with .Program}}
  case {{.FullPathLiteral}}:
    return {{.DocText}};
{{end}}
{{range .Commands}}
  case {{.FullPathLiteral}}:
    return {{.DocText}};
{{end}}
  default:
    throw Exception("invalid subcommands: ${subcommands}");
  }
}
