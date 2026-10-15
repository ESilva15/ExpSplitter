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

devdb_start:
	podman compose up database --force-recreate

devdb_clear:
	podman compose down -v database

devdb_status:
	podman compose stats database

devdb_stop:
	podman compose stop database

build_image: build_release_alpine
	podman build -t $(registry)/$(repo)/expensesweb:$(tag) .

launch_container:
	podman run \
		-p 8081:8081 \
		-v ./data/data.db:/usr/share/webapps/expensesweb/data/data.db \
		-it $(registry)/$(repo)/expensesweb:$(tag) zsh

push_image:
	podman push $(registry)/$(repo)/expensesweb:$(tag)

test:
	# go test  ./... -cover -bench=
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm ./coverage.out
