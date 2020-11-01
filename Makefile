build: main.go documents.go
	go build main.go documents.go

run: main.go documents.go
	go run main.go documents.go

test: main_test.go documents.go
	go test main_test.go documents.go