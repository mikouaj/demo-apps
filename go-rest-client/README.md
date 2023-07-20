# go-rest-client

The folder contains source code of a simple Go application that exposes the REST endpoint.
The data on this endpoint are either fetched from the remote URL or generated locally.

[![Build](https://github.com/mikouaj/demo-apps/actions/workflows/build-go-rest-client.yaml/badge.svg)](https://github.com/mikouaj/demo-apps/actions/workflows/build-go-rest-client.yaml)

## Usage

The service exposes two endpoints:

* `/healthz` reports health on the application
* `/data` returns the data originating either from remote URL or local data provider

### Source code

Prerequisites:

* [Go](https://go.dev/doc/install) 1.20 or newer
* [GNU Make](https://www.gnu.org/software/make)

Build application from the source code.

```sh
make
./go-rest-client
```
