build = ./build/expensesweb
registry = dir.esilva.org
repo = esilva
tag = v0.0.46-alpha

build_release_alpine:
	rm -rf build && mkdir -p $(build)
	CGO_ENABLED=1 GOOS=linux CC=musl-gcc \
					 go build -ldflags "-X main.ginMode=release" -o $(build)
	cp -r ./assets $(build)
	cp ./config.yaml $(build)
	mkdir $(build)/data/
	find $(build) -name "*.ttf" ! -name DejaVuSansMono.ttf -delete

build_image:
	docker buildx build -t $(registry)/$(repo)/expensesweb:$(tag) .

push_image:
	docker push $(registry)/$(repo)/expensesweb:$(tag)
