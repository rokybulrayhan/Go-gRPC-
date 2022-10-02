1. protoc --version
2. go get -u google.golang.org/grpc
3. go get -u github.com/golang/protobuf/protoc-gen-go
4. protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto


http://localhost:8080/mult/5/6
http://localhost:8080/add/5/6




/*

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    service.proto




*/