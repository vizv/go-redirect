FROM ubuntu:latest
MAINTAINER Viz <viz@linux.com>

# add the app
ADD go-redirect /app/

# setup the app
EXPOSE 8080
ENTRYPOINT /app/go-redirect
