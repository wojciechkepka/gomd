this:
	./scripts/add_assets.sh
	go build -o build/gomd .

styles:
	./scripts/add_assets.sh

buildall:
	./scripts/build.sh $(VER)
