# go-victron
[![Audit & Test](https://github.com/koestler/go-victron/actions/workflows/audit.yml/badge.svg)](https://github.com/koestler/go-victron/actions/workflows/audit.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/koestler/go-victron.svg)](https://pkg.go.dev/github.com/koestler/go-victron)

A go implementation of Victron Energy's VE.Direct and BLE protocols.

## Usage
See https://pkg.go.dev/github.com/koestler/go-victron for documentation.

To install the cli run:
```bash
go install github.com/koestler/go-victron/vecli@latest
```

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