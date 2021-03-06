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

//product category

//AddProductCategory AddProductCategory
func (d *Six910Mysql) AddProductCategory(pc *mdb.ProductCategory) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, pc.CategoryID, pc.ProductID)
	suc, id := d.DB.Insert(insertProductCategory, a...)
	d.Log.Debug("suc in add ProductCategory", suc)
	d.Log.Debug("id in add ProductCategory", id)
	return suc
}

//GetProductCategoryList GetProductCategoryList
func (d *Six910Mysql) GetProductCategoryList(productID int64) *[]int64 {
	var rtn []int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, productID)
	rows := d.DB.GetList(getProductCategory, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			cid, err := strconv.ParseInt((foundRow)[0], 10, 64)
			d.Log.Debug("id err in get GetProductCategory", err)
			rtn = append(rtn, cid)
		}
	}
	return &rtn
}

//DeleteProductCategory DeleteProductCategory
func (d *Six910Mysql) DeleteProductCategory(pc *mdb.ProductCategory) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, pc.CategoryID, pc.ProductID)
	return d.DB.Delete(deleteProductCategory, a...)
}
