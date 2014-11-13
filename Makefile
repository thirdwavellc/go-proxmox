default: build

build: test clean
	mkdir bin && cd bin && gox ../

test:
	go test

clean:
	rm -rf ./bin
