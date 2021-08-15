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

//address

//AddAddress AddAddress
func (d *Six910Mysql) AddAddress(ad *mdb.Address) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ad.Address, ad.City, ad.State, ad.Zip, ad.County, ad.Country, ad.Type, ad.CustomerID,
		ad.Attr1, ad.Attr2, ad.Attr3, ad.Attr4)
	suc, id := d.DB.Insert(insertAddress, a...)
	d.Log.Debug("suc in add address", suc)
	d.Log.Debug("id in add address", id)
	return suc, id
}

//UpdateAddress UpdateAddress
func (d *Six910Mysql) UpdateAddress(ad *mdb.Address) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, ad.Address, ad.City, ad.State, ad.Zip, ad.County, ad.Country, ad.Type,
		ad.Attr1, ad.Attr2, ad.Attr3, ad.Attr4, ad.ID)
	suc := d.DB.Update(updateAddress, a...)
	return suc
}

//GetAddress GetAddress
func (d *Six910Mysql) GetAddress(id int64) *mdb.Address {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getAddress, a...)
	rtn := d.parseAddressRow(&row.Row)
	return rtn
}

//GetAddressList GetAddressList
func (d *Six910Mysql) GetAddressList(cid int64) *[]mdb.Address {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Address{}
	var a []interface{}
	a = append(a, cid)
	rows := d.DB.GetList(getAddressList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseAddressRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteAddress DeleteAddress
func (d *Six910Mysql) DeleteAddress(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteAddress, a...)
}

func (d *Six910Mysql) parseAddressRow(foundRow *[]string) *mdb.Address {
	d.Log.Debug("foundRow in get Address", *foundRow)
	var rtn mdb.Address
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get address", err)
		if err == nil {
			cid, cerr := strconv.ParseInt((*foundRow)[8], 10, 64)
			d.Log.Debug("cid err in get address", cerr)
			if cerr == nil {
				rtn.ID = id
				rtn.CustomerID = cid
				rtn.Address = (*foundRow)[1]
				rtn.City = (*foundRow)[2]
				rtn.State = (*foundRow)[3]
				rtn.Zip = (*foundRow)[4]
				rtn.County = (*foundRow)[5]
				rtn.Country = (*foundRow)[6]
				rtn.Type = (*foundRow)[7]
				rtn.Attr1 = (*foundRow)[9]
				rtn.Attr2 = (*foundRow)[10]
				rtn.Attr3 = (*foundRow)[11]
				rtn.Attr4 = (*foundRow)[12]
			}
		}
	}
	return &rtn
}
