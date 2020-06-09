package six910mysql

import (
	"strconv"
	"time"

	mdb "github.com/Ulbora/six910-database-interface"
)

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

//Orders

//AddOrder AddOrder
func (d *Six910Mysql) AddOrder(o *mdb.Order) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, o.OrderDate, o.Status, o.Subtotal, o.ShippingHandling, o.Insurance, o.Taxes, o.Total,
		o.CustomerID, o.BillingAddressID, o.ShippingAddressID, o.CustomerName, o.BillingAddress,
		o.ShippingAddress, o.StoreID, o.OrderNumber, o.OrderType, o.Pickup, o.Username)
	suc, id := d.DB.Insert(insertOrder, a...)
	d.Log.Debug("suc in add Order", suc)
	d.Log.Debug("id in add Order", id)
	return suc, id
}

//UpdateOrder UpdateOrder
func (d *Six910Mysql) UpdateOrder(o *mdb.Order) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, o.Updated, o.Status, o.Subtotal, o.ShippingHandling, o.Insurance, o.Taxes, o.Total,
		o.BillingAddressID, o.ShippingAddressID, o.CustomerName, o.BillingAddress,
		o.ShippingAddress, o.OrderType, o.Pickup, o.Username, o.ID)
	suc := d.DB.Update(updateOrder, a...)
	return suc
}

//GetOrder GetOrder
func (d *Six910Mysql) GetOrder(id int64) *mdb.Order {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getOrder, a...)
	rtn := d.parseOrderRow(&row.Row)
	return rtn
}

//GetOrderList GetOrderList
func (d *Six910Mysql) GetOrderList(cid int64) *[]mdb.Order {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.Order
	var a []interface{}
	a = append(a, cid)
	rows := d.DB.GetList(getOrderByCid, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseOrderRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteOrder DeleteOrder
func (d *Six910Mysql) DeleteOrder(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteOrder, a...)
}

func (d *Six910Mysql) parseOrderRow(foundRow *[]string) *mdb.Order {
	d.Log.Debug("foundRow in get Order", *foundRow)
	var rtn mdb.Order
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Order", err)
		if err == nil {
			cid, oerr := strconv.ParseInt((*foundRow)[9], 10, 64)
			d.Log.Debug("oauthID err in get Order", oerr)
			if oerr == nil {
				oTime, eerr := time.Parse(timeFormat, (*foundRow)[1])
				d.Log.Debug("eTime err in get store", eerr)
				if eerr == nil {
					uTime, _ := time.Parse(timeFormat, (*foundRow)[2])
					pickup, enerr := strconv.ParseBool((*foundRow)[18])
					if enerr == nil {
						baid, err := strconv.ParseInt((*foundRow)[10], 10, 64)
						d.Log.Debug("oauthID err in get Order", oerr)
						if err == nil {
							said, err := strconv.ParseInt((*foundRow)[11], 10, 64)
							d.Log.Debug("oauthID err in get Order", oerr)
							if err == nil {
								sid, err := strconv.ParseInt((*foundRow)[15], 10, 64)
								d.Log.Debug("oauthID err in get Order", oerr)
								if err == nil {
									subtot, err := strconv.ParseFloat((*foundRow)[4], 64)
									d.Log.Debug("cost err in get Product", err)
									if err == nil {
										sandh, err := strconv.ParseFloat((*foundRow)[5], 64)
										d.Log.Debug("cost err in get Product", err)
										if err == nil {
											ins, err := strconv.ParseFloat((*foundRow)[6], 64)
											d.Log.Debug("cost err in get Product", err)
											if err == nil {
												tax, err := strconv.ParseFloat((*foundRow)[7], 64)
												d.Log.Debug("cost err in get Product", err)
												if err == nil {
													tot, err := strconv.ParseFloat((*foundRow)[8], 64)
													d.Log.Debug("cost err in get Product", err)
													if err == nil {
														rtn.ID = id
														rtn.CustomerID = cid
														rtn.OrderDate = oTime
														rtn.Updated = uTime
														rtn.Pickup = pickup
														rtn.BillingAddressID = baid
														rtn.ShippingAddressID = said
														rtn.StoreID = sid
														rtn.Subtotal = subtot
														rtn.ShippingHandling = sandh
														rtn.Insurance = ins
														rtn.Taxes = tax
														rtn.Total = tot
														rtn.Status = (*foundRow)[3]
														rtn.CustomerName = (*foundRow)[12]
														rtn.BillingAddress = (*foundRow)[13]
														rtn.ShippingAddress = (*foundRow)[14]
														rtn.OrderNumber = (*foundRow)[16]
														rtn.OrderType = (*foundRow)[17]
														rtn.Username = (*foundRow)[19]
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
