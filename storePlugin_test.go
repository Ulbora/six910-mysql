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

func TestSix910Mysql_AddStorePlugin(t *testing.T) {
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

	var pi sdbi.Plugins

	pi.ActivateURL = "test"
	pi.Category = "catalog"
	pi.ContactPhone = "123-789-4555"
	pi.Developer = "Ulbora Labs LLC"
	pi.DeveloperAddress = "atlanta, GA"
	pi.Enabled = true
	pi.Fee = 0
	pi.IsPGW = false
	pi.OauthRedirectURL = "/redirect"
	pi.PluginName = "Catalog Easy"

	pisuc, piid := si.AddPlugin(&pi)
	if !pisuc || piid == 0 {
		t.Fail()
	}

	var str sdbi.Store
	str.City = "test vill"
	str.Company = "test com"
	str.Currency = "USD"
	str.Email = "tester@tester.com"
	str.FirstName = "Tester"
	str.LastName = "Bill"
	str.LocalDomain = "localhost18:8080"
	str.Logo = "some logo"
	str.OauthClientID = 22
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart18.com"
	str.State = "GA"
	str.StoreName = "testers18 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var spi sdbi.StorePlugins
	spi.APIKey = "123"
	spi.ActivateURL = "/test"
	spi.Active = false
	spi.Category = "inventory"
	spi.IframeURL = "/iframe"
	spi.IsPGW = true
	spi.MenuIconURL = "/menu"
	spi.MenuTitle = "somemenu"
	spi.OauthClientID = 334
	spi.OauthRedirectURL = "/redirect"
	spi.OauthSecret = "123456kjh"
	spi.PluginID = piid
	spi.PluginName = "the great plugin"
	spi.StoreID = sid

	spisuc, spiid := si.AddStorePlugin(&spi)
	if !spisuc || spiid == 0 {
		t.Fail()
	}

	spi.ID = spiid
	spi.APIKey = "1235"
	spi.ActivateURL = "/test5"
	spi.Active = true
	spi.IframeURL = "/iframe5"
	spi.MenuIconURL = "/menu5"
	spi.MenuTitle = "somemenu5"
	spi.OauthClientID = 3345
	spi.OauthRedirectURL = "/redirect5"
	spi.OauthSecret = "123456kjh5"

	uspisuc := si.UpdateStorePlugin(&spi)
	if !uspisuc {
		t.Fail()
	}

	spi.APIKey = "12357"
	spi.ActivateURL = "/test57"
	spi.Active = true
	spi.IframeURL = "/iframe57"
	spi.MenuIconURL = "/menu57"
	spi.MenuTitle = "somemenu57"
	spi.OauthClientID = 33457
	spi.OauthRedirectURL = "/redirect57"
	spi.OauthSecret = "123456kjh57"
	spi.RekeyDate = time.Now()
	spi.RekeyTryCount = 1

	uspisuc2 := si.UpdateStorePlugin(&spi)
	if !uspisuc2 {
		t.Fail()
	}

	fspi := si.GetStorePlugin(spiid)
	fmt.Println("fspi: ", fspi)
	if fspi.IframeURL != spi.IframeURL {
		t.Fail()
	}

	fspilist := si.GetStorePluginList(sid)
	fmt.Println("fspilist: ", fspilist)
	if len(*fspilist) != 1 {
		t.Fail()
	}

	dlspisuc := si.DeleteStorePlugin(spiid)
	if !dlspisuc {
		t.Fail()
	}

	dlpi := si.DeletePlugin(piid)
	if !dlpi {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}
