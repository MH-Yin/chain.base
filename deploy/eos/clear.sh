#!/bin/bash
docker stop nodeos
docker stop keosd
docker rm nodeos
docker rm keosd