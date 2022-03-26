#!/usr/bin/env bash

set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# SET enviromental variables and WARN user
[[ -z "${BUILD_OS}" ]] &&
    echo "BUILD_OS environmental variable not found. Defaulting to 'linux'" &&
[[ -z "${BUILD_ARCH}" ]] &&
    echo "BUILD_ARCH environmental variable not found. Defaulting to 'amd64'. " &&

# SET MAIN Variables
BUILD_OS="${BUILD_OS:=linux}"
BUILD_ARCH="${BUILD_ARCH:=amd64}"
BIN_NAME="as207414_${BUILD_OS}_${BUILD_ARCH}"

while getopts f: flag
do
    case "${flag}" in
        f) IMAGENAME=${OPTARG};;
    esac
done

echo "CLEAN existing binaries in build/docker" && \
    rm -f build/docker/$BIN_NAME

echo "COPY bin to build/docker" && \
    cp bin/$BIN_NAME build/docker/$BIN_NAME && \
    chmod +x build/docker/$BIN_NAME

echo "BUILD docker image to image ${IMAGENAME}" && \
    docker build build/docker -t $IMAGENAME