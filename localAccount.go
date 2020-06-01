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

//Local Accounts when oauth not used

//AddLocalAccount AddLocalAccount
func (d *Six910Mysql) AddLocalAccount(a *mdb.LocalAccount) bool {
	return false
}

//UpdateLocalAccount UpdateLocalAccount
func (d *Six910Mysql) UpdateLocalAccount(a *mdb.LocalAccount) bool {
	return false
}

//GetLocalAccount GetLocalAccount
func (d *Six910Mysql) GetLocalAccount(username string, storeID int64) *mdb.LocalAccount {
	return nil
}

//GetLocalAccountList GetLocalAccountList
func (d *Six910Mysql) GetLocalAccountList(store int64) *[]mdb.LocalAccount {
	return nil
}

//DeleteLocalAccount DeleteLocalAccount
func (d *Six910Mysql) DeleteLocalAccount(username string, storeID int64) bool {
	return false
}
