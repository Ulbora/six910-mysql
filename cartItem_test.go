package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddCartItem(t *testing.T) {
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
	str.LocalDomain = "localhost8:8080"
	str.Logo = "some logo"
	str.OauthClientID = 12
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart8.com"
	str.State = "GA"
	str.StoreName = "testers8 fantastic store"
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

	var cart sdbi.Cart
	cart.CustomerID = cid
	cart.StoreID = sid
	carsuc, carid := si.AddCart(&cart)
	if !carsuc || carid == 0 {
		t.Fail()
	}

	var dis sdbi.Distributor
	dis.Company = "abc supply"
	dis.ContactName = "Ricky Bobby"
	dis.Phone = "123-456-7891"
	dis.StoreID = sid

	dsuc, did := si.AddDistributor(&dis)
	if !dsuc || did == 0 {
		t.Fail()
	}

	var prod sdbi.Product
	prod.Color = "red"
	prod.Cost = 100.51
	prod.Currency = "USD"
	prod.Depth = 5.4
	prod.Desc = "some long desc about product"
	prod.DistributorID = did
	prod.Dropship = false
	prod.FreeShipping = true
	prod.Gtin = "44555ggggg"
	prod.Height = 22.3
	prod.Image1 = "image1"
	prod.Image2 = "image2"
	prod.Image3 = "image3"
	prod.Image4 = "image4"
	prod.Manufacturer = "some mfg"
	prod.Map = 150.99
	prod.Msrp = 185.99
	prod.MultiBox = false
	prod.Name = "some top product that sale well"
	//prod.ParentProductID
	prod.Price = 170.99
	prod.Promoted = true
	prod.SalePrice = 160.99
	prod.Searchable = true
	prod.ShipSeparately = false
	prod.ShippingMarkup = 3.40
	prod.ShortDesc = "short desc"
	prod.Size = "XL"
	prod.Sku = "123456789"
	prod.SpecialProcessing = true
	prod.SpecialProcessingType = "CODE 4"
	prod.Stock = 55
	prod.StockAlert = 10
	prod.StoreID = sid
	prod.Thumbnail = "someimage"
	prod.Visible = true
	prod.Weight = 15.4
	prod.Width = 22.4

	psuc, pid := si.AddProduct(&prod)
	if !psuc || pid == 0 {
		t.Fail()
	}

	dbi.Close()
	var citm sdbi.CartItem
	citm.CartID = carid
	citm.ProductID = pid
	citm.Quantity = 5

	itemSuc, itemid := si.AddCartItem(&citm)
	if !itemSuc || itemid == 0 {
		t.Fail()
	}

	citm.ID = itemid
	citm.Quantity = 15

	dbi.Close()
	uitemSuc := si.UpdateCartItem(&citm)
	if !uitemSuc {
		t.Fail()
	}

	dbi.Close()
	fitem1 := si.GetCarItem(carid, pid)
	fmt.Println("fitem1", fitem1)
	if fitem1.Quantity != 15 {
		t.Fail()
	}

	dbi.Close()
	fitemlist := si.GetCartItemList(carid)
	fmt.Println("fitemlist", fitemlist)
	if len(*fitemlist) != 1 {
		t.Fail()
	}

	dbi.Close()
	dlsuc := si.DeleteCartItem(itemid)
	if !dlsuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}
