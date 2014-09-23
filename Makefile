all: run
run:
	cd src/ustackweb && bee run
prepare:
	go get github.com/astaxie/beego github.com/beego/bee
