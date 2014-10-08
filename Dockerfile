FROM  grengojbo/go:latest
MAINTAINER Jens Bissinger "mail@jens-bissinger."

ADD . /go/src/github.com/UserStack/ustackweb
WORKDIR /go/src/github.com/UserStack/ustackweb
RUN make deps
RUN go install github.com/beego/bee
RUN cd /go/src/github.com/astaxie/beego && git remote add bsingr https://github.com/bsingr/beego.git && git fetch bsingr && git checkout -b bsingr-develop bsingr/develop

CMD ["go", "run", "bee", "run"]
