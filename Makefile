BACKEND=github.com/UserStack/ustackd

all: run
watch:
	npm run watch
run:
	cd src/ustackweb && bee run
prepare:
	go get -u github.com/astaxie/beego \
				    github.com/beego/bee \
				    github.com/beego/i18n \
				    github.com/beego/wetalk \
				    github.com/codegangsta/gin \
				    ${BACKEND}
	make backend
	npm install
	npm run bower
backend:
	go get -u ${BACKEND}
	go install ${BACKEND}
	ustackd -f
test:
	go test ./src/ustackweb/...
