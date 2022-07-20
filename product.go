package six910mysql

import (
	"strconv"
	"strings"
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

//product

//AddProduct AddProduct
func (d *Six910Mysql) AddProduct(p *mdb.Product) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, p.Sku, p.Gtin, p.Name, p.ShortDesc, p.Desc, p.Cost, p.Msrp, p.Map, p.Price, p.SalePrice,
		p.Currency, p.ManufacturerID, p.Manufacturer, p.Stock, p.StockAlert, p.Weight, p.Width, p.Height, p.Depth,
		p.ShippingMarkup, p.Visible, p.Searchable, p.MultiBox, p.ShipSeparately, p.FreeShipping,
		time.Now(), p.DistributorID, p.Promoted, p.Dropship, p.Size, p.Color, p.ParentProductID,
		p.StoreID, p.Thumbnail, p.Image1, p.Image2, p.Image3, p.Image4, p.SpecialProcessing,
		p.SpecialProcessingType)
	suc, id := d.DB.Insert(insertProduct, a...)
	d.Log.Debug("suc in add Product", suc)
	d.Log.Debug("id in add Product", id)
	return suc, id
}

//UpdateProduct UpdateProduct
func (d *Six910Mysql) UpdateProduct(p *mdb.Product) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, p.Sku, p.Gtin, p.Name, p.ShortDesc, p.Desc, p.Cost, p.Msrp, p.Map, p.Price, p.SalePrice,
		p.Currency, p.ManufacturerID, p.Manufacturer, p.Stock, p.StockAlert, p.Weight, p.Width, p.Height, p.Depth,
		p.ShippingMarkup, p.Visible, p.Searchable, p.MultiBox, p.ShipSeparately, p.FreeShipping,
		time.Now(), p.DistributorID, p.Promoted, p.Dropship, p.Size, p.Color, p.ParentProductID,
		p.Thumbnail, p.Image1, p.Image2, p.Image3, p.Image4, p.SpecialProcessing,
		p.SpecialProcessingType, p.ID)
	suc := d.DB.Update(updateProduct, a...)
	return suc
}

//UpdateProductQuantity UpdateProductQuantity
func (d *Six910Mysql) UpdateProductQuantity(p *mdb.Product) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, p.Stock, p.ID)
	suc := d.DB.Update(updateProductQuantity, a...)
	return suc
}

//GetProductByID GetProductByID
func (d *Six910Mysql) GetProductByID(id int64) *mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getProduct, a...)
	rtn := d.parseProductRow(&row.Row)
	return rtn
}

//GetProductBySku GetProductBySku
func (d *Six910Mysql) GetProductBySku(sku string, distributorID int64, storeID int64) *mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, sku, distributorID, storeID)
	row := d.DB.Get(getProductBySku, a...)
	rtn := d.parseProductRow(&row.Row)
	return rtn
}

//GetProductsByPromoted GetProductsByPromoted
func (d *Six910Mysql) GetProductsByPromoted(storeID int64, start int64, end int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, storeID, start, end)
	rows := d.DB.GetList(getProductByPromoted, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetProductsByName GetProductsByName
func (d *Six910Mysql) GetProductsByName(name string, storeID int64, start int64, end int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, "%"+name+"%", storeID, start, end)
	rows := d.DB.GetList(getProductByName, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetProductsByCaterory GetProductsByCaterory
func (d *Six910Mysql) GetProductsByCaterory(catID int64, start int64, end int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, catID, start, end)
	rows := d.DB.GetList(getProductByCat, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetProductList GetProductList
func (d *Six910Mysql) GetProductList(storeID int64, start int64, end int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, storeID, start, end)
	rows := d.DB.GetList(getProductByStore, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetProductSubSkuList GetProductSubSkuList
func (d *Six910Mysql) GetProductSubSkuList(storeID int64, parentProdID int64) *[]mdb.Product {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []mdb.Product{}
	var a []interface{}
	a = append(a, storeID, parentProdID)
	rows := d.DB.GetList(getProductByParentSku, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseProductRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//GetProductIDList GetProductIDList
func (d *Six910Mysql) GetProductIDList(storeID int64) *[]int64 {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []int64{}
	var a []interface{}
	a = append(a, storeID)
	rows := d.DB.GetList(getProductIDList, a...)
	//d.Log.Debug("rows Order sales", *rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			if len(foundRow) > 0 {
				id, err := strconv.ParseInt(foundRow[0], 10, 64)
				d.Log.Debug("id err in get product id list", err)
				if err == nil {
					rtn = append(rtn, id)
				}
			}
		}
	}
	return &rtn
}

//GetProductIDListByCategories GetProductIDListByCategories
func (d *Six910Mysql) GetProductIDListByCategories(storeID int64, catList *[]int64) *[]int64 {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []int64{}
	var a []interface{}
	a = append(a, storeID)
	for _, id := range *catList {
		a = append(a, id)
	}
	var getProductIDListByCategory = " SELECT pc.product_id " +
		" FROM product_category pc " +
		" inner join category c " +
		" on c.id = pc.category_id " +
		" WHERE c.store_id = ? and pc.category_id IN (?" + strings.Repeat(",?", len(*catList)-1) + ")"
	rows := d.DB.GetList(getProductIDListByCategory, a...)
	//d.Log.Debug("rows Order sales", *rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			if len(foundRow) > 0 {
				id, err := strconv.ParseInt(foundRow[0], 10, 64)
				d.Log.Debug("product_id err in get product id by cat list", err)
				if err == nil {
					rtn = append(rtn, id)
				}
			}
		}
	}
	return &rtn
}

//DeleteProduct DeleteProduct
func (d *Six910Mysql) DeleteProduct(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return d.DB.Delete(deleteProduct, a...)
}

//DeleteSubProduct DeleteSubProduct
func (d *Six910Mysql) DeleteSubProduct(parentProdID int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, parentProdID)
	return d.DB.Delete(deleteSubProduct, a...)
}

func (d *Six910Mysql) parseProductRow(foundRow *[]string) *mdb.Product {
	d.Log.Debug("foundRow in get Product", *foundRow)
	var rtn mdb.Product
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Product", err)
		if err == nil {
			sid, serr := strconv.ParseInt((*foundRow)[33], 10, 64)
			d.Log.Debug("sid err in get Product", serr)
			if serr == nil {
				eTime, eerr := time.Parse(timeFormat, (*foundRow)[25])
				d.Log.Debug("eTime err in get Product", eerr)
				if eerr == nil {
					uTime, _ := time.Parse(timeFormat, (*foundRow)[26])
					visible, enerr := strconv.ParseBool((*foundRow)[20])
					if enerr == nil {
						cost, err := strconv.ParseFloat((*foundRow)[6], 64)
						d.Log.Debug("cost err in get Product", err)
						if err == nil {
							msrp, err := strconv.ParseFloat((*foundRow)[7], 64)
							d.Log.Debug("msrp err in get Product", err)
							if err == nil {
								mapPrice, err := strconv.ParseFloat((*foundRow)[8], 64)
								d.Log.Debug("mapPrice err in get Product", err)
								if err == nil {
									price, err := strconv.ParseFloat((*foundRow)[9], 64)
									d.Log.Debug("price err in get Product", err)
									if err == nil {
										salePrice, err := strconv.ParseFloat((*foundRow)[10], 64)
										d.Log.Debug("salePrice err in get Product", err)
										if err == nil {
											stock, err := strconv.ParseInt((*foundRow)[13], 10, 64)
											d.Log.Debug("stock err in get Product", err)
											if err == nil {
												stockAlert, err := strconv.ParseInt((*foundRow)[14], 10, 64)
												d.Log.Debug("stockAlert err in get Product", err)
												if err == nil {
													weight, err := strconv.ParseFloat((*foundRow)[15], 64)
													d.Log.Debug("weight err in get Product", err)
													if err == nil {
														width, err := strconv.ParseFloat((*foundRow)[16], 64)
														d.Log.Debug("width err in get Product", err)
														if err == nil {
															height, err := strconv.ParseFloat((*foundRow)[17], 64)
															d.Log.Debug("height err in get Product", err)
															if err == nil {
																depth, err := strconv.ParseFloat((*foundRow)[18], 64)
																d.Log.Debug("depth err in get Product", err)
																if err == nil {
																	sMarkup, err := strconv.ParseFloat((*foundRow)[19], 64)
																	d.Log.Debug("sMarkup err in get Product", err)
																	if err == nil {
																		searchable, err := strconv.ParseBool((*foundRow)[21])
																		if err == nil {
																			mbox, err := strconv.ParseBool((*foundRow)[22])
																			if err == nil {
																				sSep, err := strconv.ParseBool((*foundRow)[23])
																				if err == nil {
																					sFree, err := strconv.ParseBool((*foundRow)[24])
																					if err == nil {
																						promoted, err := strconv.ParseBool((*foundRow)[28])
																						if err == nil {
																							dship, err := strconv.ParseBool((*foundRow)[29])
																							if err == nil {
																								sproc, err := strconv.ParseBool((*foundRow)[39])
																								if err == nil {
																									did, err := strconv.ParseInt((*foundRow)[27], 10, 64)
																									d.Log.Debug("did err in get Product", err)
																									if err == nil {
																										ppid, err := strconv.ParseInt((*foundRow)[32], 10, 64)
																										d.Log.Debug("ppid err in get Product", err)
																										if err == nil {
																											//hspproc, err := strconv.ParseBool((*foundRow)[42])
																											//if err == nil {
																											rtn.ID = id
																											rtn.Cost = cost
																											rtn.DateEntered = eTime
																											rtn.DateUpdated = uTime
																											rtn.Depth = depth
																											rtn.DistributorID = did
																											rtn.Dropship = dship
																											rtn.FreeShipping = sFree
																											rtn.Height = height
																											rtn.Map = mapPrice
																											rtn.Msrp = msrp
																											rtn.MultiBox = mbox
																											rtn.ParentProductID = ppid
																											rtn.Price = price
																											rtn.Promoted = promoted
																											rtn.SalePrice = salePrice
																											rtn.Searchable = searchable
																											rtn.ShipSeparately = sSep
																											rtn.ShippingMarkup = sMarkup
																											rtn.SpecialProcessing = sproc
																											rtn.Stock = stock
																											rtn.StockAlert = stockAlert
																											rtn.StoreID = sid
																											rtn.Visible = visible
																											rtn.Weight = weight
																											rtn.Width = width
																											//rtn.SubSku = hspproc
																											rtn.Sku = (*foundRow)[1]
																											rtn.Gtin = (*foundRow)[2]
																											rtn.Name = (*foundRow)[3]
																											rtn.ShortDesc = (*foundRow)[4]
																											rtn.Desc = (*foundRow)[5]
																											rtn.Currency = (*foundRow)[11]
																											rtn.Manufacturer = (*foundRow)[12]
																											rtn.Size = (*foundRow)[30]
																											rtn.Color = (*foundRow)[31]
																											rtn.Thumbnail = (*foundRow)[34]
																											rtn.Image1 = (*foundRow)[35]
																											rtn.Image2 = (*foundRow)[36]
																											rtn.Image3 = (*foundRow)[37]
																											rtn.Image4 = (*foundRow)[38]
																											rtn.SpecialProcessingType = (*foundRow)[40]
																											rtn.ManufacturerID = (*foundRow)[41]
																											rtn.Gender = (*foundRow)[42]
																											//}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return &rtn
}
