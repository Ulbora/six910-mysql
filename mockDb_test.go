package six910mysql

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestMockSix910Mysql_Mocks(t *testing.T) {
	var sdb MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l

	si := sdb.GetNew()

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockAddSecuritySuccess = true
	sdb.MockSecurityID = 5
	secsuc, secid := si.AddSecurity(&sec)
	if !secsuc || secid == 0 {
		t.Fail()
	}

	sdb.MockUpdateSecuritySuccess = true
	usecsuc := si.UpdateSecurity(&sec)
	if !usecsuc {
		t.Fail()
	}

	sdb.MockSecurity = &sec
	fsec := si.GetSecurity()
	if !fsec.OauthOn {
		t.Fail()
	}

	sdb.MockDeleteSecuritySuccess = true
	dlsec := si.DeleteSecurity()
	if !dlsec {
		t.Fail()
	}

	var str sdbi.Store
	str.City = "test vill"
	str.StoreName = "test"

	sdb.MockAddStoreSuccess = true
	sdb.MockStoreID = 3

	strsuc, sid := si.AddStore(&str)
	if !strsuc || sid != 3 {
		t.Fail()
	}

	sdb.MockUpdateStoreSuccess = true
	mstrsuc := si.UpdateStore(&str)
	if !mstrsuc {
		t.Fail()
	}

	sdb.MockStore = &str
	fstr := si.GetStore("test")
	if fstr.StoreName != "test" {
		t.Fail()
	}

	sdb.MockStore = &str
	flstr := si.GetLocalStore()
	if flstr.StoreName != "test" {
		t.Fail()
	}

	fstr2 := si.GetStoreID(1)
	if fstr2.StoreName != "test" {
		t.Fail()
	}

	fstr3 := si.GetStoreByLocal("test")
	if fstr3.StoreName != "test" {
		t.Fail()
	}

	sdb.MockStoreCount = 3
	scnt := si.GetStoreCount()
	if scnt != 3 {
		t.Fail()
	}

	sdb.MockDeleteStoreSuccess = true

	dlstr := si.DeleteStore(1)
	if !dlstr {
		t.Fail()
	}

	var cus sdbi.Customer
	cus.Email = "tester@test.com"

	sdb.MockAddCustomerSuccess = true
	sdb.MockCustomerID = 4

	cussuc, cid := si.AddCustomer(&cus)
	if !cussuc || cid == 0 {
		t.Fail()
	}

	sdb.MockUpdateCustomerSuccess = true
	ucussuc := si.UpdateCustomer(&cus)
	if !ucussuc {
		t.Fail()
	}

	sdb.MockCustomer = &cus
	fcus := si.GetCustomer("test", 1)
	if fcus.Email != cus.Email {
		t.Fail()
	}

	fcus2 := si.GetCustomerID(1)
	if fcus2.Email != cus.Email {
		t.Fail()
	}

	var cusList []sdbi.Customer
	cusList = append(cusList, cus)
	sdb.MockCustomerList = &cusList
	fcuslist := si.GetCustomerList(3)
	if len(*fcuslist) != 1 {
		t.Fail()
	}

	sdb.MockDeleteCustomerSuccess = true
	dlcus := si.DeleteCustomer(1)
	if !dlcus {
		t.Fail()
	}

	var lac sdbi.LocalAccount
	lac.Enabled = true

	sdb.MockAddLocalAccountSuccess = true
	lacsuc := si.AddLocalAccount(&lac)
	if !lacsuc {
		t.Fail()
	}

	sdb.MockUpdateLocalAccountSuccess = true
	ulacsuc := si.UpdateLocalAccount(&lac)
	if !ulacsuc {
		t.Fail()
	}

	sdb.MockLocalAccount = &lac
	flac := si.GetLocalAccount("test", 1)
	if flac.Enabled != lac.Enabled {
		t.Fail()
	}

	var lclist []sdbi.LocalAccount
	lclist = append(lclist, lac)
	sdb.MockLocalAccountList = &lclist
	flaclist := si.GetLocalAccountList(2)
	if len(*flaclist) != 1 {
		t.Fail()
	}

	sdb.MockDeleteLocalAccountSuccess = true
	dllacsuc := si.DeleteLocalAccount("test", 4)
	if !dllacsuc {
		t.Fail()
	}

	var dis sdbi.Distributor
	dis.Company = "abc supply"

	sdb.MockAddDistributorSuccess = true
	sdb.MockDistributorID = 4
	dissuc, disid := si.AddDistributor(&dis)
	if !dissuc || disid == 0 {
		t.Fail()
	}

	sdb.MockUpdateDistributorSuccess = true
	udissuc := si.UpdateDistributor(&dis)
	if !udissuc {
		t.Fail()
	}

	sdb.MockDistributor = &dis
	fdis := si.GetDistributor(5)
	if fdis.Company != dis.Company {
		t.Fail()
	}

	var dislist []sdbi.Distributor
	dislist = append(dislist, dis)
	sdb.MockDistributorList = &dislist
	fdislst := si.GetDistributorList(1)
	if len(*fdislst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteDistributorSuccess = true
	dldissuc := si.DeleteDistributor(3)
	if !dldissuc {
		t.Fail()
	}

	var cart sdbi.Cart
	cart.CustomerID = 5

	sdb.MockAddCartSuccess = true
	sdb.MockCartID = 2
	cartsuc, cartid := si.AddCart(&cart)
	if !cartsuc || cartid == 0 {
		t.Fail()
	}

	sdb.MockUpdateCartSuccess = true
	ucart := si.UpdateCart(&cart)
	if !ucart {
		t.Fail()
	}
	sdb.MockCart = &cart
	fcart := si.GetCart(4)
	if fcart.CustomerID != cart.CustomerID {
		t.Fail()
	}

	sdb.MockDeleteCartSuccess = true
	dlcartsuc := si.DeleteCart(5)
	if !dlcartsuc {
		t.Fail()
	}

	var citm sdbi.CartItem
	citm.CartID = 6

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 4
	citmsuc, citmid := si.AddCartItem(&citm)
	if !citmsuc || citmid == 0 {
		t.Fail()
	}

	sdb.MockUpdateCartItemSuccess = true
	ucitmsuc := si.UpdateCartItem(&citm)
	if !ucitmsuc {
		t.Fail()
	}

	sdb.MockCartItem = &citm
	fcitm := si.GetCarItem(3, 4)
	if fcitm.CartID != citm.CartID {
		t.Fail()
	}

	var citmlist []sdbi.CartItem
	citmlist = append(citmlist, citm)
	sdb.MockCartItemList = &citmlist

	fcitmlist := si.GetCartItemList(4)
	if len(*fcitmlist) != 1 {
		t.Fail()
	}

	sdb.MockDeleteCartItemSuccess = true
	dlcitmsuc := si.DeleteCartItem(3)
	if !dlcitmsuc {
		t.Fail()
	}

	var add sdbi.Address
	add.Address = "12345 shootout lane"

	sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 6

	addsuc, addid := si.AddAddress(&add)
	if !addsuc || addid == 0 {
		t.Fail()
	}

	sdb.MockUpdateAddressSuccess = true
	uaddsuc := si.UpdateAddress(&add)
	if !uaddsuc {
		t.Fail()
	}

	sdb.MockAddress = &add
	fadd := si.GetAddress(4)
	if fadd.Address != add.Address {
		t.Fail()
	}

	var addlst []sdbi.Address

	addlst = append(addlst, add)

	sdb.MockAddressList = &addlst

	addlist := si.GetAddressList(5)
	if len(*addlist) != 1 {
		t.Fail()
	}

	sdb.MockDeleteAddressSuccess = true
	dladdsuc := si.DeleteAddress(4)
	if !dladdsuc {
		t.Fail()
	}

	var cat sdbi.Category
	cat.Description = "this is a car category"

	sdb.MockAddCategorySuccess = true
	sdb.MockCategoryID = 7
	catsuc, catid := si.AddCategory(&cat)
	if !catsuc || catid == 0 {
		t.Fail()
	}

	sdb.MockUpdateCategorySuccess = true
	ucatsuc := si.UpdateCategory(&cat)
	if !ucatsuc {
		t.Fail()
	}

	sdb.MockCategory = &cat
	fcat := si.GetCategory(4)
	if fcat.Description != cat.Description {
		t.Fail()
	}

	var catlst []sdbi.Category
	catlst = append(catlst, cat)
	sdb.MockCategoryList = &catlst
	fcatlst := si.GetCategoryList(4)
	if len(*fcatlst) != 1 {
		t.Fail()
	}

	fcatlst2 := si.GetSubCategoryList(3)
	if len(*fcatlst2) != 1 {
		t.Fail()
	}

	sdb.MockDeleteCategorySuccess = true
	dlcatsuc := si.DeleteCategory(4)
	if !dlcatsuc {
		t.Fail()
	}

	var sm sdbi.ShippingMethod
	sm.Cost = 12.95

	sdb.MockAddShippingMethodSuccess = true
	sdb.MockShippingMethodID = 9

	smsuc, smid := si.AddShippingMethod(&sm)
	if !smsuc || smid == 0 {
		t.Fail()
	}

	sdb.MockUpdateShippingMethodSuccess = true
	usmsuc := si.UpdateShippingMethod(&sm)
	if !usmsuc {
		t.Fail()
	}

	sdb.MockShippingMethod = &sm
	fsm := si.GetShippingMethod(3)
	if fsm.Cost != sm.Cost {
		t.Fail()
	}

	var smlst []sdbi.ShippingMethod
	smlst = append(smlst, sm)
	sdb.MockShippingMethodList = &smlst

	fsmlist := si.GetShippingMethodList(5)
	if len(*fsmlist) != 1 {
		t.Fail()
	}

	sdb.MockDeleteShippingMethodSuccess = true

	dlsmsuc := si.DeleteShippingMethod(5)
	if !dlsmsuc {
		t.Fail()
	}

	var ins sdbi.Insurance
	ins.Cost = 4.95

	sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 4

	inssuc, insid := si.AddInsurance(&ins)
	if !inssuc || insid == 0 {
		t.Fail()
	}

	sdb.MockUpdateInsuranceSuccess = true

	uinssuc := si.UpdateInsurance(&ins)
	if !uinssuc {
		t.Fail()
	}

	sdb.MockInsurance = &ins
	fins := si.GetInsurance(4)
	if fins.Cost != ins.Cost {
		t.Fail()
	}

	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)

	sdb.MockInsuranceList = &inslst
	finslst := si.GetInsuranceList(4)
	if len(*finslst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteInsuranceSuccess = true
	dlins := si.DeleteInsurance(2)
	if !dlins {
		t.Fail()
	}

	var prod sdbi.Product
	prod.Color = "red"

	sdb.MockAddProductSuccess = true
	sdb.MockProductID = 5
	prodsuc, pid := si.AddProduct(&prod)
	if !prodsuc || pid == 0 {
		t.Fail()
	}

	sdb.MockUpdateProductSuccess = true
	uprodsuc := si.UpdateProduct(&prod)
	if !uprodsuc {
		t.Fail()
	}

	sdb.MockProduct = &prod
	fprod := si.GetProductByID(2)
	if fprod.Color != prod.Color {
		t.Fail()
	}

	var prodlst []sdbi.Product
	prodlst = append(prodlst, prod)
	sdb.MockProductList = &prodlst

	fprodlst := si.GetProductList(3, 3, 3)
	if len(*fprodlst) != 1 {
		t.Fail()
	}

	fprodlst2 := si.GetProductsByCaterory(3, 3, 3)
	if len(*fprodlst2) != 1 {
		t.Fail()
	}

	fprodlst3 := si.GetProductsByName("tst", 5, 3, 3)
	if len(*fprodlst3) != 1 {
		t.Fail()
	}

	sdb.MockDeleteProductSuccess = true
	dlprod := si.DeleteProduct(3)
	if !dlprod {
		t.Fail()
	}

	var rgn sdbi.Region
	rgn.Name = "USA"

	sdb.MockAddRegionSuccess = true
	sdb.MockRegionID = 2
	rgnsuc, rgnid := si.AddRegion(&rgn)
	if !rgnsuc || rgnid == 0 {
		t.Fail()
	}

	sdb.MockUpdateRegionSuccess = true
	urgnsuc := si.UpdateRegion(&rgn)
	if !urgnsuc {
		t.Fail()
	}

	sdb.MockRegion = &rgn
	frgn := si.GetRegion(3)
	if frgn.Name != rgn.Name {
		t.Fail()
	}

	var rgnlst []sdbi.Region
	rgnlst = append(rgnlst, rgn)
	sdb.MockRegionList = &rgnlst

	frgnlst := si.GetRegionList(3)
	if len(*frgnlst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteRegionSuccess = true
	dlrgn := si.DeleteRegion(4)
	if !dlrgn {
		t.Fail()
	}

	var srgn sdbi.SubRegion
	srgn.Name = "Georgia"

	sdb.MockAddSubRegionSuccess = true
	sdb.MockSubRegionID = 4
	srgnsuc, srgnid := si.AddSubRegion(&srgn)
	if !srgnsuc || srgnid == 0 {
		t.Fail()
	}

	sdb.MockUpdateSubRegionSuccess = true
	usrgnsuc := si.UpdateSubRegion(&srgn)
	if !usrgnsuc {
		t.Fail()
	}

	sdb.MockSubRegion = &srgn
	fsrgn := si.GetSubRegion(4)
	if fsrgn.Name != srgn.Name {
		t.Fail()
	}

	var srgnlst []sdbi.SubRegion
	srgnlst = append(srgnlst, srgn)
	sdb.MockSubRegionList = &srgnlst

	fsrgnlst := si.GetSubRegionList(4)
	if len(*fsrgnlst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteSubRegionSuccess = true
	dlsrgn := si.DeleteSubRegion(5)
	if !dlsrgn {
		t.Fail()
	}

	var esr sdbi.ExcludedSubRegion
	esr.RegionID = 6

	sdb.MockAddExcludedSubRegionSuccess = true
	sdb.MockExcludedSubRegionID = 4

	esrsuc, esrid := si.AddExcludedSubRegion(&esr)
	if !esrsuc || esrid == 0 {
		t.Fail()
	}

	sdb.MockUpdateExcludedSubRegionSuccess = true
	uesr := si.UpdateExcludedSubRegion(&esr)
	if !uesr {
		t.Fail()
	}

	sdb.MockExcludedSubRegion = &esr
	fesr := si.GetExcludedSubRegion(3)
	if fesr.RegionID != esr.RegionID {
		t.Fail()
	}

	var esrlst []sdbi.ExcludedSubRegion
	esrlst = append(esrlst, esr)
	sdb.MockExcludedSubRegionList = &esrlst
	fesrlst := si.GetExcludedSubRegionList(3)

	if len(*fesrlst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteExcludedSubRegionSuccess = true
	dlesr := si.DeleteExcludedSubRegion(4)
	if !dlesr {
		t.Fail()
	}

	var isr sdbi.IncludedSubRegion
	isr.RegionID = 5

	sdb.MockAddIncludedSubRegionSuccess = true
	sdb.MockIncludedSubRegionID = 3

	isrsuc, isrid := si.AddIncludedSubRegion(&isr)
	if !isrsuc || isrid == 0 {
		t.Fail()
	}

	sdb.MockUpdateIncludedSubRegionSuccess = true
	uisr := si.UpdateIncludedSubRegion(&isr)
	if !uisr {
		t.Fail()
	}

	sdb.MockIncludedSubRegion = &isr
	fisr := si.GetIncludedSubRegion(9)
	if fisr.RegionID != isr.RegionID {
		t.Fail()
	}

	var isrlst []sdbi.IncludedSubRegion
	isrlst = append(isrlst, isr)
	sdb.MockIncludedSubRegionList = &isrlst

	fisrlst := si.GetIncludedSubRegionList(4)
	if len(*fisrlst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteIncludedSubRegionSuccess = true
	dlisr := si.DeleteIncludedSubRegion(3)
	if !dlisr {
		t.Fail()
	}

	var zip sdbi.ZoneZip
	zip.ZipCode = "30066"

	sdb.MockAddZoneZipSuccess = true
	sdb.MockZoneZipID = 4

	zipsuc, zipid := si.AddZoneZip(&zip)
	if !zipsuc || zipid == 0 {
		t.Fail()
	}

	var ziplst []sdbi.ZoneZip
	ziplst = append(ziplst, zip)
	sdb.MockZoneZipList = &ziplst

	fziplst := si.GetZoneZipListByExclusion(4)
	if len(*fziplst) != 1 {
		t.Fail()
	}

	fziplst2 := si.GetZoneZipListByInclusion(4)
	if len(*fziplst2) != 1 {
		t.Fail()
	}

	sdb.MockDeleteZoneZipSuccess = true
	dlzip := si.DeleteZoneZip(3)
	if !dlzip {
		t.Fail()
	}

	var pcat sdbi.ProductCategory
	pcat.CategoryID = 3
	pcat.ProductID = 4

	sdb.MockAddProductCategorySuccess = true
	pcatsuc := si.AddProductCategory(&pcat)
	if !pcatsuc {
		t.Fail()
	}

	sdb.MockDeleteProductCategorySuccess = true
	dlpcat := si.DeleteProductCategory(&pcat)
	if !dlpcat {
		t.Fail()
	}

	var odr sdbi.Order
	odr.BillingAddress = "133 test st. Dallas, TX"

	sdb.MockAddOrderSuccess = true
	sdb.MockOrderID = 5
	odrsuc, oid := si.AddOrder(&odr)
	if !odrsuc || oid == 0 {
		t.Fail()
	}

	sdb.MockUpdateOrderSuccess = true
	uodrsuc := si.UpdateOrder(&odr)
	if !uodrsuc {
		t.Fail()
	}

	sdb.MockOrder = &odr
	fodr := si.GetOrder(3)
	if fodr.BillingAddress != odr.BillingAddress {
		t.Fail()
	}

	var odrlst []sdbi.Order
	odrlst = append(odrlst, odr)
	sdb.MockOrderList = &odrlst
	fodrlst := si.GetOrderList(4, 6)
	if len(*fodrlst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteOrderSuccess = true
	dlodr := si.DeleteOrder(4)
	if !dlodr {
		t.Fail()
	}

	var oi sdbi.OrderItem
	oi.BackOrdered = false

	sdb.MockAddOrderItemSuccess = true
	sdb.MockOrderItemID = 3
	oisuc, oiid := si.AddOrderItem(&oi)
	if !oisuc || oiid == 0 {
		t.Fail()
	}

	sdb.MockUpdateOrderItemSuccess = true
	uoi := si.UpdateOrderItem(&oi)
	if !uoi {
		t.Fail()
	}

	sdb.MockOrderItem = &oi
	foi := si.GetOrderItem(2)
	if foi.BackOrdered != oi.BackOrdered {
		t.Fail()
	}

	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sdb.MockOrderItemList = &oilst

	foilst := si.GetOrderItemList(5)
	if len(*foilst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteOrderItemSuccess = true
	dloi := si.DeleteOrderItem(3)
	if !dloi {
		t.Fail()
	}

	var oc sdbi.OrderComment
	oc.Comment = "this is way cooler"

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 4

	ocsuc, ocid := si.AddOrderComments(&oc)
	if !ocsuc || ocid == 0 {
		t.Fail()
	}

	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sdb.MockOrderCommentList = &oclst
	foclst := si.GetOrderCommentList(1)
	if len(*foclst) != 1 {
		t.Fail()
	}

	var ot sdbi.OrderTransaction
	ot.Amount = 15.23

	sdb.MockAddOrderTransactionSuccess = true
	sdb.MockOrderTransactionID = 3

	otsuc, otid := si.AddOrderTransaction(&ot)
	if !otsuc || otid == 0 {
		t.Fail()
	}

	var otlst []sdbi.OrderTransaction
	otlst = append(otlst, ot)
	sdb.MockOrderTransactionList = &otlst

	fotlst := si.GetOrderTransactionList(5)
	if len(*fotlst) != 1 {
		t.Fail()
	}

	var shp sdbi.Shipment
	shp.Boxes = 1

	sdb.MockAddShipmentSuccess = true
	sdb.MockShipmentID = 3

	shpsuc, shpid := si.AddShipment(&shp)
	if !shpsuc || shpid == 0 {
		t.Fail()
	}

	sdb.MockUpdateShipmentSuccess = true
	ushp := si.UpdateShipment(&shp)
	if !ushp {
		t.Fail()
	}

	sdb.MockShipment = &shp
	fshp := si.GetShipment(4)
	if fshp.Boxes != shp.Boxes {
		t.Fail()
	}

	var shplst []sdbi.Shipment
	shplst = append(shplst, shp)
	sdb.MockShipmentList = &shplst
	fshplst := si.GetShipmentList(4)
	if len(*fshplst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteShipmentSuccess = true
	dlshp := si.DeleteShipment(4)
	if !dlshp {
		t.Fail()
	}

	var sbox sdbi.ShipmentBox
	sbox.BoxNumber = 1

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 5
	sboxsuc, sboxid := si.AddShipmentBox(&sbox)
	if !sboxsuc || sboxid == 0 {
		t.Fail()
	}

	sdb.MockUpdateShipmentBoxSuccess = true
	usbox := si.UpdateShipmentBox(&sbox)
	if !usbox {
		t.Fail()
	}

	sdb.MockShipmentBox = &sbox
	fsbox := si.GetShipmentBox(5)
	if fsbox.BoxNumber != sbox.BoxNumber {
		t.Fail()
	}

	var sboxlst []sdbi.ShipmentBox
	sboxlst = append(sboxlst, sbox)
	sdb.MockShipmentBoxList = &sboxlst
	fsboxlst := si.GetShipmentBoxList(4)
	if len(*fsboxlst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteShipmentBoxSuccess = true
	dlsbox := si.DeleteShipmentBox(4)
	if !dlsbox {
		t.Fail()
	}

	var shpitm sdbi.ShipmentItem
	shpitm.OrderItemID = 3

	sdb.MockAddShipmentItemSuccess = true
	sdb.MockShipmentItemID = 4
	sitmsuc, sitmid := si.AddShipmentItem(&shpitm)
	if !sitmsuc || sitmid == 0 {
		t.Fail()
	}

	sdb.MockUpdateShipmentItemSuccess = true
	usitm := si.UpdateShipmentItem(&shpitm)
	if !usitm {
		t.Fail()
	}

	sdb.MockShipmentItem = &shpitm
	fshpitm := si.GetShipmentItem(6)
	if fshpitm.OrderItemID != shpitm.OrderItemID {
		t.Fail()
	}

	var shpitlst []sdbi.ShipmentItem
	shpitlst = append(shpitlst, shpitm)
	sdb.MockShipmentItemList = &shpitlst

	fshpitlst := si.GetShipmentItemList(5)
	if len(*fshpitlst) != 1 {
		t.Fail()
	}

	fshpitlst2 := si.GetShipmentItemListByBox(1, 5)
	if len(*fshpitlst2) != 1 {
		t.Fail()
	}

	sdb.MockDeleteShipmentItemSuccess = true
	dlshpitm := si.DeleteShipmentItem(4)
	if !dlshpitm {
		t.Fail()
	}

	var pi sdbi.Plugins
	pi.ActivateURL = "test"

	sdb.MockAddPluginSuccess = true
	sdb.MockPluginID = 1
	pisuc, piid := si.AddPlugin(&pi)
	if !pisuc || piid == 0 {
		t.Fail()
	}

	sdb.MockUpdatePluginSuccess = true
	upi := si.UpdatePlugin(&pi)
	if !upi {
		t.Fail()
	}

	sdb.MockPlugin = &pi
	fpi := si.GetPlugin(3)
	if fpi.ActivateURL != pi.ActivateURL {
		t.Fail()
	}

	var fpilst []sdbi.Plugins
	fpilst = append(fpilst, pi)
	sdb.MockPluginList = &fpilst

	pilst := si.GetPluginList(6, 8)
	if len(*pilst) != 1 {
		t.Fail()
	}

	sdb.MockDeletePluginSuccess = true
	dlpi := si.DeletePlugin(4)
	if !dlpi {
		t.Fail()
	}

	var spi sdbi.StorePlugins
	spi.APIKey = "123"

	sdb.MockAddStorePluginSuccess = true
	sdb.MockStorePluginID = 2

	spisuc, spiid := si.AddStorePlugin(&spi)
	if !spisuc || spiid == 0 {
		t.Fail()
	}

	sdb.MockUpdateStorePluginSuccess = true
	uspi := si.UpdateStorePlugin(&spi)
	if !uspi {
		t.Fail()
	}

	sdb.MockStorePlugin = &spi
	fspi := si.GetStorePlugin(5)
	if fspi.APIKey != spi.APIKey {
		t.Fail()
	}

	var splst []sdbi.StorePlugins
	splst = append(splst, spi)

	sdb.MockStorePluginList = &splst

	fsplst := si.GetStorePluginList(4)
	if len(*fsplst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteStorePluginSuccess = true

	dspisuc := si.DeleteStorePlugin(3)
	if !dspisuc {
		t.Fail()
	}

	var pgw sdbi.PaymentGateway
	pgw.CheckoutURL = "/test"

	sdb.MockAddPaymentGatewaySuccess = true
	sdb.MockPaymentGatewayID = 3

	pgwsuc, pgwid := si.AddPaymentGateway(&pgw)
	if !pgwsuc || pgwid == 0 {
		t.Fail()
	}

	sdb.MockUpdatePaymentGatewaySuccess = true
	upgw := si.UpdatePaymentGateway(&pgw)
	if !upgw {
		t.Fail()
	}

	sdb.MockPaymentGateway = &pgw
	fpgw := si.GetPaymentGateway(4)
	if fpgw.CheckoutURL != pgw.CheckoutURL {
		t.Fail()
	}
	var pgwlst []sdbi.PaymentGateway
	pgwlst = append(pgwlst, pgw)

	sdb.MockPaymentGatewayList = &pgwlst
	fpgwlst := si.GetPaymentGateways(5)
	if len(*fpgwlst) != 1 {
		t.Fail()
	}

	sdb.MockDeletePaymentGatewaySuccess = true
	dpgw := si.DeletePaymentGateway(4)
	if !dpgw {
		t.Fail()
	}

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"

	sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 3
	scsuc, scid := si.AddShippingCarrier(&sc)
	if !scsuc || scid == 0 {
		t.Fail()
	}

	sdb.MockUpdateShippingCarrierSuccess = true
	usc := si.UpdateShippingCarrier(&sc)
	if !usc {
		t.Fail()
	}

	var sclst []sdbi.ShippingCarrier

	sdb.MockShippingCarrier = &sc
	fsc := si.GetShippingCarrier(5)
	if fsc.Carrier != sc.Carrier {
		t.Fail()
	}
	sclst = append(sclst, sc)
	sdb.MockShippingCarrierList = &sclst
	fsclst := si.GetShippingCarrierList(4)
	if len(*fsclst) != 1 {
		t.Fail()
	}

	sdb.MockDeleteShippingCarrierSuccess = true
	dlsc := si.DeleteShippingCarrier(4)
	if !dlsc {
		t.Fail()
	}

	var lds sdbi.LocalDataStore
	lds.DataStoreName = "content"

	sdb.MockAddLocalDataStoreSuccess = true
	ldssuc := si.AddLocalDatastore(&lds)
	if !ldssuc {
		t.Fail()
	}

	sdb.MockUpdateLocalDataStoreSuccess = true
	ulds := si.UpdateLocalDatastore(&lds)
	if !ulds {
		t.Fail()
	}

	sdb.MockLocalDataStore = &lds
	flds := si.GetLocalDatastore(4, "test")
	if flds.DataStoreName != lds.DataStoreName {
		t.Fail()
	}

	var inst sdbi.Instances
	inst.DataStoreName = "content"

	sdb.MockAddInstancesSuccess = true
	instsuc := si.AddInstance(&inst)
	if !instsuc {
		t.Fail()
	}

	sdb.MockUpdateInstancesSuccess = true
	uinst := si.UpdateInstance(&inst)
	if !uinst {
		t.Fail()
	}

	sdb.MockInstances = &inst
	finst := si.GetInstance("test", "store", 3)
	if finst.DataStoreName != inst.DataStoreName {
		t.Fail()
	}

	var instlst []sdbi.Instances
	instlst = append(instlst, inst)
	sdb.MockInstancesList = &instlst
	finstl := si.GetInstanceList("store", 3)
	if (*finstl)[0].DataStoreName != inst.DataStoreName {
		t.Fail()
	}

	var lc sdbi.DataStoreWriteLock
	lc.DataStoreName = "content"

	sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 2
	lcsuc := si.AddDataStoreWriteLock(&lc)
	if !lcsuc {
		t.Fail()
	}

	sdb.MockUpdateDataStoreWriteLockSuccess = true
	ulc := si.UpdateDataStoreWriteLock(&lc)
	if !ulc {
		t.Fail()
	}

	sdb.MockDataStoreWriteLock = &lc
	flc := si.GetDataStoreWriteLock("test", 1)
	if flc.DataStoreName != lc.DataStoreName {
		t.Fail()
	}

}
