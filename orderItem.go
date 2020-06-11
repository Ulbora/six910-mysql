package six910mysql

import (
	"strconv"

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

//Order Items

//AddOrderItem AddOrderItem
func (d *Six910Mysql) AddOrderItem(i *mdb.OrderItem) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, i.OrderID, i.ProductID, i.ProductName, i.ProductShortDesc, i.Quantity,
		i.Dropship, i.BackOrdered)
	suc, id := d.DB.Insert(insertOrderItem, a...)
	d.Log.Debug("suc in add OrderItem", suc)
	d.Log.Debug("id in add OrderItem", id)
	return suc, id
}

//UpdateOrderItem UpdateOrderItem
func (d *Six910Mysql) UpdateOrderItem(i *mdb.OrderItem) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, i.Quantity, i.Dropship, i.BackOrdered, i.ID)
	suc := d.DB.Update(updateOrderItem, a...)
	return suc
}

//GetOrderItem GetOrderItem
func (d *Six910Mysql) GetOrderItem(id int64) *mdb.OrderItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getOrderItem, a...)
	rtn := d.parseOrderItemRow(&row.Row)
	return rtn
}

//GetOrderItemList GetOrderItemList
func (d *Six910Mysql) GetOrderItemList(orderID int64) *[]mdb.OrderItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.OrderItem
	var a []interface{}
	a = append(a, orderID)
	rows := d.DB.GetList(getOrderItemList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseOrderItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteOrderItem DeleteOrderItem
func (d *Six910Mysql) DeleteOrderItem(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteOrderItem, a...)
}

func (d *Six910Mysql) parseOrderItemRow(foundRow *[]string) *mdb.OrderItem {
	d.Log.Debug("foundRow in get OrderItem", *foundRow)
	var rtn mdb.OrderItem
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get OrderItem", err)
		if err == nil {
			oid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("oid err in get OrderItem", err)
			if err == nil {
				pid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
				d.Log.Debug("pid err in get OrderItem", err)
				if err == nil {
					qty, err := strconv.ParseInt((*foundRow)[5], 10, 64)
					d.Log.Debug("qty err in get OrderItem", err)
					if err == nil {
						dship, enerr := strconv.ParseBool((*foundRow)[6])
						if enerr == nil {
							back, enerr := strconv.ParseBool((*foundRow)[7])
							if enerr == nil {
								rtn.ID = id
								rtn.OrderID = oid
								rtn.ProductID = pid
								rtn.Quantity = qty
								rtn.Dropship = dship
								rtn.BackOrdered = back
								rtn.ProductName = (*foundRow)[3]
								rtn.ProductShortDesc = (*foundRow)[4]
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
