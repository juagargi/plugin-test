.PHONY: all build build-app1 build-plugins clean


all: build

build: build-app1 build-plugins

build-app1:
	@go build -o ./bin/ ./app1/

build-plugins: external-plugins
	@for dir in ./plugins/*/; do \
		echo "building $$dir" ; \
		(cd $$dir && go build -buildmode=plugin -o ../../bin/plugins/ . ); \
	done

external-plugins:
	$(foreach external,$(shell cat external_plugins.txt),\
		@cp $(external) ./bin/plugins/; \
	)

clean:
	@find ./bin/ -type f ! -name ".*" -exec rm -f {} +
	@find ./bin/plugins -type f ! -name ".*" -exec rm -f {} +


run: build
	@./bin/app1

