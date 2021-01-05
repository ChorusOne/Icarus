.PHONY: build

build:
	go build -o build/icarus-api -ldflags "-extldflags -pthread" cmd/api.go
	go build -o build/icarus -ldflags "-extldflags -pthread" cmd/icarus.go

build-alpine:
	go build -o build/icarus-api -tags musl -ldflags "-extldflags -pthread" cmd/api.go
	go build -o build/icarus -tags musl -ldflags "-extldflags -pthread" cmd/icarus.go

