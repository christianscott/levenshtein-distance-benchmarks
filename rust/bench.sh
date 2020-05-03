#!/usr/bin/env bash

set -e

run() {
  cargo build --release 2> /dev/null
  echo "rust: $(./target/release/rust)"
}

run;
