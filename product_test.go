package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddProduct(t *testing.T) {
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
	str.LocalDomain = "localhost6:8080"
	str.Logo = "some logo"
	str.OauthClientID = 10
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart6.com"
	str.State = "GA"
	str.StoreName = "testers6 fantastic store"
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
	prod.ShipSeperately = false
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

	dbi.Close()
	psuc, pid := si.AddProduct(&prod)
	if !psuc || pid == 0 {
		t.Fail()
	}

	prod.Color = "green"
	prod.ParentProductID = pid
	psuc2, pid2 := si.AddProduct(&prod)
	if !psuc2 || pid2 == 0 {
		t.Fail()
	}

	prod.ID = pid
	prod.Color = "red2"
	prod.Cost = 100.52
	prod.Currency = "US2"
	prod.Depth = 5.42
	prod.Desc = "some long desc about product2"
	prod.DistributorID = did
	prod.Dropship = true
	prod.FreeShipping = false
	prod.Gtin = "44555ggggg2"
	prod.Height = 22.32
	prod.Image1 = "image12"
	prod.Image2 = "image22"
	prod.Image3 = "image32"
	prod.Image4 = "image42"
	prod.Manufacturer = "some mfg2"
	prod.Map = 150.92
	prod.Msrp = 185.92
	prod.MultiBox = true
	prod.Name = "some top product that sale well2"
	//prod.ParentProductID
	prod.Price = 170.92
	prod.Promoted = false
	prod.SalePrice = 160.92
	prod.Searchable = false
	prod.ShipSeperately = true
	prod.ShippingMarkup = 3.42
	prod.ShortDesc = "short desc2"
	prod.Size = "XL2"
	prod.Sku = "1234567892"
	prod.SpecialProcessing = false
	prod.SpecialProcessingType = "CODE 42"
	prod.Stock = 552
	prod.StockAlert = 102
	prod.StoreID = sid
	prod.Thumbnail = "someimage2"
	prod.Visible = false
	prod.Weight = 15.42
	prod.Width = 22.42
	prod.ParentProductID = 0

	dbi.Close()
	usuc := si.UpdateProduct(&prod)
	if !usuc {
		t.Fail()
	}

	dbi.Close()
	fprod := si.GetProductByID(pid)
	fmt.Println("fprod", fprod)
	if fprod.ID != pid {
		t.Fail()
	}
	if fprod.DistributorID != did {
		t.Fail()
	}
	if fprod.StoreID != sid {
		t.Fail()
	}
	if fprod.Map != prod.Map {
		t.Fail()
	}
	if fprod.MultiBox != prod.MultiBox {
		t.Fail()
	}

	dbi.Close()
	fprodn := si.GetProductsByName("well2", sid, 0, 100)
	fmt.Println("fprodn", fprodn)
	if len(*fprodn) != 1 {
		t.Fail()
	}

	//add cat here

	var cat sdbi.Category
	cat.Description = "this is a car category"
	cat.Image = "http://car/123/desc.png"
	cat.Thumbnail = "http://car/123/desc.png"
	cat.Name = "cars"
	cat.StoreID = sid

	dbi.Close()
	ctsuc, ctid := si.AddCategory(&cat)
	if !ctsuc || ctid == 0 {
		t.Fail()
	}

	dbi.Close()
	var pcat sdbi.ProductCategory
	pcat.CategoryID = ctid
	pcat.ProductID = pid

	pcatsuc := si.AddProductCategory(&pcat)
	if !pcatsuc {
		t.Fail()
	}

	dbi.Close()
	prodCatlist := si.GetProductsByCaterory(ctid, 0, 100)
	fmt.Println("prodCatlist", prodCatlist)
	if len(*prodCatlist) != 1 {
		t.Fail()
	}

	dbi.Close()
	prodStr := si.GetProductList(sid, 0, 100)
	fmt.Println("prodStr", prodStr)
	if len(*prodStr) != 2 {
		t.Fail()
	}

	dbi.Close()
	dprodCatSuc := si.DeleteProductCategory(&pcat)
	if !dprodCatSuc {
		t.Fail()
	}

	dbi.Close()
	dprodSuc := si.DeleteProduct(pid)
	if !dprodSuc {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}
