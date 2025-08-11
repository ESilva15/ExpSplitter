build = ./build/expensesweb
registry = dir.esilva.org
repo = esilva
tag = v0.0.51-alpha

build_release_alpine:
	rm -rf build && mkdir -p $(build)
	CGO_ENABLED=1 GOOS=linux CC=musl-gcc \
		go build -ldflags "-X expenses/cmd.ginMode=release" \
		-o $(build)/expensesweb
	cp -r ./assets $(build)
	cp ./config.yaml $(build)
	mkdir -p $(build)/expenses/db/migrations/ && \
		cp -r ./expenses/db/migrations/ $(build)/expenses/db/migrations/
	mkdir -p $(build)/data/
	find $(build) -name "*.ttf" ! -name DejaVuSansMono.ttf -delete

build_image:
	docker buildx build -t $(registry)/$(repo)/expensesweb:$(tag) .

launch_container:
	docker run \
		-p 8081:8081 \
		-v ./data/data.db:/usr/share/webapps/expensesweb/data/data.db \
		-it $(registry)/$(repo)/expensesweb:$(tag) zsh

push_image:
	docker push $(registry)/$(repo)/expensesweb:$(tag)
