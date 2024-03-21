import cli_gen
import sys


def func(args, err):
    if err is not None:
        raise err
    print(args)


cli = cli_gen.CLI()
cli.FUNC = func
cli.list.FUNC = func
cli.fetch.FUNC = func

cli_gen.run(cli, sys.argv)
