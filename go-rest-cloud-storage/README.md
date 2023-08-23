# go-rest-cloud-storage

The folder contains source code of a simple Go application that exposes the REST endpoint.
The endpoint serves the data from the given Google Cloud Storage bucket.

[![Build](https://github.com/mikouaj/demo-apps/actions/workflows/build-go-rest-cloud-storage.yaml/badge.svg)](https://github.com/mikouaj/demo-apps/actions/workflows/build-go-rest-cloud-storage.yaml)

## Usage

The service exposes two endpoints:

* `/healthz` reports health on the application
* `/buckets/<bucket name>` returns the list of objects in a given bucket

### Source code

Prerequisites:

* [Go](https://go.dev/doc/install) 1.20 or newer
* [GNU Make](https://www.gnu.org/software/make)

Build application from the source code.

```sh
make
./go-rest-cloud-storage
```
