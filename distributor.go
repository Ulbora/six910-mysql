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

//distributors

//AddDistributor AddDistributor
func (d *Six910Mysql) AddDistributor(ds *mdb.Distributor) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ds.Company, ds.ContactName, ds.Phone, ds.StoreID)
	suc, id := d.DB.Insert(insertDistributor, a...)
	d.Log.Debug("suc in add Distributor", suc)
	d.Log.Debug("id in add Distributor", id)
	return suc, id
}

//UpdateDistributor UpdateDistributor
func (d *Six910Mysql) UpdateDistributor(ds *mdb.Distributor) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ds.Company, ds.ContactName, ds.Phone, ds.ID)
	suc := d.DB.Update(updateDistributor, a...)
	return suc
}

//GetDistributor GetDistributor
func (d *Six910Mysql) GetDistributor(id int64) *mdb.Distributor {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getDistributor, a...)
	rtn := d.parseDistributorRow(&row.Row)
	return rtn
}

//GetDistributorList GetDistributorList
func (d *Six910Mysql) GetDistributorList(store int64) *[]mdb.Distributor {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Distributor{}
	var a []interface{}
	a = append(a, store)
	rows := d.DB.GetList(getDistributorList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseDistributorRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteDistributor DeleteDistributor
func (d *Six910Mysql) DeleteDistributor(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteDistributor, a...)
}

func (d *Six910Mysql) parseDistributorRow(foundRow *[]string) *mdb.Distributor {
	d.Log.Debug("foundRow in get Distributor", *foundRow)
	var rtn mdb.Distributor
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Distributor", err)
		if err == nil {
			sid, serr := strconv.ParseInt((*foundRow)[4], 10, 64)
			d.Log.Debug("sid err in get Distributor", serr)
			if serr == nil {
				rtn.ID = id
				rtn.StoreID = sid
				rtn.Company = (*foundRow)[1]
				rtn.ContactName = (*foundRow)[2]
				rtn.Phone = (*foundRow)[3]
			}
		}
	}
	return &rtn
}
