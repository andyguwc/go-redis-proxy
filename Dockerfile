# DEVELOPMENT ONLY
FROM golang:alpine

ENV CAPACITY=100
ENV GLOBAL_EXPIRY=200
ENV PORT=8080
ENV REDIS_ADDRESS=redis:6379

# Install tools required for the project
RUN apk add --no-cache git mercurial

# Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/andyguwc/go-redis-cache
WORKDIR /go/src/github.com/andyguwc/go-redis-cache

EXPOSE 8080
