.PHONY: postgres adminer migrate
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD={pass} postgres

adminer:
	docker run --rm -ti --network host adminer 

migrate:
	migrate -source file://migrations/ \
			-database postgres://postgres:Sc0rpion@localhost/yogo?sslmode=disable up
