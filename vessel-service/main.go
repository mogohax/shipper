package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/mogohax/shipper/vessel-service/proto/vessel"
)



const (
	defaultHost = "localhost:27017"
)

func CreateDummyData(repo IRepository) {
	defer repo.Close()

	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
}1

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecing to datastore '%s': %v", host, err)
	}

	repo := &VesselRepository{session.Copy()}

	CreateDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}