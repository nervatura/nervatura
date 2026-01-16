CREATE TABLE IF NOT EXISTS usref(
  id INTEGER,
  refnumber TEXT NOT NULL,
  value TEXT NOT NULL,
  changed TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT usref_refnumber_key UNIQUE (refnumber)
);

CREATE TRIGGER usref_changed_update
  AFTER UPDATE ON usref
  FOR EACH ROW
BEGIN
  UPDATE usref SET changed = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

CREATE TABLE IF NOT EXISTS config(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  config_type TEXT NOT NULL DEFAULT 'CONFIG_MAP', 
  data JSONB NOT NULL,
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT config_code_key UNIQUE (code),
  CHECK( config_type IN ('CONFIG_MAP', 'CONFIG_SHORTCUT', 'CONFIG_MESSAGE', 'CONFIG_PATTERN', 'CONFIG_REPORT', 'CONFIG_PRINT_QUEUE', 'CONFIG_DATA') ) 
);

CREATE INDEX idx_config_type ON config (config_type);

CREATE TRIGGER config_default_code
AFTER INSERT ON config
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE config SET code = 'CNF'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW config_map AS
  SELECT id, code,
    data->>'field_name' AS field_name, data->>'field_type' AS field_type, 
	  data->>'description' AS description, data->'tags' AS tags, data->'filter' AS filter
  FROM config
  WHERE config_type = 'CONFIG_MAP' AND deleted = 0;
  
CREATE VIEW config_shortcut AS
  SELECT id, code,
    data->>'shortcut_key' AS shortcut_key, data->>'description' AS description, 
	  data->>'modul' AS modul, data->>'method' AS method, data->>'func_name' AS func_name,
    data->>'address' AS address, data->'fields' AS fields
  FROM config
  WHERE config_type = 'CONFIG_SHORTCUT' AND deleted = 0;

CREATE VIEW config_message AS
  SELECT id, code,
    data->>'section' AS section, data->>'key' AS message_key, 
	  data->>'lang' AS lang, data->>'value' AS message_value
  FROM config
  WHERE config_type = 'CONFIG_MESSAGE' AND deleted = 0;

CREATE VIEW config_pattern AS
  SELECT id, code,
    data->>'trans_type' AS trans_type, data->>'description' AS description, 
	  data->>'notes' AS notes, data->>'default_pattern' AS default_pattern
  FROM config
  WHERE config_type = 'CONFIG_PATTERN' AND deleted = 0;

CREATE VIEW config_print_queue AS
  SELECT id, code,
    data->>'ref_type' AS ref_type, data->>'ref_code' AS ref_code, 
	  CAST(data->>'qty' AS FLOAT) AS qty, data->>'report_code' AS report_code,
	  data->>'orientation' AS orientation, data->>'paper_size' AS paper_size,
	  data->>'auth_code' AS auth_code, time_stamp
  FROM config
  WHERE config_type = 'CONFIG_PRINT_QUEUE' AND deleted = 0;

CREATE VIEW config_report AS
  SELECT id, code,
    data->>'report_key' AS report_key, data->>'report_type' AS report_type, 
	  data->>'trans_type' AS trans_type, data->>'direction' AS direction,
	  data->>'report_name' AS report_name, data->>'description' AS description,
	  data->>'label' AS label, data->>'file_type' AS file_type,
	  data->'template' AS template
  FROM config
  WHERE config_type = 'CONFIG_REPORT' AND deleted = 0;
  
CREATE VIEW config_data AS
  SELECT ROW_NUMBER() OVER (ORDER BY code, key) as id, code AS config_code,
    key as config_key, value as config_value, type as config_type
  FROM config, json_each(config.data)
  WHERE config_type = 'CONFIG_DATA' AND deleted = 0;

CREATE TABLE IF NOT EXISTS auth(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL', 
  user_name TEXT NOT NULL,
  user_group TEXT NOT NULL DEFAULT 'GROUP_USER',
  disabled BOOLEAN NOT NULL DEFAULT 0,
  auth_meta JSONB NOT NULL DEFAULT (json_object()),
  auth_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT auth_code_key UNIQUE (code),
  CHECK( user_group IN ('GROUP_ADMIN', 'GROUP_USER', 'GROUP_GUEST') )
);

CREATE UNIQUE INDEX idx_auth_user_name ON auth (user_name);
CREATE INDEX idx_auth_deleted ON auth (deleted);
CREATE INDEX idx_auth_disabled ON auth (disabled);
CREATE INDEX idx_auth_tags ON auth (json_extract(auth_meta, '$.tags'));

CREATE TRIGGER auth_default_code
AFTER INSERT ON auth
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE auth SET code = 'USR'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW auth_view AS
  SELECT id, code, user_name, user_group, disabled,
    auth_meta->'tags' AS tags, REPLACE(REPLACE(REPLACE(auth_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst,
    auth_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'user_name', user_name, 'user_group', user_group, 'disabled', 
      json(case when disabled = 1 then 'true' else 'false' end),
      'auth_meta', json(auth_meta), 'auth_map', json(auth_map), 'time_stamp', time_stamp
    ) AS auth_object
  FROM auth
  WHERE deleted = 0;

CREATE VIEW auth_map AS
  SELECT auth.id AS id, auth.code, user_name, user_group,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM auth, json_each(auth.auth_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS currency(
  id INTEGER,
  code TEXT NOT NULL,
  currency_meta JSONB NOT NULL DEFAULT (json_object()),
  currency_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT currency_code_key UNIQUE (code)
);

CREATE INDEX idx_currency_deleted ON currency (deleted);
CREATE INDEX idx_currency_tags ON currency (json_extract(currency_meta, '$.tags'));

CREATE VIEW currency_view AS
  SELECT id, code,
    currency_meta->>'description' AS description, 
    CAST(currency_meta->>'digit' AS INTEGER) AS digit, 
    CAST(currency_meta->>'cash_round' AS INTEGER) AS cash_round, 
    currency_meta->'tags' AS tags, REPLACE(REPLACE(REPLACE(currency_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst,
    currency_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'currency_meta', json(currency_meta), 'currency_map', json(currency_map), 'time_stamp', time_stamp
    ) AS currency_object
  FROM currency
  WHERE deleted = 0;

CREATE VIEW currency_map AS
  SELECT currency.id AS id, currency.code, currency_meta->>'description' AS currency_description,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM currency, json_each(currency.currency_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW currency_tags AS
  SELECT currency.id AS id, code, value as tag
  FROM currency, json_each(currency.currency_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS customer(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  customer_type TEXT NOT NULL DEFAULT 'CUSTOMER_COMPANY',
  customer_name TEXT NOT NULL,
  addresses JSONB NOT NULL DEFAULT (json_array()),
  contacts JSONB NOT NULL DEFAULT (json_array()),
  events JSONB NOT NULL DEFAULT (json_array()), 
  customer_meta JSONB NOT NULL DEFAULT (json_object()),
  customer_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT customer_code_key UNIQUE (code),
  CHECK( customer_type IN ('CUSTOMER_OWN', 'CUSTOMER_COMPANY', 'CUSTOMER_PRIVATE', 'CUSTOMER_OTHER') )
);

CREATE INDEX idx_customer_customer_type ON customer (customer_type);
CREATE INDEX idx_customer_deleted ON customer (deleted);
CREATE INDEX idx_customer_name ON customer (customer_name);
CREATE INDEX idx_customer_inactive ON customer (json_extract(customer_meta, '$.inactive'));
CREATE INDEX idx_customer_tags ON customer (json_extract(customer_meta, '$.tags'));

CREATE TRIGGER customer_default_code
AFTER INSERT ON customer
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE customer SET code = 'CUS'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW customer_view AS
  SELECT id, code, customer_type, customer_name, customer_meta->>'tax_number' AS tax_number, customer_meta->>'account' AS account,
    CAST(customer_meta->>'tax_free' AS BOOLEAN) AS tax_free, CAST(customer_meta->>'terms' AS INTEGER) AS terms, 
    CAST(customer_meta->>'credit_limit' AS FLOAT) AS credit_limit,
    CAST(customer_meta->>'discount' AS FLOAT) AS discount, customer_meta->>'notes' AS notes, customer_meta->>'inactive' AS inactive, 
    customer_meta->'tags' AS tags, REPLACE(REPLACE(REPLACE(customer_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst,
    addresses, contacts, events, customer_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'customer_type', customer_type, 'customer_name', customer_name, 
      'addresses', json(addresses), 'contacts', json(contacts), 'events', json(events), 
      'customer_meta', json(customer_meta), 'customer_map', json(customer_map), 'time_stamp', time_stamp
    ) AS customer_object
  FROM customer
  WHERE deleted = 0;

CREATE VIEW customer_contacts AS
  SELECT customer.id AS id, code, customer_name, 
    value->>'first_name' AS first_name, value->>'surname' AS surname, value->>'status' AS status,
    value->>'phone' AS phone, value->>'mobile' AS mobile, value->>'email' AS email, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'contact_map' AS contact_map
  FROM customer, json_each(customer.contacts)
  WHERE deleted = 0;

CREATE VIEW customer_addresses AS
  SELECT customer.id AS id, code, customer_name, 
    value->>'country' AS country, value->>'state' AS state, value->>'zip_code' AS zip_code,
    value->>'city' AS city, value->>'street' AS street, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'address_map' AS address_map
  FROM customer, json_each(customer.addresses)
  WHERE deleted = 0;

CREATE VIEW customer_events AS
  SELECT customer.id AS id, code, customer_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, value->>'start_time' AS start_time,
    value->>'end_time' AS end_time, value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'event_map' AS event_map
  FROM customer, json_each(customer.events)
  WHERE deleted = 0;

CREATE VIEW customer_map AS
  SELECT customer.id AS id, customer.code, customer.customer_name,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM customer, json_each(customer.customer_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW customer_tags AS
  SELECT customer.id AS id, code, value as tag
  FROM customer, json_each(customer.customer_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS employee(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL', 
  address JSONB NOT NULL DEFAULT (json_object()),
  contact JSONB NOT NULL DEFAULT (json_object()),
  events JSONB NOT NULL DEFAULT (json_array()), 
  employee_meta JSONB NOT NULL DEFAULT (json_object()),
  employee_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT employee_code_key UNIQUE (code)
);

CREATE INDEX idx_employee_deleted ON employee (deleted);
CREATE INDEX idx_employee_name ON employee (json_extract(employee_meta, '$.name'));
CREATE INDEX idx_employee_inactive ON employee (json_extract(employee_meta, '$.inactive'));
CREATE INDEX idx_employee_tags ON employee (json_extract(employee_meta, '$.tags'));

CREATE TRIGGER employee_default_code
AFTER INSERT ON employee
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE employee SET code = 'EMP'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW employee_view AS
  SELECT id, code, employee_meta->>'start_date' AS start_date, employee_meta->>'end_date' AS end_date, 
    CAST(employee_meta->>'inactive' AS BOOLEAN) AS inactive, employee_meta->>'notes' AS notes,
    contact->>'first_name' AS first_name, contact->>'surname' AS surname, contact->>'status' AS status,
    contact->>'phone' AS phone, contact->>'mobile' AS mobile, contact->>'email' AS email, 
    address->>'country' AS country, address->>'state' AS state, address->>'zip_code' AS zip_code,
    address->>'city' AS city, address->>'street' AS street, 
    employee_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(employee_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    employee_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'contact', json(contact), 'address', json(address), 'events', json(events), 
      'employee_meta', json(employee_meta), 'employee_map', json(employee_map), 'time_stamp', time_stamp
    ) AS employee_object
  FROM employee
  WHERE deleted = 0;

CREATE VIEW employee_events AS
  SELECT employee.id AS id, code,
    contact->>'first_name' AS first_name, contact->>'surname' AS surname,
    value->>'uid' AS uid, value->>'subject' AS subject, value->>'start_time' AS start_time,
    value->>'end_time' AS end_time, value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'event_map' AS event_map
  FROM employee, json_each(employee.events)
  WHERE deleted = 0;

CREATE VIEW employee_map AS
  SELECT employee.id AS id, employee.code, 
    contact->>'first_name' AS first_name, contact->>'surname' AS surname,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM employee, json_each(employee.employee_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW employee_tags AS
  SELECT employee.id AS id, code, value as tag
  FROM employee, json_each(employee.employee_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS place(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  place_type TEXT NOT NULL DEFAULT 'PLACE_WAREHOUSE',
  place_name TEXT NOT NULL,
  currency_code TEXT,
  address JSONB NOT NULL DEFAULT (json_object()),
  contacts JSONB NOT NULL DEFAULT (json_array()),
  events JSONB NOT NULL DEFAULT (json_array()),
  place_meta JSONB NOT NULL DEFAULT (json_object()),
  place_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT place_code_key UNIQUE (code),
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  CHECK( place_type IN ('PLACE_BANK', 'PLACE_CASH', 'PLACE_WAREHOUSE', 'PLACE_OTHER') )
);

CREATE INDEX idx_place_deleted ON place (deleted);
CREATE INDEX idx_place_place_name ON place (place_name);
CREATE INDEX idx_place_inactive ON place (json_extract(place_meta, '$.inactive'));
CREATE INDEX idx_place_tags ON place (json_extract(place_meta, '$.tags'));

CREATE TRIGGER place_default_code
AFTER INSERT ON place
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE place SET code = 'PLA'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW place_view AS
  SELECT id, code, place_type, place_name, currency_code,
    CAST(place_meta->>'inactive' AS BOOLEAN) AS inactive, place_meta->>'notes' AS notes,
    address->>'country' AS country, address->>'state' AS state, address->>'zip_code' AS zip_code,
    address->>'city' AS city, address->>'street' AS street, 
    place_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(place_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    place_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'place_type', place_type, 'place_name', place_name, 'currency_code', currency_code,
      'address', json(address), 'contacts', json(contacts), 'events', json(events), 'place_meta', json(place_meta), 
      'place_map', json(place_map), 'time_stamp', time_stamp
    ) AS place_object
  FROM place
  WHERE deleted = 0;

CREATE VIEW place_contacts AS
  SELECT place.id AS id, code, place_name, 
    value->>'first_name' AS first_name, value->>'surname' AS surname, value->>'status' AS status,
    value->>'phone' AS phone, value->>'mobile' AS mobile, value->>'email' AS email, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'contact_map' AS contact_map
  FROM place, json_each(place.contacts)
  WHERE deleted = 0;

CREATE VIEW place_events AS
  SELECT place.id AS id, code, place_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, value->>'start_time' AS start_time,
    value->>'end_time' AS end_time, value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'event_map' AS event_map
  FROM place, json_each(place.events)
  WHERE deleted = 0;

CREATE VIEW place_map AS
  SELECT place.id AS id, place.code, place.place_name,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM place, json_each(place.place_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW place_tags AS
  SELECT place.id AS id, code, value as tag
  FROM place, json_each(place.place_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS tax(
  id INTEGER,
  code TEXT NOT NULL, 
  tax_meta JSONB NOT NULL DEFAULT (json_object()),
  tax_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT tax_code_key UNIQUE (code)
);

CREATE INDEX idx_tax_deleted ON tax (deleted);
CREATE INDEX idx_tax_inactive ON tax (json_extract(tax_meta, '$.inactive'));
CREATE INDEX idx_tax_tags ON tax (json_extract(tax_meta, '$.tags'));

CREATE VIEW tax_view AS
  SELECT id, code,
    tax_meta->>'description' AS description, CAST(tax_meta->>'rate_value' AS FLOAT) AS rate_value, 
    CAST(tax_meta->>'inactive' AS BOOLEAN) AS inactive, tax_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(tax_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    tax_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'tax_meta', json(tax_meta), 'tax_map', json(tax_map), 'time_stamp', time_stamp
    ) AS tax_object
  FROM tax
  WHERE deleted = 0;

CREATE VIEW tax_map AS
  SELECT tax.id AS id, tax.code, tax_meta->>'description' AS tax_description,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM tax, json_each(tax.tax_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW tax_tags AS
  SELECT tax.id AS id, code, value as tag
  FROM tax, json_each(tax.tax_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS link(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  link_type_1 TEXT NOT NULL DEFAULT 'LINK_TRANS',
  link_code_1 TEXT NOT NULL,
  link_type_2 TEXT NOT NULL DEFAULT 'LINK_TRANS',
  link_code_2 TEXT NOT NULL, 
  link_meta JSONB NOT NULL DEFAULT (json_object()),
  link_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT link_code_key UNIQUE (code),
  CHECK( link_type_1 IN (
    'LINK_CUSTOMER', 'LINK_EMPLOYEE', 'LINK_ITEM', 'LINK_MOVEMENT', 'LINK_PAYMENT', 'LINK_PLACE', 'LINK_PRODUCT', 'LINK_PROJECT', 'LINK_TOOL', 'LINK_TRANS'
  ) ),
  CHECK( link_type_2 IN (
    'LINK_CUSTOMER', 'LINK_EMPLOYEE', 'LINK_ITEM', 'LINK_MOVEMENT', 'LINK_PAYMENT', 'LINK_PLACE', 'LINK_PRODUCT', 'LINK_PROJECT', 'LINK_TOOL', 'LINK_TRANS'
  ) )
);

CREATE INDEX idx_link_link_code_1 ON link (link_type_1, link_code_1);
CREATE INDEX idx_link_link_code_2 ON link (link_type_2, link_code_2);
CREATE INDEX idx_link_deleted ON link (deleted);
CREATE INDEX idx_link_tags ON link (json_extract(link_meta, '$.tags'));

CREATE TRIGGER link_default_code
AFTER INSERT ON link
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE link SET code = 'LNK'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE TRIGGER link_insert
BEFORE INSERT ON link
FOR EACH ROW
BEGIN
  SELECT CASE 
    WHEN NEW.link_type_1 = 'LINK_CUSTOMER' AND (SELECT id FROM customer WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid customer code')
    WHEN NEW.link_type_2 = 'LINK_CUSTOMER' AND (SELECT id FROM customer WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid customer code')
    WHEN NEW.link_type_1 = 'LINK_EMPLOYEE' AND (SELECT id FROM employee WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid employee code')
    WHEN NEW.link_type_2 = 'LINK_EMPLOYEE' AND (SELECT id FROM employee WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid employee code')
    WHEN NEW.link_type_1 = 'LINK_ITEM' AND (SELECT id FROM item WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid item code')
    WHEN NEW.link_type_2 = 'LINK_ITEM' AND (SELECT id FROM item WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid item code')
    WHEN NEW.link_type_1 = 'LINK_MOVEMENT' AND (SELECT id FROM movement WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid movement code')
    WHEN NEW.link_type_2 = 'LINK_MOVEMENT' AND (SELECT id FROM movement WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid movement code')
    WHEN NEW.link_type_1 = 'LINK_PAYMENT' AND (SELECT id FROM payment WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid payment code')
    WHEN NEW.link_type_2 = 'LINK_PAYMENT' AND (SELECT id FROM payment WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid payment code')
    WHEN NEW.link_type_1 = 'LINK_PLACE' AND (SELECT id FROM place WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid place code')
    WHEN NEW.link_type_2 = 'LINK_PLACE' AND (SELECT id FROM place WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid place code')
    WHEN NEW.link_type_1 = 'LINK_PRODUCT' AND (SELECT id FROM product WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid product code')
    WHEN NEW.link_type_2 = 'LINK_PRODUCT' AND (SELECT id FROM product WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid product code')
    WHEN NEW.link_type_1 = 'LINK_PROJECT' AND (SELECT id FROM project WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid project code')
    WHEN NEW.link_type_2 = 'LINK_PROJECT' AND (SELECT id FROM project WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid project code')
    WHEN NEW.link_type_1 = 'LINK_TOOL' AND (SELECT id FROM tool WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid tool code')
    WHEN NEW.link_type_2 = 'LINK_TOOL' AND (SELECT id FROM tool WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid tool code')
    WHEN NEW.link_type_1 = 'LINK_TRANS' AND (SELECT id FROM trans WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid trans code')
    WHEN NEW.link_type_2 = 'LINK_TRANS' AND (SELECT id FROM trans WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid trans code')
  END;
END;

CREATE TRIGGER link_update
BEFORE UPDATE ON link
FOR EACH ROW
BEGIN
  SELECT CASE 
    WHEN NEW.link_type_1 = 'LINK_CUSTOMER' AND (SELECT id FROM customer WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid customer code')
    WHEN NEW.link_type_2 = 'LINK_CUSTOMER' AND (SELECT id FROM customer WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid customer code')
    WHEN NEW.link_type_1 = 'LINK_EMPLOYEE' AND (SELECT id FROM employee WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid employee code')
    WHEN NEW.link_type_2 = 'LINK_EMPLOYEE' AND (SELECT id FROM employee WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid employee code')
    WHEN NEW.link_type_1 = 'LINK_ITEM' AND (SELECT id FROM item WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid item code')
    WHEN NEW.link_type_2 = 'LINK_ITEM' AND (SELECT id FROM item WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid item code')
    WHEN NEW.link_type_1 = 'LINK_MOVEMENT' AND (SELECT id FROM movement WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid movement code')
    WHEN NEW.link_type_2 = 'LINK_MOVEMENT' AND (SELECT id FROM movement WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid movement code')
    WHEN NEW.link_type_1 = 'LINK_PAYMENT' AND (SELECT id FROM payment WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid payment code')
    WHEN NEW.link_type_2 = 'LINK_PAYMENT' AND (SELECT id FROM payment WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid payment code')
    WHEN NEW.link_type_1 = 'LINK_PLACE' AND (SELECT id FROM place WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid place code')
    WHEN NEW.link_type_2 = 'LINK_PLACE' AND (SELECT id FROM place WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid place code')
    WHEN NEW.link_type_1 = 'LINK_PRODUCT' AND (SELECT id FROM product WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid product code')
    WHEN NEW.link_type_2 = 'LINK_PRODUCT' AND (SELECT id FROM product WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid product code')
    WHEN NEW.link_type_1 = 'LINK_PROJECT' AND (SELECT id FROM project WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid project code')
    WHEN NEW.link_type_2 = 'LINK_PROJECT' AND (SELECT id FROM project WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid project code')
    WHEN NEW.link_type_1 = 'LINK_TOOL' AND (SELECT id FROM tool WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid tool code')
    WHEN NEW.link_type_2 = 'LINK_TOOL' AND (SELECT id FROM tool WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid tool code')
    WHEN NEW.link_type_1 = 'LINK_TRANS' AND (SELECT id FROM trans WHERE code = NEW.link_code_1) IS NULL THEN
      RAISE(ABORT, 'Invalid trans code')
    WHEN NEW.link_type_2 = 'LINK_TRANS' AND (SELECT id FROM trans WHERE code = NEW.link_code_2) IS NULL THEN
      RAISE(ABORT, 'Invalid trans code')
  END;
END;

CREATE VIEW link_view AS
  SELECT id, code, 
    link_type_1, link_code_1, link_type_2, link_code_2, 
    CAST(link_meta->>'qty' AS FLOAT) AS qty, CAST(link_meta->>'amount' AS FLOAT) AS amount, CAST(link_meta->>'rate' AS FLOAT) AS rate,
    COALESCE(link_meta->>'notes', '') AS notes,
    link_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(link_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    link_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'link_type_1', link_type_1, 'link_code_1', link_code_1, 'link_type_2', link_type_2,
      'link_code_2', link_code_2, 'link_meta', json(link_meta), 'link_map', json(link_map), 'time_stamp', time_stamp
    ) AS link_object
  FROM link
  WHERE deleted = 0;

CREATE VIEW link_map AS
  SELECT link.id AS id, link.code, 
    link_type_1, link_code_1, link_type_2, link_code_2,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM link, json_each(link.link_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW link_tags AS
  SELECT link.id AS id, code, value as tag
  FROM link, json_each(link.link_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS product(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  product_type TEXT NOT NULL DEFAULT 'PRODUCT_ITEM',
  product_name TEXT NOT NULL,
  tax_code TEXT NOT NULL,
  events JSONB NOT NULL DEFAULT (json_array()),
  product_meta JSONB NOT NULL DEFAULT (json_object()),
  product_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT product_code_key UNIQUE (code),
  FOREIGN KEY (tax_code) REFERENCES tax(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  CHECK( product_type IN ('PRODUCT_ITEM', 'PRODUCT_SERVICE', 'PRODUCT_VIRTUAL') )
);

CREATE INDEX idx_product_product_type ON product (product_type);
CREATE INDEX idx_product_tax_code ON product (tax_code);
CREATE INDEX idx_product_deleted ON product (deleted);
CREATE INDEX idx_product_name ON product (product_name);
CREATE INDEX idx_product_barcode_type ON product (product_meta->>'barcode');
CREATE INDEX idx_product_inactive ON product (json_extract(product_meta, '$.inactive'));
CREATE INDEX idx_product_tags ON product (json_extract(product_meta, '$.tags'));

CREATE TRIGGER product_default_code
AFTER INSERT ON product
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE product SET code = 'PRD'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW product_view AS
  SELECT id, code, product_type, product_name, tax_code,
    product_meta->>'unit' AS unit, CAST(product_meta->>'inactive' AS BOOLEAN) AS inactive, product_meta->>'notes' AS notes,
    product_meta->>'barcode_type' AS barcode_type, product_meta->>'barcode' AS barcode, 
    CAST(product_meta->>'barcode_qty' AS FLOAT) AS barcode_qty,
    product_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(product_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    events, product_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'product_type', product_type, 'product_name', product_name, 'tax_code', tax_code,
      'events', json(events), 'product_meta', json(product_meta), 'product_map', json(product_map), 'time_stamp', time_stamp
    ) AS product_object
  FROM product
  WHERE deleted = 0;

CREATE VIEW product_events AS
  SELECT product.id AS id, code, product_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, value->>'start_time' AS start_time,
    value->>'end_time' AS end_time, value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'event_map' AS event_map
  FROM product, json_each(product.events)
  WHERE deleted = 0;

CREATE VIEW product_map AS
  SELECT product.id AS id, product.code, product.product_name,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM product, json_each(product.product_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW product_tags AS
  SELECT product.id AS id, code, value as tag
  FROM product, json_each(product.product_meta->'tags')
  WHERE deleted = 0;

CREATE VIEW product_components AS
  SELECT p.id, p.code AS product_code, p.product_name, COALESCE(p.product_meta->>'unit','') as unit, 
  c.code as ref_product_code, c.product_name as component_name, COALESCE(c.product_meta->>'unit','') as component_unit, 
  c.product_type as component_type,
  CAST(l.link_meta->>'qty' AS FLOAT) AS qty, COALESCE(l.link_meta->>'notes', '') AS notes
  FROM product p INNER JOIN link l ON l.link_code_1 = p.code
  INNER JOIN product c ON l.link_code_2 = c.code
  WHERE p.product_type = 'PRODUCT_VIRTUAL' AND link_type_1 = 'LINK_PRODUCT' AND link_type_2 = 'LINK_PRODUCT' 
  AND p.deleted = false AND l.deleted = false AND c.deleted = false;

CREATE TABLE IF NOT EXISTS project(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  project_name TEXT NOT NULL,
  customer_code TEXT,
  addresses JSONB NOT NULL DEFAULT (json_array()),
  contacts JSONB NOT NULL DEFAULT (json_array()),
  events JSONB NOT NULL DEFAULT (json_array()),
  project_meta JSONB NOT NULL DEFAULT (json_object()),
  project_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT project_code_key UNIQUE (code),
  FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_project_customer_code ON project (customer_code);
CREATE INDEX idx_project_deleted ON project (deleted);
CREATE INDEX idx_project_project_name ON project (project_name);
CREATE INDEX idx_project_inactive ON project (json_extract(project_meta, '$.inactive'));
CREATE INDEX idx_project_tags ON project (json_extract(project_meta, '$.tags'));

CREATE TRIGGER project_default_code
AFTER INSERT ON project
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE project SET code = 'PRJ'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW project_view AS
  SELECT id, code, project_name, customer_code,
    project_meta->>'start_date' AS start_date, project_meta->>'end_date' AS end_date, 
    CAST(project_meta->>'inactive' AS BOOLEAN) AS inactive, project_meta->>'notes' AS notes, 
    project_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(project_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    addresses, contacts, events, project_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'project_name', project_name, 'customer_code', customer_code,
      'addresses', json(addresses), 'contacts', json(contacts), 'events', json(events), 
      'project_meta', json(project_meta), 'project_map', json(project_map), 'time_stamp', time_stamp
    ) AS project_object
  FROM project
  WHERE deleted = 0;

CREATE VIEW project_contacts AS
  SELECT project.id AS id, code, project_name, 
    value->>'first_name' AS first_name, value->>'surname' AS surname, value->>'status' AS status,
    value->>'phone' AS phone, value->>'mobile' AS mobile, value->>'email' AS email, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'contact_map' AS contact_map
  FROM project, json_each(project.contacts)
  WHERE deleted = 0;

CREATE VIEW project_addresses AS
  SELECT project.id AS id, code, project_name, 
    value->>'country' AS country, value->>'state' AS state, value->>'zip_code' AS zip_code,
    value->>'city' AS city, value->>'street' AS street, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'address_map' AS address_map
  FROM project, json_each(project.addresses)
  WHERE deleted = 0;

CREATE VIEW project_events AS
  SELECT project.id AS id, code, project_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, value->>'start_time' AS start_time,
    value->>'end_time' AS end_time, value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'event_map' AS event_map
  FROM project, json_each(project.events)
  WHERE deleted = 0;

CREATE VIEW project_map AS
  SELECT project.id AS id, project.code, project.project_name,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM project, json_each(project.project_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW project_tags AS
  SELECT project.id AS id, code, value as tag
  FROM project, json_each(project.project_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS rate(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  rate_type TEXT NOT NULL DEFAULT 'RATE_RATE',
  rate_date TEXT NOT NULL DEFAULT CURRENT_DATE,
  place_code TEXT,
  currency_code TEXT NOT NULL, 
  rate_meta JSONB NOT NULL DEFAULT (json_object()),
  rate_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT rate_code_key UNIQUE (code),
  FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  CHECK( rate_type IN ('RATE_RATE', 'RATE_BUY', 'RATE_SELL', 'RATE_AVERAGE') )
);

CREATE UNIQUE INDEX idx_rate_tdcp ON rate (rate_type, rate_date, currency_code, place_code);
CREATE INDEX idx_rate_rate_type ON rate (rate_type);
CREATE INDEX idx_rate_rate_date ON rate (rate_date);
CREATE INDEX idx_rate_place_code ON rate (place_code);
CREATE INDEX idx_rate_currency_code ON rate (currency_code);
CREATE INDEX idx_rate_deleted ON rate (deleted);
CREATE INDEX idx_rate_tags ON rate (json_extract(rate_meta, '$.tags'));

CREATE TRIGGER rate_default_code
AFTER INSERT ON rate
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE rate SET code = 'RAT'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW rate_view AS
  SELECT id, code, rate_type, rate_date, place_code, currency_code,
    CAST(rate_meta->>'rate_value' AS FLOAT) AS rate_value, 
    rate_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(rate_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    rate_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'rate_type', rate_type, 'rate_date', rate_date, 'place_code', place_code,
      'currency_code', currency_code, 'rate_meta', json(rate_meta), 'rate_map', json(rate_map), 'time_stamp', time_stamp
    ) AS rate_object
  FROM rate
  WHERE deleted = 0;

CREATE VIEW rate_map AS
  SELECT rate.id AS id, rate.code,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM rate, json_each(rate.rate_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW rate_tags AS
  SELECT rate.id AS id, code, value as tag
  FROM rate, json_each(rate.rate_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS tool(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  description TEXT NOT NULL,
  product_code TEXT NOT NULL,
  events JSONB NOT NULL DEFAULT (json_array()),
  tool_meta JSONB NOT NULL DEFAULT (json_object()),
  tool_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT tool_code_key UNIQUE (code),
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_tool_product_code ON tool (product_code);
CREATE INDEX idx_tool_deleted ON tool (deleted);
CREATE INDEX idx_tool_description ON tool (description);
CREATE INDEX idx_tool_inactive ON tool (json_extract(tool_meta, '$.inactive'));
CREATE INDEX idx_tool_tags ON tool (json_extract(tool_meta, '$.tags'));

CREATE TRIGGER tool_default_code
AFTER INSERT ON tool
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE tool SET code = 'SER'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW tool_view AS
  SELECT id, code, product_code, description,
    tool_meta->>'serial_number' AS serial_number,
    CAST(tool_meta->>'inactive' AS BOOLEAN) AS inactive, tool_meta->>'notes' AS notes, 
    tool_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(tool_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    tool_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'description', description, 'product_code', product_code,
      'events', json(events), 'tool_meta', json(tool_meta), 'tool_map', json(tool_map), 'time_stamp', time_stamp
    ) AS tool_object
  FROM tool
  WHERE deleted = 0;

CREATE VIEW tool_events AS
  SELECT tool.id AS id, code, description as tool_description, 
    value->>'uid' AS uid, value->>'subject' AS subject, value->>'start_time' AS start_time,
    value->>'end_time' AS end_time, value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(value->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    value->'event_map' AS event_map
  FROM tool, json_each(tool.events)
  WHERE deleted = 0;

CREATE VIEW tool_map AS
  SELECT tool.id AS id, tool.code, tool.description as tool_description,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM tool, json_each(tool.tool_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW tool_tags AS
  SELECT tool.id AS id, code, value as tag
  FROM tool, json_each(tool.tool_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS price(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  price_type TEXT NOT NULL DEFAULT 'PRICE_CUSTOMER',
  valid_from TEXT NOT NULL,
  valid_to TEXT,
  product_code TEXT NOT NULL,
  currency_code TEXT NOT NULL,
  customer_code TEXT,
  qty INTEGER NOT NULL DEFAULT 0,
  price_meta JSONB NOT NULL DEFAULT (json_object()),
  price_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT price_code_key UNIQUE (code),
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  CHECK( price_type IN ('PRICE_CUSTOMER', 'PRICE_VENDOR') )
);

CREATE UNIQUE INDEX idx_price_pvpc ON price (price_type, valid_from, product_code, currency_code, qty);
CREATE INDEX idx_price_price_type ON price (price_type);
CREATE INDEX idx_price_valid_from ON price (valid_from);
CREATE INDEX idx_price_valid_to ON price (valid_to);
CREATE INDEX idx_price_product_code ON price (product_code);
CREATE INDEX idx_price_currency_code ON price (currency_code);
CREATE INDEX idx_price_customer_code ON price (customer_code);
CREATE INDEX idx_price_deleted ON price (deleted);
CREATE INDEX idx_price_tags ON price (json_extract(price_meta, '$.tags'));

CREATE TRIGGER price_default_code
AFTER INSERT ON price
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE price SET code = 'PRC'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW price_view AS
  SELECT id, code, price_type, valid_from, valid_to, product_code, currency_code, customer_code, qty,
    CAST(price_meta->>'price_value' AS FLOAT) AS price_value, 
    price_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(price_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    price_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'price_type', price_type, 'valid_from', valid_from, 'valid_to', valid_to,
      'product_code', product_code, 'currency_code', currency_code, 'customer_code', customer_code, 'qty', qty, 
      'price_meta', json(price_meta), 'price_map', json(price_map), 'time_stamp', time_stamp
    ) AS price_object
  FROM price
  WHERE deleted = 0;

CREATE VIEW price_map AS
  SELECT price.id AS id, price.code,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM price, json_each(price.price_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW price_tags AS
  SELECT price.id AS id, code, value as tag
  FROM price, json_each(price.price_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS trans(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  trans_type TEXT NOT NULL DEFAULT 'TRANS_INVOICE',
  direction TEXT NOT NULL DEFAULT 'DIRECTION_OUT',
  trans_date TEXT NOT NULL,
  trans_code TEXT,
  customer_code TEXT,
  employee_code TEXT,
  project_code TEXT,
  place_code TEXT,
  currency_code TEXT,
  auth_code TEXT NOT NULL, 
  trans_meta JSONB NOT NULL DEFAULT (json_object()),
  trans_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT trans_code_key UNIQUE (code),
  FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE RESTRICT ON DELETE NO ACTION,
  FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (employee_code) REFERENCES employee(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (project_code) REFERENCES project(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (auth_code) REFERENCES auth(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  CHECK( trans_type IN (
    'TRANS_INVOICE', 'TRANS_RECEIPT', 'TRANS_ORDER', 'TRANS_OFFER', 'TRANS_WORKSHEET', 'TRANS_RENT', 'TRANS_DELIVERY', 'TRANS_INVENTORY', 
    'TRANS_WAYBILL', 'TRANS_PRODUCTION', 'TRANS_FORMULA', 'TRANS_BANK', 'TRANS_CASH') )
  CHECK( direction IN ('DIRECTION_OUT', 'DIRECTION_IN', 'DIRECTION_TRANSFER') )
);

CREATE INDEX idx_trans_trans_type ON trans (trans_type);
CREATE INDEX idx_trans_direction ON trans (direction);
CREATE INDEX idx_trans_trans_date ON trans (trans_date);
CREATE INDEX idx_trans_trans_code ON trans (trans_code);
CREATE INDEX idx_trans_customer_code ON trans (customer_code);
CREATE INDEX idx_trans_employee_code ON trans (employee_code);
CREATE INDEX idx_trans_project_code ON trans (project_code);
CREATE INDEX idx_trans_place_code ON trans (place_code);
CREATE INDEX idx_trans_currency_code ON trans (currency_code);
CREATE INDEX idx_trans_auth_code ON trans (auth_code);
CREATE INDEX idx_trans_due_time ON trans (json_extract(trans_meta, '$.due_time'));
CREATE INDEX idx_trans_paid_type ON trans (json_extract(trans_meta, '$.paid_type'));
CREATE INDEX idx_trans_trans_state ON trans (json_extract(trans_meta, '$.trans_state'));
CREATE INDEX idx_trans_tags ON trans (json_extract(trans_meta, '$.tags'));

CREATE TRIGGER trans_default_code_inventory
AFTER INSERT ON trans
FOR EACH ROW
WHEN NEW.code = 'NULL' AND NEW.trans_type = 'TRANS_INVENTORY'
BEGIN
  UPDATE trans SET code = 'COR'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE TRIGGER trans_default_code_delivery
AFTER INSERT ON trans
FOR EACH ROW
WHEN NEW.code = 'NULL' AND NEW.trans_type = 'TRANS_DELIVERY'
BEGIN
  UPDATE trans SET code = CASE WHEN NEW.direction = 'DIRECTION_TRANSFER' THEN 'TRF' ELSE 'DEL' END||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE TRIGGER trans_default_code
  AFTER INSERT ON trans
  FOR EACH ROW
WHEN NEW.code = 'NULL' AND NEW.trans_type <> 'TRANS_INVENTORY' AND NEW.trans_type <> 'TRANS_DELIVERY'
BEGIN
  UPDATE trans SET code = substr(NEW.trans_type, 7, 3)||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE TRIGGER trans_invoice_customer_insert
  AFTER INSERT ON trans
  FOR EACH ROW
WHEN ((NEW.trans_type = 'TRANS_INVOICE') AND (NEW.direction = 'DIRECTION_OUT'))
BEGIN
  UPDATE trans SET trans_meta = json_set(trans_meta, '$.invoice', json_extract(sl.idata, '$'))
  FROM (SELECT json_object(
      'customer_name', cu.customer_name, 'customer_tax_number', cu.customer_meta->>'tax_number', 
      'customer_account', cu.customer_meta->>'account',
      'customer_address', trim(coalesce(json_extract(cu.addresses, '$[0]')->>'zip_code','')||' '||
        coalesce(json_extract(cu.addresses, '$[0]')->>'city','')||' '||
        coalesce(json_extract(cu.addresses, '$[0]')->>'street','')),
      'company_name', co.customer_name, 'company_tax_number', co.customer_meta->>'tax_number',
      'company_account', co.customer_meta->>'account',
      'company_address', trim(coalesce(json_extract(co.addresses, '$[0]')->>'zip_code','')||' '||
        coalesce(json_extract(co.addresses, '$[0]')->>'city','')||' '||
        coalesce(json_extract(co.addresses, '$[0]')->>'street',''))
    ) idata FROM customer cu INNER JOIN customer co ON co.customer_type = 'CUSTOMER_OWN'
    WHERE cu.code = NEW.customer_code LIMIT 1) sl
  WHERE id = NEW.id;
END;

CREATE TRIGGER trans_invoice_customer_update
  AFTER UPDATE ON trans
  FOR EACH ROW
WHEN ((NEW.trans_type = 'TRANS_INVOICE') AND (NEW.direction = 'DIRECTION_OUT') AND (OLD.customer_code <> NEW.customer_code))
BEGIN
  UPDATE trans SET trans_meta = json_set(trans_meta, '$.invoice', json_extract(sl.idata, '$'))
  FROM (SELECT json_object(
      'customer_name', cu.customer_name, 'customer_tax_number', cu.customer_meta->>'tax_number', 
      'customer_account', cu.customer_meta->>'account',
      'customer_address', trim(coalesce(json_extract(cu.addresses, '$[0]')->>'zip_code','')||' '||
        coalesce(json_extract(cu.addresses, '$[0]')->>'city','')||' '||
        coalesce(json_extract(cu.addresses, '$[0]')->>'street','')),
      'company_name', co.customer_name, 'company_tax_number', co.customer_meta->>'tax_number',
      'company_account', co.customer_meta->>'account',
      'company_address', trim(coalesce(json_extract(co.addresses, '$[0]')->>'zip_code','')||' '||
        coalesce(json_extract(co.addresses, '$[0]')->>'city','')||' '||
        coalesce(json_extract(co.addresses, '$[0]')->>'street',''))
    ) idata FROM customer cu INNER JOIN customer co ON co.customer_type = 'CUSTOMER_OWN'
    WHERE cu.code = NEW.customer_code LIMIT 1) sl
  WHERE id = NEW.id;
END;

CREATE VIEW trans_view AS
  SELECT id, code, trans_type, direction, trans_date, trans_code, customer_code, employee_code,
    project_code, place_code, currency_code, auth_code,
    trans_meta->>'due_time' AS due_time, trans_meta->>'ref_number' AS ref_number,
    trans_meta->>'paid_type' AS paid_type, CAST(trans_meta->>'tax_free' AS BOOLEAN) AS tax_free,
    CAST(trans_meta->>'paid' AS BOOLEAN) AS paid, CAST(trans_meta->>'rate' AS FLOAT) AS rate,
    CAST(trans_meta->>'closed' AS BOOLEAN) AS closed, 
    trans_meta->>'status' AS status,
    trans_meta->>'trans_state' AS trans_state,
    trans_meta->>'notes' AS notes, trans_meta->>'internal_notes' AS internal_notes,
    trans_meta->>'report_notes' AS report_notes,
    CASE json_type(trans_meta->'worksheet') WHEN 'object' THEN CAST(trans_meta->'worksheet'->>'distance' AS FLOAT) ELSE 0 END AS worksheet_distance,
    CASE json_type(trans_meta->'worksheet') WHEN 'object' THEN CAST(trans_meta->'worksheet'->>'repair' AS FLOAT) ELSE 0 END AS worksheet_repair,
    CASE json_type(trans_meta->'worksheet') WHEN 'object' THEN CAST(trans_meta->'worksheet'->>'total' AS FLOAT) ELSE 0 END AS worksheet_total,
    CASE json_type(trans_meta->'worksheet') WHEN 'object' THEN trans_meta->'worksheet'->>'justification' ELSE '' END AS worksheet_justification,
    CASE json_type(trans_meta->'rent') WHEN 'object' THEN CAST(trans_meta->'rent'->>'holiday' AS FLOAT) ELSE 0 END AS rent_holiday,
    CASE json_type(trans_meta->'rent') WHEN 'object' THEN CAST(trans_meta->'rent'->>'bad_tool' AS FLOAT) ELSE 0 END AS rent_bad_tool,
    CASE json_type(trans_meta->'rent') WHEN 'object' THEN CAST(trans_meta->'rent'->>'other' AS FLOAT) ELSE 0 END AS rent_other,
    CASE json_type(trans_meta->'rent') WHEN 'object' THEN trans_meta->'rent'->>'justification' END AS rent_justification,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'company_name' END AS invoice_company_name,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'company_account' END AS invoice_company_account,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'company_address' END AS invoice_company_address,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'company_tax_number' END AS invoice_company_tax_number,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'customer_name' END AS invoice_customer_name,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'customer_account' END AS invoice_customer_account,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'customer_address' END AS invoice_customer_address,
    CASE json_type(trans_meta->'invoice') WHEN 'object' THEN trans_meta->'invoice'->>'customer_tax_number' END AS invoice_customer_tax_number,
    trans_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(trans_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    trans_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'trans_type', trans_type, 'direction', direction, 'trans_date', trans_date,
      'trans_code', trans_code, 'customer_code', customer_code, 'employee_code', employee_code, 'project_code', project_code,
      'place_code', place_code, 'currency_code', currency_code, 'auth_code', auth_code, 'trans_meta', json(trans_meta),
      'trans_map', json(trans_map), 'time_stamp', time_stamp
    ) AS trans_object
  FROM trans
  WHERE deleted = 0;

CREATE VIEW trans_map AS
  SELECT trans.id AS id, trans.code, trans.trans_type, trans.direction, trans.trans_date,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM trans, json_each(trans.trans_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0;

CREATE VIEW trans_tags AS
  SELECT trans.id AS id, code, value as tag
  FROM trans, json_each(trans.trans_meta->'tags')
  WHERE deleted = 0;

CREATE TABLE IF NOT EXISTS item(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  trans_code TEXT NOT NULL,
  product_code TEXT NOT NULL,
  tax_code TEXT NOT NULL, 
  item_meta JSONB NOT NULL DEFAULT (json_object()),
  item_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT item_code_key UNIQUE (code),
  FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (tax_code) REFERENCES tax(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_item_trans_code ON item (trans_code);
CREATE INDEX idx_item_product_code ON item (product_code);
CREATE INDEX idx_item_tax_code ON item (tax_code);
CREATE INDEX idx_item_deleted ON item (deleted);
CREATE INDEX idx_item_tags ON item (json_extract(item_meta, '$.tags'));

CREATE TRIGGER item_default_code
AFTER INSERT ON item
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE item SET code = 'ITM'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW item_view AS
  SELECT id, code, trans_code, product_code, tax_code,
    item_meta->>'unit' AS unit, CAST(item_meta->>'qty' AS FLOAT) AS qty, CAST(item_meta->>'fx_price' AS FLOAT) AS fx_price,
    CAST(item_meta->>'net_amount' AS FLOAT) AS net_amount, CAST(item_meta->>'discount' AS FLOAT) AS discount, CAST(item_meta->>'vat_amount' AS FLOAT) AS vat_amount,
    CAST(item_meta->>'amount' AS FLOAT) AS amount, item_meta->>'description' AS description, CAST(item_meta->>'deposit' AS BOOLEAN) AS deposit,
    CAST(item_meta->>'own_stock' AS FLOAT) AS own_stock, CAST(item_meta->>'action_price' AS BOOLEAN) AS action_price, 
    item_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(item_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    item_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'trans_code', trans_code, 'product_code', product_code, 'tax_code', tax_code,
      'item_meta', json(item_meta), 'item_map', json(item_map), 'time_stamp', time_stamp
    ) AS item_object
  FROM item
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW item_map AS
  SELECT item.id AS id, item.code, item.trans_code, item.product_code,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM item, json_each(item.item_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW item_tags AS
  SELECT item.id AS id, code, value as tag
  FROM item, json_each(item.item_meta->'tags')
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE TABLE IF NOT EXISTS movement(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  movement_type TEXT NOT NULL DEFAULT 'MOVEMENT_INVENTORY',
  shipping_time TEXT NOT NULL,
  trans_code TEXT NOT NULL,
  product_code TEXT,
  tool_code TEXT,
  place_code TEXT, 
  item_code TEXT,
  movement_code TEXT,
  movement_meta JSONB NOT NULL DEFAULT (json_object()),
  movement_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT movement_code_key UNIQUE (code),
  FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (tool_code) REFERENCES tool(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (item_code) REFERENCES item(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (movement_code) REFERENCES movement(code)
    ON UPDATE RESTRICT ON DELETE SET NULL,
  CHECK( movement_type IN ('MOVEMENT_INVENTORY', 'MOVEMENT_TOOL', 'MOVEMENT_PLAN', 'MOVEMENT_HEAD') )
);

CREATE INDEX idx_movement_trans_code ON movement (trans_code);
CREATE INDEX idx_movement_product_code ON movement (product_code);
CREATE INDEX idx_movement_tool_code ON movement (tool_code);
CREATE INDEX idx_movement_place_code ON movement (place_code);
CREATE INDEX idx_movement_item_code ON movement (item_code);
CREATE INDEX idx_movement_movement_code ON movement (movement_code);
CREATE INDEX idx_movement_deleted ON movement (deleted);
CREATE INDEX idx_movement_tags ON movement (json_extract(movement_meta, '$.tags'));

CREATE TRIGGER movement_default_code
AFTER INSERT ON movement
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE movement SET code = 'MOV'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW movement_view AS
  SELECT id, code, movement_type, shipping_time,
    trans_code, product_code, tool_code, place_code, item_code, movement_code,
    CAST(movement_meta->>'qty' AS FLOAT) AS qty, movement_meta->>'notes' AS notes, 
    CAST(movement_meta->>'shared' AS BOOLEAN) AS shared,
    movement_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(movement_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    movement_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'movement_type', movement_type, 'shipping_time', shipping_time,
      'trans_code', trans_code, 'product_code', product_code, 'tool_code', tool_code, 'place_code', place_code,
      'item_code', item_code, 'movement_code', movement_code,
      'movement_meta', json(movement_meta), 'movement_map', json(movement_map), 
      'time_stamp', time_stamp
    ) AS movement_object
  FROM movement
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW movement_map AS
  SELECT movement.id AS id, movement.code, movement.trans_code,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM movement, json_each(movement.movement_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW movement_tags AS
  SELECT movement.id AS id, code, value as tag
  FROM movement, json_each(movement.movement_meta->'tags')
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW movement_stock AS
  SELECT ROW_NUMBER() OVER (ORDER BY pl.place_name, p.product_name) as id,
    mv.place_code, pl.place_name, mv.product_code, p.product_name, 
    p.product_meta->>'unit' AS unit, movement_meta->>'notes' AS batch_no, 
    SUM(CAST(mv.movement_meta->>'qty' AS FLOAT)) AS qty, 
    MAX(date(mv.shipping_time)) AS posdate
  FROM movement mv INNER JOIN place pl ON mv.place_code = pl.code
  INNER JOIN product p ON mv.product_code = p.code
  WHERE mv.movement_type = 'MOVEMENT_INVENTORY' AND mv.deleted = false AND p.deleted = false AND pl.deleted = false
    AND mv.trans_code IN (SELECT code FROM trans WHERE deleted = false)
  GROUP BY mv.place_code, pl.place_name, mv.product_code, p.product_name, p.product_meta->>'unit', movement_meta->>'notes'
  HAVING SUM(CAST(mv.movement_meta->>'qty' AS FLOAT)) <> 0
  ORDER BY pl.place_name, p.product_name;

CREATE VIEW movement_inventory AS
  SELECT mt.id, mt.code, mt.trans_code, t.trans_type, t.direction, DATE(mt.shipping_time) AS shipping_date,
    mt.place_code, pl.place_name, mt.product_code, p.product_name,
    p.product_meta->>'unit' AS unit, mt.movement_meta->>'notes' AS batch_no, 
    CAST(mt.movement_meta->>'qty' AS FLOAT) AS qty, 
    it.customer_code, ci.customer_name, 
    coalesce(i.trans_code,  mr.trans_code, t.trans_code) AS ref_trans_code
  FROM movement mt INNER JOIN trans t ON mt.trans_code = t.code
  INNER JOIN place pl ON mt.place_code = pl.code
  INNER JOIN product p ON mt.product_code = p.code
  LEFT JOIN item i ON mt.item_code = i.code AND i.deleted = false
  LEFT JOIN trans it ON i.trans_code = it.code AND it.deleted = false
  LEFT JOIN customer ci ON it.customer_code = ci.code AND ci.deleted = false
  LEFT JOIN movement mr ON mt.movement_code = mr.code AND mr.deleted = false
  WHERE mt.movement_type = 'MOVEMENT_INVENTORY' AND mt.deleted = false AND t.deleted = false AND pl.deleted = false AND p.deleted = false
  ORDER BY mt.id;

CREATE VIEW movement_waybill AS
  SELECT mv.id, mv.code, t.code AS trans_code, t.direction, t.trans_code AS ref_trans_code, mv.shipping_time, 
    mv.tool_code, tl.tool_meta->>'serial_number' as serial_number, tl.description,
    mv.movement_meta->>'notes' as mvnotes,
    t.employee_code, t.customer_code, c.customer_name, 
    t.trans_meta->>'trans_state' as trans_state, trans_meta->>'notes' AS notes, 
    trans_meta->>'internal_notes' AS internal_notes, 
    CAST(trans_meta->>'closed' AS BOOLEAN) AS closed,t.time_stamp
  FROM trans t INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN tool tl ON mv.tool_code = tl.code
  LEFT JOIN customer c ON t.customer_code = c.code
  WHERE t.trans_type = 'TRANS_WAYBILL' and mv.deleted = false and t.deleted = false;

CREATE VIEW movement_formula AS
  SELECT mv.id, mv.code, t.code AS trans_code, CASE WHEN mv.movement_type = 'MOVEMENT_HEAD' THEN 'IN' ELSE 'OUT' END as direction, 
    mv.product_code, p.product_name, p.product_meta->>'unit' as unit,
    CAST(mv.movement_meta->>'qty' AS FLOAT) AS qty, mv.movement_meta->>'notes' as batch_no,
    mv.place_code, pl.place_name, CAST(mv.movement_meta->>'shared' AS BOOLEAN) AS shared
  FROM trans t INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN product p ON mv.product_code = p.code
  LEFT JOIN place pl ON mv.place_code = pl.code
  WHERE t.trans_type = 'TRANS_FORMULA' and mv.deleted = false and t.deleted = false;

CREATE VIEW item_shipping AS
  SELECT iv.id, iv.code, iv.trans_code, iv.direction, iv.product_code, iv.product_name, iv.unit, iv.item_qty,
    SUM(cast(COALESCE(movement_meta->>'qty','0') AS FLOAT)) as movement_qty
  FROM (
  SELECT i.id, i.code, i.trans_code, t.direction, i.product_code, p.product_name, 
    COALESCE(p.product_meta->>'unit','') as unit, 
    cast(COALESCE(item_meta->>'qty','0') AS FLOAT) item_qty
  FROM item i
  INNER JOIN trans t ON i.trans_code = t.code
  INNER JOIN product p ON i.product_code = p.code
  WHERE t.deleted = false AND i.deleted = false AND p.product_type = 'PRODUCT_ITEM' 
    AND t.trans_type IN('TRANS_ORDER', 'TRANS_WORKSHEET', 'TRANS_RENT')
  UNION 
  SELECT i.id, i.code, i.trans_code, t.direction, pc.ref_product_code as product_code, pc.component_name as product_name, 
    pc.component_unit as unit, 
    cast(COALESCE(item_meta->>'qty','0') AS FLOAT)*pc.qty item_qty
  FROM item i
  INNER JOIN trans t ON i.trans_code = t.code
  INNER JOIN product_components pc ON i.product_code = pc.product_code AND pc.component_type = 'PRODUCT_ITEM'
  WHERE t.deleted = false AND i.deleted = false 
    AND t.trans_type IN('TRANS_ORDER', 'TRANS_WORKSHEET', 'TRANS_RENT')
  ) iv
  LEFT JOIN movement mv ON mv.item_code = iv.code AND mv.product_code = iv.product_code
  GROUP BY iv.id, iv.code, iv.trans_code, iv.direction, iv.product_code, iv.product_name, iv.unit, iv.item_qty;

CREATE TABLE IF NOT EXISTS payment(
  id INTEGER,
  code TEXT NOT NULL DEFAULT 'NULL',
  paid_date TEXT NOT NULL DEFAULT CURRENT_DATE,
  trans_code TEXT NOT NULL, 
  payment_meta JSONB NOT NULL DEFAULT (json_object()),
  payment_map JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT payment_code_key UNIQUE (code),
  FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_payment_trans_code ON payment (trans_code);
CREATE INDEX idx_payment_deleted ON payment (deleted);
CREATE INDEX idx_payment_tags ON payment (json_extract(payment_meta, '$.tags'));

CREATE TRIGGER payment_default_code
AFTER INSERT ON payment
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE payment SET code = 'PMT'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;

CREATE VIEW payment_view AS
  SELECT id, code, paid_date, trans_code, 
    CAST(payment_meta->>'amount' AS FLOAT) AS amount, payment_meta->>'notes' AS notes,
    payment_meta->'tags' AS tags, 
    REPLACE(REPLACE(REPLACE(payment_meta->>'tags', '"', ''), '[', ''), ']', '') AS tag_lst, 
    payment_map, time_stamp,
    json_object(
      'id', id, 'code', code, 'paid_date', paid_date, 'trans_code', trans_code,
      'payment_meta', json(payment_meta), 'payment_map', json(payment_map), 'time_stamp', time_stamp
    ) AS payment_object
  FROM payment
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW payment_map AS
  SELECT payment.id AS id, payment.code, payment.paid_date, payment.trans_code,
    key as map_key, value as map_value, type as map_type,
    COALESCE(cf.description, key) AS description,
    COALESCE(cf.field_type, 'FIELD_STRING') AS field_type,
    CASE WHEN cf.field_type = 'FIELD_BOOL' THEN 'bool'
			WHEN cf.field_type = 'FIELD_INTEGER' THEN 'integer'
			WHEN cf.field_type = 'FIELD_NUMBER' THEN 'float'
			WHEN cf.field_type = 'FIELD_DATE' THEN 'date'
			WHEN cf.field_type = 'FIELD_DATETIME' THEN 'datetime'
			WHEN cf.field_type IN (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			ELSE 'string' END AS value_meta
  FROM payment, json_each(payment.payment_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE VIEW payment_invoice AS
  SELECT pm.id, pm.code, pm.trans_code, pt.trans_type, pt.direction, pm.paid_date, pl.place_name, pl.currency_code,
    cast(l.link_meta->>'amount' AS FLOAT) AS paid_amount, cast(l.link_meta->>'rate' AS FLOAT) AS paid_rate, 
    it.code AS ref_trans_code, it.currency_code AS invoice_curr,
    im.amount AS invoice_amount, pm.payment_meta->>'notes' AS description
  FROM link l 
  INNER JOIN payment pm ON l.link_code_1 = pm.code
  INNER JOIN trans pt ON pm.trans_code = pt.code INNER JOIN place pl ON pt.place_code = pl.code
  INNER JOIN trans it ON l.link_code_2 = it.code
  INNER JOIN(
    SELECT trans_code, sum(cast(item_meta->>'amount' AS FLOAT)) AS amount FROM item GROUP BY trans_code) im ON it.code = im.trans_code
  WHERE l.link_type_1 = 'LINK_PAYMENT' AND l.link_type_2 = 'LINK_TRANS' AND it.trans_type IN('TRANS_INVOICE','TRANS_RECEIPT');

CREATE VIEW payment_tags AS
  SELECT payment.id AS id, code, value as tag
  FROM payment, json_each(payment.payment_meta->'tags')
  WHERE deleted = 0 AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE TABLE IF NOT EXISTS log(
  id INTEGER,
  code TEXT UNIQUE NOT NULL DEFAULT 'NULL',
  log_type TEXT NOT NULL DEFAULT 'LOG_UPDATE',
  ref_type TEXT NOT NULL DEFAULT 'TRANS',
  ref_code TEXT NOT NULL,
  auth_code TEXT NOT NULL,
  data JSONB NOT NULL DEFAULT (json_object()),
  time_stamp TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY("id" AUTOINCREMENT),
  CONSTRAINT log_code_key UNIQUE (code),
  CHECK( log_type IN ('LOG_INSERT', 'LOG_UPDATE', 'LOG_DELETE') )
);

CREATE INDEX idx_log_ref_type ON log (ref_type);
CREATE INDEX idx_log_ref_code ON log (ref_code);
CREATE INDEX idx_log_auth_code ON log (auth_code);
CREATE INDEX idx_log_time_stamp ON log (time_stamp);
CREATE INDEX idx_log_deleted ON log (deleted);

CREATE TRIGGER log_default_code
AFTER INSERT ON log
FOR EACH ROW
WHEN NEW.code = 'NULL'
BEGIN
  UPDATE log SET code = 'LOG'||unixepoch()||'N'||NEW.id WHERE id = NEW.id;
END;