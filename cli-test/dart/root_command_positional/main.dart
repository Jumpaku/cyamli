import 'cli.g.dart';

void main(List<String> args) {
  final cli = CLI();
  cli.FUNC = (subcommand, input, inputErr) {
    if (inputErr != null) {
      throw inputErr;
    }
    print(
        "${subcommand.join("-")}_${input?.argArgInteger}_${input?.argArgBoolean}_${input?.argArgString}_${input?.argArgVariadic.join(",")}");
  };
  run(cli, args);
}
