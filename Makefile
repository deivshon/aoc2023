debug:
	mkdir -p ./out/debug
	go mod tidy && make lint && go build -o ./out/debug/aoc ./src/cmd/*

release:
	mkdir -p ./out/release
	go mod tidy && go build -ldflags "-w -s" -o ./out/release/aoc ./src/cmd/*

lint:
	golangci-lint run

clean:
	rm -rf ./out

.PHONY: all install uninstall clean
