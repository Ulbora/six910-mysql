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

//shipment boxes

//AddShipmentBox AddShipmentBox
func (d *Six910Mysql) AddShipmentBox(sb *mdb.ShipmentBox) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, sb.BoxNumber, sb.ShippingMethodID, time.Now(), sb.ShippingAddressID,
		sb.ShippingAddress, sb.ShipmentID)
	suc, id := d.DB.Insert(insertShipmentBox, a...)
	d.Log.Debug("suc in add Shipment", suc)
	d.Log.Debug("id in add Shipment", id)
	return suc, id
}

//UpdateShipmentBox UpdateShipmentBox
func (d *Six910Mysql) UpdateShipmentBox(sb *mdb.ShipmentBox) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, sb.Dropship, sb.Cost, sb.Insurance, sb.Weight, sb.Width, sb.Height, sb.Depth,
		time.Now(), sb.TrackingNumber, sb.ID)
	suc := d.DB.Update(updateShipmentBox, a...)
	return suc
}

//GetShipmentBox GetShipmentBox
func (d *Six910Mysql) GetShipmentBox(id int64) *mdb.ShipmentBox {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getShipmentBox, a...)
	rtn := d.parseShipmentBoxRow(&row.Row)
	return rtn
}

//GetShipmentBoxList GetShipmentBoxList
func (d *Six910Mysql) GetShipmentBoxList(shipmentID int64) *[]mdb.ShipmentBox {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.ShipmentBox{}
	var a []interface{}
	a = append(a, shipmentID)
	rows := d.DB.GetList(getShipmentBoxList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseShipmentBoxRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteShipmentBox DeleteShipmentBox
func (d *Six910Mysql) DeleteShipmentBox(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteShipmentBox, a...)
}

func (d *Six910Mysql) parseShipmentBoxRow(foundRow *[]string) *mdb.ShipmentBox {
	d.Log.Debug("foundRow in get ShipmentBox", *foundRow)
	var rtn mdb.ShipmentBox
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get ShipmentBox", err)
		if err == nil {
			boxnum, oerr := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("boxnum err in get ShipmentBox", oerr)
			if oerr == nil {
				sTime, eerr := time.Parse(timeFormat, (*foundRow)[10])
				d.Log.Debug("eTime err in get ShipmentBox", eerr)
				if eerr == nil {
					uTime, _ := time.Parse(timeFormat, (*foundRow)[11])
					drops, enerr := strconv.ParseBool((*foundRow)[2])
					if enerr == nil {
						methID, err := strconv.ParseInt((*foundRow)[5], 10, 64)
						d.Log.Debug("methID err in get ShipmentBox", err)
						if err == nil {
							said, err := strconv.ParseInt((*foundRow)[13], 10, 64)
							d.Log.Debug("said err in get ShipmentBox", err)
							if err == nil {
								sid, err := strconv.ParseInt((*foundRow)[15], 10, 64)
								d.Log.Debug("sid err in get ShipmentBox", err)
								if err == nil {
									cost, err := strconv.ParseFloat((*foundRow)[3], 64)
									d.Log.Debug("cost err in get ShipmentBox", err)
									if err == nil {
										ins, err := strconv.ParseFloat((*foundRow)[4], 64)
										d.Log.Debug("ins err in get ShipmentBox", err)
										if err == nil {
											weight, err := strconv.ParseFloat((*foundRow)[6], 64)
											d.Log.Debug("weight err in get ShipmentBox", err)
											if err == nil {
												width, err := strconv.ParseFloat((*foundRow)[7], 64)
												d.Log.Debug("width err in get ShipmentBox", err)
												if err == nil {
													height, err := strconv.ParseFloat((*foundRow)[8], 64)
													d.Log.Debug("height err in get ShipmentBox", err)
													if err == nil {
														depth, err := strconv.ParseFloat((*foundRow)[9], 64)
														d.Log.Debug("depth err in get ShipmentBox", err)
														if err == nil {
															rtn.ID = id
															rtn.BoxNumber = boxnum
															rtn.ShipDate = sTime
															rtn.Updated = uTime
															rtn.Dropship = drops
															rtn.ShippingMethodID = methID
															rtn.ShippingAddressID = said
															rtn.ShipmentID = sid
															rtn.Cost = cost
															rtn.Insurance = ins
															rtn.Weight = weight
															rtn.Width = width
															rtn.Height = height
															rtn.Depth = depth
															rtn.TrackingNumber = (*foundRow)[12]
															rtn.ShippingAddress = (*foundRow)[14]
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
			}
		}
	}
	return &rtn
}
