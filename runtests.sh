#!/bin/bash

EXIT_CODE=0

for go_dir in $(find . -type f -name '*.go' -printf '%h\n' | sort | uniq); do
  if go build -tags testing -o /tmp/go_build_tmp ${go_dir}; then
    echo "${go_dir} builds!"
  else
    EXIT_CODE=$?
  fi
done

rm /tmp/go_build_tmp || true

for test_dir in $(find . -type f -name '*_test.go' -printf '%h\n' | sort | uniq); do
   if go test -tags testing ${test_dir}; then
    echo "${test_dir} ok!"
  else
    EXIT_CODE=$?
  fi
done

exit ${EXIT_CODE}