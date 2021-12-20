FROM golang:1.17-alpine3.15 AS grpc-client

COPY script.sh /script.sh
COPY grpcServer.go /home/grpc/server.go
COPY management /home/grpc/management

CMD ["/script.sh"]