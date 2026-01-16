INSERT INTO config(code, config_type, data) 
VALUES('setting', 'CONFIG_DATA', jsonb_build_object(
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
VALUES('orientation', 'CONFIG_DATA', jsonb_build_object(
  'P','Portrait', 
  'L', 'Landscape'
));
INSERT INTO config(code, config_type, data) 
VALUES('paper_size', 'CONFIG_DATA', jsonb_build_object(
  'a3','A3', 'a4', 'A4', 'a5', 'A5', 'letter', 'Letter', 'legal', 'Legal'
));

INSERT INTO auth(code, user_name, user_group, disabled, auth_meta) 
VALUES('USR0000000000N1', 'admin', 'GROUP_ADMIN'::user_group, false, jsonb_build_object(
  'tags', '[]'::JSONB
));

INSERT INTO currency(code, currency_meta) 
VALUES('EUR', jsonb_build_object(
  'description','euro', 'digit', 2, 'cash_round', 0, 'tags', '[]'::JSONB
));
INSERT INTO currency(code, currency_meta) 
VALUES('USD', jsonb_build_object(
  'description','dollar', 'digit', 2, 'cash_round', 0, 'tags', '[]'::JSONB
));

INSERT INTO customer(code, customer_type, customer_name, customer_meta) 
VALUES('CUS0000000000N1', 'CUSTOMER_OWN', 'COMPANY NAME',
  jsonb_build_object(
    'tax_number', '12345678-1-12', 'account', '', 'tax_free', false, 
    'terms', 0, 'credit_limit', 0, 'discount', 0, 'notes', '', 'inactive', false, 'tags', '[]'::JSONB
  )
);

INSERT INTO place(code, place_type, place_name, place_meta) 
VALUES('PLA0000000000N1', 'PLACE_WAREHOUSE', 'Warehouse',
  jsonb_build_object(
    'notes', '', 'inactive', false, 'tags', '[]'::JSONB
  )
);
INSERT INTO place(code, place_type, place_name, currency_code, place_meta) 
VALUES('PLA0000000000N2', 'PLACE_BANK', 'Bank', 'EUR',
  jsonb_build_object(
    'notes', '', 'inactive', false, 'tags', '[]'::JSONB
  )
);
INSERT INTO place(code, place_type, place_name, currency_code, place_meta) 
VALUES('PLA0000000000N3', 'PLACE_CASH', 'Cash', 'EUR',
  jsonb_build_object(
    'notes', '', 'inactive', false, 'tags', '[]'::JSONB
  )
);

INSERT INTO tax(code, tax_meta) 
VALUES('VAT00',
  jsonb_build_object(
    'description','VAT 0%', 'rate_value', 0, 'tags', '[]'::JSONB
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT05',
  jsonb_build_object(
    'description','VAT 5%', 'rate_value', 0.05, 'tags', '[]'::JSONB
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT10',
  jsonb_build_object(
    'description','VAT 10%', 'rate_value', 0.1, 'tags', '[]'::JSONB
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT15',
  jsonb_build_object(
    'description','VAT 15%', 'rate_value', 0.15, 'tags', '[]'::JSONB
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT20',
  jsonb_build_object(
    'description','VAT 20%', 'rate_value', 0.2, 'tags', '[]'::JSONB
  )
);
INSERT INTO tax(code, tax_meta) 
VALUES('VAT25',
  jsonb_build_object(
    'description','VAT 25%', 'rate_value', 0.25, 'tags', '[]'::JSONB
  )
);