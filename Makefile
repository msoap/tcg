test:
	go test ./...

test-race:
	go test -race ./...

lint:
	go vet ./...
	errcheck ./...

golangci-lint:
	golangci-lint run

example-game-of-life:
	go run examples/game_of_life/game_of_life.go -mode 2x3

example-fractal-demo:
	go run examples/fractal-demo/fractal-demo.go -mode 2x3
