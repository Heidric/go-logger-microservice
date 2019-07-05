#!/bin/bash

docker stop logger
docker stop loggerdb

docker rm logger
docker rm loggerdb

docker rmi logger
docker rmi loggerdb

docker rmi iron/go
docker rmi postgres
