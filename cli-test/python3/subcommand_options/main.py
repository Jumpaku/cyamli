import sys

from cli_gen import *


def cli_func(subcommand: list[str], input: CLI_Sub_Input, ex: Exception):
    if ex:
        raise ex
    print(f"{''.join(subcommand)}_{input.opt_opt_integer}_{str(input.opt_opt_boolean).lower()}_{input.opt_opt_string}")


cli = CLI()

cli.sub.FUNC = cli_func

run(cli, sys.argv)
