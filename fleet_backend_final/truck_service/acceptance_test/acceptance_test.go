package acceptance_test

import (
	"context"
	"fleet_backend_final/common"
	proto2 "fleet_backend_final/common/proto"
	"fleet_backend_final/truck_service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"testing"
)

func TestSomething(t *testing.T) {
	service := proto.NewTruckService("truckservice", client.NewClient(common.UseConsul))

	response, err := service.CreateTruck(context.Background(), &proto.Truck{
		Id:             uuid.New().String(),
		LicensePlate:   "RJ14JS5522",
		ClockedInUser:  "ashish",
		Miles:          25.00,
		FleetCompanyId: uuid.New().String(),
		CorporationId:  uuid.New().String(),
		RegionId:       uuid.New().String(),
		DistrictId:     uuid.New().String(),
		LocationId:     uuid.New().String(),
	})

	//Test for CreateTruck
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, response.Truck.LicensePlate, is.EqualTo("RJ14JS5522"))
	then.AssertThat(t, response.Truck.FleetCompanyId == "", is.False())
	then.AssertThat(t, response.Truck.CorporationId == "", is.False())
	then.AssertThat(t, response.Truck.RegionId == "", is.False())
	then.AssertThat(t, response.Truck.DistrictId == "", is.False())
	then.AssertThat(t, response.Truck.LocationId == "", is.False())

	truckId := response.Truck.Id
	truckResponse, err := service.GetTruckById(context.Background(), &proto2.IdRequest{Id: truckId})

	//Test for GetTruckById
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, truckResponse.Truck.Id, is.EqualTo(truckId))
	then.AssertThat(t, truckResponse.Truck.Id == "", is.False())

	fleetcompanyid := response.Truck.FleetCompanyId
	trucksByFleetCompany, err := service.GetAllTrucksByFleetCompanyId(context.Background(), &proto2.IdRequest{Id: fleetcompanyid})

	//Test for GetAllTruckByFleetCompanyId
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, trucksByFleetCompany.Trucks, has.Length(1))

	updatedTruck, err := service.UpdateTruck(context.Background(), &proto.Truck{
		Id:             truckId,
		LicensePlate:   response.Truck.LicensePlate,
		ClockedInUser:  response.Truck.ClockedInUser,
		CorporationId:  uuid.New().String(),
		Miles:          50.00,
		FleetCompanyId: response.Truck.FleetCompanyId,
		RegionId:       uuid.New().String(),
		DistrictId:     uuid.New().String(),
		LocationId:     uuid.New().String(),
	})

	//Test for UpdateTruck
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updatedTruck.Truck.Miles == 50.00, is.True())
	then.AssertThat(t, updatedTruck.Truck.FleetCompanyId , is.EqualTo(fleetcompanyid))

	clockOutResponse, err := service.ClockOut(context.Background(), &proto.ClockOperation{
		DriverId: uuid.New().String(),
		TruckId: truckId,
	})

	//Test for ClockOut
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, clockOutResponse.Truck.ClockedInUser == "", is.True())
	then.AssertThat(t, clockOutResponse.Truck.FleetCompanyId , is.EqualTo(fleetcompanyid))

	clockInResponse, err := service.ClockIn(context.Background(), &proto.ClockOperation{
		DriverId: uuid.New().String(),
		TruckId: truckId,
	})

	//Test for Clock In
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, clockInResponse.Truck.ClockedInUser == "", is.False())
	then.AssertThat(t, clockInResponse.Truck.FleetCompanyId , is.EqualTo(fleetcompanyid))
}