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

//Geographic Sub Regions

//AddSubRegion AddSubRegion
func (d *Six910Mysql) AddSubRegion(s *mdb.SubRegion) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Name, s.SubRegionCode, s.RegionID)
	suc, id := d.DB.Insert(insertSubRegion, a...)
	d.Log.Debug("suc in add SubRegion", suc)
	d.Log.Debug("id in add SubRegion", id)
	return suc, id
}

//UpdateSubRegion UpdateSubRegion
func (d *Six910Mysql) UpdateSubRegion(s *mdb.SubRegion) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Name, s.SubRegionCode, s.ID)
	suc := d.DB.Update(updateSubRegion, a...)
	return suc
}

//GetSubRegion GetSubRegion
func (d *Six910Mysql) GetSubRegion(id int64) *mdb.SubRegion {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getSubRegion, a...)
	rtn := d.parseSubRegionRow(&row.Row)
	return rtn
}

//GetSubRegionList GetSubRegionList
func (d *Six910Mysql) GetSubRegionList(regionID int64) *[]mdb.SubRegion {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.SubRegion{}
	var a []interface{}
	a = append(a, regionID)
	rows := d.DB.GetList(getSubRegionList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseSubRegionRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteSubRegion DeleteSubRegion
func (d *Six910Mysql) DeleteSubRegion(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteSubRegion, a...)
}

func (d *Six910Mysql) parseSubRegionRow(foundRow *[]string) *mdb.SubRegion {
	d.Log.Debug("foundRow in get SubRegion", *foundRow)
	var rtn mdb.SubRegion
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get SubRegion", err)
		if err == nil {
			rid, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
			d.Log.Debug("rid err in get SubRegion", cerr)
			if cerr == nil {
				rtn.ID = id
				rtn.RegionID = rid
				rtn.Name = (*foundRow)[1]
				rtn.SubRegionCode = (*foundRow)[2]
			}
		}
	}
	return &rtn
}
