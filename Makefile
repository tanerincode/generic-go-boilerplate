serve-application:
	./bin/application

build: clean
	go build -ldflags="-s -w" -o ./bin/application cmd/application/main.go

clean:
	go clean
	-rm ./bin/*