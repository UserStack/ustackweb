BACKEND=github.com/UserStack/ustackd

all: run
watch:
	npm run watch
run:
	bee run
prepare:
	make deps
	make beego_develop
	make assets
prepare_test:
	make deps
	make beego_develop
deps:
	go get -u github.com/astaxie/beego \
						github.com/beego/bee \
				    github.com/beego/i18n \
				    github.com/beego/wetalk \
				    github.com/smartystreets/goconvey \
				    github.com/tools/godep \
				    ${BACKEND}
beego_develop:
	cd $(firstword $(subst :, ,${GOPATH}))/src/github.com/astaxie/beego && git checkout develop
assets:
	bundle install
	npm install
	npm run bower
	npm run compile
ustackd:
	go get -u ${BACKEND}
	go install ${BACKEND}
	make run_ustackd
run_ustackd:
	ustackd --foreground --config '${GOPATH}/github.com/UserStack/ustackd/config/ustackd.conf'
test:
	go test ./tests/...
convey:
	goconvey -depth=10 -host="0.0.0.0" -port="8081"
