FROM  grengojbo/go:latest
MAINTAINER Jens Bissinger "mail@jens-bissinger."

ADD . /go/src/github.com/UserStack/ustackweb
WORKDIR /go/src/github.com/UserStack/ustackweb
RUN make setup_prod build

CMD ["./ustackweb"]
