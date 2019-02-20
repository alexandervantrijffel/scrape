#!/bin/sh -x

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
pushd $DIR

export $(cat .env | xargs)

go run -tags=debug main.go "$1"

popd
