#!/bin/sh
set -ex

DATE_STR=$(date +%Y-%m-%d)
GITC=$(git rev-parse HEAD)
TRUNC_GITC=${GITC:35:40}
RELEASE_NAME=$DATE_STR-$TRUNC_GITC
sh build.sh $RELEASE_NAME

pushd artifacts
for PLATFORM in {darwin,linux,windows}; do
  FOLDERNAME="verify-service_${PLATFORM}"
  ZIPNAME="verify-service_${PLATFORM}_amd64.zip"
  mkdir -p $FOLDERNAME
  for FILE in $(ls *$PLATFORM*); do
    if [[ -f $FILE ]]; then
      cp $FILE $FOLDERNAME
    fi
  done
  cp -r ../../schema $FOLDERNAME/schema
  cp -r ../../retailops-asp-dotnet-api-example $FOLDERNAME/retailops-asp-dotnet-api-example
  cp ../README.md $FOLDERNAME/README.md
  zip -r $ZIPNAME $FOLDERNAME
done
popd