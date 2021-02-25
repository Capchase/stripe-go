#!/bin/bash

find . -type f -name \*.go | xargs -n1 -IX \
  bash -c "sed -re 's/\.ID = id/\.ID = \&id/' X > X.tmp && mv X{.tmp,}"
