default: styles pack tests compile

styles:
	./scripts/add_css.sh

tests:
	go test -v ./...

pack:
	pkger -include /assets/html -include /assets/js

compile:
	go build -v -o build/gomd .

run: pack compile
	./build/gomd

dbg: pack compile
	./build/gomd --debug

buildall:
	./scripts/build.sh $(VER)
