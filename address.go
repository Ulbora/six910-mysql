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

//address

//AddAddress AddAddress
func (d *Six910Mysql) AddAddress(a *mdb.Address) (bool, int64) {
	return false, 0
}

//UpdateAddress UpdateAddress
func (d *Six910Mysql) UpdateAddress(a *mdb.Address) bool {
	return false
}

//GetAddress GetAddress
func (d *Six910Mysql) GetAddress(id int64) *mdb.Address {
	return nil
}

//GetAddressList GetAddressList
func (d *Six910Mysql) GetAddressList(cid int64) *[]mdb.Address {
	return nil
}

//DeleteAddress DeleteAddress
func (d *Six910Mysql) DeleteAddress(id int64) bool {
	return false
}
