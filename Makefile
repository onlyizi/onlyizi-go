MAIN_PATH=cmd/example.go

run:
	go run $(MAIN_PATH)

swag:
	swag init -g cmd/example.go --parseDependency --parseInternal