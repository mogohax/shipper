package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	vesselProto "github.com/mogohax/shipper/vessel-service/proto/vessel"
	pb "github.com/mogohax/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore whith host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
