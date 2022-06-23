FROM golang:1.18-alpine AS builder
WORKDIR /go/src/github.com/pentohq/grpc-stub
COPY . .

RUN go build -o cli ./cmd/cli/

FROM alpine
COPY --from=builder /go/src/github.com/pentohq/grpc-stub/cli /cli
ENTRYPOINT /cli