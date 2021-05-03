build:
	docker run -d -p 33061:3306 -v ${PWD}/migrations/mysql/:/docker-entrypoint-initdb.d --name video-mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest --default-authentication-plugin=mysql_native_password
clean:
	docker stop video-mysql
	docker rm video-mysql


