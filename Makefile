all: run
watch:
	npm run watch
run:
	cd src/ustackweb && bee run
prepare:
	go get -u github.com/astaxie/beego \
				    github.com/beego/bee \
				    github.com/beego/i18n \
				    github.com/codegangsta/gin
	make backend
	npm install
	npm run bower
backend:
	go get -u github.com/UserStack/ustackd
test:
	go test ./src/ustackweb/...
