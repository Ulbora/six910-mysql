package six910mysql

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

// go mod init github.com/Ulbora/six910-mysql

import (
	"fmt"
	"strconv"

	lg "github.com/Ulbora/Level_Logger"
	dbi "github.com/Ulbora/dbinterface"
	sdbi "github.com/Ulbora/six910-database-interface"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

//Six910Mysql Six910Mysql
type Six910Mysql struct {
	DB  dbi.Database
	Log *lg.Logger
}

//GetNew GetNew
func (d *Six910Mysql) GetNew() sdbi.Six910DB {
	return d
}

func (d *Six910Mysql) testConnection() bool {
	d.Log.Debug("in testConnection")
	var rtn = false
	var a []interface{}
	d.Log.Debug("d.DB: ", fmt.Sprintln(d.DB))
	rowPtr := d.DB.Test(cartTest, a...)
	d.Log.Debug("rowPtr", *rowPtr)
	d.Log.Debug("after testConnection test", *rowPtr)
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		//log.Print("Records found during test ")
		//log.Println("Records found during test :", int64Val)
		if err != nil {
			d.Log.Error(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}
