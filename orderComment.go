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

//Order Comments

//AddOrderComments AddOrderComments
func (d *Six910Mysql) AddOrderComments(c *mdb.OrderComment) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Comment, c.Username, c.OrderID)
	suc, id := d.DB.Insert(insertOrderComment, a...)
	d.Log.Debug("suc in add OrderComment", suc)
	d.Log.Debug("id in add OrderComment", id)
	return suc, id
}

//GetOrderCommentList GetOrderCommentList
func (d *Six910Mysql) GetOrderCommentList(orderID int64) *[]mdb.OrderComment {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.OrderComment
	var a []interface{}
	a = append(a, orderID)
	rows := d.DB.GetList(getOrderCommentList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseOrderCommentRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *Six910Mysql) parseOrderCommentRow(foundRow *[]string) *mdb.OrderComment {
	d.Log.Debug("foundRow in get OrderComment", *foundRow)
	var rtn mdb.OrderComment
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get OrderComment", err)
		if err == nil {
			oid, err := strconv.ParseInt((*foundRow)[3], 10, 64)
			d.Log.Debug("oid err in get OrderComment", err)
			if err == nil {
				rtn.ID = id
				rtn.OrderID = oid
				rtn.Comment = (*foundRow)[1]
				rtn.Username = (*foundRow)[2]
			}
		}
	}
	return &rtn
}
