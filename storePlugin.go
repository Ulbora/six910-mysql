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

//store plugins installed

//AddStorePlugin AddStorePlugin
func (d *Six910Mysql) AddStorePlugin(sp *mdb.StorePlugins) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, sp.PluginID, sp.PluginName, sp.Category, sp.Active, sp.OauthClientID, sp.OauthSecret,
		sp.ActivateURL, sp.OauthRedirectURL, sp.APIKey, sp.RekeyTryCount, sp.IframeURL,
		sp.MenuTitle, sp.MenuIconURL, sp.IsPGW, sp.StoreID)
	suc, id := d.DB.Insert(insertStorePlugin, a...)
	d.Log.Debug("suc in add StorePlugins", suc)
	d.Log.Debug("id in add StorePlugins", id)
	return suc, id
}

//UpdateStorePlugin UpdateStorePlugin
func (d *Six910Mysql) UpdateStorePlugin(sp *mdb.StorePlugins) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	if sp.RekeyDate.IsZero() {
		a = append(a, sp.Active, sp.OauthClientID, sp.OauthSecret,
			sp.ActivateURL, sp.OauthRedirectURL, sp.APIKey, sp.RekeyTryCount, sp.IframeURL,
			sp.MenuTitle, sp.MenuIconURL, sp.ID)
		suc = d.DB.Update(updateStorePluginNoDate, a...)
	} else {
		a = append(a, sp.Active, sp.OauthClientID, sp.OauthSecret,
			sp.ActivateURL, sp.OauthRedirectURL, sp.APIKey, sp.RekeyTryCount, sp.IframeURL,
			sp.MenuTitle, sp.MenuIconURL, sp.RekeyDate, sp.ID)
		suc = d.DB.Update(updateStorePluginDate, a...)
	}

	return suc
}

//GetStorePlugin GetStorePlugin
func (d *Six910Mysql) GetStorePlugin(id int64) *mdb.StorePlugins {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getStorePlugin, a...)
	rtn := d.parseStorePluginRow(&row.Row)
	return rtn
}

//GetStorePluginList GetStorePluginList
func (d *Six910Mysql) GetStorePluginList(storeID int64) *[]mdb.StorePlugins {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.StorePlugins
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getStorePluginList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseStorePluginRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteStorePlugin DeleteStorePlugin
func (d *Six910Mysql) DeleteStorePlugin(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteStorePlugin, a...)
}

func (d *Six910Mysql) parseStorePluginRow(foundRow *[]string) *mdb.StorePlugins {
	d.Log.Debug("foundRow in get StorePlugins", *foundRow)
	var rtn mdb.StorePlugins
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get StorePlugins", err)
		if err == nil {
			pid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("cost err in get StorePlugins", err)
			if err == nil {
				active, err := strconv.ParseBool((*foundRow)[4])
				d.Log.Debug("min err in get StorePlugins", err)
				if err == nil {
					pgw, err := strconv.ParseBool((*foundRow)[15])
					d.Log.Debug("max err in get StorePlugins", err)
					if err == nil {
						oid, err := strconv.ParseInt((*foundRow)[5], 10, 64)
						d.Log.Debug("cost err in get StorePlugins", err)
						if err == nil {
							rtcnt, err := strconv.ParseInt((*foundRow)[10], 10, 64)
							d.Log.Debug("cost err in get StorePlugins", err)
							if err == nil {
								rtTime, err := time.Parse(timeFormat, (*foundRow)[11])
								d.Log.Debug("eTime err in get StorePlugins", err)
								if err == nil {
									sid, err := strconv.ParseInt((*foundRow)[16], 10, 64)
									d.Log.Debug("cost err in get StorePlugins", err)
									if err == nil {
										rtn.ID = id
										rtn.PluginID = pid
										rtn.Active = active
										rtn.IsPGW = pgw
										rtn.OauthClientID = oid
										rtn.RekeyTryCount = rtcnt
										rtn.RekeyDate = rtTime
										rtn.StoreID = sid
										rtn.PluginName = (*foundRow)[2]
										rtn.Category = (*foundRow)[3]
										rtn.OauthSecret = (*foundRow)[6]
										rtn.ActivateURL = (*foundRow)[7]
										rtn.OauthRedirectURL = (*foundRow)[8]
										rtn.APIKey = (*foundRow)[9]
										rtn.IframeURL = (*foundRow)[12]
										rtn.MenuTitle = (*foundRow)[13]
										rtn.MenuIconURL = (*foundRow)[14]
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
