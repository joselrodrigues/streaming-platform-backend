generate-mock:
	go generate -v ./...
run:
	docker compose up
build:
	 go build -o server main.go	