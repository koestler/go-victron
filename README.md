# go-victron
[![Audit & Test](https://github.com/koestler/go-victron/actions/workflows/audit.yml/badge.svg)](https://github.com/koestler/go-victron/actions/workflows/audit.yml)

A go implementatino of Victron Energy's VE.Direct and BLE protocols.

## Development

Install local dependencies / build tools:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
go install go.uber.org/mock/mockgen@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

Build test and audit the code:
```bash
go generate ./...
go build ./...
go vet ./...
go test ./...
staticcheck ./...
```