FROM golang:1.10.4 as builder

WORKDIR /go/src/github.com/mogohax/shipper/user-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/mogohax/shipper/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]