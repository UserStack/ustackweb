all: run
watch:
	npm run watch
run:
	cd src/ustackweb && bee run
prepare:
	go get -u github.com/astaxie/beego \
				    github.com/beego/bee \
				    github.com/beego/i18n \
				    github.com/codegangsta/gin \
				    github.com/UserStack/ustackd
	npm install
	npm run bower
test:
	cd src/ustackweb && go test tests/default_test.go
