package implement

import (
	"context"
	"fleet_backend_final/customer_service/proto"
	"github.com/google/uuid"
)

type Service struct {
	Repository *CustomerServiceRepository
}


func (s Service) Signup(ctx context.Context, request *proto.SignUpRequest)(*proto.Driver, error){
	client := &proto.FleetCompany{
		Id:uuid.New().String(),
		Name:request.FleetCompanyName,
	}
	if err := s.Repository.AddFleetCompany(ctx, client); err != nil {
		return nil, err
	}

	driver := &proto.Driver{
		Name: request.Name,
		Id: uuid.New().String(),
		Email:request.Email,
		Password: request.Password,
		FleetCompanyId: client.Id,
	}
	if err := s.Repository.AddDriver(ctx, driver); err != nil {
		return nil, err
	}
	return driver, nil
}


func (s Service) CreateDriver(ctx context.Context, request *proto.Driver)(*proto.Driver, error){


	driver := &proto.Driver{
		Name: request.Name,
		Id: uuid.New().String(),
		Email:request.Email,
		Password: request.Password,
		FleetCompanyId: request.FleetCompanyId,
	}
	if err := s.Repository.AddDriver(ctx, driver); err != nil {
		return nil, err
	}
	return driver, nil
}


func (s Service) UpdateDriver(ctx context.Context, request *proto.Driver)(*proto.Driver, error){

	if err := s.Repository.UpdateDriver(ctx, request); err != nil {
		return nil, err
	}
	return request, nil
}


func (s Service) GetDriverById(ctx context.Context, id string)(*proto.Driver, error){
	if driver, err := s.Repository.GetDriverById(ctx, id); err != nil{
		return nil, err
	}else {
		return driver, nil
	}
}


func (s Service) GetDriversByFleetCompanyId(ctx context.Context, fleetCompanyId string)([] *proto.Driver, error){

	if drivers, err := s.Repository.GetDriversByFleetCompanyId(ctx, fleetCompanyId); err != nil{
		return nil, err
	}else {
		return drivers, nil
	}
}


func (s Service) AddCorporation(ctx context.Context, request *proto.Corporation)(*proto.Corporation, error){


	corporation := &proto.Corporation{
		Name: request.Name,
		Id: uuid.New().String(),
		FleetCompanyId: request.FleetCompanyId,
	}
	if err := s.Repository.AddCorporation(ctx, corporation); err != nil {
		return nil, err
	}
	return corporation, nil
}


func (s Service) GetCorporationById(ctx context.Context, id string)(*proto.Corporation, error){
	if corporation, err := s.Repository.GetCorporationById(ctx, id); err != nil{
		return nil, err
	}else {
		return corporation, nil
	}
}


func (s Service) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetCompanyId string)([] *proto.Corporation, error){

	if corporations, err := s.Repository.GetAllCorporationsByFleetCompanyId(ctx, fleetCompanyId); err != nil{
		return nil, err
	}else {
		return corporations, nil
	}
}


func (s Service) AddRegion(ctx context.Context, request *proto.Region)(*proto.Region, error){

	region := &proto.Region{
		Name: request.Name,
		Id: uuid.New().String(),
		CorporationId: request.CorporationId,
	}
	if err := s.Repository.AddRegion(ctx, region); err != nil {
		return nil, err
	}
	return region, nil
}


func (s Service) GetRegionById(ctx context.Context, id string)(*proto.Region, error){
	if region, err := s.Repository.GetRegionById(ctx, id); err != nil{
		return nil, err
	}else {
		return region, nil
	}
}


func (s Service) GetAllRegionsByCorporationId(ctx context.Context, corporationId string)([] *proto.Region, error){

	if regions, err := s.Repository.GetAllRegionsByCorporationId(ctx, corporationId); err != nil{
		return nil, err
	}else {
		return regions, nil
	}
}


func (s Service) AddDistrict(ctx context.Context, request *proto.District)(*proto.District, error){

	district := &proto.District{
		Name: request.Name,
		Id: uuid.New().String(),
		RegionId: request.RegionId,
	}
	if err := s.Repository.AddDistrict(ctx, district); err != nil {
		return nil, err
	}
	return district, nil
}


func (s Service) GetDistrictById(ctx context.Context, id string)(*proto.District, error){
	if district, err := s.Repository.GetDistrictById(ctx, id); err != nil{
		return nil, err
	}else {
		return district, nil
	}
}


func (s Service) GetAllDistrictsByRegionId(ctx context.Context, regionId string)([] *proto.District, error){

	if districts, err := s.Repository.GetAllDistrictsByRegionId(ctx, regionId); err != nil{
		return nil, err
	}else {
		return districts, nil
	}
}


func (s Service) AddLocation(ctx context.Context, request *proto.Location)(*proto.Location, error){


	location := &proto.Location{
		Name: request.Name,
		Id: uuid.New().String(),
		DistrictId: request.DistrictId,
	}
	if err := s.Repository.AddLocation(ctx, location); err != nil {
		return nil, err
	}
	return location, nil
}


func (s Service) GetLocationById(ctx context.Context, id string)(*proto.Location, error){
	if location, err := s.Repository.GetLocationById(ctx, id); err != nil{
		return nil, err
	}else {
		return location, nil
	}
}


func (s Service) GetAllLocationsByDistrictId(ctx context.Context, DistrictId string)([] *proto.Location, error){

	if locations, err := s.Repository.GetAllLocationsByDistrictId(ctx, DistrictId); err != nil{
		return nil, err
	}else {
		return locations, nil
	}
}