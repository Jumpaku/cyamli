import cli_gen
import sys


def func(subcommand: list[str], args: any, err: Exception):
    if err is not None:
        raise err
    print(args)
    print(cli_gen.get_doc(subcommand))


cli = cli_gen.CLI()
cli.FUNC = func
cli.list.FUNC = func
cli.fetch.FUNC = func

cli_gen.run(cli, sys.argv)
