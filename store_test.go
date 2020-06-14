package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddStore(t *testing.T) {
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
	str.LocalDomain = "localhost:8080"
	str.Logo = "some logo"
	str.OauthClientID = 5
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart.com"
	str.State = "GA"
	str.StoreName = "testers fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false

	dbi.Close()
	suc, id := si.AddStore(&str)
	if !suc || id == 0 {
		t.Fail()
	}

	dbi.Close()
	fstr := si.GetStore("testers fantastic store")
	fmt.Println("found store: ", fstr)
	if fstr.ID == 0 {
		t.Fail()
	}

	dbi.Close()
	fstrid := si.GetStoreID(id)
	fmt.Println("found store by id: ", fstrid)
	if fstrid.ID == 0 {
		t.Fail()
	}

	str.ID = id
	str.City = "test vill2"
	str.Company = "test com2"
	str.Currency = "USD2"
	str.Email = "tester@tester.com2"
	str.FirstName = "Tester2"
	str.LastName = "Bill2"
	str.LocalDomain = "localhost:80802"
	str.Logo = "some logo2"
	str.OauthClientID = 55
	str.OauthSecret = "this is secret2"
	str.RemoteDomain = "www.someCart.com2"
	str.State = "GA2"
	str.StoreName = "testers fantastic store2"
	str.StoreSlogan = "we test for less2"
	str.Zip = "300362"
	str.Enabled = true

	dbi.Close()
	sucu := si.UpdateStore(&str)
	if !sucu {
		t.Fail()
	}

	dbi.Close()
	fstr2 := si.GetStoreByLocal("localhost:80802")
	fmt.Println("found store by local domain: ", fstr2)
	if fstr2.ID == 0 {
		t.Fail()
	}

	dbi.Close()
	cnt := si.GetStoreCount()
	fmt.Println("store count: ", cnt)
	if cnt == 0 {
		t.Fail()
	}

	dbi.Close()
	dsuc := si.DeleteStore(id)
	fmt.Println("delete suc: ", dsuc)
	if !dsuc {
		t.Fail()
	}

	dbi.Close()
}
