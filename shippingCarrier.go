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

//store shipment carrier like UPS and FEDex

//AddShippingCarrier AddShippingCarrier
func (d *Six910Mysql) AddShippingCarrier(c *mdb.ShippingCarrier) (bool, int64) {
	return false, 0
}

//UpdateShippingCarrier UpdateShippingCarrier
func (d *Six910Mysql) UpdateShippingCarrier(c *mdb.ShippingCarrier) bool {
	return false
}

//GetShippingCarrierList GetShippingCarrierList
func (d *Six910Mysql) GetShippingCarrierList(storeID int64) *[]mdb.ShippingCarrier {
	return nil
}

//DeleteShippingCarrier DeleteShippingCarrier
func (d *Six910Mysql) DeleteShippingCarrier(id int64) bool {
	return false
}
