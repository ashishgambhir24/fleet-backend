package implement

import (
	"context"
	"fleet_backend_final/common"
	"fleet_backend_final/customer_service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "customer"
const driverCollection = "drivers"
const fleetcompanyCollection = "fleetcompany"
const corporationCollection = "Corporation"
const regionCollection = "Region"
const districtCollection = "District"
const locationCollection = "Location"

type CustomerServiceRepository struct {
	fleetcompany *mgo.Session
	dbName string
}


func NewCustomerServiceRepository()(*CustomerServiceRepository, error){
	if session , err := common.ConnectToMongo(); err != nil {
		return nil,err
	}else {
		return &CustomerServiceRepository{
			fleetcompany:session,
		},err
	}
}


func (c *CustomerServiceRepository) AddDriver(ctx context.Context, driver *proto.Driver) error{
	if err:= c.driverCollection().Insert(driver); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) UpdateDriver(ctx context.Context, driverUpdated *proto.Driver) error{
	if _, err:= c.driverCollection().Upsert(bson.M{"id": driverUpdated.Id}, driverUpdated); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) AddFleetCompany(ctx context.Context, fleetcompany *proto.FleetCompany) error{
	if err:= c.fleetcompanyCollection().Insert(fleetcompany); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) driverCollection() *mgo.Collection{
	return c.fleetcompany.DB(dbName).C(driverCollection)
}


func (c *CustomerServiceRepository) fleetcompanyCollection() *mgo.Collection{
	return c.fleetcompany.DB(dbName).C(fleetcompanyCollection)
}


func (c *CustomerServiceRepository) GetDriverById(ctx context.Context, id string)(*proto.Driver, error){
	driver := &proto.Driver{}
	if err := c.driverCollection().Find(bson.M{"id": id}).One(driver); err != nil {
		return nil, err
	}else {
		return driver, nil
	}
}


func (c *CustomerServiceRepository) GetDriversByFleetCompanyId(ctx context.Context, fleetcompanyid string)([]*proto.Driver, error){
	var drivers [] *proto.Driver
	if err := c.driverCollection().Find(bson.M{"fleetcompanyid": fleetcompanyid}).All(&drivers); err != nil {
		return nil, err
	}else {
		return drivers , nil
	}
}


func (c *CustomerServiceRepository) AddCorporation(ctx context.Context, corporation *proto.Corporation) error{
	if err:= c.corporationCollection().Insert(corporation); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) GetCorporationById(ctx context.Context, id string)(*proto.Corporation, error){
	corporation := &proto.Corporation{}
	if err := c.corporationCollection().Find(bson.M{"id": id}).One(corporation); err != nil {
		return nil, err
	}else {
		return corporation, nil
	}
}


func (c *CustomerServiceRepository) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetcompanyId string)([]*proto.Corporation, error){
	var corporations [] *proto.Corporation
	if err := c.corporationCollection().Find(bson.M{"fleetcompanyid": fleetcompanyId}).All(&corporations); err != nil {
		return nil, err
	}else {
		return corporations , nil
	}
}


func (c *CustomerServiceRepository) corporationCollection() *mgo.Collection{
	return c.fleetcompany.DB(dbName).C(corporationCollection)
}


func (c *CustomerServiceRepository) AddRegion(ctx context.Context, region *proto.Region) error{
	if err:= c.regionCollection().Insert(region); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) GetRegionById(ctx context.Context, id string)(*proto.Region, error){
	region := &proto.Region{}
	if err := c.regionCollection().Find(bson.M{"id": id}).One(region); err != nil {
		return nil, err
	}else {
		return region, nil
	}
}


func (c *CustomerServiceRepository) GetAllRegionsByCorporationId(ctx context.Context, corporationId string)([]*proto.Region, error){
	var regions [] *proto.Region
	if err := c.regionCollection().Find(bson.M{"corporationid": corporationId}).All(&regions); err != nil {
		return nil, err
	}else {
		return regions , nil
	}
}


func (c *CustomerServiceRepository) regionCollection() *mgo.Collection{
	return c.fleetcompany.DB(dbName).C(regionCollection)
}


func (c *CustomerServiceRepository) AddDistrict(ctx context.Context, district *proto.District) error{
	if err:= c.districtCollection().Insert(district); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) GetDistrictById(ctx context.Context, id string)(*proto.District, error){
	district := &proto.District{}
	if err := c.districtCollection().Find(bson.M{"id": id}).One(district); err != nil {
		return nil, err
	}else {
		return district, nil
	}
}


func (c *CustomerServiceRepository) GetAllDistrictsByRegionId(ctx context.Context, regionId string)([]*proto.District, error){
	var districts [] *proto.District
	if err := c.districtCollection().Find(bson.M{"regionid": regionId}).All(&districts); err != nil {
		return nil, err
	}else {
		return districts , nil
	}
}


func (c *CustomerServiceRepository) districtCollection() *mgo.Collection{
	return c.fleetcompany.DB(dbName).C(districtCollection)
}


func (c *CustomerServiceRepository) AddLocation(ctx context.Context, location *proto.Location) error{
	if err:= c.locationCollection().Insert(location); err != nil{
		return err
	}else{
		return nil
	}
}


func (c *CustomerServiceRepository) GetLocationById(ctx context.Context, id string)(*proto.Location, error){
	location := &proto.Location{}
	if err := c.locationCollection().Find(bson.M{"id": id}).One(location); err != nil {
		return nil, err
	}else {
		return location, nil
	}
}


func (c *CustomerServiceRepository) GetAllLocationsByDistrictId(ctx context.Context, districtId string)([]*proto.Location, error){
	var locations [] *proto.Location
	if err := c.locationCollection().Find(bson.M{"districtid": districtId}).All(&locations); err != nil {
		return nil, err
	}else {
		return locations , nil
	}
}


func (c *CustomerServiceRepository) locationCollection() *mgo.Collection{
	return c.fleetcompany.DB(dbName).C(locationCollection)
}


func(c *CustomerServiceRepository) Close(){
	c.fleetcompany.Close()
}