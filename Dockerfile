FROM ubuntu:18.04

COPY . /code
FROM golang

COPY . /code
WORKDIR /code
EXPOSE 8080