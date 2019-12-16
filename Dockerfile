FROM golang:1.11-alpine AS build


WORKDIR /go/src/github.com/rastringer/artist_studio 
COPY . .
COPY templates templates

ENV VERBOSE=0
ENV PKG=github.com/rastringer/artist_studio
ENV ARCH=amd64
ENV VERSION=test

RUN CG0_ENABLED=0 go build -o /artist_studio

FROM alpine

COPY --from=build /artist_studio /artist_studio
COPY ./templates ./templates
EXPOSE 8080
ENTRYPOINT ["/artist_studio"]
