FROM golang:1.19.3

WORKDIR /usr/home

COPY . .

RUN apt-get update  && \
    apt-get install -y protobuf-compiler && \
    apt-get install sqlite3 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/ktr0731/evans@latest

EXPOSE 50051


