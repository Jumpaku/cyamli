#!/bin/sh

WORKDIR=$(pwd)
for TESTCASE in */testcases.sh; do
  cd "${WORKDIR}" || exit 1
  GROUP=$(dirname "${TESTCASE}")
  cd "${GROUP}" || exit 1
  if sh testcases.sh | tee test.log | grep -E '^NG'; then
    cat test.log
    exit 1
  fi
done