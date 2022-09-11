# Crypto Block Indexer

## Feature

| Title | Value |
|:-----:|:-----:|
|language| golang (v1.18) |
|architecture| full clean arch implemented |
|transport| gRPC |
|db| clickhouse |
|container| docker |
|orchestrator| docker-compose |
|documentation| openAPI |

## Development

### Runing application

```bash
go run init/main.go grpc --port 3000 --host 0.0.0.0 --env development
```

## Use in prodution

```bash
docker build -t app:1.0.0 .
docker run \
  --name app
  -p 3000:3000 \
  -v ${PWD}/config.yaml:/app/config/config.yaml \
  app:1.0.0
```

## Usage

### last block data

```bash
grpcurl -plaintext localhost:3000 io.app.price.api.BlockService/GetBlock
```

### Specific block
```bash
grpcurl -plaintext -d '{"blockNumber":15509900}' localhost:3000 io.app.price.api.BlockService/GetBlock
```

### Stat of blocks

```bash
grpcurl -plaintext localhost:3000 io.app.price.api.StatsService/Stat
```

### Stat of between two block range

```bash
grpcurl -plaintext -d '{"startBlock": 5000000, "endBlock": 15000000000}' localhost:3000 io.app.price.api.StatsService/Stat
```

### Get last transaction

```bash
grpcurl -plaintext localhost:3000 io.app.price.api.TransactionService/GetTransaction
```

## Test

```bash
go test ./internal/transport/grpc
go test internal/pkg/networkScaning/sync_network
```