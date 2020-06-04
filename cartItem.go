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

//cart item

//AddCartItem AddCartItem
func (d *Six910Mysql) AddCartItem(ci *mdb.CartItem) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ci.CartID, ci.Quantity, ci.ProductID)
	suc, id := d.DB.Insert(insertCartItem, a...)
	d.Log.Debug("suc in add CartItem", suc)
	d.Log.Debug("id in add CartItem", id)
	return suc, id
}

//UpdateCartItem UpdateCartItem
func (d *Six910Mysql) UpdateCartItem(ci *mdb.CartItem) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ci.Quantity, ci.ID)
	suc := d.DB.Update(updateCartItem, a...)
	return suc
}

//GetCarItem GetCarItem
func (d *Six910Mysql) GetCarItem(cartID int64, prodID int64) *mdb.CartItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, cartID, prodID)
	row := d.DB.Get(getCartItem, a...)
	rtn := d.parseCartItemRow(&row.Row)
	return rtn
}

//GetCartItemList GetCartItemList
func (d *Six910Mysql) GetCartItemList(cartID int64) *[]mdb.CartItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.CartItem
	var a []interface{}
	a = append(a, cartID)
	rows := d.DB.GetList(getCartItemList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseCartItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteCartItem DeleteCartItem
func (d *Six910Mysql) DeleteCartItem(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteCartItem, a...)
}

func (d *Six910Mysql) parseCartItemRow(foundRow *[]string) *mdb.CartItem {
	d.Log.Debug("foundRow in get CartItem", *foundRow)
	var rtn mdb.CartItem
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get CartItem", err)
		if err == nil {
			qty, cerr := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("qty err in get CartItem", cerr)
			if cerr == nil {
				ctid, cerr := strconv.ParseInt((*foundRow)[2], 10, 64)
				d.Log.Debug("ctid err in get CartItem", cerr)
				if cerr == nil {
					pid, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
					d.Log.Debug("pid err in get CartItem", cerr)
					if cerr == nil {
						rtn.ID = id
						rtn.CartID = ctid
						rtn.ProductID = pid
						rtn.Quantity = qty
					}
				}
			}
		}
	}
	return &rtn
}
