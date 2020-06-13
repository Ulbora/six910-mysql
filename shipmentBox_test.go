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

func TestSix910Mysql_AddShipmentBox(t *testing.T) {
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
	str.LocalDomain = "localhost28:8080"
	str.Logo = "some logo"
	str.OauthClientID = 32
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart28.com"
	str.State = "GA"
	str.StoreName = "testers28 fantastic store"
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

	dbi.Close()
	sboxsuc, sboxid := si.AddShipmentBox(&sbox)
	if !sboxsuc || sboxid == 0 {
		t.Fail()
	}

	sbox.ID = sboxid
	sbox.Dropship = false
	sbox.Cost = 5.55

	dbi.Close()
	usboxsuc := si.UpdateShipmentBox(&sbox)
	if !usboxsuc {
		t.Fail()
	}

	sbox.TrackingNumber = "123track"
	sbox.Weight = 4.2
	sbox.Width = 12
	sbox.Height = 18
	sbox.Depth = 16
	sbox.Insurance = 4.56

	usboxsuc2 := si.UpdateShipmentBox(&sbox)
	if !usboxsuc2 {
		t.Fail()
	}

	dbi.Close()
	fsbox := si.GetShipmentBox(sboxid)
	fmt.Println("fsbox: ", fsbox)
	if fsbox.BoxNumber != sbox.BoxNumber {
		t.Fail()
	}

	var sbox2 sdbi.ShipmentBox
	sbox2.BoxNumber = 2
	sbox2.ShipmentID = shpid
	sbox2.ShippingAddress = "123 bates st, bates, CA"
	sbox2.ShippingAddressID = 4
	sbox2.ShippingMethodID = 55

	sboxsuc2, sboxid2 := si.AddShipmentBox(&sbox2)
	if !sboxsuc2 || sboxid2 == 0 {
		t.Fail()
	}

	dbi.Close()
	fsboxlist := si.GetShipmentBoxList(shpid)
	fmt.Println("fsboxlist: ", fsboxlist)
	if len(*fsboxlist) != 2 {
		t.Fail()
	}

	fsbox2 := si.GetShipmentBox(sboxid2)
	fmt.Println("fsbox2: ", fsbox2)
	if fsbox2.BoxNumber != sbox2.BoxNumber {
		t.Fail()
	}

	dbi.Close()
	dlsboxsuc := si.DeleteShipmentBox(sboxid2)
	if !dlsboxsuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}
