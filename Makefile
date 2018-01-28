test:
	go test .

test-race:
	go test -race .

lint:
	golint .
	go vet .
	errcheck .

gometalinter:
	gometalinter --vendor --cyclo-over=20 --line-length=150 --dupl-threshold=150 --min-occurrences=2 --enable=misspell --deadline=10m ./...
