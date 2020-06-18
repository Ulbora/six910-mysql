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

//shipping method

//AddShippingMethod AddShippingMethod
func (d *Six910Mysql) AddShippingMethod(s *mdb.ShippingMethod) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Name, s.Cost, s.MaxWeight, s.Handling, s.MinOrderAmount, s.MaxOrderAmount,
		s.RegionID, s.ShippingCarrierID, s.InsuranceID, s.StoreID)
	suc, id := d.DB.Insert(insertShippingMethod, a...)
	d.Log.Debug("suc in add ShippingMethod", suc)
	d.Log.Debug("id in add ShippingMethod", id)
	return suc, id
}

//UpdateShippingMethod UpdateShippingMethod
func (d *Six910Mysql) UpdateShippingMethod(s *mdb.ShippingMethod) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Name, s.Cost, s.MaxWeight, s.Handling, s.MinOrderAmount, s.MaxOrderAmount, s.ID)
	suc := d.DB.Update(updateShippingMethod, a...)
	return suc
}

//GetShippingMethod GetShippingMethod
func (d *Six910Mysql) GetShippingMethod(id int64) *mdb.ShippingMethod {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getShippingMethod, a...)
	rtn := d.parseShippingMethodRow(&row.Row)
	return rtn
}

//GetShippingMethodList GetShippingMethodList
func (d *Six910Mysql) GetShippingMethodList(storeID int64) *[]mdb.ShippingMethod {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ShippingMethod{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getShippingMethodList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseShippingMethodRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteShippingMethod DeleteShippingMethod
func (d *Six910Mysql) DeleteShippingMethod(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteShippingMethod, a...)
}

func (d *Six910Mysql) parseShippingMethodRow(foundRow *[]string) *mdb.ShippingMethod {
	d.Log.Debug("foundRow in get ShippingMethod", *foundRow)
	var rtn mdb.ShippingMethod
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get ShippingMethod", err)
		if err == nil {
			rid, cerr := strconv.ParseInt((*foundRow)[7], 10, 64)
			d.Log.Debug("rid err in get ShippingMethod", cerr)
			if cerr == nil {
				cost, err := strconv.ParseFloat((*foundRow)[2], 64)
				d.Log.Debug("cost err in get ShippingMethod", err)
				if err == nil {
					min, err := strconv.ParseFloat((*foundRow)[5], 64)
					d.Log.Debug("min err in get ShippingMethod", err)
					if err == nil {
						max, err := strconv.ParseFloat((*foundRow)[6], 64)
						d.Log.Debug("max err in get ShippingMethod", err)
						if err == nil {
							hand, err := strconv.ParseFloat((*foundRow)[4], 64)
							d.Log.Debug("hand err in get ShippingMethod", err)
							if err == nil {
								maxw, err := strconv.ParseInt((*foundRow)[3], 10, 64)
								d.Log.Debug("maxw err in get ShippingMethod", err)
								if err == nil {
									scid, err := strconv.ParseInt((*foundRow)[8], 10, 64)
									d.Log.Debug("scid err in get ShippingMethod", err)
									if err == nil {
										iid, err := strconv.ParseInt((*foundRow)[9], 10, 64)
										d.Log.Debug("iid err in get ShippingMethod", err)
										if err == nil {
											sid, err := strconv.ParseInt((*foundRow)[10], 10, 64)
											d.Log.Debug("sid err in get ShippingMethod", err)
											if err == nil {
												rtn.ID = id
												rtn.Name = (*foundRow)[1]
												rtn.Cost = cost
												rtn.Handling = hand
												rtn.InsuranceID = iid
												rtn.MaxOrderAmount = max
												rtn.MinOrderAmount = min
												rtn.MaxWeight = maxw
												rtn.RegionID = rid
												rtn.ShippingCarrierID = scid
												rtn.StoreID = sid
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
