package six910mysql

import (
	"fmt"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_GetProductManufacturerListByProductName(t *testing.T) {
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
	str.LocalDomain = "localhost66:8080"
	str.Logo = "some logo"
	str.OauthClientID = 10
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart66.com"
	str.State = "GA"
	str.StoreName = "testers6 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
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

	dbi.Close()
	psuc, pid := si.AddProduct(&prod)
	if !psuc || pid == 0 {
		t.Fail()
	}

	dbi.Close()
	mlst := si.GetProductManufacturerListByProductName("top product", sid)
	fmt.Println("mlst: ", mlst)
	if len(*mlst) != 1 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()

}

func TestSix910Mysql_GetProductByNameAndManufacturerName(t *testing.T) {
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
	str.LocalDomain = "localhost67:8080"
	str.Logo = "some logo"
	str.OauthClientID = 10
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart67.com"
	str.State = "GA"
	str.StoreName = "testers6 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
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

	dbi.Close()
	psuc, pid := si.AddProduct(&prod)
	if !psuc || pid == 0 {
		t.Fail()
	}

	dbi.Close()
	mlst := si.GetProductByNameAndManufacturerName("some mfg", "top product", sid, 0, 10)
	fmt.Println("mlst: ", mlst)
	if len(*mlst) != 1 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}

func TestSix910Mysql_GetProductManufacturerListByCatID(t *testing.T) {
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
	str.LocalDomain = "localhost68:8080"
	str.Logo = "some logo"
	str.OauthClientID = 10
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart68.com"
	str.State = "GA"
	str.StoreName = "testers6 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
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

	dbi.Close()
	psuc, pid := si.AddProduct(&prod)
	if !psuc || pid == 0 {
		t.Fail()
	}

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
	mlst := si.GetProductManufacturerListByCatID(ctid, sid)
	fmt.Println("mlst: ", mlst)
	if len(*mlst) != 1 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}

func TestSix910Mysql_GetProductByCatAndManufacturer(t *testing.T) {
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
	str.LocalDomain = "localhost69:8080"
	str.Logo = "some logo"
	str.OauthClientID = 10
	str.OauthSecret = "this is secret"
	str.RemoteDomain = "www.someCart69.com"
	str.State = "GA"
	str.StoreName = "testers6 fantastic store"
	str.StoreSlogan = "we test for less"
	str.Zip = "30036"
	str.Enabled = false
	suc, sid := si.AddStore(&str)
	if !suc || sid == 0 {
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

	dbi.Close()
	psuc, pid := si.AddProduct(&prod)
	if !psuc || pid == 0 {
		t.Fail()
	}

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
	mlst := si.GetProductByCatAndManufacturer(ctid, "some mfg", sid, 0, 10)
	fmt.Println("mlst: ", mlst)
	if len(*mlst) != 1 {
		t.Fail()
	}

	dssuc := si.DeleteStore(sid)
	fmt.Println("delete store in customer suc: ", dssuc)
	if !dssuc {
		t.Fail()
	}

	dbi.Close()
}
