#!/bin/bash

find . -type f -name \*.go | xargs -n1 -IX \
  bash -c "sed -re 's/([a-zA-Z]+)\{([a-zA-Z]+): \"(.+)\"\}/\1\{\2: String\(\"\3\"\)\}/' X > X.tmp && mv X{.tmp,}"
