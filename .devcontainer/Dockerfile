FROM ubuntu:latest

# Install Go
COPY --from=golang:1.21.6-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

# Install Git
RUN apt update
RUN apt install -y git