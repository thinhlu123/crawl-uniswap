run:
	go mod vendor
	go mod tidy
	go vet -all
	go run main.go