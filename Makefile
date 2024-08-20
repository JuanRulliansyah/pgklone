all: build

build:
	go build -o pgklone ./cmd

build-exe:
	go build -o pgklone.exe ./cmd

clean:
	rm -f pgklone