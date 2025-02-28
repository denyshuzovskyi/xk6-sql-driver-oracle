all: test build example

test: *.go testdata/*.js
	go test -count 1 ./...

build: k6

k6: *.go go.mod go.sum
	CGO_ENABLED=1 xk6 build --with github.com/grafana/xk6-sql@latest --with github.com/denyshuzovskyi/xk6-sql-driver-oracle=.

example: k6
	./k6 run examples/example.js

.PHONY: test all example