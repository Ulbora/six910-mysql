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

//product

//AddProduct AddProduct
func (d *Six910Mysql) AddProduct(p *mdb.Product) (bool, int64) {
	return false, 0
}

//UpdateProduct UpdateProduct
func (d *Six910Mysql) UpdateProduct(p *mdb.Product) bool {
	return false
}

//GetProductByID GetProductByID
func (d *Six910Mysql) GetProductByID(id int64) *mdb.Product {
	return nil
}

//GetProductsByName GetProductsByName
func (d *Six910Mysql) GetProductsByName(name string, start int64, end int64) *[]mdb.Product {
	return nil
}

//GetProductsByCaterory GetProductsByCaterory
func (d *Six910Mysql) GetProductsByCaterory(catID int64, start int64, end int64) *[]mdb.Product {
	return nil
}

//GetProductList GetProductList
func (d *Six910Mysql) GetProductList(storeID int64, start int64, end int64) *[]mdb.Product {
	return nil
}

//DeleteProduct DeleteProduct
func (d *Six910Mysql) DeleteProduct(id int64) bool {
	return false
}
