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

func TestSix910Mysql_AddDataStoreWriteLock(t *testing.T) {
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
	str.LocalDomain = "localhost22:8080"
	str.Logo = "some logo"
	str.OauthClientID = 26
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart22.com"
	str.State = "GA"
	str.StoreName = "testers22 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var lc sdbi.DataStoreWriteLock
	lc.DataStoreName = "content"
	lc.Locked = true
	lc.LockedByUser = "tester"
	lc.LockedInstanceName = "inst-1"
	lc.LockedTime = time.Now()
	lc.StoreID = sid

	dbi.Close()
	lcsuc := si.AddDataStoreWriteLock(&lc)
	if !lcsuc {
		t.Fail()
	}

	lc.Locked = false
	lc.LockedByUser = "tester2"
	lc.LockedInstanceName = "inst-12"
	lc.LockedTime = time.Now().Add(time.Minute * 10)

	dbi.Close()
	ulcsuc := si.UpdateDataStoreWriteLock(&lc)
	if !ulcsuc {
		t.Fail()
	}

	dbi.Close()
	flc := si.GetDataStoreWriteLock("content", sid)
	fmt.Println("flc: ", flc)
	if flc.LockedInstanceName != lc.LockedInstanceName {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
