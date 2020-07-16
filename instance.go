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

//---------------------start instance--------------------
// this gets called when each instance is started and added only if never added before
//The instance name is pulled from Docker or an manually entered env variable

//AddInstance AddInstance
func (d *Six910Mysql) AddInstance(i *mdb.Instances) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, i.InstanceName, i.ReloadDate, i.StoreID, i.DataStoreName)
	suc, id := d.DB.Insert(insertInstances, a...)
	d.Log.Debug("suc in add Instances", suc)
	d.Log.Debug("id in add Instances", id)
	return suc
}

//This gets called after instance gets reloaded

//UpdateInstance UpdateInstance
func (d *Six910Mysql) UpdateInstance(i *mdb.Instances) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, i.ReloadDate, i.InstanceName, i.StoreID, i.DataStoreName)
	suc := d.DB.Update(updateInstances, a...)
	return suc
}

//Gets called before updating an instance

//GetInstance GetInstance
func (d *Six910Mysql) GetInstance(name string, dataStoreName string, storeID int64) *mdb.Instances {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, name, storeID, dataStoreName)
	row := d.DB.Get(getInstances, a...)
	rtn := d.parseInstanceRow(&row.Row)
	return rtn
}

//GetInstanceList GetInstanceList
func (d *Six910Mysql) GetInstanceList(dataStoreName string, storeID int64) *[]mdb.Instances {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Instances{}
	var a []interface{}
	a = append(a, storeID, dataStoreName)
	rows := d.DB.GetList(getInstancesList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseInstanceRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *Six910Mysql) parseInstanceRow(foundRow *[]string) *mdb.Instances {
	d.Log.Debug("foundRow in get Instances", *foundRow)
	var rtn mdb.Instances
	if len(*foundRow) > 0 {
		sid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
		d.Log.Debug("sid err in get Instance", err)
		if err == nil {
			reltime, err := time.Parse(timeFormat, (*foundRow)[1])
			d.Log.Debug("reload time err in get Instance", err)
			if err == nil {
				rtn.StoreID = sid
				rtn.ReloadDate = reltime
				rtn.InstanceName = (*foundRow)[0]
				rtn.DataStoreName = (*foundRow)[3]
			}
		}
	}
	return &rtn
}
