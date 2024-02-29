.PHONY: all build clean


all: build



build:
	@go build -o ./bin/ ./app1/

clean:
	@rm -f ./bin/*


