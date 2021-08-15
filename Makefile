.PHONY: doc
doc:
	swag init


.PHONY: run
run:
	swag init
	go run main.go
