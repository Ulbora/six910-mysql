package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddTaxRate(t *testing.T) {
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
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var txr sdbi.TaxRate
	txr.Country = "USA"
	txr.State = "GA"
	txr.PercentRate = 7
	txr.TaxType = "Sales"
	txr.StoreID = sid

	dbi.Close()
	psuc, tid := si.AddTaxRate(&txr)
	if !psuc || tid == 0 {
		t.Fail()
	}

	txr.ID = tid
	txr.PercentRate = 7
	txr.IncludeHandling = true
	txr.IncludeShipping = true
	txr.ProductCategoryID = 3
	txr.ZipStart = "12345"
	txr.ZipEnd = "23456"

	dbi.Close()
	usuc := si.UpdateTaxRate(&txr)
	if !usuc {
		t.Fail()
	}

	dbi.Close()
	txlst1 := si.GetTaxRate("USA", "GA", sid)
	fmt.Println("txlst1", txlst1)
	if len(*txlst1) != 1 {
		t.Fail()
	}

	dbi.Close()
	txlst2 := si.GetTaxRateList(sid)
	fmt.Println("txlst2", txlst2)
	if len(*txlst2) != 1 {
		t.Fail()
	}

	dbi.Close()
	dprodSuc := si.DeleteTaxRate(tid)
	if !dprodSuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
