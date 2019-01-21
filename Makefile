clean: 
	rm main

test: 
	cd pkg && go test ./...

build:
	go build main.go
