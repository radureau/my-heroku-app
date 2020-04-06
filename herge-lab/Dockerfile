FROM golang:1.14 AS builder

RUN umask 755
RUN apt-get install -y ca-certificates

RUN mkdir /app
ADD . /app/
WORKDIR /app

ARG GITHUB_ACCESS_TOKEN
RUN apt-get install -y git
RUN git config --global url."https://radureau:$GITHUB_ACCESS_TOKEN@github.com".insteadOf "https://github.com"

RUN go env -w GOPRIVATE=github.com/radureau/my-heroku-app-pkg

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o ./herge-lab main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/herge-lab /

ENTRYPOINT ["/herge-lab"]
