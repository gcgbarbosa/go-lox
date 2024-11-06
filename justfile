build:
  go run main.go

test:
  go test ./...

pack:
  llm-prepare -p . -f "*.go" -o "go-lox.txt"
