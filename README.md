# Chainstream Go SDK

Official Go client library for Chainstream API.

## Installation

```bash
go get github.com/chainstream-io/chainstream-go-sdk
```

## Usage

```go
import "github.com/chainstream-io/chainstream-go-sdk"

client, err := chainstream.NewDexClient("your-access-token", &chainstream.DexAggregatorOptions{
    ServerUrl: "https://api-dex.chainstream.io",
    StreamUrl: "wss://realtime-dex.chainstream.io/connection/websocket",
})
if err != nil {
    log.Fatal(err)
}
defer client.Close()

// Use the client...
```

## Documentation

For detailed documentation, please visit [https://docs.chainstream.io](https://docs.chainstream.io)

## License

MIT License
