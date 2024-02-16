build:
	@go build -o bin/ddacc cmd/ddacc.go

test:
	@go test -v ./tests/...

clean:
	@rm -rf bin