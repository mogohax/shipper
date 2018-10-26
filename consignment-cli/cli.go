package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/mogohax/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return  nil, err
	}

	json.Unmarshal(data, &consignment)

	return consignment, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	response, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("could not create %v", err)
	}
	log.Printf("created %t", response.Created)
}
