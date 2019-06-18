package acceptance_test

import (
	"context"
	"fleet_backend_final/common"
	proto2 "fleet_backend_final/common/proto"
	"fleet_backend_final/customer_service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"testing"
)

func TestSomething(t *testing.T){
	service := proto.NewCustomerService("customerservice", client.NewClient(common.UseConsul))
	response, err := service.SignUp(context.Background(), &proto.SignUpRequest{
		Name: "ashish",
		Email: "abc",
		Password: "def",
		FleetCompanyName:"fleetup",
	})

	//Test for SignUp
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, response.Driver.Email, is.EqualTo("abc"))
	then.AssertThat(t, response.Driver.FleetCompanyId == "", is.False())


	driverId := response.Driver.Id
	driverResponse, err := service.GetDriverById(context.Background(), &proto2.IdRequest{Id:driverId})

	//Test for GetDriverById
	then.AssertThat(t ,err, is.Nil())
	then.AssertThat(t, driverResponse.Driver.Id, is.EqualTo(driverId))
	then.AssertThat(t, driverResponse.Driver.Id == "", is.False())


	fleetcompanyid := response.Driver.FleetCompanyId
	driversByFleetCompany, err := service.GetDriversByFleetCompanyId(context.Background(), &proto2.IdRequest{Id:fleetcompanyid})

	//Test for GetDriversByFleetCompanyId
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, driversByFleetCompany.Drivers, has.Length(1))

	//fleetcompanyid := driverResponse.Driver.FleetCompanyId
	//driversByFleetCompany, err := service.GetDriversByFleetCompanyId(context.Background(), &proto2.IdRequest{Id:fleetcompanyid})
	//then.AssertThat(t, err, is.Nil())
	//then.AssertThat(t, driversByFleetCompany.Drivers, has.Length(1))

	updatedDriver, err := service.UpdateDriver(context.Background(),&proto.Driver{
		Id:driverId,
		Name:"NewName",
		Email:"NewEmail",
		Password:"NewPassword",
		FleetCompanyId: uuid.New().String(),
	})

	//Test for UpdateDriver
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updatedDriver.Driver.Email, is.EqualTo("NewEmail"))

	corporationResponse, err1 := service.CreateCorporation(context.Background(), &proto.Corporation{
		Name: "CorporationAlpha",
		FleetCompanyId: fleetcompanyid,})

	//Test for CreateCorporation
	then.AssertThat(t, err1, is.Nil())
	then.AssertThat(t, corporationResponse.Corporation.Name, is.EqualTo("CorporationAlpha"))

	corporationid := corporationResponse.Corporation.Id
	corporationById, err2 := service.GetCorporationById(context.Background(), &proto2.IdRequest{Id: corporationid})

	//Test for GetCorporationById
	then.AssertThat(t, err2, is.Nil())
	then.AssertThat(t, corporationById.Corporation.Id == "", is.False())

	//fleetcompanyid := corporationResponse.Corporation.FleetCompanyId
	corporationsByFleetCompanyId, err3 := service.GetAllCorporationsByFleetCompanyId(context.Background(), &proto2.IdRequest{Id:fleetcompanyid})

	//Test for GetAllCorporationsByFleetCompanyId
	then.AssertThat(t, err3, is.Nil())
	then.AssertThat(t, corporationsByFleetCompanyId.Corporations, has.Length(1))


	regionResponse, err := service.CreateRegion(context.Background(), &proto.Region{
		Name: "Rajasthan",
		CorporationId: corporationid,
	})

	//Test for CreateRegion
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, regionResponse.Region.Name, is.EqualTo("Rajasthan"))

	RegionId := regionResponse.Region.Id
	regionById, err := service.GetRegionById(context.Background(), &proto2.IdRequest{Id: RegionId})

	//Test for GetRegionById
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, regionById.Region.Id == RegionId, is.True())

	//corporationid := regionResponse.Region.CorporationId
	regionsByCorporationId, err := service.GetAllRegionsByCorporationId(context.Background(), &proto2.IdRequest{Id:corporationid})

	//Test for GetAllRegionsByCorporationId
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, regionsByCorporationId.Regions, has.Length(1))

	districtResponse, err := service.CreateDistrict(context.Background(), &proto.District{
		Name: "Jaipur",
		RegionId: RegionId,
	})

	//Test for CreateDistrict
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, districtResponse.District.Name, is.EqualTo("Jaipur"))

	DistrictId := districtResponse.District.Id
	districtById, err := service.GetDistrictById(context.Background(), &proto2.IdRequest{Id: DistrictId})

	//Test for GetDistrictById
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, districtById.District.Id == "", is.False())

	//regionId := districtResponse.District.RegionId
	districtsByRegionId, err := service.GetAllDistrictsByRegionId(context.Background(), &proto2.IdRequest{Id:RegionId})

	//Test for GetAllDistrictsByRegionId
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, districtsByRegionId.Districts, has.Length(1))

	locationResponse, err := service.CreateLocation(context.Background(), &proto.Location{
		Name: "Mansarover",
		DistrictId: DistrictId,
	})

	//Test for CreateLocation
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, locationResponse.Location.Name, is.EqualTo("Mansarover"))

	LocationId := locationResponse.Location.Id
	locationById, err := service.GetLocationById(context.Background(), &proto2.IdRequest{Id: LocationId})

	//Test for GetLocationById
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, locationById.Location.Id == LocationId, is.True())

	//districtId := locationResponse.Location.DistrictId
	locationsByDistrictId, err := service.GetAllLocationsByDistrictId(context.Background(), &proto2.IdRequest{Id:DistrictId})

	//Test for GetAllLocationsByDistrictId
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, locationsByDistrictId.Locations, has.Length(1))

}
