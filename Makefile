all: run
livereload:
	npm run livereload
run:
	cd src/ustackweb && bee run
prepare:
	go get github.com/astaxie/beego \
		github.com/beego/bee \
		github.com/codegangsta/gin
	npm install
