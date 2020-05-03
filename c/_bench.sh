#!/usr/bin/env bash

set -e

run() {
  rm main
  gcc -O3 main.c -o main
  echo "c: $(./main)"
}

run;
