import 'cli.g.dart';

void main(List<String> args) {
  final cli = CLI();
  cli.sub.FUNC = (subcommand, input, inputErr) {
    if (inputErr != null) {
      throw inputErr;
    }
    print(
        "${subcommand.join("-")}_${input?.optOptInteger}_${input?.optOptBoolean}_${input?.optOptString}");
  };
  run(cli, args);
}
