FROM golang:1.13 AS builder

RUN umask 755
RUN apt-get install -y ca-certificates

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./herge-lab.sh herge-lab/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/herge-lab.sh /

ENTRYPOINT ["/herge-lab.sh"]