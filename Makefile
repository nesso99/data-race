try-javascript:
	docker run --rm -v ${PWD}/javascript:/app node:14-alpine sh -c 'cd /app && node data-race.js'

try-go:
	docker run --rm -v ${PWD}/go:/go/data-race golang:1.16-alpine sh -c 'cd /go/data-race && go test -v'