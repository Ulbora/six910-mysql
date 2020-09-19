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

//AddTaxRate AddTaxRate
func (d *Six910Mysql) AddTaxRate(t *mdb.TaxRate) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, t.Country, t.State, t.ZipStart, t.ZipEnd, t.PercentRate,
		t.ProductCategoryID, t.IncludeHandling, t.IncludeShipping, t.TaxType, t.StoreID)
	suc, id := d.DB.Insert(insertTaxRate, a...)
	d.Log.Debug("suc in add AddTaxRate", suc)
	d.Log.Debug("id in add AddTaxRate", id)
	return suc, id
}

//UpdateTaxRate UpdateTaxRate
func (d *Six910Mysql) UpdateTaxRate(t *mdb.TaxRate) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, t.ZipStart, t.ZipEnd, t.PercentRate,
		t.ProductCategoryID, t.IncludeHandling, t.IncludeShipping, t.TaxType, t.ID)
	suc := d.DB.Update(updateTaxRate, a...)
	return suc
}

//GetTaxRate GetTaxRate
func (d *Six910Mysql) GetTaxRate(country string, state string, storeID int64) *[]mdb.TaxRate {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.TaxRate{}
	var a []interface{}
	a = append(a, country, state, storeID)
	rows := d.DB.GetList(getTaxRate, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseTaxRateRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetTaxRateList GetTaxRateList
func (d *Six910Mysql) GetTaxRateList(storeID int64) *[]mdb.TaxRate {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.TaxRate{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getTaxRateList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseTaxRateRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteTaxRate DeleteTaxRate
func (d *Six910Mysql) DeleteTaxRate(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteTaxRate, a...)
}

func (d *Six910Mysql) parseTaxRateRow(foundRow *[]string) *mdb.TaxRate {
	d.Log.Debug("foundRow in get TaxRate", *foundRow)
	var rtn mdb.TaxRate
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get TaxRate", err)
		if err == nil {
			rate, err := strconv.ParseFloat((*foundRow)[5], 64)
			d.Log.Debug("rate err in get TaxRate", err)
			if err == nil {
				pid, err := strconv.ParseInt((*foundRow)[6], 10, 64)
				d.Log.Debug("pid err in get TaxRate", err)
				if err == nil {
					sid, err := strconv.ParseInt((*foundRow)[10], 10, 64)
					d.Log.Debug("qty err in get OrderItem", err)
					if err == nil {
						useHandling, err := strconv.ParseBool((*foundRow)[7])
						if err == nil {
							useShipping, err := strconv.ParseBool((*foundRow)[8])
							if err == nil {
								rtn.ID = id
								rtn.Country = (*foundRow)[1]
								rtn.State = (*foundRow)[2]
								rtn.ZipStart = (*foundRow)[3]
								rtn.ZipEnd = (*foundRow)[4]
								rtn.PercentRate = rate
								rtn.ProductCategoryID = pid
								rtn.IncludeHandling = useHandling
								rtn.IncludeShipping = useShipping
								rtn.TaxType = (*foundRow)[9]
								rtn.StoreID = sid
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
