package six910mysql

import (
	"strconv"
	"time"

	mdb "github.com/Ulbora/six910-database-interface"
)

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
//stores

//AddStore AddStore
func (d *Six910Mysql) AddStore(s *mdb.Store) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Company, s.FirstName, s.LastName, s.LocalDomain, s.RemoteDomain, s.OauthClientID,
		s.OauthSecret, s.Email, s.City, s.State, s.Zip, time.Now(), s.StoreName, s.StoreSlogan, s.Logo, s.Currency, s.Enabled)
	suc, id := d.DB.Insert(insertStore, a...)
	d.Log.Debug("suc in add store", suc)
	d.Log.Debug("id in add store", id)
	return suc, id
}

//UpdateStore UpdateStore
func (d *Six910Mysql) UpdateStore(s *mdb.Store) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, s.Company, s.FirstName, s.LastName, s.LocalDomain, s.RemoteDomain,
		s.OauthClientID, s.OauthSecret, s.Email, s.City, s.State, s.Zip, s.StoreName, s.StoreSlogan, s.Logo, s.Currency, time.Now(), s.Enabled, s.ID)
	suc := d.DB.Update(updateStore, a...)
	return suc
}

//GetStore GetStore
func (d *Six910Mysql) GetStore(sname string) *mdb.Store {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, sname)
	row := d.DB.Get(getStoreByName, a...)
	rtn := d.parseStoreRow(&row.Row)
	return rtn
}

//GetLocalStore GetLocalStore
func (d *Six910Mysql) GetLocalStore() *mdb.Store {
	var rtn *mdb.Store
	if !d.testConnection() {
		d.DB.Connect()
	}
	if !d.GetSecurity().OauthOn {
		var a []interface{}
		//a = append(a, sname)
		rows := d.DB.GetList(getAllStores, a...)
		if rows != nil && len(rows.Rows) == 1 {
			foundRows := rows.Rows
			for r := range foundRows {
				foundRow := foundRows[r]
				rtn = d.parseStoreRow(&foundRow)
				break
				//rowContent := d.parseShippingMethodRow(&foundRow)
				//rtn = append(rtn, *rowContent)
			}
		}
		//rtn = d.parseStoreRow(&row.Row)
	}
	return rtn
}

//GetStoreID GetStoreID
func (d *Six910Mysql) GetStoreID(id int64) *mdb.Store {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getStoreByID, a...)
	rtn := d.parseStoreRow(&row.Row)
	return rtn
}

//GetStoreByLocal GetStoreByLocal
func (d *Six910Mysql) GetStoreByLocal(localDomain string) *mdb.Store {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, localDomain)
	row := d.DB.Get(getStoreByLocalDomain, a...)
	rtn := d.parseStoreRow(&row.Row)
	return rtn
}

//GetStoreCount GetStoreCount
func (d *Six910Mysql) GetStoreCount() int64 {
	var rtn int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	row := d.DB.Get(getStoreCount, a...)
	if len(row.Row) > 0 {
		cnt, err := strconv.ParseInt((row.Row)[0], 10, 64)
		d.Log.Debug("store count err", err)
		if err == nil {
			rtn = cnt
		}
	}
	return rtn
}

//DeleteStore DeleteStore
func (d *Six910Mysql) DeleteStore(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteStore, a...)
}

func (d *Six910Mysql) parseStoreRow(foundRow *[]string) *mdb.Store {
	d.Log.Debug("foundRow in get store", *foundRow)
	var rtn mdb.Store
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get store", err)
		if err == nil {
			oauthID, oerr := strconv.ParseInt((*foundRow)[6], 10, 64)
			d.Log.Debug("oauthID err in get store", oerr)
			if oerr == nil {
				eTime, eerr := time.Parse(timeFormat, (*foundRow)[12])
				d.Log.Debug("eTime err in get store", eerr)
				if eerr == nil {
					uTime, _ := time.Parse(timeFormat, (*foundRow)[13])
					enabled, enerr := strconv.ParseBool((*foundRow)[18])
					if enerr == nil {
						rtn.ID = id
						rtn.DateEntered = eTime
						rtn.DateUpdated = uTime
						rtn.OauthClientID = oauthID
						rtn.Enabled = enabled
						rtn.Company = (*foundRow)[1]
						rtn.FirstName = (*foundRow)[2]
						rtn.LastName = (*foundRow)[3]
						rtn.LocalDomain = (*foundRow)[4]
						rtn.RemoteDomain = (*foundRow)[5]
						rtn.OauthSecret = (*foundRow)[7]
						rtn.Email = (*foundRow)[8]
						rtn.City = (*foundRow)[9]
						rtn.State = (*foundRow)[10]
						rtn.Zip = (*foundRow)[11]
						rtn.StoreName = (*foundRow)[14]
						rtn.StoreSlogan = (*foundRow)[15]
						rtn.Logo = (*foundRow)[16]
						rtn.Currency = (*foundRow)[17]
					}
				}
			}
		}
	}
	return &rtn
}
