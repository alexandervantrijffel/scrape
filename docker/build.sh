#!/bin/sh -x

./.githooks/pre-commit
RESULT=$?
if [ $RESULT -ne 0 ]; then
    echo pre-commit script failed with exit code $RESULT
    exit $RESULT
fi

rm docker/scrape || true
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags \"-static\"' -o docker/scrape
RESULT=$?
if [ $RESULT -ne 0 ]; then
    echo go build failed with exit code $RESULT
    exit $RESULT
fi

docker build -t scrape docker/
