package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddInsurance(t *testing.T) {
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
	str.LocalDomain = "localhost12:8080"
	str.Logo = "some logo"
	str.OauthClientID = 17
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart12.com"
	str.State = "GA"
	str.StoreName = "testers12 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var ins sdbi.Insurance
	ins.Cost = 4.95
	ins.MinOrderAmount = 50.00
	ins.MaxOrderAmount = 500.00
	ins.StoreID = sid

	inssuc, insid := si.AddInsurance(&ins)
	if !inssuc || insid == 0 {
		t.Fail()
	}

	ins.ID = insid
	ins.Cost = 6.25
	ins.MinOrderAmount = 80.00
	ins.MaxOrderAmount = 1000.00

	uinssuc := si.UpdateInsurance(&ins)
	if !uinssuc {
		t.Fail()
	}

	fins := si.GetInsurance(insid)
	fmt.Println("fins: ", fins)
	if fins.MaxOrderAmount != ins.MaxOrderAmount {
		t.Fail()
	}

	finslist := si.GetInsuranceList(sid)
	fmt.Println("finslist: ", finslist)
	if len(*finslist) != 1 {
		t.Fail()
	}

	dlins := si.DeleteInsurance(insid)
	if !dlins {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
