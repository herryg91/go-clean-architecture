build: clean
	docker run -d -p 33062:3306 -v ${PWD}/migrations/mysql/:/docker-entrypoint-initdb.d --name book-mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest --default-authentication-plugin=mysql_native_password
	docker run -d -p 33792:6379 --name book-redis -d redis
clean:
	docker rm -f book-mysql || echo 'no mysql docker clean needed'
	docker rm -f book-redis || echo 'no mysql docker clean needed'
test: build
	go test ./...