con-run: 
	@docker run --name=mysql \
		-p 3306:3306 \
		-e MYSQL_ROOT_PASSWORD=pass \
		-v mysql:/var/lib/mysql \
		-d mysql:latest

con-exec:
	@docker exec -it mysql mysql -uroot -p

stop:
	@docker stop mysql

con-rm:
	@docker container rm mysql

install-goose:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

load-env:
	export GOOSE_DRIVER="mysql" && \
	export GOOSE_DB_STRING="mysql://root:pass@tcp(127.0.0.1:3306)" && \
	export GOOSE_MIGRATION_DIR="./migrations/"

up:
	@goose up

down:
	@goose down

serve:
	@go run .
