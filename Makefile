default: build

build: clean
	mkdir bin && cd bin && gox ../

clean:
	rm -rf ./bin
