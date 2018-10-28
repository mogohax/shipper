FROM golang:1.10.4 as builder

WORKDIR /go/src/github.com/mogohax/shipper/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN make compile

#
#   Build end
#

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/src/github.com/mogohax/shipper/user-service/user-service .

CMD ["./user-service"]