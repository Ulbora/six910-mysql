package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddRegion(t *testing.T) {
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
	str.LocalDomain = "localhost9:8080"
	str.Logo = "some logo"
	str.OauthClientID = 14
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart9.com"
	str.State = "GA"
	str.StoreName = "testers8 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
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

	rgn.ID = rgnid
	rgn.Name = "United States of America"
	rgn.RegionCode = "USA"

	urgnsuc := si.UpdateRegion(&rgn)
	if !urgnsuc {
		t.Fail()
	}

	frgn := si.GetRegion(rgnid)
	fmt.Println("frgn", frgn)
	if frgn.RegionCode != rgn.RegionCode {
		t.Fail()
	}

	frgnList := si.GetRegionList(sid)
	fmt.Println("frgnList", frgnList)
	if len(*frgnList) != 1 {
		t.Fail()
	}

	dlsuc := si.DeleteRegion(rgnid)
	if !dlsuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
