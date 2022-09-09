# Crypto price indexer

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
|CICD| gitlab friendly |

## Development

### Runing application
```bash
go run init/main.go \
  grpc \
  --app app.io \
  --port 3000
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

> https://github.com/grpc/grpc-go/blob/master/examples/helloworld/