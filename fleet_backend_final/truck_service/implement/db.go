package implement

import (
	"context"
	"fleet_backend_final/common"
	"fleet_backend_final/truck_service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "Truck_service"
const truckCollection = "Truck"

type TruckServiceRepository struct{
	fleetcompany *mgo.Session
	dbName string
}

func NewTruckServiceRepository()(*TruckServiceRepository, error){
	if session , err := common.ConnectToMongo(); err != nil {
		return nil,err
	}else {
		return &TruckServiceRepository{
			fleetcompany:session,
		},err
	}
}


func (t *TruckServiceRepository) CreateTruck(ctx context.Context, truck *proto.Truck) error{
	if err:= t.truckCollection().Insert(truck); err != nil{
		return err
	}else{
		return nil
	}
}


func (t *TruckServiceRepository) UpdateTruck(ctx context.Context, truckUpdated *proto.Truck) error{
	if _, err:= t.truckCollection().Upsert(bson.M{"id": truckUpdated.Id}, truckUpdated); err != nil{
		return err
	}else{
		return nil
	}
}


func (t *TruckServiceRepository) truckCollection() *mgo.Collection{
	return t.fleetcompany.DB(dbName).C(truckCollection)
}


func (t *TruckServiceRepository) GetTruckById(ctx context.Context, id string)(*proto.Truck, error){
	truck := &proto.Truck{}
	if err := t.truckCollection().Find(bson.M{"id": id}).One(truck); err != nil {
		return nil, err
	}else {
		return truck, nil
	}
}


func (t *TruckServiceRepository) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetcompanyid string)([]*proto.Truck, error){
	var trucks [] *proto.Truck
	if err := t.truckCollection().Find(bson.M{"fleetcompanyid": fleetcompanyid}).All(&trucks); err != nil {
		return nil, err
	}else {
		return trucks , nil
	}
}


func(c *TruckServiceRepository) Close(){
	c.fleetcompany.Close()
}