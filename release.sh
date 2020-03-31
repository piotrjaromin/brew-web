#!/bin/bash

npm ci --prefix web-ui
npm run build --prefix web-ui

go get github.com/rakyll/statik
statik -f -src=web-ui/build

go test -v

go build

export Commit=$(git rev-list -1 HEAD)
export Version=$(git describe --tags $(git rev-list --tags --max-count=1))

# brew install goreleaser/tap/goreleaser
# brew install goreleaser
goreleaser --rm-dist