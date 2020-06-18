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

//limit exclusions and inclusions to a zip code

//AddZoneZip AddZoneZip
func (d *Six910Mysql) AddZoneZip(z *mdb.ZoneZip) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, z.ZipCode, z.IncludedSubRegionID, z.ExcludedSubRegionID)
	suc, id := d.DB.Insert(insertZoneZip, a...)
	d.Log.Debug("suc in add ZoneZip", suc)
	d.Log.Debug("id in add ZoneZip", id)
	return suc, id
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (d *Six910Mysql) GetZoneZipListByExclusion(exID int64) *[]mdb.ZoneZip {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ZoneZip{}
	var a []interface{}
	a = append(a, exID)
	rows := d.DB.GetList(getZipExclustionList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseZoneZipRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (d *Six910Mysql) GetZoneZipListByInclusion(incID int64) *[]mdb.ZoneZip {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ZoneZip{}
	var a []interface{}
	a = append(a, incID)
	rows := d.DB.GetList(getZipInclustionList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseZoneZipRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteZoneZip DeleteZoneZip
func (d *Six910Mysql) DeleteZoneZip(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteZoneZip, a...)
}

func (d *Six910Mysql) parseZoneZipRow(foundRow *[]string) *mdb.ZoneZip {
	d.Log.Debug("foundRow in get ZoneZip", *foundRow)
	var rtn mdb.ZoneZip
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get ZoneZip", err)
		if err == nil {
			inc, cerr := strconv.ParseInt((*foundRow)[2], 10, 64)
			d.Log.Debug("inc err in get ZoneZip", cerr)
			if cerr == nil {
				exc, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
				d.Log.Debug("exc err in get ZoneZip", cerr)
				if cerr == nil {
					rtn.ID = id
					rtn.IncludedSubRegionID = inc
					rtn.ExcludedSubRegionID = exc
					rtn.ZipCode = (*foundRow)[1]
				}
			}
		}
	}
	return &rtn
}
