## ENTERTAINMENT SERVICE

### Description
This repo contains project that act as a service of microservices system.
This service is part of a big system. 
The whole system will be used to present technology show case.

### Features
- Serve movie data

This service serve data that is mentioned above through GRPC.

### How to run
#### Docker
- Install docker
- Create `config-dev.json` under `config` dir which contains following content
```json
{
  "movie": {
    "server_address": "http://www.omdbapi.com/",
    "api_key": "faf7e5bb"
  }
}
```
- Build and run docker image as below
```shell script
$ docker build -t entertainment-service .
$ docker run -p 8082:8080 entertainment-service
```
