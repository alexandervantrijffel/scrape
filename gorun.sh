#!/bin/zsh
DIR=$(dirname $0:A)
pushd $DIR

export $(cat .env | xargs)

go run -tags=debug main.go "$1"

popd
