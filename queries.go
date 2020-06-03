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

	getStoreByLocalDomain = "SELECT id, company, first_name, last_name, local_domain, remote_domain, " +
		" oauth_client_id, oauth_secret, email, city, state, zip, date_entered, date_updated, " +
		" store_name, store_slogan, logo, currency, enabled " +
		" FROM store " +
		" WHERE local_domain = ? "

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
		" WHERE store_id = ? "

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
)
