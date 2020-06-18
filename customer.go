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

//customer

//AddCustomer AddCustomer
func (d *Six910Mysql) AddCustomer(c *mdb.Customer) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Email, c.ResetPassword, c.FirstName, c.LastName, c.Company, c.City,
		c.State, c.Zip, c.Phone, c.StoreID, time.Now())
	suc, id := d.DB.Insert(insertCustomer, a...)
	d.Log.Debug("suc in add customer", suc)
	d.Log.Debug("id in add customer", id)
	return suc, id
}

//UpdateCustomer UpdateCustomer
func (d *Six910Mysql) UpdateCustomer(c *mdb.Customer) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Email, c.ResetPassword, c.FirstName, c.LastName, c.Company, c.City,
		c.State, c.Zip, c.Phone, time.Now(), c.ID)
	suc := d.DB.Update(updateCustomer, a...)
	return suc
}

//GetCustomer GetCustomer
func (d *Six910Mysql) GetCustomer(email string, storeID int64) *mdb.Customer {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, email, storeID)
	row := d.DB.Get(getCustemerByEmail, a...)
	rtn := d.parseCustomerRow(&row.Row)
	return rtn
}

//GetCustomerList GetCustomerList
func (d *Six910Mysql) GetCustomerList(storeID int64) *[]mdb.Customer {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Customer{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getCustemerList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseCustomerRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetCustomerID GetCustomerID
func (d *Six910Mysql) GetCustomerID(id int64) *mdb.Customer {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getCustemerByID, a...)
	rtn := d.parseCustomerRow(&row.Row)
	return rtn
}

//DeleteCustomer DeleteCustomer
func (d *Six910Mysql) DeleteCustomer(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteCustomer, a...)
}

func (d *Six910Mysql) parseCustomerRow(foundRow *[]string) *mdb.Customer {
	d.Log.Debug("foundRow in get customer", *foundRow)
	var rtn mdb.Customer
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get customer", err)
		if err == nil {
			sid, serr := strconv.ParseInt((*foundRow)[10], 10, 64)
			d.Log.Debug("sid err in get customer", serr)
			if serr == nil {
				eTime, eerr := time.Parse(timeFormat, (*foundRow)[11])
				d.Log.Debug("eTime err in get customer", eerr)
				if eerr == nil {
					uTime, _ := time.Parse(timeFormat, (*foundRow)[12])
					rpass, enerr := strconv.ParseBool((*foundRow)[2])
					if enerr == nil {
						rtn.ID = id
						rtn.StoreID = sid
						rtn.DateEntered = eTime
						rtn.DateUpdated = uTime
						rtn.ResetPassword = rpass
						rtn.Email = (*foundRow)[1]
						rtn.FirstName = (*foundRow)[3]
						rtn.LastName = (*foundRow)[4]
						rtn.Company = (*foundRow)[5]
						rtn.City = (*foundRow)[6]
						rtn.State = (*foundRow)[7]
						rtn.Zip = (*foundRow)[8]
						rtn.Phone = (*foundRow)[9]
					}
				}
			}
		}
	}
	return &rtn
}
