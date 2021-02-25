#!/bin/bash

find . -type f -name \*.go | xargs -n1 -IX \
  bash -c "sed -re 's/github.com\/stripe/github.com\/Capchase/' X > X.tmp && mv X{.tmp,}"
