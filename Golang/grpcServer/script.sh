#! /bin/sh

cd ..
cd home
ls
cd grpc
go mod init server
go mod tidy
go run server.go
