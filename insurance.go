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

//shipping insurance

//AddInsurance AddInsurance
func (d *Six910Mysql) AddInsurance(i *mdb.Insurance) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, i.Cost, i.MinOrderAmount, i.MaxOrderAmount, i.StoreID)
	suc, id := d.DB.Insert(insertInsurance, a...)
	d.Log.Debug("suc in add Insurance", suc)
	d.Log.Debug("id in add Insurance", id)
	return suc, id
}

//UpdateInsurance UpdateInsurance
func (d *Six910Mysql) UpdateInsurance(i *mdb.Insurance) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, i.Cost, i.MinOrderAmount, i.MaxOrderAmount, i.ID)
	suc := d.DB.Update(updateInsurance, a...)
	return suc
}

//GetInsurance GetInsurance
func (d *Six910Mysql) GetInsurance(id int64) *mdb.Insurance {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getInsurance, a...)
	rtn := d.parseInsuranceRow(&row.Row)
	return rtn
}

//GetInsuranceList GetInsuranceList
func (d *Six910Mysql) GetInsuranceList(storeID int64) *[]mdb.Insurance {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Insurance{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getInsuranceList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseInsuranceRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteInsurance DeleteInsurance
func (d *Six910Mysql) DeleteInsurance(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteInsurance, a...)
}

func (d *Six910Mysql) parseInsuranceRow(foundRow *[]string) *mdb.Insurance {
	d.Log.Debug("foundRow in get Insurance", *foundRow)
	var rtn mdb.Insurance
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Insurance", err)
		if err == nil {
			sid, cerr := strconv.ParseInt((*foundRow)[4], 10, 64)
			d.Log.Debug("sid err in get Insurance", cerr)
			if cerr == nil {
				cost, err := strconv.ParseFloat((*foundRow)[1], 64)
				d.Log.Debug("cost err in get Insurance", err)
				if err == nil {
					min, err := strconv.ParseFloat((*foundRow)[2], 64)
					d.Log.Debug("min err in get Insurance", err)
					if err == nil {
						max, err := strconv.ParseFloat((*foundRow)[3], 64)
						d.Log.Debug("max err in get Insurance", err)
						if err == nil {
							rtn.ID = id
							rtn.StoreID = sid
							rtn.Cost = cost
							rtn.MinOrderAmount = min
							rtn.MaxOrderAmount = max
						}
					}
				}
			}
		}
	}
	return &rtn
}
