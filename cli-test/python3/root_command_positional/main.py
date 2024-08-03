import sys

from cli_gen import *


def cli_func(subcommand: list[str], input: CLI_Input, ex: Exception):
    if ex:
        raise ex
    print(
        f"{''.join(subcommand)}_{input.arg_arg_integer}_{str(input.arg_arg_boolean).lower()}_{input.arg_arg_string}_{','.join(input.arg_arg_variadic)}")


cli = CLI()

cli.FUNC = cli_func

run(cli, sys.argv)
