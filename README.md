1. protoc --version
2. go get -u google.golang.org/grpc
3. go get -u github.com/golang/protobuf/protoc-gen-go
4. protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto