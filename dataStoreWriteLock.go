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

//-------------------start write lock-------------
//gets called after UI makes changes to one of the datastores
//If the datastore already exists, the Update method is called from within add

//AddDataStoreWriteLock AddDataStoreWriteLock
func (d *Six910Mysql) AddDataStoreWriteLock(w *mdb.DataStoreWriteLock) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, w.DataStoreName, w.Locked, w.LockedInstanceName, w.LockedByUser, w.LockedTime,
		w.StoreID)
	suc, id := d.DB.Insert(insertDataStoreWriteLock, a...)
	d.Log.Debug("suc in add DataStoreWriteLock", suc)
	d.Log.Debug("id in add DataStoreWriteLock", id)
	return suc
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (d *Six910Mysql) UpdateDataStoreWriteLock(w *mdb.DataStoreWriteLock) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, w.Locked, w.LockedInstanceName, w.LockedByUser, w.LockedTime,
		w.DataStoreName, w.StoreID)
	suc := d.DB.Update(updateDataStoreWriteLock, a...)
	return suc
}

//gets called from within the add method and by any node trying to update a datastore

//GetDataStoreWriteLock GetDataStoreWriteLock
func (d *Six910Mysql) GetDataStoreWriteLock(dataStore string, storeID int64) *mdb.DataStoreWriteLock {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, dataStore, storeID)
	row := d.DB.Get(getDataStoreWriteLock, a...)
	rtn := d.parseDataStoreWriteLockRow(&row.Row)
	return rtn
}

func (d *Six910Mysql) parseDataStoreWriteLockRow(foundRow *[]string) *mdb.DataStoreWriteLock {
	d.Log.Debug("foundRow in get DataStoreWriteLock", *foundRow)
	var rtn mdb.DataStoreWriteLock
	if len(*foundRow) > 0 {
		loctime, err := time.Parse(timeFormat, (*foundRow)[4])
		d.Log.Debug("reload time err in get DataStoreWriteLock", err)
		if err == nil {
			locked, err := strconv.ParseBool((*foundRow)[1])
			if err == nil {
				sid, err := strconv.ParseInt((*foundRow)[5], 10, 64)
				d.Log.Debug("sid err in get DataStoreWriteLock", err)
				if err == nil {
				}
				rtn.StoreID = sid
				rtn.LockedTime = loctime
				rtn.Locked = locked
				rtn.DataStoreName = (*foundRow)[0]
				rtn.LockedInstanceName = (*foundRow)[2]
				rtn.LockedByUser = (*foundRow)[3]
			}
		}
	}
	return &rtn
}
