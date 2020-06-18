package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddShippingCarrier(t *testing.T) {
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
	str.LocalDomain = "localhost11:8080"
	str.Logo = "some logo"
	str.OauthClientID = 16
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart11.com"
	str.State = "GA"
	str.StoreName = "testers11 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.Type = "GROUND"
	sc.StoreID = sid

	dbi.Close()
	scsuc, scid := si.AddShippingCarrier(&sc)
	if !scsuc || scid == 0 {
		t.Fail()
	}

	sc.ID = scid
	sc.Carrier = "FEXex"
	sc.Type = "AIR"

	dbi.Close()
	uscsuc := si.UpdateShippingCarrier(&sc)
	if !uscsuc {
		t.Fail()
	}

	dbi.Close()
	fsc := si.GetShippingCarrier(scid)
	fmt.Println("fsc: ", fsc)
	if fsc.Carrier != sc.Carrier {
		t.Fail()
	}

	dbi.Close()
	fsclist := si.GetShippingCarrierList(sid)
	fmt.Println("fsclist: ", fsclist)
	if len(*fsclist) != 1 {
		t.Fail()
	}

	dbi.Close()
	dlsc := si.DeleteShippingCarrier(scid)
	if !dlsc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
