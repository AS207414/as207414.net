#!/usr/bin/env bash

set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

while getopts s: flag
do
    case "${flag}" in
        s) APPSOURCE=${OPTARG};;
    esac
done

LF_SHA=$(git describe --always) &&
    echo "Setting linkerflag:git_sha to $LF_SHA"

LF_VERSION="${SEMVER}" &&
    echo "Setting linkerflag:version to $LF_VERSION"

LF_GOOS=$(go env GOOS) &&
    echo "Setting linkerflag:GOOS to $LF_GOOS"

LF_GOARCH=$(go env GOARCH) &&
    echo "Setting linkerflag:GOARCH to $LF_GOARCH"

LF_GOVER=$(go env GOVERSION) &&
    echo -e "Setting linkerflag:GOVER to $LF_GOVER \n"


LF_CMD="-s -w -X main.version=$LF_VERSION -X main.os_ver=$LF_GOOS -X main.os_arc=$LF_GOARCH -X main.go_ver=$LF_GOVER -X main.git_sha=$LF_SHA"

BUILD_DIR="bin"
OUT_FILE="$BUILD_DIR/as207414_${LF_GOOS}_${LF_GOARCH}"

echo -e "Linkerflag flag is: \n ${LF_CMD} \n"

echo "Source application is cmd/$APPSOURCE"
echo "Checking source application exists"

if [[ ! -d cmd/$APPSOURCE ]]; then
    echo "cmd/$APPSOURCE does not exist."
    return 1
else
    echo "cmd/$APPSOURCE exist."
fi

echo -e "\nBuilding GO application cmd/$APPSOURCE:\n"

CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS GOARCH=$GOARCH go build -a -installsuffix cgo -ldflags="$LF_CMD" -o=$OUT_FILE ./cmd/$APPSOURCE