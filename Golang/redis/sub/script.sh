#! /bin/sh

cd ..
cd home
ls
cd redis
go mod init redis
go mod tidy
go run sub.go