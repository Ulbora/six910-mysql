package six910mysql

import (
	"strconv"
	"time"

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

//shipment

//AddShipment AddShipment
func (d *Six910Mysql) AddShipment(s *mdb.Shipment) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.OrderID, time.Now(), s.Status, s.Boxes, s.ShippingHandling, s.Insurance)
	suc, id := d.DB.Insert(insertShipment, a...)
	d.Log.Debug("suc in add Shipment", suc)
	d.Log.Debug("id in add Shipment", id)
	return suc, id
}

//UpdateShipment UpdateShipment
func (d *Six910Mysql) UpdateShipment(s *mdb.Shipment) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Status, s.Boxes, s.ShippingHandling, s.Insurance, time.Now(), s.ID)
	suc := d.DB.Update(updateShipment, a...)
	return suc
}

//GetShipment GetShipment
func (d *Six910Mysql) GetShipment(id int64) *mdb.Shipment {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getShipment, a...)
	rtn := d.parseShipmentRow(&row.Row)
	return rtn
}

//GetShipmentList GetShipmentList
func (d *Six910Mysql) GetShipmentList(orderID int64) *[]mdb.Shipment {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.Shipment
	var a []interface{}
	a = append(a, orderID)
	rows := d.DB.GetList(getShipmentList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseShipmentRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteShipment DeleteShipment
func (d *Six910Mysql) DeleteShipment(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteShipment, a...)
}

func (d *Six910Mysql) parseShipmentRow(foundRow *[]string) *mdb.Shipment {
	d.Log.Debug("foundRow in get Shipment", *foundRow)
	var rtn mdb.Shipment
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Shipment", err)
		if err == nil {
			oid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("id err in get Shipment", err)
			if err == nil {
				boxes, err := strconv.ParseInt((*foundRow)[4], 10, 64)
				d.Log.Debug("id err in get Shipment", err)
				if err == nil {
					sanh, err := strconv.ParseFloat((*foundRow)[5], 64)
					d.Log.Debug("id err in get Shipment", err)
					if err == nil {
						ins, enerr := strconv.ParseFloat((*foundRow)[6], 64)
						if enerr == nil {
							eTime, err := time.Parse(timeFormat, (*foundRow)[2])
							d.Log.Debug("eTime err in get Shipment", err)
							if err == nil {
								uTime, _ := time.Parse(timeFormat, (*foundRow)[7])
								rtn.ID = id
								rtn.OrderID = oid
								rtn.Boxes = boxes
								rtn.CreateDate = eTime
								rtn.Insurance = ins
								rtn.ShippingHandling = sanh
								rtn.Updated = uTime
								rtn.Status = (*foundRow)[3]
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
