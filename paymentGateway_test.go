package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddPaymentGateway(t *testing.T) {
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
	str.LocalDomain = "localhost19:8080"
	str.Logo = "some logo"
	str.OauthClientID = 23
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart19.com"
	str.State = "GA"
	str.StoreName = "testers19 fantastic store"
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

	var pgw sdbi.PaymentGateway
	pgw.CheckoutURL = "/test"
	pgw.ClientID = "45"
	pgw.ClientKey = "5616516"
	pgw.LogoURL = "/login"
	pgw.PostOrderURL = "/post"
	pgw.StorePluginsID = spiid

	dbi.Close()
	pgwsuc, pgwid := si.AddPaymentGateway(&pgw)
	if !pgwsuc || pgwid == 0 {
		t.Fail()
	}

	pgw.ID = pgwid
	pgw.CheckoutURL = "/test2"
	pgw.ClientID = "452"
	pgw.ClientKey = "56165162"
	pgw.LogoURL = "/login2"
	pgw.PostOrderURL = "/post2"

	dbi.Close()
	upgwsuc := si.UpdatePaymentGateway(&pgw)
	if !upgwsuc {
		t.Fail()
	}

	dbi.Close()
	fpgw := si.GetPaymentGateway(pgwid)
	fmt.Println("fpgw: ", fpgw)
	if fpgw.ClientID != pgw.ClientID {
		t.Fail()
	}

	dbi.Close()
	fpgwliststr := si.GetPaymentGateways(sid)
	fmt.Println("fpgw: ", fpgwliststr)
	if len(*fpgwliststr) != 1 {
		t.Fail()
	}

	dbi.Close()
	dlpgw := si.DeletePaymentGateway(pgwid)
	if !dlpgw {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dlpi := si.DeletePlugin(piid)
	if !dlpi {
		t.Fail()
	}

	dbi.Close()

}
