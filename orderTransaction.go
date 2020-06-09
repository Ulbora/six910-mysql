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

//Order Payment Transactions

//AddOrderTransaction AddOrderTransaction
func (d *Six910Mysql) AddOrderTransaction(t *mdb.OrderTransaction) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, t.Amount, t.Approval, t.Avs, time.Now(), t.Gwid, t.Method, t.OrderID,
		t.ReferenceNumber, t.ResponseCode, t.ResponseMessage, t.Success, t.TransactionID, t.Type)
	suc, id := d.DB.Insert(insertOrderTransaction, a...)
	d.Log.Debug("suc in add OrderTransaction", suc)
	d.Log.Debug("id in add OrderTransaction", id)
	return suc, id
}

//GetOrderTransactionList GetOrderTransactionList
func (d *Six910Mysql) GetOrderTransactionList(orderID int64) *[]mdb.OrderTransaction {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn []mdb.OrderTransaction
	var a []interface{}
	a = append(a, orderID)
	rows := d.DB.GetList(getOrderTransaction, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseOrderTransactionRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *Six910Mysql) parseOrderTransactionRow(foundRow *[]string) *mdb.OrderTransaction {
	d.Log.Debug("foundRow in get OrderTransaction", *foundRow)
	var rtn mdb.OrderTransaction
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get OrderTransaction", err)
		if err == nil {
			oid, err := strconv.ParseInt((*foundRow)[7], 10, 64)
			d.Log.Debug("id err in get OrderTransaction", err)
			if err == nil {
				gwid, err := strconv.ParseInt((*foundRow)[5], 10, 64)
				d.Log.Debug("id err in get OrderTransaction", err)
				if err == nil {
					amount, err := strconv.ParseFloat((*foundRow)[1], 64)
					d.Log.Debug("id err in get OrderTransaction", err)
					if err == nil {
						suc, enerr := strconv.ParseBool((*foundRow)[11])
						if enerr == nil {
							eTime, err := time.Parse(timeFormat, (*foundRow)[4])
							d.Log.Debug("eTime err in get OrderTransaction", err)
							if err == nil {
								rtn.ID = id
								rtn.OrderID = oid
								rtn.Gwid = gwid
								rtn.Amount = amount
								rtn.Success = suc
								rtn.DateEntered = eTime
								rtn.Approval = (*foundRow)[2]
								rtn.Avs = (*foundRow)[3]
								rtn.Method = (*foundRow)[6]
								rtn.ReferenceNumber = (*foundRow)[8]
								rtn.ResponseCode = (*foundRow)[9]
								rtn.ResponseMessage = (*foundRow)[10]
								rtn.TransactionID = (*foundRow)[12]
								rtn.Type = (*foundRow)[13]
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
