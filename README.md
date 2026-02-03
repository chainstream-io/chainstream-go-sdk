# ChainStream Go SDK

Official Go client library for ChainStream API.

## Installation

```bash
go get github.com/chainstream-io/chainstream-go-sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    chainstream "github.com/chainstream-io/chainstream-go-sdk"
)

func main() {
    client, err := chainstream.NewClient(chainstream.ClientOptions{
        AccessToken: "your-access-token",
    })
    if err != nil {
        panic(err)
    }

    // Use the client for API calls...
    _ = client
}
```

## Documentation

For detailed documentation, visit [https://docs.chainstream.io](https://docs.chainstream.io)

## Development

```bash
# Run tests
make test

# Lint
make lint

# Generate OpenAPI client
make client
```

## License

MIT
