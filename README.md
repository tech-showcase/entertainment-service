## ENTERTAINMENT SERVICE

### Description
This repo contains project that act as a **service**.
This service is part of a big system. 
The whole system will be used to present **micro-services without an orchestrator**.

### Features
- Serve movie data through gRPC.

### API
Please refer to all proto file [here](proto) for more detail about the provided API.
You can use gRPC by:
- Installing [protoc](http://google.github.io/proto-lens/installing-protoc.html)
- Installing [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/)
- Generating code by executing `protoc -I./proto --go_out=plugins=grpc:. proto/*/*.proto`

### How to run
#### Docker
- Install docker
- Create config file `.env` under root dir which contains following content
```
MOVIE_SERVER_ADDRESS=http://www.omdbapi.com/
MOVIE_API_KEY=faf7e5bb
```
- Fill env var `CONFIG_FILEPATH` with directory path where config file is contained
- Fill env var `CONFIG_FILENAME` with the name of config file (e.g `.env`)
- Build and run docker image as below
```shell script
$ docker build -t entertainment-service .
$ docker run -p 8082:8080 entertainment-service
```

### Tech / Dependency
- [Go kit - service](https://github.com/go-kit/kit)
- [Cobra - cli app](https://github.com/spf13/cobra)
- [Viper - config](https://github.com/spf13/viper)
- [gRPC - api](https://grpc.io/)
