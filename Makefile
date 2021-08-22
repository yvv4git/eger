app_up:
	docker-compose up --scale k6=0

app_down:
	docker-compose down

k6_test_jaeger:
	docker-compose run k6 run /scripts/eger.js

k6_test_smoke:
	docker-compose run k6 run /scripts/type_smoke.js

k6_test_load_simple:
	docker-compose run k6 run /scripts/type_load_simple.js

k6_test_load_normal:
	docker-compose run k6 run /scripts/type_load_normal_day.js

k6_test_stress_long:
	docker-compose run k6 run /scripts/type_stress.js

k6_test_stress_spike:
	docker-compose run k6 run /scripts/type_stress_spike.js

k6_test_graceful_stop:
	docker-compose run k6 run /scripts/graceful_stop.js

k6_test_stress_soak:
	docker-compose run k6 run /scripts/type_stress_soak.js

curl:
	curl -X GET 'http://localhost:8080/api/v1/jaeger'