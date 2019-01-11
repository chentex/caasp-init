# Development environment for `caasp-init`

## Project structure

This project follows the conventions presented in the [standard Golang
project](https://github.com/golang-standards/project-layout).

## Dependencies

* `go >= 1.11`

Note: 
We use golang modules but you still need to work inside your $GOPATH for developing `caasp-init`.
Working outside GOPATH is currently **not supported**

## Building

A simple `make` should be enough. This should compile [the main
function](cmd/root.go) and generate a `caasp-init` binary.

## Running unit-tests

Unit tests can be run using `make test`

## Code Coverage

Run first the tests. Then use `make coverage` for visualizing coverage.

Feel free to read more about this on : https://blog.golang.org/cover.

## Generating Releases

Run first the tests. Then use `make release` to generate the release assets.

They will be created in the [release](release/) folder.

## Community

Currently the caasp-init project lives inside the kubic echosystem.

If you have a question to ask? Want to join in the discussion? Find community information including chat and mailing lists on the main [Kubic](https://en.opensuse.org/Portal:Kubic) page.

Want to get involved but don't know what to do? Try looking at our github [issues](https://github.com/kubic-project/caasp-init/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) and use the tags `good first issue` or `help wanted`!
