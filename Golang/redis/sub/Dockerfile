FROM golang:1.17-alpine3.15 AS redis-pub

COPY script.sh /script.sh
COPY sub.go /home/redis/sub.go

CMD ["/script.sh"]