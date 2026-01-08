SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;

/* =========================
   Tables
   ========================= */

CREATE TABLE usref(
  id INT IDENTITY(1,1) NOT NULL,
  refnumber NVARCHAR(255) NOT NULL,
  value NVARCHAR(255) NOT NULL,
  changed DATETIME2(0) NOT NULL CONSTRAINT DF_usref_changed DEFAULT (SYSUTCDATETIME()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_usref_time_stamp DEFAULT (SYSUTCDATETIME()),
  CONSTRAINT PK_usref PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_usref_refnumber UNIQUE (refnumber)
);

CREATE TABLE config(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  config_type NVARCHAR(32) NOT NULL CONSTRAINT DF_config_config_type DEFAULT (N'CONFIG_MAP')
    CONSTRAINT CK_config_config_type CHECK (config_type IN (
      N'CONFIG_MAP', N'CONFIG_SHORTCUT', N'CONFIG_MESSAGE', N'CONFIG_PATTERN', N'CONFIG_REPORT', N'CONFIG_PRINT_QUEUE', N'CONFIG_DATA'
    )),
  data JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_config_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_config_deleted DEFAULT (0),
  CONSTRAINT PK_config PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_config_code UNIQUE (code)
);

CREATE TABLE auth(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  user_name NVARCHAR(255) NOT NULL,
  user_group NVARCHAR(32) NOT NULL CONSTRAINT DF_auth_user_group DEFAULT (N'GROUP_USER')
    CONSTRAINT CK_auth_user_group CHECK (user_group IN (N'GROUP_ADMIN', N'GROUP_USER', N'GROUP_GUEST')),
  disabled BIT NOT NULL CONSTRAINT DF_auth_disabled DEFAULT (0),
  auth_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  auth_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_auth_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_auth_deleted DEFAULT (0),
  CONSTRAINT PK_auth PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_auth_code UNIQUE (code)
);

CREATE TABLE currency(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NOT NULL,
  currency_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  currency_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_currency_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_currency_deleted DEFAULT (0),
  CONSTRAINT PK_currency PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_currency_code UNIQUE (code)
);

CREATE TABLE customer(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  customer_type NVARCHAR(32) NOT NULL CONSTRAINT DF_customer_type DEFAULT (N'CUSTOMER_COMPANY')
    CONSTRAINT CK_customer_type CHECK (customer_type IN (N'CUSTOMER_OWN', N'CUSTOMER_COMPANY', N'CUSTOMER_PRIVATE', N'CUSTOMER_OTHER')),
  customer_name NVARCHAR(255) NOT NULL,
  addresses JSON NOT NULL DEFAULT (JSON_ARRAY()),
  contacts JSON NOT NULL DEFAULT (JSON_ARRAY()),
  events JSON NOT NULL DEFAULT (JSON_ARRAY()),
  customer_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  customer_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_customer_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_customer_deleted DEFAULT (0),
  CONSTRAINT PK_customer PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_customer_code UNIQUE (code)
);

CREATE TABLE employee(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  address JSON NOT NULL DEFAULT (JSON_OBJECT()),
    CONSTRAINT CK_employee_address_json CHECK (ISJSON(address) = 1),
  contact JSON NOT NULL DEFAULT (JSON_OBJECT()),
  events JSON NOT NULL DEFAULT (JSON_ARRAY()),
  employee_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  employee_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_employee_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_employee_deleted DEFAULT (0),
  CONSTRAINT PK_employee PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_employee_code UNIQUE (code)
);

CREATE TABLE place(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  place_type NVARCHAR(32) NOT NULL CONSTRAINT DF_place_type DEFAULT (N'PLACE_WAREHOUSE')
    CONSTRAINT CK_place_type CHECK (place_type IN (N'PLACE_BANK', N'PLACE_CASH', N'PLACE_WAREHOUSE', N'PLACE_OTHER')),
  place_name NVARCHAR(255) NOT NULL,
  currency_code NVARCHAR(255) NULL,
  address JSON NOT NULL DEFAULT (JSON_OBJECT()),
  contacts JSON NOT NULL DEFAULT (JSON_ARRAY()),
  events JSON NOT NULL DEFAULT (JSON_ARRAY()),
  place_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  place_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_place_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_place_deleted DEFAULT (0),
  CONSTRAINT PK_place PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_place_code UNIQUE (code),
  CONSTRAINT FK_place_currency_code FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE tax(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NOT NULL,
  tax_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  tax_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_tax_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_tax_deleted DEFAULT (0),
  CONSTRAINT PK_tax PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_tax_code UNIQUE (code)
);

CREATE TABLE link(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  link_type_1 NVARCHAR(32) NOT NULL CONSTRAINT DF_link_type_1 DEFAULT (N'LINK_TRANS')
    CONSTRAINT CK_link_type_1 CHECK (link_type_1 IN (
      N'LINK_CUSTOMER', N'LINK_EMPLOYEE', N'LINK_ITEM', N'LINK_MOVEMENT', N'LINK_PAYMENT', N'LINK_PLACE',
      N'LINK_PRODUCT', N'LINK_PROJECT', N'LINK_TOOL', N'LINK_TRANS'
    )),
  link_code_1 NVARCHAR(255) NOT NULL,
  link_type_2 NVARCHAR(32) NOT NULL CONSTRAINT DF_link_type_2 DEFAULT (N'LINK_TRANS')
    CONSTRAINT CK_link_type_2 CHECK (link_type_2 IN (
      N'LINK_CUSTOMER', N'LINK_EMPLOYEE', N'LINK_ITEM', N'LINK_MOVEMENT', N'LINK_PAYMENT', N'LINK_PLACE',
      N'LINK_PRODUCT', N'LINK_PROJECT', N'LINK_TOOL', N'LINK_TRANS'
    )),
  link_code_2 NVARCHAR(255) NOT NULL,
  link_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  link_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_link_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_link_deleted DEFAULT (0),
  CONSTRAINT PK_link PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_link_code UNIQUE (code)
);

CREATE TABLE product(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  product_type NVARCHAR(32) NOT NULL CONSTRAINT DF_product_type DEFAULT (N'PRODUCT_ITEM')
    CONSTRAINT CK_product_type CHECK (product_type IN (N'PRODUCT_ITEM', N'PRODUCT_SERVICE', N'PRODUCT_VIRTUAL')),
  product_name NVARCHAR(255) NOT NULL,
  tax_code NVARCHAR(255) NOT NULL,
  events JSON NOT NULL DEFAULT (JSON_ARRAY()),
  product_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  product_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_product_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_product_deleted DEFAULT (0),
  CONSTRAINT PK_product PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_product_code UNIQUE (code),
  CONSTRAINT FK_product_tax_code FOREIGN KEY (tax_code) REFERENCES tax(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE project(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  project_name NVARCHAR(255) NOT NULL,
  customer_code NVARCHAR(255) NULL,
  addresses JSON NOT NULL DEFAULT (JSON_ARRAY()),
  contacts JSON NOT NULL DEFAULT (JSON_ARRAY()),
  events JSON NOT NULL DEFAULT (JSON_ARRAY()),
  project_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  project_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_project_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_project_deleted DEFAULT (0),
  CONSTRAINT PK_project PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_project_code UNIQUE (code),
  CONSTRAINT FK_project_customer_code FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE rate(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  rate_type NVARCHAR(32) NOT NULL CONSTRAINT DF_rate_type DEFAULT (N'RATE_RATE')
    CONSTRAINT CK_rate_type CHECK (rate_type IN (N'RATE_RATE', N'RATE_BUY', N'RATE_SELL', N'RATE_AVERAGE')),
  rate_date DATE NOT NULL CONSTRAINT DF_rate_date DEFAULT (CONVERT(date, SYSUTCDATETIME())),
  place_code NVARCHAR(255) NULL,
  currency_code NVARCHAR(255) NOT NULL,
  rate_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  rate_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_rate_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_rate_deleted DEFAULT (0),
  CONSTRAINT PK_rate PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_rate_code UNIQUE (code),
  CONSTRAINT FK_rate_place_code FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_rate_currency_code FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE tool(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  description NVARCHAR(255) NOT NULL,
  product_code NVARCHAR(255) NOT NULL,
  events JSON NOT NULL DEFAULT (JSON_ARRAY()),
  tool_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  tool_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_tool_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_tool_deleted DEFAULT (0),
  CONSTRAINT PK_tool PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_tool_code UNIQUE (code),
  CONSTRAINT FK_tool_product_code FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE price(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  price_type NVARCHAR(32) NOT NULL CONSTRAINT DF_price_type DEFAULT (N'PRICE_CUSTOMER')
    CONSTRAINT CK_price_type CHECK (price_type IN (N'PRICE_CUSTOMER', N'PRICE_VENDOR')),
  valid_from DATE NOT NULL,
  valid_to DATE NULL,
  product_code NVARCHAR(255) NOT NULL,
  currency_code NVARCHAR(255) NOT NULL,
  customer_code NVARCHAR(255) NULL,
  qty INT NOT NULL CONSTRAINT DF_price_qty DEFAULT (0),
  price_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  price_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_price_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_price_deleted DEFAULT (0),
  CONSTRAINT PK_price PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_price_code UNIQUE (code),
  CONSTRAINT FK_price_product_code FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_price_currency_code FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_price_customer_code FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE trans(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  trans_type NVARCHAR(32) NOT NULL
    CONSTRAINT CK_trans_type CHECK (trans_type IN (
      N'TRANS_INVOICE', N'TRANS_RECEIPT', N'TRANS_ORDER', N'TRANS_OFFER', N'TRANS_WORKSHEET', N'TRANS_RENT', N'TRANS_DELIVERY',
      N'TRANS_INVENTORY', N'TRANS_WAYBILL', N'TRANS_PRODUCTION', N'TRANS_FORMULA', N'TRANS_BANK', N'TRANS_CASH'
    )),
  direction NVARCHAR(32) NOT NULL CONSTRAINT DF_trans_direction DEFAULT (N'DIRECTION_OUT')
    CONSTRAINT CK_trans_direction CHECK (direction IN (N'DIRECTION_OUT', N'DIRECTION_IN', N'DIRECTION_TRANSFER')),
  trans_date DATE NOT NULL,
  trans_code NVARCHAR(255) NULL,
  customer_code NVARCHAR(255) NULL,
  employee_code NVARCHAR(255) NULL,
  project_code NVARCHAR(255) NULL,
  place_code NVARCHAR(255) NULL,
  currency_code NVARCHAR(255) NULL,
  auth_code NVARCHAR(255) NOT NULL,
  trans_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  trans_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_trans_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_trans_deleted DEFAULT (0),
  CONSTRAINT PK_trans PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_trans_code UNIQUE (code),
  CONSTRAINT FK_trans_trans_code FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_trans_customer_code FOREIGN KEY (customer_code) REFERENCES customer(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_trans_employee_code FOREIGN KEY (employee_code) REFERENCES employee(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_trans_project_code FOREIGN KEY (project_code) REFERENCES project(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_trans_place_code FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_trans_currency_code FOREIGN KEY (currency_code) REFERENCES currency(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_trans_auth_code FOREIGN KEY (auth_code) REFERENCES auth(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE item(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  trans_code NVARCHAR(255) NOT NULL,
  product_code NVARCHAR(255) NOT NULL,
  tax_code NVARCHAR(255) NOT NULL,
  item_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  item_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_item_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_item_deleted DEFAULT (0),
  CONSTRAINT PK_item PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_item_code UNIQUE (code),
  CONSTRAINT FK_item_trans_code FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_item_product_code FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_item_tax_code FOREIGN KEY (tax_code) REFERENCES tax(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE movement(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  movement_type NVARCHAR(32) NOT NULL CONSTRAINT DF_movement_type DEFAULT (N'MOVEMENT_INVENTORY')
    CONSTRAINT CK_movement_type CHECK (movement_type IN (N'MOVEMENT_INVENTORY', N'MOVEMENT_TOOL', N'MOVEMENT_PLAN', N'MOVEMENT_HEAD')),
  shipping_time DATETIME2(0) NOT NULL,
  trans_code NVARCHAR(255) NOT NULL,
  product_code NVARCHAR(255) NULL,
  tool_code NVARCHAR(255) NULL,
  place_code NVARCHAR(255) NULL,
  item_code NVARCHAR(255) NULL,
  movement_code NVARCHAR(255) NULL,
  movement_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  movement_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_movement_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_movement_deleted DEFAULT (0),
  CONSTRAINT PK_movement PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_movement_code UNIQUE (code),
  CONSTRAINT FK_movement_trans_code FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_movement_product_code FOREIGN KEY (product_code) REFERENCES product(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_movement_tool_code FOREIGN KEY (tool_code) REFERENCES tool(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_movement_place_code FOREIGN KEY (place_code) REFERENCES place(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_movement_item_code FOREIGN KEY (item_code) REFERENCES item(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT FK_movement_movement_code FOREIGN KEY (movement_code) REFERENCES movement(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE payment(
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  paid_date DATE NOT NULL,
  trans_code NVARCHAR(255) NOT NULL,
  payment_meta JSON NOT NULL DEFAULT (JSON_OBJECT()),
  payment_map JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_payment_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_payment_deleted DEFAULT (0),
  CONSTRAINT PK_payment PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_payment_code UNIQUE (code),
  CONSTRAINT FK_payment_trans_code FOREIGN KEY (trans_code) REFERENCES trans(code)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE [log](
  id INT IDENTITY(1,1) NOT NULL,
  code NVARCHAR(255) NULL,
  log_type NVARCHAR(32) NOT NULL CONSTRAINT DF_log_type DEFAULT (N'LOG_UPDATE')
    CONSTRAINT CK_log_type CHECK (log_type IN (N'LOG_INSERT', N'LOG_UPDATE', N'LOG_DELETE')),
  ref_type NVARCHAR(255) NOT NULL CONSTRAINT DF_log_ref_type DEFAULT (N'TRANS'),
  ref_code NVARCHAR(255) NOT NULL,
  auth_code NVARCHAR(255) NOT NULL,
  data JSON NOT NULL DEFAULT (JSON_OBJECT()),
  time_stamp DATETIME2(0) NOT NULL CONSTRAINT DF_log_time_stamp DEFAULT (SYSUTCDATETIME()),
  deleted BIT NOT NULL CONSTRAINT DF_log_deleted DEFAULT (0),
  CONSTRAINT PK_log PRIMARY KEY CLUSTERED (id),
  CONSTRAINT UQ_log_code UNIQUE (code)
);

/* =========================
   Indexes
   ========================= */

CREATE INDEX idx_config_type ON config (config_type);
CREATE INDEX idx_config_deleted ON config (deleted);

CREATE UNIQUE INDEX idx_auth_user_name ON auth (user_name);
CREATE INDEX idx_auth_deleted ON auth (deleted);
CREATE INDEX idx_auth_disabled ON auth (disabled);

CREATE INDEX idx_currency_deleted ON currency (deleted);

CREATE INDEX idx_customer_customer_type ON customer (customer_type);
CREATE INDEX idx_customer_deleted ON customer (deleted);
CREATE INDEX idx_customer_customer_name ON customer (customer_name);

CREATE INDEX idx_employee_deleted ON employee (deleted);

CREATE INDEX idx_place_deleted ON place (deleted);
CREATE INDEX idx_place_place_name ON place (place_name);

CREATE INDEX idx_tax_deleted ON tax (deleted);

CREATE INDEX idx_link_link_code_1 ON link (link_type_1, link_code_1);
CREATE INDEX idx_link_link_code_2 ON link (link_type_2, link_code_2);
CREATE INDEX idx_link_deleted ON link (deleted);

CREATE INDEX idx_product_product_type ON product (product_type);
CREATE INDEX idx_product_product_name ON product (product_name);
CREATE INDEX idx_product_tax_code ON product (tax_code);
CREATE INDEX idx_product_deleted ON product (deleted);

CREATE INDEX idx_project_customer_code ON project (customer_code);
CREATE INDEX idx_project_project_name ON project (project_name);
CREATE INDEX idx_project_deleted ON project (deleted);

CREATE UNIQUE INDEX idx_rate_tdcp ON rate (rate_type, rate_date, currency_code, place_code);
CREATE INDEX idx_rate_rate_type ON rate (rate_type);
CREATE INDEX idx_rate_rate_date ON rate (rate_date);
CREATE INDEX idx_rate_place_code ON rate (place_code);
CREATE INDEX idx_rate_currency_code ON rate (currency_code);
CREATE INDEX idx_rate_deleted ON rate (deleted);

CREATE INDEX idx_tool_description ON tool (description);
CREATE INDEX idx_tool_product_code ON tool (product_code);
CREATE INDEX idx_tool_deleted ON tool (deleted);

CREATE UNIQUE INDEX idx_price_pvpc ON price (price_type, valid_from, product_code, currency_code, qty);
CREATE INDEX idx_price_price_type ON price (price_type);
CREATE INDEX idx_price_valid_from ON price (valid_from);
CREATE INDEX idx_price_valid_to ON price (valid_to);
CREATE INDEX idx_price_product_code ON price (product_code);
CREATE INDEX idx_price_currency_code ON price (currency_code);
CREATE INDEX idx_price_customer_code ON price (customer_code);
CREATE INDEX idx_price_deleted ON price (deleted);

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

CREATE INDEX idx_item_trans_code ON item (trans_code);
CREATE INDEX idx_item_product_code ON item (product_code);
CREATE INDEX idx_item_tax_code ON item (tax_code);
CREATE INDEX idx_item_deleted ON item (deleted);

CREATE INDEX idx_movement_trans_code ON movement (trans_code);
CREATE INDEX idx_movement_product_code ON movement (product_code);
CREATE INDEX idx_movement_tool_code ON movement (tool_code);
CREATE INDEX idx_movement_place_code ON movement (place_code);
CREATE INDEX idx_movement_item_code ON movement (item_code);
CREATE INDEX idx_movement_movement_code ON movement (movement_code);
CREATE INDEX idx_movement_deleted ON movement (deleted);

CREATE INDEX idx_payment_trans_code ON payment (trans_code);
CREATE INDEX idx_payment_deleted ON payment (deleted);

CREATE INDEX idx_log_ref_type ON [log] (ref_type);
CREATE INDEX idx_log_ref_code ON [log] (ref_code);
CREATE INDEX idx_log_auth_code ON [log] (auth_code);
CREATE INDEX idx_log_time_stamp ON [log] (time_stamp);
CREATE INDEX idx_log_deleted ON [log] (deleted);

/* =========================
   Views (T-SQL + JSON)
   ========================= */

CREATE OR ALTER VIEW config_map AS
  SELECT id, code,
    JSON_VALUE(data, N'$.field_name') AS field_name,
    JSON_VALUE(data, N'$.field_type') AS field_type,
    JSON_VALUE(data, N'$.description') AS description,
    JSON_QUERY(data, N'$.tags') AS tags,
    JSON_QUERY(data, N'$.filter') AS filter
  FROM config
  WHERE config_type = N'CONFIG_MAP' AND deleted = 0;

CREATE OR ALTER VIEW config_shortcut AS
  SELECT id, code,
    JSON_VALUE(data, N'$.shortcut_key') AS shortcut_key,
    JSON_VALUE(data, N'$.description') AS description,
    JSON_VALUE(data, N'$.modul') AS modul,
    JSON_VALUE(data, N'$.method') AS method,
    JSON_VALUE(data, N'$.func_name') AS func_name,
    JSON_VALUE(data, N'$.address') AS address,
    JSON_QUERY(data, N'$.fields') AS fields
  FROM config
  WHERE config_type = N'CONFIG_SHORTCUT' AND deleted = 0;

CREATE OR ALTER VIEW config_message AS
  SELECT id, code,
    JSON_VALUE(data, N'$.key') AS message_key,
    JSON_VALUE(data, N'$.lang') AS lang,
    JSON_VALUE(data, N'$.value') AS message_value
  FROM config
  WHERE config_type = N'CONFIG_MESSAGE' AND deleted = 0;

CREATE OR ALTER VIEW config_pattern AS
  SELECT id, code,
    JSON_VALUE(data, N'$.trans_type') AS trans_type,
    JSON_VALUE(data, N'$.description') AS description,
    JSON_VALUE(data, N'$.notes') AS notes,
    JSON_VALUE(data, N'$.default_pattern') AS default_pattern
  FROM config
  WHERE config_type = N'CONFIG_PATTERN' AND deleted = 0;

CREATE OR ALTER VIEW config_print_queue AS
  SELECT id, code,
    JSON_VALUE(data, N'$.ref_type') AS ref_type,
    JSON_VALUE(data, N'$.ref_code') AS ref_code,
    TRY_CONVERT(float, JSON_VALUE(data, N'$.qty')) AS qty,
    JSON_VALUE(data, N'$.report_code') AS report_code,
    JSON_VALUE(data, N'$.orientation') AS orientation,
    JSON_VALUE(data, N'$.paper_size') AS paper_size,
    JSON_VALUE(data, N'$.auth_code') AS auth_code,
    time_stamp
  FROM config
  WHERE config_type = N'CONFIG_PRINT_QUEUE' AND deleted = 0;

CREATE OR ALTER VIEW config_report AS
  SELECT id, code,
    JSON_VALUE(data, N'$.report_key') AS report_key,
    JSON_VALUE(data, N'$.report_type') AS report_type,
    JSON_VALUE(data, N'$.trans_type') AS trans_type,
    JSON_VALUE(data, N'$.direction') AS direction,
    JSON_VALUE(data, N'$.report_name') AS report_name,
    JSON_VALUE(data, N'$.description') AS description,
    JSON_VALUE(data, N'$.label') AS label,
    JSON_VALUE(data, N'$.file_type') AS file_type,
    JSON_QUERY(data, N'$.template') AS template
  FROM config
  WHERE config_type = N'CONFIG_REPORT' AND deleted = 0;

CREATE OR ALTER VIEW config_data AS
  SELECT
    ROW_NUMBER() OVER (ORDER BY c.code, j.[key]) AS id,
    c.code AS config_code,
    j.[key] AS config_key,
    j.value AS config_value,
    j.[type] AS config_type
  FROM config c
  CROSS APPLY OPENJSON(c.data) j
  WHERE c.config_type = N'CONFIG_DATA'
    AND c.deleted = 0
    AND ISJSON(c.data) = 1
    AND LEFT(LTRIM(JSON_VALUE(c.data, N'$')), 1) = N'{';

CREATE OR ALTER VIEW auth_view AS
  SELECT id, code, user_name, user_group, disabled,
    JSON_QUERY(auth_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(auth_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    auth_map, time_stamp,
    (SELECT
      id AS id, code AS code, user_name AS user_name, user_group AS user_group, disabled AS disabled,
      JSON_QUERY(auth_meta) AS auth_meta, JSON_QUERY(auth_map) AS auth_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS auth_object
  FROM auth
  WHERE deleted = 0;

CREATE OR ALTER VIEW auth_map AS
  SELECT tbl.id AS id, tbl.code, tbl.user_name, tbl.user_group,
    j.[key] AS map_key,
    j.value AS map_value,
    j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM auth tbl
  CROSS APPLY OPENJSON(tbl.auth_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW currency_view AS
  SELECT id, code,
    JSON_VALUE(currency_meta, N'$.description') AS description,
    TRY_CONVERT(int, JSON_VALUE(currency_meta, N'$.digit')) AS digit,
    TRY_CONVERT(int, JSON_VALUE(currency_meta, N'$.cash_round')) AS cash_round,
    JSON_QUERY(currency_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(currency_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    currency_map, time_stamp,
    (SELECT
      id AS id, code AS code, JSON_QUERY(currency_meta) AS currency_meta, JSON_QUERY(currency_map) AS currency_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS currency_object
  FROM currency
  WHERE deleted = 0;

CREATE OR ALTER VIEW currency_map AS
  SELECT tbl.id AS id, tbl.code,
    JSON_VALUE(currency_meta, N'$.description') AS currency_description,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM currency tbl
  CROSS APPLY OPENJSON(tbl.currency_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW currency_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM currency tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.currency_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW customer_view AS
  SELECT id, code, customer_type, customer_name,
    JSON_VALUE(customer_meta, N'$.tax_number') AS tax_number,
    JSON_VALUE(customer_meta, N'$.account') AS account,
    CASE WHEN JSON_VALUE(customer_meta, N'$.tax_free') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS tax_free,
    TRY_CONVERT(int, JSON_VALUE(customer_meta, N'$.terms')) AS terms,
    TRY_CONVERT(float, JSON_VALUE(customer_meta, N'$.credit_limit')) AS credit_limit,
    TRY_CONVERT(float, JSON_VALUE(customer_meta, N'$.discount')) AS discount,
    JSON_VALUE(customer_meta, N'$.notes') AS notes,
    CASE WHEN JSON_VALUE(customer_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_QUERY(customer_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(customer_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    addresses, contacts, events, customer_map, time_stamp,
    (SELECT
      id AS id, code AS code, customer_type AS customer_type, customer_name AS customer_name,
      JSON_QUERY(addresses) AS addresses, JSON_QUERY(contacts) AS contacts, JSON_QUERY(events) AS events,
      JSON_QUERY(customer_meta) AS customer_meta, JSON_QUERY(customer_map) AS customer_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS customer_object
  FROM customer
  WHERE deleted = 0;

CREATE OR ALTER VIEW customer_contacts AS
  SELECT c.id AS id, c.code, c.customer_name,
    JSON_VALUE(j.value, N'$.first_name') AS first_name,
    JSON_VALUE(j.value, N'$.surname') AS surname,
    JSON_VALUE(j.value, N'$.status') AS status,
    JSON_VALUE(j.value, N'$.phone') AS phone,
    JSON_VALUE(j.value, N'$.mobile') AS mobile,
    JSON_VALUE(j.value, N'$.email') AS email,
    JSON_VALUE(j.value, N'$.notes') AS notes,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.contact_map') AS contact_map
  FROM customer c
  CROSS APPLY OPENJSON(c.contacts) j
  WHERE c.deleted = 0;

CREATE OR ALTER VIEW customer_addresses AS
  SELECT c.id AS id, c.code, c.customer_name,
    JSON_VALUE(j.value, N'$.country') AS country,
    JSON_VALUE(j.value, N'$.state') AS state,
    JSON_VALUE(j.value, N'$.zip_code') AS zip_code,
    JSON_VALUE(j.value, N'$.city') AS city,
    JSON_VALUE(j.value, N'$.street') AS street,
    JSON_VALUE(j.value, N'$.notes') AS notes,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.address_map') AS address_map
  FROM customer c
  CROSS APPLY OPENJSON(c.addresses) j
  WHERE c.deleted = 0;

CREATE OR ALTER VIEW customer_events AS
  SELECT c.id AS id, c.code, c.customer_name,
    JSON_VALUE(j.value, N'$.uid') AS uid,
    JSON_VALUE(j.value, N'$.subject') AS subject,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.start_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.start_time'), N'T', N' '), 120)
    END AS start_time,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.end_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.end_time'), N'T', N' '), 120)
    END AS end_time,
    JSON_VALUE(j.value, N'$.place') AS place,
    JSON_VALUE(j.value, N'$.description') AS description,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.event_map') AS event_map
  FROM customer c
  CROSS APPLY OPENJSON(c.events) j
  WHERE c.deleted = 0;

CREATE OR ALTER VIEW customer_map AS
  SELECT tbl.id AS id, tbl.code, tbl.customer_name,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM customer tbl
  CROSS APPLY OPENJSON(tbl.customer_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW customer_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM customer tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.customer_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW employee_view AS
  SELECT id, code,
    CASE WHEN COALESCE(JSON_VALUE(employee_meta, N'$.start_date'), N'') = N'' THEN NULL ELSE TRY_CONVERT(date, JSON_VALUE(employee_meta, N'$.start_date')) END AS start_date,
    CASE WHEN COALESCE(JSON_VALUE(employee_meta, N'$.end_date'), N'') = N'' THEN NULL ELSE TRY_CONVERT(date, JSON_VALUE(employee_meta, N'$.end_date')) END AS end_date,
    CASE WHEN JSON_VALUE(employee_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_VALUE(employee_meta, N'$.notes') AS notes,
    JSON_VALUE(contact, N'$.first_name') AS first_name,
    JSON_VALUE(contact, N'$.surname') AS surname,
    JSON_VALUE(contact, N'$.status') AS status,
    JSON_VALUE(contact, N'$.phone') AS phone,
    JSON_VALUE(contact, N'$.mobile') AS mobile,
    JSON_VALUE(contact, N'$.email') AS email,
    JSON_VALUE(address, N'$.country') AS country,
    JSON_VALUE(address, N'$.state') AS state,
    JSON_VALUE(address, N'$.zip_code') AS zip_code,
    JSON_VALUE(address, N'$.city') AS city,
    JSON_VALUE(address, N'$.street') AS street,
    JSON_QUERY(employee_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(employee_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    employee_map, time_stamp,
    (SELECT
      id AS id, code AS code, JSON_QUERY(address) AS address, JSON_QUERY(contact) AS contact, JSON_QUERY(events) AS events,
      JSON_QUERY(employee_meta) AS employee_meta, JSON_QUERY(employee_map) AS employee_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS employee_object
  FROM employee
  WHERE deleted = 0;

CREATE OR ALTER VIEW employee_events AS
  SELECT e.id AS id, e.code,
    JSON_VALUE(e.contact, N'$.first_name') AS first_name,
    JSON_VALUE(e.contact, N'$.surname') AS surname,
    JSON_VALUE(j.value, N'$.uid') AS uid,
    JSON_VALUE(j.value, N'$.subject') AS subject,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.start_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.start_time'), N'T', N' '), 120)
    END AS start_time,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.end_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.end_time'), N'T', N' '), 120)
    END AS end_time,
    JSON_VALUE(j.value, N'$.place') AS place,
    JSON_VALUE(j.value, N'$.description') AS description,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.event_map') AS event_map
  FROM employee e
  CROSS APPLY OPENJSON(e.events) j
  WHERE e.deleted = 0;

CREATE OR ALTER VIEW employee_map AS
  SELECT tbl.id AS id, tbl.code,
    JSON_VALUE(tbl.contact, N'$.first_name') AS first_name,
    JSON_VALUE(tbl.contact, N'$.surname') AS surname,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM employee tbl
  CROSS APPLY OPENJSON(tbl.employee_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW employee_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM employee tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.employee_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW place_view AS
  SELECT id, code, place_type, place_name, currency_code,
    CASE WHEN JSON_VALUE(place_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_VALUE(place_meta, N'$.notes') AS notes,
    JSON_VALUE(address, N'$.country') AS country,
    JSON_VALUE(address, N'$.state') AS state,
    JSON_VALUE(address, N'$.zip_code') AS zip_code,
    JSON_VALUE(address, N'$.city') AS city,
    JSON_VALUE(address, N'$.street') AS street,
    JSON_QUERY(place_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(place_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    place_map, time_stamp,
    (SELECT
      id AS id, code AS code, place_type AS place_type, place_name AS place_name, currency_code AS currency_code,
      JSON_QUERY(address) AS address, JSON_QUERY(contacts) AS contacts, JSON_QUERY(events) AS events,
      JSON_QUERY(place_meta) AS place_meta, JSON_QUERY(place_map) AS place_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS place_object
  FROM place
  WHERE deleted = 0;

CREATE OR ALTER VIEW place_contacts AS
  SELECT p.id AS id, p.code, p.place_name,
    JSON_VALUE(j.value, N'$.first_name') AS first_name,
    JSON_VALUE(j.value, N'$.surname') AS surname,
    JSON_VALUE(j.value, N'$.status') AS status,
    JSON_VALUE(j.value, N'$.phone') AS phone,
    JSON_VALUE(j.value, N'$.mobile') AS mobile,
    JSON_VALUE(j.value, N'$.email') AS email,
    JSON_VALUE(j.value, N'$.notes') AS notes,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.contact_map') AS contact_map
  FROM place p
  CROSS APPLY OPENJSON(p.contacts) j
  WHERE p.deleted = 0;

CREATE OR ALTER VIEW place_events AS
  SELECT p.id AS id, p.code, p.place_name,
    JSON_VALUE(j.value, N'$.uid') AS uid,
    JSON_VALUE(j.value, N'$.subject') AS subject,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.start_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.start_time'), N'T', N' '), 120)
    END AS start_time,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.end_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.end_time'), N'T', N' '), 120)
    END AS end_time,
    JSON_VALUE(j.value, N'$.place') AS place,
    JSON_VALUE(j.value, N'$.description') AS description,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.event_map') AS event_map
  FROM place p
  CROSS APPLY OPENJSON(p.events) j
  WHERE p.deleted = 0;

CREATE OR ALTER VIEW place_map AS
  SELECT tbl.id AS id, tbl.code, tbl.place_name,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM place tbl
  CROSS APPLY OPENJSON(tbl.place_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW place_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM place tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.place_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW tax_view AS
  SELECT id, code,
    JSON_VALUE(tax_meta, N'$.description') AS description,
    TRY_CONVERT(float, JSON_VALUE(tax_meta, N'$.rate_value')) AS rate_value,
    CASE WHEN JSON_VALUE(tax_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_QUERY(tax_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(tax_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    tax_map, time_stamp,
    (SELECT
      id AS id, code AS code, JSON_QUERY(tax_meta) AS tax_meta, JSON_QUERY(tax_map) AS tax_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS tax_object
  FROM tax
  WHERE deleted = 0;

CREATE OR ALTER VIEW tax_map AS
  SELECT tbl.id AS id, tbl.code,
    JSON_VALUE(tbl.tax_meta, N'$.description') AS tax_description,
    JSON_VALUE(tbl.tax_meta, N'$.rate_value') AS rate_value,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM tax tbl
  CROSS APPLY OPENJSON(tbl.tax_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW tax_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM tax tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.tax_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW link_view AS
  SELECT id, code,
    link_type_1, link_code_1, link_type_2, link_code_2,
    TRY_CONVERT(float, JSON_VALUE(link_meta, N'$.qty')) AS qty,
    TRY_CONVERT(float, JSON_VALUE(link_meta, N'$.amount')) AS amount,
    TRY_CONVERT(float, JSON_VALUE(link_meta, N'$.rate')) AS rate,
    COALESCE(JSON_VALUE(link_meta, N'$.notes'), N'') AS notes,
    JSON_QUERY(link_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(link_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    link_map, time_stamp,
    (SELECT
      id AS id, code AS code, link_type_1 AS link_type_1, link_code_1 AS link_code_1,
      link_type_2 AS link_type_2, link_code_2 AS link_code_2,
      JSON_QUERY(link_meta) AS link_meta, JSON_QUERY(link_map) AS link_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS link_object
  FROM link
  WHERE deleted = 0;

CREATE OR ALTER VIEW link_map AS
  SELECT tbl.id AS id, tbl.code, tbl.link_type_1, tbl.link_code_1, tbl.link_type_2, tbl.link_code_2,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM link tbl
  CROSS APPLY OPENJSON(tbl.link_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW link_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM link tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.link_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW product_view AS
  SELECT id, code, product_type, product_name, tax_code,
    JSON_VALUE(product_meta, N'$.unit') AS unit,
    JSON_VALUE(product_meta, N'$.barcode_type') AS barcode_type,
    JSON_VALUE(product_meta, N'$.barcode') AS barcode,
    TRY_CONVERT(float, JSON_VALUE(product_meta, N'$.barcode_qty')) AS barcode_qty,
    CASE WHEN JSON_VALUE(product_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_VALUE(product_meta, N'$.notes') AS notes,
    JSON_QUERY(product_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(product_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    events, product_map, time_stamp,
    (SELECT
      id AS id, code AS code, product_type AS product_type, product_name AS product_name, tax_code AS tax_code,
      JSON_QUERY(events) AS events, JSON_QUERY(product_meta) AS product_meta, JSON_QUERY(product_map) AS product_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS product_object
  FROM product
  WHERE deleted = 0;

CREATE OR ALTER VIEW product_events AS
  SELECT p.id AS id, p.code, p.product_name,
    JSON_VALUE(j.value, N'$.uid') AS uid,
    JSON_VALUE(j.value, N'$.subject') AS subject,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.start_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.start_time'), N'T', N' '), 120)
    END AS start_time,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.end_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.end_time'), N'T', N' '), 120)
    END AS end_time,
    JSON_VALUE(j.value, N'$.place') AS place,
    JSON_VALUE(j.value, N'$.description') AS description,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.event_map') AS event_map
  FROM product p
  CROSS APPLY OPENJSON(p.events) j
  WHERE p.deleted = 0;

CREATE OR ALTER VIEW product_map AS
  SELECT tbl.id AS id, tbl.code, tbl.product_name,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM product tbl
  CROSS APPLY OPENJSON(tbl.product_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW product_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM product tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.product_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW product_components AS
  SELECT p.id, p.code AS product_code, p.product_name,
    COALESCE(JSON_VALUE(p.product_meta, N'$.unit'), N'') AS unit,
    c.code AS ref_product_code, c.product_name AS component_name,
    COALESCE(JSON_VALUE(c.product_meta, N'$.unit'), N'') AS component_unit,
    c.product_type AS component_type,
    TRY_CONVERT(float, JSON_VALUE(l.link_meta, N'$.qty')) AS qty,
    COALESCE(JSON_VALUE(l.link_meta, N'$.notes'), N'') AS notes
  FROM product p
  INNER JOIN link l ON l.link_code_1 = p.code
  INNER JOIN product c ON l.link_code_2 = c.code
  WHERE p.product_type = N'PRODUCT_VIRTUAL'
    AND l.link_type_1 = N'LINK_PRODUCT' AND l.link_type_2 = N'LINK_PRODUCT'
    AND p.deleted = 0 AND l.deleted = 0 AND c.deleted = 0;

CREATE OR ALTER VIEW project_view AS
  SELECT id, code, project_name, customer_code,
    CASE WHEN COALESCE(JSON_VALUE(project_meta, N'$.start_date'), N'') = N'' THEN NULL ELSE TRY_CONVERT(date, JSON_VALUE(project_meta, N'$.start_date')) END AS start_date,
    CASE WHEN COALESCE(JSON_VALUE(project_meta, N'$.end_date'), N'') = N'' THEN NULL ELSE TRY_CONVERT(date, JSON_VALUE(project_meta, N'$.end_date')) END AS end_date,
    CASE WHEN JSON_VALUE(project_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_VALUE(project_meta, N'$.notes') AS notes,
    JSON_QUERY(project_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(project_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    addresses, contacts, events, project_map, time_stamp,
    (SELECT
      id AS id, code AS code, project_name AS project_name, customer_code AS customer_code,
      JSON_QUERY(addresses) AS addresses, JSON_QUERY(contacts) AS contacts, JSON_QUERY(events) AS events,
      JSON_QUERY(project_meta) AS project_meta, JSON_QUERY(project_map) AS project_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS project_object
  FROM project
  WHERE deleted = 0;

CREATE OR ALTER VIEW project_contacts AS
  SELECT p.id AS id, p.code, p.project_name,
    JSON_VALUE(j.value, N'$.first_name') AS first_name,
    JSON_VALUE(j.value, N'$.surname') AS surname,
    JSON_VALUE(j.value, N'$.status') AS status,
    JSON_VALUE(j.value, N'$.phone') AS phone,
    JSON_VALUE(j.value, N'$.mobile') AS mobile,
    JSON_VALUE(j.value, N'$.email') AS email,
    JSON_VALUE(j.value, N'$.notes') AS notes,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.contact_map') AS contact_map
  FROM project p
  CROSS APPLY OPENJSON(p.contacts) j
  WHERE p.deleted = 0;

CREATE OR ALTER VIEW project_addresses AS
  SELECT p.id AS id, p.code, p.project_name,
    JSON_VALUE(j.value, N'$.country') AS country,
    JSON_VALUE(j.value, N'$.state') AS state,
    JSON_VALUE(j.value, N'$.zip_code') AS zip_code,
    JSON_VALUE(j.value, N'$.city') AS city,
    JSON_VALUE(j.value, N'$.street') AS street,
    JSON_VALUE(j.value, N'$.notes') AS notes,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.address_map') AS address_map
  FROM project p
  CROSS APPLY OPENJSON(p.addresses) j
  WHERE p.deleted = 0;

CREATE OR ALTER VIEW project_events AS
  SELECT p.id AS id, p.code, p.project_name,
    JSON_VALUE(j.value, N'$.uid') AS uid,
    JSON_VALUE(j.value, N'$.subject') AS subject,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.start_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.start_time'), N'T', N' '), 120)
    END AS start_time,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.end_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.end_time'), N'T', N' '), 120)
    END AS end_time,
    JSON_VALUE(j.value, N'$.place') AS place,
    JSON_VALUE(j.value, N'$.description') AS description,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.event_map') AS event_map
  FROM project p
  CROSS APPLY OPENJSON(p.events) j
  WHERE p.deleted = 0;

CREATE OR ALTER VIEW project_map AS
  SELECT tbl.id AS id, tbl.code, tbl.project_name,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM project tbl
  CROSS APPLY OPENJSON(tbl.project_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW project_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM project tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.project_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW rate_view AS
  SELECT id, code, rate_type, TRY_CONVERT(date, rate_date) as rate_date, place_code, currency_code,
    TRY_CONVERT(float, JSON_VALUE(rate_meta, N'$.rate_value')) AS rate_value,
    JSON_QUERY(rate_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(rate_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    rate_map, time_stamp,
    (SELECT
      id AS id, code AS code, rate_type AS rate_type, rate_date AS rate_date, place_code AS place_code, currency_code AS currency_code,
      JSON_QUERY(rate_meta) AS rate_meta, JSON_QUERY(rate_map) AS rate_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS rate_object
  FROM rate
  WHERE deleted = 0;

CREATE OR ALTER VIEW rate_map AS
  SELECT tbl.id AS id, tbl.code,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM rate tbl
  CROSS APPLY OPENJSON(tbl.rate_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW rate_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM rate tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.rate_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW tool_view AS
  SELECT id, code, product_code, description,
    JSON_VALUE(tool_meta, N'$.serial_number') AS serial_number,
    CASE WHEN JSON_VALUE(tool_meta, N'$.inactive') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS inactive,
    JSON_VALUE(tool_meta, N'$.notes') AS notes,
    JSON_QUERY(tool_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(tool_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    tool_map, time_stamp,
    (SELECT
      id AS id, code AS code, description AS description, product_code AS product_code,
      JSON_QUERY(events) AS events, JSON_QUERY(tool_meta) AS tool_meta, JSON_QUERY(tool_map) AS tool_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS tool_object
  FROM tool
  WHERE deleted = 0;

CREATE OR ALTER VIEW tool_events AS
  SELECT t.id AS id, t.code, t.description AS tool_description,
    JSON_VALUE(j.value, N'$.uid') AS uid,
    JSON_VALUE(j.value, N'$.subject') AS subject,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.start_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.start_time'), N'T', N' '), 120)
    END AS start_time,
    CASE
      WHEN COALESCE(JSON_VALUE(j.value, N'$.end_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(j.value, N'$.end_time'), N'T', N' '), 120)
    END AS end_time,
    JSON_VALUE(j.value, N'$.place') AS place,
    JSON_VALUE(j.value, N'$.description') AS description,
    JSON_QUERY(j.value, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_QUERY(j.value, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    JSON_QUERY(j.value, N'$.event_map') AS event_map
  FROM tool t
  CROSS APPLY OPENJSON(t.events) j
  WHERE t.deleted = 0;

CREATE OR ALTER VIEW tool_map AS
  SELECT tbl.id AS id, tbl.code, tbl.description AS tool_description,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM tool tbl
  CROSS APPLY OPENJSON(tbl.tool_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW tool_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM tool tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.tool_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW price_view AS
  SELECT id, code, price_type, TRY_CONVERT(date, valid_from) as valid_from, TRY_CONVERT(date, valid_to) as valid_to, 
    product_code, currency_code, customer_code, qty,
    TRY_CONVERT(float, JSON_VALUE(price_meta, N'$.price_value')) AS price_value,
    JSON_QUERY(price_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(price_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    price_map, time_stamp,
    (SELECT
      id AS id, code AS code, price_type AS price_type, valid_from AS valid_from, valid_to AS valid_to,
      product_code AS product_code, currency_code AS currency_code, customer_code AS customer_code, qty AS qty,
      JSON_QUERY(price_meta) AS price_meta, JSON_QUERY(price_map) AS price_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS price_object
  FROM price
  WHERE deleted = 0;

CREATE OR ALTER VIEW price_map AS
  SELECT tbl.id AS id, tbl.code,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM price tbl
  CROSS APPLY OPENJSON(tbl.price_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW price_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM price tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.price_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW trans_view AS
  SELECT id, code, trans_type, direction, TRY_CONVERT(date, trans_date) as trans_date, trans_code, customer_code, employee_code,
    project_code, place_code, currency_code, auth_code,
    CASE
      WHEN COALESCE(JSON_VALUE(trans_meta, N'$.due_time'), N'') = N'' THEN NULL
      ELSE TRY_CONVERT(datetime2(0), REPLACE(JSON_VALUE(trans_meta, N'$.due_time'), N'T', N' '), 120)
    END AS due_time,
    JSON_VALUE(trans_meta, N'$.ref_number') AS ref_number,
    JSON_VALUE(trans_meta, N'$.paid_type') AS paid_type,
    CASE WHEN JSON_VALUE(trans_meta, N'$.tax_free') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS tax_free,
    CASE WHEN JSON_VALUE(trans_meta, N'$.paid') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS paid,
    TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.rate')) AS rate,
    CASE WHEN JSON_VALUE(trans_meta, N'$.closed') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS closed,
    JSON_VALUE(trans_meta, N'$.status') AS status,
    JSON_VALUE(trans_meta, N'$.trans_state') AS trans_state,
    JSON_VALUE(trans_meta, N'$.notes') AS notes,
    JSON_VALUE(trans_meta, N'$.internal_notes') AS internal_notes,
    JSON_VALUE(trans_meta, N'$.report_notes') AS report_notes,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.worksheet.distance')), 0) AS worksheet_distance,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.worksheet.repair')), 0) AS worksheet_repair,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.worksheet.total')), 0) AS worksheet_total,
    COALESCE(JSON_VALUE(trans_meta, N'$.worksheet.justification'), N'') AS worksheet_justification,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.rent.holiday')), 0) AS rent_holiday,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.rent.bad_tool')), 0) AS rent_bad_tool,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(trans_meta, N'$.rent.other')), 0) AS rent_other,
    COALESCE(JSON_VALUE(trans_meta, N'$.rent.justification'), N'') AS rent_justification,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.company_name'), N'') AS invoice_company_name,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.company_account'), N'') AS invoice_company_account,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.company_address'), N'') AS invoice_company_address,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.company_tax_number'), N'') AS invoice_company_tax_number,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.customer_name'), N'') AS invoice_customer_name,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.customer_account'), N'') AS invoice_customer_account,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.customer_address'), N'') AS invoice_customer_address,
    COALESCE(JSON_VALUE(trans_meta, N'$.invoice.customer_tax_number'), N'') AS invoice_customer_tax_number,
    JSON_QUERY(trans_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(trans_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    trans_map, time_stamp,
    (SELECT
      id AS id, code AS code, trans_type AS trans_type, direction AS direction, trans_date AS trans_date,
      trans_code AS trans_code, customer_code AS customer_code, employee_code AS employee_code, project_code AS project_code,
      place_code AS place_code, currency_code AS currency_code, auth_code AS auth_code,
      JSON_QUERY(trans_meta) AS trans_meta, JSON_QUERY(trans_map) AS trans_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS trans_object
  FROM trans
  WHERE deleted = 0;

CREATE OR ALTER VIEW trans_map AS
  SELECT tbl.id AS id, tbl.code, tbl.trans_type, tbl.direction, tbl.trans_date,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM trans tbl
  CROSS APPLY OPENJSON(tbl.trans_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW trans_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM trans tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.trans_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0;

CREATE OR ALTER VIEW item_view AS
  SELECT id, code, trans_code, product_code, tax_code,
    JSON_VALUE(item_meta, N'$.unit') AS unit,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.qty')) AS qty,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.fx_price')) AS fx_price,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.net_amount')) AS net_amount,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.discount')) AS discount,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.vat_amount')) AS vat_amount,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.amount')) AS amount,
    JSON_VALUE(item_meta, N'$.description') AS description,
    CASE WHEN JSON_VALUE(item_meta, N'$.deposit') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS deposit,
    CASE WHEN JSON_VALUE(item_meta, N'$.action_price') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS action_price,
    TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.own_stock')) AS own_stock,
    JSON_QUERY(item_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(item_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    item_map, time_stamp,
    (SELECT
      id AS id, code AS code, trans_code AS trans_code, product_code AS product_code, tax_code AS tax_code,
      JSON_QUERY(item_meta) AS item_meta, JSON_QUERY(item_map) AS item_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS item_object
  FROM item
  WHERE deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = item.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW item_map AS
  SELECT tbl.id AS id, tbl.code, tbl.trans_code, tbl.product_code,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM item tbl
  CROSS APPLY OPENJSON(tbl.item_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = tbl.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW item_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM item tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.item_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = tbl.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW movement_view AS
  SELECT id, code, movement_type, shipping_time,
    trans_code, product_code, tool_code, place_code, item_code, movement_code,
    JSON_VALUE(movement_meta, N'$.qty') AS qty,
    JSON_VALUE(movement_meta, N'$.shared') AS shared,
    JSON_VALUE(movement_meta, N'$.notes') AS notes,
    JSON_QUERY(movement_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(movement_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    movement_map, time_stamp,
    (SELECT
      id AS id, code AS code, movement_type AS movement_type, shipping_time AS shipping_time,
      trans_code AS trans_code, product_code AS product_code, tool_code AS tool_code, place_code AS place_code,
      item_code AS item_code, movement_code AS movement_code,
      JSON_QUERY(movement_meta) AS movement_meta, JSON_QUERY(movement_map) AS movement_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS movement_object
  FROM movement
  WHERE deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = movement.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW movement_map AS
  SELECT tbl.id AS id, tbl.code, tbl.trans_code,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM movement tbl
  CROSS APPLY OPENJSON(tbl.movement_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = tbl.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW movement_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM movement tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.movement_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = tbl.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW movement_stock AS
  SELECT ROW_NUMBER() OVER (ORDER BY pl.place_name, p.product_name) AS id,
    mv.place_code, pl.place_name, mv.product_code, p.product_name,
    JSON_VALUE(p.product_meta, N'$.unit') AS unit,
    JSON_VALUE(mv.movement_meta, N'$.notes') AS batch_no,
    SUM(COALESCE(TRY_CONVERT(float, JSON_VALUE(mv.movement_meta, N'$.qty')), 0)) AS qty,
    MAX(CONVERT(date, mv.shipping_time)) AS posdate
  FROM movement mv
  INNER JOIN place pl ON mv.place_code = pl.code
  INNER JOIN product p ON mv.product_code = p.code
  WHERE mv.movement_type = N'MOVEMENT_INVENTORY'
    AND mv.deleted = 0 AND p.deleted = 0 AND pl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = mv.trans_code AND t.deleted = 0)
  GROUP BY mv.place_code, pl.place_name, mv.product_code, p.product_name, JSON_VALUE(p.product_meta, N'$.unit'), JSON_VALUE(mv.movement_meta, N'$.notes')
  HAVING SUM(COALESCE(TRY_CONVERT(float, JSON_VALUE(mv.movement_meta, N'$.qty')), 0)) <> 0;

CREATE OR ALTER VIEW movement_inventory AS
  SELECT mt.id, mt.code, mt.trans_code, t.trans_type, t.direction, CONVERT(date, mt.shipping_time) AS shipping_date,
    mt.place_code, pl.place_name, mt.product_code, p.product_name,
    JSON_VALUE(p.product_meta, N'$.unit') AS unit,
    JSON_VALUE(mt.movement_meta, N'$.notes') AS batch_no,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(mt.movement_meta, N'$.qty')), 0) AS qty,
    it.customer_code, ci.customer_name,
    COALESCE(i.trans_code, mr.trans_code, t.trans_code) AS ref_trans_code
  FROM movement mt
  INNER JOIN trans t ON mt.trans_code = t.code
  INNER JOIN place pl ON mt.place_code = pl.code
  INNER JOIN product p ON mt.product_code = p.code
  LEFT JOIN item i ON mt.item_code = i.code AND i.deleted = 0
  LEFT JOIN trans it ON i.trans_code = it.code AND it.deleted = 0
  LEFT JOIN customer ci ON it.customer_code = ci.code AND ci.deleted = 0
  LEFT JOIN movement mr ON mt.movement_code = mr.code AND mr.deleted = 0
  WHERE mt.movement_type = N'MOVEMENT_INVENTORY'
    AND mt.deleted = 0 AND t.deleted = 0 AND pl.deleted = 0 AND p.deleted = 0;

CREATE OR ALTER VIEW movement_waybill AS
  SELECT mv.id, mv.code, t.code AS trans_code, t.direction, t.trans_code AS ref_trans_code, mv.shipping_time,
    mv.tool_code, JSON_VALUE(tl.tool_meta, N'$.serial_number') AS serial_number, tl.description,
    JSON_VALUE(mv.movement_meta, N'$.notes') AS mvnotes,
    t.employee_code, t.customer_code, c.customer_name,
    JSON_VALUE(t.trans_meta, N'$.trans_state') AS trans_state,
    JSON_VALUE(t.trans_meta, N'$.notes') AS notes,
    JSON_VALUE(t.trans_meta, N'$.internal_notes') AS internal_notes,
    CASE WHEN JSON_VALUE(t.trans_meta, N'$.closed') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS closed,
    t.time_stamp
  FROM trans t
  INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN tool tl ON mv.tool_code = tl.code
  LEFT JOIN customer c ON t.customer_code = c.code
  WHERE t.trans_type = N'TRANS_WAYBILL' AND mv.deleted = 0 AND t.deleted = 0;

CREATE OR ALTER VIEW movement_formula AS
  SELECT mv.id, mv.code, t.code AS trans_code,
    CASE WHEN mv.movement_type = N'MOVEMENT_HEAD' THEN N'IN' ELSE N'OUT' END AS direction,
    mv.product_code, p.product_name, JSON_VALUE(p.product_meta, N'$.unit') AS unit,
    COALESCE(TRY_CONVERT(float, JSON_VALUE(mv.movement_meta, N'$.qty')), 0) AS qty,
    JSON_VALUE(mv.movement_meta, N'$.notes') AS batch_no,
    mv.place_code, pl.place_name,
    CASE WHEN JSON_VALUE(mv.movement_meta, N'$.shared') = N'true' THEN CAST(1 AS bit) ELSE CAST(0 AS bit) END AS shared
  FROM trans t
  INNER JOIN movement mv ON mv.trans_code = t.code
  INNER JOIN product p ON mv.product_code = p.code
  LEFT JOIN place pl ON mv.place_code = pl.code
  WHERE t.trans_type = N'TRANS_FORMULA' AND mv.deleted = 0 AND t.deleted = 0;

CREATE OR ALTER VIEW item_shipping AS
  SELECT iv.id, iv.code, iv.trans_code, iv.direction, iv.product_code, iv.product_name, iv.unit, iv.item_qty,
    SUM(COALESCE(TRY_CONVERT(float, JSON_VALUE(mv.movement_meta, N'$.qty')), 0)) AS movement_qty
  FROM (
    SELECT i.id, i.code, i.trans_code, t.direction, i.product_code, p.product_name,
      COALESCE(JSON_VALUE(p.product_meta, N'$.unit'), N'') AS unit,
      COALESCE(TRY_CONVERT(float, JSON_VALUE(i.item_meta, N'$.qty')), 0) AS item_qty
    FROM item i
    INNER JOIN trans t ON i.trans_code = t.code
    INNER JOIN product p ON i.product_code = p.code
    WHERE t.deleted = 0 AND i.deleted = 0 AND p.product_type = N'PRODUCT_ITEM'
      AND t.trans_type IN (N'TRANS_ORDER', N'TRANS_WORKSHEET', N'TRANS_RENT')
    UNION ALL
    SELECT i.id, i.code, i.trans_code, t.direction, pc.ref_product_code AS product_code, pc.component_name AS product_name,
      pc.component_unit AS unit,
      COALESCE(TRY_CONVERT(float, JSON_VALUE(i.item_meta, N'$.qty')), 0) * COALESCE(pc.qty, 0) AS item_qty
    FROM item i
    INNER JOIN trans t ON i.trans_code = t.code
    INNER JOIN product_components pc ON i.product_code = pc.product_code AND pc.component_type = N'PRODUCT_ITEM'
    WHERE t.deleted = 0 AND i.deleted = 0
      AND t.trans_type IN (N'TRANS_ORDER', N'TRANS_WORKSHEET', N'TRANS_RENT')
  ) iv
  LEFT JOIN movement mv ON mv.item_code = iv.code AND mv.product_code = iv.product_code AND mv.deleted = 0
  GROUP BY iv.id, iv.code, iv.trans_code, iv.direction, iv.product_code, iv.product_name, iv.unit, iv.item_qty;

CREATE OR ALTER VIEW payment_view AS
  SELECT id, code, TRY_CONVERT(date, paid_date) as paid_date, trans_code,
    JSON_VALUE(payment_meta, N'$.amount') AS amount,
    JSON_VALUE(payment_meta, N'$.notes') AS notes,
    JSON_QUERY(payment_meta, N'$.tags') AS tags,
    REPLACE(REPLACE(REPLACE(COALESCE(JSON_VALUE(payment_meta, N'$.tags'), N'[]'), '"', ''), '[', ''), ']', '') AS tag_lst,
    payment_map, time_stamp,
    (SELECT
      id AS id, code AS code, paid_date AS paid_date, trans_code AS trans_code,
      JSON_QUERY(payment_meta) AS payment_meta, JSON_QUERY(payment_map) AS payment_map, time_stamp AS time_stamp
     FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) AS payment_object
  FROM payment
  WHERE deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = payment.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW payment_map AS
  SELECT tbl.id AS id, tbl.code, tbl.paid_date, tbl.trans_code,
    j.[key] AS map_key, j.value AS map_value, j.[type] AS map_type,
    COALESCE(cf.description, j.[key]) AS description,
    COALESCE(cf.field_type, N'FIELD_STRING') AS field_type,
    CASE
      WHEN cf.field_type = N'FIELD_BOOL' THEN N'bool'
      WHEN cf.field_type = N'FIELD_INTEGER' THEN N'integer'
      WHEN cf.field_type = N'FIELD_NUMBER' THEN N'float'
      WHEN cf.field_type = N'FIELD_DATE' THEN N'date'
      WHEN cf.field_type = N'FIELD_DATETIME' THEN N'datetime'
      WHEN cf.field_type IN (
        N'FIELD_URL', N'FIELD_CUSTOMER', N'FIELD_EMPLOYEE', N'FIELD_PLACE', N'FIELD_PRODUCT', N'FIELD_PROJECT',
        N'FIELD_TOOL', N'FIELD_TRANS_ITEM', N'FIELD_TRANS_MOVEMENT', N'FIELD_TRANS_PAYMENT'
      ) THEN N'link'
      ELSE N'string'
    END AS value_meta
  FROM payment tbl
  CROSS APPLY OPENJSON(tbl.payment_map) j
  LEFT JOIN config_map cf ON j.[key] = cf.field_name
  WHERE tbl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = tbl.trans_code AND t.deleted = 0);

CREATE OR ALTER VIEW payment_invoice AS
  SELECT pm.id, pm.code, pm.trans_code, pt.trans_type, pt.direction, pm.paid_date, pl.place_name, pl.currency_code,
    TRY_CONVERT(float, JSON_VALUE(l.link_meta, N'$.amount')) AS paid_amount,
    TRY_CONVERT(float, JSON_VALUE(l.link_meta, N'$.rate')) AS paid_rate,
    it.code AS ref_trans_code, it.currency_code AS invoice_curr,
    im.amount AS invoice_amount,
    JSON_VALUE(pm.payment_meta, N'$.notes') AS description
  FROM link l
  INNER JOIN payment pm ON l.link_code_1 = pm.code
  INNER JOIN trans pt ON pm.trans_code = pt.code
  INNER JOIN place pl ON pt.place_code = pl.code
  INNER JOIN trans it ON l.link_code_2 = it.code
  INNER JOIN (
    SELECT trans_code, SUM(COALESCE(TRY_CONVERT(float, JSON_VALUE(item_meta, N'$.amount')), 0)) AS amount
    FROM item
    GROUP BY trans_code
  ) im ON it.code = im.trans_code
  WHERE l.link_type_1 = N'LINK_PAYMENT' AND l.link_type_2 = N'LINK_TRANS'
    AND it.trans_type IN (N'TRANS_INVOICE', N'TRANS_RECEIPT');

CREATE OR ALTER VIEW payment_tags AS
  SELECT tbl.id AS id, tbl.code, j.value AS tag
  FROM payment tbl
  CROSS APPLY OPENJSON(COALESCE(JSON_QUERY(tbl.payment_meta, N'$.tags'), N'[]')) j
  WHERE tbl.deleted = 0
    AND EXISTS (SELECT 1 FROM trans t WHERE t.code = tbl.trans_code AND t.deleted = 0);
    
/* =========================
   Triggers
   ========================= */
CREATE OR ALTER TRIGGER usref_changed_timestamp
ON usref
AFTER UPDATE
AS
BEGIN
  SET NOCOUNT ON
  UPDATE u
  SET changed = SYSUTCDATETIME()
  FROM usref u
  INNER JOIN inserted i ON i.id = u.id
END;

CREATE OR ALTER TRIGGER config_default_code
ON config
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE c
  SET c.code = CONCAT(N'CNF', @epoch, N'N', c.id)
  FROM config c
  INNER JOIN inserted i ON i.id = c.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER auth_default_code
ON auth
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE a
  SET a.code = CONCAT(N'USR', @epoch, N'N', a.id)
  FROM auth a
  INNER JOIN inserted i ON i.id = a.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER customer_default_code
ON customer
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE c
  SET c.code = CONCAT(N'CUS', @epoch, N'N', c.id)
  FROM customer c
  INNER JOIN inserted i ON i.id = c.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER employee_default_code
ON employee
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE e
  SET e.code = CONCAT(N'EMP', @epoch, N'N', e.id)
  FROM employee e
  INNER JOIN inserted i ON i.id = e.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER place_default_code
ON place
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE p
  SET p.code = CONCAT(N'PLA', @epoch, N'N', p.id)
  FROM place p
  INNER JOIN inserted i ON i.id = p.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER link_default_code
ON link
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE l
  SET l.code = CONCAT(N'LNK', @epoch, N'N', l.id)
  FROM link l
  INNER JOIN inserted i ON i.id = l.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER product_default_code
ON product
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE p
  SET p.code = CONCAT(N'PRD', @epoch, N'N', p.id)
  FROM product p
  INNER JOIN inserted i ON i.id = p.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER project_default_code
ON project
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE p
  SET p.code = CONCAT(N'PRJ', @epoch, N'N', p.id)
  FROM project p
  INNER JOIN inserted i ON i.id = p.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER rate_default_code
ON rate
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE r
  SET r.code = CONCAT(N'RAT', @epoch, N'N', r.id)
  FROM rate r
  INNER JOIN inserted i ON i.id = r.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER tool_default_code
ON tool
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE t
  SET t.code = CONCAT(N'SER', @epoch, N'N', t.id)
  FROM tool t
  INNER JOIN inserted i ON i.id = t.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER price_default_code
ON price
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE p
  SET p.code = CONCAT(N'PRC', @epoch, N'N', p.id)
  FROM price p
  INNER JOIN inserted i ON i.id = p.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

/* trans: default code + invoice JSON patch on insert/update */
CREATE OR ALTER TRIGGER trans_default_code
ON trans
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE t
  SET t.code = CONCAT(
      CASE
        WHEN t.trans_type = N'TRANS_INVENTORY' THEN N'COR'
        WHEN t.trans_type = N'TRANS_DELIVERY' AND t.direction = N'DIRECTION_TRANSFER' THEN N'TRF'
        ELSE SUBSTRING(t.trans_type, 7, 3)
      END,
      @epoch, N'N', t.id
    )
  FROM trans t
  INNER JOIN inserted i ON i.id = t.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''

  /* Invoice patch (only for outgoing invoices) */
  UPDATE t
  SET t.trans_meta =
    JSON_MODIFY(
      COALESCE(t.trans_meta, N'{}'),
      N'$.invoice',
      JSON_QUERY(inv.invoice_json)
    )
  FROM trans t
  INNER JOIN inserted i ON i.id = t.id
  OUTER APPLY (
    SELECT (
      SELECT
        cu.customer_name AS customer_name,
        JSON_VALUE(cu.customer_meta, N'$.tax_number') AS customer_tax_number,
        JSON_VALUE(cu.customer_meta, N'$.account') AS customer_account,
        LTRIM(RTRIM(CONCAT(
          COALESCE(JSON_VALUE(cu.addresses, N'$[0].zip_code'), N''), N' ',
          COALESCE(JSON_VALUE(cu.addresses, N'$[0].city'), N''), N' ',
          COALESCE(JSON_VALUE(cu.addresses, N'$[0].street'), N'')
        ))) AS customer_address,
        co.customer_name AS company_name,
        JSON_VALUE(co.customer_meta, N'$.tax_number') AS company_tax_number,
        JSON_VALUE(co.customer_meta, N'$.account') AS company_account,
        LTRIM(RTRIM(CONCAT(
          COALESCE(JSON_VALUE(co.addresses, N'$[0].zip_code'), N''), N' ',
          COALESCE(JSON_VALUE(co.addresses, N'$[0].city'), N''), N' ',
          COALESCE(JSON_VALUE(co.addresses, N'$[0].street'), N'')
        ))) AS company_address
      FOR JSON PATH, WITHOUT_ARRAY_WRAPPER
    ) AS invoice_json
    FROM customer cu
    INNER JOIN customer co ON co.customer_type = N'CUSTOMER_OWN'
    WHERE cu.code = t.customer_code
  ) inv
  WHERE i.trans_type = N'TRANS_INVOICE'
    AND i.direction = N'DIRECTION_OUT'
    AND inv.invoice_json IS NOT NULL
END;

CREATE OR ALTER TRIGGER trans_invoice_customer_update
ON trans
AFTER UPDATE
AS
BEGIN
  SET NOCOUNT ON

  /* Invoice patch for updated outgoing invoices */
  UPDATE t
  SET t.trans_meta =
    JSON_MODIFY(
      COALESCE(t.trans_meta, N'{}'),
      N'$.invoice',
      JSON_QUERY(inv.invoice_json)
    )
  FROM trans t
  INNER JOIN inserted i ON i.id = t.id
  OUTER APPLY (
    SELECT (
      SELECT
        cu.customer_name AS customer_name,
        JSON_VALUE(cu.customer_meta, N'$.tax_number') AS customer_tax_number,
        JSON_VALUE(cu.customer_meta, N'$.account') AS customer_account,
        LTRIM(RTRIM(CONCAT(
          COALESCE(JSON_VALUE(cu.addresses, N'$[0].zip_code'), N''), N' ',
          COALESCE(JSON_VALUE(cu.addresses, N'$[0].city'), N''), N' ',
          COALESCE(JSON_VALUE(cu.addresses, N'$[0].street'), N'')
        ))) AS customer_address,
        co.customer_name AS company_name,
        JSON_VALUE(co.customer_meta, N'$.tax_number') AS company_tax_number,
        JSON_VALUE(co.customer_meta, N'$.account') AS company_account,
        LTRIM(RTRIM(CONCAT(
          COALESCE(JSON_VALUE(co.addresses, N'$[0].zip_code'), N''), N' ',
          COALESCE(JSON_VALUE(co.addresses, N'$[0].city'), N''), N' ',
          COALESCE(JSON_VALUE(co.addresses, N'$[0].street'), N'')
        ))) AS company_address
      FOR JSON PATH, WITHOUT_ARRAY_WRAPPER
    ) AS invoice_json
    FROM customer cu
    INNER JOIN customer co ON co.customer_type = N'CUSTOMER_OWN'
    WHERE cu.code = t.customer_code
  ) inv
  WHERE i.trans_type = N'TRANS_INVOICE'
    AND i.direction = N'DIRECTION_OUT'
    AND inv.invoice_json IS NOT NULL
END;

CREATE OR ALTER TRIGGER item_default_code
ON item
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE it
  SET it.code = CONCAT(N'ITM', @epoch, N'N', it.id)
  FROM item it
  INNER JOIN inserted i ON i.id = it.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER movement_default_code
ON movement
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE mv
  SET mv.code = CONCAT(N'MOV', @epoch, N'N', mv.id)
  FROM movement mv
  INNER JOIN inserted i ON i.id = mv.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER payment_default_code
ON payment
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE p
  SET p.code = CONCAT(N'PMT', @epoch, N'N', p.id)
  FROM payment p
  INNER JOIN inserted i ON i.id = p.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

CREATE OR ALTER TRIGGER log_default_code
ON [log]
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON
  DECLARE @epoch BIGINT = DATEDIFF(SECOND, '1970-01-01', SYSUTCDATETIME())
  UPDATE l
  SET l.code = CONCAT(N'LOG', @epoch, N'N', l.id)
  FROM [log] l
  INNER JOIN inserted i ON i.id = l.id
  WHERE i.code IS NULL OR LTRIM(RTRIM(i.code)) = N''
END;

/* link validation triggers (mimic postgres link_changed) */
CREATE OR ALTER TRIGGER link_insert
ON link
AFTER INSERT
AS
BEGIN
  SET NOCOUNT ON

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_CUSTOMER' AND NOT EXISTS (SELECT 1 FROM customer c WHERE c.code = i.link_code_1))
    THROW 50000, 'Invalid customer code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_CUSTOMER' AND NOT EXISTS (SELECT 1 FROM customer c WHERE c.code = i.link_code_2))
    THROW 50000, 'Invalid customer code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_EMPLOYEE' AND NOT EXISTS (SELECT 1 FROM employee e WHERE e.code = i.link_code_1))
    THROW 50000, 'Invalid employee code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_EMPLOYEE' AND NOT EXISTS (SELECT 1 FROM employee e WHERE e.code = i.link_code_2))
    THROW 50000, 'Invalid employee code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_ITEM' AND NOT EXISTS (SELECT 1 FROM item t WHERE t.code = i.link_code_1))
    THROW 50000, 'Invalid item code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_ITEM' AND NOT EXISTS (SELECT 1 FROM item t WHERE t.code = i.link_code_2))
    THROW 50000, 'Invalid item code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_MOVEMENT' AND NOT EXISTS (SELECT 1 FROM movement m WHERE m.code = i.link_code_1))
    THROW 50000, 'Invalid movement code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_MOVEMENT' AND NOT EXISTS (SELECT 1 FROM movement m WHERE m.code = i.link_code_2))
    THROW 50000, 'Invalid movement code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PAYMENT' AND NOT EXISTS (SELECT 1 FROM payment p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid payment code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PAYMENT' AND NOT EXISTS (SELECT 1 FROM payment p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid payment code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PLACE' AND NOT EXISTS (SELECT 1 FROM place p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid place code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PLACE' AND NOT EXISTS (SELECT 1 FROM place p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid place code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PRODUCT' AND NOT EXISTS (SELECT 1 FROM product p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid product code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PRODUCT' AND NOT EXISTS (SELECT 1 FROM product p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid product code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PROJECT' AND NOT EXISTS (SELECT 1 FROM project p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid project code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PROJECT' AND NOT EXISTS (SELECT 1 FROM project p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid project code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_TOOL' AND NOT EXISTS (SELECT 1 FROM tool t WHERE t.code = i.link_code_1))
    THROW 50000, 'Invalid tool code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_TOOL' AND NOT EXISTS (SELECT 1 FROM tool t WHERE t.code = i.link_code_2))
    THROW 50000, 'Invalid tool code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_TRANS' AND NOT EXISTS (SELECT 1 FROM trans t WHERE t.code = i.link_code_1))
    THROW 50000, 'Invalid trans code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_TRANS' AND NOT EXISTS (SELECT 1 FROM trans t WHERE t.code = i.link_code_2))
    THROW 50000, 'Invalid trans code (link_code_2).', 1
END;

CREATE OR ALTER TRIGGER link_update
ON link
AFTER UPDATE
AS
BEGIN
  SET NOCOUNT ON

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_CUSTOMER' AND NOT EXISTS (SELECT 1 FROM customer c WHERE c.code = i.link_code_1))
    THROW 50000, 'Invalid customer code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_CUSTOMER' AND NOT EXISTS (SELECT 1 FROM customer c WHERE c.code = i.link_code_2))
    THROW 50000, 'Invalid customer code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_EMPLOYEE' AND NOT EXISTS (SELECT 1 FROM employee e WHERE e.code = i.link_code_1))
    THROW 50000, 'Invalid employee code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_EMPLOYEE' AND NOT EXISTS (SELECT 1 FROM employee e WHERE e.code = i.link_code_2))
    THROW 50000, 'Invalid employee code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_ITEM' AND NOT EXISTS (SELECT 1 FROM item t WHERE t.code = i.link_code_1))
    THROW 50000, 'Invalid item code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_ITEM' AND NOT EXISTS (SELECT 1 FROM item t WHERE t.code = i.link_code_2))
    THROW 50000, 'Invalid item code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_MOVEMENT' AND NOT EXISTS (SELECT 1 FROM movement m WHERE m.code = i.link_code_1))
    THROW 50000, 'Invalid movement code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_MOVEMENT' AND NOT EXISTS (SELECT 1 FROM movement m WHERE m.code = i.link_code_2))
    THROW 50000, 'Invalid movement code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PAYMENT' AND NOT EXISTS (SELECT 1 FROM payment p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid payment code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PAYMENT' AND NOT EXISTS (SELECT 1 FROM payment p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid payment code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PLACE' AND NOT EXISTS (SELECT 1 FROM place p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid place code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PLACE' AND NOT EXISTS (SELECT 1 FROM place p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid place code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PRODUCT' AND NOT EXISTS (SELECT 1 FROM product p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid product code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PRODUCT' AND NOT EXISTS (SELECT 1 FROM product p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid product code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_PROJECT' AND NOT EXISTS (SELECT 1 FROM project p WHERE p.code = i.link_code_1))
    THROW 50000, 'Invalid project code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_PROJECT' AND NOT EXISTS (SELECT 1 FROM project p WHERE p.code = i.link_code_2))
    THROW 50000, 'Invalid project code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_TOOL' AND NOT EXISTS (SELECT 1 FROM tool t WHERE t.code = i.link_code_1))
    THROW 50000, 'Invalid tool code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_TOOL' AND NOT EXISTS (SELECT 1 FROM tool t WHERE t.code = i.link_code_2))
    THROW 50000, 'Invalid tool code (link_code_2).', 1

  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_1 = N'LINK_TRANS' AND NOT EXISTS (SELECT 1 FROM trans t WHERE t.code = i.link_code_1))
    THROW 50000, 'Invalid trans code (link_code_1).', 1
  IF EXISTS (SELECT 1 FROM inserted i WHERE i.link_type_2 = N'LINK_TRANS' AND NOT EXISTS (SELECT 1 FROM trans t WHERE t.code = i.link_code_2))
    THROW 50000, 'Invalid trans code (link_code_2).', 1
END;


