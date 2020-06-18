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

//included sub regions

//AddIncludedSubRegion AddIncludedSubRegion
func (d *Six910Mysql) AddIncludedSubRegion(e *mdb.IncludedSubRegion) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, e.RegionID, e.SubRegionID, e.ShippingMethodID)
	suc, id := d.DB.Insert(insertIncludedSubRegion, a...)
	d.Log.Debug("suc in add IncludedSubRegion", suc)
	d.Log.Debug("id in add IncludedSubRegion", id)
	return suc, id
}

//UpdateIncludedSubRegion UpdateIncludedSubRegion
func (d *Six910Mysql) UpdateIncludedSubRegion(e *mdb.IncludedSubRegion) bool {
	//will not be implemented at this stage
	return false
}

//GetIncludedSubRegion GetIncludedSubRegion
func (d *Six910Mysql) GetIncludedSubRegion(id int64) *mdb.IncludedSubRegion {
	//will not be implemented at this stage
	return nil
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (d *Six910Mysql) GetIncludedSubRegionList(regionID int64) *[]mdb.IncludedSubRegion {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.IncludedSubRegion{}
	var a []interface{}
	a = append(a, regionID)
	rows := d.DB.GetList(getIncludedSubRegionList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseIncludedSubRegionRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (d *Six910Mysql) DeleteIncludedSubRegion(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteIncludedSubRegion, a...)
}

func (d *Six910Mysql) parseIncludedSubRegionRow(foundRow *[]string) *mdb.IncludedSubRegion {
	d.Log.Debug("foundRow in get IncludedSubRegion", *foundRow)
	var rtn mdb.IncludedSubRegion
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get IncludedSubRegion", err)
		if err == nil {
			rid, cerr := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("rid err in get IncludedSubRegion", cerr)
			if cerr == nil {
				srid, cerr := strconv.ParseInt((*foundRow)[2], 10, 64)
				d.Log.Debug("srid err in get IncludedSubRegion", cerr)
				if cerr == nil {
					smid, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
					d.Log.Debug("smid err in get IncludedSubRegion", cerr)
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
