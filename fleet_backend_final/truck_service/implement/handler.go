package implement

import (
	"context"
	proto2 "fleet_backend_final/common/proto"
	"fleet_backend_final/truck_service/proto"
)

type Handler struct{
Service Service
}


func (h Handler) CreateTruck(ctx context.Context, req *proto.Truck, rsp *proto.TruckResponse) error {
	if truck, err := h.Service.CreateTruck(ctx, req); err != nil{
		return err
	}else{
		rsp.Truck = truck;
		return nil
	}
}


func (h Handler) UpdateTruck(ctx context.Context, req *proto.Truck, rsp *proto.TruckResponse) error {
	if truck, err := h.Service.UpdateTruck(ctx, req); err != nil{
		return err
	}else{
		rsp.Truck = truck;
		return nil
	}
}


func (h Handler) GetTruckById(ctx context.Context, req *proto2.IdRequest, rsp *proto.TruckResponse) error {
	if truck, err := h.Service.GetTruckById(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Truck = truck;
		return nil
	}
}


func (h Handler) GetAllTrucksByFleetCompanyId(ctx context.Context, req *proto2.IdRequest, rsp *proto.TrucksResponse) error {
	if trucks, err := h.Service.GetAllTrucksByFleetCompanyId(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Trucks = trucks;
		return nil
	}
}


func (h Handler) ClockIn(ctx context.Context, req *proto.ClockOperation, rsp *proto.TruckResponse) error {
	if truck, err := h.Service.ClockIn(ctx, req); err!=nil{
		return err
	}else{
		rsp.Truck = truck;
		return nil
	}
}


func (h Handler) ClockOut(ctx context.Context, req *proto.ClockOperation, rsp *proto.TruckResponse) error {
	if truck, err := h.Service.ClockOut(ctx, req); err!=nil{
		return err
	}else{
		rsp.Truck = truck;
		return nil
	}
}