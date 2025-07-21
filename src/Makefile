build_release_alpine:
	rm -rf build && mkdir build
	GIN_MODE=RELEASE CGO_ENABLED=1 GOOS=linux CC=musl-gcc \
					 go build -o ./build/expensesweb
	cp -r ./assets ./build/
	cp ./config.yaml ./build/
	cp ./data/data.db ./build/
	find build -name "*.ttf" ! -name DejaVuSansMono.ttf -delete
