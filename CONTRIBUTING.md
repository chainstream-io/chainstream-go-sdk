# Contributing to ChainStream Go SDK

## Development Setup

### Prerequisites

- Go 1.23 or higher
- Make

### Setup

```bash
cd go
make clean
```

### Development Commands

```bash
# Run tests
make test

# Lint code
make lint

# Fix lint issues
make lint-fix

# Generate OpenAPI client
make client

# Generate documentation
make docs
```

## Code Style

- Follow Go idioms and conventions
- Use `golangci-lint` for linting
- Format code with `gofmt`

## Pull Request Process

1. Create a feature branch
2. Make your changes
3. Run tests and linting
4. Submit a pull request

## License

MIT
