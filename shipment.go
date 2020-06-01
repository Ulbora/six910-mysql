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

//shipment

//AddShipment AddShipment
func (d *Six910Mysql) AddShipment(s *mdb.Shipment) (bool, int64) {
	return false, 0
}

//UpdateShipment UpdateShipment
func (d *Six910Mysql) UpdateShipment(s *mdb.Shipment) bool {
	return false
}

//GetShipment GetShipment
func (d *Six910Mysql) GetShipment(id int64) *mdb.Shipment {
	return nil
}

//GetShipmentList GetShipmentList
func (d *Six910Mysql) GetShipmentList(orderID int64) *[]mdb.Shipment {
	return nil
}

//DeleteShipment DeleteShipment
func (d *Six910Mysql) DeleteShipment(id int64) *mdb.Shipment {
	return nil
}
