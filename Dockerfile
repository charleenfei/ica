FROM golang:1.18 as builder
ARG arch=x86_64

ENV GOPATH=""
ENV GOMODULE="on"

COPY . .

RUN go mod download
RUN make build

RUN cp /go/build/icad /usr/bin/

RUN apt-get install bash