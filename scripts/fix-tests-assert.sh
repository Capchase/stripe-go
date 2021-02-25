#!/bin/bash

find . -type f -name \*.go | xargs -n1 -IX \
  bash -c "sed -re 's/assert\.Equal\(t, \"(.+)\", ([^\*][a-zA-Z]+(\.[a-zA-Z0-9]+)*)\)/assert\.Equal\(t, \"\1\", \*\2\)/' X > X.tmp && mv X{.tmp,}"
