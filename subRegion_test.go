package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_SubRegion(t *testing.T) {
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
	str.LocalDomain = "localhost10:8080"
	str.Logo = "some logo"
	str.OauthClientID = 15
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart10.com"
	str.State = "GA"
	str.StoreName = "testers10 fantastic store"
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

	var srgn sdbi.SubRegion
	srgn.Name = "Georgia"
	srgn.SubRegionCode = "G"
	srgn.RegionID = rgnid

	dbi.Close()
	srgnsuc, srgnid := si.AddSubRegion(&srgn)
	if !srgnsuc || srgnid == 0 {
		t.Fail()
	}

	srgn.ID = srgnid
	srgn.Name = "State of Georgia"
	srgn.SubRegionCode = "GA"

	dbi.Close()
	usrgnsuc := si.UpdateSubRegion(&srgn)
	if !usrgnsuc {
		t.Fail()
	}

	dbi.Close()
	fsrgn := si.GetSubRegion(srgnid)
	fmt.Println("fsrgn", fsrgn)
	if fsrgn.SubRegionCode != srgn.SubRegionCode {
		t.Fail()
	}

	dbi.Close()
	fsrgnlist := si.GetSubRegionList(rgnid)
	fmt.Println("fsrgnlist", fsrgnlist)
	if len(*fsrgnlist) != 1 {
		t.Fail()
	}

	dbi.Close()
	dlsuc := si.DeleteSubRegion(srgnid)
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
