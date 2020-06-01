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

//---------------------start instance--------------------
// this gets called when each instance is started and added only if never added before
//The instance name is pulled from Docker or an manually entered env variable

//AddInstance AddInstance
func (d *Six910Mysql) AddInstance(i *mdb.Instances) bool {
	return false
}

//This gets called after instance gets reloaded

//UpdateInstance UpdateInstance
func (d *Six910Mysql) UpdateInstance(i *mdb.Instances) bool {
	return false
}

//Gets called before updating an instance

//GetInstance GetInstance
func (d *Six910Mysql) GetInstance(name string, dataStoreName string, storeID int64) *mdb.Instances {
	return nil
}
