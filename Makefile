.PHONY: postgres adminer migrate
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD={Password} postgres

adminer:
	docker run --rm -ti --network host adminer 

migrate:
	migrate -source file://migrations/ \
			-database postgres://postgres:{Password}@localhost/yogo?sslmode=disable up
dropdb:
	migrate -source file://migrations/ \
			-database postgres://postgres:{Password}@localhost/yogo?sslmode=disable down