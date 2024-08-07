# Code generated by github.com/Jumpaku/cyamli v1.1.7, DO NOT EDIT.

from dataclasses import dataclass
import typing


@dataclass
class CLI_List_Input:
    opt_config: str = ""
    
    
    pass


FuncType_CLI_List = typing.Callable[[None,list[str],CLI_List_Input,Exception],None]
class CLI_List:
    
    FUNC: FuncType_CLI_List = None


def resolve_CLI_List_Input(rest_args: list[str])->CLI_List_Input:
    input = CLI_List_Input()
    arguments = []
    for i, arg in enumerate(rest_args):
        if arg == "--":
            arguments += rest_args[i+1:]
            break
        if not arg.startswith("-"):
            arguments.append(arg)
            continue
        split = arg.split("=", 1)
        opt_name, assign = split[0], len(split) > 1
        
        if opt_name == "-config" or opt_name == "-c":
            if not assign:
                raise Exception("value is not specified to option "+ opt_name)
                
            input.opt_config = parse_value(str, split[1])
            continue
        
        raise Exception("unsupported option " + opt_name)
    
    return input


@dataclass
class CLI_Fetch_Input:
    opt_config: str = ""
    opt_verbose: bool = False
    
    arg_tables: tuple[str,...] = tuple[str,...]()
    
    pass


FuncType_CLI_Fetch = typing.Callable[[None,list[str],CLI_Fetch_Input,Exception],None]
class CLI_Fetch:
    
    FUNC: FuncType_CLI_Fetch = None


def resolve_CLI_Fetch_Input(rest_args: list[str])->CLI_Fetch_Input:
    input = CLI_Fetch_Input()
    arguments = []
    for i, arg in enumerate(rest_args):
        if arg == "--":
            arguments += rest_args[i+1:]
            break
        if not arg.startswith("-"):
            arguments.append(arg)
            continue
        split = arg.split("=", 1)
        opt_name, assign = split[0], len(split) > 1
        
        if opt_name == "-config" or opt_name == "-c":
            if not assign:
                raise Exception("value is not specified to option "+ opt_name)
                
            input.opt_config = parse_value(str, split[1])
            continue
        
        if opt_name == "-verbose" or opt_name == "-v":
            if not assign:
                split.append("True")
                
            input.opt_verbose = parse_value(bool, split[1])
            continue
        
        raise Exception("unsupported option " + opt_name)
    if len(arguments) <= 0 - 1:
        raise Exception("too few arguments")
    input.arg_tables = parse_value(tuple[str,...], *arguments[0:])
    
    return input



@dataclass
class CLI_Input:
    
    
    pass


FuncType_CLI = typing.Callable[[None,list[str],CLI_Input,Exception],None]
class CLI:
    list: CLI_List = CLI_List()
    fetch: CLI_Fetch = CLI_Fetch()
    
    FUNC: FuncType_CLI = None


def resolve_CLI_Input(rest_args: list[str])->CLI_Input:
    input = CLI_Input()
    arguments = []
    for i, arg in enumerate(rest_args):
        if arg == "--":
            arguments += rest_args[i+1:]
            break
        if not arg.startswith("-"):
            arguments.append(arg)
            continue
        split = arg.split("=", 1)
        opt_name, assign = split[0], len(split) > 1
        
        raise Exception("unsupported option " + opt_name)
    
    return input


def run(cli: CLI, args: list[str]):
    r = resolve_subcommand(args)
    subcommand_path, rest_args = r.subcommand, r.rest_args
    joined_subcommand = " ".join(subcommand_path)
    
    if joined_subcommand == "":
        if not cli.FUNC:
            raise Exception("unsupported subcommand \"" + "" + "\": cli.FUNC not assigned")
        ex: Exception = None
        input: CLI_Input = None
        try:
            input = resolve_CLI_Input(rest_args)
        except Exception as e:
            ex = e
        cli.FUNC(subcommand_path, input, ex)
        return
    
    
    if joined_subcommand == "list":
        if not cli.list.FUNC:
            raise Exception("unsupported subcommand \"" + "list" + "\": cli.list.FUNC not assigned")
        ex: Exception = None
        input: CLI_List_Input = None
        try:
            input = resolve_CLI_List_Input(rest_args)
        except Exception as e:
            ex = e
        cli.list.FUNC(subcommand_path, input, ex)
        return
    
    if joined_subcommand == "fetch":
        if not cli.fetch.FUNC:
            raise Exception("unsupported subcommand \"" + "fetch" + "\": cli.fetch.FUNC not assigned")
        ex: Exception = None
        input: CLI_Fetch_Input = None
        try:
            input = resolve_CLI_Fetch_Input(rest_args)
        except Exception as e:
            ex = e
        cli.fetch.FUNC(subcommand_path, input, ex)
        return
    
    raise Exception("subcommand not found: " + joined_subcommand)

@dataclass
class ResolveSubcommandResult:
    subcommand: list[str]
    rest_args: list[str]


def resolve_subcommand(args: list[str])->ResolveSubcommandResult:
    if not args:
        raise Exception("command line arguments are too few")
    
    subcommand_set = {
        "",
        "list","fetch",
    }

    subcommand_path = []
    for arg in args[1:]:
        if arg == "--":
            break
        if " ".join(subcommand_path + [arg]) not in subcommand_set:
            break
        subcommand_path.append(arg)
    
    return ResolveSubcommandResult(subcommand_path, args[1+len(subcommand_path):])


def parse_value(typ, *str_values: str) -> typing.Union[str, bool, float, int, tuple[str, ...], tuple[bool, ...], tuple[float, ...], tuple[int, ...]]:
    try: 
        if typ == str:
            return str(str_values[0])
        if typ == bool:
            if str_values[0] in {"", "0", "f", "F", "FALSE", "false", "False"}:
                return False
            if str_values[0] in {"1", "t", "T", "TRUE", "true", "True"}:
                return True
            raise Exception("could not convert string to bool: '" + str_values[0] + "'")
        if typ == float:
            return float(str_values[0])
        if typ == int:
            return int(str_values[0], base=0)
        if typ == tuple[str,...]:
            return tuple([parse_value(str, s) for s in str_values])
        if typ == tuple[bool,...]:
            return tuple([parse_value(bool, s) for s in str_values])
        if typ == tuple[float,...]:
            return tuple([parse_value(float, s) for s in str_values])
        if typ == tuple[int,...]:
            return tuple([parse_value(int, s) for s in str_values])
        raise Exception("unsupported type")
    except Exception as e:
        e.add_note('fail to parse string value as ' + typ.__name__)
        raise


def get_doc(subcommand: list[str]) -> str:
    joined_subcommand = " ".join(subcommand)
    
    if joined_subcommand == "":
        return "demo\n\ndemo\n\n    Description:\n        demo app to get table information from databases\n\n    Syntax:\n        $ demo \n\n    Subcommands:\n        fetch:\n            show information of tables\n\n        list:\n            list tables\n\n\n"
    
    
    if joined_subcommand == "list":
        return "demo\n\ndemo list\n\n    Description:\n        list tables\n\n    Syntax:\n        $ demo list [<option>]...\n\n    Options:\n        -config=<string>, -c=<string>  (default=\"\"):\n            path to config file\n\n\n"
    
    if joined_subcommand == "fetch":
        return "demo\n\ndemo fetch\n\n    Description:\n        show information of tables\n\n    Syntax:\n        $ demo fetch [<option>|<argument>]... [-- [<argument>]...]\n\n    Options:\n        -config=<string>, -c=<string>  (default=\"\"):\n            path to config file\n\n        -verbose[=<boolean>], -v[=<boolean>]  (default=false):\n            shows detailed log\n\n    Arguments:\n        1. [<tables:string>]...\n            names of tables to be described\n\n\n"
    
    raise Exception("subcommand not found: " + joined_subcommand)
