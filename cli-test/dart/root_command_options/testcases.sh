#!/bin/sh
# Code generated by go run ./cli-test/internal/gen-testcases/... DO NOT EDIT!

run_test() {
    message=$1
    want=$2
    entrypoint=$3
    shift 3
    got=$( ${entrypoint} $@ )

    if [ "${got}" = "${want}" ]; then
        echo "OK: ${message}"
    else
        echo "NG: ${message}"
        echo "  execution: ${entrypoint} $@"
        echo "  want:      '${want}'"
        echo "  got:       '${got}'"
    fi
}


run_test 'root command should handle optional arguments of type integer' \
    '_123_false_' \
    dart main.dart -opt-integer=123

run_test 'root command should handle optional arguments of type boolean' \
    '_0_true_' \
    dart main.dart -opt-boolean=true

run_test 'root command should handle optional arguments of type boolean' \
    '_0_false_abc' \
    dart main.dart -opt-string=abc

run_test 'root command should handle optional arguments of type integer' \
    '_123_false_' \
    dart main.dart  -i=123

run_test 'root command should handle short optional arguments of type boolean' \
    '_0_true_' \
    dart main.dart -b=true

run_test 'root command should handle short optional arguments of type boolean' \
    '_0_false_abc' \
    dart main.dart -s=abc

run_test 'root command should handle default optional argument values' \
    '_0_false_' \
    dart main.dart 

run_test 'root command should handle optional arguments of type boolean without value' \
    '_0_true_' \
    dart main.dart -opt-boolean

run_test 'root command should handle short optional arguments of type boolean without value' \
    '_0_true_' \
    dart main.dart -b
