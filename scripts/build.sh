#!/bin/bash

docker_repo="mayaleng/engine"
docker_tag=""

echo "Starting Docker publish process"

build_version=$(git rev-parse --short HEAD)

if [[ ! -z "${TRAVIS_TAG}" ]]; then
	docker_tag=${TRAVIS_TAG}
  build_version=$(git describe --tags --abbrev=0)
elif [[ $TRAVIS_BRANCH == "develop" ]]; then
	docker_tag="latest"
elif [[ $TRAVIS_BRANCH =~ feature/.+|hotfix/.+ ]]; then
	docker_tag=$(echo "$TRAVIS_BRANCH" | sed "s/\//-/")
fi

if [[ -z $docker_tag ]]; then
	echo "Skipping publish docker image. Valid when tag or push to [develop, feature/, hotfix/]"
	exit 0
fi

echo "Tag to be published: ${docker_tag}"

time=$(date)

docker build --build-arg BUILD_VERSION="$build_version" --build-arg BUILD_TIME="$time" -f ./build/package/Dockerfile -t "${docker_repo}:${docker_tag}" .

docker push "${docker_repo}:${docker_tag}"
