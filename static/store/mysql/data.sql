INSERT INTO config(code, config_type, data) 
VALUES('setting', 'CONFIG_DATA', JSON_OBJECT(
  'default_bank','PLA0000000000N2', 
  'default_chest', 'PLA0000000000N3', 
  'default_warehouse', 'PLA0000000000N1',
  'default_country', 'EU',
  'default_lang', 'en',
  'default_currency', 'EUR',
  'default_deadline', 8,
  'default_paidtype', 'PAID_TRANSFER',
  'default_unit', 'piece',
  'default_taxcode', 'VAT20'
));
INSERT INTO config(code, config_type, data) 
VALUES('orientation', 'CONFIG_DATA', JSON_OBJECT(
  'P','Portrait', 
  'L', 'Landscape'
));
INSERT INTO config(code, config_type, data) 
VALUES('paper_size', 'CONFIG_DATA', JSON_OBJECT(
  'a3','A3', 'a4', 'A4', 'a5', 'A5', 'letter', 'Letter', 'legal', 'Legal'
));

INSERT INTO auth(code, user_name, user_group, disabled, auth_meta, auth_map) 
VALUES('USR0000000000N1', 'admin', 'GROUP_ADMIN', false, JSON_OBJECT(
  'tags', JSON_ARRAY()
), JSON_OBJECT());

INSERT INTO currency(code, currency_meta, currency_map) 
VALUES('EUR', JSON_OBJECT(
  'description','euro', 'digit', 2, 'cash_round', 0, 'tags', JSON_ARRAY()
), JSON_OBJECT());
INSERT INTO currency(code, currency_meta, currency_map) 
VALUES('USD', JSON_OBJECT(
  'description','dollar', 'digit', 2, 'cash_round', 0, 'tags', JSON_ARRAY()
), JSON_OBJECT());

INSERT INTO customer(code, customer_type, customer_name, addresses, contacts, events, customer_meta, customer_map) 
VALUES('CUS0000000000N1', 'CUSTOMER_OWN', 'COMPANY NAME', 
  JSON_ARRAY(), JSON_ARRAY(), JSON_ARRAY(),
  JSON_OBJECT(
    'tax_number', '12345678-1-12', 'account', '', 'tax_free', false, 
    'terms', 0, 'credit_limit', 0, 'discount', 0, 'notes', '', 'inactive', false, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);

INSERT INTO place(code, place_type, place_name, address, contacts, events, place_meta, place_map) 
VALUES('PLA0000000000N1', 'PLACE_WAREHOUSE', 'Warehouse',
 JSON_OBJECT(), JSON_ARRAY(), JSON_ARRAY(),
  JSON_OBJECT(
    'notes', '', 'inactive', false, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO place(code, place_type, place_name, currency_code, address, contacts, events, place_meta, place_map) 
VALUES('PLA0000000000N2', 'PLACE_BANK', 'Bank', 'EUR',
 JSON_OBJECT(), JSON_ARRAY(), JSON_ARRAY(),
  JSON_OBJECT(
    'notes', '', 'inactive', false, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO place(code, place_type, place_name, currency_code, address, contacts, events, place_meta, place_map) 
VALUES('PLA0000000000N3', 'PLACE_CASH', 'Cash', 'EUR',
 JSON_OBJECT(), JSON_ARRAY(), JSON_ARRAY(),
  JSON_OBJECT(
    'notes', '', 'inactive', false, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);

INSERT INTO tax(code, tax_meta, tax_map) 
VALUES('VAT00',
  JSON_OBJECT(
    'description','VAT 0%', 'rate_value', 0, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO tax(code, tax_meta, tax_map ) 
VALUES('VAT05',
  JSON_OBJECT(
    'description','VAT 5%', 'rate_value', 0.05, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO tax(code, tax_meta, tax_map ) 
VALUES('VAT10',
  JSON_OBJECT(
    'description','VAT 10%', 'rate_value', 0.1, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO tax(code, tax_meta, tax_map ) 
VALUES('VAT15',
  JSON_OBJECT(
    'description','VAT 15%', 'rate_value', 0.15, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO tax(code, tax_meta, tax_map ) 
VALUES('VAT20',
  JSON_OBJECT(
    'description','VAT 20%', 'rate_value', 0.2, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);
INSERT INTO tax(code, tax_meta, tax_map   ) 
VALUES('VAT25',
  JSON_OBJECT(
    'description','VAT 25%', 'rate_value', 0.25, 'tags', JSON_ARRAY()
  ), JSON_OBJECT()
);