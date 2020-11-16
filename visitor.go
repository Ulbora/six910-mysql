package six910mysql

import (
	mdb "github.com/Ulbora/six910-database-interface"
	"strconv"
	"time"
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

//Visitors

//AddVisit AddVisit
func (d *Six910Mysql) AddVisit(v *mdb.Visitor) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, v.Origin, v.Host, v.IPAddress, time.Now(), v.StoreID)
	suc, id := d.DB.Insert(insertVisitor, a...)
	d.Log.Debug("suc in add Visit", suc)
	d.Log.Debug("id in add Visit", id)
	return suc
}

//GetVisitorData GetVisitorData
func (d *Six910Mysql) GetVisitorData(storeID int64) *[]mdb.VisitorData {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.VisitorData{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getVisitorCharteInfo, a...)
	//d.Log.Debug("rows Order count", *rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			var vct mdb.VisitorData
			if len(foundRow) > 0 {
				dt, err := time.Parse(dateOnlyFormat, foundRow[0])
				d.Log.Debug("dt err in get Order count", err)
				if err == nil {
					cnt, err := strconv.ParseInt(foundRow[1], 10, 64)
					d.Log.Debug("cnt err in get visit count", err)
					if err == nil {
						vct.VisitDate = dt
						vct.VisitCount = cnt
						rtn = append(rtn, vct)
					}
				}
			}
		}
	}
	return &rtn
}
