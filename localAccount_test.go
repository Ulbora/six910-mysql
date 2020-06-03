package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddLocalAccount(t *testing.T) {
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
	str.LocalDomain = "localhost3:8080"
	str.Logo = "some logo"
	str.OauthClientID = 7
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart3.com"
	str.State = "GA"
	str.StoreName = "testers3 fantastic store"
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

	var lac sdbi.LocalAccount
	lac.CustomerID = cid
	lac.StoreID = sid
	lac.Enabled = true
	lac.Password = "password"
	lac.UserName = "someuser"
	lac.Role = "customer"

	lacsuc := si.AddLocalAccount(&lac)
	if !lacsuc {
		t.Fail()
	}

	lac.Enabled = false
	lac.Password = "password2"
	lac.Role = "customer2"

	lacUsuc := si.UpdateLocalAccount(&lac)
	if !lacUsuc {
		t.Fail()
	}

	rlac := si.GetLocalAccount("someuser", sid)
	fmt.Println("rlac", rlac)
	if rlac.CustomerID != cid {
		t.Fail()
	}
}
