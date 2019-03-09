#!/bin/bash

build_sensors() {
    cd $(pwd)/sensors
    GOOS=linux GOARCH=amd64 go build
    cd -
}

build_coordinator() {
    cd $(pwd)/coordinator/exec
    GOOS=linux GOARCH=amd64 go build -o coordinator
    cd -
}

build_metric_manager(){
    cd $(pwd)/monitoring/exec
    GOOS=linux GOARCH=amd64 go build -o metricmanager
    cd -
}

docker_compose(){
    docker-compose up -d --build
}

main (){
    build_sensors
    build_coordinator
    build_metric_manager

    docker_compose
}

main