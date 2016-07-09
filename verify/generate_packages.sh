#!/bin/bash
set -ex

if [[ $TRAVIS_TAG == "" && $TRAVIS ]]; then
  echo "TRAVIS_TAG not set. skipping artifact build."
  exit 0
fi

if [[ $TRAVIS_TAG == "" ]]; then
  DATE_STR=$(date +%Y-%m-%d)
  GITC=$(git rev-parse HEAD)
  TRUNC_GITC=${GITC:35:40}
  RELEASE_NAME=$DATE_STR-$TRUNC_GITC
else
  RELEASE_NAME=$TRAVIS_TAG
fi
GO_VERSION=1.6

# docker build -t verify .
if [[ -d artifacts ]]; then
  rm -rf $PWD/artifacts/*
else
  mkdir $PWD/artifacts
fi

build_platform_binary()
{
  docker run \
    -v $PWD:/go/src/github.com/gudtech/retailops-sdk/verify \
    -v $PWD/artifacts:/artifacts \
    -e GOOS=$1 \
    -e GOARCH=$2 \
    golang:$GO_VERSION \
    go build -o /artifacts/verify-$1 /go/src/github.com/gudtech/retailops-sdk/verify/bin/verify.go
}
build_platform_binary windows amd64
build_platform_binary linux amd64
build_platform_binary darwin amd64

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
  cp -r ./schema $FOLDERNAME/schema
  cp -r ./schema $FOLDERNAME/tests
  cp -r ../../retailops-asp-dotnet-api-example $FOLDERNAME/retailops-asp-dotnet-api-example
  cp ../README.${PLATFORM}.md $FOLDERNAME/README.md
  zip -r $ZIPNAME $FOLDERNAME
done

rm -rf $(ls | grep -v zip)

popd
