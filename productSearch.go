package six910mysql

import (
	//"strconv"
	//"time"

	// "fmt"

	"strings"

	"github.com/Ulbora/dbinterface"
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

//GetProductManufacturerListByProductName GetProductManufacturerListByProductName
func (d *Six910Mysql) GetProductManufacturerListByProductName(name string, storeID int64) *[]string {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []string{}
	var a []interface{}
	a = append(a, "%"+name+"%", storeID)
	rows := d.DB.GetList(getProductManufacturerListByProductName, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			manf := (foundRow)[0]
			rtn = append(rtn, manf)
		}
	}
	return &rtn
}

//GetProductManufacturerListByProductSearch GetProductManufacturerListByProductSearch
func (d *Six910Mysql) GetProductManufacturerListByProductSearch(attrs string, storeID int64) *[]string {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []string{}
	var a []interface{}
	var attrbs = strings.Split(attrs, " ")
	var revind []string
	var sdesc = "%"
	for _, da := range attrbs {
		sdesc += da + "%"
		revind = append(revind, da)
	}
	var rsdesc = "%"
	for i := len(revind) - 1; i >= 0; i-- {
		rsdesc += revind[i] + "%"
	}
	a = append(a, sdesc, rsdesc, storeID)
	rows := d.DB.GetList(getProductManufacturerListByProductSearch, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			manf := (foundRow)[0]
			rtn = append(rtn, manf)
		}
	}
	return &rtn
}

//GetProductByNameAndManufacturerName GetProductByNameAndManufacturerName
func (d *Six910Mysql) GetProductByNameAndManufacturerName(manf string, name string, storeID int64,
	start int64, end int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, manf, "%"+name+"%", storeID, start, end)
	rows := d.DB.GetList(getProductByNameAndManfName, a...)
	if rows != nil && len(rows.Rows) != 0 {
		sfoundRows := rows.Rows
		for r := range sfoundRows {
			foundRow := sfoundRows[r]
			srowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *srowContent)
		}
	}
	return &rtn
}

//GetProductManufacturerListByCatID GetProductManufacturerListByCatID
func (d *Six910Mysql) GetProductManufacturerListByCatID(catID int64, storeID int64) *[]string {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []string{}
	var a []interface{}
	a = append(a, catID, storeID)
	rows := d.DB.GetList(getProductManufacturerListByCatID, a...)
	if rows != nil && len(rows.Rows) != 0 {
		csfoundRows := rows.Rows
		for r := range csfoundRows {
			foundRow := csfoundRows[r]
			manf := (foundRow)[0]
			rtn = append(rtn, manf)
		}
	}
	return &rtn
}

//GetProductByCatAndManufacturer GetProductByCatAndManufacturer
func (d *Six910Mysql) GetProductByCatAndManufacturer(catID int64, manf string, storeID int64,
	start int64, end int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, catID, manf, storeID, start, end)
	rows := d.DB.GetList(getProductByCatAndManufacturer, a...)
	if rows != nil && len(rows.Rows) != 0 {
		cssfoundRows := rows.Rows
		for r := range cssfoundRows {
			foundRow := cssfoundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//ProductSearch ProductSearch
func (d *Six910Mysql) ProductSearch(p *mdb.ProductSearch) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var rows *dbinterface.DbRows

	if p.ProductID == 0 {
		var revind []string
		var sdesc = "%"
		for _, da := range *(p).DescAttributes {
			sdesc += da + "%"
			revind = append(revind, da)
		}
		var rsdesc = "%"
		for i := len(revind) - 1; i >= 0; i-- {
			rsdesc += revind[i] + "%"
		}
		var a []interface{}
		a = append(a, sdesc, rsdesc, p.StoreID, p.Start, p.End)
		rows = d.DB.GetList(productSearch, a...)
	} else {
		var a []interface{}
		a = append(a, p.ProductID, p.StoreID, p.Start, p.End)
		rows = d.DB.GetList(subProductSearch, a...)
	}
	if rows != nil && len(rows.Rows) != 0 {
		cssfoundRows := rows.Rows
		for r := range cssfoundRows {
			foundRow := cssfoundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}
