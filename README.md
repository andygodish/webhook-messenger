# Golang Base

This repo is used as a scaffolding for new Golang projects within a Docker container. 

## Docker

Use the dockerfile to build a development environment locally. You can update the version of go you install in your container by updating the version of goland used in the base images of the Dockerfile. 

Once built, simply run `docker run -it -v ${PWD}:/work go sh`. The use of the volume will allow you to use your IDE outside of your running docker container. 

### Build Local Development Container

```
docker build -t goland:dev --target dev .
```

The `--target dev` will only build the dev base defined in your Dockerfile. 

## Go Module

```
go mod init github.com/andygodish/golang-base
```

## Build

From your interactive terminal (running docker container), build you hello world app:



