// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package proto

import (
	_ "fleet_backend_final/common/proto"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type SignUpRequest struct {
	FleetCompanyName     string   `protobuf:"bytes,1,opt,name=fleetCompanyName,proto3" json:"fleetCompanyName,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetFleetCompanyName() string {
	if m != nil {
		return m.FleetCompanyName
	}
	return ""
}

func (m *SignUpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type FleetCompany struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FleetCompany) Reset()         { *m = FleetCompany{} }
func (m *FleetCompany) String() string { return proto.CompactTextString(m) }
func (*FleetCompany) ProtoMessage()    {}
func (*FleetCompany) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *FleetCompany) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FleetCompany.Unmarshal(m, b)
}
func (m *FleetCompany) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FleetCompany.Marshal(b, m, deterministic)
}
func (m *FleetCompany) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FleetCompany.Merge(m, src)
}
func (m *FleetCompany) XXX_Size() int {
	return xxx_messageInfo_FleetCompany.Size(m)
}
func (m *FleetCompany) XXX_DiscardUnknown() {
	xxx_messageInfo_FleetCompany.DiscardUnknown(m)
}

var xxx_messageInfo_FleetCompany proto.InternalMessageInfo

func (m *FleetCompany) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FleetCompany) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Driver struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	FleetCompanyId       string   `protobuf:"bytes,5,opt,name=fleetCompanyId,proto3" json:"fleetCompanyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Driver) Reset()         { *m = Driver{} }
func (m *Driver) String() string { return proto.CompactTextString(m) }
func (*Driver) ProtoMessage()    {}
func (*Driver) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *Driver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Driver.Unmarshal(m, b)
}
func (m *Driver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Driver.Marshal(b, m, deterministic)
}
func (m *Driver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Driver.Merge(m, src)
}
func (m *Driver) XXX_Size() int {
	return xxx_messageInfo_Driver.Size(m)
}
func (m *Driver) XXX_DiscardUnknown() {
	xxx_messageInfo_Driver.DiscardUnknown(m)
}

var xxx_messageInfo_Driver proto.InternalMessageInfo

func (m *Driver) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Driver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Driver) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Driver) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Driver) GetFleetCompanyId() string {
	if m != nil {
		return m.FleetCompanyId
	}
	return ""
}

type DriverResponse struct {
	Driver               *Driver  `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverResponse) Reset()         { *m = DriverResponse{} }
func (m *DriverResponse) String() string { return proto.CompactTextString(m) }
func (*DriverResponse) ProtoMessage()    {}
func (*DriverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *DriverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverResponse.Unmarshal(m, b)
}
func (m *DriverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverResponse.Marshal(b, m, deterministic)
}
func (m *DriverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverResponse.Merge(m, src)
}
func (m *DriverResponse) XXX_Size() int {
	return xxx_messageInfo_DriverResponse.Size(m)
}
func (m *DriverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DriverResponse proto.InternalMessageInfo

func (m *DriverResponse) GetDriver() *Driver {
	if m != nil {
		return m.Driver
	}
	return nil
}

type DriversResponse struct {
	Drivers              []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DriversResponse) Reset()         { *m = DriversResponse{} }
func (m *DriversResponse) String() string { return proto.CompactTextString(m) }
func (*DriversResponse) ProtoMessage()    {}
func (*DriversResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{4}
}

func (m *DriversResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriversResponse.Unmarshal(m, b)
}
func (m *DriversResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriversResponse.Marshal(b, m, deterministic)
}
func (m *DriversResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriversResponse.Merge(m, src)
}
func (m *DriversResponse) XXX_Size() int {
	return xxx_messageInfo_DriversResponse.Size(m)
}
func (m *DriversResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DriversResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DriversResponse proto.InternalMessageInfo

func (m *DriversResponse) GetDrivers() []*Driver {
	if m != nil {
		return m.Drivers
	}
	return nil
}

type Corporation struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FleetCompanyId       string   `protobuf:"bytes,3,opt,name=fleetCompanyId,proto3" json:"fleetCompanyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Corporation) Reset()         { *m = Corporation{} }
func (m *Corporation) String() string { return proto.CompactTextString(m) }
func (*Corporation) ProtoMessage()    {}
func (*Corporation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{5}
}

func (m *Corporation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Corporation.Unmarshal(m, b)
}
func (m *Corporation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Corporation.Marshal(b, m, deterministic)
}
func (m *Corporation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Corporation.Merge(m, src)
}
func (m *Corporation) XXX_Size() int {
	return xxx_messageInfo_Corporation.Size(m)
}
func (m *Corporation) XXX_DiscardUnknown() {
	xxx_messageInfo_Corporation.DiscardUnknown(m)
}

var xxx_messageInfo_Corporation proto.InternalMessageInfo

func (m *Corporation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Corporation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Corporation) GetFleetCompanyId() string {
	if m != nil {
		return m.FleetCompanyId
	}
	return ""
}

type CorporationResponse struct {
	Corporation          *Corporation `protobuf:"bytes,1,opt,name=corporation,proto3" json:"corporation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CorporationResponse) Reset()         { *m = CorporationResponse{} }
func (m *CorporationResponse) String() string { return proto.CompactTextString(m) }
func (*CorporationResponse) ProtoMessage()    {}
func (*CorporationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{6}
}

func (m *CorporationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorporationResponse.Unmarshal(m, b)
}
func (m *CorporationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorporationResponse.Marshal(b, m, deterministic)
}
func (m *CorporationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorporationResponse.Merge(m, src)
}
func (m *CorporationResponse) XXX_Size() int {
	return xxx_messageInfo_CorporationResponse.Size(m)
}
func (m *CorporationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CorporationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CorporationResponse proto.InternalMessageInfo

func (m *CorporationResponse) GetCorporation() *Corporation {
	if m != nil {
		return m.Corporation
	}
	return nil
}

type CorporationsResponse struct {
	Corporations         []*Corporation `protobuf:"bytes,1,rep,name=corporations,proto3" json:"corporations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CorporationsResponse) Reset()         { *m = CorporationsResponse{} }
func (m *CorporationsResponse) String() string { return proto.CompactTextString(m) }
func (*CorporationsResponse) ProtoMessage()    {}
func (*CorporationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{7}
}

func (m *CorporationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorporationsResponse.Unmarshal(m, b)
}
func (m *CorporationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorporationsResponse.Marshal(b, m, deterministic)
}
func (m *CorporationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorporationsResponse.Merge(m, src)
}
func (m *CorporationsResponse) XXX_Size() int {
	return xxx_messageInfo_CorporationsResponse.Size(m)
}
func (m *CorporationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CorporationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CorporationsResponse proto.InternalMessageInfo

func (m *CorporationsResponse) GetCorporations() []*Corporation {
	if m != nil {
		return m.Corporations
	}
	return nil
}

type Region struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CorporationId        string   `protobuf:"bytes,3,opt,name=corporationId,proto3" json:"corporationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Region) Reset()         { *m = Region{} }
func (m *Region) String() string { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()    {}
func (*Region) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{8}
}

func (m *Region) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Region.Unmarshal(m, b)
}
func (m *Region) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Region.Marshal(b, m, deterministic)
}
func (m *Region) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Region.Merge(m, src)
}
func (m *Region) XXX_Size() int {
	return xxx_messageInfo_Region.Size(m)
}
func (m *Region) XXX_DiscardUnknown() {
	xxx_messageInfo_Region.DiscardUnknown(m)
}

var xxx_messageInfo_Region proto.InternalMessageInfo

func (m *Region) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Region) GetCorporationId() string {
	if m != nil {
		return m.CorporationId
	}
	return ""
}

type RegionResponse struct {
	Region               *Region  `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionResponse) Reset()         { *m = RegionResponse{} }
func (m *RegionResponse) String() string { return proto.CompactTextString(m) }
func (*RegionResponse) ProtoMessage()    {}
func (*RegionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{9}
}

func (m *RegionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionResponse.Unmarshal(m, b)
}
func (m *RegionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionResponse.Marshal(b, m, deterministic)
}
func (m *RegionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionResponse.Merge(m, src)
}
func (m *RegionResponse) XXX_Size() int {
	return xxx_messageInfo_RegionResponse.Size(m)
}
func (m *RegionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegionResponse proto.InternalMessageInfo

func (m *RegionResponse) GetRegion() *Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type RegionsResponse struct {
	Regions              []*Region `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegionsResponse) Reset()         { *m = RegionsResponse{} }
func (m *RegionsResponse) String() string { return proto.CompactTextString(m) }
func (*RegionsResponse) ProtoMessage()    {}
func (*RegionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{10}
}

func (m *RegionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionsResponse.Unmarshal(m, b)
}
func (m *RegionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionsResponse.Marshal(b, m, deterministic)
}
func (m *RegionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionsResponse.Merge(m, src)
}
func (m *RegionsResponse) XXX_Size() int {
	return xxx_messageInfo_RegionsResponse.Size(m)
}
func (m *RegionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegionsResponse proto.InternalMessageInfo

func (m *RegionsResponse) GetRegions() []*Region {
	if m != nil {
		return m.Regions
	}
	return nil
}

type District struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RegionId             string   `protobuf:"bytes,3,opt,name=regionId,proto3" json:"regionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *District) Reset()         { *m = District{} }
func (m *District) String() string { return proto.CompactTextString(m) }
func (*District) ProtoMessage()    {}
func (*District) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{11}
}

func (m *District) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_District.Unmarshal(m, b)
}
func (m *District) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_District.Marshal(b, m, deterministic)
}
func (m *District) XXX_Merge(src proto.Message) {
	xxx_messageInfo_District.Merge(m, src)
}
func (m *District) XXX_Size() int {
	return xxx_messageInfo_District.Size(m)
}
func (m *District) XXX_DiscardUnknown() {
	xxx_messageInfo_District.DiscardUnknown(m)
}

var xxx_messageInfo_District proto.InternalMessageInfo

func (m *District) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *District) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *District) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

type DistrictResponse struct {
	District             *District `protobuf:"bytes,1,opt,name=district,proto3" json:"district,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DistrictResponse) Reset()         { *m = DistrictResponse{} }
func (m *DistrictResponse) String() string { return proto.CompactTextString(m) }
func (*DistrictResponse) ProtoMessage()    {}
func (*DistrictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{12}
}

func (m *DistrictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistrictResponse.Unmarshal(m, b)
}
func (m *DistrictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistrictResponse.Marshal(b, m, deterministic)
}
func (m *DistrictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrictResponse.Merge(m, src)
}
func (m *DistrictResponse) XXX_Size() int {
	return xxx_messageInfo_DistrictResponse.Size(m)
}
func (m *DistrictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DistrictResponse proto.InternalMessageInfo

func (m *DistrictResponse) GetDistrict() *District {
	if m != nil {
		return m.District
	}
	return nil
}

type DistrictsResponse struct {
	Districts            []*District `protobuf:"bytes,1,rep,name=districts,proto3" json:"districts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DistrictsResponse) Reset()         { *m = DistrictsResponse{} }
func (m *DistrictsResponse) String() string { return proto.CompactTextString(m) }
func (*DistrictsResponse) ProtoMessage()    {}
func (*DistrictsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{13}
}

func (m *DistrictsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistrictsResponse.Unmarshal(m, b)
}
func (m *DistrictsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistrictsResponse.Marshal(b, m, deterministic)
}
func (m *DistrictsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrictsResponse.Merge(m, src)
}
func (m *DistrictsResponse) XXX_Size() int {
	return xxx_messageInfo_DistrictsResponse.Size(m)
}
func (m *DistrictsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrictsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DistrictsResponse proto.InternalMessageInfo

func (m *DistrictsResponse) GetDistricts() []*District {
	if m != nil {
		return m.Districts
	}
	return nil
}

type Location struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DistrictId           string   `protobuf:"bytes,3,opt,name=districtId,proto3" json:"districtId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{14}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetDistrictId() string {
	if m != nil {
		return m.DistrictId
	}
	return ""
}

type LocationResponse struct {
	Location             *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LocationResponse) Reset()         { *m = LocationResponse{} }
func (m *LocationResponse) String() string { return proto.CompactTextString(m) }
func (*LocationResponse) ProtoMessage()    {}
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{15}
}

func (m *LocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationResponse.Unmarshal(m, b)
}
func (m *LocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationResponse.Marshal(b, m, deterministic)
}
func (m *LocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationResponse.Merge(m, src)
}
func (m *LocationResponse) XXX_Size() int {
	return xxx_messageInfo_LocationResponse.Size(m)
}
func (m *LocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LocationResponse proto.InternalMessageInfo

func (m *LocationResponse) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type LocationsResponse struct {
	Locations            []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LocationsResponse) Reset()         { *m = LocationsResponse{} }
func (m *LocationsResponse) String() string { return proto.CompactTextString(m) }
func (*LocationsResponse) ProtoMessage()    {}
func (*LocationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{16}
}

func (m *LocationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationsResponse.Unmarshal(m, b)
}
func (m *LocationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationsResponse.Marshal(b, m, deterministic)
}
func (m *LocationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationsResponse.Merge(m, src)
}
func (m *LocationsResponse) XXX_Size() int {
	return xxx_messageInfo_LocationsResponse.Size(m)
}
func (m *LocationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LocationsResponse proto.InternalMessageInfo

func (m *LocationsResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "proto.SignUpRequest")
	proto.RegisterType((*FleetCompany)(nil), "proto.FleetCompany")
	proto.RegisterType((*Driver)(nil), "proto.Driver")
	proto.RegisterType((*DriverResponse)(nil), "proto.DriverResponse")
	proto.RegisterType((*DriversResponse)(nil), "proto.DriversResponse")
	proto.RegisterType((*Corporation)(nil), "proto.Corporation")
	proto.RegisterType((*CorporationResponse)(nil), "proto.CorporationResponse")
	proto.RegisterType((*CorporationsResponse)(nil), "proto.CorporationsResponse")
	proto.RegisterType((*Region)(nil), "proto.Region")
	proto.RegisterType((*RegionResponse)(nil), "proto.RegionResponse")
	proto.RegisterType((*RegionsResponse)(nil), "proto.RegionsResponse")
	proto.RegisterType((*District)(nil), "proto.District")
	proto.RegisterType((*DistrictResponse)(nil), "proto.DistrictResponse")
	proto.RegisterType((*DistrictsResponse)(nil), "proto.DistrictsResponse")
	proto.RegisterType((*Location)(nil), "proto.Location")
	proto.RegisterType((*LocationResponse)(nil), "proto.LocationResponse")
	proto.RegisterType((*LocationsResponse)(nil), "proto.LocationsResponse")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcb, 0x4f, 0xdb, 0x4e,
	0x10, 0x26, 0x3c, 0xf2, 0x0b, 0x03, 0x49, 0x60, 0x7f, 0xb4, 0x4d, 0x4d, 0x55, 0xa1, 0x55, 0x5f,
	0xa2, 0x2d, 0x48, 0xb4, 0x2a, 0x12, 0xaa, 0x5a, 0xd5, 0x49, 0xa1, 0x81, 0x8a, 0x43, 0x10, 0x87,
	0xf6, 0x82, 0x8c, 0xbd, 0x20, 0xab, 0x89, 0xed, 0x7a, 0x0d, 0x55, 0x8e, 0xbd, 0xf4, 0x6f, 0xee,
	0xb1, 0xf2, 0x3e, 0xc6, 0xeb, 0x07, 0x21, 0x9c, 0x1c, 0x7f, 0x33, 0xf3, 0x79, 0xbe, 0x9d, 0x9d,
	0x2f, 0xd0, 0x72, 0xaf, 0x78, 0x12, 0x8e, 0x58, 0xbc, 0x15, 0xc5, 0x61, 0x12, 0x92, 0x05, 0xf1,
	0xb0, 0x96, 0xdd, 0x70, 0x34, 0x0a, 0x03, 0x09, 0xd2, 0xdf, 0x35, 0x68, 0x9e, 0xf8, 0x97, 0xc1,
	0x69, 0x34, 0x60, 0x3f, 0xaf, 0x18, 0x4f, 0xc8, 0x26, 0xac, 0x5c, 0x0c, 0x19, 0x4b, 0xba, 0xe1,
	0x28, 0x72, 0x82, 0xf1, 0xb1, 0x33, 0x62, 0x9d, 0xda, 0x46, 0xed, 0xc5, 0xe2, 0xa0, 0x84, 0x13,
	0x02, 0xf3, 0x41, 0x1a, 0x9f, 0x15, 0x71, 0xf1, 0x9b, 0xac, 0xc1, 0x02, 0x1b, 0x39, 0xfe, 0xb0,
	0x33, 0x27, 0x40, 0xf9, 0x42, 0x2c, 0x68, 0x44, 0x0e, 0xe7, 0xbf, 0xc2, 0xd8, 0xeb, 0xcc, 0x8b,
	0x00, 0xbe, 0xd3, 0x1d, 0x58, 0xde, 0x37, 0x98, 0x49, 0x0b, 0x66, 0x7d, 0x4f, 0x7d, 0x73, 0xd6,
	0xf7, 0xaa, 0xbe, 0x42, 0xff, 0xd4, 0xa0, 0xde, 0x8b, 0xfd, 0x6b, 0x16, 0x4f, 0x93, 0x7e, 0xf7,
	0xa6, 0xc8, 0x33, 0x68, 0x99, 0x72, 0xfb, 0x5e, 0x67, 0x41, 0x64, 0x14, 0x50, 0xba, 0x0b, 0x2d,
	0xd9, 0xc7, 0x80, 0xf1, 0x28, 0x0c, 0x38, 0x23, 0x4f, 0xa1, 0xee, 0x09, 0x44, 0xf4, 0xb4, 0xb4,
	0xd3, 0x94, 0x47, 0xbd, 0xa5, 0xd2, 0x54, 0x90, 0xee, 0x41, 0x5b, 0x22, 0x1c, 0x2b, 0x9f, 0xc3,
	0x7f, 0x32, 0xc8, 0x3b, 0xb5, 0x8d, 0xb9, 0x72, 0xa9, 0x8e, 0xd2, 0x6f, 0xb0, 0xd4, 0x0d, 0xe3,
	0x28, 0x8c, 0x9d, 0xc4, 0x0f, 0x83, 0xa9, 0x4e, 0xa0, 0xac, 0x67, 0xae, 0x52, 0xcf, 0x11, 0xfc,
	0x6f, 0x50, 0x63, 0x6b, 0x6f, 0x61, 0xc9, 0xcd, 0x60, 0xa5, 0x8c, 0xa8, 0xf6, 0xcc, 0x02, 0x33,
	0x8d, 0x1e, 0xc3, 0x9a, 0x11, 0xcb, 0x84, 0xbe, 0x83, 0x65, 0x23, 0x4d, 0xab, 0xad, 0xa2, 0xcb,
	0xe5, 0xd1, 0x01, 0xd4, 0x07, 0xec, 0x72, 0x5a, 0xc9, 0x4f, 0xa0, 0x69, 0x54, 0xa3, 0xe2, 0x3c,
	0x98, 0x0e, 0x50, 0x72, 0x9a, 0x03, 0x8c, 0x05, 0x52, 0x18, 0xa0, 0x4a, 0x53, 0xc1, 0x74, 0x80,
	0x12, 0xc9, 0x0d, 0x50, 0x06, 0x8b, 0x03, 0x54, 0xa5, 0x3a, 0x4a, 0x0f, 0xa1, 0xd1, 0xf3, 0x79,
	0x12, 0xfb, 0x6e, 0x32, 0x95, 0x14, 0x0b, 0x1a, 0xb2, 0x14, 0x55, 0xe0, 0x3b, 0xfd, 0x08, 0x2b,
	0x9a, 0x0b, 0x1b, 0x79, 0x09, 0x0d, 0x4f, 0x61, 0x4a, 0x44, 0x5b, 0x5f, 0x25, 0x9d, 0x8a, 0x09,
	0xd4, 0x86, 0x55, 0x8d, 0x66, 0x52, 0x5e, 0xc3, 0xa2, 0x4e, 0xd0, 0x62, 0x4a, 0x14, 0x59, 0x06,
	0x3d, 0x86, 0xc6, 0xd7, 0xd0, 0x9d, 0xfe, 0x3a, 0x3e, 0x06, 0xd0, 0xc5, 0x28, 0xc9, 0x40, 0x52,
	0x51, 0x9a, 0xcf, 0x14, 0x35, 0x54, 0x58, 0x41, 0x14, 0xa6, 0x62, 0x42, 0x2a, 0x4a, 0xa3, 0x39,
	0x51, 0x3a, 0xa1, 0x28, 0x0a, 0x29, 0xb2, 0x8c, 0x9d, 0xbf, 0x0d, 0x68, 0x6b, 0x13, 0x3d, 0x61,
	0xf1, 0xb5, 0xef, 0x32, 0xb2, 0x0b, 0x75, 0xe9, 0x97, 0x64, 0x4d, 0x55, 0xe6, 0xec, 0xd3, 0xba,
	0x97, 0x5f, 0x59, 0xf5, 0x65, 0x3a, 0x93, 0xde, 0xf9, 0x6e, 0xcc, 0x9c, 0x84, 0x29, 0xdb, 0xca,
	0xef, 0xf6, 0xc4, 0xba, 0xd3, 0xc8, 0xbb, 0x7b, 0xdd, 0x1e, 0x34, 0x0f, 0x58, 0x22, 0x61, 0x7b,
	0xdc, 0xf7, 0xc8, 0x8a, 0xca, 0xec, 0x7b, 0xb7, 0xf6, 0xfa, 0x05, 0x2c, 0xac, 0xe5, 0xf6, 0x78,
	0x3f, 0x67, 0x11, 0x15, 0x44, 0xf7, 0x73, 0x44, 0xdc, 0x60, 0xfa, 0x0c, 0xab, 0x52, 0xb5, 0xe9,
	0x57, 0x15, 0x8b, 0x6e, 0x59, 0x15, 0xcb, 0x9f, 0xd1, 0x9c, 0x02, 0x3d, 0x60, 0xc9, 0xa7, 0xe1,
	0xd0, 0xb4, 0x93, 0x29, 0x1a, 0x5b, 0x2f, 0xb3, 0x9a, 0xdd, 0xf5, 0x80, 0x1c, 0xa4, 0x04, 0x18,
	0xbc, 0xe1, 0xa0, 0x26, 0x37, 0x87, 0x93, 0x55, 0xde, 0x94, 0x5f, 0x7a, 0x3c, 0xe5, 0xbc, 0xcb,
	0xd0, 0x19, 0x72, 0x08, 0x8f, 0xa4, 0x28, 0x65, 0x23, 0xf6, 0xb8, 0x6b, 0x3a, 0xd3, 0x84, 0x73,
	0x2e, 0xf8, 0x0e, 0x4e, 0x5b, 0xe2, 0xb7, 0x4c, 0xbb, 0xd4, 0xc7, 0x7b, 0x68, 0xa9, 0x9b, 0xa9,
	0x2d, 0xa9, 0xb8, 0xe9, 0xd6, 0x83, 0xe2, 0xea, 0x67, 0xd5, 0x7d, 0x78, 0x28, 0x55, 0xa0, 0x87,
	0xd8, 0xe3, 0x81, 0xf2, 0xa6, 0x8a, 0x2e, 0x3a, 0x05, 0x26, 0x53, 0xc4, 0x07, 0x68, 0xa7, 0xd7,
	0x4e, 0x45, 0x6e, 0x90, 0x31, 0xa1, 0x15, 0x14, 0x82, 0x56, 0x54, 0xdc, 0x6e, 0xac, 0x2e, 0x9a,
	0x0b, 0x9d, 0x21, 0x47, 0xb0, 0x2e, 0x85, 0xa0, 0x6f, 0xd8, 0xe3, 0x1e, 0x3a, 0xd2, 0x04, 0x29,
	0x25, 0x9f, 0x41, 0x29, 0x3a, 0x72, 0x8b, 0x94, 0x72, 0x33, 0xf6, 0xab, 0xef, 0x9b, 0xe2, 0x8f,
	0xf9, 0xec, 0xdc, 0x71, 0x7f, 0xb0, 0xc0, 0x3b, 0xbb, 0xf0, 0x03, 0x67, 0xb8, 0xad, 0xdd, 0xe8,
	0x8c, 0x4b, 0x3b, 0xda, 0x16, 0x04, 0xe7, 0x75, 0xf1, 0x78, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x96, 0xfa, 0x9c, 0xf3, 0x09, 0x00, 0x00,
}
