# grpc-logger-bin

docker run -d -p 27017:27017 --name loggerbin -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=localdbpass mongo:latest

protoc proto/loggerbin.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.