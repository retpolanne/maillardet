test: clean 
	cd pkg && go test ./... -cover

test-coverage: clean 
	-cd pkg && go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

build: clean
	go build main.go

clean:
	#rm main
