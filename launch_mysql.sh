#!/usr/bin/env bash

docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root --name test --rm  -v "$(pwd)"/data:/docker-entrypoint-initdb.d mariadb:latest