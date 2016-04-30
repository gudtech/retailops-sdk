#!/bin/sh
set -ex

DATE_STR=$(date +%Y-%m-%d)
GITC=$(git rev-parse HEAD)
TRUNC_GITC=${GITC:35:40}
RELEASE_NAME=$DATE_STR-$TRUNC_GITC
# sh build.sh $RELEASE_NAME

pushd artifacts
for FILE in $(ls); do
  if [[ -f $FILE ]]; then
    echo $FILE
  fi
done
popd