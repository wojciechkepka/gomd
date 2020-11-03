default: tests styles compile

compile:
	go build -v -o build/gomd .

styles:
	./scripts/add_css.sh

tests:
	go test -v ./...

buildall:
	./scripts/build.sh $(VER)
