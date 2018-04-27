# drone-greenkeeper

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-greenkeeper/status.svg)](http://beta.drone.io/drone-plugins/drone-greenkeeper)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-greenkeeper?status.svg)](http://godoc.org/github.com/drone-plugins/drone-greenkeeper)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-greenkeeper)](https://goreportcard.com/report/github.com/drone-plugins/drone-greenkeeper)
[![](https://images.microbadger.com/badges/image/plugins/greenkeeper.svg)](https://microbadger.com/images/plugins/greenkeeper "Get your own image badge on microbadger.com")

Drone plugin to use [greenkeeper-lockfile](https://github.com/greenkeeperio/greenkeeper-lockfile) when using a `package.lock` or `yarn.lock` within a JavaScript project. For the usage information and a listing of available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-greenkeeper/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-greenkeeper
docker build --rm -t plugins/greenkeeper .
```

## Usage

Update a lockfile:

```sh
docker run --rm \
  -e DRONE=true \
  -e DRONE_REMOTE_URL=https://github.com/octocat/hello-world.git \
  -e DRONE_REPO=octocat/hello-world \
  -e DRONE_COMMIT_BRANCH=greenkeeper/foo-1.0.0 \
  -e DRONE_COMMIT_MESSAGE=chore(package) update foo to version 1.0.0 \
  -e DRONE_BUILD_EVENT=push \
  -e DRONE_JOB_NUMBER=1 \
  -e PLUGIN_UPDATE=true \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/greenkeeper
```

Upload a lockfile:

```sh
docker run --rm \
  -e DRONE=true \
  -e DRONE_REMOTE_URL=https://github.com/octocat/hello-world.git \
  -e DRONE_REPO=octocat/hello-world \
  -e DRONE_COMMIT_BRANCH=greenkeeper/foo-1.0.0 \
  -e DRONE_COMMIT_MESSAGE=chore(package) update foo to version 1.0.0 \
  -e DRONE_BUILD_EVENT=push \
  -e DRONE_JOB_NUMBER=1 \
  -e PLUGIN_UPDATE=true \
  -e GREENKEEPER_GITHUB_TOKEN=@@@@@@@ \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/greenkeeper
```