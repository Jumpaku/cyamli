from cli_gen import *
import sys

cli = CLI()

cli.FUNC = lambda subcommand, args, ex: print(subcommand, input, ex)
cli.sub1.FUNC = lambda subcommand, args, ex: print(subcommand, input, ex)
cli.sub2.FUNC = lambda subcommand, args, ex: print(subcommand, input, ex)
cli.sub3.FUNC = lambda subcommand, args, ex: print(subcommand, input, ex)
run(cli, sys.argv)
