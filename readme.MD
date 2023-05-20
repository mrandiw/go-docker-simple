# HOW TO CREATE & DEPLOY GOLANG TO DOCKER

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

Setup

    1. Build Script Dockerfile -> run : docker build --tag go-docker-simple .
    2. Create Network -> run : docker network create go-docker-simple-network
    3. Run Docker Container Golang -> run : docker run -d -p 8080:8080 --name go-docker-simple --network go-docker-simple-network go-docker-simple
