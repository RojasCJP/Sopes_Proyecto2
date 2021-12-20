#! /bin/sh

cd ..
cd home
ls
cd grpc
go mod init client
go mod tidy
go run client.go
