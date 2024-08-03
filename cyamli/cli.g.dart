// GENERATED CODE - DO NOT MODIFY BY HAND
// generator: github.com/Jumpaku/cyamli v1.1.5

// ignore_for_file: unused_local_variable

typedef Func<Input> = Function(List<String> subcommand, Input? input, Exception? inputErr);





class CLI {
  Func<CLI_Input>? FUNC;
  
  final CLI_Generate generate = CLI_Generate();
  
  final CLI_List list = CLI_List();
  
  final CLI_Validate validate = CLI_Validate();
  
}


typedef CLI_Input = ({
  
  bool optHelp,
  
  bool optVersion,
  
  
});


CLI_Input _resolve_CLI_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
  bool var_optVersion = false;
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-version" || "-v":
      if (!cut) {
        lit = "true";
        
      }
      var_optVersion = _parseValue(var_optVersion.runtimeType, [lit]) as bool;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
    optVersion: var_optVersion,
  
  
  );
}






class CLI_Generate {
  Func<CLI_Generate_Input>? FUNC;
  
  final CLI_GenerateDart dart = CLI_GenerateDart();
  
  final CLI_GenerateDocs docs = CLI_GenerateDocs();
  
  final CLI_GenerateGolang golang = CLI_GenerateGolang();
  
  final CLI_GeneratePython3 python3 = CLI_GeneratePython3();
  
}


typedef CLI_Generate_Input = ({
  
  bool optHelp,
  
  
});


CLI_Generate_Input _resolve_CLI_Generate_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
  
  );
}



class CLI_GenerateDart {
  Func<CLI_GenerateDart_Input>? FUNC;
  
}


typedef CLI_GenerateDart_Input = ({
  
  bool optHelp,
  
  String optOutPath,
  
  String optSchemaPath,
  
  
});


CLI_GenerateDart_Input _resolve_CLI_GenerateDart_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
  String var_optOutPath = "";
  
  String var_optSchemaPath = "";
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-out-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOutPath = _parseValue(var_optOutPath.runtimeType, [lit]) as String;
    
    case "-schema-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optSchemaPath = _parseValue(var_optSchemaPath.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
    optOutPath: var_optOutPath,
  
    optSchemaPath: var_optSchemaPath,
  
  
  );
}



class CLI_GenerateDocs {
  Func<CLI_GenerateDocs_Input>? FUNC;
  
}


typedef CLI_GenerateDocs_Input = ({
  
  bool optAll,
  
  String optFormat,
  
  bool optHelp,
  
  String optOutPath,
  
  String optSchemaPath,
  
  
  List<String> argSubcommands,
  
});


CLI_GenerateDocs_Input _resolve_CLI_GenerateDocs_Input(List<String> restArgs) {
  
  bool var_optAll = false;
  
  String var_optFormat = "text";
  
  bool var_optHelp = false;
  
  String var_optOutPath = "";
  
  String var_optSchemaPath = "";
  
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
    
    case "-all" || "-a":
      if (!cut) {
        lit = "true";
        
      }
      var_optAll = _parseValue(var_optAll.runtimeType, [lit]) as bool;
    
    case "-format" || "-f":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optFormat = _parseValue(var_optFormat.runtimeType, [lit]) as String;
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-out-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOutPath = _parseValue(var_optOutPath.runtimeType, [lit]) as String;
    
    case "-schema-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optSchemaPath = _parseValue(var_optSchemaPath.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  
  if (arguments.length <= 0 - 1) {
    throw Exception("too few arguments");
  }
  List<String> var_argSubcommands = _parseValue(List<String>, arguments.sublist(0)) as List<String>;
  

  

  return (
  
    optAll: var_optAll,
  
    optFormat: var_optFormat,
  
    optHelp: var_optHelp,
  
    optOutPath: var_optOutPath,
  
    optSchemaPath: var_optSchemaPath,
  
  
    argSubcommands: var_argSubcommands,
  
  );
}



class CLI_GenerateGolang {
  Func<CLI_GenerateGolang_Input>? FUNC;
  
}


typedef CLI_GenerateGolang_Input = ({
  
  bool optHelp,
  
  String optOutPath,
  
  String optPackage,
  
  String optSchemaPath,
  
  
});


CLI_GenerateGolang_Input _resolve_CLI_GenerateGolang_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
  String var_optOutPath = "";
  
  String var_optPackage = "main";
  
  String var_optSchemaPath = "";
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-out-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOutPath = _parseValue(var_optOutPath.runtimeType, [lit]) as String;
    
    case "-package":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optPackage = _parseValue(var_optPackage.runtimeType, [lit]) as String;
    
    case "-schema-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optSchemaPath = _parseValue(var_optSchemaPath.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
    optOutPath: var_optOutPath,
  
    optPackage: var_optPackage,
  
    optSchemaPath: var_optSchemaPath,
  
  
  );
}



class CLI_GeneratePython3 {
  Func<CLI_GeneratePython3_Input>? FUNC;
  
}


typedef CLI_GeneratePython3_Input = ({
  
  bool optHelp,
  
  String optOutPath,
  
  String optSchemaPath,
  
  
});


CLI_GeneratePython3_Input _resolve_CLI_GeneratePython3_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
  String var_optOutPath = "";
  
  String var_optSchemaPath = "";
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-out-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOutPath = _parseValue(var_optOutPath.runtimeType, [lit]) as String;
    
    case "-schema-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optSchemaPath = _parseValue(var_optSchemaPath.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
    optOutPath: var_optOutPath,
  
    optSchemaPath: var_optSchemaPath,
  
  
  );
}



class CLI_List {
  Func<CLI_List_Input>? FUNC;
  
}


typedef CLI_List_Input = ({
  
  bool optHelp,
  
  String optSchemaPath,
  
  
});


CLI_List_Input _resolve_CLI_List_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
  String var_optSchemaPath = "";
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-schema-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optSchemaPath = _parseValue(var_optSchemaPath.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
    optSchemaPath: var_optSchemaPath,
  
  
  );
}



class CLI_Validate {
  Func<CLI_Validate_Input>? FUNC;
  
}


typedef CLI_Validate_Input = ({
  
  bool optHelp,
  
  String optSchemaPath,
  
  
});


CLI_Validate_Input _resolve_CLI_Validate_Input(List<String> restArgs) {
  
  bool var_optHelp = false;
  
  String var_optSchemaPath = "";
  
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
    
    case "-help" || "-h":
      if (!cut) {
        lit = "true";
        
      }
      var_optHelp = _parseValue(var_optHelp.runtimeType, [lit]) as bool;
    
    case "-schema-path":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optSchemaPath = _parseValue(var_optSchemaPath.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  return (
  
    optHelp: var_optHelp,
  
    optSchemaPath: var_optSchemaPath,
  
  
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


  case "generate":
    final funcMethod = cli.generate.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "generate" }' is unsupported: cli.generate.FUNC not assigned");
    }

    CLI_Generate_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Generate_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "generate dart":
    final funcMethod = cli.generate.dart.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "generate dart" }' is unsupported: cli.generate.dart.FUNC not assigned");
    }

    CLI_GenerateDart_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_GenerateDart_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "generate docs":
    final funcMethod = cli.generate.docs.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "generate docs" }' is unsupported: cli.generate.docs.FUNC not assigned");
    }

    CLI_GenerateDocs_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_GenerateDocs_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "generate golang":
    final funcMethod = cli.generate.golang.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "generate golang" }' is unsupported: cli.generate.golang.FUNC not assigned");
    }

    CLI_GenerateGolang_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_GenerateGolang_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "generate python3":
    final funcMethod = cli.generate.python3.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "generate python3" }' is unsupported: cli.generate.python3.FUNC not assigned");
    }

    CLI_GeneratePython3_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_GeneratePython3_Input(restArgs);
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

  case "validate":
    final funcMethod = cli.validate.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "validate" }' is unsupported: cli.validate.FUNC not assigned");
    }

    CLI_Validate_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Validate_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  }
}


({List<String> subcommandPath, List<String> restArgs}) _resolveSubcommand(List<String> args) {
  final subcommandSet = {
    "": true,
    "generate": true,
    "generate dart": true,
    "generate docs": true,
    "generate golang": true,
    "generate python3": true,
    "list": true,
    "validate": true,
    
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
    return "cyamli (v1.1.5)\n\ncyamli\n\n    Description:\n        A command line tool to generate CLI for your app from YAML-based schema.\n\n    Syntax:\n        \$ cyamli  [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of this app.\n\n        -version[=<boolean>], -v[=<boolean>]  (default=false):\n            shows version of this app.\n\n    Subcommands:\n        generate:\n            holds subcommands to generate CLI code.\n\n        list:\n            shows subcommands\n\n        validate:\n            validates CLI schema.\n\n\n";


  case "generate":
    return "cyamli (v1.1.5)\n\ncyamli generate\n\n    Description:\n        holds subcommands to generate CLI code.\n\n    Syntax:\n        \$ cyamli generate [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of generate subcommand.\n\n    Subcommands:\n        dart:\n            generates CLI for your app written in Dart.\n\n        docs:\n            generates documentation for your CLI app.\n\n        golang:\n            generates CLI for your app written in Go.\n\n        python3:\n            generates CLI for your app written in Python3.\n\n\n";

  case "generate dart":
    return "cyamli (v1.1.5)\n\ncyamli generate dart\n\n    Description:\n        generates CLI for your app written in Dart.\n\n    Syntax:\n        \$ cyamli generate dart [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of dart subcommand.\n\n        -out-path=<string>  (default=\"\"):\n            if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.\n\n        -schema-path=<string>  (default=\"\"):\n            if specified then reads schema file from the path, otherwise reads from stdin.\n\n\n";

  case "generate docs":
    return "cyamli (v1.1.5)\n\ncyamli generate docs\n\n    Description:\n        generates documentation for your CLI app.\n\n    Syntax:\n        \$ cyamli generate docs [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -all[=<boolean>], -a[=<boolean>]  (default=false):\n            if specified then outputs documentation for all subcommands, otherwise in text format.\n\n        -format=<string>, -f=<string>  (default=\"text\"):\n            specifies output format of the documentation in text or markdown.\n\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of docs subcommand.\n\n        -out-path=<string>  (default=\"\"):\n            if specified then creates a file at the path and writes generated documentation, otherwise outputs to stdout.\n\n        -schema-path=<string>  (default=\"\"):\n            if specified then reads schema file from the path, otherwise reads from stdin.\n\n    Arguments:\n        1. [<subcommands:string>]...\n            selects subcommand for which the documentation is output.\n\n\n";

  case "generate golang":
    return "cyamli (v1.1.5)\n\ncyamli generate golang\n\n    Description:\n        generates CLI for your app written in Go.\n\n    Syntax:\n        \$ cyamli generate golang [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of golang subcommand.\n\n        -out-path=<string>  (default=\"\"):\n            if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.\n\n        -package=<string>  (default=\"main\"):\n            package name where the generated file will be placed.\n\n        -schema-path=<string>  (default=\"\"):\n            if specified then reads schema file from the path, otherwise reads from stdin.\n\n\n";

  case "generate python3":
    return "cyamli (v1.1.5)\n\ncyamli generate python3\n\n    Description:\n        generates CLI for your app written in Python3.\n\n    Syntax:\n        \$ cyamli generate python3 [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of python3 subcommand.\n\n        -out-path=<string>  (default=\"\"):\n            if specified then creates a file at the path and writes generated code, otherwise outputs to stdout.\n\n        -schema-path=<string>  (default=\"\"):\n            if specified then reads schema file from the path, otherwise reads from stdin.\n\n\n";

  case "list":
    return "cyamli (v1.1.5)\n\ncyamli list\n\n    Description:\n        shows subcommands\n\n    Syntax:\n        \$ cyamli list [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of list subcommand.\n\n        -schema-path=<string>  (default=\"\"):\n            if specified then reads schema file from the path, otherwise reads from stdin.\n\n\n";

  case "validate":
    return "cyamli (v1.1.5)\n\ncyamli validate\n\n    Description:\n        validates CLI schema.\n\n    Syntax:\n        \$ cyamli validate [<option>]...\n\n    Options:\n        -help[=<boolean>], -h[=<boolean>]  (default=false):\n            shows description of validates subcommand.\n\n        -schema-path=<string>  (default=\"\"):\n            if specified then reads schema file from the path, otherwise reads from stdin.\n\n\n";

  default:
    throw Exception("invalid subcommands: ${subcommands}");
  }
}
