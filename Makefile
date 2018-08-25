clean:
	rm -rfv bin/*

build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/errorRate errorrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/error error/main.go

deploy: clean build
	sls deploy -v