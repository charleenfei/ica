FROM golang:1.18 as builder
ARG arch=x86_64

ENV GOPATH=""
ENV GOMODULE="on"

COPY . .

RUN go mod download
RUN make build

FROM ubuntu:20.04

COPY --from=builder /go/build/icad /bin/icad

ENTRYPOINT ["icad"]
