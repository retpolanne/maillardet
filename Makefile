# TODO - make this fail by switching - to @
test: clean 
	go test ./... -coverprofile=coverage.out
	cat coverage.out >> coverage.txt

build: clean
	go build main.go

clean:
	#rm main
