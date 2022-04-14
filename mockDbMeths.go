package six910mysql

import mdb "github.com/Ulbora/six910-database-interface"

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//AddSecurity AddSecurity
func (d *MockSix910Mysql) AddSecurity(s *mdb.Security) (bool, int64) {
	return d.MockAddSecuritySuccess, d.MockSecurityID
}

//UpdateSecurity UpdateSecurity
func (d *MockSix910Mysql) UpdateSecurity(s *mdb.Security) bool {
	return d.MockUpdateSecuritySuccess
}

//GetSecurity GetSecurity
func (d *MockSix910Mysql) GetSecurity() *mdb.Security {
	return d.MockSecurity
}

//DeleteSecurity DeleteSecurity
func (d *MockSix910Mysql) DeleteSecurity() bool {
	return d.MockDeleteSecuritySuccess
}

//stores

//AddStore AddStore
func (d *MockSix910Mysql) AddStore(s *mdb.Store) (bool, int64) {
	return d.MockAddStoreSuccess, d.MockStoreID
}

//UpdateStore UpdateStore
func (d *MockSix910Mysql) UpdateStore(s *mdb.Store) bool {
	return d.MockUpdateStoreSuccess
}

//GetStore GetStore
func (d *MockSix910Mysql) GetStore(sname string) *mdb.Store {
	return d.MockStore
}

//GetLocalStore GetLocalStore
func (d *MockSix910Mysql) GetLocalStore() *mdb.Store {
	return d.MockStore
}

//GetStoreID GetStoreID
func (d *MockSix910Mysql) GetStoreID(id int64) *mdb.Store {
	return d.MockStore
}

//GetStoreByLocal GetStoreByLocal
func (d *MockSix910Mysql) GetStoreByLocal(localDomain string) *mdb.Store {
	return d.MockStore
}

//GetStoreCount GetStoreCount
func (d *MockSix910Mysql) GetStoreCount() int64 {
	return d.MockStoreCount
}

//DeleteStore DeleteStore
func (d *MockSix910Mysql) DeleteStore(id int64) bool {
	return d.MockDeleteStoreSuccess
}

//customer

//AddCustomer AddCustomer
func (d *MockSix910Mysql) AddCustomer(c *mdb.Customer) (bool, int64) {
	return d.MockAddCustomerSuccess, d.MockCustomerID
}

//UpdateCustomer UpdateCustomer
func (d *MockSix910Mysql) UpdateCustomer(c *mdb.Customer) bool {
	return d.MockUpdateCustomerSuccess
}

//GetCustomer GetCustomer
func (d *MockSix910Mysql) GetCustomer(email string, storeID int64) *mdb.Customer {
	return d.MockCustomer
}

//GetCustomerList GetCustomerList
func (d *MockSix910Mysql) GetCustomerList(storeID int64, start int64, end int64) *[]mdb.Customer {
	return d.MockCustomerList
}

//GetCustomerUsers GetCustomerUsers
func (d *MockSix910Mysql) GetCustomerUsers(cid int64, storeID int64) *[]mdb.LocalAccount {
	return d.MockLocalAccountList
}

//GetCustomerID GetCustomerID
func (d *MockSix910Mysql) GetCustomerID(id int64) *mdb.Customer {
	return d.MockCustomer
}

//DeleteCustomer DeleteCustomer
func (d *MockSix910Mysql) DeleteCustomer(id int64) bool {
	return d.MockDeleteCustomerSuccess
}

//Local Accounts when oauth not used

//AddLocalAccount AddLocalAccount
func (d *MockSix910Mysql) AddLocalAccount(a *mdb.LocalAccount) bool {
	return d.MockAddLocalAccountSuccess
}

//UpdateLocalAccount UpdateLocalAccount
func (d *MockSix910Mysql) UpdateLocalAccount(a *mdb.LocalAccount) bool {
	return d.MockUpdateLocalAccountSuccess
}

//GetLocalAccount GetLocalAccount
func (d *MockSix910Mysql) GetLocalAccount(username string, storeID int64) *mdb.LocalAccount {
	return d.MockLocalAccount
}

//GetLocalAccountList GetLocalAccountList
func (d *MockSix910Mysql) GetLocalAccountList(store int64) *[]mdb.LocalAccount {
	return d.MockLocalAccountList
}

//DeleteLocalAccount DeleteLocalAccount
func (d *MockSix910Mysql) DeleteLocalAccount(username string, storeID int64) bool {
	return d.MockDeleteLocalAccountSuccess
}

//distributors

//AddDistributor AddDistributor
func (d *MockSix910Mysql) AddDistributor(ds *mdb.Distributor) (bool, int64) {
	return d.MockAddDistributorSuccess, d.MockDistributorID
}

//UpdateDistributor UpdateDistributor
func (d *MockSix910Mysql) UpdateDistributor(ds *mdb.Distributor) bool {
	return d.MockUpdateDistributorSuccess
}

//GetDistributor GetDistributor
func (d *MockSix910Mysql) GetDistributor(id int64) *mdb.Distributor {
	return d.MockDistributor
}

//GetDistributorList GetDistributorList
func (d *MockSix910Mysql) GetDistributorList(store int64) *[]mdb.Distributor {
	return d.MockDistributorList
}

//DeleteDistributor DeleteDistributor
func (d *MockSix910Mysql) DeleteDistributor(id int64) bool {
	return d.MockDeleteDistributorSuccess
}

//Cart

//AddCart AddCart
func (d *MockSix910Mysql) AddCart(c *mdb.Cart) (bool, int64) {
	return d.MockAddCartSuccess, d.MockCartID
}

//UpdateCart UpdateCart
func (d *MockSix910Mysql) UpdateCart(c *mdb.Cart) bool {
	return d.MockUpdateCartSuccess
}

//GetCart GetCart
func (d *MockSix910Mysql) GetCart(cid int64) *mdb.Cart {
	return d.MockCart
}

//DeleteCart DeleteCart
func (d *MockSix910Mysql) DeleteCart(id int64) bool {
	return d.MockDeleteCartSuccess
}

//cart item

//AddCartItem AddCartItem
func (d *MockSix910Mysql) AddCartItem(ci *mdb.CartItem) (bool, int64) {
	return d.MockAddCartItemSuccess, d.MockCartItemID
}

//UpdateCartItem UpdateCartItem
func (d *MockSix910Mysql) UpdateCartItem(ci *mdb.CartItem) bool {
	return d.MockUpdateCartItemSuccess
}

//GetCarItem GetCarItem
func (d *MockSix910Mysql) GetCarItem(cartID int64, prodID int64) *mdb.CartItem {
	return d.MockCartItem
}

//GetCartItemList GetCartItemList
func (d *MockSix910Mysql) GetCartItemList(cartID int64) *[]mdb.CartItem {
	return d.MockCartItemList
}

//DeleteCartItem DeleteCartItem
func (d *MockSix910Mysql) DeleteCartItem(id int64) bool {
	return d.MockDeleteCartItemSuccess
}

//address

//AddAddress AddAddress
func (d *MockSix910Mysql) AddAddress(a *mdb.Address) (bool, int64) {
	return d.MockAddAddressSuccess, d.MockAddressID
}

//UpdateAddress UpdateAddress
func (d *MockSix910Mysql) UpdateAddress(a *mdb.Address) bool {
	return d.MockUpdateAddressSuccess
}

//GetAddress GetAddress
func (d *MockSix910Mysql) GetAddress(id int64) *mdb.Address {
	return d.MockAddress
}

//GetAddressList GetAddressList
func (d *MockSix910Mysql) GetAddressList(cid int64) *[]mdb.Address {
	return d.MockAddressList
}

//DeleteAddress DeleteAddress
func (d *MockSix910Mysql) DeleteAddress(id int64) bool {
	return d.MockDeleteAddressSuccess
}

//category

//AddCategory AddCategory
func (d *MockSix910Mysql) AddCategory(c *mdb.Category) (bool, int64) {
	return d.MockAddCategorySuccess, d.MockCategoryID
}

//UpdateCategory UpdateCategory
func (d *MockSix910Mysql) UpdateCategory(c *mdb.Category) bool {
	return d.MockUpdateCategorySuccess
}

//GetCategory GetCategory
func (d *MockSix910Mysql) GetCategory(id int64) *mdb.Category {
	return d.MockCategory
}

//GetHierarchicalCategoryList GetHierarchicalCategoryList
func (d *MockSix910Mysql) GetHierarchicalCategoryList(storeID int64) *[]mdb.Category {
	return d.MockCategoryList
}

//GetCategoryList GetCategoryList
func (d *MockSix910Mysql) GetCategoryList(storeID int64) *[]mdb.Category {
	return d.MockCategoryList
}

//GetSubCategoryList GetSubCategoryList
func (d *MockSix910Mysql) GetSubCategoryList(catID int64) *[]mdb.Category {
	return d.MockCategoryList
}

//DeleteCategory DeleteCategory
func (d *MockSix910Mysql) DeleteCategory(id int64) bool {
	return d.MockDeleteCategorySuccess
}

//shipping method

//AddShippingMethod AddShippingMethod
func (d *MockSix910Mysql) AddShippingMethod(s *mdb.ShippingMethod) (bool, int64) {
	return d.MockAddShippingMethodSuccess, d.MockShippingMethodID
}

//UpdateShippingMethod UpdateShippingMethod
func (d *MockSix910Mysql) UpdateShippingMethod(s *mdb.ShippingMethod) bool {
	return d.MockUpdateShippingMethodSuccess
}

//GetShippingMethod GetShippingMethod
func (d *MockSix910Mysql) GetShippingMethod(id int64) *mdb.ShippingMethod {
	return d.MockShippingMethod
}

//GetShippingMethodList GetShippingMethodList
func (d *MockSix910Mysql) GetShippingMethodList(storeID int64) *[]mdb.ShippingMethod {
	return d.MockShippingMethodList
}

//DeleteShippingMethod DeleteShippingMethod
func (d *MockSix910Mysql) DeleteShippingMethod(id int64) bool {
	return d.MockDeleteShippingMethodSuccess
}

//shipping insurance

//AddInsurance AddInsurance
func (d *MockSix910Mysql) AddInsurance(i *mdb.Insurance) (bool, int64) {
	return d.MockAddInsuranceSuccess, d.MockInsuranceID
}

//UpdateInsurance UpdateInsurance
func (d *MockSix910Mysql) UpdateInsurance(i *mdb.Insurance) bool {
	return d.MockUpdateInsuranceSuccess
}

//GetInsurance GetInsurance
func (d *MockSix910Mysql) GetInsurance(id int64) *mdb.Insurance {
	return d.MockInsurance
}

//GetInsuranceList GetInsuranceList
func (d *MockSix910Mysql) GetInsuranceList(storeID int64) *[]mdb.Insurance {
	return d.MockInsuranceList
}

//DeleteInsurance DeleteInsurance
func (d *MockSix910Mysql) DeleteInsurance(id int64) bool {
	return d.MockDeleteInsuranceSuccess
}

//product

//AddProduct AddProduct
func (d *MockSix910Mysql) AddProduct(p *mdb.Product) (bool, int64) {
	return d.MockAddProductSuccess, d.MockProductID
}

//UpdateProduct UpdateProduct
func (d *MockSix910Mysql) UpdateProduct(p *mdb.Product) bool {
	return d.MockUpdateProductSuccess
}

//UpdateProductQuantity UpdateProductQuantity
func (d *MockSix910Mysql) UpdateProductQuantity(p *mdb.Product) bool {
	return d.MockUpdateProductQuantitySuccess
}

//GetProductByID GetProductByID
func (d *MockSix910Mysql) GetProductByID(id int64) *mdb.Product {
	return d.MockProduct
}

//GetProductBySku GetProductBySku
func (d *MockSix910Mysql) GetProductBySku(sku string, distributorID int64, storeID int64) *mdb.Product {
	return d.MockProduct
}

//GetProductsByPromoted GetProductsByPromoted
func (d *MockSix910Mysql) GetProductsByPromoted(storeID int64, start int64, end int64) *[]mdb.Product {
	return d.MockProductList
}

//GetProductsByName GetProductsByName
func (d *MockSix910Mysql) GetProductsByName(name string, storeID int64, start int64, end int64) *[]mdb.Product {
	return d.MockProductList
}

//GetProductsByCaterory GetProductsByCaterory
func (d *MockSix910Mysql) GetProductsByCaterory(catID int64, start int64, end int64) *[]mdb.Product {
	return d.MockProductList
}

//GetProductList GetProductList
func (d *MockSix910Mysql) GetProductList(storeID int64, start int64, end int64) *[]mdb.Product {
	return d.MockProductList
}

//GetProductIDList GetProductIDList
func (d *MockSix910Mysql) GetProductIDList(storeID int64) *[]int64 {
	return d.MockProductIDList
}

//GetProductIDListByCategories GetProductIDListByCategories
func (d *MockSix910Mysql) GetProductIDListByCategories(storeID int64, catList *[]int64) *[]int64 {
	return d.MockProductIDList
}

//GetProductSubSkuList GetProductSubSkuList
func (d *MockSix910Mysql) GetProductSubSkuList(parentProdID int64) *[]mdb.Product {
	return d.MockProductSubSkuList
}

//DeleteProduct DeleteProduct
func (d *MockSix910Mysql) DeleteProduct(id int64) bool {
	return d.MockDeleteProductSuccess
}

//Geographic Regions

//AddRegion AddRegion
func (d *MockSix910Mysql) AddRegion(r *mdb.Region) (bool, int64) {
	return d.MockAddRegionSuccess, d.MockRegionID
}

//UpdateRegion UpdateRegion
func (d *MockSix910Mysql) UpdateRegion(r *mdb.Region) bool {
	return d.MockUpdateRegionSuccess
}

//GetRegion GetRegion
func (d *MockSix910Mysql) GetRegion(id int64) *mdb.Region {
	return d.MockRegion
}

//GetRegionList GetRegionList
func (d *MockSix910Mysql) GetRegionList(storeID int64) *[]mdb.Region {
	return d.MockRegionList
}

//DeleteRegion DeleteRegion
func (d *MockSix910Mysql) DeleteRegion(id int64) bool {
	return d.MockDeleteRegionSuccess
}

//Geographic Sub Regions

//AddSubRegion AddSubRegion
func (d *MockSix910Mysql) AddSubRegion(s *mdb.SubRegion) (bool, int64) {
	return d.MockAddSubRegionSuccess, d.MockSubRegionID
}

//UpdateSubRegion UpdateSubRegion
func (d *MockSix910Mysql) UpdateSubRegion(s *mdb.SubRegion) bool {
	return d.MockUpdateSubRegionSuccess
}

//GetSubRegion GetSubRegion
func (d *MockSix910Mysql) GetSubRegion(id int64) *mdb.SubRegion {
	return d.MockSubRegion
}

//GetSubRegionList GetSubRegionList
func (d *MockSix910Mysql) GetSubRegionList(regionID int64) *[]mdb.SubRegion {
	return d.MockSubRegionList
}

//DeleteSubRegion DeleteSubRegion
func (d *MockSix910Mysql) DeleteSubRegion(id int64) bool {
	return d.MockDeleteSubRegionSuccess
}

//excluded sub regions

//AddExcludedSubRegion AddExcludedSubRegion
func (d *MockSix910Mysql) AddExcludedSubRegion(e *mdb.ExcludedSubRegion) (bool, int64) {
	return d.MockAddExcludedSubRegionSuccess, d.MockExcludedSubRegionID
}

//UpdateExcludedSubRegion UpdateExcludedSubRegion
func (d *MockSix910Mysql) UpdateExcludedSubRegion(e *mdb.ExcludedSubRegion) bool {
	return d.MockUpdateExcludedSubRegionSuccess
}

//GetExcludedSubRegion GetExcludedSubRegion
func (d *MockSix910Mysql) GetExcludedSubRegion(id int64) *mdb.ExcludedSubRegion {
	return d.MockExcludedSubRegion
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (d *MockSix910Mysql) GetExcludedSubRegionList(regionID int64) *[]mdb.ExcludedSubRegion {
	return d.MockExcludedSubRegionList
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (d *MockSix910Mysql) DeleteExcludedSubRegion(id int64) bool {
	return d.MockDeleteExcludedSubRegionSuccess
}

//included sub regions

//AddIncludedSubRegion AddIncludedSubRegion
func (d *MockSix910Mysql) AddIncludedSubRegion(e *mdb.IncludedSubRegion) (bool, int64) {
	return d.MockAddIncludedSubRegionSuccess, d.MockIncludedSubRegionID
}

//UpdateIncludedSubRegion UpdateIncludedSubRegion
func (d *MockSix910Mysql) UpdateIncludedSubRegion(e *mdb.IncludedSubRegion) bool {
	return d.MockUpdateIncludedSubRegionSuccess
}

//GetIncludedSubRegion GetIncludedSubRegion
func (d *MockSix910Mysql) GetIncludedSubRegion(id int64) *mdb.IncludedSubRegion {
	return d.MockIncludedSubRegion
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (d *MockSix910Mysql) GetIncludedSubRegionList(regionID int64) *[]mdb.IncludedSubRegion {
	return d.MockIncludedSubRegionList
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (d *MockSix910Mysql) DeleteIncludedSubRegion(id int64) bool {
	return d.MockDeleteIncludedSubRegionSuccess
}

//limit exclusions and inclusions to a zip code

//AddZoneZip AddZoneZip
func (d *MockSix910Mysql) AddZoneZip(z *mdb.ZoneZip) (bool, int64) {
	return d.MockAddZoneZipSuccess, d.MockZoneZipID
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (d *MockSix910Mysql) GetZoneZipListByExclusion(exID int64) *[]mdb.ZoneZip {
	return d.MockZoneZipList
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (d *MockSix910Mysql) GetZoneZipListByInclusion(incID int64) *[]mdb.ZoneZip {
	return d.MockZoneZipList
}

//DeleteZoneZip DeleteZoneZip
func (d *MockSix910Mysql) DeleteZoneZip(id int64) bool {
	return d.MockDeleteZoneZipSuccess
}

//product category

//AddProductCategory AddProductCategory
func (d *MockSix910Mysql) AddProductCategory(pc *mdb.ProductCategory) bool {
	return d.MockAddProductCategorySuccess
}

//GetProductCategoryList GetProductCategoryList
func (d *MockSix910Mysql) GetProductCategoryList(productID int64) *[]int64 {
	return d.MockCategoryIDList
}

//DeleteProductCategory DeleteProductCategory
func (d *MockSix910Mysql) DeleteProductCategory(pc *mdb.ProductCategory) bool {
	return d.MockDeleteProductCategorySuccess
}

//Orders

//AddOrder AddOrder
func (d *MockSix910Mysql) AddOrder(o *mdb.Order) (bool, int64) {
	return d.MockAddOrderSuccess, d.MockOrderID
}

//UpdateOrder UpdateOrder
func (d *MockSix910Mysql) UpdateOrder(o *mdb.Order) bool {
	return d.MockUpdateOrderSuccess
}

//GetOrder GetOrder
func (d *MockSix910Mysql) GetOrder(id int64) *mdb.Order {
	return d.MockOrder
}

//GetOrderList GetOrderList
func (d *MockSix910Mysql) GetOrderList(cid int64, storeID int64) *[]mdb.Order {
	return d.MockOrderList
}

//GetStoreOrderList GetStoreOrderList
func (d *MockSix910Mysql) GetStoreOrderList(storeID int64) *[]mdb.Order {
	return d.MockOrderList
}

//GetStoreOrderListByStatus GetStoreOrderListByStatus
func (d *MockSix910Mysql) GetStoreOrderListByStatus(status string, storeID int64) *[]mdb.Order {
	return d.MockOrderList
}

//GetOrderCountData GetOrderCountData
func (d *MockSix910Mysql) GetOrderCountData(storeID int64) *[]mdb.OrderCountData {
	return d.MockOrderCountData
}

//GetOrderSalesData GetOrderSalesData
func (d *MockSix910Mysql) GetOrderSalesData(storeID int64) *[]mdb.OrderSalesData {
	return d.MockOrderSalesData
}

//DeleteOrder DeleteOrder
func (d *MockSix910Mysql) DeleteOrder(id int64) bool {
	return d.MockDeleteOrderSuccess
}

//Order Items

//AddOrderItem AddOrderItem
func (d *MockSix910Mysql) AddOrderItem(i *mdb.OrderItem) (bool, int64) {
	return d.MockAddOrderItemSuccess, d.MockOrderItemID
}

//UpdateOrderItem UpdateOrderItem
func (d *MockSix910Mysql) UpdateOrderItem(i *mdb.OrderItem) bool {
	return d.MockUpdateOrderItemSuccess
}

//GetOrderItem GetOrderItem
func (d *MockSix910Mysql) GetOrderItem(id int64) *mdb.OrderItem {
	return d.MockOrderItem
}

//GetOrderItemList GetOrderItemList
func (d *MockSix910Mysql) GetOrderItemList(orderID int64) *[]mdb.OrderItem {
	return d.MockOrderItemList
}

//DeleteOrderItem DeleteOrderItem
func (d *MockSix910Mysql) DeleteOrderItem(id int64) bool {
	return d.MockDeleteOrderItemSuccess
}

//Order Comments

//AddOrderComments AddOrderComments
func (d *MockSix910Mysql) AddOrderComments(c *mdb.OrderComment) (bool, int64) {
	return d.MockAddOrderCommentSuccess, d.MockOrderCommentID
}

//GetOrderCommentList GetOrderCommentList
func (d *MockSix910Mysql) GetOrderCommentList(orderID int64) *[]mdb.OrderComment {
	return d.MockOrderCommentList
}

//Order Payment Transactions

//AddOrderTransaction AddOrderTransaction
func (d *MockSix910Mysql) AddOrderTransaction(t *mdb.OrderTransaction) (bool, int64) {
	return d.MockAddOrderTransactionSuccess, d.MockOrderTransactionID
}

//GetOrderTransactionList GetOrderTransactionList
func (d *MockSix910Mysql) GetOrderTransactionList(orderID int64) *[]mdb.OrderTransaction {
	return d.MockOrderTransactionList
}

//shipment

//AddShipment AddShipment
func (d *MockSix910Mysql) AddShipment(s *mdb.Shipment) (bool, int64) {
	return d.MockAddShipmentSuccess, d.MockShipmentID
}

//UpdateShipment UpdateShipment
func (d *MockSix910Mysql) UpdateShipment(s *mdb.Shipment) bool {
	return d.MockUpdateShipmentSuccess
}

//GetShipment GetShipment
func (d *MockSix910Mysql) GetShipment(id int64) *mdb.Shipment {
	return d.MockShipment
}

//GetShipmentList GetShipmentList
func (d *MockSix910Mysql) GetShipmentList(orderID int64) *[]mdb.Shipment {
	return d.MockShipmentList
}

//DeleteShipment DeleteShipment
func (d *MockSix910Mysql) DeleteShipment(id int64) bool {
	return d.MockDeleteShipmentSuccess
}

//shipment boxes

//AddShipmentBox AddShipmentBox
func (d *MockSix910Mysql) AddShipmentBox(sb *mdb.ShipmentBox) (bool, int64) {
	return d.MockAddShipmentBoxSuccess, d.MockShipmentBoxID
}

//UpdateShipmentBox UpdateShipmentBox
func (d *MockSix910Mysql) UpdateShipmentBox(sb *mdb.ShipmentBox) bool {
	return d.MockUpdateShipmentBoxSuccess
}

//GetShipmentBox GetShipmentBox
func (d *MockSix910Mysql) GetShipmentBox(id int64) *mdb.ShipmentBox {
	return d.MockShipmentBox
}

//GetShipmentBoxList GetShipmentBoxList
func (d *MockSix910Mysql) GetShipmentBoxList(shipmentID int64) *[]mdb.ShipmentBox {
	return d.MockShipmentBoxList
}

//DeleteShipmentBox DeleteShipmentBox
func (d *MockSix910Mysql) DeleteShipmentBox(id int64) bool {
	return d.MockDeleteShipmentBoxSuccess
}

//Shipment Items in box

//AddShipmentItem AddShipmentItem
func (d *MockSix910Mysql) AddShipmentItem(si *mdb.ShipmentItem) (bool, int64) {
	return d.MockAddShipmentItemSuccess, d.MockShipmentItemID
}

//UpdateShipmentItem UpdateShipmentItem
func (d *MockSix910Mysql) UpdateShipmentItem(si *mdb.ShipmentItem) bool {
	return d.MockUpdateShipmentItemSuccess
}

//GetShipmentItem GetShipmentItem
func (d *MockSix910Mysql) GetShipmentItem(id int64) *mdb.ShipmentItem {
	return d.MockShipmentItem
}

//GetShipmentItemList GetShipmentItemList
func (d *MockSix910Mysql) GetShipmentItemList(shipmentID int64) *[]mdb.ShipmentItem {
	return d.MockShipmentItemList
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (d *MockSix910Mysql) GetShipmentItemListByBox(boxNumber int64, shipmentID int64) *[]mdb.ShipmentItem {
	return d.MockShipmentItemList
}

//DeleteShipmentItem DeleteShipmentItem
func (d *MockSix910Mysql) DeleteShipmentItem(id int64) bool {
	return d.MockDeleteShipmentItemSuccess
}

//Global Plugins

//AddPlugin AddPlugin
func (d *MockSix910Mysql) AddPlugin(p *mdb.Plugins) (bool, int64) {
	return d.MockAddPluginSuccess, d.MockPluginID
}

//UpdatePlugin UpdatePlugin
func (d *MockSix910Mysql) UpdatePlugin(p *mdb.Plugins) bool {
	return d.MockUpdatePluginSuccess
}

//GetPlugin GetPlugin
func (d *MockSix910Mysql) GetPlugin(id int64) *mdb.Plugins {
	return d.MockPlugin
}

//GetPluginList GetPluginList
func (d *MockSix910Mysql) GetPluginList(start int64, end int64) *[]mdb.Plugins {
	return d.MockPluginList
}

//DeletePlugin DeletePlugin
func (d *MockSix910Mysql) DeletePlugin(id int64) bool {
	return d.MockDeletePluginSuccess
}

//store plugins installed

//AddStorePlugin AddStorePlugin
func (d *MockSix910Mysql) AddStorePlugin(sp *mdb.StorePlugins) (bool, int64) {
	return d.MockAddStorePluginSuccess, d.MockStorePluginID
}

//UpdateStorePlugin UpdateStorePlugin
func (d *MockSix910Mysql) UpdateStorePlugin(sp *mdb.StorePlugins) bool {
	return d.MockUpdateStorePluginSuccess
}

//GetStorePlugin GetStorePlugin
func (d *MockSix910Mysql) GetStorePlugin(id int64) *mdb.StorePlugins {
	return d.MockStorePlugin
}

//GetStorePluginList GetStorePluginList
func (d *MockSix910Mysql) GetStorePluginList(storeID int64) *[]mdb.StorePlugins {
	return d.MockStorePluginList
}

//DeleteStorePlugin DeleteStorePlugin
func (d *MockSix910Mysql) DeleteStorePlugin(id int64) bool {
	return d.MockDeleteStorePluginSuccess
}

//Plugins that are payment gateways

//AddPaymentGateway AddPaymentGateway
func (d *MockSix910Mysql) AddPaymentGateway(pgw *mdb.PaymentGateway) (bool, int64) {
	return d.MockAddPaymentGatewaySuccess, d.MockPaymentGatewayID
}

//UpdatePaymentGateway UpdatePaymentGateway
func (d *MockSix910Mysql) UpdatePaymentGateway(pgw *mdb.PaymentGateway) bool {
	return d.MockUpdatePaymentGatewaySuccess
}

//GetPaymentGateway GetPaymentGateway
func (d *MockSix910Mysql) GetPaymentGateway(id int64) *mdb.PaymentGateway {
	return d.MockPaymentGateway
}

//GetPaymentGatewayByName GetPaymentGatewayByName
func (d *MockSix910Mysql) GetPaymentGatewayByName(name string, storeID int64) *mdb.PaymentGateway {
	return d.MockPaymentGateway
}

//GetPaymentGateways GetPaymentGateways
func (d *MockSix910Mysql) GetPaymentGateways(storeID int64) *[]mdb.PaymentGateway {
	return d.MockPaymentGatewayList
}

//DeletePaymentGateway DeletePaymentGateway
func (d *MockSix910Mysql) DeletePaymentGateway(id int64) bool {
	return d.MockDeletePaymentGatewaySuccess
}

//store shipment carrier like UPS and FEDex

//AddShippingCarrier AddShippingCarrier
func (d *MockSix910Mysql) AddShippingCarrier(c *mdb.ShippingCarrier) (bool, int64) {
	return d.MockAddShippingCarrierSuccess, d.MockShippingCarrierID
}

//UpdateShippingCarrier UpdateShippingCarrier
func (d *MockSix910Mysql) UpdateShippingCarrier(c *mdb.ShippingCarrier) bool {
	return d.MockUpdateShippingCarrierSuccess
}

//GetShippingCarrier GetShippingCarrier
func (d *MockSix910Mysql) GetShippingCarrier(id int64) *mdb.ShippingCarrier {
	return d.MockShippingCarrier
}

//GetShippingCarrierList GetShippingCarrierList
func (d *MockSix910Mysql) GetShippingCarrierList(storeID int64) *[]mdb.ShippingCarrier {
	return d.MockShippingCarrierList
}

//DeleteShippingCarrier DeleteShippingCarrier
func (d *MockSix910Mysql) DeleteShippingCarrier(id int64) bool {
	return d.MockDeleteShippingCarrierSuccess
}

//----UI Cluster installation: this is only called if UI is running in a cluster---
//Handle the situation where clients are running in a cluster
//This gives a way to make sure the json datastores are update on each node in the cluster

//----------------start datastore------------------------------------
//this gets called when a node start up and add only if it doesn't already exist

//AddLocalDatastore AddLocalDatastore
func (d *MockSix910Mysql) AddLocalDatastore(ds *mdb.LocalDataStore) bool {
	return d.MockAddLocalDataStoreSuccess
}

//This get get called when a change is made to a datastore from a node in the cluster
//Or after all reloads have completed and then get set to Reload = false

//UpdateLocalDatastore UpdateLocalDatastore
func (d *MockSix910Mysql) UpdateLocalDatastore(ds *mdb.LocalDataStore) bool {
	return d.MockUpdateLocalDataStoreSuccess
}

//This gets call by cluster nodes to see if there are pending reload

//GetLocalDatastore GetLocalDatastore
func (d *MockSix910Mysql) GetLocalDatastore(storeID int64, dataStoreName string) *mdb.LocalDataStore {
	return d.MockLocalDataStore
}

//---------------------start instance--------------------
// this gets called when each instance is started and added only if never added before
//The instance name is pulled from Docker or an manually entered env variable

//AddInstance AddInstance
func (d *MockSix910Mysql) AddInstance(i *mdb.Instances) bool {
	return d.MockAddInstancesSuccess
}

//This gets called after instance gets reloaded

//UpdateInstance UpdateInstance
func (d *MockSix910Mysql) UpdateInstance(i *mdb.Instances) bool {
	return d.MockUpdateInstancesSuccess
}

//Gets called before updating an instance

//GetInstance GetInstance
func (d *MockSix910Mysql) GetInstance(name string, dataStoreName string, storeID int64) *mdb.Instances {
	return d.MockInstances
}

//GetInstanceList GetInstanceList
func (d *MockSix910Mysql) GetInstanceList(dataStoreName string, storeID int64) *[]mdb.Instances {
	return d.MockInstancesList
}

//-------------------start write lock-------------
//gets called after UI makes changes to one of the datastores
//If the datastore already exists, the Update method is called from within add

//AddDataStoreWriteLock AddDataStoreWriteLock
func (d *MockSix910Mysql) AddDataStoreWriteLock(w *mdb.DataStoreWriteLock) bool {
	return d.MockAddDataStoreWriteLockSuccess
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (d *MockSix910Mysql) UpdateDataStoreWriteLock(w *mdb.DataStoreWriteLock) bool {
	return d.MockUpdateDataStoreWriteLockSuccess
}

//gets called from within the add method and by any node trying to update a datastore

//GetDataStoreWriteLock GetDataStoreWriteLock
func (d *MockSix910Mysql) GetDataStoreWriteLock(dataStore string, storeID int64) *mdb.DataStoreWriteLock {
	return d.MockDataStoreWriteLock
}

//AddTaxRate AddTaxRate
func (d *MockSix910Mysql) AddTaxRate(t *mdb.TaxRate) (bool, int64) {
	return d.MockAddTaxRateSuccess, d.MockTaxRateID
}

//UpdateTaxRate UpdateTaxRate
func (d *MockSix910Mysql) UpdateTaxRate(t *mdb.TaxRate) bool {
	return d.MockUpdateTaxRateSuccess
}

//GetTaxRate GetTaxRate
func (d *MockSix910Mysql) GetTaxRate(country string, state string, storeID int64) *[]mdb.TaxRate {
	return d.MockTaxRateList
}

//GetTaxRateList GetTaxRateList
func (d *MockSix910Mysql) GetTaxRateList(storeID int64) *[]mdb.TaxRate {
	return d.MockTaxRateList
}

//DeleteTaxRate DeleteTaxRate
func (d *MockSix910Mysql) DeleteTaxRate(id int64) bool {
	return d.MockDeleteTaxRateSuccess
}

//GetProductManufacturerListByProductName GetProductManufacturerListByProductName
func (d *MockSix910Mysql) GetProductManufacturerListByProductName(name string, storeID int64) *[]string {
	return d.MockManufacturerList
}

//GetProductByNameAndManufacturerName GetProductByNameAndManufacturerName
func (d *MockSix910Mysql) GetProductByNameAndManufacturerName(manf string, name string, storeID int64,
	start int64, end int64) *[]mdb.Product {
	return d.MockProductList
}

//GetProductManufacturerListByCatID GetProductManufacturerListByCatID
func (d *MockSix910Mysql) GetProductManufacturerListByCatID(catID int64, storeID int64) *[]string {
	return d.MockManufacturerList
}

//GetProductByCatAndManufacturer GetProductByCatAndManufacturer
func (d *MockSix910Mysql) GetProductByCatAndManufacturer(catID int64, manf string, storeID int64,
	start int64, end int64) *[]mdb.Product {
	return d.MockProductList
}

//AddVisit AddVisit
func (d *MockSix910Mysql) AddVisit(v *mdb.Visitor) bool {
	return d.MockAddVisitorResp
}

//GetVisitorData GetVisitorData
func (d *MockSix910Mysql) GetVisitorData(storeID int64) *[]mdb.VisitorData {
	return d.MockVisitorData
}
