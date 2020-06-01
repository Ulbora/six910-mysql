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
//stores

//AddStore AddStore
func (d *Six910Mysql) AddStore(s *mdb.Store) (bool, int64) {
	return false, 0
}

//UpdateStore UpdateStore
func (d *Six910Mysql) UpdateStore(s *mdb.Store) bool {
	return false
}

//GetStore GetStore
func (d *Six910Mysql) GetStore(sname string) *mdb.Store {
	return nil
}

//GetStoreID GetStoreID
func (d *Six910Mysql) GetStoreID(id int64) *mdb.Store {
	return nil
}

//GetStoreByLocal GetStoreByLocal
func (d *Six910Mysql) GetStoreByLocal(localDomain string) *mdb.Store {
	return nil
}

//DeleteStore DeleteStore
func (d *Six910Mysql) DeleteStore(id int64) bool {
	return false
}
