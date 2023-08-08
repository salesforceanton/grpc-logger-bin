# grpc-logger-bin
## gRPC server to collect logs from another microservices system component

### Description:

This is example of creation simple gRPC server to collect logs from another microservices system component and store in MongoDB database instance.
To generate go-files from .proto wer can use command below

```protoc proto/loggerbin.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.```

### Stack:

```
1. GO
2. gRPC
3. MongoDB
```

### Setting up:

Set on your `.env` file variables according with `.env.example` file
In this example I used MongoDB which we can run with Docker Container

```docker run -d -p {DB_PORT:DB_PORT} --name {LOGGERBIN_DB_DATABASE_NAME} -e MONGO_INITDB_ROOT_USERNAME={LOGGERBIN_DB_USERNAME} -e MONGO_INITDB_ROOT_PASSWORD={LOGGERBIN_DB_PASSWORD} mongo:latest```

### Running
To run server use command

```go run cmd/main.go```

