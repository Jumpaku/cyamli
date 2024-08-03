// GENERATED CODE - DO NOT MODIFY BY HAND
// generator: github.com/Jumpaku/cyamli v1.1.5

// ignore_for_file: unused_local_variable

typedef Func<Input> = Function(List<String> subcommand, Input? input, Exception? inputErr);





class CLI {
  Func<CLI_Input>? FUNC;
  
}


typedef CLI_Input = ({
  
  bool optOptBoolean,
  
  int optOptInteger,
  
  String optOptString,
  
  
});


CLI_Input _resolve_CLI_Input(List<String> restArgs) {
  
  bool var_optOptBoolean = false;
  
  int var_optOptInteger = 0;
  
  String var_optOptString = "";
  
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
    
    case "-opt-boolean" || "-b":
      if (!cut) {
        lit = "true";
        
      }
      var_optOptBoolean = _parseValue(var_optOptBoolean.runtimeType, [lit]) as bool;
    
    case "-opt-integer" || "-i":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptInteger = _parseValue(var_optOptInteger.runtimeType, [lit]) as int;
    
    case "-opt-string" || "-s":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptString = _parseValue(var_optOptString.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optOptBoolean: var_optOptBoolean,
  
    optOptInteger: var_optOptInteger,
  
    optOptString: var_optOptString,
  
  
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
    return "<program>\n\n<program>\n\n    Syntax:\n        \$ <program>  [<option>]...\n\n    Options:\n        -opt-boolean[=<boolean>], -b[=<boolean>]  (default=false):\n\n        -opt-integer=<integer>, -i=<integer>  (default=0):\n\n        -opt-string=<string>, -s=<string>  (default=\"\"):\n\n\n";


  default:
    throw Exception("invalid subcommands: ${subcommands}");
  }
}
