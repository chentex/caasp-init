# caasp-init

[![Release](https://img.shields.io/github/release/kubic-project/caasp-init.svg)](https://github.com/kubic-project/caasp-init/releases/latest)
[![Build Status](https://img.shields.io/travis/kubic-project/caasp-init/master.svg)](https://travis-ci.org/kubic-project/caasp-init)
[![codecov](https://codecov.io/gh/kubic-project/caasp-init/branch/master/graph/badge.svg)](https://codecov.io/gh/kubic-project/caasp-init)

![License: Apache 2.0](https://img.shields.io/github/license/kubic-project/caasp-init.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubic-project/caasp-init)](https://goreportcard.com/report/github.com/kubic-project/caasp-init)

Initialize docker daemon mirrors configuration

## Usage

caasp-init will create the daemon.json configuration file and the necessary certificates for the mirror you want to config based on the kubic-init.yaml configuration file.

```shell
Usage:
  caasp-init [flags]
  caasp-init [command]

Available Commands:
  help        Help about any command
  version     Show version of caasp-init

Flags:
  -c, --config string   kubibc-init.yaml config file (default "/etc/kubic/kubic-init.yaml")
  -h, --help            help for caasp-init

Use "caasp-init [command] --help" for more information about a command.
```

## Example

`$ caasp-init`

This will use default value for configuration file

`/etc/kubic/kubic-init.yaml`

if you want to indicate a file run

`$ caasp-init -c /path/to/config/file.yaml`

If the configuration file has mirrors declared ot will generate the daemon.json
file with the following structure:

```
{
  "registries": [
    {
    "Mirrors": [
      {
      "URL": "https://airgappedregistry.com"
      }
    ],
    "Prefix": "https://registry.suse.com"
    }
  ],
  "iptables":false,
  "log-level": "warn"
}
```

If there is no mirror declared the configuration file will just be the default:

```
{
  "iptables":false,
  "log-level": "warn"
}
```

For help use `caasp-init help`

### help

Displays the current version of caasp-init.

### version

Displays the current version of caasp-init.

## Install

Download the latest version from [releases](https://github.com/kubic-project/caasp-init/releases/latest).

`$ tar -xvf caasp-init-linux.tgz -C /opt`

The create a symbolic link

`$ ln -s /opt/caasp-init/bin/caasp-init /usr/local/bin/caasp-init`

## Test

Clone repository into your $GOPATH. You can also use go get:

`go get github.com/kubic-project/caasp-init`

### Prerequisites

- Have [GO](https://golang.org/) installed.

### Running tests

To run test on this package simply run:

`make test`

#### Testing with Docker

`make test.unit`

## Building

Be sure you have all prerequisites, then build the binary by simply running

`make caasp-init`

your binary will be stored under `bin` folder

[1]: https://github.com/SUSE/skopeo/blob/sync/docs/skopeo.1.md#skopeo-sync