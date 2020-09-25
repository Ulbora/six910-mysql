package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_Customer(t *testing.T) {
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
	str.LocalDomain = "localhost2:8080"
	str.Logo = "some logo"
	str.OauthClientID = 6
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart2.com"
	str.State = "GA"
	str.StoreName = "testers2 fantastic store"
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

	dbi.Close()
	csuc, cid := si.AddCustomer(&cus)
	if !csuc || cid == 0 {
		t.Fail()
	}

	cus.ID = cid
	cus.Email = "tester2@test.com"
	cus.ResetPassword = true
	cus.FirstName = "tester2"
	cus.LastName = "bob2"
	cus.Company = "teste company2"
	cus.City = "testvill2"
	cus.State = "G2"
	cus.Zip = "301472"
	cus.Phone = "123-456-7892"

	dbi.Close()
	ucsuc := si.UpdateCustomer(&cus)
	if !ucsuc {
		t.Fail()
	}

	dbi.Close()
	fcus := si.GetCustomer("tester2@test.com", sid)
	if fcus.ID != cid {
		t.Fail()
	}

	dbi.Close()
	cusList := si.GetCustomerList(sid, 0, 10)
	fmt.Println("cusList", cusList)
	if len(*cusList) != 1 {
		t.Fail()
	}

	dbi.Close()
	fcus2 := si.GetCustomerID(cid)
	fmt.Println("fcus2", fcus2)
	if fcus2.StoreID != sid {
		t.Fail()
	}

	dbi.Close()
	dsuc := si.DeleteCustomer(cid)
	if !dsuc {
		t.Fail()
	}

	cusList2 := si.GetCustomerList(sid, 0, 10)
	fmt.Println("cusList", cusList2)
	if len(*cusList2) != 0 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
