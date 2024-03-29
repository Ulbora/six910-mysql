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

func TestSix910Mysql_Order(t *testing.T) {
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
	str.LocalDomain = "localhost23:8080"
	str.Logo = "some logo"
	str.OauthClientID = 27
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart23.com"
	str.State = "GA"
	str.StoreName = "testers23 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
		t.Fail()
	}

	var odr sdbi.Order
	odr.BillingAddress = "133 test st. Dallas, TX"
	odr.BillingAddressID = 3
	odr.CustomerID = 44
	odr.CustomerName = "Bills Cars"
	odr.Insurance = 0.0
	odr.OrderDate = time.Now()
	odr.OrderNumber = "S23433LC456"
	odr.OrderType = "Standard"
	odr.Pickup = false
	odr.ShippingAddress = "133 test st. Dallas, TX"
	odr.ShippingAddressID = 44
	odr.ShippingHandling = 5.55
	odr.Status = "Processing"
	odr.StoreID = sid
	odr.Subtotal = 141.05
	odr.Taxes = 5.27
	odr.Total = 146.32
	odr.Username = "billybob"
	odr.FFLShippingAddressID = 21
	odr.FFLShippingAddress = "My FFL"
	odr.FFLName = "bobs guns"
	odr.FFLLic = "abcd"
	odr.FFLExpDate = "tomorrow"
	odr.FFLPhone = "123456789"

	dbi.Close()
	odrsuc, oid := si.AddOrder(&odr)
	if !odrsuc || oid == 0 {
		t.Fail()
	}

	odr.ID = oid
	odr.BillingAddress = "133 test st. Dallas, TX2"
	odr.BillingAddressID = 32
	odr.CustomerID = 44
	odr.CustomerName = "Bills Cars2"
	odr.Insurance = 0.20
	odr.Updated = time.Now().Add(time.Minute * 45)
	odr.OrderNumber = "S23433LC4562"
	odr.OrderType = "Standard2"
	odr.Pickup = true
	odr.ShippingAddress = "133 test st. Dallas, TX2"
	odr.ShippingAddressID = 442
	odr.ShippingHandling = 5.52
	odr.Status = "Processing2"
	odr.StoreID = sid
	odr.Subtotal = 141.02
	odr.Taxes = 5.22
	odr.Total = 146.22
	odr.Username = "billybob2"
	odr.ShippingMethodID = 23
	odr.ShippingMethodName = "UPS Ground"
	odr.Refunded = 15.00
	odr.FFLShippingAddressID = 22
	odr.FFLShippingAddress = "My New FFL"
	odr.FFLName = "bobs guns 2"
	odr.FFLLic = "abcd2"
	odr.FFLExpDate = "tomorrow2"
	odr.FFLPhone = "123456789222"

	dbi.Close()
	uodrsuc := si.UpdateOrder(&odr)
	if !uodrsuc {
		t.Fail()
	}

	dbi.Close()
	fodr := si.GetOrder(oid)
	fmt.Println("fodr: ", fodr)
	fmt.Println("fodr fflid: ", fodr.FFLShippingAddressID)
	fmt.Println("fodr ffl: ", fodr.FFLShippingAddress)
	fmt.Println("Refunded: ", fodr.Refunded)
	if fodr.CustomerID != odr.CustomerID || fodr.OrderType != odr.OrderType || fodr.Refunded != 15.00 {
		t.Fail()
	}

	dbi.Close()
	fodrlist := si.GetOrderList(44, sid)
	fmt.Println("fodrlist: ", fodrlist)
	if len(*fodrlist) != 1 {
		t.Fail()
	}

	dbi.Close()
	fodrlist2 := si.GetStoreOrderList(sid)
	fmt.Println("fodrlist2: ", fodrlist2)
	if len(*fodrlist2) != 1 {
		t.Fail()
	}

	dbi.Close()
	fodrlist3 := si.GetStoreOrderListByStatus("Processing2", sid)
	fmt.Println("fodrlist3: ", fodrlist3)
	if len(*fodrlist3) != 1 {
		t.Fail()
	}

	dbi.Close()
	fodrlist4 := si.GetStoreOrderListByStatus("new", sid)
	fmt.Println("fodrlist4: ", fodrlist4)
	if len(*fodrlist4) != 0 {
		t.Fail()
	}

	dbi.Close()
	cntLst := si.GetOrderCountData(sid)
	fmt.Println("cntLst: ", cntLst)
	if len(*cntLst) != 1 {
		t.Fail()
	}

	dbi.Close()
	salLst := si.GetOrderSalesData(sid)
	fmt.Println("salLst: ", salLst)
	if len(*salLst) != 1 || (*salLst)[0].OrderTotal != 146.22 {
		t.Fail()
	}

	dbi.Close()
	dlodr := si.DeleteOrder(oid)
	if !dlodr {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
	// t.Fail()
}
