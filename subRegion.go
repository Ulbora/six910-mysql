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

//Geographic Sub Regions

//AddSubRegion AddSubRegion
func (d *Six910Mysql) AddSubRegion(s *mdb.SubRegion) (bool, int64) {
	return false, 0
}

//UpdateSubRegion UpdateSubRegion
func (d *Six910Mysql) UpdateSubRegion(s *mdb.SubRegion) bool {
	return false
}

//GetSubRegion GetSubRegion
func (d *Six910Mysql) GetSubRegion(id int64) *mdb.SubRegion {
	return nil
}

//GetSubRegionList GetSubRegionList
func (d *Six910Mysql) GetSubRegionList(regionID int64) *[]mdb.SubRegion {
	return nil
}

//DeleteSubRegion DeleteSubRegion
func (d *Six910Mysql) DeleteSubRegion(id int64) bool {
	return false
}
