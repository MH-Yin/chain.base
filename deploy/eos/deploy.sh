#!/bin/bash

if [ $(docker images |grep 'private/eos') -z ]
  then  docker build -t private/eos .
fi

docker-compose -f docker-compose.yaml up -d