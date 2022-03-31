#!/bin/bash

# https://hub.docker.com/_/redis?tab=description

docker run \
  --detach \
  --name device-db \
  --env MONGO_INITDB_ROOT_USERNAME=mongoadmin \
  --env MONGO_INITDB_ROOT_PASSWORD=secret \
  --volume $(pwd)/temp/device-db:/data \
  --publish 6379:6379 \
  redis:6.2-alpine
