#!/bin/sh
set -ex

docker build -t verify-service .

if [[ -d artifacts ]]; then
  rm -f artifacts/*
else
  mkdir artifacts
fi

docker run --rm -it -v $PWD/artifacts:/artifacts verify-service $1
