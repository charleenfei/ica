FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
  && apt-get install -y python3-pip python3-dev \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /src/app/

COPY . .

RUN pip3 install -r requirements.txt