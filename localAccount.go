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

//Local Accounts when oauth not used

//AddLocalAccount AddLocalAccount
func (d *Six910Mysql) AddLocalAccount(c *mdb.LocalAccount) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.UserName, c.Password, c.Enabled, c.Role, c.StoreID, c.CustomerID,
		time.Now())
	suc, _ := d.DB.Insert(insertLocalAccount, a...)
	d.Log.Debug("suc in add local account", suc)
	return suc
}

//UpdateLocalAccount UpdateLocalAccount
func (d *Six910Mysql) UpdateLocalAccount(c *mdb.LocalAccount) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, c.Password, c.Enabled, c.Role, time.Now(), c.UserName, c.StoreID)
	suc := d.DB.Update(updateLocalAccount, a...)
	return suc
}

//GetLocalAccount GetLocalAccount
func (d *Six910Mysql) GetLocalAccount(username string, storeID int64) *mdb.LocalAccount {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, username, storeID)
	row := d.DB.Get(getLocalAccount, a...)
	rtn := d.parseLocalAccountRow(&row.Row)
	return rtn
}

//GetLocalAccountList GetLocalAccountList
func (d *Six910Mysql) GetLocalAccountList(store int64) *[]mdb.LocalAccount {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.LocalAccount
	var a []interface{}
	a = append(a, store)
	rows := d.DB.GetList(getLocalAccountList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseLocalAccountRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteLocalAccount DeleteLocalAccount
func (d *Six910Mysql) DeleteLocalAccount(username string, storeID int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, username, storeID)
	return d.DB.Delete(deleteLocalAccount, a...)
}

func (d *Six910Mysql) parseLocalAccountRow(foundRow *[]string) *mdb.LocalAccount {
	d.Log.Debug("foundRow in get local account", *foundRow)
	var rtn mdb.LocalAccount
	if len(*foundRow) > 0 {
		sid, err := strconv.ParseInt((*foundRow)[4], 10, 64)
		d.Log.Debug("id err in get local account", err)
		if err == nil {
			cid, serr := strconv.ParseInt((*foundRow)[5], 10, 64)
			d.Log.Debug("id err in get local account", serr)
			if serr == nil {
				eTime, eerr := time.Parse(timeFormat, (*foundRow)[6])
				d.Log.Debug("eTime err in get local account", eerr)
				if eerr == nil {
					uTime, _ := time.Parse(timeFormat, (*foundRow)[7])
					enabled, enerr := strconv.ParseBool((*foundRow)[2])
					if enerr == nil {
						rtn.StoreID = sid
						rtn.CustomerID = cid
						rtn.Enabled = enabled
						rtn.DateEntered = eTime
						rtn.DateUpdated = uTime
						rtn.UserName = (*foundRow)[0]
						rtn.Password = (*foundRow)[1]
						rtn.Role = (*foundRow)[3]
					}
				}
			}
		}
	}
	return &rtn
}
