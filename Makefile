jaeger:
	docker-compose up

curl:
	curl -X GET 'http://localhost:8080/api/v1/jaeger'

power:
	ab -v -k -c 2 -n 200 http://localhost:8080/api/v1/jaeger