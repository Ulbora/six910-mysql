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

//store shipment carrier like UPS and FEDex

//AddShippingCarrier AddShippingCarrier
func (d *Six910Mysql) AddShippingCarrier(c *mdb.ShippingCarrier) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Carrier, c.Type, c.StoreID)
	suc, id := d.DB.Insert(insertShippingCarrier, a...)
	d.Log.Debug("suc in add ShippingCarrier", suc)
	d.Log.Debug("id in add ShippingCarrier", id)
	return suc, id
}

//UpdateShippingCarrier UpdateShippingCarrier
func (d *Six910Mysql) UpdateShippingCarrier(c *mdb.ShippingCarrier) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Carrier, c.Type, c.ID)
	suc := d.DB.Update(updateShippingCarrier, a...)
	return suc
}

//GetShippingCarrierList GetShippingCarrierList
func (d *Six910Mysql) GetShippingCarrierList(storeID int64) *[]mdb.ShippingCarrier {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.ShippingCarrier
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getShippingCarrierList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseShippingCarrierRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteShippingCarrier DeleteShippingCarrier
func (d *Six910Mysql) DeleteShippingCarrier(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteShippingCarrier, a...)
}

func (d *Six910Mysql) parseShippingCarrierRow(foundRow *[]string) *mdb.ShippingCarrier {
	d.Log.Debug("foundRow in get Region", *foundRow)
	var rtn mdb.ShippingCarrier
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Region", err)
		if err == nil {
			sid, cerr := strconv.ParseInt((*foundRow)[3], 10, 64)
			d.Log.Debug("qty err in get Region", cerr)
			if cerr == nil {
				rtn.ID = id
				rtn.StoreID = sid
				rtn.Carrier = (*foundRow)[1]
				rtn.Type = (*foundRow)[2]
			}
		}
	}
	return &rtn
}
