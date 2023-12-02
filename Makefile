debug:
	mkdir -p ./out/debug
	go mod tidy && make lint && go build -o ./out/debug/aoc ./src/cmd/main.go

release:
	mkdir -p ./out/release
	go mod tidy && go build -ldflags "-w -s" -o ./out/release/aoc ./src/cmd/main.go

lint:
	golangci-lint run

clean:
	rm -rf ./out

.PHONY: all install uninstall clean
