build: clean
	docker run -d -p 33061:3306 -v ${PWD}/migrations/mysql/:/docker-entrypoint-initdb.d --name video-mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest --default-authentication-plugin=mysql_native_password
clean:
	docker rm -f video-mysql || echo 'no clean needed'
test: build
	go test ./...