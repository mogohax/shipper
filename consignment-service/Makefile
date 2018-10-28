compile:
		env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

build:
		protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/mogohax/shipper/consignment-service \
				proto/consignment/consignment.proto

		docker build -t consignment-service .

run:
		docker run -p 50051:50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=mdns consignment-service