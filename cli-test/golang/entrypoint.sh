#!/bin/sh
set -eu

WORKDIR=$(pwd)
for TESTCASE in */testcases.sh; do
  cd "${WORKDIR}"
  GROUP=$(dirname "${TESTCASE}")
  cd "${GROUP}"
  sh testcases.sh
done