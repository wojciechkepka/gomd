default: styles pack compile

compile:
	go build -v -o build/gomd .

styles:
	./scripts/add_css.sh

tests:
	go test -v ./...

pack:
	pkger -include /assets/html -include /assets/js

buildall:
	./scripts/build.sh $(VER)
