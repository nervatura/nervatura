CREATE TABLE IF NOT EXISTS usref(
  id INTEGER AUTO_INCREMENT NOT NULL,
  refnumber VARCHAR(255) NOT NULL,
  value VARCHAR(255) NOT NULL,
  changed TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT usref_pkey PRIMARY KEY (id),
  CONSTRAINT usref_refnumber_key UNIQUE (refnumber)
);

CREATE TRIGGER usref_changed_timestamp 
  BEFORE UPDATE ON usref
  FOR EACH ROW
BEGIN
  SET NEW.changed = CURRENT_TIMESTAMP;
END;

CREATE TABLE IF NOT EXISTS config(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  config_type ENUM ('CONFIG_MAP', 'CONFIG_SHORTCUT', 'CONFIG_MESSAGE', 'CONFIG_PATTERN', 'CONFIG_REPORT', 'CONFIG_PRINT_QUEUE', 'CONFIG_DATA') 
    NOT NULL DEFAULT 'CONFIG_MAP', 
  data JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT config_pkey PRIMARY KEY (id),
  CONSTRAINT config_code_key UNIQUE (code)
);

CREATE INDEX idx_config_type ON config (config_type);
CREATE INDEX idx_config_deleted ON config (deleted);

CREATE TRIGGER config_default_code 
  BEFORE INSERT ON config
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM config);
    SET NEW.code = CONCAT('CNF', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW config_map AS
  SELECT id, code,
    data->>"$.field_name" AS field_name,
    data->>"$.field_type" AS field_type,
    data->>"$.description" AS description,
    data->"$.tags" AS tags,
    data->"$.filter" AS filter
  FROM config
  WHERE config_type = 'CONFIG_MAP' AND deleted = false;

CREATE VIEW config_shortcut AS
  SELECT id, code,
    data->>"$.shortcut_key" AS shortcut_key,
    data->>"$.description" AS description,
    data->>"$.modul" AS modul,
    data->>"$.method" AS method,
    data->>"$.func_name" AS func_name,
    data->>"$.address" AS address,
    data->"$.fields" AS fields
  FROM config
  WHERE config_type = 'CONFIG_SHORTCUT' AND deleted = false;

CREATE VIEW config_message AS
  SELECT id, code,
    data->>"$.key" AS message_key,
    data->>"$.lang" AS lang,
    data->>"$.value" AS message_value
  FROM config
  WHERE config_type = 'CONFIG_MESSAGE' AND deleted = false;

CREATE VIEW config_pattern AS
  SELECT id, code,
    data->>"$.trans_type" AS trans_type,
    data->>"$.description" AS description,
    data->>"$.notes" AS notes,
    data->"$.default_pattern" AS default_pattern
  FROM config
  WHERE config_type = 'CONFIG_PATTERN' AND deleted = false;

CREATE VIEW config_print_queue AS
  SELECT id, code,
    data->>"$.ref_type" AS ref_type,
    data->>"$.ref_code" AS ref_code,
    CAST(data->>"$.qty" AS FLOAT) AS qty,
    data->>"$.report_code" AS report_code,
    data->>"$.orientation" AS orientation,
    data->>"$.paper_size" AS paper_size,
    data->>"$.auth_code" AS auth_code,
    data->>"$.time_stamp" AS time_stamp
  FROM config
  WHERE config_type = 'CONFIG_PRINT_QUEUE' AND deleted = false;

CREATE VIEW config_report AS
  SELECT id, code,
    data->>"$.report_key" AS report_key,
    data->>"$.report_type" AS report_type,
    data->>"$.trans_type" AS trans_type,
    data->>"$.direction" AS direction,
    data->>"$.report_name" AS report_name,
    data->>"$.description" AS description,
    data->>"$.label" AS label,
    data->>"$.file_type" AS file_type,
    data->"$.template" AS template
  FROM config
  WHERE config_type = 'CONFIG_REPORT' AND deleted = false;

CREATE VIEW config_data AS
  SELECT ROW_NUMBER() OVER (ORDER BY c.code, jt.config_key) as id, c.code AS config_code, jt.config_key,
    JSON_UNQUOTE(JSON_EXTRACT(c.data, CONCAT('$.', jt.config_key))) AS config_value,
    JSON_TYPE(JSON_EXTRACT(c.data, CONCAT('$.', jt.config_key))) AS config_type
  FROM config c
  CROSS JOIN JSON_TABLE(
      JSON_KEYS(c.data),
      '$[*]' COLUMNS (config_key VARCHAR(255) PATH '$')
  ) AS jt
  WHERE c.config_type = 'CONFIG_DATA' 
  AND JSON_TYPE(c.data) = 'OBJECT' AND c.deleted = false;

CREATE TABLE IF NOT EXISTS auth(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL, 
  user_name VARCHAR(255) NOT NULL,
  user_group ENUM ('GROUP_ADMIN', 'GROUP_USER', 'GROUP_GUEST') NOT NULL DEFAULT 'GROUP_USER',
  disabled BOOLEAN NOT NULL DEFAULT false,
  auth_meta JSON NOT NULL,
  auth_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT auth_pkey PRIMARY KEY (id),
  CONSTRAINT auth_code_key UNIQUE (code)
);

CREATE UNIQUE INDEX idx_auth_user_name ON auth (user_name);
CREATE INDEX idx_auth_deleted ON auth (deleted);
CREATE INDEX idx_auth_disabled ON auth (disabled);

CREATE TRIGGER auth_default_code 
  BEFORE INSERT ON auth
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM auth);
    SET NEW.code = CONCAT('USR', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW auth_view AS
  SELECT id, code, user_name, user_group, disabled,
    auth_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(auth_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    auth_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'user_name', user_name, 'user_group', user_group, 'disabled', disabled,
      'auth_meta', auth_meta, 'auth_map', auth_map, 'time_stamp', time_stamp
    ) AS auth_object
  FROM auth
  WHERE deleted = false;

CREATE VIEW auth_map AS
  SELECT tbl.id AS id, tbl.code, user_name, user_group, jt.map_key,
    JSON_UNQUOTE(JSON_EXTRACT(tbl.auth_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.auth_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM auth tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.auth_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS currency(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  currency_meta JSON NOT NULL,
  currency_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT currency_pkey PRIMARY KEY (id),
  CONSTRAINT currency_code_key UNIQUE (code)
);

CREATE INDEX idx_currency_deleted ON currency (deleted);

CREATE VIEW currency_view AS
  SELECT id, code,
    currency_meta->>"$.description" AS description,
    CAST(currency_meta->>"$.digit" AS UNSIGNED) AS digit, CAST(currency_meta->>"$.cash_round" AS UNSIGNED) AS cash_round, 
    currency_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(currency_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst, 
    currency_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'currency_meta', currency_meta, 'currency_map', currency_map, 'time_stamp', time_stamp
    ) AS currency_object
  FROM currency
  WHERE deleted = false;

CREATE VIEW currency_map AS
  SELECT tbl.id AS id, tbl.code, currency_meta->>"$.description" AS currency_description,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.currency_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.currency_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM currency tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.currency_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW currency_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM currency tbl
  CROSS JOIN JSON_TABLE(
    tbl.currency_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS customer(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  customer_type ENUM ('CUSTOMER_OWN', 'CUSTOMER_COMPANY', 'CUSTOMER_PRIVATE', 'CUSTOMER_OTHER') NOT NULL DEFAULT 'CUSTOMER_COMPANY',
  customer_name VARCHAR(255) NOT NULL,
  addresses JSON NOT NULL,
  contacts JSON NOT NULL,
  events JSON NOT NULL, 
  customer_meta JSON NOT NULL,
  customer_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT customer_pkey PRIMARY KEY (id),
  CONSTRAINT customer_code_key UNIQUE (code)
);

CREATE INDEX idx_customer_customer_type ON customer (customer_type);
CREATE INDEX idx_customer_deleted ON customer (deleted);
CREATE INDEX idx_customer_customer_name ON customer (customer_name);

CREATE TRIGGER customer_default_code 
  BEFORE INSERT ON customer
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM customer);
    SET NEW.code = CONCAT('CUS', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW customer_view AS
  SELECT id, code, customer_type, customer_name, customer_meta->>"$.tax_number" AS tax_number, customer_meta->>"$.account" AS account,
    (customer_meta->>"$.tax_free" = 'true') AS tax_free, CAST(customer_meta->>"$.terms" AS UNSIGNED) AS terms,
    CAST(customer_meta->>"$.credit_limit" AS FLOAT) AS credit_limit, CAST(customer_meta->>"$.discount" AS FLOAT) AS discount,
    customer_meta->>"$.notes" AS notes, (customer_meta->>"$.inactive" = 'true') AS inactive,
    customer_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(customer_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    addresses, contacts, events, customer_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'customer_type', customer_type, 'customer_name', customer_name,
      'addresses', addresses, 'contacts', contacts, 'events', events,
      'customer_meta', customer_meta, 'customer_map', customer_map, 'time_stamp', time_stamp
    ) AS customer_object
  FROM customer
  WHERE deleted = false;

CREATE VIEW customer_contacts AS
  SELECT c.id AS id, c.code, 
    customer_name,
    jt.value->>"$.first_name" AS first_name,
    jt.value->>"$.surname" AS surname,
    jt.value->>"$.status" AS status,
    jt.value->>"$.phone" AS phone,
    jt.value->>"$.mobile" AS mobile,
    jt.value->>"$.email" AS email,
    jt.value->>"$.notes" AS notes,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.contact_map" AS contact_map
  FROM customer c
  CROSS JOIN JSON_TABLE(c.contacts,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW customer_addresses AS
  SELECT c.id AS id, c.code, 
    customer_name,
    jt.value->>"$.country" AS country,
    jt.value->>"$.state" AS state,
    jt.value->>"$.zip_code" AS zip_code,
    jt.value->>"$.city" AS city,
    jt.value->>"$.street" AS street,
    jt.value->>"$.notes" AS notes,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.address_map" AS address_map
  FROM customer c
  CROSS JOIN JSON_TABLE(c.addresses,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW customer_events AS
  SELECT c.id AS id, c.code, 
    customer_name,
    jt.value->>"$.uid" AS uid,
    jt.value->>"$.subject" AS subject,
    CASE WHEN jt.value->>"$.start_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.start_time", '%Y-%m-%dT%H:%i:%s') END AS start_time,
    CASE WHEN jt.value->>"$.end_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.end_time", '%Y-%m-%dT%H:%i:%s') END AS end_time,
    jt.value->>"$.place" AS place,
    jt.value->>"$.description" AS description,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.event_map" AS event_map
  FROM customer c
  CROSS JOIN JSON_TABLE(c.events,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW customer_map AS
  SELECT tbl.id AS id, tbl.code, tbl.customer_name,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.customer_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.customer_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM customer tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.customer_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW customer_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM customer tbl
  CROSS JOIN JSON_TABLE(
    tbl.customer_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS employee(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL, 
  address JSON NOT NULL,
  contact JSON NOT NULL,
  events JSON NOT NULL, 
  employee_meta JSON NOT NULL,
  employee_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT employee_pkey PRIMARY KEY (id),
  CONSTRAINT employee_code_key UNIQUE (code)
);

CREATE INDEX idx_employee_deleted ON employee (deleted);

CREATE TRIGGER employee_default_code 
  BEFORE INSERT ON employee
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM employee);
    SET NEW.code = CONCAT('EMP', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW employee_view AS
  SELECT id, code, 
    CASE WHEN employee_meta->>"$.start_date" = '' THEN NULL ELSE STR_TO_DATE(employee_meta->>"$.start_date", '%Y-%m-%d') END AS start_date,
    CASE WHEN employee_meta->>"$.end_date" = '' THEN NULL ELSE STR_TO_DATE(employee_meta->>"$.end_date", '%Y-%m-%d') END AS end_date,
    (employee_meta->>"$.inactive" = 'true') AS inactive, employee_meta->>"$.notes" AS notes,
    contact->>"$.first_name" AS first_name, contact->>"$.surname" AS surname, contact->>"$.status" AS status,
    contact->>"$.phone" AS phone, contact->>"$.mobile" AS mobile, contact->>"$.email" AS email, 
    address->>"$.country" AS country, address->>"$.state" AS state, address->>"$.zip_code" AS zip_code,
    address->>"$.city" AS city, address->>"$.street" AS street, 
    employee_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(employee_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    employee_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'address', address, 'contact', contact, 'events', events,
      'employee_meta', employee_meta, 'employee_map', employee_map, 'time_stamp', time_stamp
    ) AS employee_object
  FROM employee
  WHERE deleted = false;

CREATE VIEW employee_events AS
  SELECT c.id AS id, c.code, 
    contact->>"$.first_name" AS first_name, contact->>"$.surname" AS surname,
    jt.value->>"$.uid" AS uid,
    jt.value->>"$.subject" AS subject,
    CASE WHEN jt.value->>"$.start_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.start_time", '%Y-%m-%dT%H:%i:%s') END AS start_time,
    CASE WHEN jt.value->>"$.end_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.end_time", '%Y-%m-%dT%H:%i:%s') END AS end_time,
    jt.value->>"$.place" AS place,
    jt.value->>"$.description" AS description,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.event_map" AS event_map
  FROM employee c
  CROSS JOIN JSON_TABLE(c.events,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW employee_map AS
  SELECT tbl.id AS id, tbl.code, 
    contact->>"$.first_name" AS first_name, contact->>"$.surname" AS surname,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.employee_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.employee_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM employee tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.employee_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW employee_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM employee tbl
  CROSS JOIN JSON_TABLE(
    tbl.employee_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS place(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  place_type ENUM ('PLACE_BANK', 'PLACE_CASH', 'PLACE_WAREHOUSE', 'PLACE_OTHER') NOT NULL DEFAULT 'PLACE_WAREHOUSE',
  place_name VARCHAR(255) NOT NULL,
  currency_code VARCHAR(255),
  address JSON NOT NULL,
  contacts JSON NOT NULL,
  place_meta JSON NOT NULL,
  place_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT place_pkey PRIMARY KEY (id),
  CONSTRAINT place_code_key UNIQUE (code),
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_place_deleted ON place (deleted);
CREATE INDEX idx_place_place_name ON place (place_name);

CREATE TRIGGER place_default_code 
  BEFORE INSERT ON place
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM place);
    SET NEW.code = CONCAT('PLA', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW place_view AS
  SELECT id, code, place_type, place_name, currency_code,
    (place_meta->>"$.inactive" = 'true') AS inactive, place_meta->>"$.notes" AS notes,
    address->>"$.country" AS country, address->>"$.state" AS state, address->>"$.zip_code" AS zip_code,
    address->>"$.city" AS city, address->>"$.street" AS street, 
    place_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(place_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    place_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'place_type', place_type, 'place_name', place_name,
      'currency_code', currency_code, 'address', address, 'contacts', contacts,
      'place_meta', place_meta, 'place_map', place_map, 'time_stamp', time_stamp
    ) AS place_object
  FROM place
  WHERE deleted = false;

CREATE VIEW place_contacts AS
  SELECT c.id AS id, c.code, 
    place_name,
    jt.value->>"$.first_name" AS first_name,
    jt.value->>"$.surname" AS surname,
    jt.value->>"$.status" AS status,
    jt.value->>"$.phone" AS phone,
    jt.value->>"$.mobile" AS mobile,
    jt.value->>"$.email" AS email,
    jt.value->>"$.notes" AS notes,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.contact_map" AS contact_map
  FROM place c
  CROSS JOIN JSON_TABLE(c.contacts,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW place_map AS
  SELECT tbl.id AS id, tbl.code, place_name,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.place_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.place_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM place tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.place_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW place_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM place tbl
  CROSS JOIN JSON_TABLE(
    tbl.place_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS tax(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL, 
  tax_meta JSON NOT NULL,
  tax_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT tax_pkey PRIMARY KEY (id),
  CONSTRAINT tax_code_key UNIQUE (code)
);

CREATE INDEX idx_tax_deleted ON tax (deleted);

CREATE VIEW tax_view AS
  SELECT id, code,
    tax_meta->>"$.description" AS description, CAST(tax_meta->>"$.rate_value" AS FLOAT) AS rate_value,
    (tax_meta->>"$.inactive" = 'true') AS inactive,
    tax_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(tax_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    tax_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'tax_meta', tax_meta, 'tax_map', tax_map, 'time_stamp', time_stamp
    ) AS tax_object
  FROM tax
  WHERE deleted = false;

CREATE VIEW tax_map AS
  SELECT tbl.id AS id, tbl.code, 
    tax_meta->>"$.description" AS tax_description, tax_meta->>"$.rate_value" AS rate_value,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.tax_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.tax_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM tax tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.tax_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW tax_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM tax tbl
  CROSS JOIN JSON_TABLE(
    tbl.tax_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS link(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  link_type_1 ENUM ('LINK_CUSTOMER', 'LINK_EMPLOYEE', 'LINK_ITEM', 'LINK_MOVEMENT', 'LINK_PAYMENT', 'LINK_PLACE', 'LINK_PRODUCT', 'LINK_PROJECT', 'LINK_TOOL', 'LINK_TRANS') 
    NOT NULL DEFAULT 'LINK_TRANS',
  link_code_1 VARCHAR(255) NOT NULL,
  link_type_2 ENUM ('LINK_CUSTOMER', 'LINK_EMPLOYEE', 'LINK_ITEM', 'LINK_MOVEMENT', 'LINK_PAYMENT', 'LINK_PLACE', 'LINK_PRODUCT', 'LINK_PROJECT', 'LINK_TOOL', 'LINK_TRANS') 
    NOT NULL DEFAULT 'LINK_TRANS',
  link_code_2 VARCHAR(255) NOT NULL, 
  link_meta JSON NOT NULL,
  link_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT link_pkey PRIMARY KEY (id),
  CONSTRAINT link_code_key UNIQUE (code)
);

CREATE INDEX idx_link_link_code_1 ON link (link_type_1, link_code_1);
CREATE INDEX idx_link_link_code_2 ON link (link_type_2, link_code_2);
CREATE INDEX idx_link_deleted ON link (deleted);

CREATE TRIGGER link_default_code 
  BEFORE INSERT ON link
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM link);
    SET NEW.code = CONCAT('LNK', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE TRIGGER link_insert 
BEFORE INSERT ON link
FOR EACH ROW
BEGIN
  -- Check customer codes
  IF NEW.link_type_1 = "CUSTOMER" AND NOT EXISTS (SELECT 1 FROM customer WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid customer code";
  END IF;
  IF NEW.link_type_2 = "CUSTOMER" AND NOT EXISTS (SELECT 1 FROM customer WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid customer code";
  END IF;
  
  -- Check employee codes
  IF NEW.link_type_1 = "EMPLOYEE" AND NOT EXISTS (SELECT 1 FROM employee WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid employee code";
  END IF;
  IF NEW.link_type_2 = "EMPLOYEE" AND NOT EXISTS (SELECT 1 FROM employee WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid employee code";
  END IF;
  
  -- Check item codes
  IF NEW.link_type_1 = "ITEM" AND NOT EXISTS (SELECT 1 FROM item WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid item code";
  END IF;
  IF NEW.link_type_2 = "ITEM" AND NOT EXISTS (SELECT 1 FROM item WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid item code";
  END IF;
  
  -- Check movement codes
  IF NEW.link_type_1 = "MOVEMENT" AND NOT EXISTS (SELECT 1 FROM movement WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid movement code";
  END IF;
  IF NEW.link_type_2 = "MOVEMENT" AND NOT EXISTS (SELECT 1 FROM movement WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid movement code";
  END IF;
  
  -- Check payment codes
  IF NEW.link_type_1 = "PAYMENT" AND NOT EXISTS (SELECT 1 FROM payment WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid payment code";
  END IF;
  IF NEW.link_type_2 = "PAYMENT" AND NOT EXISTS (SELECT 1 FROM payment WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid payment code";
  END IF;
  
  -- Check place codes
  IF NEW.link_type_1 = "PLACE" AND NOT EXISTS (SELECT 1 FROM place WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid place code";
  END IF;
  IF NEW.link_type_2 = "PLACE" AND NOT EXISTS (SELECT 1 FROM place WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid place code";
  END IF;
  
  -- Check product codes
  IF NEW.link_type_1 = "PRODUCT" AND NOT EXISTS (SELECT 1 FROM product WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid product code";
  END IF;
  IF NEW.link_type_2 = "PRODUCT" AND NOT EXISTS (SELECT 1 FROM product WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid product code";
  END IF;
  
  -- Check project codes
  IF NEW.link_type_1 = "PROJECT" AND NOT EXISTS (SELECT 1 FROM project WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid project code";
  END IF;
  IF NEW.link_type_2 = "PROJECT" AND NOT EXISTS (SELECT 1 FROM project WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid project code";
  END IF;
  
  -- Check tool codes
  IF NEW.link_type_1 = "TOOL" AND NOT EXISTS (SELECT 1 FROM tool WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid tool code";
  END IF;
  IF NEW.link_type_2 = "TOOL" AND NOT EXISTS (SELECT 1 FROM tool WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid tool code";
  END IF;
  
  -- Check trans codes
  IF NEW.link_type_1 = "TRANS" AND NOT EXISTS (SELECT 1 FROM trans WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid trans code";
  END IF;
  IF NEW.link_type_2 = "TRANS" AND NOT EXISTS (SELECT 1 FROM trans WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid trans code";
  END IF;
END;

CREATE TRIGGER link_update 
BEFORE UPDATE ON link
FOR EACH ROW
BEGIN
  -- Check customer codes
  IF NEW.link_type_1 = "CUSTOMER" AND NOT EXISTS (SELECT 1 FROM customer WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid customer code";
  END IF;
  IF NEW.link_type_2 = "CUSTOMER" AND NOT EXISTS (SELECT 1 FROM customer WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid customer code";
  END IF;
  
  -- Check employee codes
  IF NEW.link_type_1 = "EMPLOYEE" AND NOT EXISTS (SELECT 1 FROM employee WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid employee code";
  END IF;
  IF NEW.link_type_2 = "EMPLOYEE" AND NOT EXISTS (SELECT 1 FROM employee WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid employee code";
  END IF;
  
  -- Check item codes
  IF NEW.link_type_1 = "ITEM" AND NOT EXISTS (SELECT 1 FROM item WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid item code";
  END IF;
  IF NEW.link_type_2 = "ITEM" AND NOT EXISTS (SELECT 1 FROM item WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid item code";
  END IF;
  
  -- Check movement codes
  IF NEW.link_type_1 = "MOVEMENT" AND NOT EXISTS (SELECT 1 FROM movement WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid movement code";
  END IF;
  IF NEW.link_type_2 = "MOVEMENT" AND NOT EXISTS (SELECT 1 FROM movement WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid movement code";
  END IF;
  
  -- Check payment codes
  IF NEW.link_type_1 = "PAYMENT" AND NOT EXISTS (SELECT 1 FROM payment WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid payment code";
  END IF;
  IF NEW.link_type_2 = "PAYMENT" AND NOT EXISTS (SELECT 1 FROM payment WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid payment code";
  END IF;
  
  -- Check place codes
  IF NEW.link_type_1 = "PLACE" AND NOT EXISTS (SELECT 1 FROM place WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid place code";
  END IF;
  IF NEW.link_type_2 = "PLACE" AND NOT EXISTS (SELECT 1 FROM place WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid place code";
  END IF;
  
  -- Check product codes
  IF NEW.link_type_1 = "PRODUCT" AND NOT EXISTS (SELECT 1 FROM product WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid product code";
  END IF;
  IF NEW.link_type_2 = "PRODUCT" AND NOT EXISTS (SELECT 1 FROM product WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid product code";
  END IF;
  
  -- Check project codes
  IF NEW.link_type_1 = "PROJECT" AND NOT EXISTS (SELECT 1 FROM project WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid project code";
  END IF;
  IF NEW.link_type_2 = "PROJECT" AND NOT EXISTS (SELECT 1 FROM project WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid project code";
  END IF;
  
  -- Check tool codes
  IF NEW.link_type_1 = "TOOL" AND NOT EXISTS (SELECT 1 FROM tool WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid tool code";
  END IF;
  IF NEW.link_type_2 = "TOOL" AND NOT EXISTS (SELECT 1 FROM tool WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid tool code";
  END IF;
  
  -- Check trans codes
  IF NEW.link_type_1 = "TRANS" AND NOT EXISTS (SELECT 1 FROM trans WHERE code = NEW.link_code_1) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid trans code";
  END IF;
  IF NEW.link_type_2 = "TRANS" AND NOT EXISTS (SELECT 1 FROM trans WHERE code = NEW.link_code_2) THEN
    SIGNAL SQLSTATE "45000" SET MESSAGE_TEXT = "Invalid trans code";
  END IF;
END;

CREATE VIEW link_view AS
  SELECT id, code, 
    link_type_1, link_code_1, link_type_2, link_code_2,
    CAST(link_meta->>"$.qty" AS FLOAT) AS qty, CAST(link_meta->>"$.amount" AS FLOAT) AS amount, CAST(link_meta->>"$.rate" AS FLOAT) AS rate,
    COALESCE(link_meta->>"$.notes", '') AS notes,
    link_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(link_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    link_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'link_type_1', link_type_1, 'link_code_1', link_code_1,
      'link_type_2', link_type_2, 'link_code_2', link_code_2, 'link_meta', link_meta, 'link_map', link_map, 'time_stamp', time_stamp
    ) AS link_object
  FROM link
  WHERE deleted = false;

CREATE VIEW link_map AS
  SELECT tbl.id AS id, tbl.code, link_type_1, link_code_1, link_type_2, link_code_2,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.link_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.link_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM link tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.link_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW link_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM link tbl
  CROSS JOIN JSON_TABLE(
    tbl.link_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS product(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  product_type ENUM ('PRODUCT_ITEM', 'PRODUCT_SERVICE', 'PRODUCT_VIRTUAL') NOT NULL DEFAULT 'PRODUCT_ITEM',
  product_name VARCHAR(255) NOT NULL,
  tax_code VARCHAR(255) NOT NULL,
  events JSON NOT NULL,
  product_meta JSON NOT NULL,
  product_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT product_pkey PRIMARY KEY (id),
  CONSTRAINT product_code_key UNIQUE (code),
  FOREIGN KEY (tax_code) REFERENCES tax(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_product_product_type ON product (product_type);
CREATE INDEX idx_product_product_name ON product (product_name);
CREATE INDEX idx_product_tax_code ON product (tax_code);
CREATE INDEX idx_product_deleted ON product (deleted);

CREATE TRIGGER product_default_code 
  BEFORE INSERT ON product
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM product);
    SET NEW.code = CONCAT('PRD', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW product_view AS
  SELECT id, code, product_type, product_name, tax_code,
    product_meta->>"$.unit" AS unit,
    product_meta->>"$.barcode_type" AS barcode_type, product_meta->>"$.barcode" AS barcode, 
    CAST(product_meta->>'$.barcode_qty' AS FLOAT) AS barcode_qty,
    (product_meta->>"$.inactive" = 'true') AS inactive, product_meta->>"$.notes" AS notes, 
    product_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(product_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    events, product_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'product_type', product_type, 'product_name', product_name,
      'tax_code', tax_code, 'events', events, 'product_meta', product_meta, 'product_map', product_map, 'time_stamp', time_stamp
    ) AS product_object
  FROM product
  WHERE deleted = false;

CREATE VIEW product_events AS
  SELECT c.id AS id, c.code, 
    product_name,
    jt.value->>"$.uid" AS uid,
    jt.value->>"$.subject" AS subject,
    CASE WHEN jt.value->>"$.start_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.start_time", '%Y-%m-%dT%H:%i:%s') END AS start_time,
    CASE WHEN jt.value->>"$.end_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.end_time", '%Y-%m-%dT%H:%i:%s') END AS end_time,
    jt.value->>"$.place" AS place,
    jt.value->>"$.description" AS description,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.event_map" AS event_map
  FROM product c
  CROSS JOIN JSON_TABLE(c.events,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW product_map AS
  SELECT tbl.id AS id, tbl.code, product_name,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.product_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.product_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM product tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.product_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW product_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM product tbl
  CROSS JOIN JSON_TABLE(
    tbl.product_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE OR REPLACE VIEW product_components AS
  SELECT p.id, p.code AS product_code, p.product_name, COALESCE(p.product_meta->>"$.unit",'') as unit, 
  c.code as ref_product_code, c.product_name as component_name, COALESCE(c.product_meta->>"$.unit",'') as component_unit, 
  c.product_type as component_type,
  CAST(l.link_meta->>"$.qty" AS FLOAT) AS qty, COALESCE(l.link_meta->>"$.notes", '') AS notes
  FROM product p INNER JOIN link l ON l.link_code_1 = p.code
  INNER JOIN product c ON l.link_code_2 = c.code
  WHERE p.product_type = 'PRODUCT_VIRTUAL' AND link_type_1 = 'LINK_PRODUCT' AND link_type_2 = 'LINK_PRODUCT' 
  AND p.deleted = false AND l.deleted = false AND c.deleted = false;

CREATE TABLE IF NOT EXISTS project(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  project_name VARCHAR(255) NOT NULL,
  customer_code VARCHAR(255),
  addresses JSON NOT NULL,
  contacts JSON NOT NULL,
  events JSON NOT NULL,
  project_meta JSON NOT NULL,
  project_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT project_pkey PRIMARY KEY (id),
  CONSTRAINT project_code_key UNIQUE (code),
  FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_project_customer_code ON project (customer_code);
CREATE INDEX idx_project_project_name ON project (project_name);
CREATE INDEX idx_project_deleted ON project (deleted);

CREATE TRIGGER project_default_code 
  BEFORE INSERT ON project
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM project);
    SET NEW.code = CONCAT('PRJ', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW project_view AS
  SELECT id, code, project_name, customer_code,
    CASE WHEN project_meta->>"$.start_date" = '' THEN NULL ELSE STR_TO_DATE(project_meta->>"$.start_date", '%Y-%m-%d') END AS start_date,
    CASE WHEN project_meta->>"$.end_date" = '' THEN NULL ELSE STR_TO_DATE(project_meta->>"$.end_date", '%Y-%m-%d') END AS end_date,
    (project_meta->>"$.inactive" = 'true') AS inactive, project_meta->>"$.notes" AS notes, 
    project_meta->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(project_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    addresses, contacts, events, project_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'project_name', project_name, 'customer_code', customer_code,
      'addresses', addresses, 'contacts', contacts, 'events', events, 'project_map', project_map, 'time_stamp', time_stamp
    ) AS project_object
  FROM project
  WHERE deleted = false;

CREATE VIEW project_contacts AS
  SELECT c.id AS id, c.code, 
    project_name,
    jt.value->>"$.first_name" AS first_name,
    jt.value->>"$.surname" AS surname,
    jt.value->>"$.status" AS status,
    jt.value->>"$.phone" AS phone,
    jt.value->>"$.mobile" AS mobile,
    jt.value->>"$.email" AS email,
    jt.value->>"$.notes" AS notes,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.contact_map" AS contact_map
  FROM project c
  CROSS JOIN JSON_TABLE(c.contacts,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW project_addresses AS
  SELECT c.id AS id, c.code, 
    project_name,
    jt.value->>"$.country" AS country,
    jt.value->>"$.state" AS state,
    jt.value->>"$.zip_code" AS zip_code,
    jt.value->>"$.city" AS city,
    jt.value->>"$.street" AS street,
    jt.value->>"$.notes" AS notes,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.address_map" AS address_map
  FROM project c
  CROSS JOIN JSON_TABLE(c.addresses,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW project_events AS
  SELECT c.id AS id, c.code, 
    project_name,
    jt.value->>"$.uid" AS uid,
    jt.value->>"$.subject" AS subject,
    CASE WHEN jt.value->>"$.start_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.start_time", '%Y-%m-%dT%H:%i:%s') END AS start_time,
    CASE WHEN jt.value->>"$.end_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.end_time", '%Y-%m-%dT%H:%i:%s') END AS end_time,
    jt.value->>"$.place" AS place,
    jt.value->>"$.description" AS description,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.event_map" AS event_map
  FROM project c
  CROSS JOIN JSON_TABLE(c.events,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW project_map AS
  SELECT tbl.id AS id, tbl.code, project_name,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.project_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.project_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM project tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.project_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW project_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM project tbl
  CROSS JOIN JSON_TABLE(
    tbl.project_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS rate(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  rate_type ENUM ('RATE_RATE', 'RATE_BUY', 'RATE_SELL', 'RATE_AVERAGE') NOT NULL DEFAULT 'RATE_RATE',
  rate_date DATE NOT NULL DEFAULT (CURRENT_DATE),
  place_code VARCHAR(255),
  currency_code VARCHAR(255) NOT NULL, 
  rate_meta JSON NOT NULL,
  rate_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT rate_pkey PRIMARY KEY (id),
  CONSTRAINT rate_code_key UNIQUE (code),
  FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE UNIQUE INDEX idx_rate_tdcp ON rate (rate_type, rate_date, currency_code, place_code);
CREATE INDEX idx_rate_rate_type ON rate (rate_type);
CREATE INDEX idx_rate_rate_date ON rate (rate_date);
CREATE INDEX idx_rate_place_code ON rate (place_code);
CREATE INDEX idx_rate_currency_code ON rate (currency_code);
CREATE INDEX idx_rate_deleted ON rate (deleted);

CREATE TRIGGER rate_default_code 
  BEFORE INSERT ON rate
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM rate);
    SET NEW.code = CONCAT('RAT', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW rate_view AS
  SELECT id, code, rate_type, rate_date, place_code, currency_code,
    CAST(rate_meta->>"$.rate_value" AS FLOAT) AS rate_value,
    rate_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(rate_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    rate_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'rate_type', rate_type, 'rate_date', rate_date, 'place_code', place_code,
      'currency_code', currency_code, 'rate_meta', rate_meta, 'rate_map', rate_map, 'time_stamp', time_stamp
    ) AS rate_object
  FROM rate
  WHERE deleted = false;

CREATE VIEW rate_map AS
  SELECT tbl.id AS id, tbl.code,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.rate_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.rate_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM rate tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.rate_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW rate_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM rate tbl
  CROSS JOIN JSON_TABLE(
    tbl.rate_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS tool(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  product_code VARCHAR(255) NOT NULL,
  events JSON NOT NULL,
  tool_meta JSON NOT NULL,
  tool_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT tool_pkey PRIMARY KEY (id),
  CONSTRAINT tool_code_key UNIQUE (code),
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_tool_description ON tool (description);
CREATE INDEX idx_tool_product_code ON tool (product_code);
CREATE INDEX idx_tool_deleted ON tool (deleted);

CREATE TRIGGER tool_default_code 
  BEFORE INSERT ON tool
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM tool);
    SET NEW.code = CONCAT('SER', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW tool_view AS
  SELECT id, code, product_code, description,
    tool_meta->>"$.serial_number" AS serial_number,
    (tool_meta->>"$.inactive" = 'true') AS inactive, tool_meta->>"$.notes" AS notes, 
    tool_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(tool_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    tool_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'description', description, 'product_code', product_code,
      'events', events, 'tool_meta', tool_meta, 'tool_map', tool_map, 'time_stamp', time_stamp
    ) AS tool_object
  FROM tool
  WHERE deleted = false;

CREATE VIEW tool_events AS
  SELECT c.id AS id, c.code, description as tool_description,
    jt.value->>"$.uid" AS uid,
    jt.value->>"$.subject" AS subject,
    CASE WHEN jt.value->>"$.start_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.start_time", '%Y-%m-%dT%H:%i:%s') END AS start_time,
    CASE WHEN jt.value->>"$.end_time" = '' THEN NULL ELSE STR_TO_DATE(jt.value->>"$.end_time", '%Y-%m-%dT%H:%i:%s') END AS end_time,
    jt.value->>"$.place" AS place,
    jt.value->>"$.description" AS description,
    jt.value->"$.tags" AS tags,
    REPLACE(REPLACE(REPLACE(jt.value->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    jt.value->"$.event_map" AS event_map
  FROM tool c
  CROSS JOIN JSON_TABLE(c.events,
    '$[*]' COLUMNS (value JSON PATH '$')
  ) AS jt
  WHERE c.deleted = false;

CREATE VIEW tool_map AS
  SELECT tbl.id AS id, tbl.code, tbl.description as tool_description,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.tool_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.tool_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM tool tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.tool_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW tool_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM tool tbl
  CROSS JOIN JSON_TABLE(
    tbl.tool_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS price(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  price_type ENUM ('PRICE_CUSTOMER', 'PRICE_VENDOR') NOT NULL DEFAULT 'PRICE_CUSTOMER',
  valid_from DATE NOT NULL,
  valid_to DATE,
  product_code VARCHAR(255) NOT NULL,
  currency_code VARCHAR(255) NOT NULL,
  customer_code VARCHAR(255),
  qty INTEGER NOT NULL DEFAULT 0,
  price_meta JSON NOT NULL,
  price_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT price_pkey PRIMARY KEY (id),
  CONSTRAINT price_code_key UNIQUE (code),
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE UNIQUE INDEX idx_price_pvpc ON price (price_type, valid_from, product_code, currency_code, qty);
CREATE INDEX idx_price_price_type ON price (price_type);
CREATE INDEX idx_price_valid_from ON price (valid_from);
CREATE INDEX idx_price_valid_to ON price (valid_to);
CREATE INDEX idx_price_product_code ON price (product_code);
CREATE INDEX idx_price_currency_code ON price (currency_code);
CREATE INDEX idx_price_customer_code ON price (customer_code);
CREATE INDEX idx_price_deleted ON price (deleted);

CREATE TRIGGER price_default_code 
  BEFORE INSERT ON price
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM price);
    SET NEW.code = CONCAT('PRC', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW price_view AS
  SELECT id, code, price_type, valid_from, valid_to, product_code, currency_code, customer_code, qty,
    CAST(price_meta->>"$.price_value" AS FLOAT) AS price_value,
    price_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(price_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    price_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'price_type', price_type, 'valid_from', valid_from, 'valid_to', valid_to,
      'product_code', product_code, 'currency_code', currency_code, 'customer_code', customer_code, 
      'qty', qty, 'price_meta', price_meta, 'price_map', price_map, 'time_stamp', time_stamp
    ) AS price_object
  FROM price
  WHERE deleted = false;

CREATE VIEW price_map AS
  SELECT tbl.id AS id, tbl.code,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.price_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.price_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM price tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.price_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW price_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM price tbl
  CROSS JOIN JSON_TABLE(
    tbl.price_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS trans(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  trans_type ENUM (
  'TRANS_INVOICE', 'TRANS_RECEIPT', 'TRANS_ORDER', 'TRANS_OFFER', 'TRANS_WORKSHEET', 'TRANS_RENT', 'TRANS_DELIVERY', 'TRANS_INVENTORY', 
  'TRANS_WAYBILL', 'TRANS_PRODUCTION', 'TRANS_FORMULA', 'TRANS_BANK', 'TRANS_CASH') NOT NULL,
  direction ENUM ('DIRECTION_OUT', 'DIRECTION_IN', 'DIRECTION_TRANSFER') NOT NULL DEFAULT 'DIRECTION_OUT',
  trans_date DATE NOT NULL,
  trans_code VARCHAR(255),
  customer_code VARCHAR(255),
  employee_code VARCHAR(255),
  project_code VARCHAR(255),
  place_code VARCHAR(255),
  currency_code VARCHAR(255),
  auth_code VARCHAR(255) NOT NULL, 
  trans_meta JSON NOT NULL,
  trans_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT trans_pkey PRIMARY KEY (id),
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
    ON UPDATE RESTRICT ON DELETE RESTRICT
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

CREATE TRIGGER trans_default_code 
  BEFORE INSERT ON trans
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM trans);
    IF NEW.trans_type = 'TRANS_INVENTORY' THEN
      SET NEW.code = CONCAT('COR', UNIX_TIMESTAMP(), 'N', @A);
    ELSEIF NEW.trans_type = 'TRANS_DELIVERY' AND NEW.direction = 'DIRECTION_TRANSFER' THEN
      SET NEW.code = CONCAT('TRF', UNIX_TIMESTAMP(), 'N', @A);
    ELSE
      SET NEW.code = CONCAT(SUBSTR(NEW.trans_type, 7, 3), UNIX_TIMESTAMP(), 'N', @A);
    END IF;
  END IF;
END;
  
CREATE TRIGGER trans_invoice_customer_insert
  BEFORE INSERT ON trans
  FOR EACH ROW
BEGIN
  IF (NEW.trans_type = 'TRANS_INVOICE' AND NEW.direction = 'DIRECTION_OUT') THEN
    SET NEW.trans_meta = JSON_SET(NEW.trans_meta, '$.invoice', (
      SELECT JSON_OBJECT(
        'customer_name', cu.customer_name,
        'customer_tax_number', cu.customer_meta->>"$.tax_number",
        'customer_account', cu.customer_meta->>"$.account",
        'customer_address', TRIM(CONCAT_WS(' ',
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(cu.addresses, '$[0].zip_code')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(cu.addresses, '$[0].city')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(cu.addresses, '$[0].street')), '')
        )),
        'company_name', co.customer_name,
        'company_tax_number', co.customer_meta->>"$.tax_number",
        'company_account', co.customer_meta->>"$.account",
        'company_address', TRIM(CONCAT_WS(' ',
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(co.addresses, '$[0].zip_code')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(co.addresses, '$[0].city')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(co.addresses, '$[0].street')), '')
        ))
      )
      FROM customer cu 
      INNER JOIN customer co ON co.customer_type = 'CUSTOMER_OWN'
      WHERE cu.code = NEW.customer_code 
      LIMIT 1
    ));
  END IF;
END;

CREATE TRIGGER trans_invoice_customer_update
  BEFORE UPDATE ON trans
  FOR EACH ROW
BEGIN
  IF (NEW.trans_type = 'TRANS_INVOICE' AND NEW.direction = 'DIRECTION_OUT') THEN
    SET NEW.trans_meta = JSON_SET(NEW.trans_meta, '$.invoice', (
      SELECT JSON_OBJECT(
        'customer_name', cu.customer_name,
        'customer_tax_number', cu.customer_meta->>"$.tax_number",
        'customer_account', cu.customer_meta->>"$.account",
        'customer_address', TRIM(CONCAT_WS(' ',
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(cu.addresses, '$[0].zip_code')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(cu.addresses, '$[0].city')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(cu.addresses, '$[0].street')), '')
        )),
        'company_name', co.customer_name,
        'company_tax_number', co.customer_meta->>"$.tax_number",
        'company_account', co.customer_meta->>"$.account",
        'company_address', TRIM(CONCAT_WS(' ',
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(co.addresses, '$[0].zip_code')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(co.addresses, '$[0].city')), ''),
          COALESCE(JSON_UNQUOTE(JSON_EXTRACT(co.addresses, '$[0].street')), '')
        ))
      )
      FROM customer cu 
      INNER JOIN customer co ON co.customer_type = 'CUSTOMER_OWN'
      WHERE cu.code = NEW.customer_code 
      LIMIT 1
    ));
  END IF;
END;

CREATE VIEW trans_view AS
  SELECT id, code, trans_type, direction, trans_date, trans_code, customer_code, employee_code,
    project_code, place_code, currency_code, auth_code,
    CASE WHEN trans_meta->>"$.due_time" = '' THEN NULL ELSE STR_TO_DATE(trans_meta->>"$.due_time", '%Y-%m-%dT%H:%i:%s') END AS due_time,
    trans_meta->>"$.ref_number" AS ref_number,
    trans_meta->>"$.paid_type" AS paid_type, 
    (trans_meta->>"$.tax_free" = 'true') AS tax_free, (trans_meta->>"$.paid" = 'true') AS paid,
    CAST(trans_meta->>"$.rate" AS FLOAT) AS rate,
    (trans_meta->>"$.closed" = 'true') AS closed,
    trans_meta->>"$.status" AS status,
    trans_meta->>"$.trans_state" AS trans_state,
    trans_meta->>"$.notes" AS notes, trans_meta->>"$.internal_notes" AS internal_notes,
    trans_meta->>"$.report_notes" AS report_notes,
    CASE WHEN json_type(trans_meta->"$.worksheet") = 'OBJECT' THEN CAST(JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.worksheet"),"$.distance")) AS FLOAT) ELSE 0 END AS worksheet_distance,
    CASE WHEN json_type(trans_meta->"$.worksheet") = 'OBJECT' THEN CAST(JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.worksheet"),"$.repair")) AS FLOAT) ELSE 0 END AS worksheet_repair,
    CASE WHEN json_type(trans_meta->"$.worksheet") = 'OBJECT' THEN CAST(JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.worksheet"),"$.total")) AS FLOAT) ELSE 0 END AS worksheet_total,
    CASE WHEN json_type(trans_meta->"$.worksheet") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.worksheet"),"$.notes")) ELSE '' END AS worksheet_notes,
    CASE WHEN json_type(trans_meta->"$.rent") = 'OBJECT' THEN CAST(JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.rent"),"$.holiday")) AS FLOAT) ELSE 0 END AS rent_holiday,
    CASE WHEN json_type(trans_meta->"$.rent") = 'OBJECT' THEN CAST(JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.rent"),"$.bad_tool")) AS FLOAT) ELSE 0 END AS rent_bad_tool,
    CASE WHEN json_type(trans_meta->'$.rent') = 'OBJECT' THEN CAST(JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.rent"),"$.other")) AS FLOAT) ELSE 0 END AS rent_other,
    CASE WHEN json_type(trans_meta->"$.rent") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.rent"),"$.notes")) ELSE '' END AS rent_notes,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.company_name")) ELSE '' END AS invoice_company_name,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.company_account")) ELSE '' END AS invoice_company_account,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.company_address")) ELSE '' END AS invoice_company_address,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.company_tax_number")) ELSE '' END AS invoice_company_tax_number,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.customer_name")) ELSE '' END AS invoice_customer_name,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.customer_account")) ELSE '' END AS invoice_customer_account,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.customer_address")) ELSE '' END AS invoice_customer_address,
    CASE WHEN json_type(trans_meta->"$.invoice") = 'OBJECT' THEN JSON_UNQUOTE(JSON_EXTRACT(JSON_EXTRACT(trans_meta,"$.invoice"),"$.customer_tax_number")) ELSE '' END AS invoice_customer_tax_number,
    trans_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(trans_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    trans_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'trans_type', trans_type, 'direction', direction, 'trans_date', trans_date,
      'trans_code', trans_code, 'customer_code', customer_code, 'employee_code', employee_code,
      'project_code', project_code, 'place_code', place_code, 'currency_code', currency_code, 'auth_code', auth_code,
      'trans_meta', trans_meta, 'trans_map', trans_map, 'time_stamp', time_stamp
    ) AS trans_object
  FROM trans
  WHERE deleted = false;

CREATE VIEW trans_map AS
  SELECT tbl.id AS id, tbl.code, trans_type, direction, trans_date,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.trans_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.trans_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM trans tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.trans_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW trans_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM trans tbl
  CROSS JOIN JSON_TABLE(
    tbl.trans_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS item(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  trans_code VARCHAR(255) NOT NULL,
  product_code VARCHAR(255) NOT NULL,
  tax_code VARCHAR(255) NOT NULL, 
  item_meta JSON NOT NULL,
  item_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT item_pkey PRIMARY KEY (id),
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

CREATE TRIGGER item_default_code 
  BEFORE INSERT ON item
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM item);
    SET NEW.code = CONCAT('ITM', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW item_view AS
  SELECT id, code, trans_code, product_code, tax_code,
    item_meta->>"$.unit" AS unit, CAST(item_meta->>"$.qty" AS FLOAT) AS qty, CAST(item_meta->>"$.fx_price" AS FLOAT) AS fx_price,
    CAST(item_meta->>"$.net_amount" AS FLOAT) AS net_amount, CAST(item_meta->>"$.discount" AS FLOAT) AS discount, 
    CAST(item_meta->>"$.vat_amount" AS FLOAT) AS vat_amount, CAST(item_meta->>"$.amount" AS FLOAT) AS amount,
    item_meta->>"$.description" AS description, (item_meta->>"$.deposit" = 'true') AS deposit,
    (item_meta->>"$.action_price" = 'true') AS action_price, CAST(item_meta->>"$.own_stock" AS FLOAT) AS own_stock, 
    item_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(item_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    item_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'trans_code', trans_code, 'product_code', product_code,
      'tax_code', tax_code, 'item_meta', item_meta, 'item_map', item_map, 'time_stamp', time_stamp
    ) AS item_object
  FROM item
  WHERE deleted = false;

CREATE VIEW item_map AS
  SELECT tbl.id AS id, tbl.code, trans_code, product_code,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.item_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.item_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM item tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.item_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW item_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM item tbl
  CROSS JOIN JSON_TABLE(
    tbl.item_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS movement(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  movement_type ENUM ('MOVEMENT_INVENTORY', 'MOVEMENT_TOOL', 'MOVEMENT_PLAN', 'MOVEMENT_HEAD') NOT NULL DEFAULT 'MOVEMENT_INVENTORY',
  shipping_time TIMESTAMP NOT NULL,
  trans_code VARCHAR(255) NOT NULL,
  product_code VARCHAR(255),
  tool_code VARCHAR(255),
  place_code VARCHAR(255), 
  item_code VARCHAR(255),
  movement_code VARCHAR(255),
  movement_meta JSON NOT NULL,
  movement_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT movement_pkey PRIMARY KEY (id),
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
    ON UPDATE RESTRICT ON DELETE SET NULL
);

CREATE INDEX idx_movement_trans_code ON movement (trans_code);
CREATE INDEX idx_movement_product_code ON movement (product_code);
CREATE INDEX idx_movement_tool_code ON movement (tool_code);
CREATE INDEX idx_movement_place_code ON movement (place_code);
CREATE INDEX idx_movement_item_code ON movement (item_code);
CREATE INDEX idx_movement_movement_code ON movement (movement_code);
CREATE INDEX idx_movement_deleted ON movement (deleted);

CREATE TRIGGER movement_default_code 
  BEFORE INSERT ON movement
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM movement);
    SET NEW.code = CONCAT('MOV', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW movement_view AS
  SELECT id, code, movement_type, shipping_time,
    trans_code, product_code, tool_code, place_code, item_code, movement_code,
    movement_meta->>"$.qty" AS qty, movement_meta->>"$.shared" AS shared,
    movement_meta->>"$.notes" AS notes,
    movement_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(movement_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    movement_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'movement_type', movement_type, 'shipping_time', shipping_time,
      'trans_code', trans_code, 'product_code', product_code, 'tool_code', tool_code, 'place_code', place_code,
      'item_code', item_code, 'movement_code', movement_code,
      'movement_meta', movement_meta, 'movement_map', movement_map, 'time_stamp', time_stamp
    ) AS movement_object
  FROM movement
  WHERE deleted = false;

CREATE VIEW movement_map AS
  SELECT tbl.id AS id, tbl.code, trans_code,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.movement_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.movement_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM movement tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.movement_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW movement_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM movement tbl
  CROSS JOIN JSON_TABLE(
    tbl.movement_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE VIEW movement_stock AS
  SELECT ROW_NUMBER() OVER (ORDER BY pl.place_name, p.product_name) as id,
    mv.place_code, pl.place_name, mv.product_code, p.product_name, 
    p.product_meta->>"$.unit" AS unit, movement_meta->>"$.notes" AS batch_no, 
    SUM(CAST(movement_meta->>"$.qty" AS FLOAT)) AS qty, 
    MAX(date(mv.shipping_time)) AS posdate
  FROM movement mv INNER JOIN place pl ON mv.place_code = pl.code
  INNER JOIN product p ON mv.product_code = p.code
  WHERE mv.movement_type = 'MOVEMENT_INVENTORY' AND mv.deleted = false AND p.deleted = false AND pl.deleted = false
  GROUP BY mv.place_code, pl.place_name, mv.product_code, p.product_name, p.product_meta->>"$.unit", movement_meta->>"$.notes"
  HAVING SUM(CAST(mv.movement_meta->>"$.qty" AS FLOAT)) <> 0
  ORDER BY pl.place_name, p.product_name;

CREATE VIEW movement_inventory AS
  SELECT mt.id, mt.code, mt.trans_code, t.trans_type, t.direction, DATE(mt.shipping_time) AS shipping_date,
    mt.place_code, pl.place_name, mt.product_code, p.product_name,
    p.product_meta->>"$.unit" AS unit, mt.movement_meta->>"$.notes" AS batch_no, 
    CAST(mt.movement_meta->>"$.qty" AS FLOAT) AS qty, 
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
    mv.tool_code, tl.tool_meta->>"$.serial_number" as serial_number, tl.description,
    mv.movement_meta->>"$.notes" as mvnotes,
    t.employee_code, t.customer_code, c.customer_name, 
    t.trans_meta->>"$.trans_state" as trans_state, trans_meta->>"$.notes" AS notes, 
    trans_meta->>"$.internal_notes" AS internal_notes, 
    (trans_meta->>"$.closed" = 'true') AS closed,t.time_stamp
  FROM trans t INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN tool tl ON mv.tool_code = tl.code
  LEFT JOIN customer c ON t.customer_code = c.code
  WHERE t.trans_type = 'TRANS_WAYBILL' and mv.deleted = false and t.deleted = false;

CREATE VIEW movement_formula AS
  SELECT mv.id, mv.code, t.code AS trans_code, CASE WHEN mv.movement_type = 'MOVEMENT_HEAD' THEN 'IN' ELSE 'OUT' END as direction, 
    mv.product_code, p.product_name, p.product_meta->>"$.unit" as unit,
    CAST(mv.movement_meta->>"$.qty" AS FLOAT) AS qty, mv.movement_meta->>"$.notes" as batch_no,
    mv.place_code, pl.place_name, (mv.movement_meta->>"$.shared" = 'true') AS shared
  FROM trans t INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN product p ON mv.product_code = p.code
  LEFT JOIN place pl ON mv.place_code = pl.code
  WHERE t.trans_type = 'TRANS_FORMULA' and mv.deleted = false and t.deleted = false;

CREATE VIEW item_shipping AS
  SELECT iv.id, iv.code, iv.trans_code, iv.direction, iv.product_code, iv.product_name, iv.unit, iv.item_qty,
    SUM(cast(COALESCE(movement_meta->>"$.qty",'0') AS FLOAT)) as movement_qty
  FROM (
  SELECT i.id, i.code, i.trans_code, t.direction, i.product_code, p.product_name, 
    COALESCE(p.product_meta->>"$.unit",'') as unit, 
    cast(COALESCE(item_meta->>"$.qty",'0') AS FLOAT) item_qty
  FROM item i
  INNER JOIN trans t ON i.trans_code = t.code
  INNER JOIN product p ON i.product_code = p.code
  WHERE t.deleted = false AND i.deleted = false AND p.product_type = 'PRODUCT_ITEM' 
    AND t.trans_type IN('TRANS_ORDER', 'TRANS_WORKSHEET', 'TRANS_RENT')
  UNION 
  SELECT i.id, i.code, i.trans_code, t.direction, pc.ref_product_code as product_code, pc.component_name as product_name, 
    pc.component_unit as unit, 
    cast(COALESCE(item_meta->>"$.qty",'0') AS FLOAT)*pc.qty item_qty
  FROM item i
  INNER JOIN trans t ON i.trans_code = t.code
  INNER JOIN product_components pc ON i.product_code = pc.product_code AND pc.component_type = 'PRODUCT_ITEM'
  WHERE t.deleted = false AND i.deleted = false 
    AND t.trans_type IN('TRANS_ORDER', 'TRANS_WORKSHEET', 'TRANS_RENT')
  ) iv
  LEFT JOIN movement mv ON mv.item_code = iv.code AND mv.product_code = iv.product_code
  GROUP BY iv.id, iv.code, iv.trans_code, iv.direction, iv.product_code, iv.product_name, iv.unit, iv.item_qty;

CREATE TABLE IF NOT EXISTS payment(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  paid_date DATE NOT NULL,
  trans_code VARCHAR(255) NOT NULL, 
  payment_meta JSON NOT NULL,
  payment_map JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT payment_pkey PRIMARY KEY (id),
  CONSTRAINT payment_code_key UNIQUE (code),
  FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_payment_trans_code ON payment (trans_code);
CREATE INDEX idx_payment_deleted ON payment (deleted);

CREATE TRIGGER payment_default_code 
  BEFORE INSERT ON payment
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM payment);
    SET NEW.code = CONCAT('PMT', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;

CREATE VIEW payment_view AS
  SELECT id, code, paid_date, trans_code,
    payment_meta->>"$.amount" AS amount, payment_meta->>"$.notes" AS notes,
    payment_meta->"$.tags" AS tags, 
    REPLACE(REPLACE(REPLACE(payment_meta->>"$.tags", '"', ''), '[', ''), ']', '') as tag_lst,
    payment_map, time_stamp,
    JSON_OBJECT(
      'id', id, 'code', code, 'paid_date', paid_date, 'trans_code', trans_code,
      'payment_meta', payment_meta, 'payment_map', payment_map, 'time_stamp', time_stamp
    ) AS payment_object
  FROM payment
  WHERE deleted = false;

CREATE VIEW payment_map AS
  SELECT tbl.id AS id, tbl.code, paid_date, trans_code,
    jt.map_key, JSON_UNQUOTE(JSON_EXTRACT(tbl.payment_map, CONCAT('$.', jt.map_key))) AS map_value,
    JSON_TYPE(JSON_EXTRACT(tbl.payment_map, CONCAT('$.', jt.map_key))) AS map_type,
    COALESCE(cf.description, jt.map_key) AS description,
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
  FROM payment tbl
  CROSS JOIN JSON_TABLE(
    JSON_KEYS(tbl.payment_map),
    '$[*]' COLUMNS (map_key VARCHAR(255) PATH '$')
  ) AS jt
  LEFT JOIN config_map cf on jt.map_key = cf.field_name
  WHERE deleted = false;

CREATE VIEW payment_invoice AS
  SELECT pm.id, pm.code, pm.trans_code, pt.trans_type, pt.direction, pm.paid_date, pl.place_name, pl.currency_code,
    cast(l.link_meta->>"$.amount" AS FLOAT) AS paid_amount, cast(l.link_meta->>"$.rate" AS FLOAT) AS paid_rate, 
    it.code AS ref_trans_code, it.currency_code AS invoice_curr,
    im.amount AS invoice_amount, pm.payment_meta->>"$.notes" AS description
  FROM link l 
  INNER JOIN payment pm ON l.link_code_1 = pm.code
  INNER JOIN trans pt ON pm.trans_code = pt.code INNER JOIN place pl ON pt.place_code = pl.code
  INNER JOIN trans it ON l.link_code_2 = it.code
  INNER JOIN(
    SELECT trans_code, sum(cast(item_meta->>"$.amount" AS FLOAT)) AS amount FROM item GROUP BY trans_code) im ON it.code = im.trans_code
  WHERE l.link_type_1 = 'LINK_PAYMENT' AND l.link_type_2 = 'LINK_TRANS' AND it.trans_type IN('TRANS_INVOICE','TRANS_RECEIPT');

CREATE VIEW payment_tags AS
  SELECT tbl.id AS id, tbl.code, jt.value as tag
  FROM payment tbl
  CROSS JOIN JSON_TABLE(
    tbl.payment_meta->"$.tags", "$[*]" COLUMNS(value VARCHAR(255) PATH "$")
  ) AS jt
  WHERE tbl.deleted = false;

CREATE TABLE IF NOT EXISTS log(
  id INTEGER AUTO_INCREMENT NOT NULL,
  code VARCHAR(255) NOT NULL,
  log_type ENUM ('LOG_INSERT', 'LOG_UPDATE', 'LOG_DELETE') NOT NULL DEFAULT 'LOG_UPDATE',
  ref_type VARCHAR(255) NOT NULL DEFAULT 'TRANS',
  ref_code VARCHAR(255) NOT NULL,
  auth_code VARCHAR(255) NOT NULL,
  data JSON NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT log_pkey PRIMARY KEY (id),
  CONSTRAINT log_code_key UNIQUE (code)
);

CREATE INDEX idx_log_ref_type ON log (ref_type);
CREATE INDEX idx_log_ref_code ON log (ref_code);
CREATE INDEX idx_log_auth_code ON log (auth_code);
CREATE INDEX idx_log_time_stamp ON log (time_stamp);
CREATE INDEX idx_log_deleted ON log (deleted);

CREATE TRIGGER log_default_code 
  BEFORE INSERT ON log
  FOR EACH ROW
BEGIN
  IF NEW.code IS NULL THEN
    SET @A = (SELECT CASE WHEN COUNT(*)=0 THEN 1 ELSE MAX(id) + 1 END FROM log);
    SET NEW.code = CONCAT('LOG', UNIX_TIMESTAMP(), 'N', @A);
  END IF;
END;
