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

//excluded sub regions

//AddExcludedSubRegion AddExcludedSubRegion
func (d *Six910Mysql) AddExcludedSubRegion(e *mdb.ExcludedSubRegion) (bool, int64) {
	return false, 0
	//add exclusion
}

//UpdateExcludedSubRegion UpdateExcludedSubRegion
func (d *Six910Mysql) UpdateExcludedSubRegion(e *mdb.ExcludedSubRegion) bool {
	return false
}

//GetExcludedSubRegion GetExcludedSubRegion
func (d *Six910Mysql) GetExcludedSubRegion(id int64) *mdb.ExcludedSubRegion {
	return nil
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (d *Six910Mysql) GetExcludedSubRegionList(regionID int64) *[]mdb.ExcludedSubRegion {
	return nil
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (d *Six910Mysql) DeleteExcludedSubRegion(id int64) bool {
	return false
}
