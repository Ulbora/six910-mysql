package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_Address(t *testing.T) {
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
	str.LocalDomain = "localhost4:8080"
	str.Logo = "some logo"
	str.OauthClientID = 8
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart4.com"
	str.State = "GA"
	str.StoreName = "testers4 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var cus sdbi.Customer
	cus.Email = "tester@test.com"
	cus.ResetPassword = false
	cus.FirstName = "tester"
	cus.LastName = "bob"
	cus.Company = "teste company"
	cus.City = "testvill"
	cus.State = "GA"
	cus.Zip = "30147"
	cus.Phone = "123-456-7897"
	cus.StoreID = sid

	csuc, cid := si.AddCustomer(&cus)
	if !csuc || cid == 0 {
		t.Fail()
	}

	var add sdbi.Address
	add.Address = "12345 shootout lane"
	add.City = "gunville"
	add.State = "LA"
	add.Zip = "12345"
	add.County = "Polk"
	add.Country = "US"
	add.Type = "BILLING"
	add.CustomerID = cid
	add.Attr1 = "one"
	add.Attr2 = "two"
	add.Attr3 = "three"
	add.Attr4 = "four"

	dbi.Close()
	asuc, aid := si.AddAddress(&add)
	if !asuc || aid == 0 {
		t.Fail()
	}

	add.ID = aid
	add.Address = "12345 shootout st"
	add.City = "gunerville"
	add.State = "CA"
	add.Zip = "45678"
	add.County = "Smoke"
	add.Country = "CA"
	add.Type = "SHIPPING"

	dbi.Close()
	ausuc := si.UpdateAddress(&add)
	if !ausuc {
		t.Fail()
	}

	dbi.Close()
	radd := si.GetAddress(aid)
	fmt.Println("radd", radd)
	if radd.CustomerID != cid || radd.Attr4 != "four" {
		t.Fail()
	}

	dbi.Close()
	raddlist := si.GetAddressList(cid)
	fmt.Println("raddlist", raddlist)
	if len(*raddlist) != 1 {
		t.Fail()
	}

	dbi.Close()
	dsuc := si.DeleteAddress(aid)
	if !dsuc {
		t.Fail()
	}

	raddlist2 := si.GetAddressList(cid)
	fmt.Println("raddlist", raddlist2)
	if len(*raddlist2) != 0 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}
