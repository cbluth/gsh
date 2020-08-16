#!/usr/bin/env bash
#
# builds the project in temp docker env
#
set -e

PROJECT="$(grep module go.mod | awk '{print $2}')"
SCRIPTPATH="$(dirname "$(readlink -f "${0}")")"
DOCKERTAG="${PROJECT}-build:tmp"
BINARY="${PROJECT}"
SILENT="${1}" # ./build.sh -s
DOCKERFILE="$(cat << EOF
FROM golang:latest as build
SHELL ["/bin/bash", "-c"]
WORKDIR /build
ADD . /build
RUN GOOS=linux GOARCH=amd64 go build \
    -a \
    -o /build/${PROJECT}-linux \
    .
RUN GOOS=darwin GOARCH=amd64 go build \
    -a \
    -o /build/${PROJECT}-darwin \
    .
FROM golang:latest
COPY --from=build /build/${PROJECT}-linux /build/
COPY --from=build /build/${PROJECT}-darwin /build/
CMD ["bash", "-c", "tar -cv /build/ | gzip"]
EOF
)"

spinner()
{
    # http://fitnr.com/showing-a-bash-spinner.html
    local pid=$1
    local delay=0.15
    local spinstr='|/-\'
    while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
}

build()
{
    pushd "${SCRIPTPATH}/" > /dev/null 2>&1
    echo
    echo "Using temporary docker build environment..."
    echo
    docker build \
        -t "${DOCKERTAG}" \
        -f - \
        . <<< "${DOCKERFILE}"
    docker run --rm "${DOCKERTAG}" > "release.tgz"
    # docker rmi "${DOCKERTAG}"
    # chmod a+rx "${BINARY}"
    echo
    # echo "Removed temporary build environment"
    # echo
    popd > /dev/null 2>&1
    echo -n "Build time:"
}

echo -n "Building... "
if [[ ! "${SILENT}" == "-s" ]] ; then
    (
        time build
        echo
        # echo "${PROJECT^^} project executable: ${HOME}/.local/bin/${BINARY}"
        # echo
    ) &
else
    (build) > /dev/null 2>&1 &
fi

spinner $!

# cp "${BINARY}" "${HOME}/.local/bin/"

echo "Done!"
