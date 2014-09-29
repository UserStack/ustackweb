BACKEND=github.com/UserStack/ustackd

all: run
watch:
	npm run watch
run:
	bee run
prepare:
	make deps
	make assets
deps:
	go get -u github.com/astaxie/beego \
				    github.com/beego/bee \
				    github.com/beego/i18n \
				    github.com/beego/wetalk \
				    github.com/codegangsta/gin \
				    github.com/smartystreets/goconvey \
				    ${BACKEND}
assets:
	bundle install
	npm install
	npm run bower
	npm run compile
ustackd:
	go get -u ${BACKEND}
	go install ${BACKEND}
	ustackd --foreground --config '../ustackd/config/ustackd.conf'
test:
	go test ./...
convey:
	goconvey -depth=10 -host="0.0.0.0" -port="8081"
