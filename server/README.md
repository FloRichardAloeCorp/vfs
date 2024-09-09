# Server

## Introduction

Gin HTTP server developed in Golang to interact with the `vfs` through HTTP.

## Features

* CRUD operations on `VFS`

## Usage

### Running the server

Bare metal:

```
go run main.go
```

Docker:

```
docker compose up
```

### API Endpoints

A swagger describing the API is available [here](./docs/swagger.yaml)

## Configuration

### Env variables

* **VFS_CONFIG**: Path to config file (default: ./config/config.yaml)

### Config file

The API is configuration with `.yaml` config file. Check [config.yaml](./config/config.yaml) to see configuration keys.

All keys can be overriden with environment variable (VFS_KEY_PATH).

## Architecture

The API is designed following the principles of "Screaming Architecture." This architecture ensures that developers clearly understand what they are building because the architecture "screams" its purpose. It organizes feature-related entities and separates ingress and datasource handling, resulting in a simple and scalable API structure.

### Packages

* **internal/app**: Initializes the API and launches the server.
* **internal/config**: Contains the server configuration parsing functions.
* **internal/features/feature_name**: Groups features in directories, each containing:
  * An interface
  * A controller implementing the interface
  * A repository for data source manipulation
  * A data model
  * A factory method
* **internal/interfaces**: Defines the server's interfaces, including ingress and datasources.

**NOTE:** I've developed a CLI to create such an API. It initializes the project, generates boilerplate code, and more. You can find it [here](https://github.com/Aloe-Corporation/screamer) if I've finished its development.
