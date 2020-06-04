package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddCategory(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "this is a car category"
	cat.Image = "http://car/123/desc.png"
	cat.Thumbnail = "http://car/123/desc.png"
	cat.Name = "cars"
	cat.StoreID = sid

	ctsuc, ctid := si.AddCategory(&cat)
	if !ctsuc || ctid == 0 {
		t.Fail()
	}

	var cat2 sdbi.Category
	cat2.Description = "this is a car category2"
	cat2.Image = "http://car/123/desc2.png"
	cat2.Thumbnail = "http://car/123/desc2.png"
	cat2.Name = "cars2"
	cat2.StoreID = sid
	//cat2.ParentCategoryID = ctid

	ctsuc2, ctid2 := si.AddCategory(&cat2)
	if !ctsuc2 || ctid2 == 0 {
		t.Fail()
	}

	cat2.ID = ctid2
	cat2.Description = "this is a car category3"
	cat2.Image = "http://car/123/desc3.png"
	cat2.Thumbnail = "http://car/123/desc3.png"
	cat2.Name = "cars3"
	cat2.StoreID = sid
	cat2.ParentCategoryID = ctid
	uctsuc := si.UpdateCategory(&cat2)
	if !uctsuc {
		t.Fail()
	}

	fcat := si.GetCategory(ctid)
	fmt.Println("fcat", fcat)
	if fcat.StoreID != sid {
		t.Fail()
	}

	fcatList := si.GetCategoryList(sid)
	fmt.Println("fcatList", fcatList)
	if len(*fcatList) != 2 {
		t.Fail()
	}

	fscatList := si.GetSubCategoryList(ctid)
	fmt.Println("fscatList", fscatList)
	if len(*fscatList) != 1 {
		t.Fail()
	}

	dcatsuc := si.DeleteCategory(ctid)
	if !dcatsuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
