#!/bin/bash

./build_images.sh

DOCKER_CONF="${PWD}/.docker"
mkdir -p "${DOCKER_CONF}"
docker --config="${DOCKER_CONF}" login -u="${QUAY_USER}" -p="${QUAY_TOKEN}" quay.io

docker --config="${DOCKER_CONF}" push "${ASSISTED_SERVICE_IMAGE}:latest"
docker --config="${DOCKER_CONF}" push "${ASSISTED_SERVICE_IMAGE}:${TAG}"

docker --config="${DOCKER_CONF}" push "${ASSISTED_ISO_GENERATOR_IMAGE}:latest"
docker --config="${DOCKER_CONF}" push "${ASSISTED_ISO_GENERATOR_IMAGE}:${TAG}"

