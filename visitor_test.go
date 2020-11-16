package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddVisit(t *testing.T) {
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
	str.LocalDomain = "localhost6:8080"
	str.Logo = "some logo"
	str.OauthClientID = 10
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart6.com"
	str.State = "GA"
	str.StoreName = "testers6 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false

	dbi.Close()
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var vis sdbi.Visitor
	vis.Origin = "test"
	vis.Host = "test"
	vis.IPAddress = "123.123.456.777"
	vis.StoreID = sid

	vsuc := si.AddVisit(&vis)
	if !vsuc {
		t.Fail()
	}

	var vis2 sdbi.Visitor
	vis2.Origin = "test"
	vis2.Host = "test"
	vis2.IPAddress = "123.123.456.777"
	vis2.StoreID = sid

	vsuc2 := si.AddVisit(&vis2)
	if !vsuc2 {
		t.Fail()
	}

	dbi.Close()
	vlst := si.GetVisitorData(sid)
	fmt.Println("vlst: ", *vlst)
	if len(*vlst) != 1 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
