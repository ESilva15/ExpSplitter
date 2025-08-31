service = expenses-pgdb
version = v0.0.81-alpha
image = $(service)
container = $(image)-container

.PHONY: clean attach-test test

clean:
	@docker image inspect $(image) > /dev/null 2>&1 && \
	docker image rm $(image) || true

build:
	cp ../expenses/db/schema.sql .
	docker buildx build \
		-t $(image) .

attach:
	docker exec -it $(container) zsh

attach-test:
	docker exec -it $(container)-dev zsh

remove-dev-container:
	docker container rm $(container)-dev --force > /dev/null 2>&1

test: remove-dev-container
	docker run -p 5431:5432 \
		--name $(container)-dev \
		$$(cat .env | sed 's/^/-e /') \
		-v ./pgsql-data-dev/:/var/lib/postgresql/data/ \
		-it $(image)
