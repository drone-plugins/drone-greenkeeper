# drone-greenkeeper

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-greenkeeper/status.svg)](http://cloud.drone.io/drone-plugins/drone-greenkeeper)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/greenkeeper.svg)](https://microbadger.com/images/plugins/greenkeeper "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-greenkeeper?status.svg)](http://godoc.org/github.com/drone-plugins/drone-greenkeeper)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-greenkeeper)](https://goreportcard.com/report/github.com/drone-plugins/drone-greenkeeper)

Drone plugin to use [greenkeeper-lockfile](https://github.com/greenkeeperio/greenkeeper-lockfile) when using a `package.lock` or `yarn.lock` within a JavaScript project. For the usage information and a listing of available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-greenkeeper/).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-greenkeeper
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/greenkeeper .
```

## Usage

```console
docker run --rm \
  -e DRONE_REPO=octocat/hello-world \
  -e DRONE_REMOTE_URL=https://github.com/octocat/hello-world.git \
  -e DRONE_BUILD_EVENT=push \
  -e DRONE_COMMIT_BRANCH=greenkeeper/foo-1.0.0 \
  -e DRONE_COMMIT_MESSAGE=chore(package) update foo to version 1.0.0 \
  -e DRONE_JOB_NUMBER=1 \
  -e PLUGIN_UPDATE=true \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/greenkeeper
```
