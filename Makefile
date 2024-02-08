up:
	docker-compose up -d --remove-orphan --force-recreate
down:
	docker-compose down
init:
	make up
	sleep 2
	@make db-init
db-init:
	docker cp ./mysql/init.sql gotutor-mysql:/init.sql
	docker exec -ti gotutor-mysql sh -c 'echo "source /init.sql" | mysql -u user --password=secret-hello123 --database=database'
db-console:
	docker exec -ti gotutor-mysql sh -c 'mysql -u user --password=secret-hello123 --database=database'
db:
	docker exec -ti gotutor-mysql sh