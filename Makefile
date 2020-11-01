build: main.go
	go build main.go

run: main.go
	go run main.go

test: main_test.go documents.go
	go test main_test.go documents.go