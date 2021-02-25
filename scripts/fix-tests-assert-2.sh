#!/bin/bash

find . -type f -name \*_test.go | xargs -n1 -IX \
  bash -c "sed -re 's/assert\.Equal\(t, (true|false), ([^\*][a-zA-Z]+(\.[a-zA-Z0-9]+)*)\)/assert\.Equal\(t, Bool\(\1\), \2\)/' X > X.tmp && mv X{.tmp,}"

find . -type f -name \*_test.go | xargs -n1 -IX \
  bash -c "sed -re 's/assert\.Equal\(t, (\".+\"), ([^\*].+(\..+)*)\)/assert\.Equal\(t, \1, \*\2\)/' X > X.tmp && mv X{.tmp,}"

rm **/*.tmp
