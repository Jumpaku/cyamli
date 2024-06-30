// GENERATED CODE - DO NOT MODIFY BY HAND
// generator: github.com/Jumpaku/cyamli v1.1.5

// ignore_for_file: unused_local_variable

typedef Func<Input> = Function(List<String> subcommand, Input? input, Exception? inputErr);





class CLI {
  Func<CLI_Input>? FUNC;
  
  final CLI_Fetch fetch = CLI_Fetch();
  
  final CLI_List list = CLI_List();
  
}


typedef CLI_Input = ();


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

  

  return (
  
  
  );
}






class CLI_Fetch {
  Func<CLI_Fetch_Input>? FUNC;
  
}


typedef CLI_Fetch_Input = ({
  
  String optConfig,
  
  bool optVerbose,
  
  
  List<String> argTables,
  
});


CLI_Fetch_Input _resolve_CLI_Fetch_Input(List<String> restArgs) {
  
  String var_optConfig = "";
  
  bool var_optVerbose = false;
  
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
    
    case "-config" || "-c":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optConfig = _parseValue(var_optConfig.runtimeType, [lit]) as String;
    
    case "-verbose" || "-v":
      if (!cut) {
        lit = "true";
        
      }
      var_optVerbose = _parseValue(var_optVerbose.runtimeType, [lit]) as bool;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  
  if (arguments.length <= 0 - 1) {
    throw Exception("too few arguments");
  }
  List<String> var_argTables = _parseValue(List<String>, arguments.sublist(0)) as List<String>;
  

  

  return (
  
    optConfig: var_optConfig,
  
    optVerbose: var_optVerbose,
  
  
    argTables: var_argTables,
  
  );
}



class CLI_List {
  Func<CLI_List_Input>? FUNC;
  
}


typedef CLI_List_Input = ({
  
  String optConfig,
  
  
});


CLI_List_Input _resolve_CLI_List_Input(List<String> restArgs) {
  
  String var_optConfig = "";
  
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
    
    case "-config" || "-c":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optConfig = _parseValue(var_optConfig.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optConfig: var_optConfig,
  
  
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


  case "fetch":
    final funcMethod = cli.fetch.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "fetch" }' is unsupported: cli.fetch.FUNC not assigned");
    }

    CLI_Fetch_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Fetch_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "list":
    final funcMethod = cli.list.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "list" }' is unsupported: cli.list.FUNC not assigned");
    }

    CLI_List_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_List_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  }
}


({List<String> subcommandPath, List<String> restArgs}) _resolveSubcommand(List<String> args) {
  final subcommandSet = {
    "": true,
    "fetch": true,
    "list": true,
    
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
    return "demo\n\ndemo\n\n    Description:\n        demo app to get table information from databases\n\n    Syntax:\n        \$ demo \n\n    Subcommands:\n        fetch:\n            show information of tables\n\n        list:\n            list tables\n\n\n";


  case "fetch":
    return "demo\n\ndemo fetch\n\n    Description:\n        show information of tables\n\n    Syntax:\n        \$ demo fetch [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -config=<string>, -c=<string>  (default=\"\"):\n            path to config file\n\n        -verbose[=<boolean>], -v[=<boolean>]  (default=false):\n            shows detailed log\n\n    Arguments:\n        1. [<tables:string>]...\n            names of tables to be described\n\n\n";

  case "list":
    return "demo\n\ndemo list\n\n    Description:\n        list tables\n\n    Syntax:\n        \$ demo list [<option>]...\n\n    Options:\n        -config=<string>, -c=<string>  (default=\"\"):\n            path to config file\n\n\n";

  default:
    throw Exception("invalid subcommands: ${subcommands}");
  }
}
