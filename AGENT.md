# AGENT.md

## Build/Test Commands
- **Build**: `go build -o init_project.exe .` (create executable)
- **Run**: `go run main.go` or `./init_project.exe`
- **Test**: `go test ./...` (run all tests)
- **Single Test**: `go test -run TestName ./package`
- **Check**: `go vet ./...` && `go fmt ./...`

## Architecture
- **Main**: Single `main.go` with algorithm demos, CLI tools, Gin server, GORM examples
- **DataLib**: Database models and operations (`dataLib/` - Student, Account, Transactions)
- **Database**: MySQL with GORM ORM (connection: root:123@tcp(127.0.0.1:3306)/new_schema)
- **Web**: Gin HTTP server (optional, port 8080)
- **CLI**: urfave/cli v2 calculator commands

## Code Style
- **Package**: `github.com/zxNewBee/init_project`
- **Imports**: Standard lib first, then third-party, then local (dataLib)
- **Naming**: CamelCase for public, camelCase for private
- **Structs**: GORM tags for DB models (`gorm:"primarykey"`, `gorm:"not null"`)
- **Errors**: Return errors with descriptive messages in Chinese
- **Functions**: Single responsibility, clear naming (e.g., `findElement`, `TransferFunds`)
- **Database**: Use transactions for multi-step operations, row-level locking for concurrent access
