FROM ubuntu:18.04

COPY . /code
FROM go

COPY . /code
WORKDIR /code
EXPOSE 80
ENTRYPOINT ["./test"]