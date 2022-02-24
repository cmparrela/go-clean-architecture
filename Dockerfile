FROM golang:alpine
WORKDIR /app
COPY . .
ENTRYPOINT [ "docker/php/entrypoint.sh" ]
