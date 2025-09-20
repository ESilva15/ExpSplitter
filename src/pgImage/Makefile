include ../vars.mk
service = expenses-pgdb
image = $(service)
container = $(image)-container

.PHONY: clean attach-test test

clean:
	@docker image inspect $(image) > /dev/null 2>&1 && \
	docker image rm $(image) || true

build:
	cp ../expenses/repo/pgdb/schema.sql .
	docker buildx build \
		-t $(registry)/$(repo)/expensesdb:$(tag) .

push_image:
	docker push $(registry)/$(repo)/expensesdb:$(tag)

attach:
	docker exec -it $(container) zsh

attach-test:
	docker exec -it $(container)-dev zsh

remove-dev-container:
	docker container rm $(container)-dev --force > /dev/null 2>&1

test: remove-dev-container
	docker run -p 5431:5432 \
		--name $(container)-dev \
		--env-file .env \
		-v ./pgsql-data-dev/:/var/lib/postgresql/data/ \
		-it $(registry)/$(repo)/expensesdb:$(tag)

build_image:
	make -C ./
