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

//Geographic Regions

//AddRegion AddRegion
func (d *Six910Mysql) AddRegion(r *mdb.Region) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, r.Name, r.RegionCode, r.StoreID)
	suc, id := d.DB.Insert(insertRegion, a...)
	d.Log.Debug("suc in add Region", suc)
	d.Log.Debug("id in add Region", id)
	return suc, id
}

//UpdateRegion UpdateRegion
func (d *Six910Mysql) UpdateRegion(r *mdb.Region) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, r.Name, r.RegionCode, r.ID)
	suc := d.DB.Update(updateRegion, a...)
	return suc
}

//GetRegion GetRegion
func (d *Six910Mysql) GetRegion(id int64) *mdb.Region {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getRegion, a...)
	rtn := d.parseRegionRow(&row.Row)
	return rtn
}

//GetRegionList GetRegionList
func (d *Six910Mysql) GetRegionList(storeID int64) *[]mdb.Region {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.Region
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getRegionList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseRegionRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteRegion DeleteRegion
func (d *Six910Mysql) DeleteRegion(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteRegion, a...)
}

func (d *Six910Mysql) parseRegionRow(foundRow *[]string) *mdb.Region {
	d.Log.Debug("foundRow in get Region", *foundRow)
	var rtn mdb.Region
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Region", err)
		if err == nil {
			sid, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
			d.Log.Debug("qty err in get Region", cerr)
			if cerr == nil {
				rtn.ID = id
				rtn.StoreID = sid
				rtn.Name = (*foundRow)[1]
				rtn.RegionCode = (*foundRow)[2]
			}
		}
	}
	return &rtn
}
