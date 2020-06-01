package six910mysql

import mdb "github.com/Ulbora/six910-database-interface"

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
func (d *Six910Mysql) AddDataStoreWriteLock(w *mdb.DataStoreWriteLock) (bool, int64) {
	return false, 0
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (d *Six910Mysql) UpdateDataStoreWriteLock(w *mdb.DataStoreWriteLock) bool {
	return false
}

//gets called from within the add method and by any node trying to update a datastore

//GetDataStoreWriteLock GetDataStoreWriteLock
func (d *Six910Mysql) GetDataStoreWriteLock(dataStore string, storeID int64) *mdb.DataStoreWriteLock {
	return nil
}
