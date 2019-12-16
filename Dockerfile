FROM golang:latest

ADD . /go/src/github.com/rastringer/artist_studio

WORKDIR /go/src/github.com/rastringer/artist_studio 

RUN go install github.com/rastringer/artist_studio

ENTRYPOINT /go/bin/artist_studio

EXPOSE 8080 