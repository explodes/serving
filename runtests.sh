#!/bin/bash

EXIT_CODE=0

for test_dir in $(find . -type f -name '*_test.go' -printf '%h\n' | sort | uniq); do
   if go test -tags testing ${test_dir}; then
    echo "${test_dir} ok!"
  else
    EXIT_CODE=$?
  fi
done

exit ${EXIT_CODE}