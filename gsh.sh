#!/usr/bin/env bash

cd "$(dirname "$(readlink -f "${0}")")" > /dev/null 2>&1
go run . "${@}"
