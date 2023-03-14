
.PHONY: run
run:
	go run .

.PHONY: cp-env
cp-env:
	cp ./docs/app.env.example ./app.env

.PHONY: curl-insert	
insert-record:
	curl -X POST -H "Content-Type: application/json" -d '{"npm":5119992, "name":"student1"}' http://localhost:8000/students


