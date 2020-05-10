accentis-core-service
===

The Core service provides an API to manipulate the core resources for the Accentis system.

## Documentation

The API documentation is generated from the Swagger 2.0 specification.

```bash
$ swagger server ./swagger/swagger.yaml
```

## Building

This project includes code generated from the Swagger 2.0 specification.  This code resides in the **models/** and **restapi/** packages.

```bash
$ swagger generate server -f ./swagger/swagger.yaml --name core
```

Once the code has been regenerated, it can be compiled with the Go compiler.

```bash
$ go build ./...
$ go build ./cmd/core-server
```

## Running

This service can be run locally on an randomly selected ephemeral port.

```bash
$ ./core-server
```
