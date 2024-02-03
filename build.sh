#!/bin/bash

WORKSPACE="$(dirname "$(realpath "$0")")"
BIN_DIR="${WORKSPACE}/binary"
BIN_NAME="gen-ag-table"

echo "Using WORKSPACE: $WORKSPACE"

mkdir -p "${BIN_DIR}"

set -x

LD_FLAGS="\
-X github.com/alexditu/ag-csv-table-generator/version.commitHash=$(git rev-parse --short HEAD) \
-X github.com/alexditu/ag-csv-table-generator/version.branchName=$(git rev-parse --abbrev-ref HEAD) \
-X github.com/alexditu/ag-csv-table-generator/version.binaryName=${BIN_NAME}"

GOARCH=amd64 GOOS=linux \
go build -x -o "${BIN_DIR}/${BIN_NAME}" \
		-ldflags "${LD_FLAGS}" \
		.