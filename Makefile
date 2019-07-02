MAKEFLAGS += --no-print-directory
GOOS=linux
GOARCH=amd64

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o thanks-server

publish:
	@make build
	heroku container:push web
	heroku container:release web
