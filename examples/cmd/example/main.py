from cli_gen import *
import sys

cli = CLI()

cli.FUNC = lambda input, ex: print(input, ex) 
cli.sub1.FUNC = lambda input, ex: print(input, ex) 
cli.sub2.FUNC = lambda input, ex: print(input, ex) 
cli.sub3.FUNC = lambda input, ex: print(input, ex) 
run(cli, sys.argv)