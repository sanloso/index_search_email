run:
	go run main.go

deps:
	go install github.com/swaggo/swag/cmd/swag@latest && \
	goproxy=direct \
	GOSUMDB=off \
	go mod tidy