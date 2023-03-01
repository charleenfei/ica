FROM ubuntu:20.04

WORKDIR /home/ubuntu

env PATH /usr/local/go/bin:$PATH

RUN apt-get update && apt-get install -y wget make jq
RUN wget https://go.dev/dl/go1.18.9.linux-amd64.tar.gz
RUN tar -C /usr/local -xzvf go1.18.9.linux-amd64.tar.gz

ENV GOPATH=""
ENV GOMODULE="on"

COPY . .

#RUN rm -rf ./build

#RUN go mod download
#RUN make build

RUN cp ./build/icad /bin/icad
