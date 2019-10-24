#!/bin/bash
protoc -I/usr/local/include -I.  -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --go_out=plugins=grpc:.  ./gateway/app/protobuf/app_gateway/app_gateway.proto

#protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --ruby_out=.   ./gateway/app/protobuf/app_gateway/app_gateway.proto

#protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --plugin=protoc-gen-grpc=grpc_ruby_plugin   --grpc-ruby_out=.   ./gateway/app/protobuf/app_gateway/app_gateway.proto

#protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true:.   ./gateway/app/protobuf/app_gateway/app_gateway.proto

#protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.   ./gateway/app/protobuf/app_gateway/app_gateway.proto

protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./gateway/app/protobuf/app_gateway/app_gateway.yaml:.   ./gateway/app/protobuf/app_gateway/app_gateway.proto