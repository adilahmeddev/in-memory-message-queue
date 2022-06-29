DEFAULT_GOAl: test-and-push

test:
	@go test ./...

test-and-push: test push

push:
	@git push