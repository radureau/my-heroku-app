FROM golang:1.13 AS builder

RUN umask 755
RUN apt-get install -y ca-certificates

RUN mkdir /app
ADD ./herge-lab /app/
WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV buildcmd="go build -o ./herge-lab main.go"
RUN $buildcmd -mod vendor || $buildcmd

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/herge-lab /

ENTRYPOINT ["/herge-lab"]