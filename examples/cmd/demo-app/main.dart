import 'cli.g.dart';

void main(List<String> args) {
  final cli = CLI();
  cli.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  cli.fetch.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  cli.list.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  run(cli, args);
}
