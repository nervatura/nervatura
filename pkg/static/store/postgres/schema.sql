CREATE TYPE config_type AS ENUM ('CONFIG_MAP', 'CONFIG_SHORTCUT', 'CONFIG_MESSAGE', 'CONFIG_PATTERN', 'CONFIG_REPORT', 'CONFIG_PRINT_QUEUE', 'CONFIG_DATA');
CREATE TYPE user_group AS ENUM ('GROUP_ADMIN', 'GROUP_USER', 'GROUP_GUEST');
CREATE TYPE customer_type AS ENUM ('CUSTOMER_OWN', 'CUSTOMER_COMPANY', 'CUSTOMER_PRIVATE', 'CUSTOMER_OTHER');
CREATE TYPE place_type AS ENUM ('PLACE_BANK', 'PLACE_CASH', 'PLACE_WAREHOUSE', 'PLACE_OTHER');
CREATE TYPE product_type AS ENUM ('PRODUCT_ITEM', 'PRODUCT_SERVICE', 'PRODUCT_VIRTUAL');
CREATE TYPE rate_type AS ENUM ('RATE_RATE', 'RATE_BUY', 'RATE_SELL', 'RATE_AVERAGE');
CREATE TYPE price_type AS ENUM ('PRICE_CUSTOMER', 'PRICE_VENDOR');
CREATE TYPE trans_type AS ENUM (
  'TRANS_INVOICE', 'TRANS_RECEIPT', 'TRANS_ORDER', 'TRANS_OFFER', 'TRANS_WORKSHEET', 'TRANS_RENT', 'TRANS_DELIVERY', 'TRANS_INVENTORY', 
  'TRANS_WAYBILL', 'TRANS_PRODUCTION', 'TRANS_FORMULA', 'TRANS_BANK', 'TRANS_CASH');
CREATE TYPE direction AS ENUM ('DIRECTION_OUT', 'DIRECTION_IN', 'DIRECTION_TRANSFER'); 
CREATE TYPE movement_type AS ENUM ('MOVEMENT_INVENTORY', 'MOVEMENT_TOOL', 'MOVEMENT_PLAN', 'MOVEMENT_HEAD');
CREATE TYPE log_type AS ENUM ('LOG_INSERT', 'LOG_UPDATE', 'LOG_DELETE');
CREATE TYPE link_type AS ENUM ('LINK_CUSTOMER', 'LINK_EMPLOYEE', 'LINK_ITEM', 'LINK_MOVEMENT', 'LINK_PAYMENT', 'LINK_PLACE', 'LINK_PRODUCT', 'LINK_PROJECT', 'LINK_TOOL', 'LINK_TRANS');

CREATE OR REPLACE FUNCTION set_new_code() 
  RETURNS TRIGGER LANGUAGE PLPGSQL AS $BODY$
BEGIN
  NEW.code := TG_ARGV[0] || EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)::INTEGER || 'N' || NEW.id;
  RETURN NEW;
END;
$BODY$;

CREATE OR REPLACE FUNCTION set_trans_code() 
  RETURNS TRIGGER LANGUAGE PLPGSQL AS $BODY$
DECLARE
  _prefix TEXT; 
BEGIN
  _prefix := substr(NEW.trans_type::text, 7, 3);
  IF NEW.trans_type = 'TRANS_INVENTORY'::trans_type THEN
    _prefix = 'COR';
  ELSIF NEW.trans_type = 'TRANS_DELIVERY' AND NEW.direction = 'DIRECTION_TRANSFER' THEN
    _prefix = 'TRF';
  END IF;
  NEW.code := _prefix || EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)::INTEGER || 'N' || NEW.id;
  RETURN NEW;
END;
$BODY$;

CREATE OR REPLACE FUNCTION set_invoice_customer() 
  RETURNS TRIGGER LANGUAGE PLPGSQL AS $BODY$
DECLARE
  _id INTEGER;
  _name TEXT;
  _tax_number TEXT;
  _address TEXT;
  _account TEXT;
  _invoice JSONB;
BEGIN
  -- Get company data
  SELECT id, customer_name, customer_meta->>'tax_number', customer_meta->>'account', concat_ws(' ',
    case when addresses[0] is null then '' else addresses[0]->>'zip_code' end,
	  case when addresses[0] is null then '' else addresses[0]->>'city' end,
	  case when addresses[0] is null then '' else addresses[0]->>'street' end)
    INTO _id, _name, _tax_number, _account, _address
  FROM customer WHERE customer_type = 'CUSTOMER_OWN'::customer_type;
  -- Check if company data exists
  IF _id IS NULL THEN
    RAISE EXCEPTION 'Missing company data'
      USING HINT = 'Missing customer_type OWN (own company) data';
  END IF;
  -- Build invoice JSON
  _invoice = jsonb_build_object('company_name', _name, 'company_tax_number', _tax_number, 'company_account', _account, 'company_address', _address);
  -- Get customer data
  SELECT id, customer_name, customer_meta->>'tax_number', customer_meta->>'account', concat_ws(' ',
    case when addresses[0] is null then '' else addresses[0]->>'zip_code' end,
	  case when addresses[0] is null then '' else addresses[0]->>'city' end,
	  case when addresses[0] is null then '' else addresses[0]->>'street' end)
    INTO _id, _name, _tax_number, _account, _address
  FROM customer WHERE code = NEW.customer_code;
  -- Check if customer code is valid
  IF _id IS NULL THEN
    RAISE EXCEPTION 'Invalid customer code: %', _arr_id
      USING HINT = 'Valid values: customer table code.';
  END IF;
  -- Append customer data to invoice JSON
  _invoice = _invoice || jsonb_build_object('customer_name', _name, 'customer_tax_number', _tax_number, 'customer_address', _address);
  -- Update the trans_meta field in the NEW row
  NEW.trans_meta = NEW.trans_meta || jsonb_build_object('invoice', _invoice);
  
  RETURN NEW;
END;
$BODY$;

CREATE OR REPLACE FUNCTION link_changed()
  RETURNS TRIGGER LANGUAGE PLPGSQL AS $BODY$
DECLARE
  _table_name varchar; 
  _ref_id integer;
BEGIN
  _table_name := SPLIT_PART(LOWER(NEW.link_type_1::text), '_', 2);
  EXECUTE 'SELECT id FROM ' || _table_name || ' WHERE code=$1'
    INTO _ref_id USING NEW.link_code_1;
  IF _ref_id IS NULL THEN
    RAISE EXCEPTION 'Invalid % code: %', _table_name, NEW.link_code_1;
  END IF;

  _table_name := SPLIT_PART(LOWER(NEW.link_type_2::text), '_', 2);
  EXECUTE 'SELECT id FROM ' || _table_name || ' WHERE code=$1'
    INTO _ref_id USING NEW.link_code_2;
  IF _ref_id IS NULL THEN
    RAISE EXCEPTION 'Invalid % code: %', _table_name, NEW.link_code_2;
  END IF;

  RETURN NEW;
END;
$BODY$;

CREATE OR REPLACE FUNCTION set_changed_timestamp()
  RETURNS TRIGGER LANGUAGE PLPGSQL AS $BODY$
BEGIN
  NEW.changed = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$BODY$;

CREATE TABLE IF NOT EXISTS usref(
  id SERIAL NOT NULL,
  refnumber TEXT NOT NULL,
  value TEXT NOT NULL,
  changed TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT usref_pkey PRIMARY KEY (id),
  CONSTRAINT usref_refnumber_key UNIQUE (refnumber)
);

CREATE OR REPLACE TRIGGER usref_changed_update
  AFTER UPDATE ON usref
  FOR EACH ROW
  EXECUTE FUNCTION set_changed_timestamp();

CREATE TABLE IF NOT EXISTS config(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  config_type config_type NOT NULL DEFAULT 'CONFIG_MAP'::config_type, 
  data JSONB NOT NULL,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT config_pkey PRIMARY KEY (id),
  CONSTRAINT config_code_key UNIQUE (code)
);

CREATE INDEX idx_config_type ON config (config_type);
CREATE INDEX idx_config_deleted ON config (deleted);

CREATE OR REPLACE TRIGGER config_default_code
  BEFORE INSERT ON config
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('CNF');

CREATE OR REPLACE VIEW config_map AS
  SELECT id, code,
    data->>'field_name' AS field_name, data->>'field_type' AS field_type, 
	  data->>'description' AS description, data->'tags' AS tags, data->'filter' AS filter
  FROM config
  WHERE config_type = 'CONFIG_MAP' AND deleted = false;

CREATE OR REPLACE VIEW config_shortcut AS
  SELECT id, code,
    data->>'shortcut_key' AS shortcut_key, data->>'description' AS description, 
	  data->>'modul' AS modul, data->>'method' AS method, data->>'func_name' AS func_name,
    data->>'address' AS address, data->'fields' AS fields
  FROM config
  WHERE config_type = 'CONFIG_SHORTCUT' AND deleted = false;

CREATE OR REPLACE VIEW config_message AS
  SELECT id, code,
    data->>'section' AS section, data->>'key' AS message_key, 
	  data->>'lang' AS lang, data->>'value' AS message_value
  FROM config
  WHERE config_type = 'CONFIG_MESSAGE' AND deleted = false;

CREATE OR REPLACE VIEW config_pattern AS
  SELECT id, code,
    data->>'trans_type' AS trans_type, data->>'description' AS description, 
	  data->>'notes' AS notes, data->>'default_pattern' AS default_pattern
  FROM config
  WHERE config_type = 'CONFIG_PATTERN' AND deleted = false;

CREATE OR REPLACE VIEW config_print_queue AS
  SELECT id, code,
    data->>'ref_type' AS ref_type, data->>'ref_code' AS ref_code,
    cast(data->>'qty' AS FLOAT) AS qty, data->>'report_code' AS report_code,
    data->>'orientation' AS orientation, data->>'paper_size' AS paper_size,
	  data->>'auth_code' AS auth_code, data->>'time_stamp' AS time_stamp
  FROM config
  WHERE config_type = 'CONFIG_PRINT_QUEUE' AND deleted = false;

CREATE OR REPLACE VIEW config_report AS
  SELECT id, code,
    data->>'report_key' AS report_key, data->>'report_type' AS report_type, 
	  data->>'trans_type' AS trans_type, data->>'direction' AS direction,
	  data->>'report_name' AS report_name, data->>'description' AS description,
	  data->>'label' AS label, data->>'file_type' AS file_type,
	  data->'template' AS template
  FROM config
  WHERE config_type = 'CONFIG_REPORT' AND deleted = false;
  
CREATE OR REPLACE VIEW config_data AS
  SELECT ROW_NUMBER() OVER (ORDER BY code, key) as id, code AS config_code,
    key as config_key, value as config_value, jsonb_typeof(value) as config_type
  FROM config, jsonb_each(config.data)
  WHERE config_type = 'CONFIG_DATA' AND jsonb_typeof(data) = 'object' AND deleted = false;

CREATE TABLE IF NOT EXISTS auth(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL, 
  user_name VARCHAR NOT NULL,
  user_group user_group NOT NULL DEFAULT 'GROUP_USER'::user_group,
  disabled BOOLEAN NOT NULL DEFAULT false,
  auth_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  auth_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT auth_pkey PRIMARY KEY (id),
  CONSTRAINT auth_code_key UNIQUE (code)
);

CREATE UNIQUE INDEX idx_auth_user_name ON auth (user_name);
CREATE INDEX idx_auth_deleted ON auth (deleted);
CREATE INDEX idx_auth_disabled ON auth (disabled);
CREATE INDEX idx_auth_tags ON auth ((auth_meta->>'tags'));

CREATE OR REPLACE TRIGGER auth_default_code
  BEFORE INSERT ON auth
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('USR');

CREATE OR REPLACE VIEW auth_view AS
  SELECT id, code, user_name, user_group, disabled,
    auth_meta->'tags' AS tags, REGEXP_REPLACE(auth_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, auth_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'user_name', user_name, 'user_group', user_group, 'disabled', disabled,
      'auth_meta', auth_meta, 'auth_map', auth_map, 'time_stamp', time_stamp
    ) AS auth_object
  FROM auth
  WHERE deleted = false;

CREATE OR REPLACE VIEW auth_map AS
  SELECT auth.id AS id, auth.code, user_name, user_group,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM auth, jsonb_each(auth.auth_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS currency(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  currency_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  currency_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT currency_pkey PRIMARY KEY (id),
  CONSTRAINT currency_code_key UNIQUE (code)
);

CREATE INDEX idx_currency_deleted ON currency (deleted);
CREATE INDEX idx_currency_tags ON currency ((currency_meta->>'tags'));

CREATE OR REPLACE VIEW currency_view AS
  SELECT id, code,
    currency_meta->>'description' AS description,
    cast(currency_meta->>'digit' AS INTEGER) AS digit, cast(currency_meta->>'cash_round' AS INTEGER) AS cash_round, 
    currency_meta->'tags' AS tags, REGEXP_REPLACE(currency_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, currency_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'currency_meta', currency_meta, 'currency_map', currency_map, 'time_stamp', time_stamp
    ) AS currency_object
  FROM currency
  WHERE deleted = false;

CREATE OR REPLACE VIEW currency_map AS
  SELECT currency.id AS id, currency.code, currency_meta->>'description' AS currency_description,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM currency, jsonb_each(currency.currency_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW currency_tags AS
  SELECT currency.id AS id, code, value as tag
  FROM currency, jsonb_array_elements_text(currency.currency_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS customer(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  customer_type customer_type NOT NULL DEFAULT 'CUSTOMER_COMPANY'::customer_type,
  customer_name VARCHAR NOT NULL,
  addresses JSONB NOT NULL DEFAULT '[]'::JSONB,
  contacts JSONB NOT NULL DEFAULT '[]'::JSONB,
  events JSONB NOT NULL DEFAULT '[]'::JSONB, 
  customer_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  customer_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT customer_pkey PRIMARY KEY (id),
  CONSTRAINT customer_code_key UNIQUE (code)
);

CREATE INDEX idx_customer_customer_type ON customer (customer_type);
CREATE INDEX idx_customer_deleted ON customer (deleted);
CREATE INDEX idx_customer_name ON customer (customer_name);
CREATE INDEX idx_customer_inactive ON customer ((customer_meta->>'inactive'));
CREATE INDEX idx_customer_tags ON customer ((customer_meta->>'tags'));

CREATE OR REPLACE TRIGGER customer_default_code
  BEFORE INSERT ON customer
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('CUS');

CREATE OR REPLACE VIEW customer_view AS
  SELECT id, code, customer_type, customer_name, customer_meta->>'tax_number' AS tax_number, customer_meta->>'account' AS account,
    cast(customer_meta->>'tax_free' AS BOOLEAN) AS tax_free, cast(customer_meta->>'terms' AS INTEGER) AS terms,
    cast(customer_meta->>'credit_limit' AS FLOAT) AS credit_limit, cast(customer_meta->>'discount' AS FLOAT) AS discount,
    customer_meta->>'notes' AS notes, cast(customer_meta->>'inactive' AS BOOLEAN) AS inactive,
    customer_meta->'tags' AS tags, REGEXP_REPLACE(customer_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst,
    addresses, contacts, events, customer_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'customer_type', customer_type, 'customer_name', customer_name,
      'addresses', addresses, 'contacts', contacts, 'events', events, 
      'customer_meta', customer_meta, 'customer_map', customer_map, 'time_stamp', time_stamp
    ) AS customer_object
  FROM customer
  WHERE deleted = false;

CREATE OR REPLACE VIEW customer_contacts AS
  SELECT customer.id AS id, code, customer_name, 
    value->>'first_name' AS first_name, value->>'surname' AS surname, value->>'status' AS status,
    value->>'phone' AS phone, value->>'mobile' AS mobile, value->>'email' AS email, 
    value->>'notes' AS notes, value->'tags' AS tags, REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'contact_map' AS contact_map
  FROM customer, jsonb_array_elements(customer.contacts)
  WHERE deleted = false;

CREATE OR REPLACE VIEW customer_addresses AS
  SELECT customer.id AS id, code, customer_name, 
    value->>'country' AS country, value->>'state' AS state, value->>'zip_code' AS zip_code,
    value->>'city' AS city, value->>'street' AS street, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'address_map' AS address_map
  FROM customer, jsonb_array_elements(customer.addresses)
  WHERE deleted = false;

CREATE OR REPLACE VIEW customer_events AS
  SELECT customer.id AS id, code, customer_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, 
	  CASE WHEN value->>'start_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'start_time', 'YYYY-MM-DDTHH24:MI:SS') END AS start_time,
	  CASE WHEN value->>'end_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'end_time', 'YYYY-MM-DDTHH24:MI:SS') END AS end_time,
	  value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'event_map' AS event_map
  FROM customer, jsonb_array_elements(customer.events)
  WHERE deleted = false;

CREATE OR REPLACE VIEW customer_map AS
  SELECT customer.id AS id, customer.code, customer.customer_name,
    key AS map_key, value AS map_value, jsonb_typeof(value) AS map_type,
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
  FROM customer, jsonb_each(customer.customer_map) LEFT JOIN config_map cf ON key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW customer_tags AS
  SELECT customer.id AS id, code, value as tag
  FROM customer, jsonb_array_elements_text(customer.customer_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS employee(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL, 
  address JSONB NOT NULL DEFAULT '{}'::JSONB,
  contact JSONB NOT NULL DEFAULT '{}'::JSONB,
  events JSONB NOT NULL DEFAULT '[]'::JSONB, 
  employee_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  employee_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT employee_pkey PRIMARY KEY (id),
  CONSTRAINT employee_code_key UNIQUE (code)
);

CREATE INDEX idx_employee_deleted ON employee (deleted);
CREATE INDEX idx_employee_name ON employee ((employee_meta->>'name'));
CREATE INDEX idx_employee_inactive ON employee ((employee_meta->>'inactive'));
CREATE INDEX idx_employee_tags ON employee ((employee_meta->>'tags'));

CREATE OR REPLACE TRIGGER employee_default_code
  BEFORE INSERT ON employee
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('EMP');

CREATE OR REPLACE VIEW employee_view AS
  SELECT id, code, 
	  CASE WHEN employee_meta->>'start_date' = '' THEN NULL ELSE TO_DATE(employee_meta->>'start_date', 'YYYY-MM-DD') END AS start_date, 
	  CASE WHEN employee_meta->>'end_date' = '' THEN NULL ELSE TO_DATE(employee_meta->>'end_date', 'YYYY-MM-DD') END AS end_date, 
    cast(employee_meta->>'inactive' AS BOOLEAN) AS inactive, employee_meta->>'notes' AS notes,
    contact->>'first_name' AS first_name, contact->>'surname' AS surname, contact->>'status' AS status,
    contact->>'phone' AS phone, contact->>'mobile' AS mobile, contact->>'email' AS email, 
    address->>'country' AS country, address->>'state' AS state, address->>'zip_code' AS zip_code,
    address->>'city' AS city, address->>'street' AS street, 
    employee_meta->'tags' AS tags, REGEXP_REPLACE(employee_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    employee_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'contact', contact, 'address', address, 'events', events,
      'employee_meta', employee_meta, 'employee_map', employee_map, 'time_stamp', time_stamp
    ) AS employee_object
  FROM employee
  WHERE deleted = false;

CREATE OR REPLACE VIEW employee_events AS
  SELECT employee.id AS id, code,
    contact->>'first_name' AS first_name, contact->>'surname' AS surname,
    value->>'uid' AS uid, value->>'subject' AS subject, 
    CASE WHEN value->>'start_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'start_time', 'YYYY-MM-DDTHH24:MI:SS') END AS start_time,
	  CASE WHEN value->>'end_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'end_time', 'YYYY-MM-DDTHH24:MI:SS') END AS end_time,
    value->>'place' AS place, value->>'description' AS description, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'event_map' AS event_map
  FROM employee, jsonb_array_elements(employee.events)
  WHERE deleted = false;

CREATE OR REPLACE VIEW employee_map AS
  SELECT employee.id AS id, employee.code, 
    contact->>'first_name' AS first_name, contact->>'surname' AS surname,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM employee, jsonb_each(employee.employee_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW employee_tags AS
  SELECT employee.id AS id, code, value as tag
  FROM employee, jsonb_array_elements_text(employee.employee_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS place(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  place_type place_type NOT NULL DEFAULT 'PLACE_WAREHOUSE'::place_type,
  place_name VARCHAR NOT NULL,
  currency_code VARCHAR,
  address JSONB NOT NULL DEFAULT '{}'::JSONB,
  contacts JSONB NOT NULL DEFAULT '[]'::JSONB,
  events JSONB NOT NULL DEFAULT '[]'::JSONB,
  place_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  place_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT place_pkey PRIMARY KEY (id),
  CONSTRAINT place_code_key UNIQUE (code),
  FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_place_deleted ON place (deleted);
CREATE INDEX idx_place_place_name ON place (place_name);
CREATE INDEX idx_place_inactive ON place ((place_meta->>'inactive'));
CREATE INDEX idx_place_tags ON place ((place_meta->>'tags'));

CREATE OR REPLACE TRIGGER place_default_code
  BEFORE INSERT ON place
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('PLA');

CREATE OR REPLACE VIEW place_view AS
  SELECT id, code, place_type, place_name, currency_code,
    cast(place_meta->>'inactive' AS BOOLEAN) AS inactive, place_meta->>'notes' AS notes,
    address->>'country' AS country, address->>'state' AS state, address->>'zip_code' AS zip_code,
    address->>'city' AS city, address->>'street' AS street, 
    place_meta->'tags' AS tags, REGEXP_REPLACE(place_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    place_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'place_type', place_type, 'place_name', place_name,
      'currency_code', currency_code, 'address', address, 'contacts', contacts,
      'events', events,
      'place_meta', place_meta, 'place_map', place_map, 'time_stamp', time_stamp
    ) AS place_object
  FROM place
  WHERE deleted = false;

CREATE OR REPLACE VIEW place_contacts AS
  SELECT place.id AS id, code, place_name, 
    value->>'first_name' AS first_name, value->>'surname' AS surname, value->>'status' AS status,
    value->>'phone' AS phone, value->>'mobile' AS mobile, value->>'email' AS email, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'contact_map' AS contact_map
  FROM place, jsonb_array_elements(place.contacts)
  WHERE deleted = false;

CREATE OR REPLACE VIEW place_events AS
  SELECT place.id AS id, code, place_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, 
	  CASE WHEN value->>'start_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'start_time', 'YYYY-MM-DDTHH24:MI:SS') END AS start_time,
	  CASE WHEN value->>'end_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'end_time', 'YYYY-MM-DDTHH24:MI:SS') END AS end_time,
	  value->>'place' AS place, 
    value->>'description' AS description, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'event_map' AS event_map
  FROM place, jsonb_array_elements(place.events)
  WHERE deleted = false;

CREATE OR REPLACE VIEW place_map AS
  SELECT place.id AS id, place.code, place.place_name,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM place, jsonb_each(place.place_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW place_tags AS
  SELECT place.id AS id, code, value as tag
  FROM place, jsonb_array_elements_text(place.place_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS tax(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL, 
  tax_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  tax_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT tax_pkey PRIMARY KEY (id),
  CONSTRAINT tax_code_key UNIQUE (code)
);

CREATE INDEX idx_tax_deleted ON tax (deleted);
CREATE INDEX idx_tax_inactive ON tax ((tax_meta->>'inactive'));
CREATE INDEX idx_tax_tags ON tax ((tax_meta->>'tags'));

CREATE OR REPLACE VIEW tax_view AS
  SELECT id, code,
    tax_meta->>'description' AS description, cast(tax_meta->>'rate_value' AS FLOAT) AS rate_value,
    cast(tax_meta->>'inactive' AS BOOLEAN) AS inactive,
    tax_meta->'tags' AS tags, REGEXP_REPLACE(tax_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    tax_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'tax_meta', tax_meta, 'tax_map', tax_map, 'time_stamp', time_stamp
    ) AS tax_object
  FROM tax
  WHERE deleted = false;

CREATE OR REPLACE VIEW tax_map AS
  SELECT tax.id AS id, tax.code, tax_meta->>'description' AS tax_description,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM tax, jsonb_each(tax.tax_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW tax_tags AS
  SELECT tax.id AS id, code, value as tag
  FROM tax, jsonb_array_elements_text(tax.tax_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS link(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  link_type_1 link_type NOT NULL DEFAULT 'LINK_TRANS'::link_type,
  link_code_1 VARCHAR NOT NULL,
  link_type_2 link_type NOT NULL DEFAULT 'LINK_TRANS'::link_type,
  link_code_2 VARCHAR NOT NULL, 
  link_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  link_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT link_pkey PRIMARY KEY (id),
  CONSTRAINT link_code_key UNIQUE (code)
);

CREATE INDEX idx_link_link_code_1 ON link (link_type_1, link_code_1);
CREATE INDEX idx_link_link_code_2 ON link (link_type_2, link_code_2);
CREATE INDEX idx_link_deleted ON link (deleted);
CREATE INDEX idx_link_tags ON link ((link_meta->>'tags'));

CREATE OR REPLACE TRIGGER link_default_code
  BEFORE INSERT ON link
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('LNK');

CREATE TRIGGER update_link_changed BEFORE INSERT OR UPDATE ON link
  FOR EACH ROW EXECUTE FUNCTION link_changed();

CREATE OR REPLACE VIEW link_view AS
  SELECT id, code, 
    link_type_1, link_code_1, link_type_2, link_code_2,
    cast(link_meta->>'qty' AS FLOAT) AS qty, cast(link_meta->>'amount' AS FLOAT) AS amount, cast(link_meta->>'rate' AS FLOAT) AS rate,
    COALESCE(link_meta->>'notes', '') AS notes,
    link_meta->'tags' AS tags, REGEXP_REPLACE(link_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    link_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'link_type_1', link_type_1, 'link_code_1', link_code_1,
      'link_type_2', link_type_2, 'link_code_2', link_code_2, 'link_meta', link_meta,
      'link_map', link_map, 'time_stamp', time_stamp
    ) AS link_object
  FROM link
  WHERE deleted = false;

CREATE OR REPLACE VIEW link_map AS
  SELECT link.id AS id, link.code, 
    link_type_1, link_code_1, link_type_2, link_code_2,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM link, jsonb_each(link.link_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW link_tags AS
  SELECT link.id AS id, code, value as tag
  FROM link, jsonb_array_elements_text(link.link_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS product(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  product_type product_type NOT NULL DEFAULT 'PRODUCT_ITEM'::product_type, 
  product_name VARCHAR NOT NULL,
  tax_code VARCHAR NOT NULL,
  events JSONB NOT NULL DEFAULT '[]'::JSONB,
  product_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  product_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT product_pkey PRIMARY KEY (id),
  CONSTRAINT product_code_key UNIQUE (code),
  FOREIGN KEY (tax_code) REFERENCES tax(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_product_product_type ON product (product_type);
CREATE INDEX idx_product_tax_code ON product (tax_code);
CREATE INDEX idx_product_deleted ON product (deleted);
CREATE INDEX idx_product_name ON product (product_name);
CREATE INDEX idx_product_barcode ON product ((product_meta->>'barcode'));
CREATE INDEX idx_product_inactive ON product ((product_meta->>'inactive'));
CREATE INDEX idx_product_tags ON product ((product_meta->>'tags'));

CREATE OR REPLACE TRIGGER product_default_code
  BEFORE INSERT ON product
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('PRD');

CREATE OR REPLACE VIEW product_view AS
  SELECT id, code, product_type, tax_code, product_name,
    product_meta->>'unit' AS unit,
    product_meta->>'barcode_type' AS barcode_type, product_meta->>'barcode' AS barcode, 
    cast(product_meta->>'barcode_qty' AS FLOAT) AS barcode_qty,
    cast(product_meta->>'inactive' AS BOOLEAN) AS inactive, product_meta->>'notes' AS notes, 
    product_meta->'tags' AS tags, REGEXP_REPLACE(product_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    events, product_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'product_type', product_type, 'tax_code', tax_code,
      'product_name', product_name, 'product_meta', product_meta, 'events', events,
      'product_map', product_map, 'time_stamp', time_stamp
    ) AS product_object
  FROM product
  WHERE deleted = false;

CREATE OR REPLACE VIEW product_events AS
  SELECT product.id AS id, code, product_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, 
    CASE WHEN value->>'start_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'start_time', 'YYYY-MM-DDTHH24:MI:SS') END AS start_time,
	  CASE WHEN value->>'end_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'end_time', 'YYYY-MM-DDTHH24:MI:SS') END AS end_time,
    value->>'place' AS place, value->>'description' AS description, value->'tags' AS tags,
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'event_map' AS event_map
  FROM product, jsonb_array_elements(product.events)
  WHERE deleted = false;

CREATE OR REPLACE VIEW product_map AS
  SELECT product.id AS id, product.code, product.product_name,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM product, jsonb_each(product.product_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW product_tags AS
  SELECT product.id AS id, code, value as tag
  FROM product, jsonb_array_elements_text(product.product_meta->'tags')
  WHERE deleted = false;

CREATE OR REPLACE VIEW product_components AS
  SELECT p.id, p.code AS product_code, p.product_name, COALESCE(p.product_meta->>'unit','') as unit, 
  c.code as ref_product_code, c.product_name as component_name, COALESCE(c.product_meta->>'unit','') as component_unit, 
  c.product_type as component_type,
  CAST(l.link_meta->>'qty' AS FLOAT) AS qty, COALESCE(l.link_meta->>'notes', '') AS notes
  FROM product p INNER JOIN link l ON l.link_code_1 = p.code
  INNER JOIN product c ON l.link_code_2 = c.code
  WHERE p.product_type = 'PRODUCT_VIRTUAL' AND link_type_1 = 'LINK_PRODUCT' AND link_type_2 = 'LINK_PRODUCT' 
  AND p.deleted = false AND l.deleted = false AND c.deleted = false;

CREATE TABLE IF NOT EXISTS project(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  project_name VARCHAR NOT NULL,
  customer_code VARCHAR,
  addresses JSONB NOT NULL DEFAULT '[]'::JSONB,
  contacts JSONB NOT NULL DEFAULT '[]'::JSONB,
  events JSONB NOT NULL DEFAULT '[]'::JSONB,
  project_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  project_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT project_pkey PRIMARY KEY (id),
  CONSTRAINT project_code_key UNIQUE (code),
  FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_project_customer_code ON project (customer_code);
CREATE INDEX idx_project_deleted ON project (deleted);
CREATE INDEX idx_project_project_name ON project (project_name);
CREATE INDEX idx_project_inactive ON project ((project_meta->>'inactive'));
CREATE INDEX idx_project_tags ON project ((project_meta->>'tags'));

CREATE OR REPLACE TRIGGER project_default_code
  BEFORE INSERT ON project
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('PRJ');

CREATE OR REPLACE VIEW project_view AS
  SELECT id, code, customer_code, project_name,
	CASE WHEN project_meta->>'start_date' = '' THEN NULL ELSE TO_DATE(project_meta->>'start_date', 'YYYY-MM-DD') END AS start_date, 
	CASE WHEN project_meta->>'end_date' = '' THEN NULL ELSE TO_DATE(project_meta->>'end_date', 'YYYY-MM-DD') END AS end_date, 
    cast(project_meta->>'inactive' AS BOOLEAN) AS inactive, project_meta->>'notes' AS notes, 
    project_meta->'tags' AS tags, REGEXP_REPLACE(project_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    addresses, contacts, events, project_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'customer_code', customer_code, 'project_name', project_name,
      'project_meta', project_meta, 'addresses', addresses, 'contacts', contacts,
      'events', events, 'project_map', project_map, 'time_stamp', time_stamp
    ) AS project_object
  FROM project
  WHERE deleted = false;

CREATE OR REPLACE VIEW project_contacts AS
  SELECT project.id AS id, code, project_name, 
    value->>'first_name' AS first_name, value->>'surname' AS surname, value->>'status' AS status,
    value->>'phone' AS phone, value->>'mobile' AS mobile, value->>'email' AS email, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'contact_map' AS contact_map
  FROM project, jsonb_array_elements(project.contacts)
  WHERE deleted = false;

CREATE OR REPLACE VIEW project_addresses AS
  SELECT project.id AS id, code, project_name, 
    value->>'country' AS country, value->>'state' AS state, value->>'zip_code' AS zip_code,
    value->>'city' AS city, value->>'street' AS street, 
    value->>'notes' AS notes, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'address_map' AS address_map
  FROM project, jsonb_array_elements(project.addresses)
  WHERE deleted = false;

CREATE OR REPLACE VIEW project_events AS
  SELECT project.id AS id, code, project_name, 
    value->>'uid' AS uid, value->>'subject' AS subject, 
    CASE WHEN value->>'start_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'start_time', 'YYYY-MM-DDTHH24:MI:SS') END AS start_time,
	  CASE WHEN value->>'end_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'end_time', 'YYYY-MM-DDTHH24:MI:SS') END AS end_time,
    value->>'place' AS place, value->>'description' AS description, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'event_map' AS event_map
  FROM project, jsonb_array_elements(project.events)
  WHERE deleted = false;

CREATE OR REPLACE VIEW project_map AS
  SELECT project.id AS id, project.code, project.project_name,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM project, jsonb_each(project.project_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW project_tags AS
  SELECT project.id AS id, code, value as tag
  FROM project, jsonb_array_elements_text(project.project_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS rate(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  rate_type rate_type NOT NULL DEFAULT 'RATE_RATE'::rate_type,
  rate_date DATE NOT NULL DEFAULT CURRENT_DATE,
  place_code VARCHAR,
  currency_code VARCHAR NOT NULL, 
  rate_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  rate_map JSONB NOT NULL DEFAULT '{}'::JSONB,
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
CREATE INDEX idx_rate_tags ON rate ((rate_meta->>'tags'));

CREATE OR REPLACE TRIGGER rate_default_code
  BEFORE INSERT ON rate
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('RAT');

CREATE OR REPLACE VIEW rate_view AS
  SELECT id, code, rate_type, rate_date, place_code, currency_code,
    cast(rate_meta->>'rate_value' AS FLOAT) AS rate_value,
    rate_meta->'tags' AS tags, REGEXP_REPLACE(rate_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    rate_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'rate_type', rate_type, 'rate_date', rate_date,
      'place_code', place_code, 'currency_code', currency_code, 'rate_meta', rate_meta,
      'rate_map', rate_map, 'time_stamp', time_stamp
    ) AS rate_object
  FROM rate
  WHERE deleted = false;

CREATE OR REPLACE VIEW rate_map AS
  SELECT rate.id AS id, rate.code,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM rate, jsonb_each(rate.rate_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW rate_tags AS
  SELECT rate.id AS id, code, value as tag
  FROM rate, jsonb_array_elements_text(rate.rate_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS tool(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  description VARCHAR NOT NULL,
  product_code VARCHAR NOT NULL,
  events JSONB NOT NULL DEFAULT '[]'::JSONB,
  tool_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  tool_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT tool_pkey PRIMARY KEY (id),
  CONSTRAINT tool_code_key UNIQUE (code),
  FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_tool_product_code ON tool (product_code);
CREATE INDEX idx_tool_deleted ON tool (deleted);
CREATE INDEX idx_tool_description ON tool (description);
CREATE INDEX idx_tool_inactive ON tool ((tool_meta->>'inactive'));
CREATE INDEX idx_tool_tags ON tool ((tool_meta->>'tags'));

CREATE OR REPLACE TRIGGER tool_default_code
  BEFORE INSERT ON tool
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('SER');

CREATE OR REPLACE VIEW tool_view AS
  SELECT id, code, product_code, description,
    tool_meta->>'serial_number' AS serial_number,
    cast(tool_meta->>'inactive' AS BOOLEAN) AS inactive, 
    tool_meta->>'notes' AS notes, tool_meta->'tags' AS tags, 
    REGEXP_REPLACE(tool_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    tool_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'product_code', product_code, 'description', description,
      'tool_meta', tool_meta, 'tool_map', tool_map, 'time_stamp', time_stamp
    ) AS tool_object
  FROM tool
  WHERE deleted = false;

CREATE OR REPLACE VIEW tool_events AS
  SELECT tool.id AS id, code, description as tool_description, 
    value->>'uid' AS uid, value->>'subject' AS subject, 
    CASE WHEN value->>'start_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'start_time', 'YYYY-MM-DDTHH24:MI:SS') END AS start_time,
	  CASE WHEN value->>'end_time' = '' THEN NULL ELSE TO_TIMESTAMP(value->>'end_time', 'YYYY-MM-DDTHH24:MI:SS') END AS end_time,
    value->>'place' AS place, value->>'description' AS description, value->'tags' AS tags, 
    REGEXP_REPLACE(value->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    value->'event_map' AS event_map
  FROM tool, jsonb_array_elements(tool.events)
  WHERE deleted = false;

CREATE OR REPLACE VIEW tool_map AS
  SELECT tool.id AS id, tool.code, tool.description as tool_description,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM tool, jsonb_each(tool.tool_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW tool_tags AS
  SELECT tool.id AS id, code, value as tag
  FROM tool, jsonb_array_elements_text(tool.tool_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS price(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  price_type price_type NOT NULL DEFAULT 'PRICE_CUSTOMER'::price_type,
  valid_from DATE NOT NULL,
  valid_to DATE,
  product_code VARCHAR NOT NULL,
  currency_code VARCHAR NOT NULL,
  customer_code VARCHAR,
  qty INTEGER NOT NULL DEFAULT 0,
  price_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  price_map JSONB NOT NULL DEFAULT '{}'::JSONB,
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
CREATE INDEX idx_price_tags ON price ((price_meta->>'tags'));

CREATE OR REPLACE TRIGGER price_default_code
  BEFORE INSERT ON price
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('PRC');

CREATE OR REPLACE VIEW price_view AS
  SELECT id, code, price_type, valid_from, valid_to, product_code, currency_code, customer_code, qty,
    cast(price_meta->>'price_value' AS FLOAT) AS price_value,
    price_meta->'tags' AS tags, REGEXP_REPLACE(price_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    price_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'price_type', price_type, 'valid_from', valid_from,
      'valid_to', valid_to, 'product_code', product_code, 'currency_code', currency_code,
      'customer_code', customer_code, 'qty', qty, 'price_meta', price_meta, 'price_map', price_map,
      'time_stamp', time_stamp
    ) AS price_object
  FROM price
  WHERE deleted = false;

CREATE OR REPLACE VIEW price_map AS
  SELECT price.id AS id, price.code, 
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM price, jsonb_each(price.price_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW price_tags AS
  SELECT price.id AS id, code, value as tag
  FROM price, jsonb_array_elements_text(price.price_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS trans(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  trans_type trans_type NOT NULL DEFAULT 'TRANS_INVOICE'::trans_type,
  direction direction NOT NULL DEFAULT 'DIRECTION_OUT'::direction,
  trans_date DATE NOT NULL,
  trans_code VARCHAR,
  customer_code VARCHAR,
  employee_code VARCHAR,
  project_code VARCHAR,
  place_code VARCHAR,
  currency_code VARCHAR,
  auth_code VARCHAR NOT NULL, 
  trans_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  trans_map JSONB NOT NULL DEFAULT '{}'::JSONB,
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
CREATE INDEX idx_trans_due_time ON trans ((trans_meta->>'due_time'));
CREATE INDEX idx_trans_paid_type ON trans ((trans_meta->>'paid_type'));
CREATE INDEX idx_trans_trans_state ON trans ((trans_meta->>'trans_state'));
CREATE INDEX idx_trans_tags ON trans ((trans_meta->>'tags'));

CREATE OR REPLACE TRIGGER trans_default_code
  BEFORE INSERT ON trans
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_trans_code();

CREATE OR REPLACE TRIGGER trans_invoice_customer_insert
  BEFORE INSERT ON trans
  FOR EACH ROW
  WHEN ((NEW.trans_type = 'TRANS_INVOICE'::trans_type) AND (NEW.direction = 'DIRECTION_OUT'::direction))
  EXECUTE FUNCTION set_invoice_customer();

CREATE OR REPLACE TRIGGER trans_invoice_customer_update
  BEFORE UPDATE ON trans
  FOR EACH ROW
  WHEN ((NEW.trans_type = 'TRANS_INVOICE'::trans_type) AND (NEW.direction = 'DIRECTION_OUT'::direction) AND (NEW.customer_code <> OLD.customer_code))
  EXECUTE FUNCTION set_invoice_customer();

CREATE OR REPLACE VIEW trans_view AS
  SELECT id, code, trans_type, direction, trans_date, trans_code, customer_code, employee_code,
    project_code, place_code, currency_code, auth_code,
    CASE WHEN trans_meta->>'due_time' = '' THEN NULL ELSE TO_TIMESTAMP(trans_meta->>'due_time', 'YYYY-MM-DDTHH24:MI:SS') END AS due_time, 
    trans_meta->>'ref_number' AS ref_number,
    trans_meta->>'paid_type' AS paid_type, 
    cast(trans_meta->>'tax_free' AS BOOLEAN) AS tax_free, cast(trans_meta->>'paid' AS BOOLEAN) AS paid,
    cast(trans_meta->>'rate' AS FLOAT) AS rate,
    cast(trans_meta->>'closed' AS BOOLEAN) AS closed,
    trans_meta->>'status' AS status,
    trans_meta->>'trans_state' AS trans_state,
    trans_meta->>'notes' AS notes, trans_meta->>'internal_notes' AS internal_notes,
    trans_meta->>'report_notes' AS report_notes,
    CASE WHEN jsonb_typeof(trans_meta->'worksheet') = 'object' THEN cast(trans_meta->'worksheet'->>'distance' as FLOAT) ELSE 0 END AS worksheet_distance,
    CASE WHEN jsonb_typeof(trans_meta->'worksheet') = 'object' THEN cast(trans_meta->'worksheet'->>'repair' as FLOAT) ELSE 0 END AS worksheet_repair,
    CASE WHEN jsonb_typeof(trans_meta->'worksheet') = 'object' THEN cast(trans_meta->'worksheet'->>'total' as FLOAT) ELSE 0 END AS worksheet_total,
    CASE WHEN jsonb_typeof(trans_meta->'worksheet') = 'object' THEN trans_meta->'worksheet'->>'justification' ELSE '' END AS worksheet_justification,
    CASE WHEN jsonb_typeof(trans_meta->'rent') = 'object' THEN cast(trans_meta->'rent'->>'holiday' as FLOAT) ELSE 0 END AS rent_holiday,
    CASE WHEN jsonb_typeof(trans_meta->'rent') = 'object' THEN cast(trans_meta->'rent'->>'bad_tool' as FLOAT) ELSE 0 END AS rent_bad_tool,
    CASE WHEN jsonb_typeof(trans_meta->'rent') = 'object' THEN cast(trans_meta->'rent'->>'other' as FLOAT) ELSE 0 END AS rent_other,
    CASE WHEN jsonb_typeof(trans_meta->'rent') = 'object' THEN trans_meta->'rent'->>'justification' ELSE '' END AS rent_justification,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'company_name' ELSE '' END AS invoice_company_name,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'company_account' ELSE '' END AS invoice_company_account,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'company_address' ELSE '' END AS invoice_company_address,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'company_tax_number' ELSE '' END AS invoice_company_tax_number,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'customer_name' ELSE '' END AS invoice_customer_name,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'customer_account' ELSE '' END AS invoice_customer_account,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'customer_address' ELSE '' END AS invoice_customer_address,
    CASE WHEN jsonb_typeof(trans_meta->'invoice') = 'object' THEN trans_meta->'invoice'->>'customer_tax_number' ELSE '' END AS invoice_customer_tax_number,
    trans_meta->'tags' AS tags, REGEXP_REPLACE(trans_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    trans_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'trans_type', trans_type, 'direction', direction,
      'trans_date', trans_date, 'trans_code', trans_code, 'customer_code', customer_code,
      'employee_code', employee_code, 'project_code', project_code, 'place_code', place_code,
      'currency_code', currency_code, 'auth_code', auth_code, 'trans_meta', trans_meta,
      'trans_map', trans_map, 'time_stamp', time_stamp
    ) AS trans_object
  FROM trans
  WHERE deleted = false;

CREATE OR REPLACE VIEW trans_map AS
  SELECT trans.id AS id, trans.code, trans.trans_type, trans.direction, trans.trans_date,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM trans, jsonb_each(trans.trans_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false;

CREATE OR REPLACE VIEW trans_tags AS
  SELECT trans.id AS id, code, value as tag
  FROM trans, jsonb_array_elements_text(trans.trans_meta->'tags')
  WHERE deleted = false;

CREATE TABLE IF NOT EXISTS item(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  trans_code VARCHAR NOT NULL,
  product_code VARCHAR NOT NULL,
  tax_code VARCHAR NOT NULL, 
  item_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  item_map JSONB NOT NULL DEFAULT '{}'::JSONB,
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
CREATE INDEX idx_item_tags ON item ((item_meta->>'tags'));

CREATE OR REPLACE TRIGGER item_default_code
  BEFORE INSERT ON item
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('ITM');

CREATE OR REPLACE VIEW item_view AS
  SELECT id, code, trans_code, product_code, tax_code,
    item_meta->>'unit' AS unit, cast(item_meta->>'qty' AS FLOAT) AS qty, cast(item_meta->>'fx_price' AS FLOAT) AS fx_price,
    cast(item_meta->>'net_amount' AS FLOAT) AS net_amount, cast(item_meta->>'qtdiscounty' AS FLOAT) AS discount, 
    cast(item_meta->>'vat_amount' AS FLOAT) AS vat_amount, cast(item_meta->>'amount' AS FLOAT) AS amount,
    item_meta->>'description' AS description, cast(item_meta->>'deposit' AS BOOLEAN) AS deposit,
    cast(item_meta->>'action_price' AS BOOLEAN) AS action_price, cast(item_meta->>'own_stock' AS FLOAT) AS own_stock, 
    item_meta->'tags' AS tags, REGEXP_REPLACE(item_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    item_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'trans_code', trans_code, 'product_code', product_code,
      'tax_code', tax_code, 'item_meta', item_meta, 'item_map', item_map, 'time_stamp', time_stamp
    ) AS item_object
  FROM item
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW item_map AS
  SELECT item.id AS id, item.code, item.trans_code, item.product_code,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM item, jsonb_each(item.item_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW item_tags AS
  SELECT item.id AS id, code, value as tag
  FROM item, jsonb_array_elements_text(item.item_meta->'tags')
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE TABLE IF NOT EXISTS movement(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  movement_type movement_type NOT NULL DEFAULT 'MOVEMENT_INVENTORY'::movement_type,
  shipping_time TIMESTAMP NOT NULL,
  trans_code VARCHAR NOT NULL,
  product_code VARCHAR,
  tool_code VARCHAR,
  place_code VARCHAR,
  item_code VARCHAR,
  movement_code VARCHAR,
  movement_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  movement_map JSONB NOT NULL DEFAULT '{}'::JSONB,
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
CREATE INDEX idx_movement_tags ON movement ((movement_meta->>'tags'));

CREATE OR REPLACE TRIGGER movement_default_code
  BEFORE INSERT ON movement
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('MOV');

CREATE OR REPLACE VIEW movement_view AS
  SELECT id, code, movement_type, shipping_time,
    trans_code, product_code, tool_code, place_code, item_code, movement_code,
    cast(movement_meta->>'qty' AS FLOAT) AS qty, cast(movement_meta->>'shared' AS BOOLEAN) AS shared,
    movement_meta->>'notes' AS notes,
    movement_meta->'tags' AS tags, REGEXP_REPLACE(movement_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    movement_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'movement_type', movement_type, 'shipping_time', shipping_time,
      'trans_code', trans_code, 'product_code', product_code, 'tool_code', tool_code, 'place_code', place_code,
      'item_code', item_code, 'movement_code', movement_code,
      'movement_meta', movement_meta, 'movement_map', movement_map, 'time_stamp', time_stamp
    ) AS movement_object
  FROM movement
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW movement_map AS
  SELECT movement.id AS id, movement.code, movement.trans_code,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM movement, jsonb_each(movement.movement_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW movement_tags AS
  SELECT movement.id AS id, code, value as tag
  FROM movement, jsonb_array_elements_text(movement.movement_meta->'tags')
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW movement_stock AS
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

CREATE OR REPLACE VIEW movement_inventory AS
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

CREATE OR REPLACE VIEW movement_waybill AS
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

CREATE OR REPLACE VIEW movement_formula AS
  SELECT mv.id, mv.code, t.code AS trans_code, CASE WHEN mv.movement_type = 'MOVEMENT_HEAD' THEN 'IN' ELSE 'OUT' END as direction, 
    mv.product_code, p.product_name, p.product_meta->>'unit' as unit,
    CAST(mv.movement_meta->>'qty' AS FLOAT) AS qty, mv.movement_meta->>'notes' as batch_no,
    mv.place_code, pl.place_name, CAST(mv.movement_meta->>'shared' AS BOOLEAN) AS shared
  FROM trans t INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN product p ON mv.product_code = p.code
  LEFT JOIN place pl ON mv.place_code = pl.code
  WHERE t.trans_type = 'TRANS_FORMULA' and mv.deleted = false and t.deleted = false;

CREATE OR REPLACE VIEW item_shipping AS
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
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  paid_date DATE NOT NULL,
  trans_code VARCHAR NOT NULL, 
  payment_meta JSONB NOT NULL DEFAULT '{}'::JSONB,
  payment_map JSONB NOT NULL DEFAULT '{}'::JSONB,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT payment_pkey PRIMARY KEY (id),
  CONSTRAINT payment_code_key UNIQUE (code),
  FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE INDEX idx_payment_trans_code ON payment (trans_code);
CREATE INDEX idx_payment_deleted ON payment (deleted);
CREATE INDEX idx_payment_tags ON payment ((payment_meta->>'tags'));

CREATE OR REPLACE TRIGGER payment_default_code
  BEFORE INSERT ON payment
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('PMT');

CREATE OR REPLACE VIEW payment_view AS
  SELECT id, code, paid_date, trans_code,
    cast(payment_meta->>'amount' AS FLOAT) AS amount, payment_meta->>'notes' AS notes,
    payment_meta->'tags' AS tags, REGEXP_REPLACE(payment_meta->>'tags', '[\[\]"]', '', 'g') as tag_lst, 
    payment_map, time_stamp,
    jsonb_build_object(
      'id', id, 'code', code, 'paid_date', paid_date, 'trans_code', trans_code,
      'payment_meta', payment_meta, 'payment_map', payment_map, 'time_stamp', time_stamp
    ) AS payment_object
  FROM payment
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW payment_map AS
  SELECT payment.id AS id, payment.code, payment.paid_date, payment.trans_code,
    key as map_key, value as map_value, jsonb_typeof(value) as map_type,
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
  FROM payment, jsonb_each(payment.payment_map) LEFT JOIN config_map cf on key = cf.field_name
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE OR REPLACE VIEW payment_invoice AS
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

CREATE OR REPLACE VIEW payment_tags AS
  SELECT payment.id AS id, code, value as tag
  FROM payment, jsonb_array_elements_text(payment.payment_meta->'tags')
  WHERE deleted = false AND trans_code IN (SELECT code FROM trans WHERE deleted = false);

CREATE TABLE IF NOT EXISTS log(
  id SERIAL NOT NULL,
  code VARCHAR NOT NULL,
  log_type log_type NOT NULL DEFAULT 'LOG_UPDATE'::log_type,
  ref_type VARCHAR NOT NULL DEFAULT 'TRANS'::VARCHAR,
  ref_code VARCHAR NOT NULL,
  auth_code VARCHAR NOT NULL,
  data JSONB NOT NULL DEFAULT '{}'::JSONB,
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

CREATE OR REPLACE TRIGGER log_default_code
  BEFORE INSERT ON log
  FOR EACH ROW
  WHEN (NEW.code IS NULL)
  EXECUTE FUNCTION set_new_code('LOG');
