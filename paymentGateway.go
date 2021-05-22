package six910mysql

import (
	"strconv"

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

//Plugins that are payment gateways

//AddPaymentGateway AddPaymentGateway
func (d *Six910Mysql) AddPaymentGateway(pgw *mdb.PaymentGateway) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, pgw.StorePluginsID, pgw.CheckoutURL, pgw.PostOrderURL, pgw.LogoURL, pgw.ClientID,
		pgw.ClientKey, pgw.Name, pgw.Token)
	suc, id := d.DB.Insert(insertPaymentGateway, a...)
	d.Log.Debug("suc in add PaymentGateway", suc)
	d.Log.Debug("id in add PaymentGateway", id)
	return suc, id
}

//UpdatePaymentGateway UpdatePaymentGateway
func (d *Six910Mysql) UpdatePaymentGateway(pgw *mdb.PaymentGateway) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, pgw.CheckoutURL, pgw.PostOrderURL, pgw.LogoURL, pgw.ClientID,
		pgw.ClientKey, pgw.Token, pgw.ID)
	suc := d.DB.Update(updatePaymentGateway, a...)
	return suc
}

//GetPaymentGateway GetPaymentGateway
func (d *Six910Mysql) GetPaymentGateway(id int64) *mdb.PaymentGateway {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getPaymentGateway, a...)
	rtn := d.parsePaymentGatewayRow(&row.Row)
	return rtn
}

//GetPaymentGatewayByName GetPaymentGatewayByName
func (d *Six910Mysql) GetPaymentGatewayByName(name string, storeID int64) *mdb.PaymentGateway {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, name, storeID)
	row := d.DB.Get(getPaymentGatewayByName, a...)
	rtn := d.parsePaymentGatewayRow(&row.Row)
	return rtn
}

//GetPaymentGateways GetPaymentGateways
func (d *Six910Mysql) GetPaymentGateways(storeID int64) *[]mdb.PaymentGateway {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.PaymentGateway{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getPaymentGatewayByStore, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parsePaymentGatewayRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeletePaymentGateway DeletePaymentGateway
func (d *Six910Mysql) DeletePaymentGateway(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deletePaymentGateway, a...)
}

func (d *Six910Mysql) parsePaymentGatewayRow(foundRow *[]string) *mdb.PaymentGateway {
	d.Log.Debug("foundRow in get PaymentGateway", *foundRow)
	var rtn mdb.PaymentGateway
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get PaymentGateway", err)
		if err == nil {
			spiid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			d.Log.Debug("spiid err in get PaymentGateway", err)
			if err == nil {
				rtn.ID = id
				rtn.StorePluginsID = spiid
				rtn.CheckoutURL = (*foundRow)[2]
				rtn.PostOrderURL = (*foundRow)[3]
				rtn.LogoURL = (*foundRow)[4]
				rtn.ClientID = (*foundRow)[5]
				rtn.ClientKey = (*foundRow)[6]
				rtn.Name = (*foundRow)[7]
				rtn.Token = (*foundRow)[8]
			}
		}
	}
	return &rtn
}
