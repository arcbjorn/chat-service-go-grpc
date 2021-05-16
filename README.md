## Chat service using gRPC

```bash
# 1nd terminal
go run server.go

# 2nd terminal
go run client.go

# if you change chat.proto -> generate new protobuf
protoc --go_out=plugins=grpc:./ chat.proto
```
