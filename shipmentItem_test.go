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

func TestSix910Mysql_ShipmentItem(t *testing.T) {
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
	str.LocalDomain = "localhost29:8080"
	str.Logo = "some logo"
	str.OauthClientID = 33
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart29.com"
	str.State = "GA"
	str.StoreName = "testers29 fantastic store"
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

	var shp sdbi.Shipment
	shp.Boxes = 1
	shp.Insurance = 0.00
	shp.OrderID = oid
	shp.ShippingHandling = 5.26
	shp.Status = "Processing"

	shpsuc, shpid := si.AddShipment(&shp)
	if !shpsuc || shpid == 0 {
		t.Fail()
	}

	var sbox sdbi.ShipmentBox
	sbox.BoxNumber = 1
	sbox.ShipmentID = shpid
	sbox.ShippingAddress = "123 bates st, bates, CA"
	sbox.ShippingAddressID = 4
	sbox.ShippingMethodID = 55

	sboxsuc, sboxid := si.AddShipmentBox(&sbox)
	if !sboxsuc || sboxid == 0 {
		t.Fail()
	}

	var shpitm sdbi.ShipmentItem
	shpitm.OrderItemID = oiid
	shpitm.ShipmentID = shpid
	shpitm.ShipmentBoxID = sboxid

	shpitmsuc, shpitmid := si.AddShipmentItem(&shpitm)
	if !shpitmsuc || shpitmid == 0 {
		t.Fail()
	}

	shpitm.ID = shpitmid
	shpitm.Quantity = 2

	ushpitmsuc := si.UpdateShipmentItem(&shpitm)
	if !ushpitmsuc {
		t.Fail()
	}

	fshpitm := si.GetShipmentItem(shpitmid)
	fmt.Println("fshpitm: ", fshpitm)
	if fshpitm.ShipmentBoxID != shpitm.ShipmentBoxID {
		t.Fail()
	}

	fsitmList1 := si.GetShipmentItemList(shpid)
	fmt.Println("fsitmList1: ", fsitmList1)
	if len(*fsitmList1) != 1 {
		t.Fail()
	}

	fsitmList2 := si.GetShipmentItemListByBox(1)
	fmt.Println("fsitmList2: ", fsitmList2)
	if len(*fsitmList2) != 1 {
		t.Fail()
	}

	dlshpitmsuc := si.DeleteShipmentItem(shpitmid)
	if !dlshpitmsuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
