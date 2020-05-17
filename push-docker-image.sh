#!/bin/bash

REPO="mayaleng/engine"
TAG=

echo "Starting Docker publish process"

if [[ ! -z "${TRAVIS_TAG}" ]]; then
	TAG=${TRAVIS_TAG}
elif [[ $TRAVIS_BRANCH == "develop" ]]; then
	TAG="latest"
elif [[ $TRAVIS_BRANCH =~ feature/.+|hotfix/.+ ]]; then
	TAG=$(echo "$TRAVIS_BRANCH" | sed "s/\//-/")
fi

if [[ -z $TAG ]]; then
	echo "Skipping publish docker image. Valid when tag or push to [develop, feature/, hotfix/]"
	exit 0
fi

echo "Tag to be published: ${TAG}"

docker build -f ./build/package/Dockerfile -t "${REPO}:${TAG}" .

docker push "${REPO}:${TAG}"
