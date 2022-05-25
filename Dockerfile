FROM golang:alpine
RUN apk add -U --no-cache make gcc g++
WORKDIR /app
COPY . .
RUN go mod download
RUN  go install github.com/cespare/reflex@latest
CMD tail -f /dev/null