FROM golang:latest

RUN mkdir /price-generator

COPY . /price-generator

WORKDIR /price-generator

RUN go build -o main

CMD ["/price-generator/main"]

EXPOSE 8083