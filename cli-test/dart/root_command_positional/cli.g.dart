// GENERATED CODE - DO NOT MODIFY BY HAND
// generator: github.com/Jumpaku/cyamli v1.1.5

// ignore_for_file: unused_local_variable

typedef Func<Input> = Function(List<String> subcommand, Input? input, Exception? inputErr);





class CLI {
  Func<CLI_Input>? FUNC;
  
}


typedef CLI_Input = ({
  
  
  int argArgInteger,
  
  bool argArgBoolean,
  
  String argArgString,
  
  List<String> argArgVariadic,
  
});


CLI_Input _resolve_CLI_Input(List<String> restArgs) {
  
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
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  
  if (arguments.length <= 0) {
    throw Exception("too few arguments");
  }
  int var_argArgInteger = _parseValue(int, [arguments[0]]) as int;
  

  

  
  if (arguments.length <= 1) {
    throw Exception("too few arguments");
  }
  bool var_argArgBoolean = _parseValue(bool, [arguments[1]]) as bool;
  

  

  
  if (arguments.length <= 2) {
    throw Exception("too few arguments");
  }
  String var_argArgString = _parseValue(String, [arguments[2]]) as String;
  

  

  
  if (arguments.length <= 3 - 1) {
    throw Exception("too few arguments");
  }
  List<String> var_argArgVariadic = _parseValue(List<String>, arguments.sublist(3)) as List<String>;
  

  

  return (
  
  
    argArgInteger: var_argArgInteger,
  
    argArgBoolean: var_argArgBoolean,
  
    argArgString: var_argArgString,
  
    argArgVariadic: var_argArgVariadic,
  
  );
}







void run(CLI cli, List<String> args) {
  var (subcommandPath: subcommandPath, restArgs: restArgs) = _resolveSubcommand(args);
  switch (subcommandPath.join(" ")) {

  case "":
    final funcMethod = cli.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "" }' is unsupported: cli.FUNC not assigned");
    }

    CLI_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);


  }
}


({List<String> subcommandPath, List<String> restArgs}) _resolveSubcommand(List<String> args) {
  final subcommandSet = {
    "": true,
    
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


String getDoc(List<String> subcommands) {
  switch (subcommands.join(" ")) {

  case "":
    return "<program>\n\n<program>\n\n    Syntax:\n        \$ <program>  [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1.  <arg_integer:integer>\n\n        2.  <arg_boolean:boolean>\n\n        3.  <arg_string:string>\n\n        4. [<arg_variadic:string>]...\n\n\n";


  default:
    throw Exception("invalid subcommands: ${subcommands}");
  }
}
