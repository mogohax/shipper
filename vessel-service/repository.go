package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
import pb "github.com/mogohax/shipper/vessel-service/proto/vessel"

const (
	dbName = "shipper"
	vesselCollection = "vessels"
)

type IRepository interface {
	FindAvailable(spec *pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	// Here we define a more complex query than our consignment-service's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than and equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.

	err := repo.collection().Find(bson.M{
		"capacity": bson.M{ "$gte": spec.Capacity},
		"maxweight": bson.M{ "$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

