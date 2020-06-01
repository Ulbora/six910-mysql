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

//Geographic Regions

//AddRegion AddRegion
func (d *Six910Mysql) AddRegion(r *mdb.Region) (bool, int64) {
	return false, 0
}

//UpdateRegion UpdateRegion
func (d *Six910Mysql) UpdateRegion(r *mdb.Region) bool {
	return false
}

//GetRegion GetRegion
func (d *Six910Mysql) GetRegion(id int64) *mdb.Region {
	return nil
}

//GetRegionList GetRegionList
func (d *Six910Mysql) GetRegionList(storeID int64) *[]mdb.Region {
	return nil
}

//DeleteRegion DeleteRegion
func (d *Six910Mysql) DeleteRegion(id int64) bool {
	return false
}
