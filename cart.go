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

//Cart

//AddCart AddCart
func (d *Six910Mysql) AddCart(c *mdb.Cart) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.StoreID, c.CustomerID, time.Now())
	suc, id := d.DB.Insert(insertCart, a...)
	d.Log.Debug("suc in add Cart", suc)
	d.Log.Debug("id in add Cart", id)
	return suc, id
}

//UpdateCart UpdateCart
func (d *Six910Mysql) UpdateCart(c *mdb.Cart) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, time.Now(), c.CustomerID, c.ID)
	suc := d.DB.Update(updateCart, a...)
	return suc
}

//GetCart GetCart
func (d *Six910Mysql) GetCart(cid int64) *mdb.Cart {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn *mdb.Cart
	var a []interface{}
	a = append(a, cid)
	rows := d.DB.GetList(getCart, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rtn = d.parseCartRow(&foundRow)
			break
		}
	}
	return rtn
}

//DeleteCart DeleteCart
func (d *Six910Mysql) DeleteCart(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteCart, a...)
}

func (d *Six910Mysql) parseCartRow(foundRow *[]string) *mdb.Cart {
	d.Log.Debug("foundRow in get Cart", *foundRow)
	var rtn mdb.Cart
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Cart", err)
		if err == nil {
			cid, cerr := strconv.ParseInt((*foundRow)[2], 10, 64)
			d.Log.Debug("cid err in get Cart", cerr)
			if cerr == nil {
				sid, cerr := strconv.ParseInt((*foundRow)[1], 10, 64)
				d.Log.Debug("sid err in get Cart", cerr)
				if cerr == nil {
					eTime, eerr := time.Parse(timeFormat, (*foundRow)[3])
					d.Log.Debug("eTime err in get store", eerr)
					if eerr == nil {
						uTime, _ := time.Parse(timeFormat, (*foundRow)[4])
						rtn.ID = id
						rtn.CustomerID = cid
						rtn.StoreID = sid
						rtn.DateEntered = eTime
						rtn.DateUpdated = uTime
					}
				}
			}
		}
	}
	return &rtn
}
