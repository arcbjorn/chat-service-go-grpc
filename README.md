## Chat service using gRPC

Make sure Protoc compiler (protobuf), protoc-go-gen package are installed on your system.

```bash
# 1nd terminal
go run server.go

# 2nd terminal
go run client.go

# if you change chat.proto -> generate new protobuf
protoc --go_out=plugins=grpc:./ chat.proto
```
