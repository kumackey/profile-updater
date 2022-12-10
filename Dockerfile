FROM golang:1.18-alpine

WORKDIR /go/src/app
COPY . .
RUN go build -o profile ./cmd

ENTRYPOINT ["/go/src/app/profile"]