server:
	go run main.go

sqlc:
	sqlc generate


.PHONY: server sqlc