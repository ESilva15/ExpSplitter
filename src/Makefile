build_release_alpine:
	rm -rf build && mkdir -p build/data
	CGO_ENABLED=1 GOOS=linux CC=musl-gcc \
					 go build -ldflags "-X main.ginMode=release" -o ./build/expensesweb
	cp -r ./assets ./build/
	cp ./config.yaml ./build/
	cp ./data/data.db ./build/data/
	find build -name "*.ttf" ! -name DejaVuSansMono.ttf -delete
