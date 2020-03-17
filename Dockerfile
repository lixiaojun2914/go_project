FROM golang

COPY . /code
WORKDIR /code
EXPOSE 8080
CMD [ "./test" ]