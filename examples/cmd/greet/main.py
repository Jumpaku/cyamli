import cli_gen
import sys

# Create the CLI object
cli = cli_gen.CLI()

def main():
    # Overwrite behaviors
    cli.FUNC = showHelp
    cli.hello.FUNC = sayHello
    # Run with command line arguments
    cli_gen.run(cli, sys.argv)

def showHelp(input: cli_gen.CLI_Input, inputErr: Exception): 
	if inputErr != None:
		print(cli.desc_simple)
		raise inputErr
	if input.opt_help: 
		print(cli.desc_detail)

def sayHello(input: cli_gen.CLI_Hello_Input, inputErr: Exception):
	if inputErr != None:
		print(cli.hello.desc_simple)
		raise inputErr
	if input.opt_target_name != "":
		print("Hello, {}! My name is {}!".format(input.opt_target_name, input.arg_greeter))
	else:
		print("Hello! My name is {}!".format(input.arg_greeter))

main()