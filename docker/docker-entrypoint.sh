#!/bin/sh

set -eux

if [ "$1" = 'cyamli' ]; then
    exec "$@"
else
    exec cyamli "$@"
fi