#!/bin/bash

find . -type f -name \*.go | xargs -n1 -IX \
  bash -c "sed -re 's/ ((u)*int(8|16|32|64)*|float(32|64)|bool|string)[  ]+\`json/ \*\1 \`json/' X > X.tmp && mv X{.tmp,}"
