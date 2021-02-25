#!/bin/bash

find . -type f -name \*.go | xargs -n1 -IX \
  bash -c "sed -re 's/\`json:\"([a-zA-Z0-9\_\-]+)\"\`/\`json:\"\1,omitempty\"\`/' X > X.tmp && mv X{.tmp,}"
