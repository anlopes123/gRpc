FROM golang:1.19.3

WORKDIR /usr/home

COPY . .

RUN apt-get update  && \
    apt-get install -y protobuf-compiler




