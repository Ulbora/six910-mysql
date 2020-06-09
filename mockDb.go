package six910mysql

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

// go mod init github.com/Ulbora/six910-mysql

import (
	lg "github.com/Ulbora/Level_Logger"
	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/six910-database-interface"
	sdbi "github.com/Ulbora/six910-database-interface"
)

//MockSix910Mysql MockSix910Mysql
type MockSix910Mysql struct {
	DB  dbi.Database
	Log *lg.Logger
}

//GetNew GetNew
func (d *MockSix910Mysql) GetNew() sdbi.Six910DB {
	return d
}

//AddSecurity AddSecurity
func (d *MockSix910Mysql) AddSecurity(s *mdb.Security) (bool, int64) {
	return false, 0
}

//UpdateSecurity UpdateSecurity
func (d *MockSix910Mysql) UpdateSecurity(s *mdb.Security) bool {
	return false
}

//GetSecurity GetSecurity
func (d *MockSix910Mysql) GetSecurity() *mdb.Security {
	return nil
}

//DeleteSecurity DeleteSecurity
func (d *MockSix910Mysql) DeleteSecurity() bool {
	return false
}

//stores

//AddStore AddStore
func (d *MockSix910Mysql) AddStore(s *mdb.Store) (bool, int64) {
	return false, 0
}

//UpdateStore UpdateStore
func (d *MockSix910Mysql) UpdateStore(s *mdb.Store) bool {
	return false
}

//GetStore GetStore
func (d *MockSix910Mysql) GetStore(sname string) *mdb.Store {
	return nil
}

//GetStoreID GetStoreID
func (d *MockSix910Mysql) GetStoreID(id int64) *mdb.Store {
	return nil
}

//GetStoreByLocal GetStoreByLocal
func (d *MockSix910Mysql) GetStoreByLocal(localDomain string) *mdb.Store {
	return nil
}

//DeleteStore DeleteStore
func (d *MockSix910Mysql) DeleteStore(id int64) bool {
	return false
}

//customer

//AddCustomer AddCustomer
func (d *MockSix910Mysql) AddCustomer(c *mdb.Customer) (bool, int64) {
	return false, 0
}

//UpdateCustomer UpdateCustomer
func (d *MockSix910Mysql) UpdateCustomer(c *mdb.Customer) bool {
	return false
}

//GetCustomer GetCustomer
func (d *MockSix910Mysql) GetCustomer(email string, storeID int64) *mdb.Customer {
	return nil
}

//GetCustomerList GetCustomerList
func (d *MockSix910Mysql) GetCustomerList(storeID int64) *[]mdb.Customer {
	return nil
}

//GetCustomerID GetCustomerID
func (d *MockSix910Mysql) GetCustomerID(id int64) *mdb.Customer {
	return nil
}

//DeleteCustomer DeleteCustomer
func (d *MockSix910Mysql) DeleteCustomer(id int64) bool {
	return false
}

//Local Accounts when oauth not used

//AddLocalAccount AddLocalAccount
func (d *MockSix910Mysql) AddLocalAccount(a *mdb.LocalAccount) bool {
	return false
}

//UpdateLocalAccount UpdateLocalAccount
func (d *MockSix910Mysql) UpdateLocalAccount(a *mdb.LocalAccount) bool {
	return false
}

//GetLocalAccount GetLocalAccount
func (d *MockSix910Mysql) GetLocalAccount(username string, storeID int64) *mdb.LocalAccount {
	return nil
}

//GetLocalAccountList GetLocalAccountList
func (d *MockSix910Mysql) GetLocalAccountList(store int64) *[]mdb.LocalAccount {
	return nil
}

//DeleteLocalAccount DeleteLocalAccount
func (d *MockSix910Mysql) DeleteLocalAccount(username string, storeID int64) bool {
	return false
}

//distributors

//AddDistributor AddDistributor
func (d *MockSix910Mysql) AddDistributor(ds *mdb.Distributor) (bool, int64) {
	return false, 0
}

//UpdateDistributor UpdateDistributor
func (d *MockSix910Mysql) UpdateDistributor(ds *mdb.Distributor) bool {
	return false
}

//GetDistributor GetDistributor
func (d *MockSix910Mysql) GetDistributor(id int64) *mdb.Distributor {
	return nil
}

//GetDistributorList GetDistributorList
func (d *MockSix910Mysql) GetDistributorList(store int64) *[]mdb.Distributor {
	return nil
}

//DeleteDistributor DeleteDistributor
func (d *MockSix910Mysql) DeleteDistributor(id int64) bool {
	return false
}

//Cart

//AddCart AddCart
func (d *MockSix910Mysql) AddCart(c *mdb.Cart) (bool, int64) {
	return false, 0
}

//UpdateCart UpdateCart
func (d *MockSix910Mysql) UpdateCart(c *mdb.Cart) bool {
	return false
}

//GetCart GetCart
func (d *MockSix910Mysql) GetCart(cid int64) *mdb.Cart {
	return nil
}

//DeleteCart DeleteCart
func (d *MockSix910Mysql) DeleteCart(id int64) bool {
	return false
}

//cart item

//AddCartItem AddCartItem
func (d *MockSix910Mysql) AddCartItem(ci *mdb.CartItem) (bool, int64) {
	return false, 0
}

//UpdateCartItem UpdateCartItem
func (d *MockSix910Mysql) UpdateCartItem(ci *mdb.CartItem) bool {
	return false
}

//GetCarItem GetCarItem
func (d *MockSix910Mysql) GetCarItem(cartID int64, prodID int64) *mdb.CartItem {
	return nil
}

//GetCartItemList GetCartItemList
func (d *MockSix910Mysql) GetCartItemList(cartID int64) *[]mdb.CartItem {
	return nil
}

//DeleteCartItem DeleteCartItem
func (d *MockSix910Mysql) DeleteCartItem(id int64) bool {
	return false
}

//address

//AddAddress AddAddress
func (d *MockSix910Mysql) AddAddress(a *mdb.Address) (bool, int64) {
	return false, 0
}

//UpdateAddress UpdateAddress
func (d *MockSix910Mysql) UpdateAddress(a *mdb.Address) bool {
	return false
}

//GetAddress GetAddress
func (d *MockSix910Mysql) GetAddress(id int64) *mdb.Address {
	return nil
}

//GetAddressList GetAddressList
func (d *MockSix910Mysql) GetAddressList(cid int64) *[]mdb.Address {
	return nil
}

//DeleteAddress DeleteAddress
func (d *MockSix910Mysql) DeleteAddress(id int64) bool {
	return false
}

//category

//AddCategory AddCategory
func (d *MockSix910Mysql) AddCategory(c *mdb.Category) (bool, int64) {
	return false, 0
}

//UpdateCategory UpdateCategory
func (d *MockSix910Mysql) UpdateCategory(c *mdb.Category) bool {
	return false
}

//GetCategory GetCategory
func (d *MockSix910Mysql) GetCategory(id int64) *mdb.Category {
	return nil
}

//GetCategoryList GetCategoryList
func (d *MockSix910Mysql) GetCategoryList(storeID int64) *[]mdb.Category {
	return nil
}

//GetSubCategoryList GetSubCategoryList
func (d *MockSix910Mysql) GetSubCategoryList(catID int64) *[]mdb.Category {
	return nil
}

//DeleteCategory DeleteCategory
func (d *MockSix910Mysql) DeleteCategory(id int64) bool {
	return false
}

//shipping method

//AddShippingMethod AddShippingMethod
func (d *MockSix910Mysql) AddShippingMethod(s *mdb.ShippingMethod) (bool, int64) {
	return false, 0
}

//UpdateShippingMethod UpdateShippingMethod
func (d *MockSix910Mysql) UpdateShippingMethod(s *mdb.ShippingMethod) bool {
	return false
}

//GetShippingMethod GetShippingMethod
func (d *MockSix910Mysql) GetShippingMethod(id int64) *mdb.ShippingMethod {
	return nil
}

//GetShippingMethodList GetShippingMethodList
func (d *MockSix910Mysql) GetShippingMethodList(storeID int64) *[]mdb.ShippingMethod {
	return nil
}

//DeleteShippingMethod DeleteShippingMethod
func (d *MockSix910Mysql) DeleteShippingMethod(id int64) bool {
	return false
}

//shipping insurance

//AddInsurance AddInsurance
func (d *MockSix910Mysql) AddInsurance(i *mdb.Insurance) (bool, int64) {
	return false, 0
}

//UpdateInsurance UpdateInsurance
func (d *MockSix910Mysql) UpdateInsurance(i *mdb.Insurance) bool {
	return false
}

//GetInsurance GetInsurance
func (d *MockSix910Mysql) GetInsurance(id int64) *mdb.Insurance {
	return nil
}

//GetInsuranceList GetInsuranceList
func (d *MockSix910Mysql) GetInsuranceList(storeID int64) *[]mdb.Insurance {
	return nil
}

//DeleteInsurance DeleteInsurance
func (d *MockSix910Mysql) DeleteInsurance(id int64) bool {
	return false
}

//product

//AddProduct AddProduct
func (d *MockSix910Mysql) AddProduct(p *mdb.Product) (bool, int64) {
	return false, 0
}

//UpdateProduct UpdateProduct
func (d *MockSix910Mysql) UpdateProduct(p *mdb.Product) bool {
	return false
}

//GetProductByID GetProductByID
func (d *MockSix910Mysql) GetProductByID(id int64) *mdb.Product {
	return nil
}

//GetProductsByName GetProductsByName
func (d *MockSix910Mysql) GetProductsByName(name string, start int64, end int64) *[]mdb.Product {
	return nil
}

//GetProductsByCaterory GetProductsByCaterory
func (d *MockSix910Mysql) GetProductsByCaterory(catID int64, start int64, end int64) *[]mdb.Product {
	return nil
}

//GetProductList GetProductList
func (d *MockSix910Mysql) GetProductList(storeID int64, start int64, end int64) *[]mdb.Product {
	return nil
}

//DeleteProduct DeleteProduct
func (d *MockSix910Mysql) DeleteProduct(id int64) bool {
	return false
}

//Geographic Regions

//AddRegion AddRegion
func (d *MockSix910Mysql) AddRegion(r *mdb.Region) (bool, int64) {
	return false, 0
}

//UpdateRegion UpdateRegion
func (d *MockSix910Mysql) UpdateRegion(r *mdb.Region) bool {
	return false
}

//GetRegion GetRegion
func (d *MockSix910Mysql) GetRegion(id int64) *mdb.Region {
	return nil
}

//GetRegionList GetRegionList
func (d *MockSix910Mysql) GetRegionList(storeID int64) *[]mdb.Region {
	return nil
}

//DeleteRegion DeleteRegion
func (d *MockSix910Mysql) DeleteRegion(id int64) bool {
	return false
}

//Geographic Sub Regions

//AddSubRegion AddSubRegion
func (d *MockSix910Mysql) AddSubRegion(s *mdb.SubRegion) (bool, int64) {
	return false, 0
}

//UpdateSubRegion UpdateSubRegion
func (d *MockSix910Mysql) UpdateSubRegion(s *mdb.SubRegion) bool {
	return false
}

//GetSubRegion GetSubRegion
func (d *MockSix910Mysql) GetSubRegion(id int64) *mdb.SubRegion {
	return nil
}

//GetSubRegionList GetSubRegionList
func (d *MockSix910Mysql) GetSubRegionList(regionID int64) *[]mdb.SubRegion {
	return nil
}

//DeleteSubRegion DeleteSubRegion
func (d *MockSix910Mysql) DeleteSubRegion(id int64) bool {
	return false
}

//excluded sub regions

//AddExcludedSubRegion AddExcludedSubRegion
func (d *MockSix910Mysql) AddExcludedSubRegion(e *mdb.ExcludedSubRegion) (bool, int64) {
	return false, 0
}

//UpdateExcludedSubRegion UpdateExcludedSubRegion
func (d *MockSix910Mysql) UpdateExcludedSubRegion(e *mdb.ExcludedSubRegion) bool {
	return false
}

//GetExcludedSubRegion GetExcludedSubRegion
func (d *MockSix910Mysql) GetExcludedSubRegion(id int64) *mdb.ExcludedSubRegion {
	return nil
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (d *MockSix910Mysql) GetExcludedSubRegionList(regionID int64) *[]mdb.ExcludedSubRegion {
	return nil
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (d *MockSix910Mysql) DeleteExcludedSubRegion(id int64) bool {
	return false
}

//included sub regions

//AddIncludedSubRegion AddIncludedSubRegion
func (d *MockSix910Mysql) AddIncludedSubRegion(e *mdb.IncludedSubRegion) (bool, int64) {
	return false, 0
}

//UpdateIncludedSubRegion UpdateIncludedSubRegion
func (d *MockSix910Mysql) UpdateIncludedSubRegion(e *mdb.IncludedSubRegion) bool {
	return false
}

//GetIncludedSubRegion GetIncludedSubRegion
func (d *MockSix910Mysql) GetIncludedSubRegion(id int64) *mdb.IncludedSubRegion {
	return nil
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (d *MockSix910Mysql) GetIncludedSubRegionList(regionID int64) *[]mdb.IncludedSubRegion {
	return nil
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (d *MockSix910Mysql) DeleteIncludedSubRegion(id int64) bool {
	return false
}

//limit exclusions and inclusions to a zip code

//AddZoneZip AddZoneZip
func (d *MockSix910Mysql) AddZoneZip(z *mdb.ZoneZip) (bool, int64) {
	return false, 0
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (d *MockSix910Mysql) GetZoneZipListByExclusion(exID int64) *[]mdb.ZoneZip {
	return nil
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (d *MockSix910Mysql) GetZoneZipListByInclusion(incID int64) *[]mdb.ZoneZip {
	return nil
}

//DeleteZoneZip DeleteZoneZip
func (d *MockSix910Mysql) DeleteZoneZip(id int64) bool {
	return false
}

//product category

//AddProductCategory AddProductCategory
func (d *MockSix910Mysql) AddProductCategory(pc *mdb.ProductCategory) bool {
	return false
}

//DeleteProductCategory DeleteProductCategory
func (d *MockSix910Mysql) DeleteProductCategory(pc *mdb.ProductCategory) bool {
	return false
}

//Orders

//AddOrder AddOrder
func (d *MockSix910Mysql) AddOrder(o *mdb.Order) (bool, int64) {
	return false, 0
}

//UpdateOrder UpdateOrder
func (d *MockSix910Mysql) UpdateOrder(o *mdb.Order) bool {
	return false
}

//GetOrder GetOrder
func (d *MockSix910Mysql) GetOrder(id int64) *mdb.Order {
	return nil
}

//GetOrderList GetOrderList
func (d *MockSix910Mysql) GetOrderList(cid int64) *[]mdb.Order {
	return nil
}

//DeleteOrder DeleteOrder
func (d *MockSix910Mysql) DeleteOrder(id int64) bool {
	return false
}

//Order Items

//AddOrderItem AddOrderItem
func (d *MockSix910Mysql) AddOrderItem(i *mdb.OrderItem) (bool, int64) {
	return false, 0
}

//UpdateOrderItem UpdateOrderItem
func (d *MockSix910Mysql) UpdateOrderItem(i *mdb.OrderItem) bool {
	return false
}

//GetOrderItem GetOrderItem
func (d *MockSix910Mysql) GetOrderItem(id int64) *mdb.OrderItem {
	return nil
}

//GetOrderItemList GetOrderItemList
func (d *MockSix910Mysql) GetOrderItemList(orderID int64) *[]mdb.OrderItem {
	return nil
}

//DeleteOrderItem DeleteOrderItem
func (d *MockSix910Mysql) DeleteOrderItem(id int64) bool {
	return false
}

//Order Comments

//AddOrderComments AddOrderComments
func (d *MockSix910Mysql) AddOrderComments(c *mdb.OrderComment) (bool, int64) {
	return false, 0
}

//GetOrderCommentList GetOrderCommentList
func (d *MockSix910Mysql) GetOrderCommentList(orderID int64) *[]mdb.OrderComment {
	return nil
}

//Order Payment Transactions

//AddOrderTransaction AddOrderTransaction
func (d *MockSix910Mysql) AddOrderTransaction(t *mdb.OrderTransaction) (bool, int64) {
	return false, 0
}

//GetOrderTransactionList GetOrderTransactionList
func (d *MockSix910Mysql) GetOrderTransactionList(orderID int64) *[]mdb.OrderTransaction {
	return nil
}

//shipment

//AddShipment AddShipment
func (d *MockSix910Mysql) AddShipment(s *mdb.Shipment) (bool, int64) {
	return false, 0
}

//UpdateShipment UpdateShipment
func (d *MockSix910Mysql) UpdateShipment(s *mdb.Shipment) bool {
	return false
}

//GetShipment GetShipment
func (d *MockSix910Mysql) GetShipment(id int64) *mdb.Shipment {
	return nil
}

//GetShipmentList GetShipmentList
func (d *MockSix910Mysql) GetShipmentList(orderID int64) *[]mdb.Shipment {
	return nil
}

//DeleteShipment DeleteShipment
func (d *MockSix910Mysql) DeleteShipment(id int64) bool {
	return false
}

//shipment boxes

//AddShipmentBox AddShipmentBox
func (d *MockSix910Mysql) AddShipmentBox(sb *mdb.ShipmentBox) (bool, int64) {
	return false, 0
}

//UpdateShipmentBox UpdateShipmentBox
func (d *MockSix910Mysql) UpdateShipmentBox(sb *mdb.ShipmentBox) bool {
	return false
}

//GetShipmentBox GetShipmentBox
func (d *MockSix910Mysql) GetShipmentBox(id int64) *mdb.ShipmentBox {
	return nil
}

//GetShipmentBoxList GetShipmentBoxList
func (d *MockSix910Mysql) GetShipmentBoxList(shipmentID int64) *[]mdb.ShipmentBox {
	return nil
}

//DeleteShipmentBox DeleteShipmentBox
func (d *MockSix910Mysql) DeleteShipmentBox(id int64) bool {
	return false
}

//Shipment Items in box

//AddShipmentItem AddShipmentItem
func (d *MockSix910Mysql) AddShipmentItem(si *mdb.ShipmentItem) (bool, int64) {
	return false, 0
}

//UpdateShipmentItem UpdateShipmentItem
func (d *MockSix910Mysql) UpdateShipmentItem(si *mdb.ShipmentItem) bool {
	return false
}

//GetShipmentItem GetShipmentItem
func (d *MockSix910Mysql) GetShipmentItem(id int64) *mdb.ShipmentItem {
	return nil
}

//GetShipmentItemList GetShipmentItemList
func (d *MockSix910Mysql) GetShipmentItemList(shipmentID int64) *[]mdb.ShipmentItem {
	return nil
}

//DeleteShipmentItem DeleteShipmentItem
func (d *MockSix910Mysql) DeleteShipmentItem(id int64) bool {
	return false
}

//Global Plugins

//AddPlugin AddPlugin
func (d *MockSix910Mysql) AddPlugin(p *mdb.Plugins) (bool, int64) {
	return false, 0
}

//UpdatePlugin UpdatePlugin
func (d *MockSix910Mysql) UpdatePlugin(p *mdb.Plugins) bool {
	return false
}

//GetPlugin GetPlugin
func (d *MockSix910Mysql) GetPlugin(id int64) *mdb.Plugins {
	return nil
}

//GetPluginList GetPluginList
func (d *MockSix910Mysql) GetPluginList(start int64, end int64) *[]mdb.Plugins {
	return nil
}

//DeletePlugin DeletePlugin
func (d *MockSix910Mysql) DeletePlugin(id int64) bool {
	return false
}

//store plugins installed

//AddStorePlugin AddStorePlugin
func (d *MockSix910Mysql) AddStorePlugin(sp *mdb.StorePlugins) (bool, int64) {
	return false, 0
}

//UpdateStorePlugin UpdateStorePlugin
func (d *MockSix910Mysql) UpdateStorePlugin(sp *mdb.StorePlugins) bool {
	return false
}

//GetStorePlugin GetStorePlugin
func (d *MockSix910Mysql) GetStorePlugin(id int64) *mdb.StorePlugins {
	return nil
}

//GetStorePluginList GetStorePluginList
func (d *MockSix910Mysql) GetStorePluginList(storeID int64) *[]mdb.StorePlugins {
	return nil
}

//DeleteStorePlugin DeleteStorePlugin
func (d *MockSix910Mysql) DeleteStorePlugin(id int64) bool {
	return false
}

//Plugins that are payment gateways

//AddPaymentGateway AddPaymentGateway
func (d *MockSix910Mysql) AddPaymentGateway(pgw *mdb.PaymentGateway) (bool, int64) {
	return false, 0
}

//UpdatePaymentGateway UpdatePaymentGateway
func (d *MockSix910Mysql) UpdatePaymentGateway(pgw *mdb.PaymentGateway) bool {
	return false
}

//GetPaymentGateways GetPaymentGateways
func (d *MockSix910Mysql) GetPaymentGateways(storeID int64) *[]mdb.PaymentGateway {
	return nil
}

//DeletePaymentGateway DeletePaymentGateway
func (d *MockSix910Mysql) DeletePaymentGateway(id int64) bool {
	return false
}

//store shipment carrier like UPS and FEDex

//AddShippingCarrier AddShippingCarrier
func (d *MockSix910Mysql) AddShippingCarrier(c *mdb.ShippingCarrier) (bool, int64) {
	return false, 0
}

//UpdateShippingCarrier UpdateShippingCarrier
func (d *MockSix910Mysql) UpdateShippingCarrier(c *mdb.ShippingCarrier) bool {
	return false
}

//GetShippingCarrierList GetShippingCarrierList
func (d *MockSix910Mysql) GetShippingCarrierList(storeID int64) *[]mdb.ShippingCarrier {
	return nil
}

//DeleteShippingCarrier DeleteShippingCarrier
func (d *MockSix910Mysql) DeleteShippingCarrier(id int64) bool {
	return false
}

//----UI Cluster installation: this is only called if UI is running in a cluster---
//Handle the situation where clients are running in a cluster
//This gives a way to make sure the json datastores are update on each node in the cluster

//----------------start datastore------------------------------------
//this gets called when a node start up and add only if it doesn't already exist

//AddLocalDatastore AddLocalDatastore
func (d *MockSix910Mysql) AddLocalDatastore(ds *mdb.LocalDataStore) bool {
	return false
}

//This get get called when a change is made to a datastore from a node in the cluster
//Or after all reloads have completed and then get set to Reload = false

//UpdateLocalDatastore UpdateLocalDatastore
func (d *MockSix910Mysql) UpdateLocalDatastore(ds *mdb.LocalDataStore) bool {
	return false
}

//This gets call by cluster nodes to see if there are pending reload

//GetLocalDatastore GetLocalDatastore
func (d *MockSix910Mysql) GetLocalDatastore(storeID int64, dataStoreName string) *mdb.LocalDataStore {
	return nil
}

//---------------------start instance--------------------
// this gets called when each instance is started and added only if never added before
//The instance name is pulled from Docker or an manually entered env variable

//AddInstance AddInstance
func (d *MockSix910Mysql) AddInstance(i *mdb.Instances) bool {
	return false
}

//This gets called after instance gets reloaded

//UpdateInstance UpdateInstance
func (d *MockSix910Mysql) UpdateInstance(i *mdb.Instances) bool {
	return false
}

//Gets called before updating an instance

//GetInstance GetInstance
func (d *MockSix910Mysql) GetInstance(name string, dataStoreName string, storeID int64) *mdb.Instances {
	return nil
}

//-------------------start write lock-------------
//gets called after UI makes changes to one of the datastores
//If the datastore already exists, the Update method is called from within add

//AddDataStoreWriteLock AddDataStoreWriteLock
func (d *MockSix910Mysql) AddDataStoreWriteLock(w *mdb.DataStoreWriteLock) (bool, int64) {
	return false, 0
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (d *MockSix910Mysql) UpdateDataStoreWriteLock(w *mdb.DataStoreWriteLock) bool {
	return false
}

//gets called from within the add method and by any node trying to update a datastore

//GetDataStoreWriteLock GetDataStoreWriteLock
func (d *MockSix910Mysql) GetDataStoreWriteLock(dataStore string, storeID int64) *mdb.DataStoreWriteLock {
	return nil
}
