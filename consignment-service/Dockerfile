FROM golang:1.10.4 as builder

WORKDIR /go/src/github.com/mogohax/shipper/consignment-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOS=linux go build -a -installsuffix cgo .

#
#   Build end
#

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/src/github.com/mogohax/shipper/consignment-service/consignment-service .

CMD ["./consignment-service"]