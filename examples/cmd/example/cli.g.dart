// GENERATED CODE - DO NOT MODIFY BY HAND
// generator: github.com/Jumpaku/cyamli v1.1.7

// ignore_for_file: unused_local_variable

typedef Func<Input> = Function(List<String> subcommand, Input? input, Exception? inputErr);





class CLI {
  Func<CLI_Input>? FUNC;
  
  final CLI_Sub1 sub1 = CLI_Sub1();
  
  final CLI_Sub2 sub2 = CLI_Sub2();
  
  final CLI_Sub3 sub3 = CLI_Sub3();
  
}


typedef CLI_Input = ({
  
  String optOptionA,
  
  int optOptionB,
  
  bool optOptionC,
  
  double optOptionD,
  
  String optOptionE,
  
  
  String argArgA,
  
  int argArgB,
  
  bool argArgC,
  
  double argArgD,
  
  String argArgE,
  
  List<String> argArgV,
  
});


CLI_Input _resolve_CLI_Input(List<String> restArgs) {
  
  String var_optOptionA = "abc";
  
  int var_optOptionB = -123;
  
  bool var_optOptionC = true;
  
  double var_optOptionD = -123.456;
  
  String var_optOptionE = "";
  
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
    
    case "-option-a" || "-a":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionA = _parseValue(var_optOptionA.runtimeType, [lit]) as String;
    
    case "-option-b" || "-b":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionB = _parseValue(var_optOptionB.runtimeType, [lit]) as int;
    
    case "-option-c" || "-c":
      if (!cut) {
        lit = "true";
        
      }
      var_optOptionC = _parseValue(var_optOptionC.runtimeType, [lit]) as bool;
    
    case "-option-d" || "-d":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionD = _parseValue(var_optOptionD.runtimeType, [lit]) as double;
    
    case "-option-e":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionE = _parseValue(var_optOptionE.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  
  if (arguments.length <= 0) {
    throw Exception("too few arguments");
  }
  String var_argArgA = _parseValue(String, [arguments[0]]) as String;
  

  

  
  if (arguments.length <= 1) {
    throw Exception("too few arguments");
  }
  int var_argArgB = _parseValue(int, [arguments[1]]) as int;
  

  

  
  if (arguments.length <= 2) {
    throw Exception("too few arguments");
  }
  bool var_argArgC = _parseValue(bool, [arguments[2]]) as bool;
  

  

  
  if (arguments.length <= 3) {
    throw Exception("too few arguments");
  }
  double var_argArgD = _parseValue(double, [arguments[3]]) as double;
  

  

  
  if (arguments.length <= 4) {
    throw Exception("too few arguments");
  }
  String var_argArgE = _parseValue(String, [arguments[4]]) as String;
  

  

  
  if (arguments.length <= 5 - 1) {
    throw Exception("too few arguments");
  }
  List<String> var_argArgV = _parseValue(List<String>, arguments.sublist(5)) as List<String>;
  

  

  return (
  
    optOptionA: var_optOptionA,
  
    optOptionB: var_optOptionB,
  
    optOptionC: var_optOptionC,
  
    optOptionD: var_optOptionD,
  
    optOptionE: var_optOptionE,
  
  
    argArgA: var_argArgA,
  
    argArgB: var_argArgB,
  
    argArgC: var_argArgC,
  
    argArgD: var_argArgD,
  
    argArgE: var_argArgE,
  
    argArgV: var_argArgV,
  
  );
}






class CLI_Sub1 {
  Func<CLI_Sub1_Input>? FUNC;
  
}


typedef CLI_Sub1_Input = ();


CLI_Sub1_Input _resolve_CLI_Sub1_Input(List<String> restArgs) {
  
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



class CLI_Sub2 {
  Func<CLI_Sub2_Input>? FUNC;
  
}


typedef CLI_Sub2_Input = ();


CLI_Sub2_Input _resolve_CLI_Sub2_Input(List<String> restArgs) {
  
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



class CLI_Sub3 {
  Func<CLI_Sub3_Input>? FUNC;
  
  final CLI_Sub3Suba suba = CLI_Sub3Suba();
  
  final CLI_Sub3Subb subb = CLI_Sub3Subb();
  
  final CLI_Sub3Subc subc = CLI_Sub3Subc();
  
  final CLI_Sub3Subd subd = CLI_Sub3Subd();
  
}


typedef CLI_Sub3_Input = ({
  
  String optOptionA,
  
  int optOptionB,
  
  bool optOptionC,
  
  double optOptionD,
  
  String optOptionE,
  
  
  String argArgA,
  
  int argArgB,
  
  bool argArgC,
  
  double argArgD,
  
  String argArgE,
  
  List<String> argArgV,
  
});


CLI_Sub3_Input _resolve_CLI_Sub3_Input(List<String> restArgs) {
  
  String var_optOptionA = "abc";
  
  int var_optOptionB = -123;
  
  bool var_optOptionC = true;
  
  double var_optOptionD = -123.456;
  
  String var_optOptionE = "";
  
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
    
    case "-option-a" || "-a":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionA = _parseValue(var_optOptionA.runtimeType, [lit]) as String;
    
    case "-option-b" || "-b":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionB = _parseValue(var_optOptionB.runtimeType, [lit]) as int;
    
    case "-option-c" || "-c":
      if (!cut) {
        lit = "true";
        
      }
      var_optOptionC = _parseValue(var_optOptionC.runtimeType, [lit]) as bool;
    
    case "-option-d" || "-d":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionD = _parseValue(var_optOptionD.runtimeType, [lit]) as double;
    
    case "-option-e":
      if (!cut) {
         throw Exception("value is not specified to option ${optName}");
        
      }
      var_optOptionE = _parseValue(var_optOptionE.runtimeType, [lit]) as String;
    
    default:
      throw Exception("unknown option ${optName}");
    }
  }

  

  
  if (arguments.length <= 0) {
    throw Exception("too few arguments");
  }
  String var_argArgA = _parseValue(String, [arguments[0]]) as String;
  

  

  
  if (arguments.length <= 1) {
    throw Exception("too few arguments");
  }
  int var_argArgB = _parseValue(int, [arguments[1]]) as int;
  

  

  
  if (arguments.length <= 2) {
    throw Exception("too few arguments");
  }
  bool var_argArgC = _parseValue(bool, [arguments[2]]) as bool;
  

  

  
  if (arguments.length <= 3) {
    throw Exception("too few arguments");
  }
  double var_argArgD = _parseValue(double, [arguments[3]]) as double;
  

  

  
  if (arguments.length <= 4) {
    throw Exception("too few arguments");
  }
  String var_argArgE = _parseValue(String, [arguments[4]]) as String;
  

  

  
  if (arguments.length <= 5 - 1) {
    throw Exception("too few arguments");
  }
  List<String> var_argArgV = _parseValue(List<String>, arguments.sublist(5)) as List<String>;
  

  

  return (
  
    optOptionA: var_optOptionA,
  
    optOptionB: var_optOptionB,
  
    optOptionC: var_optOptionC,
  
    optOptionD: var_optOptionD,
  
    optOptionE: var_optOptionE,
  
  
    argArgA: var_argArgA,
  
    argArgB: var_argArgB,
  
    argArgC: var_argArgC,
  
    argArgD: var_argArgD,
  
    argArgE: var_argArgE,
  
    argArgV: var_argArgV,
  
  );
}



class CLI_Sub3Suba {
  Func<CLI_Sub3Suba_Input>? FUNC;
  
}


typedef CLI_Sub3Suba_Input = ({
  
  
  List<bool> argArgV,
  
});


CLI_Sub3Suba_Input _resolve_CLI_Sub3Suba_Input(List<String> restArgs) {
  
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

  

  
  if (arguments.length <= 0 - 1) {
    throw Exception("too few arguments");
  }
  List<bool> var_argArgV = _parseValue(List<bool>, arguments.sublist(0)) as List<bool>;
  

  

  return (
  
  
    argArgV: var_argArgV,
  
  );
}



class CLI_Sub3Subb {
  Func<CLI_Sub3Subb_Input>? FUNC;
  
}


typedef CLI_Sub3Subb_Input = ({
  
  
  List<int> argArgV,
  
});


CLI_Sub3Subb_Input _resolve_CLI_Sub3Subb_Input(List<String> restArgs) {
  
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

  

  
  if (arguments.length <= 0 - 1) {
    throw Exception("too few arguments");
  }
  List<int> var_argArgV = _parseValue(List<int>, arguments.sublist(0)) as List<int>;
  

  

  return (
  
  
    argArgV: var_argArgV,
  
  );
}



class CLI_Sub3Subc {
  Func<CLI_Sub3Subc_Input>? FUNC;
  
}


typedef CLI_Sub3Subc_Input = ({
  
  
  List<double> argArgV,
  
});


CLI_Sub3Subc_Input _resolve_CLI_Sub3Subc_Input(List<String> restArgs) {
  
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

  

  
  if (arguments.length <= 0 - 1) {
    throw Exception("too few arguments");
  }
  List<double> var_argArgV = _parseValue(List<double>, arguments.sublist(0)) as List<double>;
  

  

  return (
  
  
    argArgV: var_argArgV,
  
  );
}



class CLI_Sub3Subd {
  Func<CLI_Sub3Subd_Input>? FUNC;
  
}


typedef CLI_Sub3Subd_Input = ({
  
  
  List<String> argArgV,
  
});


CLI_Sub3Subd_Input _resolve_CLI_Sub3Subd_Input(List<String> restArgs) {
  
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

  

  
  if (arguments.length <= 0 - 1) {
    throw Exception("too few arguments");
  }
  List<String> var_argArgV = _parseValue(List<String>, arguments.sublist(0)) as List<String>;
  

  

  return (
  
  
    argArgV: var_argArgV,
  
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


  case "sub1":
    final funcMethod = cli.sub1.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub1" }' is unsupported: cli.sub1.FUNC not assigned");
    }

    CLI_Sub1_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub1_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "sub2":
    final funcMethod = cli.sub2.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub2" }' is unsupported: cli.sub2.FUNC not assigned");
    }

    CLI_Sub2_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub2_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "sub3":
    final funcMethod = cli.sub3.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub3" }' is unsupported: cli.sub3.FUNC not assigned");
    }

    CLI_Sub3_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub3_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "sub3 suba":
    final funcMethod = cli.sub3.suba.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub3 suba" }' is unsupported: cli.sub3.suba.FUNC not assigned");
    }

    CLI_Sub3Suba_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub3Suba_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "sub3 subb":
    final funcMethod = cli.sub3.subb.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub3 subb" }' is unsupported: cli.sub3.subb.FUNC not assigned");
    }

    CLI_Sub3Subb_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub3Subb_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "sub3 subc":
    final funcMethod = cli.sub3.subc.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub3 subc" }' is unsupported: cli.sub3.subc.FUNC not assigned");
    }

    CLI_Sub3Subc_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub3Subc_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  case "sub3 subd":
    final funcMethod = cli.sub3.subd.FUNC;
    if (funcMethod == null) {
      throw Exception("'${ "sub3 subd" }' is unsupported: cli.sub3.subd.FUNC not assigned");
    }

    CLI_Sub3Subd_Input? input;
    Exception? err;
    try {
      input = _resolve_CLI_Sub3Subd_Input(restArgs);
    } on Exception catch (e) {
      err = e;
    }
    funcMethod(subcommandPath, input, err);

  }
}


({List<String> subcommandPath, List<String> restArgs}) _resolveSubcommand(List<String> args) {
  final subcommandSet = {
    "": true,
    "sub1": true,
    "sub2": true,
    "sub3": true,
    "sub3 suba": true,
    "sub3 subb": true,
    "sub3 subc": true,
    "sub3 subd": true,
    
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
    return "example (v1.0.0)\n\nexample\n\n    Description:\n        this is an example command\n\n    Syntax:\n        \$ example  [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -option-a=<string>, -a=<string>  (default=\"abc\"):\n            a - this is an option for root command\n\n        -option-b=<integer>, -b=<integer>  (default=-123):\n            b - this is an option for root command\n\n        -option-c[=<boolean>], -c[=<boolean>]  (default=true):\n            c - this is an option for root command\n\n        -option-d=<float>, -d=<float>  (default=-123.456):\n            d - this is an option for root command\n\n        -option-e=<string>  (default=\"\"):\n\n    Arguments:\n        1.  <arg_a:string>\n            a - this is an argument for root command\n\n        2.  <arg_b:integer>\n            b - this is an argument for root command\n\n        3.  <arg_c:boolean>\n            c - this is an argument for root command\n\n        4.  <arg_d:float>\n            d - this is an argument for root command\n\n        5.  <arg_e:string>\n\n        6. [<arg_v:string>]...\n            v - this is an argument for root command\n\n    Subcommands:\n        sub1:\n            1 - this is a sub command\n\n        sub2:\n            2 - this is a sub command\n\n        sub3:\n            3 - this is a sub command\n\n\n";


  case "sub1":
    return "example (v1.0.0)\n\nexample sub1\n\n    Description:\n        1 - this is a sub command\n\n    Syntax:\n        \$ example sub1\n\n\n";

  case "sub2":
    return "example (v1.0.0)\n\nexample sub2\n\n    Description:\n        2 - this is a sub command\n\n    Syntax:\n        \$ example sub2\n\n\n";

  case "sub3":
    return "example (v1.0.0)\n\nexample sub3\n\n    Description:\n        3 - this is a sub command\n\n    Syntax:\n        \$ example sub3 [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -option-a=<string>, -a=<string>  (default=\"abc\"):\n            3 - a - this is an option for root command\n\n        -option-b=<integer>, -b=<integer>  (default=-123):\n            3 - b - this is an option for root command\n\n        -option-c[=<boolean>], -c[=<boolean>]  (default=true):\n            3 - c - this is an option for root command\n\n        -option-d=<float>, -d=<float>  (default=-123.456):\n            3 - d - this is an option for root command\n\n        -option-e=<string>  (default=\"\"):\n\n    Arguments:\n        1.  <arg_a:string>\n            3 - a - this is an argument for root command\n\n        2.  <arg_b:integer>\n            3 - b - this is an argument for root command\n\n        3.  <arg_c:boolean>\n            3 - c - this is an argument for root command\n\n        4.  <arg_d:float>\n            3 - d - this is an argument for root command\n\n        5.  <arg_e:string>\n\n        6. [<arg_v:string>]...\n            3 - v - this is an argument for root command\n\n    Subcommands:\n        suba:\n\n        subb:\n\n        subc:\n\n        subd:\n\n\n";

  case "sub3 suba":
    return "example (v1.0.0)\n\nexample sub3 suba\n\n    Syntax:\n        \$ example sub3 suba [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:boolean>]...\n\n\n";

  case "sub3 subb":
    return "example (v1.0.0)\n\nexample sub3 subb\n\n    Syntax:\n        \$ example sub3 subb [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:integer>]...\n\n\n";

  case "sub3 subc":
    return "example (v1.0.0)\n\nexample sub3 subc\n\n    Syntax:\n        \$ example sub3 subc [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:float>]...\n\n\n";

  case "sub3 subd":
    return "example (v1.0.0)\n\nexample sub3 subd\n\n    Syntax:\n        \$ example sub3 subd [<argument>]... [-- [<argument>]...]\n\n    Arguments:\n        1. [<arg_v:string>]...\n\n\n";

  default:
    throw Exception("invalid subcommands: ${subcommands}");
  }
}
