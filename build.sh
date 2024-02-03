#!/bin/bash

WORKSPACE="$(dirname "$(realpath "$0")")"
BIN_DIR="${WORKSPACE}/binary"
BIN_NAME="gen-ag-table"

echo "Using WORKSPACE: $WORKSPACE"

mkdir -p "${BIN_DIR}"

set -x

LD_FLAGS="\
-X github.com/alexditu/go-utils/pkg/version.commitHash=$(git rev-parse --short HEAD) \
-X github.com/alexditu/go-utils/pkg/version.branchName=$(git rev-parse --abbrev-ref HEAD) \
-X github.com/alexditu/go-utils/pkg/version.binaryName=${BIN_NAME} \
-X github.com/alexditu/go-utils/pkg/version.major=1 \
-X github.com/alexditu/go-utils/pkg/version.minor=0"

GOARCH=amd64 GOOS=linux \
go build -x -o "${BIN_DIR}/${BIN_NAME}" \
		-ldflags "${LD_FLAGS}" \
		.