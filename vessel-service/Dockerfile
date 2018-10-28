#
# Build
#
FROM golang:1.10.4 as builder

WORKDIR /go/src/github.com/mogohax/shipper/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN make compile


#
# Run
#
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/src/github.com/mogohax/shipper/vessel-service/vessel-service .

CMD ["./vessel-service"]