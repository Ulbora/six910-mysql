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

//cart item

//AddCartItem AddCartItem
func (d *Six910Mysql) AddCartItem(ci *mdb.CartItem) (bool, int64) {
	return false, 0
}

//UpdateCartItem UpdateCartItem
func (d *Six910Mysql) UpdateCartItem(ci *mdb.CartItem) bool {
	return false
}

//GetCarItem GetCarItem
func (d *Six910Mysql) GetCarItem(cartID int64, prodID int64) *mdb.CartItem {
	return nil
}

//GetCartItemList GetCartItemList
func (d *Six910Mysql) GetCartItemList(cartID int64) *[]mdb.CartItem {
	return nil
}

//DeleteCartItem DeleteCartItem
func (d *Six910Mysql) DeleteCartItem(id int64) bool {
	return false
}
