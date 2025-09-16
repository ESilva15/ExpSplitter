include vars.mk
build = ./build/expensesweb

build_release_alpine:
	rm -rf build && mkdir -p $(build)
	CGO_ENABLED=1 GOOS=linux CC=musl-gcc \
		go build -ldflags "-X expenses/cmd.ginMode=release" \
		-o $(build)/expensesweb
	cp -r ./assets $(build)
	cp ./config.yaml $(build)
	mkdir -p $(build)/expenses/repo/pgdb/migrations/ && \
		cp -r ./expenses/repo/pgdb/migrations/ $(build)/expenses/repo/pgdb/
	mkdir -p $(build)/data/
	find $(build) -name "*.ttf" ! -name DejaVuSansMono.ttf -delete

build_image: build_release_alpine
	docker buildx build -t $(registry)/$(repo)/expensesweb:$(tag) .

launch_container:
	docker run \
		-p 8081:8081 \
		-v ./data/data.db:/usr/share/webapps/expensesweb/data/data.db \
		-it $(registry)/$(repo)/expensesweb:$(tag) zsh

push_image:
	docker push $(registry)/$(repo)/expensesweb:$(tag)

test:
	# go test  ./... -cover -bench=
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm ./coverage.out

build_db_image:
	make -C ./pgImage build

test_db:
	make -C ./pgImage test
