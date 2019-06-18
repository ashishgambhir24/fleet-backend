package implement

import (
	"context"
	"fleet_backend_final/truck_service/proto"
	"github.com/google/uuid"
)

type Service struct{
	Repository *TruckServiceRepository
}

func (s Service) CreateTruck(ctx context.Context, request *proto.Truck)(*proto.Truck, error){


	truck := &proto.Truck{
		LicensePlate: request.LicensePlate,
		Id: uuid.New().String(),
		ClockedInUser:request.ClockedInUser,
		Miles: request.Miles,
		FleetCompanyId: request.FleetCompanyId,
		CorporationId: request.CorporationId,
		RegionId: request.RegionId,
		DistrictId: request.DistrictId,
		LocationId: request.LocationId,
	}
	if err := s.Repository.CreateTruck(ctx, truck); err != nil {
		return nil, err
	}
	return truck, nil
}

func (s Service) UpdateTruck(ctx context.Context, request *proto.Truck)(*proto.Truck, error){

	if err := s.Repository.UpdateTruck(ctx, request); err != nil {
		return nil, err
	}
	return request, nil
}


func (s Service) GetTruckById(ctx context.Context, id string)(*proto.Truck, error){
	if truck, err := s.Repository.GetTruckById(ctx, id); err != nil{
		return nil, err
	}else {
		return truck, nil
	}
}


func (s Service) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string)([] *proto.Truck, error){

	if trucks, err := s.Repository.GetAllTrucksByFleetCompanyId(ctx, fleetCompanyId) ; err != nil{
		return nil, err
	}else {
		return trucks, nil
	}
}


func (s Service) ClockIn(ctx context.Context, request *proto.ClockOperation)(*proto.Truck, error){

	if truck, err := s.Repository.GetTruckById(ctx, request.TruckId); err != nil{
		return nil, err
	}else{
		if truck.ClockedInUser != ""{
			return truck, nil
		}else{

			truckupdated := &proto.Truck{
				LicensePlate: truck.LicensePlate,
				Id: truck.Id,
				ClockedInUser:request.DriverId,
				Miles: truck.Miles,
				FleetCompanyId: truck.FleetCompanyId,
				CorporationId: truck.CorporationId,
				RegionId: truck.RegionId,
				DistrictId: truck.DistrictId,
				LocationId: truck.LocationId,
			}
			if err2 := s.Repository.UpdateTruck(ctx, truckupdated); err2 != nil {
				return nil, err2
			}
			return truckupdated, nil
		}
	}

}


func (s Service) ClockOut(ctx context.Context, request *proto.ClockOperation)(*proto.Truck, error){

	if truck, err := s.Repository.GetTruckById(ctx, request.TruckId); err != nil{
		return nil, err
	}else{

		truckupdated := &proto.Truck{
			LicensePlate: truck.LicensePlate,
			Id: truck.Id,
			ClockedInUser:"",
			Miles: truck.Miles,
			FleetCompanyId: truck.FleetCompanyId,
			CorporationId: truck.CorporationId,
			RegionId: truck.RegionId,
			DistrictId: truck.DistrictId,
			LocationId: truck.LocationId,
		}
		if err2 := s.Repository.UpdateTruck(ctx, truckupdated); err2 != nil {
			return nil, err2
		}
		return truckupdated, nil
	}
}
