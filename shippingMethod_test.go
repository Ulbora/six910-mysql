package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_ShippingMethod(t *testing.T) {
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
	str.LocalDomain = "localhost14:8080"
	str.Logo = "some logo"
	str.OauthClientID = 18
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart14.com"
	str.State = "GA"
	str.StoreName = "testers14 fantastic store"
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

	sm.ID = smid
	sm.Cost = 15.95
	sm.Handling = 4.50
	sm.MaxOrderAmount = 2000.00
	sm.MaxWeight = 200
	sm.MinOrderAmount = 5.00
	sm.Name = "USP Ground Express"

	usmsuc := si.UpdateShippingMethod(&sm)
	if !usmsuc {
		t.Fail()
	}

	fsm := si.GetShippingMethod(smid)
	fmt.Println("fsm: ", fsm)
	if fsm.Cost != sm.Cost || fsm.MaxWeight != sm.MaxWeight || fsm.ShippingCarrierID != sm.ShippingCarrierID ||
		fsm.Name != sm.Name {
		t.Fail()
	}

	fsmlist := si.GetShippingMethodList(sid)
	fmt.Println("fsmlist: ", fsmlist)
	if len(*fsmlist) != 1 {
		t.Fail()
	}

	dlsm := si.DeleteShippingMethod(smid)
	if !dlsm {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
