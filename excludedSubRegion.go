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

//excluded sub regions

//AddExcludedSubRegion AddExcludedSubRegion
func (d *Six910Mysql) AddExcludedSubRegion(e *mdb.ExcludedSubRegion) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, e.RegionID, e.SubRegionID, e.ShippingMethodID)
	suc, id := d.DB.Insert(insertExcludedSubRegion, a...)
	d.Log.Debug("suc in add ExcludedSubRegion", suc)
	d.Log.Debug("id in add ExcludedSubRegion", id)
	return suc, id
}

//UpdateExcludedSubRegion UpdateExcludedSubRegion
func (d *Six910Mysql) UpdateExcludedSubRegion(e *mdb.ExcludedSubRegion) bool {
	//will not be implemented at this stage
	return false
}

//GetExcludedSubRegion GetExcludedSubRegion
func (d *Six910Mysql) GetExcludedSubRegion(id int64) *mdb.ExcludedSubRegion {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getExcludedSubRegion, a...)
	rtn := d.parseExcludedSubRegionRow(&row.Row)
	return rtn
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (d *Six910Mysql) GetExcludedSubRegionList(regionID int64) *[]mdb.ExcludedSubRegion {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ExcludedSubRegion{}
	var a []interface{}
	a = append(a, regionID)
	rows := d.DB.GetList(getExcludedSubRegionList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseExcludedSubRegionRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (d *Six910Mysql) DeleteExcludedSubRegion(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteExcludedSubRegion, a...)
}

func (d *Six910Mysql) parseExcludedSubRegionRow(foundRow *[]string) *mdb.ExcludedSubRegion {
	d.Log.Debug("foundRow in get ExcludedSubRegion", *foundRow)
	var rtn mdb.ExcludedSubRegion
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get ExcludedSubRegion", err)
		if err == nil {
			rid, cerr := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("rid err in get ExcludedSubRegion", cerr)
			if cerr == nil {
				srid, cerr := strconv.ParseInt((*foundRow)[2], 10, 64)
				d.Log.Debug("srid err in get ExcludedSubRegion", cerr)
				if cerr == nil {
					smid, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
					d.Log.Debug("smid err in get ExcludedSubRegion", cerr)
					if cerr == nil {
						rtn.ID = id
						rtn.RegionID = rid
						rtn.SubRegionID = srid
						rtn.ShippingMethodID = smid
					}
				}
			}
		}
	}
	return &rtn
}
