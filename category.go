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

//category

//AddCategory AddCategory
func (d *Six910Mysql) AddCategory(c *mdb.Category) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Name, c.Description, c.Image, c.Thumbnail, c.StoreID, c.ParentCategoryID)
	suc, id := d.DB.Insert(insertCategory, a...)
	d.Log.Debug("suc in add Category", suc)
	d.Log.Debug("id in add Category", id)
	return suc, id
}

//UpdateCategory UpdateCategory
func (d *Six910Mysql) UpdateCategory(c *mdb.Category) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Name, c.Description, c.Image, c.Thumbnail, c.ParentCategoryID, c.ID)
	suc := d.DB.Update(updateCategory, a...)
	return suc
}

//GetCategory GetCategory
func (d *Six910Mysql) GetCategory(id int64) *mdb.Category {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getCategory, a...)
	rtn := d.parseCategoryRow(&row.Row)
	return rtn
}

//GetCategoryList GetCategoryList
func (d *Six910Mysql) GetCategoryList(storeID int64) *[]mdb.Category {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.Category
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getCategoryList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseCategoryRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetSubCategoryList GetSubCategoryList
func (d *Six910Mysql) GetSubCategoryList(catID int64) *[]mdb.Category {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.Category
	var a []interface{}
	a = append(a, catID)
	rows := d.DB.GetList(getSubCategoryList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseCategoryRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteCategory DeleteCategory
func (d *Six910Mysql) DeleteCategory(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id, id)
	return d.DB.Delete(deleteCategory, a...)
}

func (d *Six910Mysql) parseCategoryRow(foundRow *[]string) *mdb.Category {
	d.Log.Debug("foundRow in get Category", *foundRow)
	var rtn mdb.Category
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Category", err)
		if err == nil {
			pid, perr := strconv.ParseInt((*foundRow)[5], 10, 64)
			d.Log.Debug("cid err in get Category", perr)
			if perr == nil {
				sid, serr := strconv.ParseInt((*foundRow)[6], 10, 64)
				d.Log.Debug("cid err in get Category", serr)
				if serr == nil {
					rtn.ID = id
					rtn.ParentCategoryID = pid
					rtn.StoreID = sid
					rtn.Name = (*foundRow)[1]
					rtn.Description = (*foundRow)[2]
					rtn.Image = (*foundRow)[3]
					rtn.Thumbnail = (*foundRow)[4]
				}
			}
		}
	}
	return &rtn
}
