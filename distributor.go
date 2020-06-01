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

//distributors

//AddDistributor AddDistributor
func (d *Six910Mysql) AddDistributor(ds *mdb.Distributor) (bool, int64) {
	return false, 0
}

//UpdateDistributor UpdateDistributor
func (d *Six910Mysql) UpdateDistributor(ds *mdb.Distributor) bool {
	return false
}

//GetDistributor GetDistributor
func (d *Six910Mysql) GetDistributor(id int64) *mdb.Distributor {
	return nil
}

//GetDistributorList GetDistributorList
func (d *Six910Mysql) GetDistributorList(store int64) *[]mdb.Distributor {
	return nil
}

//DeleteDistributor DeleteDistributor
func (d *Six910Mysql) DeleteDistributor(id int64) bool {
	return false
}
