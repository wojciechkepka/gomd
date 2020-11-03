this:
	./scripts/add_css.sh
	go build -o build/gomd .

styles:
	./scripts/add_css.sh

tests:
	go test -v ./...

buildall:
	./scripts/build.sh $(VER)
