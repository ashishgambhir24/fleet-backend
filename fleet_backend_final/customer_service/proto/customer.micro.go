// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: customer.proto

package proto

import (
	proto1 "fleet_backend_final/common/proto"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CustomerService service

type CustomerService interface {
	// Drivers
	SignUp(ctx context.Context, in *SignUpRequest, opts ...client.CallOption) (*DriverResponse, error)
	CreateDriver(ctx context.Context, in *Driver, opts ...client.CallOption) (*DriverResponse, error)
	UpdateDriver(ctx context.Context, in *Driver, opts ...client.CallOption) (*DriverResponse, error)
	GetDriverById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DriverResponse, error)
	GetDriversByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DriversResponse, error)
	// Corporation
	CreateCorporation(ctx context.Context, in *Corporation, opts ...client.CallOption) (*CorporationResponse, error)
	GetAllCorporationsByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*CorporationsResponse, error)
	GetCorporationById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*CorporationResponse, error)
	// Regions
	CreateRegion(ctx context.Context, in *Region, opts ...client.CallOption) (*RegionResponse, error)
	GetAllRegionsByCorporationId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*RegionsResponse, error)
	GetRegionById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*RegionResponse, error)
	// District
	CreateDistrict(ctx context.Context, in *District, opts ...client.CallOption) (*DistrictResponse, error)
	GetAllDistrictsByRegionId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DistrictsResponse, error)
	GetDistrictById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DistrictResponse, error)
	// Location
	CreateLocation(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error)
	GetAllLocationsByDistrictId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*LocationsResponse, error)
	GetLocationById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*LocationResponse, error)
}

type customerService struct {
	c    client.Client
	name string
}

func NewCustomerService(name string, c client.Client) CustomerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &customerService{
		c:    c,
		name: name,
	}
}

func (c *customerService) SignUp(ctx context.Context, in *SignUpRequest, opts ...client.CallOption) (*DriverResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.SignUp", in)
	out := new(DriverResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateDriver(ctx context.Context, in *Driver, opts ...client.CallOption) (*DriverResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateDriver", in)
	out := new(DriverResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) UpdateDriver(ctx context.Context, in *Driver, opts ...client.CallOption) (*DriverResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.UpdateDriver", in)
	out := new(DriverResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetDriverById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DriverResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetDriverById", in)
	out := new(DriverResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetDriversByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DriversResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetDriversByFleetCompanyId", in)
	out := new(DriversResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateCorporation(ctx context.Context, in *Corporation, opts ...client.CallOption) (*CorporationResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateCorporation", in)
	out := new(CorporationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetAllCorporationsByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*CorporationsResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetAllCorporationsByFleetCompanyId", in)
	out := new(CorporationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCorporationById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*CorporationResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetCorporationById", in)
	out := new(CorporationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateRegion(ctx context.Context, in *Region, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateRegion", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetAllRegionsByCorporationId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*RegionsResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetAllRegionsByCorporationId", in)
	out := new(RegionsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetRegionById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetRegionById", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateDistrict(ctx context.Context, in *District, opts ...client.CallOption) (*DistrictResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateDistrict", in)
	out := new(DistrictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetAllDistrictsByRegionId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DistrictsResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetAllDistrictsByRegionId", in)
	out := new(DistrictsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetDistrictById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*DistrictResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetDistrictById", in)
	out := new(DistrictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) CreateLocation(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.CreateLocation", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetAllLocationsByDistrictId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*LocationsResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetAllLocationsByDistrictId", in)
	out := new(LocationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetLocationById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "CustomerService.GetLocationById", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerService service

type CustomerServiceHandler interface {
	// Drivers
	SignUp(context.Context, *SignUpRequest, *DriverResponse) error
	CreateDriver(context.Context, *Driver, *DriverResponse) error
	UpdateDriver(context.Context, *Driver, *DriverResponse) error
	GetDriverById(context.Context, *proto1.IdRequest, *DriverResponse) error
	GetDriversByFleetCompanyId(context.Context, *proto1.IdRequest, *DriversResponse) error
	// Corporation
	CreateCorporation(context.Context, *Corporation, *CorporationResponse) error
	GetAllCorporationsByFleetCompanyId(context.Context, *proto1.IdRequest, *CorporationsResponse) error
	GetCorporationById(context.Context, *proto1.IdRequest, *CorporationResponse) error
	// Regions
	CreateRegion(context.Context, *Region, *RegionResponse) error
	GetAllRegionsByCorporationId(context.Context, *proto1.IdRequest, *RegionsResponse) error
	GetRegionById(context.Context, *proto1.IdRequest, *RegionResponse) error
	// District
	CreateDistrict(context.Context, *District, *DistrictResponse) error
	GetAllDistrictsByRegionId(context.Context, *proto1.IdRequest, *DistrictsResponse) error
	GetDistrictById(context.Context, *proto1.IdRequest, *DistrictResponse) error
	// Location
	CreateLocation(context.Context, *Location, *LocationResponse) error
	GetAllLocationsByDistrictId(context.Context, *proto1.IdRequest, *LocationsResponse) error
	GetLocationById(context.Context, *proto1.IdRequest, *LocationResponse) error
}

func RegisterCustomerServiceHandler(s server.Server, hdlr CustomerServiceHandler, opts ...server.HandlerOption) error {
	type customerService interface {
		SignUp(ctx context.Context, in *SignUpRequest, out *DriverResponse) error
		CreateDriver(ctx context.Context, in *Driver, out *DriverResponse) error
		UpdateDriver(ctx context.Context, in *Driver, out *DriverResponse) error
		GetDriverById(ctx context.Context, in *proto1.IdRequest, out *DriverResponse) error
		GetDriversByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, out *DriversResponse) error
		CreateCorporation(ctx context.Context, in *Corporation, out *CorporationResponse) error
		GetAllCorporationsByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, out *CorporationsResponse) error
		GetCorporationById(ctx context.Context, in *proto1.IdRequest, out *CorporationResponse) error
		CreateRegion(ctx context.Context, in *Region, out *RegionResponse) error
		GetAllRegionsByCorporationId(ctx context.Context, in *proto1.IdRequest, out *RegionsResponse) error
		GetRegionById(ctx context.Context, in *proto1.IdRequest, out *RegionResponse) error
		CreateDistrict(ctx context.Context, in *District, out *DistrictResponse) error
		GetAllDistrictsByRegionId(ctx context.Context, in *proto1.IdRequest, out *DistrictsResponse) error
		GetDistrictById(ctx context.Context, in *proto1.IdRequest, out *DistrictResponse) error
		CreateLocation(ctx context.Context, in *Location, out *LocationResponse) error
		GetAllLocationsByDistrictId(ctx context.Context, in *proto1.IdRequest, out *LocationsResponse) error
		GetLocationById(ctx context.Context, in *proto1.IdRequest, out *LocationResponse) error
	}
	type CustomerService struct {
		customerService
	}
	h := &customerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CustomerService{h}, opts...))
}

type customerServiceHandler struct {
	CustomerServiceHandler
}

func (h *customerServiceHandler) SignUp(ctx context.Context, in *SignUpRequest, out *DriverResponse) error {
	return h.CustomerServiceHandler.SignUp(ctx, in, out)
}

func (h *customerServiceHandler) CreateDriver(ctx context.Context, in *Driver, out *DriverResponse) error {
	return h.CustomerServiceHandler.CreateDriver(ctx, in, out)
}

func (h *customerServiceHandler) UpdateDriver(ctx context.Context, in *Driver, out *DriverResponse) error {
	return h.CustomerServiceHandler.UpdateDriver(ctx, in, out)
}

func (h *customerServiceHandler) GetDriverById(ctx context.Context, in *proto1.IdRequest, out *DriverResponse) error {
	return h.CustomerServiceHandler.GetDriverById(ctx, in, out)
}

func (h *customerServiceHandler) GetDriversByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, out *DriversResponse) error {
	return h.CustomerServiceHandler.GetDriversByFleetCompanyId(ctx, in, out)
}

func (h *customerServiceHandler) CreateCorporation(ctx context.Context, in *Corporation, out *CorporationResponse) error {
	return h.CustomerServiceHandler.CreateCorporation(ctx, in, out)
}

func (h *customerServiceHandler) GetAllCorporationsByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, out *CorporationsResponse) error {
	return h.CustomerServiceHandler.GetAllCorporationsByFleetCompanyId(ctx, in, out)
}

func (h *customerServiceHandler) GetCorporationById(ctx context.Context, in *proto1.IdRequest, out *CorporationResponse) error {
	return h.CustomerServiceHandler.GetCorporationById(ctx, in, out)
}

func (h *customerServiceHandler) CreateRegion(ctx context.Context, in *Region, out *RegionResponse) error {
	return h.CustomerServiceHandler.CreateRegion(ctx, in, out)
}

func (h *customerServiceHandler) GetAllRegionsByCorporationId(ctx context.Context, in *proto1.IdRequest, out *RegionsResponse) error {
	return h.CustomerServiceHandler.GetAllRegionsByCorporationId(ctx, in, out)
}

func (h *customerServiceHandler) GetRegionById(ctx context.Context, in *proto1.IdRequest, out *RegionResponse) error {
	return h.CustomerServiceHandler.GetRegionById(ctx, in, out)
}

func (h *customerServiceHandler) CreateDistrict(ctx context.Context, in *District, out *DistrictResponse) error {
	return h.CustomerServiceHandler.CreateDistrict(ctx, in, out)
}

func (h *customerServiceHandler) GetAllDistrictsByRegionId(ctx context.Context, in *proto1.IdRequest, out *DistrictsResponse) error {
	return h.CustomerServiceHandler.GetAllDistrictsByRegionId(ctx, in, out)
}

func (h *customerServiceHandler) GetDistrictById(ctx context.Context, in *proto1.IdRequest, out *DistrictResponse) error {
	return h.CustomerServiceHandler.GetDistrictById(ctx, in, out)
}

func (h *customerServiceHandler) CreateLocation(ctx context.Context, in *Location, out *LocationResponse) error {
	return h.CustomerServiceHandler.CreateLocation(ctx, in, out)
}

func (h *customerServiceHandler) GetAllLocationsByDistrictId(ctx context.Context, in *proto1.IdRequest, out *LocationsResponse) error {
	return h.CustomerServiceHandler.GetAllLocationsByDistrictId(ctx, in, out)
}

func (h *customerServiceHandler) GetLocationById(ctx context.Context, in *proto1.IdRequest, out *LocationResponse) error {
	return h.CustomerServiceHandler.GetLocationById(ctx, in, out)
}
