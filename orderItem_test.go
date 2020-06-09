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

func TestSix910Mysql_AddOrderItem(t *testing.T) {
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
	str.LocalDomain = "localhost24:8080"
	str.Logo = "some logo"
	str.OauthClientID = 28
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart24.com"
	str.State = "GA"
	str.StoreName = "testers24 fantastic store"
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
	odr.CustomerID = 54
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

	odrsuc, oid := si.AddOrder(&odr)
	if !odrsuc || oid == 0 {
		t.Fail()
	}

	var oi sdbi.OrderItem
	oi.BackOrdered = false
	oi.Dropship = false
	oi.OrderID = oid
	oi.ProductID = 54
	oi.ProductName = "Printer"
	oi.ProductShortDesc = "HP deskjet printer"
	oi.Quantity = 1

	oisuc, oiid := si.AddOrderItem(&oi)
	if !oisuc || oiid == 0 {
		t.Fail()
	}

	oi.ID = oiid
	oi.BackOrdered = true
	oi.Dropship = true
	oi.Quantity = 10

	uoissuc := si.UpdateOrderItem(&oi)
	if !uoissuc {
		t.Fail()
	}

	foi := si.GetOrderItem(oiid)
	fmt.Println("foi: ", foi)
	if foi.ProductName != oi.ProductName || foi.Quantity != oi.Quantity {
		t.Fail()
	}

	foilist := si.GetOrderItemList(oid)
	fmt.Println("foilist: ", foilist)
	if len(*foilist) != 1 {
		t.Fail()
	}

	dloisuc := si.DeleteOrderItem(oiid)
	if !dloisuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
