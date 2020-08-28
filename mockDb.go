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

	MockAddSecuritySuccess    bool
	MockSecurityID            int64
	MockUpdateSecuritySuccess bool
	MockSecurity              *mdb.Security
	MockDeleteSecuritySuccess bool

	MockAddStoreSuccess    bool
	MockStoreID            int64
	MockUpdateStoreSuccess bool
	MockStore              *mdb.Store
	MockStoreCount         int64
	MockDeleteStoreSuccess bool

	MockAddCustomerSuccess    bool
	MockCustomerID            int64
	MockUpdateCustomerSuccess bool
	MockCustomer              *mdb.Customer
	MockCustomerList          *[]mdb.Customer
	MockDeleteCustomerSuccess bool

	MockAddLocalAccountSuccess    bool
	MockLocalAccountID            int64
	MockUpdateLocalAccountSuccess bool
	MockLocalAccount              *mdb.LocalAccount
	MockLocalAccountList          *[]mdb.LocalAccount
	MockDeleteLocalAccountSuccess bool

	MockAddDistributorSuccess    bool
	MockDistributorID            int64
	MockUpdateDistributorSuccess bool
	MockDistributor              *mdb.Distributor
	MockDistributorList          *[]mdb.Distributor
	MockDeleteDistributorSuccess bool

	MockAddCartSuccess    bool
	MockCartID            int64
	MockUpdateCartSuccess bool
	MockCart              *mdb.Cart
	MockCartList          *[]mdb.Cart
	MockDeleteCartSuccess bool

	MockAddCartItemSuccess    bool
	MockCartItemID            int64
	MockUpdateCartItemSuccess bool
	MockCartItem              *mdb.CartItem
	MockCartItemList          *[]mdb.CartItem
	MockDeleteCartItemSuccess bool

	MockAddAddressSuccess    bool
	MockAddressID            int64
	MockUpdateAddressSuccess bool
	MockAddress              *mdb.Address
	MockAddressList          *[]mdb.Address
	MockDeleteAddressSuccess bool

	MockAddCategorySuccess    bool
	MockCategoryID            int64
	MockUpdateCategorySuccess bool
	MockCategory              *mdb.Category
	MockCategoryList          *[]mdb.Category
	MockDeleteCategorySuccess bool

	MockAddShippingMethodSuccess    bool
	MockShippingMethodID            int64
	MockUpdateShippingMethodSuccess bool
	MockShippingMethod              *mdb.ShippingMethod
	MockShippingMethodList          *[]mdb.ShippingMethod
	MockDeleteShippingMethodSuccess bool

	MockAddInsuranceSuccess    bool
	MockInsuranceID            int64
	MockUpdateInsuranceSuccess bool
	MockInsurance              *mdb.Insurance
	MockInsuranceList          *[]mdb.Insurance
	MockDeleteInsuranceSuccess bool

	MockAddProductSuccess    bool
	MockProductID            int64
	MockUpdateProductSuccess bool
	MockProduct              *mdb.Product
	MockProductList          *[]mdb.Product
	MockDeleteProductSuccess bool

	MockAddRegionSuccess    bool
	MockRegionID            int64
	MockUpdateRegionSuccess bool
	MockRegion              *mdb.Region
	MockRegionList          *[]mdb.Region
	MockDeleteRegionSuccess bool

	MockAddSubRegionSuccess    bool
	MockSubRegionID            int64
	MockUpdateSubRegionSuccess bool
	MockSubRegion              *mdb.SubRegion
	MockSubRegionList          *[]mdb.SubRegion
	MockDeleteSubRegionSuccess bool

	MockAddExcludedSubRegionSuccess    bool
	MockExcludedSubRegionID            int64
	MockUpdateExcludedSubRegionSuccess bool
	MockExcludedSubRegion              *mdb.ExcludedSubRegion
	MockExcludedSubRegionList          *[]mdb.ExcludedSubRegion
	MockDeleteExcludedSubRegionSuccess bool

	MockAddIncludedSubRegionSuccess    bool
	MockIncludedSubRegionID            int64
	MockUpdateIncludedSubRegionSuccess bool
	MockIncludedSubRegion              *mdb.IncludedSubRegion
	MockIncludedSubRegionList          *[]mdb.IncludedSubRegion
	MockDeleteIncludedSubRegionSuccess bool

	MockAddZoneZipSuccess    bool
	MockZoneZipID            int64
	MockUpdateZoneZipSuccess bool
	MockZoneZipRegion        *mdb.ZoneZip
	MockZoneZipList          *[]mdb.ZoneZip
	MockDeleteZoneZipSuccess bool

	MockAddProductCategorySuccess    bool
	MockProductCategoryID            int64
	MockUpdateProductCategorySuccess bool
	MockProductCategoryRegion        *mdb.ProductCategory
	MockProductCategoryList          *[]mdb.ProductCategory
	MockDeleteProductCategorySuccess bool

	MockAddOrderSuccess    bool
	MockOrderID            int64
	MockUpdateOrderSuccess bool
	MockOrder              *mdb.Order
	MockOrderList          *[]mdb.Order
	MockDeleteOrderSuccess bool

	MockAddOrderItemSuccess    bool
	MockOrderItemID            int64
	MockUpdateOrderItemSuccess bool
	MockOrderItem              *mdb.OrderItem
	MockOrderItemList          *[]mdb.OrderItem
	MockDeleteOrderItemSuccess bool

	MockAddOrderCommentSuccess    bool
	MockOrderCommentID            int64
	MockUpdateOrderCommentSuccess bool
	MockOrderComment              *mdb.OrderComment
	MockOrderCommentList          *[]mdb.OrderComment
	MockDeleteOrderCommentSuccess bool

	MockAddOrderTransactionSuccess    bool
	MockOrderTransactionID            int64
	MockUpdateOrderTransactionSuccess bool
	MockOrderTransaction              *mdb.OrderTransaction
	MockOrderTransactionList          *[]mdb.OrderTransaction
	MockDeleteOrderTransactionSuccess bool

	MockAddShipmentSuccess    bool
	MockShipmentID            int64
	MockUpdateShipmentSuccess bool
	MockShipment              *mdb.Shipment
	MockShipmentList          *[]mdb.Shipment
	MockDeleteShipmentSuccess bool

	MockAddShipmentBoxSuccess    bool
	MockShipmentBoxID            int64
	MockUpdateShipmentBoxSuccess bool
	MockShipmentBox              *mdb.ShipmentBox
	MockShipmentBoxList          *[]mdb.ShipmentBox
	MockDeleteShipmentBoxSuccess bool

	MockAddShipmentItemSuccess    bool
	MockShipmentItemID            int64
	MockUpdateShipmentItemSuccess bool
	MockShipmentItem              *mdb.ShipmentItem
	MockShipmentItemList          *[]mdb.ShipmentItem
	MockDeleteShipmentItemSuccess bool

	MockAddPluginSuccess    bool
	MockPluginID            int64
	MockUpdatePluginSuccess bool
	MockPlugin              *mdb.Plugins
	MockPluginList          *[]mdb.Plugins
	MockDeletePluginSuccess bool

	MockAddStorePluginSuccess    bool
	MockStorePluginID            int64
	MockUpdateStorePluginSuccess bool
	MockStorePlugin              *mdb.StorePlugins
	MockStorePluginList          *[]mdb.StorePlugins
	MockDeleteStorePluginSuccess bool

	MockAddPaymentGatewaySuccess    bool
	MockPaymentGatewayID            int64
	MockUpdatePaymentGatewaySuccess bool
	MockPaymentGateway              *mdb.PaymentGateway
	MockPaymentGatewayList          *[]mdb.PaymentGateway
	MockDeletePaymentGatewaySuccess bool

	MockAddShippingCarrierSuccess    bool
	MockShippingCarrierID            int64
	MockUpdateShippingCarrierSuccess bool
	MockShippingCarrier              *mdb.ShippingCarrier
	MockShippingCarrierList          *[]mdb.ShippingCarrier
	MockDeleteShippingCarrierSuccess bool

	MockAddLocalDataStoreSuccess    bool
	MockLocalDataStoreID            int64
	MockUpdateLocalDataStoreSuccess bool
	MockLocalDataStore              *mdb.LocalDataStore
	MockLocalDataStoreList          *[]mdb.LocalDataStore
	MockDeleteLocalDataStoreSuccess bool

	MockAddInstancesSuccess    bool
	MockInstancesID            int64
	MockUpdateInstancesSuccess bool
	MockInstances              *mdb.Instances
	MockInstancesList          *[]mdb.Instances
	MockDeleteInstancesSuccess bool

	MockAddDataStoreWriteLockSuccess    bool
	MockDataStoreWriteLockID            int64
	MockUpdateDataStoreWriteLockSuccess bool
	MockDataStoreWriteLock              *mdb.DataStoreWriteLock
	MockDataStoreWriteLockList          *[]mdb.DataStoreWriteLock
	MockDeleteDataStoreWriteLockSuccess bool

	MockAddTaxRateSuccess    bool
	MockTaxRateID            int64
	MockUpdateTaxRateSuccess bool
	MockTaxRate              *mdb.TaxRate
	MockTaxRateList          *[]mdb.TaxRate
	MockDeleteTaxRateSuccess bool
}

//GetNew GetNew
func (d *MockSix910Mysql) GetNew() sdbi.Six910DB {
	return d
}
