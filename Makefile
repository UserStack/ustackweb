USTACKD=github.com/UserStack/ustackd

# main tools
all: run
build:
	godep go build
run:
	bee run
watch:
	npm run watch
assets:
	npm run bower
	npm run compile
	make favicon

# restore
restore:
	godep restore

# setup project environment including dependencies and pre-compilation
prepare:
	make prepare_go
	make prepare_assets
prepare_go:
	make setup_prod
	make setup_dev
prepare_assets:
	make setup_assets
	make assets

# install dependencies
setup_prod:
	go get github.com/tools/godep
setup_dev:
	go get github.com/smartystreets/goconvey
	go get github.com/beego/bee
setup_assets:
	gem install bundler
	bundle install
	npm install
setup_ustackd:
	go get -u ${USTACKD}
	go install ${USTACKD}

# tools
ustackd:
	ustackd --foreground --config '${GOPATH}/src/${USTACKD}/config/ustackd.conf'
test:
	godep go test ./tests/...
convey:
	goconvey -depth=10 -host="0.0.0.0" -port="8081"
favicon:
	cd static/images && png2ico logo.ico logo16.png logo32.png logo48.png logo64.png
