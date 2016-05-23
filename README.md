# webrtcservice

Exercice: create server service for wevrtc connection initialisation between two peers

# Status

[![Build Status](https://travis-ci.org/cedriclam/webrtcservice.svg?branch=master)](https://travis-ci.org/cedriclam/webrtcservice)
[![Coverage Status](https://coveralls.io/repos/github/cedriclam/webrtcservice/badge.svg?branch=master)](https://coveralls.io/github/cedriclam/webrtcservice?branch=master)

# Requirement

[Requirement.md](Requirement.md)


# Install and Build

First you need to setup the project. It will download the golang tools and and vendor the dependencies project

```
make init
```

To build the server binary.

```
make build
```

Now you can run the unit-tests.

```
make test
```

Test coverage information

```
make cover
```

# Other commands

```
make help

  Usage:
    make <target>
  Targets:
    all                    Install tools and run the following targets: vendor, build, validate, test, docker
    build-docker           Build rest server docker image
    build                  Build rest server binary
    clean                  Clean project
    cover-extra            Run test coverage and generate report in standard output
    cover-xml              Run test coverage and generate report in _reports folder
    cover                  Run all test coverage
    docker                 Build docker images
    get-tools              Download all tools dependencies in hack sub-directory
    help                   Display list of targets
    init                   Init project
    run                    Run currencyconverter server process
    test-xml               Run all tests and generate reports in _reports folder
    test                   Run unit-tests
    validate               Run go style validation (golint)
    vendor                 Install all vendor dependencies
```

# Test with web browser

start the server with the following command

```
make run
```

then open 2 web browser tabs. You should be able to exchange message between the two client page.

```
open http://0.0.0.0:9090
```

# Open Questions

## How many users can connect to one server?

Already a lot of users can be connected thanks to this implementation. more than 5000 with a simple serve

## How can the system support more users?

We can run several instance of this service behind a loadbalancer.
Then we need to update the implementation in order to share a database for the connections information synchronisation.
Redis can be a good choice with his PubSub mechanism: we create a PubSub channel by connection.
