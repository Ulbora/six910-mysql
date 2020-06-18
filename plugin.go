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

//Global Plugins

//AddPlugin AddPlugin
func (d *Six910Mysql) AddPlugin(p *mdb.Plugins) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, p.PluginName, p.Developer, p.ContactPhone, p.DeveloperAddress, p.Fee, p.Category,
		p.ActivateURL, p.OauthRedirectURL, p.IsPGW, p.Enabled)
	suc, id := d.DB.Insert(insertPlugin, a...)
	d.Log.Debug("suc in add Plugins", suc)
	d.Log.Debug("id in add Plugins", id)
	return suc, id
}

//UpdatePlugin UpdatePlugin
func (d *Six910Mysql) UpdatePlugin(p *mdb.Plugins) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, p.PluginName, p.Developer, p.ContactPhone, p.DeveloperAddress, p.Fee, p.Category,
		p.ActivateURL, p.OauthRedirectURL, p.IsPGW, p.Enabled, p.ID)
	suc := d.DB.Update(updatePlugin, a...)
	return suc
}

//GetPlugin GetPlugin
func (d *Six910Mysql) GetPlugin(id int64) *mdb.Plugins {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getPlugin, a...)
	rtn := d.parsePluginRow(&row.Row)
	return rtn
}

//GetPluginList GetPluginList
func (d *Six910Mysql) GetPluginList(start int64, end int64) *[]mdb.Plugins {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Plugins{}
	var a []interface{}
	a = append(a, start, end)
	rows := d.DB.GetList(getPluginList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parsePluginRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeletePlugin DeletePlugin
func (d *Six910Mysql) DeletePlugin(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deletePlugin, a...)
}

func (d *Six910Mysql) parsePluginRow(foundRow *[]string) *mdb.Plugins {
	d.Log.Debug("foundRow in get Plugins", *foundRow)
	var rtn mdb.Plugins
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Plugins", err)
		if err == nil {
			fee, err := strconv.ParseFloat((*foundRow)[5], 64)
			d.Log.Debug("fee err in get Plugins", err)
			if err == nil {
				pgw, err := strconv.ParseBool((*foundRow)[9])
				d.Log.Debug("pgw err in get Plugins", err)
				if err == nil {
					enabled, err := strconv.ParseBool((*foundRow)[10])
					d.Log.Debug("enabled err in get Plugins", err)
					if err == nil {
						rtn.ID = id
						rtn.Fee = fee
						rtn.IsPGW = pgw
						rtn.Enabled = enabled
						rtn.PluginName = (*foundRow)[1]
						rtn.Developer = (*foundRow)[2]
						rtn.ContactPhone = (*foundRow)[3]
						rtn.DeveloperAddress = (*foundRow)[4]
						rtn.Category = (*foundRow)[6]
						rtn.ActivateURL = (*foundRow)[7]
						rtn.OauthRedirectURL = (*foundRow)[8]
					}
				}
			}
		}
	}
	return &rtn
}
