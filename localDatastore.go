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

//----UI Cluster installation: this is only called if UI is running in a cluster---
//Handle the situation where clients are running in a cluster
//This gives a way to make sure the json datastores are update on each node in the cluster

//----------------start datastore------------------------------------
//this gets called when a node start up and add only if it doesn't already exist

//AddLocalDatastore AddLocalDatastore
func (d *Six910Mysql) AddLocalDatastore(ds *mdb.LocalDataStore) bool {
	return false
}

//This get get called when a change is made to a datastore from a node in the cluster
//Or after all reloads have completed and then get set to Reload = false

//UpdateLocalDatastore UpdateLocalDatastore
func (d *Six910Mysql) UpdateLocalDatastore(ds *mdb.LocalDataStore) bool {
	return false
}

//This gets call by cluster nodes to see if there are pending reload

//GetLocalDatastore GetLocalDatastore
func (d *Six910Mysql) GetLocalDatastore(storeID int64, dataStoreName string) *mdb.LocalDataStore {
	return nil
}
