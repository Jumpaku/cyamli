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


run_test 'subcommand should handle optional arguments of type integer' \
    'sub_123_false_' \
    dart main.dart sub -opt-integer=123

run_test 'subcommand should handle optional arguments of type boolean' \
    'sub_0_true_' \
    dart main.dart sub -opt-boolean=true

run_test 'subcommand should handle optional arguments of type boolean' \
    'sub_0_false_abc' \
    dart main.dart sub -opt-string=abc

run_test 'subcommand should handle optional arguments of type integer' \
    'sub_123_false_' \
    dart main.dart sub -i=123

run_test 'subcommand should handle short optional arguments of type boolean' \
    'sub_0_true_' \
    dart main.dart sub -b=true

run_test 'subcommand should handle short optional arguments of type boolean' \
    'sub_0_false_abc' \
    dart main.dart sub -s=abc

run_test 'subcommand should handle default optional argument values' \
    'sub_0_false_' \
    dart main.dart sub

run_test 'subcommand should handle optional arguments of type boolean without value' \
    'sub_0_true_' \
    dart main.dart sub -opt-boolean

run_test 'subcommand should handle short optional arguments of type boolean without value' \
    'sub_0_true_' \
    dart main.dart sub -b
