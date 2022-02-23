FROM golang:alpine
WORKDIR /app
ADD . .
# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]
