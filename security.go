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

//AddSecurity AddSecurity
func (d *Six910Mysql) AddSecurity(s *mdb.Security) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.OauthOn)
	suc, id := d.DB.Insert(insertSecurity, a...)
	d.Log.Debug("suc in add", suc)
	d.Log.Debug("id in add", id)
	return suc, id
}

//UpdateSecurity UpdateSecurity
func (d *Six910Mysql) UpdateSecurity(s *mdb.Security) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.OauthOn, s.ID)
	suc := d.DB.Update(updateSecurity, a...)
	return suc
}

//GetSecurity GetSecurity
func (d *Six910Mysql) GetSecurity() *mdb.Security {
	var rtn mdb.Security
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	rows := d.DB.GetList(getSecurity, a...)
	d.Log.Debug("rows", rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			d.Log.Debug("foundRow", foundRow)
			if len(foundRow) > 0 {
				id, err := strconv.ParseInt((foundRow)[0], 10, 64)
				if err == nil {
					rtn.OauthOn, _ = strconv.ParseBool((foundRow)[1])
					if err == nil {
						rtn.ID = id
					}
				}
			}
			break
		}
	}
	return &rtn
}

//DeleteSecurity DeleteSecurity
func (d *Six910Mysql) DeleteSecurity() bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	//a = append(a, id)
	return d.DB.Delete(deleteSecurity, a...)
}
