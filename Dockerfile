FROM golang:1.19.4-alpine3.16 AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/productCatalog
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev
# RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates gcc libtool make musl-dev

COPY Makefile go.mod go.sum ./
# RUN make init # gRPC related protobuf commands
RUN go mod download
# RUN make proto
COPY . .
RUN make tidy build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/productCatalog/productCatalog /productCatalog
ENTRYPOINT ["/productCatalog"]
CMD []
