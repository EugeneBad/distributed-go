#!/bin/bash

build_sensors() {
    cd $(pwd)/sensors
    GOOS=linux GOARCH=amd64 go build
    cd -
}

build_coordinator() {
    cd $(pwd)/coordinator/exec
    GOOS=linux GOARCH=amd64 go build
    cd -
}

docker_compose(){
    docker-compose up -d --build
}

main (){
    build_sensors
    # build_coordinator
    docker_compose
}

main