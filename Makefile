default: styles pack tests compile

styles:
	./scripts/add_css.sh

tests:
	go test -v ./...

pack:
	pkger -include /assets/html -include /assets/js

build:
	go build -v -o build/gomd .

run: pack build
	./build/gomd

dbg: pack build
	./build/gomd --debug

buildall:
	./scripts/build.sh $(VER)
