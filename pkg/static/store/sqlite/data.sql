INSERT INTO config(code, config_type, data) 
VALUES('setting', 'CONFIG_DATA', json_object(
  'default_bank','PLA0000000000N2', 
  'default_chest', 'PLA0000000000N3', 
  'default_warehouse', 'PLA0000000000N1',
  'default_country', 'EU',
  'default_lang', 'en',
  'default_currency', 'EUR',
  'default_deadline', 8,
  'default_paidtype', 'transfer',
  'default_unit', 'piece',
  'default_taxcode', 'VAT20'
));
INSERT INTO config(code, config_type, data) 
VALUES('orientation', 'CONFIG_DATA', json_object(
  'P','Portrait', 
  'L', 'Landscape'
));
INSERT INTO config(code, config_type, data) 
VALUES('paper_size', 'CONFIG_DATA', json_object(
  'a3','A3', 'a4', 'A4', 'a5', 'A5', 'letter', 'Letter', 'legal', 'Legal'
));

INSERT INTO auth(code, user_name, user_group, disabled, auth_meta) 
VALUES('USR0000000000N1', 'admin', 'GROUP_ADMIN', false, json_object(
  'tags', json_array()
));

INSERT INTO currency(code, currency_meta) 
VALUES('EUR', json_object(
  'description','euro', 'digit', 2, 'cash_round', 0, 'tags', json_array()
));
INSERT INTO currency(code, currency_meta) 
VALUES('USD', json_object(
  'description','dollar', 'digit', 2, 'cash_round', 0, 'tags', json_array()
));

INSERT INTO customer(code, customer_type, customer_name, customer_meta) 
VALUES('CUS0000000000N1', 'CUSTOMER_OWN', 'COMPANY NAME',
  json_object(
    'tax_number', '12345678-1-12', 'account', '', 'tax_free', json('false'), 
    'terms', 0, 'credit_limit', 0, 'discount', 0, 'notes', '', 'inactive', json('false'), 'tags', json_array()
  )
);

INSERT INTO place(code, place_type, place_name, place_meta) 
VALUES('PLA0000000000N1', 'PLACE_WAREHOUSE', 'Warehouse',
  json_object(
    'notes', '', 'inactive', json('false'), 'tags', json_array()
  )
);
INSERT INTO place(code, place_type, place_name, currency_code, place_meta) 
VALUES('PLA0000000000N2', 'PLACE_BANK', 'Bank', 'EUR',
  json_object(
    'notes', '', 'inactive', json('false'), 'tags', json_array()
  )
);
INSERT INTO place(code, place_type, place_name, currency_code, place_meta) 
VALUES('PLA0000000000N3', 'PLACE_CASH', 'Cash', 'EUR',
  json_object(
    'notes', '', 'inactive', json('false'), 'tags', json_array()
  )
);

INSERT INTO tax(code, tax_meta) 
VALUES('VAT00',
  json_object(
    'description','VAT 0%', 'rate_value', 0, 'tags', json_array()
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT05',
  json_object(
    'description','VAT 5%', 'rate_value', 0.05, 'tags', json_array()
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT10',
  json_object(
    'description','VAT 10%', 'rate_value', 0.1, 'tags', json_array()
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT15',
  json_object(
    'description','VAT 15%', 'rate_value', 0.15, 'tags', json_array()
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT20',
  json_object(
    'description','VAT 20%', 'rate_value', 0.2, 'tags', json_array()
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT25',
  json_object(
    'description','VAT 25%', 'rate_value', 0.25, 'tags', json_array()
  )
);