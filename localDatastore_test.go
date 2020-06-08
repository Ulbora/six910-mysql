package six910mysql

import (
	"fmt"
	"testing"
	"time"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddLocalDatastore(t *testing.T) {
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
	str.LocalDomain = "localhost20:8080"
	str.Logo = "some logo"
	str.OauthClientID = 24
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart20.com"
	str.State = "GA"
	str.StoreName = "testers20 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var lds sdbi.LocalDataStore
	lds.DataStoreName = "content"
	lds.StoreID = sid
	lds.Reload = false
	lds.ReloadDate = time.Now()

	ldssuc := si.AddLocalDatastore(&lds)
	if !ldssuc {
		t.Fail()
	}

	ldssuc = si.AddLocalDatastore(&lds)
	if ldssuc {
		t.Fail()
	}

	lds.Reload = true
	lds.ReloadDate = time.Now().Add(time.Minute * 10)

	uldssuc := si.UpdateLocalDatastore(&lds)
	if !uldssuc {
		t.Fail()
	}

	flds := si.GetLocalDatastore(sid, "content")
	fmt.Println("flds: ", flds)
	if flds.Reload != lds.Reload {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
