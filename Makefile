test: 
	cd pkg && go test ./...

build:
	go build main.go

clean:
	rm main
