run:
	@echo "Start building..."
	#air --build.cmd "go build -o ./cmd/bin/api tmp/main.go" --build.bin "./cmd/bin/api"
	@air


init:
	@echo "Start init..."
	@go install github.com/air-verse/air@latest
	@go get -u github.com/gin-gonic/gin
