# Go Proxmox

Go Proxmox is a Go client library for accessing the Proxmox VE API.

It consists of two parts:

1. proxmox client library (found under `proxmox` dir)
2. CLI for interacting with proxmox server (found in `main.go` in root,
   reorganization TBD)

The client library is currently being used in the
[terraform-provider-proxmox](https://github.com/thirdwavellc/terraform-provider-proxmox)
repo, while the CLI was originally implemented as a quick way to test out the
client code during development. As such, the CLI is a bit messy at the moment
and will likely be refactored if/when it becomes a necessity.

## Work in Progress

This repo is currently a work in progress, with limited functionality provided
at this point.

## Setup

We are using [dep](https://github.com/golang/dep) to manage go dependencies.
Once you have dep installed, to install the project's dependencies:

```bash
$ dep ensure
```

This will install them under `vendor`.

To build the project:

```bash
$ go build
```

or

```bash
$ make
```

The difference between these is that go build will only build for your local
OS. Running `make` will also run the tests (currently there are none), before
using [gox](https://github.com/mitchellh/gox) to cross-compile for multiple OS
distribution.
