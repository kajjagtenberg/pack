run:
	go run main.go

tidy:
	go mod tidy

dev:
	docker run --rm -it -v ${PWD}:/home golang