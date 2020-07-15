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

//Shipment Items in box

//AddShipmentItem AddShipmentItem
func (d *Six910Mysql) AddShipmentItem(si *mdb.ShipmentItem) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, si.OrderItemID, si.Quantity, si.ShipmentID, time.Now(), si.ShipmentBoxID)
	suc, id := d.DB.Insert(insertShipmentItem, a...)
	d.Log.Debug("suc in add ShipmentItem", suc)
	d.Log.Debug("id in add ShipmentItem", id)
	return suc, id
}

//UpdateShipmentItem UpdateShipmentItem
func (d *Six910Mysql) UpdateShipmentItem(si *mdb.ShipmentItem) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, si.Quantity, time.Now(), si.ID)
	suc := d.DB.Update(updateShipmentItem, a...)
	return suc
}

//GetShipmentItem GetShipmentItem
func (d *Six910Mysql) GetShipmentItem(id int64) *mdb.ShipmentItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getShipmentItem, a...)
	rtn := d.parseShipmentItemRow(&row.Row)
	return rtn
}

//GetShipmentItemList GetShipmentItemList
func (d *Six910Mysql) GetShipmentItemList(shipmentID int64) *[]mdb.ShipmentItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ShipmentItem{}
	var a []interface{}
	a = append(a, shipmentID)
	rows := d.DB.GetList(getShipmentItemListByShipment, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseShipmentItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (d *Six910Mysql) GetShipmentItemListByBox(boxNumber int64, shipmentID int64) *[]mdb.ShipmentItem {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ShipmentItem{}
	var a []interface{}
	a = append(a, boxNumber, shipmentID)
	rows := d.DB.GetList(getShipmentItemsInBox, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseShipmentItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteShipmentItem DeleteShipmentItem
func (d *Six910Mysql) DeleteShipmentItem(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteShipmentItem, a...)
}

func (d *Six910Mysql) parseShipmentItemRow(foundRow *[]string) *mdb.ShipmentItem {
	d.Log.Debug("foundRow in get ShipmentItem", *foundRow)
	var rtn mdb.ShipmentItem
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get ShipmentItem", err)
		if err == nil {
			oiid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("oiid err in get ShipmentItem", err)
			if err == nil {
				qty, err := strconv.ParseInt((*foundRow)[2], 10, 64)
				d.Log.Debug("qty err in get ShipmentItem", err)
				if err == nil {
					sid, err := strconv.ParseInt((*foundRow)[3], 10, 64)
					d.Log.Debug("sid err in get ShipmentItem", err)
					if err == nil {
						sbid, enerr := strconv.ParseInt((*foundRow)[5], 10, 64)
						if enerr == nil {
							eTime, err := time.Parse(timeFormat, (*foundRow)[4])
							d.Log.Debug("eTime err in get ShipmentItem", err)
							if err == nil {
								rtn.ID = id
								rtn.OrderItemID = oiid
								rtn.Quantity = qty
								rtn.ShipmentBoxID = sbid
								rtn.ShipmentID = sid
								rtn.Updated = eTime
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
