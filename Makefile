try-go:
	docker run --rm -v "$PWD"/go:/go/data-race golang:1.16-alpine sh -c 'cd /go/data-race && go test -v'