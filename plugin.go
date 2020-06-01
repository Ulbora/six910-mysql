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

//Global Plugins

//AddPlugin AddPlugin
func (d *Six910Mysql) AddPlugin(p *mdb.Plugins) (bool, int64) {
	return false, 0
}

//UpdatePlugin UpdatePlugin
func (d *Six910Mysql) UpdatePlugin(p *mdb.Plugins) bool {
	return false
}

//GetPlugin GetPlugin
func (d *Six910Mysql) GetPlugin(id int64) *mdb.Plugins {
	return nil
}

//GetPluginList GetPluginList
func (d *Six910Mysql) GetPluginList(start int64, end int64) *[]mdb.Plugins {
	return nil
}

//DeletePlugin DeletePlugin
func (d *Six910Mysql) DeletePlugin(id int64) bool {
	return false
}
