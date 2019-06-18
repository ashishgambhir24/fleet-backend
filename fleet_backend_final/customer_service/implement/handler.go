package implement

import (
	"context"
	proto2 "fleet_backend_final/common/proto"
	"fleet_backend_final/customer_service/proto"
)

type Handler struct{
	Service Service
}






func (h Handler) SignUp(ctx context.Context, req *proto.SignUpRequest, rsp *proto.DriverResponse) error {
	if driver, err := h.Service.Signup(ctx, req); err != nil{
		return err
	}else{
		rsp.Driver = driver;
		return nil
	}
}


func (h Handler) CreateDriver(ctx context.Context, req *proto.Driver, rsp *proto.DriverResponse) error {
	if driver, err := h.Service.CreateDriver(ctx, req); err != nil{
		return err
	}else{
		rsp.Driver = driver;
		return nil
	}
}


func (h Handler) UpdateDriver(ctx context.Context, req *proto.Driver, rsp *proto.DriverResponse) error {
	if driver, err := h.Service.UpdateDriver(ctx, req); err != nil{
		return err
	}else{
		rsp.Driver = driver;
		return nil
	}
}


func (h Handler) GetDriverById(ctx context.Context, req *proto2.IdRequest, rsp *proto.DriverResponse) error {
	if driver, err := h.Service.GetDriverById(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Driver = driver;
		return nil
	}
}


func (h Handler) GetDriversByFleetCompanyId(ctx context.Context, req *proto2.IdRequest, rsp *proto.DriversResponse) error {
	if drivers, err := h.Service.GetDriversByFleetCompanyId(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Drivers = drivers;
		return nil
	}
}


func (h Handler) CreateCorporation(ctx context.Context, req *proto.Corporation, rsp *proto.CorporationResponse) error {
	if corporation, err := h.Service.AddCorporation(ctx, req); err != nil{
		return err
	}else{
		rsp.Corporation = corporation;
		return nil
	}
}


func (h Handler) GetCorporationById(ctx context.Context, req *proto2.IdRequest, rsp *proto.CorporationResponse) error {
	if corporation, err := h.Service.GetCorporationById(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Corporation = corporation;
		return nil
	}
}


func (h Handler) GetAllCorporationsByFleetCompanyId(ctx context.Context, req *proto2.IdRequest, rsp *proto.CorporationsResponse) error {
	if corporations, err := h.Service.GetAllCorporationsByFleetCompanyId(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Corporations = corporations;
		return nil
	}
}


func (h Handler) CreateRegion(ctx context.Context, req *proto.Region, rsp *proto.RegionResponse) error {
	if region, err := h.Service.AddRegion(ctx, req); err != nil{
		return err
	}else{
		rsp.Region = region;
		return nil
	}
}


func (h Handler) GetRegionById(ctx context.Context, req *proto2.IdRequest, rsp *proto.RegionResponse) error {
	if region, err := h.Service.GetRegionById(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Region = region;
		return nil
	}
}


func (h Handler) GetAllRegionsByCorporationId(ctx context.Context, req *proto2.IdRequest, rsp *proto.RegionsResponse) error {
	if regions, err := h.Service.GetAllRegionsByCorporationId(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Regions = regions;
		return nil
	}
}


func (h Handler) CreateDistrict(ctx context.Context, req *proto.District, rsp *proto.DistrictResponse) error {
	if district, err := h.Service.AddDistrict(ctx, req); err != nil{
		return err
	}else{
		rsp.District = district;
		return nil
	}
}


func (h Handler) GetDistrictById(ctx context.Context, req *proto2.IdRequest, rsp *proto.DistrictResponse) error {
	if district, err := h.Service.GetDistrictById(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.District = district;
		return nil
	}
}


func (h Handler) GetAllDistrictsByRegionId(ctx context.Context, req *proto2.IdRequest, rsp *proto.DistrictsResponse) error {
	if districts, err := h.Service.GetAllDistrictsByRegionId(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Districts = districts;
		return nil
	}
}


func (h Handler) CreateLocation(ctx context.Context, req *proto.Location, rsp *proto.LocationResponse) error {
	if location, err := h.Service.AddLocation(ctx, req); err != nil{
		return err
	}else{
		rsp.Location = location;
		return nil
	}
}


func (h Handler) GetLocationById(ctx context.Context, req *proto2.IdRequest, rsp *proto.LocationResponse) error {
	if location, err := h.Service.GetLocationById(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Location = location;
		return nil
	}
}


func (h Handler) GetAllLocationsByDistrictId(ctx context.Context, req *proto2.IdRequest, rsp *proto.LocationsResponse) error {
	if locations, err := h.Service.GetAllLocationsByDistrictId(ctx, req.Id); err != nil{
		return err
	}else{
		rsp.Locations = locations;
		return nil
	}
}