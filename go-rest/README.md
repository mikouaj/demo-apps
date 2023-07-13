# spring-rest-jpa

The folder contains source code of a simple Go application that exposes the REST endpoint.

[![Build](https://github.com/mikouaj/demo-apps/actions/workflows/build-go-rest.yaml/badge.svg)](https://github.com/mikouaj/demo-apps/actions/workflows/build-go-rest.yaml)

## Usage

The service exposes two endpoints:

* `/healthz` reports health on the application
* `/books` returns collection of books

### Source code

Prerequisites:

* [Go](https://go.dev/doc/install) 1.20 or newer
* [GNU Make](https://www.gnu.org/software/make)

Build application from the source code.

```sh
make
./go-rest
```

### Container image

Run the container locally with Docker.

```sh
docker run -d  --name spring-rest-jpa \
  -p 8080:8080 \
  ghcr.io/mikouaj/go-rest:latest 
```

### Kubernetes

1. Adjust `kustomization.yaml` according to your needs
2. Customize and apply Kubernetes manifests

   ```sh
   kubectl kustomize | kubectl apply -f
   ```
