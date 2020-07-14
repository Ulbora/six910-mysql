package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddExcludedSubRegion(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb Six910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	si := sdb.GetNew()

	var str sdbi.Store
	str.City = "test vill"
	str.Company = "test com"
	str.Currency = "USD"
	str.Email = "tester@tester.com"
	str.FirstName = "Tester"
	str.LastName = "Bill"
	str.LocalDomain = "localhost16:8080"
	str.Logo = "some logo"
	str.OauthClientID = 20
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart16.com"
	str.State = "GA"
	str.StoreName = "testers16 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var ins sdbi.Insurance
	ins.Cost = 4.95
	ins.MinOrderAmount = 50.00
	ins.MaxOrderAmount = 500.00
	ins.StoreID = sid

	inssuc, insid := si.AddInsurance(&ins)
	if !inssuc || insid == 0 {
		t.Fail()
	}

	var rgn sdbi.Region
	rgn.Name = "USA"
	rgn.RegionCode = "US"
	rgn.StoreID = sid

	rgnsuc, rgnid := si.AddRegion(&rgn)
	if !rgnsuc || rgnid == 0 {
		t.Fail()
	}

	var srgn sdbi.SubRegion
	srgn.Name = "Georgia"
	srgn.SubRegionCode = "G"
	srgn.RegionID = rgnid

	srgnsuc, srgnid := si.AddSubRegion(&srgn)
	if !srgnsuc || srgnid == 0 {
		t.Fail()
	}

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.Type = "GROUND"
	sc.StoreID = sid

	scsuc, scid := si.AddShippingCarrier(&sc)
	if !scsuc || scid == 0 {
		t.Fail()
	}

	var sm sdbi.ShippingMethod
	sm.Cost = 12.95
	sm.Handling = 2.50
	sm.InsuranceID = insid
	sm.MaxOrderAmount = 1000.00
	sm.MaxWeight = 100
	sm.MinOrderAmount = 1.00
	sm.Name = "USP Ground"
	sm.RegionID = rgnid
	sm.ShippingCarrierID = scid
	sm.StoreID = sid

	smsuc, smid := si.AddShippingMethod(&sm)
	if !smsuc || smid == 0 {
		t.Fail()
	}

	var esr sdbi.ExcludedSubRegion
	esr.RegionID = rgnid
	esr.ShippingMethodID = smid
	esr.SubRegionID = srgnid

	dbi.Close()
	esrsuc, esrid := si.AddExcludedSubRegion(&esr)
	if !esrsuc || esrid == 0 {
		t.Fail()
	}

	// this is not implemented
	uesrsuc := si.UpdateExcludedSubRegion(&esr)
	if uesrsuc {
		t.Fail()
	}

	// this is not implemented
	fesr := si.GetExcludedSubRegion(esrid)
	fmt.Println("fesr: ", fesr)
	if fesr == nil {
		t.Fail()
	}

	dbi.Close()
	fesrlist := si.GetExcludedSubRegionList(rgnid)
	fmt.Println("fesrlist: ", fesrlist)
	if len(*fesrlist) != 1 {
		t.Fail()
	}

	dbi.Close()
	dlesrsuc := si.DeleteExcludedSubRegion(esrid)
	if !dlesrsuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
