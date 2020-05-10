#!/usr/bin/env bash

set -e

run() {
  pypy3 main.py 2>/dev/null
}

run;
