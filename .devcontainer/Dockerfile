FROM ubuntu:latest
COPY --from=golang:1.21.6-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"