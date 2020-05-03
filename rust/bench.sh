#!/usr/bin/env bash

set -e

run() {
  cargo build --release 2> /dev/null
  ./target/release/rust
}

run;
