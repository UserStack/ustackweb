FROM  grengojbo/go:latest
MAINTAINER Jens Bissinger "mail@jens-bissinger."

ADD . /go/src/github.com/UserStack/ustackweb
WORKDIR /go/src/github.com/UserStack/ustackweb
RUN go get github.com/tools/godep
RUN make assets

CMD ["godep", "go", "run", "main.go"]
