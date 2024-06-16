import 'cli.g.dart';

void main(List<String> args) {
  final cli = CLI();
  cli.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  cli.sub1.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  cli.sub2.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  cli.sub3.FUNC = (subcommand, input, err) {
    print("${subcommand}, ${input}, ${err}");
    print(getDoc(subcommand));
  };
  run(cli, args);
}