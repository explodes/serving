#!/bin/bash

EXIT_CODE=0

for go_dir in $(find . -type f -name '*.go' -printf '%h\n' | sort | uniq); do
  echo "building ${go_dir}..."
  if go build -tags testing -o /tmp/go_build_tmp ${go_dir}; then
    echo "${go_dir} builds!" > /dev/null
  else
    EXIT_CODE=$?
  fi
done

rm /tmp/go_build_tmp || true

for test_dir in $(find . -type f -name '*_test.go' -printf '%h\n' | sort | uniq); do
  echo "testing ${test_dir}..."
   if go test -timeout 10s -tags testing ${test_dir}; then
    echo "${test_dir} ok!" > /dev/null
  else
    EXIT_CODE=$?
  fi
done

exit ${EXIT_CODE}