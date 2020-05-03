#!/usr/bin/env bash

set -e

bench() {
  local run_file_name="$1"
  pushd $(dirname "${run_file_name}") > /dev/null
  bash $(basename "${run_file_name}")
  popd > /dev/null
}

loop() {
  local cmd="$1"
  find . -name "$cmd.sh" | while read run_file_name; do
    $cmd "${run_file_name}"
  done
}

run() {
  local cmd="$1"
  shift
  case "${cmd}" in
    bench) loop bench ;;
    *) error "Unknown cmd '${cmd}'" ;;
  esac
}

run $1;
