#!/bin/bash
set -eux

(cd ./cmd/makebb && GO111MODULE=on go build -covermode=atomic .)
(cd ./test && go test --makebb=../cmd/makebb/makebb -v)
