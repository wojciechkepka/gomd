this:
	./scripts/add_css.sh
	go build -o build/gomd .

styles:
	./scripts/add_css.sh

buildall:
	./scripts/build.sh $(VER)
