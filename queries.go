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

const (
	cartTest = "select count(*) from store "

	insertSecurity = "INSERT INTO security (oauth_on) values(?) "
	updateSecurity = "UPDATE security SET oauth_on = ? WHERE id = ? "
	getSecurity    = "SELECT id, oauth_on from security "
	deleteSecurity = "DELETE FROM security "

	insertStore = " insert into store (company, first_name, last_name, local_domain, " +
		" remote_domain, oauth_client_id, oauth_secret, email, city, state, zip, date_entered, " +
		" store_name, store_slogan, logo, currency, enabled) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	updateStore = " UPDATE store SET company = ?, first_name = ?, last_name = ?, local_domain = ?, " +
		" remote_domain = ?, oauth_client_id = ?, oauth_secret = ?, email = ?, city = ?, " +
		"state = ?, zip = ?, store_name = ?, store_slogan = ?, logo = ?, currency = ?, date_updated = ?, enabled = ? " +
		" WHERE id = ? "

	getStoreByName = "SELECT id, company, first_name, last_name, local_domain, remote_domain, " +
		" oauth_client_id, oauth_secret, email, city, state, zip, date_entered, date_updated, " +
		" store_name, store_slogan, logo, currency, enabled " +
		" FROM store " +
		" WHERE store_name = ? "
	getStoreByID = "SELECT id, company, first_name, last_name, local_domain, remote_domain, " +
		" oauth_client_id, oauth_secret, email, city, state, zip, date_entered, date_updated, " +
		" store_name, store_slogan, logo, currency, enabled " +
		" FROM store " +
		" WHERE id = ? "
	getAllStores = "SELECT id, company, first_name, last_name, local_domain, remote_domain, " +
		" oauth_client_id, oauth_secret, email, city, state, zip, date_entered, date_updated, " +
		" store_name, store_slogan, logo, currency, enabled " +
		" FROM store "

	getStoreByLocalDomain = "SELECT id, company, first_name, last_name, local_domain, remote_domain, " +
		" oauth_client_id, oauth_secret, email, city, state, zip, date_entered, date_updated, " +
		" store_name, store_slogan, logo, currency, enabled " +
		" FROM store " +
		" WHERE local_domain = ? "

	getStoreCount = " SELECT count(*) as cnt " +
		" FROM store "

	deleteStore = "DELETE FROM store WHERE id = ? "

	insertCustomer = "INSERT into customer(email, reset_password, first_name, last_name, " +
		" company, city, state, zip, phone, store_id, date_entered) " +
		" values(?,?,?,?,?,?,?,?,?,?,?)"

	updateCustomer = " UPDATE customer SET email = ?, reset_password = ?, first_name = ?, " +
		" last_name = ?, company = ?, city = ?, state = ?, zip = ?, phone = ?, date_updated = ? " +
		" WHERE id = ? "

	getCustemerByEmail = "SELECT id, email, reset_password, first_name, last_name, " +
		" company, city, state, zip, phone, store_id, date_entered, date_updated " +
		" FROM customer " +
		" WHERE email = ? and store_id = ? "

	getCustemerByID = "SELECT id, email, reset_password, first_name, last_name, " +
		" company, city, state, zip, phone, store_id, date_entered, date_updated " +
		" FROM customer " +
		" WHERE id = ? "

	getCustemerList = "SELECT id, email, reset_password, first_name, last_name, " +
		" company, city, state, zip, phone, store_id, date_entered, date_updated " +
		" FROM customer " +
		" WHERE store_id = ? " +
		" ORDER BY email " +
		" LIMIT ?, ? "

	deleteCustomer = "DELETE FROM customer WHERE id = ? "

	insertLocalAccount = " insert into local_account (username, password, enabled, role, " +
		" store_id, customer_id, date_entered )" +
		" values(?, ?, ?, ?, ?, ?, ?)"

	updateLocalAccount = " UPDATE local_account SET password = ?, enabled = ?, role = ?, " +
		" date_updated = ? " +
		" WHERE username = ? and store_id = ?"

	getLocalAccount = "SELECT username, password, enabled, role, " +
		" store_id, customer_id, date_entered, date_updated " +
		" FROM local_account " +
		" WHERE username = ? and store_id = ?"
	getLocalAccountList = "SELECT username, password, enabled, role, " +
		" store_id, customer_id, date_entered, date_updated " +
		" FROM local_account " +
		" WHERE store_id = ?"

	deleteLocalAccount = "DELETE FROM local_account WHERE username = ? and store_id = ? "

	insertAddress = "INSERT into address(address, city, state, zip, " +
		" county, country, address_type, customer_id) " +
		" values(?,?,?,?,?,?,?,?)"

	updateAddress = " UPDATE address SET address = ?, city = ?, state = ?, " +
		" zip = ?, county = ?, country = ?, address_type = ? " +
		" WHERE id = ?"

	getAddress = "SELECT id, address, city, state, " +
		" zip, county, country, address_type, customer_id " +
		" FROM address " +
		" WHERE id = ? "
	getAddressList = "SELECT id, address, city, state, " +
		" zip, county, country, address_type, customer_id " +
		" FROM address " +
		" WHERE customer_id = ? "

	deleteAddress = "DELETE FROM address WHERE id = ? "

	insertDistributor = "INSERT into distributor(company, contact_name, phone, store_id) " +
		" values(?,?,?,?)"

	updateDistributor = " UPDATE distributor SET company = ?, contact_name = ?, phone = ? " +
		" WHERE id = ?"

	getDistributor = "SELECT id, company, contact_name, phone, " +
		" store_id " +
		" FROM distributor " +
		" WHERE id = ? "

	getDistributorList = "SELECT id, company, contact_name, phone, " +
		" store_id " +
		" FROM distributor " +
		" WHERE store_id = ? " +
		" ORDER BY company "

	deleteDistributor = "DELETE FROM distributor WHERE id = ? "

	insertCategory = "INSERT into category(name, description, image, thumbnail, " +
		" store_id, parent_category_id) " +
		" values(?,?,?,?,?,?)"

	updateCategory = " UPDATE category SET name = ?, description = ?, image = ?, " +
		" thumbnail = ?, parent_category_id = ? " +
		" WHERE id = ?"

	getCategory = "SELECT id, name, description, image, thumbnail, parent_category_id, " +
		" store_id " +
		" FROM category " +
		" WHERE id = ? "

	getCategoryList = "SELECT id, name, description, image, thumbnail, parent_category_id, " +
		" store_id " +
		" FROM category " +
		" WHERE store_id = ? " +
		" ORDER by name "

	getHierarchicalCategoryList = " SELECT c6.name as c6, c5.name as c5, c4.name as c4, " +
		" c3.name as c3, c2.name as c2, c1.name as c1, c1.id " +
		" FROM category c1 " +
		" LEFT JOIN category c2 ON c1.parent_category_id=c2.id " +
		" LEFT JOIN category c3 ON c2.parent_category_id=c3.id " +
		" LEFT JOIN category c4 ON c3.parent_category_id=c4.id " +
		" LEFT JOIN category c5 ON c4.parent_category_id=c4.id " +
		" LEFT JOIN category c6 ON c5.parent_category_id=c5.id " +
		" WHERE c1.store_id = ? "

	getSubCategoryList = "SELECT id, name, description, image, thumbnail, parent_category_id, " +
		" store_id " +
		" FROM category " +
		" WHERE parent_category_id = ? "

	deleteCategory = "DELETE FROM category WHERE id = ? or parent_category_id = ?"

	insertProduct = "INSERT into product(sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer_id, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type) " +
		" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	updateProduct = "UPDATE product SET sku = ?, gtin = ?, name = ?, short_description = ?, description = ?, " +
		" cost = ?, msrp = ?, map = ?, price = ?, sale_price = ?, currency = ?, manufacturer_id = ?,  manufacturer = ?, stock = ?, " +
		" stock_alert = ?, weight = ?, " +
		" width = ?, height = ?, depth = ?, shipping_markup = ?, visible = ?, searchable = ?, multibox = ?, " +
		" ship_separate = ?, free_shipping = ?, date_updated = ?, distributor_id = ?, promoted = ?, dropship = ?, " +
		" size = ?, color = ?, parient_product_id = ?, thumbnail = ?, image1 = ?, image2 = ?, image3 = ?, " +
		" image4 = ?, special_processing = ?, special_processing_type = ? " +
		" WHERE id = ?"

	getProduct = "SELECT id, sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, date_updated, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type, manufacturer_id " +
		" FROM product " +
		" WHERE id = ? "
	getProductBySku = "SELECT id, sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, date_updated, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type, manufacturer_id " +
		" FROM product " +
		" WHERE sku = ? and distributor_id = ? and store_id = ? "

	getProductByName = "SELECT id, sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, date_updated, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type, manufacturer_id " +
		" FROM product " +
		" WHERE name like ? and store_id = ? LIMIT ?, ? "

	// get manf list by prod name
	// "select distinct manufacturer
	// from product
	// where name like '%holster%'"

	getProductManufacturerListByProductName = " SELECT DISTINCT manufacturer " +
		" FROM product " +
		" WHERE name like ? and store_id = ? " +
		" ORDER by manufacturer"

	//get product by manf name and name
	// 	select *
	// from product
	// where manufacturer = 'Blackhawk'
	// and name like '%holster%'

	getProductByNameAndManfName = " SELECT id, sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, date_updated, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type, manufacturer_id " +
		" FROM product " +
		" WHERE manufacturer = ? and name like ? and store_id = ? LIMIT ?, ? "

	//get manf by cat id
	// 	SELECT distinct p.manufacturer
	// FROM product p
	// inner join product_category pc
	// on p.id = pc.product_id
	// inner join category c
	// on pc.category_id = c.id
	// WHERE c.id = 188
	// order by p.manufacturer

	getProductManufacturerListByCatID = " SELECT distinct p.manufacturer " +
		" FROM product p " +
		" inner join product_category pc " +
		" on p.id = pc.product_id " +
		" inner join category c " +
		" on pc.category_id = c.id " +
		" WHERE c.id = ? and p.store_id = ? " +
		" order by p.manufacturer "

	//get product by cat and manf
	// 	SELECT p.id, p.sku, p.gtin, p.name, p.short_description, p.description,
	// p.cost, p.msrp, p.map, p.price, p.sale_price, p.currency, p.manufacturer, p.stock, p.stock_alert, p.weight,
	// p.width, p.height, p.depth, p.shipping_markup, p.visible, p.searchable, p.multibox,
	// p.ship_separate, p.free_shipping, p.date_entered, p.date_updated, p.distributor_id, p.promoted, p.dropship,
	// p.size, p.color, p.parient_product_id, p.store_id, p.thumbnail, p.image1, p.image2, p.image3,
	// p.image4, p.special_processing, p.special_processing_type, p.manufacturer_id
	// FROM product p
	// inner join product_category pc
	// on p.id = pc.product_id
	// inner join category c
	// on pc.category_id = c.id
	// WHERE c.id = 200 and p.manufacturer = 'Blackhawk'
	// ORDER BY p.name

	getProductByCatAndManufacturer = "SELECT p.id, p.sku, p.gtin, p.name, p.short_description, p.description, " +
		" p.cost, p.msrp, p.map, p.price, p.sale_price, p.currency, p.manufacturer, p.stock, p.stock_alert, p.weight, " +
		" p.width, p.height, p.depth, p.shipping_markup, p.visible, p.searchable, p.multibox, " +
		" p.ship_separate, p.free_shipping, p.date_entered, p.date_updated, p.distributor_id, p.promoted, p.dropship, " +
		" p.size, p.color, p.parient_product_id, p.store_id, p.thumbnail, p.image1, p.image2, p.image3, " +
		" p.image4, p.special_processing, p.special_processing_type, p.manufacturer_id " +
		" FROM product p " +
		" inner join product_category pc " +
		" on p.id = pc.product_id " +
		" inner join category c " +
		" on pc.category_id = c.id " +
		" WHERE c.id = ? and p.manufacturer = ? and p.store_id = ? " +
		" ORDER BY p.name LIMIT ?, ? "

	getProductByPromoted = "SELECT id, sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, date_updated, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type, manufacturer_id " +
		" FROM product " +
		" WHERE searchable = true and visible = true and promoted = true and store_id = ? " +
		" ORDER BY name LIMIT ?, ? "

	getProductByCat = "SELECT p.id, p.sku, p.gtin, p.name, p.short_description, p.description, " +
		" p.cost, p.msrp, p.map, p.price, p.sale_price, p.currency, p.manufacturer, p.stock, p.stock_alert, p.weight, " +
		" p.width, p.height, p.depth, p.shipping_markup, p.visible, p.searchable, p.multibox, " +
		" p.ship_separate, p.free_shipping, p.date_entered, p.date_updated, p.distributor_id, p.promoted, p.dropship, " +
		" p.size, p.color, p.parient_product_id, p.store_id, p.thumbnail, p.image1, p.image2, p.image3, " +
		" p.image4, p.special_processing, p.special_processing_type, p.manufacturer_id " +
		" FROM product p " +
		" inner join product_category pc " +
		" on p.id = pc.product_id " +
		" inner join category c " +
		" on pc.category_id = c.id " +
		" WHERE c.id = ? ORDER BY p.name LIMIT ?, ? "

	getProductByStore = "SELECT id, sku, gtin, name, short_description, description, " +
		" cost, msrp, map, price, sale_price, currency, manufacturer, stock, stock_alert, weight, " +
		" width, height, depth, shipping_markup, visible, searchable, multibox, " +
		" ship_separate, free_shipping, date_entered, date_updated, distributor_id, promoted, dropship, " +
		" size, color, parient_product_id, store_id, thumbnail, image1, image2, image3, " +
		" image4, special_processing, special_processing_type, manufacturer_id " +
		" FROM product " +
		" WHERE store_id = ? ORDER BY name LIMIT ?, ? "

	deleteProduct = "DELETE FROM product WHERE id = ? "

	insertProductCategory = "INSERT INTO product_category (category_id, product_id) values(?, ?) "

	getProductCategory = "SELECT category_id " +
		" FROM product_category " +
		" WHERE product_id = ? "

	deleteProductCategory = "DELETE FROM product_category WHERE category_id = ? and product_id = ? "

	insertCart = "INSERT INTO cart (store_id, customer_id, date_entered) values(?, ?, ?) "

	updateCart = "UPDATE cart SET date_updated = ?, customer_id = ? WHERE id = ? "

	getCart = "SELECT id, store_id, customer_id, date_entered, date_updated " +
		" FROM cart " +
		" WHERE customer_id = ? "

	deleteCart = "DELETE FROM cart WHERE id = ? "

	insertCartItem = "INSERT INTO cart_item (cart_id, quantity, product_id) values(?, ?, ?) "

	updateCartItem = "UPDATE cart_item SET quantity = ? WHERE id = ? "

	getCartItem = "SELECT id, quantity, cart_id, product_id " +
		" FROM cart_item " +
		" WHERE cart_id = ? and product_id = ? "

	getCartItemList = "SELECT id, quantity, cart_id, product_id " +
		" FROM cart_item " +
		" WHERE cart_id = ? "

	deleteCartItem = "DELETE FROM cart_item WHERE id = ? "

	insertRegion = "INSERT INTO region (name, region_code, store_id) values(?, ?, ?) "

	updateRegion = "UPDATE region SET name = ?, region_code = ? WHERE id = ? "

	getRegion = "SELECT id, name, region_code, store_id " +
		" FROM region " +
		" WHERE id = ? "

	getRegionList = "SELECT id, name, region_code, store_id " +
		" FROM region " +
		" WHERE store_id = ? " +
		" ORDER BY name "

	deleteRegion = "DELETE FROM region WHERE id = ? "

	insertSubRegion = "INSERT INTO sub_region (name, sub_region_code, region_id) values(?, ?, ?) "

	updateSubRegion = "UPDATE sub_region SET name = ?, sub_region_code = ? WHERE id = ? "

	getSubRegion = "SELECT id, name, sub_region_code, region_id " +
		" FROM sub_region " +
		" WHERE id = ? "

	getSubRegionList = "SELECT id, name, sub_region_code, region_id " +
		" FROM sub_region " +
		" WHERE region_id = ? "

	deleteSubRegion = "DELETE FROM sub_region WHERE id = ? "

	insertShippingCarrier = "INSERT INTO shipping_carrier (carrier, type, store_id) values(?, ?, ?) "

	updateShippingCarrier = "UPDATE shipping_carrier SET carrier = ?, type = ? WHERE id = ? "

	getShipmentCarrier = "SELECT id, carrier, type, store_id " +
		" FROM shipping_carrier " +
		" WHERE id = ? "

	getShippingCarrierList = "SELECT id, carrier, type, store_id " +
		" FROM shipping_carrier " +
		" WHERE store_id = ? " +
		" ORDER BY carrier "

	deleteShippingCarrier = "DELETE FROM shipping_carrier WHERE id = ? "

	insertInsurance = "INSERT INTO insurance (cost, minimum_order_amount, maximum_order_amount, store_id) values(?, ?, ?, ?) "

	updateInsurance = "UPDATE insurance SET cost = ?, minimum_order_amount = ?, maximum_order_amount = ?  WHERE id = ? "

	getInsurance = "SELECT id, cost, minimum_order_amount, maximum_order_amount, store_id " +
		" FROM insurance " +
		" WHERE id = ? "

	getInsuranceList = "SELECT id, cost, minimum_order_amount, maximum_order_amount, store_id " +
		" FROM insurance " +
		" WHERE store_id = ? "

	deleteInsurance = "DELETE FROM insurance WHERE id = ? "

	insertShippingMethod = "INSERT INTO shipping_method (name, cost, max_weight, handling, " +
		"minimum_order, maximum_order, region_id, shipping_carrier_id, " +
		" insurance_id, store_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updateShippingMethod = "UPDATE shipping_method SET name = ?, cost = ?, max_weight = ?, " +
		" handling = ?, minimum_order = ?, maximum_order = ?  WHERE id = ? "

	getShippingMethod = "SELECT id, name, cost, max_weight, handling, " +
		"minimum_order, maximum_order, region_id, shipping_carrier_id, insurance_id, store_id" +
		" FROM shipping_method " +
		" WHERE id = ? "

	getShippingMethodList = "SELECT id, name, cost, max_weight, handling, " +
		"minimum_order, maximum_order, region_id, shipping_carrier_id, insurance_id, store_id" +
		" FROM shipping_method " +
		" WHERE store_id = ? " +
		" ORDER BY name "

	deleteShippingMethod = "DELETE FROM shipping_method WHERE id = ? "

	insertTaxRate = "INSERT INTO tax_rate (country, state, zip_start, zip_end, " +
		"percent_rate, product_category_id, include_handling, include_shipping, " +
		" tax_type, store_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updateTaxRate = "UPDATE tax_rate SET zip_start = ?, zip_end = ?, percent_rate = ?, " +
		" product_category_id = ?, include_handling = ?, include_shipping = ?, tax_type = ?  " +
		" WHERE id = ? "

	getTaxRate = "SELECT id, country, state, zip_start, zip_end, " +
		" percent_rate, product_category_id, include_handling, include_shipping, tax_type, store_id" +
		" FROM tax_rate " +
		" WHERE country = ? and state = ? and store_id = ? "

	getTaxRateList = "SELECT id, country, state, zip_start, zip_end, " +
		" percent_rate, product_category_id, include_handling, include_shipping, tax_type, store_id" +
		" FROM tax_rate " +
		" WHERE store_id = ? "

	deleteTaxRate = "DELETE FROM tax_rate WHERE id = ? "

	insertIncludedSubRegion = "INSERT INTO included_sub_region (region_id, sub_region_id, " +
		" shipping_method_id) values(?, ?, ?) "

	getIncludedSubRegion = "SELECT id, region_id, sub_region_id, shipping_method_id " +
		" FROM included_sub_region " +
		" WHERE id = ? "

	getIncludedSubRegionList = "SELECT id, region_id, sub_region_id, shipping_method_id " +
		" FROM included_sub_region " +
		" WHERE region_id = ? "

	deleteIncludedSubRegion = "DELETE FROM included_sub_region WHERE id = ? "

	insertExcludedSubRegion = "INSERT INTO excluded_sub_region (region_id, sub_region_id, " +
		" shipping_method_id) values(?, ?, ?) "

	getExcludedSubRegion = "SELECT id, region_id, sub_region_id, shipping_method_id " +
		" FROM excluded_sub_region " +
		" WHERE id = ? "

	getExcludedSubRegionList = "SELECT id, region_id, sub_region_id, shipping_method_id " +
		" FROM excluded_sub_region " +
		" WHERE region_id = ? "

	deleteExcludedSubRegion = "DELETE FROM excluded_sub_region WHERE id = ? "

	insertZoneZip = "INSERT INTO zone_zip (zip_code, included_sub_region_id, " +
		" excluded_sub_region_id) values(?, ?, ?) "

	getZipExclustionList = "SELECT id, zip_code, included_sub_region_id, excluded_sub_region_id " +
		" FROM zone_zip " +
		" WHERE excluded_sub_region_id = ? "

	getZipInclustionList = "SELECT id, zip_code, included_sub_region_id, excluded_sub_region_id " +
		" FROM zone_zip " +
		" WHERE included_sub_region_id = ? "

	deleteZoneZip = "DELETE FROM zone_zip WHERE id = ? "

	insertPlugin = "INSERT INTO plugins (plugin_name, developer, contact_phone, developer_address, " +
		"fee, category, activate_url, oauth_redirect_url, " +
		" is_pgw, enabled) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updatePlugin = "UPDATE plugins SET plugin_name = ?, developer = ?, contact_phone = ?, developer_address = ?, " +
		"fee = ?, category = ?, activate_url = ?, oauth_redirect_url = ?, " +
		" is_pgw = ?, enabled = ? " +
		" WHERE id = ? "

	getPlugin = "SELECT id, plugin_name, developer, contact_phone, developer_address, " +
		"fee, category, activate_url, oauth_redirect_url, " +
		" is_pgw, enabled  " +
		" FROM plugins " +
		" WHERE id = ? "

	getPluginList = "SELECT id, plugin_name, developer, contact_phone, developer_address, " +
		"fee, category, activate_url, oauth_redirect_url, " +
		" is_pgw, enabled  " +
		" FROM plugins " +
		" ORDER BY plugin_name " +
		" LIMIT ?, ? "

	deletePlugin = "DELETE FROM plugins WHERE id = ? "

	insertStorePlugin = "INSERT INTO store_plugins (plugins_id, plugin_name, category, active, " +
		" oauth_client_id, oauth_secret, activate_url, oauth_redirect_url, " +
		" api_key, rekey_try_count, iframe_url, menu_title, menu_icon_url, is_pgw, " +
		" store_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updateStorePluginNoDate = "UPDATE store_plugins SET  active = ?, " +
		" oauth_client_id = ?, oauth_secret = ?, activate_url = ?, oauth_redirect_url = ?, " +
		" api_key = ?, rekey_try_count = ?, iframe_url = ?, menu_title = ?, menu_icon_url = ? " +
		" WHERE id = ? "

	updateStorePluginDate = "UPDATE store_plugins SET  active = ?, " +
		" oauth_client_id = ?, oauth_secret = ?, activate_url = ?, oauth_redirect_url = ?, " +
		" api_key = ?, rekey_try_count = ?, iframe_url = ?, menu_title = ?, menu_icon_url = ?, " +
		" rekey_date = ? " +
		" WHERE id = ? "

	getStorePlugin = "SELECT id, plugins_id, plugin_name, category, active, " +
		" oauth_client_id, oauth_secret, activate_url, oauth_redirect_url, " +
		" api_key, rekey_try_count, rekey_date, iframe_url, menu_title, menu_icon_url, is_pgw, " +
		" store_id " +
		" FROM store_plugins " +
		" WHERE id = ? "

	getStorePluginList = "SELECT id, plugins_id, plugin_name, category, active, " +
		" oauth_client_id, oauth_secret, activate_url, oauth_redirect_url, " +
		" api_key, rekey_try_count, rekey_date, iframe_url, menu_title, menu_icon_url, is_pgw, " +
		" store_id " +
		" FROM store_plugins " +
		" WHERE store_id = ? " +
		" ORDER BY plugin_name "

	deleteStorePlugin = "DELETE FROM store_plugins WHERE id = ? "

	insertPaymentGateway = "INSERT into payment_gateway(store_plugins_id, checkout_url, post_order_url, logo_url, " +
		" client_id, client_key) " +
		" values(?,?,?,?,?,?)"

	updatePaymentGateway = "UPDATE payment_gateway SET checkout_url = ?, post_order_url = ?, logo_url = ?, " +
		" client_id = ?, client_key = ? " +
		" WHERE id = ? "

	getPaymentGateway = "SELECT id, store_plugins_id, checkout_url, post_order_url, logo_url, " +
		" client_id, client_key " +
		" FROM payment_gateway " +
		" WHERE id = ? "

	getPaymentGatewayByStore = " SELECT g.id, g.store_plugins_id, g.checkout_url, " +
		" g.post_order_url, g.logo_url, g.client_id, g.client_key " +
		" FROM payment_gateway g " +
		" inner join store_plugins sp " +
		" on sp.id = g.store_plugins_id " +
		" inner join store s " +
		" on s.id = sp.store_id " +
		" where s.id = ? and sp.active = true"

	deletePaymentGateway = "DELETE FROM payment_gateway WHERE id = ? "

	insertLocalDatastore = "INSERT INTO local_data_store (store_id, data_store_name, " +
		" reload, reload_date) values(?, ?, ?, ?) "

	updateLocalDatastore = "UPDATE local_data_store SET reload = ?, reload_date = ? " +
		" WHERE store_id = ? and data_store_name = ? "

	getLocalDatastore = "SELECT store_id, data_store_name, reload, reload_date " +
		" FROM local_data_store " +
		" WHERE store_id = ? and data_store_name = ? "

	insertInstances = "INSERT INTO instances (instance_name, reload_date, " +
		" store_id, data_store_name) values(?, ?, ?, ?) "

	updateInstances = "UPDATE instances SET reload_date = ? " +
		" WHERE instance_name = ? and store_id = ? and data_store_name = ? "

	getInstances = "SELECT instance_name, reload_date, store_id, data_store_name " +
		" FROM instances " +
		" WHERE instance_name = ? and store_id = ? and data_store_name = ? "

	getInstancesList = "SELECT instance_name, reload_date, store_id, data_store_name " +
		" FROM instances " +
		" WHERE store_id = ? and data_store_name = ? "

	insertDataStoreWriteLock = "INSERT INTO datastore_write_lock (datastore_name, locked, " +
		" locking_instance_name, locked_by_user, locked_time, store_id) values(?, ?, ?, ?, ?, ?) "

	updateDataStoreWriteLock = "UPDATE datastore_write_lock SET locked = ?, " +
		" locking_instance_name = ?, locked_by_user = ?, locked_time = ? " +
		" WHERE datastore_name = ? and store_id = ? "

	getDataStoreWriteLock = "SELECT datastore_name, locked, locking_instance_name, locked_by_user, " +
		" locked_time, store_id " +
		" FROM datastore_write_lock " +
		" WHERE datastore_name = ? and store_id = ?  "

	insertOrder = "INSERT INTO orders (order_date, status, subtotal, shipping_handling, " +
		" insurance, taxes, total, customer_id, billing_address_id, shipping_address_id, customer_name, " +
		" billing_address, shipping_address, store_id, order_number, order_type, pickup, username, " +
		" shipping_method_id, shipping_method_name)" +
		" values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updateOrder = "UPDATE orders SET updated = ?, status = ?, subtotal = ?, shipping_handling = ?, " +
		" insurance = ?, taxes = ?, total = ?, billing_address_id = ?, shipping_address_id = ?, customer_name = ?, " +
		" billing_address = ?, shipping_address = ?, order_type = ?, pickup = ?, username = ?, " +
		" shipping_method_id = ?, shipping_method_name = ?, refunded = ? " +
		" WHERE id = ? "

	getOrder = "SELECT id, order_date, updated, status, subtotal, shipping_handling, " +
		" insurance, taxes, total, customer_id, billing_address_id, shipping_address_id, customer_name, " +
		" billing_address, shipping_address, store_id, order_number, order_type, pickup, username, " +
		" shipping_method_id, shipping_method_name, refunded " +
		" FROM orders " +
		" WHERE id = ? "

	getOrderByCid = "SELECT id, order_date, updated, status, subtotal, shipping_handling, " +
		" insurance, taxes, total, customer_id, billing_address_id, shipping_address_id, customer_name, " +
		" billing_address, shipping_address, store_id, order_number, order_type, pickup, username, " +
		" shipping_method_id, shipping_method_name, refunded " +
		" FROM orders " +
		" WHERE customer_id = ? and store_id = ? " +
		" ORDER BY order_date "

	getOrderForStore = "SELECT id, order_date, updated, status, subtotal, shipping_handling, " +
		" insurance, taxes, total, customer_id, billing_address_id, shipping_address_id, customer_name, " +
		" billing_address, shipping_address, store_id, order_number, order_type, pickup, username, " +
		" shipping_method_id, shipping_method_name, refunded " +
		" FROM orders " +
		" WHERE store_id = ? " +
		" ORDER by status"

	getOrderForStoreByStatus = "SELECT id, order_date, updated, status, subtotal, shipping_handling, " +
		" insurance, taxes, total, customer_id, billing_address_id, shipping_address_id, customer_name, " +
		" billing_address, shipping_address, store_id, order_number, order_type, pickup, username, " +
		" shipping_method_id, shipping_method_name, refunded " +
		" FROM orders " +
		" WHERE store_id = ? and status = ? "

	deleteOrder = "DELETE FROM orders WHERE id = ? "

	insertOrderItem = "INSERT INTO order_item (order_id, product_id, product_name, product_short_desc, " +
		" quantity, dropship, backordered, price, total, image) " +
		" values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updateOrderItem = " UPDATE order_item SET quantity = ?, dropship = ?, backordered = ?, " +
		" price = ?, total = ?, image = ? " +
		" WHERE id = ? "

	getOrderItem = "SELECT id, order_id, product_id, product_name, product_short_desc, " +
		" quantity, dropship, backordered, price, total, image " +
		" FROM order_item " +
		" WHERE id = ? "

	getOrderItemList = "SELECT id, order_id, product_id, product_name, product_short_desc, " +
		" quantity, dropship, backordered, price, total, image " +
		" FROM order_item " +
		" WHERE order_id = ? "

	deleteOrderItem = "DELETE FROM order_item WHERE id = ? "

	insertOrderComment = "INSERT INTO order_comment (comment, username, order_id ) " +
		" values(?, ?, ?) "

	getOrderCommentList = "SELECT id, comment, username, order_id " +
		" FROM order_comment " +
		" WHERE order_id = ? "

	insertOrderTransaction = "INSERT INTO order_transaction (amount, approval, avs, date_entered, " +
		" gwid, method, order_id, ref_number, response_code, response_message, success, " +
		" transaction_id, type)" +
		" values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	getOrderTransaction = "SELECT id, amount, approval, avs, date_entered, " +
		" gwid, method, order_id, ref_number, response_code, response_message, success, " +
		" transaction_id, type " +
		" FROM  order_transaction " +
		" WHERE order_id = ? "

	insertShipment = "INSERT INTO shipment (order_id, create_date, status, boxes, " +
		" shipping_handling, insurance) " +
		" values(?, ?, ?, ?, ?, ?) "

	updateShipment = " UPDATE shipment SET status = ?, boxes = ?, shipping_handling = ?, " +
		" insurance = ?, updated = ? " +
		" WHERE id = ? "

	getShipment = "SELECT id, order_id, create_date, status, boxes, " +
		" shipping_handling, insurance, updated " +
		" FROM  shipment " +
		" WHERE id = ? "

	getShipmentList = "SELECT id, order_id, create_date, status, boxes, " +
		" shipping_handling, insurance, updated " +
		" FROM  shipment " +
		" WHERE order_id = ? "

	deleteShipment = "DELETE FROM shipment WHERE id = ? "

	insertShipmentBox = "INSERT INTO shipment_box (box_number, shipping_method_id, ship_date, " +
		" shipping_address_id, " +
		" shipping_address, shipment_id) " +
		" values(?, ?, ?, ?, ?, ?) "

	updateShipmentBox = "UPDATE shipment_box SET dropshipped = ?, cost = ?, insurance = ?, weight = ?, " +
		" width = ?, height = ?, depth = ?, updated = ?, tracking_number = ? " +
		" WHERE id = ? "

	getShipmentBox = "SELECT id, box_number, dropshipped, cost, insurance, shipping_method_id, " +
		" weight, width, height, depth, ship_date, updated, tracking_number, " +
		" shipping_address_id, shipping_address, shipment_id " +
		" FROM shipment_box " +
		" WHERE id = ? "

	getShipmentBoxList = "SELECT id, box_number, dropshipped, cost, insurance, shipping_method_id, " +
		" weight, width, height, depth, ship_date, updated, tracking_number, " +
		" shipping_address_id, shipping_address, shipment_id " +
		" FROM shipment_box " +
		" WHERE shipment_id = ? "

	deleteShipmentBox = "DELETE FROM shipment_box WHERE id = ? "

	insertShipmentItem = "INSERT INTO shipment_item (order_item_id, quantity, shipment_id, " +
		" updated, shipment_box_id )" +
		" values(?, ?, ?, ?, ?) "

	updateShipmentItem = "UPDATE shipment_item SET quantity = ?, updated = ? " +
		" WHERE id = ? "

	getShipmentItem = "SELECT id, order_item_id, quantity, shipment_id, updated, " +
		" shipment_box_id " +
		" FROM  shipment_item " +
		" WHERE id = ? "

	getShipmentItemListByShipment = "SELECT id, order_item_id, quantity, shipment_id, updated, " +
		" shipment_box_id " +
		" FROM  shipment_item " +
		" WHERE shipment_id = ? "

	getShipmentItemsInBox = "SELECT si.id, si.order_item_id, si.quantity, si.shipment_id, " +
		" si.updated, si.shipment_box_id " +
		" FROM  shipment_item si " +
		" inner join shipment_box sb " +
		" on sb.id = si.shipment_box_id " +
		" WHERE sb.box_number = ? and si.shipment_id = ?"

	deleteShipmentItem = "DELETE FROM shipment_item WHERE id = ? "
)
