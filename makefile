build:
	@go build -o ddacc cmd/main.go

test:
	@go test -v ./... -cover

clean:
	@rm -rf bin