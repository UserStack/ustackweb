FROM  grengojbo/go:latest
MAINTAINER Jens Bissinger "mail@jens-bissinger."

ADD . /go/src/github.com/UserStack/ustackweb
WORKDIR /go/src/github.com/UserStack/ustackweb
RUN make deps
RUN go install github.com/beego/bee
RUN make beego_develop

CMD ["go", "run", "bee", "run"]
