#!/bin/bash
set -ex

DATE_STR=$(date +%Y-%m-%d)
GITC=$(git rev-parse HEAD)
TRUNC_GITC=${GITC:35:40}
RELEASE_NAME=$DATE_STR-$TRUNC_GITC

docker build -t verify .
if [[ -d artifacts ]]; then
  rm -rf $PWD/artifacts/*
else
  mkdir $PWD/artifacts
fi
docker run --rm -it -v $PWD/artifacts:/artifacts verify

pushd artifacts
for PLATFORM in {darwin,linux,windows}; do
  FOLDERNAME="verify_${PLATFORM}_${RELEASE_NAME}"
  ZIPNAME="${FOLDERNAME}.zip"
  mkdir -p $FOLDERNAME
  for FILE in $(ls *$PLATFORM*); do
    if [[ -f $FILE ]]; then
      cp $FILE $FOLDERNAME/verify
      if [[ $PLATFORM == "windows" ]]; then
        cp $FOLDERNAME/verify $FOLDERNAME/verify.exe
      fi
    fi
  done
  cp -r ../../schema $FOLDERNAME/schema
  cp -r ../../retailops-asp-dotnet-api-example $FOLDERNAME/retailops-asp-dotnet-api-example
  cp ../README.md $FOLDERNAME/README.md
  zip -r $ZIPNAME $FOLDERNAME
done
popd