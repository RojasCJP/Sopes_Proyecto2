FROM golang:1.17-alpine3.15 AS redis-pub

COPY script.sh /script.sh
COPY pub.go /home/redis/pub.go

CMD ["/script.sh"]