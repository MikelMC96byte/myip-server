#!/bin/bash
workdir="$(pwd)"
buildir="${workdir}/build"
sourcedir="${workdir}/src"
version="0.1.0"

packagename="myipserver"

rm -rf ${buildir}
mkdir -p ${buildir}

# Compilation for Linux AMD64
env GOOS=linux go build -o ${buildir}/linux/amd64/${packagename} ${sourcedir}
cd ${buildir}/linux/amd64
zip -r ${buildir}/${packagename}-linux-amd64.zip *
cd ${workdir}

# Docker build for Linux AMD64 
# docker build -t ${packagename}:latest -t ${packagename}:${version} .