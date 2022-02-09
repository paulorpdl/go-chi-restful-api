run_service:
	go run .

install_deps:
	rm -f go.mod go.sum
	go mod init github.com/newline-sandbox/go-chi-restful-api
	go mod tidy

test:
	go test -v ./...

build:
	CGO_ENABLED=0 GOOS='linux' GOARCH='amd64' go build -v -o _output/app