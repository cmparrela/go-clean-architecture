FROM golang:alpine
RUN apk add -U --no-cache make gcc g++
WORKDIR /app
COPY . .
RUN chmod +x ./entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]