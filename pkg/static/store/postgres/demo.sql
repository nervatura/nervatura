INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_integer',
  'description', 'Example integer',
  'field_type', 'FIELD_INTEGER',
  'tags',  '[]'::JSONB, 
  'filter',  jsonb_build_array('FILTER_PRODUCT','FILTER_EMPLOYEE')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_number',
  'description', 'Example number',
  'field_type', 'FIELD_NUMBER',
  'tags',  '[]'::JSONB, 
  'filter',  jsonb_build_array('FILTER_CUSTOMER')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_date',
  'description', 'Example date',
  'field_type', 'FIELD_DATE',
  'tags',  '[]'::JSONB, 
  'filter',  jsonb_build_array('FILTER_CUSTOMER')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_color',
  'description', 'Example enum list',
  'field_type', 'FIELD_ENUM',
  'tags',  jsonb_build_array('BLUE', 'YELLOW', 'WHITE', 'BROWN', 'RED'), 
  'filter',  jsonb_build_array('FILTER_CUSTOMER', 'FILTER_PRODUCT', 'FILTER_TOOL')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_customer_reference',
  'description', 'Customer reference',
  'field_type', 'FIELD_CUSTOMER',
  'tags',  '[]'::JSONB, 
  'filter',  '[]'::JSONB
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_product_reference',
  'description', 'Product reference',
  'field_type', 'FIELD_PRODUCT',
  'tags',  '[]'::JSONB, 
  'filter',  jsonb_build_array('FILTER_PRODUCT')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_company_page',
  'description', 'Company page URL',
  'field_type', 'FIELD_URL',
  'tags',  '[]'::JSONB, 
  'filter',  '[]'::JSONB
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_status',
  'description', 'Project status',
  'field_type', 'FIELD_ENUM',
  'tags',  jsonb_build_array('10%', '20%', '30%', '40%', '50%', '60%', '70%', '80%', '90%', '100%'), 
  'filter',  jsonb_build_array('FILTER_PROJECT')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_coordinator',
  'description', 'Project coordinator',
  'field_type', 'FIELD_EMPLOYEE',
  'tags',  '[]'::JSONB, 
  'filter',  jsonb_build_array('FILTER_PROJECT')
));
INSERT INTO config(config_type, data) 
VALUES('CONFIG_MAP', jsonb_build_object(
  'field_name','demo_car_no',
  'description', 'Vehicle id.No.',
  'field_type', 'FIELD_STRING',
  'tags',  '[]'::JSONB, 
  'filter',  jsonb_build_array('FILTER_TOOL')
));

INSERT INTO customer(code, customer_type, customer_name, addresses, contacts, events, customer_meta, customer_map) 
VALUES('CUS0000000000N2', 'CUSTOMER_COMPANY', 'First Customer Co.',
jsonb_build_array(
  jsonb_build_object(
    'country', 'Country1', 'state', 'state01', 'zip_code', '1234', 'city', 'City1', 
    'street', 'street 1.', 'notes', 'address of registered office', 
    'tags', '[]'::JSONB, 'address_map', '{}'::JSONB
  ),
  jsonb_build_object(
    'country', 'Country1', 'state', 'state02', 'zip_code', '2345', 'city', 'City2', 
    'street', 'street 2.', 'notes', 'postal address', 
    'tags', '[]'::JSONB, 'address_map', '{}'::JSONB
  )  
),
jsonb_build_array(
  jsonb_build_object(
    'first_name', 'Big', 'surname', 'Man', 'status', 'manager', 
    'phone', '', 'mobile', '', 'email', 'man.big@company.co', 'notes', '', 
    'tags', '[]'::JSONB, 'contact_map', '{}'::JSONB
  ),
  jsonb_build_object(
    'first_name', 'Sales', 'surname', 'Man', 'status', 'sales', 
    'phone', '', 'mobile', '', 'email', 'man.sales@company.co', 'notes', '', 
    'tags', '[]'::JSONB, 'contact_map', '{}'::JSONB
  )  
),
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'First visit', 'start_time', '2024-04-05T08:00:00', 'end_time', '2024-04-05T10:00:00',
    'place', 'City1', 'description', 'It was long ...  ,-(', 'tags', jsonb_build_array('VISIT'),
    'event_map', jsonb_build_object('demo_company_page', 'nervatura.com')
  ),
  jsonb_build_object(
    'uid', '', 'subject', 'Second visit', 'start_time', '2024-04-06T08:00:00', 'end_time', '2024-04-06T10:00:00',
    'place', 'City1', 'description', '', 'tags', jsonb_build_array('VISIT'), 'event_map', '{}'::JSONB
  )  
), 
jsonb_build_object(
  'tax_number', '87654321-1-12', 'account', '', 'tax_free', false, 
  'terms', 8, 'credit_limit', 1000000, 'discount', 2, 'inactive', false, 'notes', '', 
  'tags', '[]'::JSONB
), 
jsonb_build_object(
  'demo_number', 123.4, 'demo_date', '2024-08-12'
));

INSERT INTO customer(code, customer_type, customer_name, addresses, contacts, events, customer_meta, customer_map) 
VALUES('CUS0000000000N3', 'CUSTOMER_PRIVATE', 'Second Customer Name',
jsonb_build_array(
  jsonb_build_object(
    'country', 'Country1', 'state', 'state03', 'zip_code', '6789', 'city', 'City3', 
    'street', 'street 3.', 'notes', '', 'tags', '[]'::JSONB, 'address_map', '{}'::JSONB
  ) 
),
jsonb_build_array(
  jsonb_build_object(
    'first_name', 'Jack', 'surname', 'Frost', 'tags', '[]'::JSONB, 'contact_map', '{}'::JSONB
  ) 
),
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'Training', 'start_time', '2024-04-07T08:00:00', 'end_time', '2024-04-07T10:00:00',
    'place', '', 'description', '', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  )  
),
jsonb_build_object(
  'tax_number', '12121212-1-12', 'account', '', 'tax_free', false, 
  'terms', 1, 'credit_limit', 0, 'discount', 6, 'inactive', false, 'notes', '', 
  'tags', '[]'::JSONB
), 
jsonb_build_object(
  'demo_number', 56789.67, 'demo_date', '2024-09-01', 
  'demo_color', 'YELLOW', 'demo_customer_reference', 'CUS0000000000N2'
));

INSERT INTO customer(code, customer_type, customer_name, addresses, contacts, events, customer_meta, customer_map) 
VALUES('CUS0000000000N4', 'CUSTOMER_OTHER', 'Third Customer Foundation',
jsonb_build_array(
  jsonb_build_object(
    'country', 'Country2', 'state', 'state04', 'zip_code', '6543', 'city', 'City4', 
    'street', 'street 4.', 'notes', '', 'tags', '[]'::JSONB, 'address_map', '{}'::JSONB
  ) 
),
jsonb_build_array(
  jsonb_build_object(
    'first_name', 'Harry', 'surname', 'Potter', 'status', '', 
    'phone', '', 'mobile', '', 'email', '', 'notes', '', 
    'tags', '[]'::JSONB, 'contact_map', '{}'::JSONB
  ) 
),
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'Training', 'start_time', '2024-04-07T08:00:00', 'end_time', '2024-04-07T10:00:00',
    'place', '', 'description', '', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  )  
),
jsonb_build_object(
  'tax_number', '23232323-1-12', 'account', '', 'tax_free', true, 
  'terms', 4, 'credit_limit', 0, 'discount', 6, 'inactive', false, 'notes', '', 
  'tags', '[]'::JSONB
), 
jsonb_build_object(
  'demo_color', 'BROWN'
));

INSERT INTO employee(code, address, contact, events, employee_meta, employee_map) 
VALUES('EMP0000000000N1',
jsonb_build_object(
  'country', 'Country', 'state', 'state', 'zip_code', '6543', 'city', 'City', 
  'street', 'street.', 'notes', '', 'tags', '[]'::JSONB, 'address_map', '{}'::JSONB
),
jsonb_build_object(
  'first_name', 'John', 'surname', 'Strong', 'status', 'heaver', 
  'phone', '', 'mobile', '', 'email', '', 'notes', 'He is a good man ...', 
  'tags', '[]'::JSONB, 'contact_map', '{}'::JSONB
),
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'Holiday', 'start_time', '2024-12-15T00:00:00', 'end_time', '2024-12-31T00:00:00',
    'place', 'On the beach', 'description', '', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  )  
),
jsonb_build_object(
  'start_date', '2020-12-01', 'end_date', '', 'inactive', false, 'notes', '',
  'tags', jsonb_build_array('PRODUCTION')
), 
jsonb_build_object(
  'demo_integer', 42
));

INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N1', 'PRODUCT_ITEM', 'Big product', 'VAT20',
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'New prices', 'start_time', '2024-04-05T08:00:00', 'end_time', '2024-04-05T15:00:00',
    'place', '', 'description', '', 'tags', jsonb_build_array('PRICING'), 'event_map', '{}'::JSONB
  )  
),
jsonb_build_object(
  'unit', 'piece', 'inactive', false, 'notes', '', 
  'barcode_type', 'BARCODE_CODE_39', 'barcode', 'BC0123456789', 'barcode_qty', 1,
  'tags', jsonb_build_array('WEBITEM')
), 
jsonb_build_object(
  'demo_color', 'RED'
));
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N2', 'PRODUCT_SERVICE', 'Good work', 'VAT20',
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'business trip', 'start_time', '2024-04-08T08:00:00', 'end_time', '2024-04-12T18:00:00',
    'place', 'Hawaii', 'description', '', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  ),
  jsonb_build_object(
    'uid', '', 'subject', 'Inventory', 'start_time', '2024-04-12T08:00:00', 'end_time', '',
    'place', '', 'description', 'Inventory check in the warehouse', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  ) 
),
jsonb_build_object(
  'unit', 'hour', 'inactive', false, 'notes', '', 
  'barcode_type', 'BARCODE_CODE_39', 'barcode', 'BC1212121212', 'barcode_qty', 5,
  'tags', jsonb_build_array()
), 
jsonb_build_object(
  'demo_integer', 100000, 'demo_product_reference', 'PRD0000000000N12'
));
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N3', 'PRODUCT_ITEM', 'Nice product', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false, 'notes', '', 
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'tags', jsonb_build_array('WEBITEM')
), 
jsonb_build_object(
  'demo_color', 'WHITE'
));
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N4', 'PRODUCT_ITEM', 'Car', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false, 
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Manufacturing products', 
  'tags', jsonb_build_array('WEBITEM')
), 
'{}'::JSONB);
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N5', 'PRODUCT_ITEM', 'Wheel', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false, 
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Raw material, component', 
  'tags', jsonb_build_array('WEBITEM', 'COMPONENT')
), 
'{}'::JSONB);
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N6', 'PRODUCT_ITEM', 'Door', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Raw material, component', 
  'tags', jsonb_build_array('WEBITEM', 'COMPONENT')
), 
'{}'::JSONB);
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N7', 'PRODUCT_ITEM', 'Paint', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'liter', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Raw material, component', 
  'tags', jsonb_build_array('WEBITEM', 'COMPONENT')
), 
'{}'::JSONB);
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N8', 'PRODUCT_ITEM', 'Pallet', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Raw material, component (not share sample)', 
  'tags', jsonb_build_array('WEBITEM', 'COMPONENT')
), 
'{}'::JSONB);
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N9', 'PRODUCT_ITEM', 'Basket', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Souvenir component', 
  'tags', jsonb_build_array('COMPONENT')
), 
jsonb_build_object());
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N10', 'PRODUCT_ITEM', 'Wine', 'VAT05',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Souvenir component', 
  'tags', jsonb_build_array('COMPONENT')
), 
jsonb_build_object());
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N11', 'PRODUCT_ITEM', 'Chocolate', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'Souvenir component', 
  'tags', jsonb_build_array('COMPONENT')
), 
jsonb_build_object());
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N12', 'PRODUCT_VIRTUAL', 'Souvenir - virtual product', 'VAT15',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'A technical package, which might include several real products or services', 
  'tags', jsonb_build_array('WEBITEM')
), 
'{}'::JSONB);
INSERT INTO product(code, product_type, product_name, tax_code, events, product_meta, product_map) 
VALUES('PRD0000000000N13', 'PRODUCT_ITEM', 'Phone', 'VAT20',
'[]'::JSONB,
jsonb_build_object(
  'unit', 'piece', 'inactive', false,
  'barcode_type', 'BARCODE_CODE_39', 'barcode', '', 'barcode_qty', 0,
  'notes', 'for tool movement...', 
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PRODUCT', 'PRD0000000000N12', 'LINK_PRODUCT', 'PRD0000000000N9',
jsonb_build_object(
  'amount', 0, 'qty', 1, 'rate', 0, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PRODUCT', 'PRD0000000000N12', 'LINK_PRODUCT', 'PRD0000000000N10',
jsonb_build_object(
  'amount', 0, 'qty', 1, 'rate', 0, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PRODUCT', 'PRD0000000000N12', 'LINK_PRODUCT', 'PRD0000000000N11',
jsonb_build_object(
  'amount', 0, 'qty', 1, 'rate', 0, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PRODUCT', 'PRD0000000000N12', 'LINK_PRODUCT', 'PRD0000000000N2',
jsonb_build_object(
  'amount', 0, 'qty', 1, 'rate', 0, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO price(price_type, valid_from, product_code, currency_code, qty, price_meta, price_map) 
VALUES('PRICE_CUSTOMER', '2024-04-05', 'PRD0000000000N1', 'EUR', 0,
jsonb_build_object(
  'price_value', 25, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO price(price_type, valid_from, product_code, currency_code, qty, price_meta, price_map) 
VALUES('PRICE_CUSTOMER', '2024-04-05', 'PRD0000000000N1', 'EUR', 10,
jsonb_build_object(
  'price_value', 20, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO project(code, project_name, customer_code, addresses, contacts, events, project_meta, project_map) 
VALUES('PRJ0000000000N1', 'Sample project', 'CUS0000000000N2',
jsonb_build_array(
  jsonb_build_object(
    'country', 'Country1', 'state', '', 'zip_code', '02230', 'city', 'City1', 
    'street', 'Address 3. AB. 5..', 'notes', '', 'tags', '[]'::JSONB, 'address_map', '{}'::JSONB
  ) 
),
jsonb_build_array(
  jsonb_build_object(
    'first_name', 'Big', 'surname', 'Man'
  ) 
),
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'Project meeting', 'start_time', '2024-12-10T09:00:00', 'end_time', '2024-12-10T11:00:00',
    'place', 'Office', 'description', '', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  )  
),
jsonb_build_object(
  'start_date', '2024-12-01', 'end_date', '',
  'inactive', false, 'notes', '', 'tags', '[]'::JSONB
), 
jsonb_build_object(
  'demo_status', '20%', 'demo_coordinator', 'EMP0000000000N1'
));

INSERT INTO tool(code, description, product_code, events, tool_meta, tool_map) 
VALUES('SER0000000000N1', 'Company car 1.', 'PRD0000000000N4',
jsonb_build_array(
  jsonb_build_object(
    'uid', '', 'subject', 'Technical inspection', 'start_time', '2024-04-05T08:00:00', 'end_time', '2024-04-05T15:00:00',
    'place', 'Service', 'description', '', 'tags', '[]'::JSONB, 'event_map', '{}'::JSONB
  )  
),
jsonb_build_object(
  'serial_number', 'ABC-123', 'inactive', false, 'notes', '', 'tags', jsonb_build_array('CAR')
), 
jsonb_build_object(
  'demo_car_no', 'VIN12345678', 'demo_color', 'RED'
));
INSERT INTO tool(code, description, product_code, events, tool_meta, tool_map) 
VALUES('SER0000000000N2', 'Company car 2.', 'PRD0000000000N4',
'[]'::JSONB,
jsonb_build_object(
  'serial_number', 'DEF-456', 'inactive', false, 'notes', '', 'tags', jsonb_build_array('CAR')
), 
jsonb_build_object(
  'demo_car_no', 'VIN87654321', 'demo_color', 'BLUE'
));
INSERT INTO tool(code, description, product_code, events, tool_meta, tool_map) 
VALUES('SER0000000000N3', 'Motorola', 'PRD0000000000N13',
'[]'::JSONB,
jsonb_build_object(
  'serial_number', 'IMEI-023456789', 'inactive', false, 'notes', '', 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO place(code, place_type, place_name, place_meta) 
VALUES('PLA0000000000N4', 'PLACE_WAREHOUSE', 'Raw material', 
  jsonb_build_object(
    'notes', '', 'inactive', false, 'tags', '[]'::JSONB
  )
);

INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('OFF0000000000N4', 'TRANS_OFFER', 'DIRECTION_OUT', '2024-11-05', NULL, 'CUS0000000000N2', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-11-30T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'DEMO invoice offer', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('SALES')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('ORD0000000000N1', 'TRANS_ORDER', 'DIRECTION_IN', '2024-11-01', NULL, 'CUS0000000000N4', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-11-10T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'We bought some basic materials for production and sale. It was invoiced on the basis of delivery, but not all were delivered yet.', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('LOGISTICS')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('ORD0000000000N2', 'TRANS_ORDER', 'DIRECTION_OUT', '2024-12-04', NULL, 'CUS0000000000N3', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-10T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'Virtual product sample.', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('SALES')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('ORD0000000000N3', 'TRANS_ORDER', 'DIRECTION_OUT', '2024-12-10', 'OFF0000000000N4', 'CUS0000000000N2', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-20T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'DEMO invoice order.', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('SALES')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('INV0000000000N5', 'TRANS_INVOICE', 'DIRECTION_OUT', '2024-12-10', 'ORD0000000000N3', 'CUS0000000000N2', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-20T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', '', 'internal_notes', '', 
  'report_notes', 'A long and <b><i>rich text</b></i> at the bottom of the invoice...<br><br>Can be multiple lines ...',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('SALES')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('INV0000000000N6', 'TRANS_INVOICE', 'DIRECTION_OUT', '2024-12-10', 'ORD0000000000N2', 'CUS0000000000N3', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-28T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'Virtual product sample.', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('SALES')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('INV0000000000N7', 'TRANS_INVOICE', 'DIRECTION_IN', '2024-11-10', 'ORD0000000000N1', 'CUS0000000000N4', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-20T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'We bought some basic materials for production and sale. It was invoiced on the basis of delivery, but not all were delivered yet.', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('LOGISTICS')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('WOR0000000000N8', 'TRANS_WORKSHEET', 'DIRECTION_OUT', '2024-12-05', NULL, 'CUS0000000000N2', 'EMP0000000000N1', NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-05T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', '',  'internal_notes', '', 'report_notes', '',
  'worksheet', jsonb_build_object(
    'distance', 200, 'repair', 0, 'total', 3, 'justification', ''
  ), 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('SALES')
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('REN0000000000N9', 'TRANS_RENT', 'DIRECTION_OUT', '2024-11-01', NULL, 'CUS0000000000N2', NULL, NULL, NULL, 'EUR', 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-11-30T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 'notes', '', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', jsonb_build_object(
    'holiday', 3, 'bad_tool', 0, 'other', 0, 'justification', ''
  ), 'invoice', '{}'::JSONB,
  'tags', jsonb_build_array('LOGISTICS')
), 
'{}'::JSONB);

INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('BAN0000000000N10', 'TRANS_BANK', 'DIRECTION_TRANSFER', '2024-12-15', NULL, NULL, NULL, NULL, 'PLA0000000000N2', NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', 'BM0123456', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 'notes', '', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('CAS0000000000N11', 'TRANS_CASH', 'DIRECTION_OUT', '2024-12-18', NULL, NULL, NULL, NULL, 'PLA0000000000N3', NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 'notes', '', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('WAY0000000000N12', 'TRANS_WAYBILL', 'DIRECTION_OUT', '2024-12-05', NULL, NULL, 'EMP0000000000N1', NULL, NULL, NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'We hand out some working tools to the employee...', 
  'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('DEL0000000000N13', 'TRANS_DELIVERY', 'DIRECTION_IN', '2024-11-08', NULL, NULL, NULL, NULL, NULL, NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', '', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('DEL0000000000N14', 'TRANS_DELIVERY', 'DIRECTION_OUT', '2024-12-10', NULL, NULL, NULL, NULL, NULL, NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', '', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('DEL0000000000N15', 'TRANS_DELIVERY', 'DIRECTION_OUT', '2024-12-10', NULL, NULL, NULL, NULL, NULL, NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', '', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('TRF0000000000N16', 'TRANS_DELIVERY', 'DIRECTION_TRANSFER', '2024-11-08', NULL, NULL, NULL, NULL, 'PLA0000000000N4', NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', '', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('COR0000000000N17', 'TRANS_INVENTORY', 'DIRECTION_TRANSFER', '2024-12-01', NULL, NULL, NULL, NULL, 'PLA0000000000N1', NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'Scrapping of some products ...', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('FOR0000000000N18', 'TRANS_FORMULA', 'DIRECTION_TRANSFER', '2024-12-01', NULL, NULL, NULL, NULL, NULL, NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'Sample formula (4 door/car)', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('FOR0000000000N19', 'TRANS_FORMULA', 'DIRECTION_TRANSFER', '2024-12-01', NULL, NULL, NULL, NULL, NULL, NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'Sample formula (3 door/car)', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO trans(
  code, trans_type, direction, trans_date, trans_code, customer_code, employee_code, project_code, 
  place_code, currency_code, auth_code, trans_meta, trans_map) 
VALUES('PRO0000000000N20', 'TRANS_PRODUCTION', 'DIRECTION_TRANSFER', '2024-12-01', NULL, NULL, NULL, NULL, 'PLA0000000000N1', NULL, 'USR0000000000N1',
jsonb_build_object(
  'due_time', '2024-12-02T00:00:00', 'ref_number', '', 'paid_type', 'PAID_TRANSFER', 
  'tax_free', false, 'paid', false, 'rate', 0, 'closed', false, 
  'status', 'STATUS_NORMAL', 'trans_state', 'STATE_OK', 
  'notes', 'formula: 4 door/car', 'internal_notes', '', 'report_notes', '',
  'worksheet', '{}'::JSONB, 'rent', '{}'::JSONB, 'invoice', '{}'::JSONB,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N1', 'ORD0000000000N1', 'PRD0000000000N5', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 40, 'fx_price', 10, 'net_amount', 400, 'discount', 0, 'vat_amount', 80, 'amount', 480,
  'description', 'Wheel', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N2', 'ORD0000000000N1', 'PRD0000000000N6', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 60, 'fx_price', 12, 'net_amount', 720, 'discount', 0, 'vat_amount', 144, 'amount', 864,
  'description', 'Door', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N3', 'ORD0000000000N1', 'PRD0000000000N7', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 50, 'fx_price', 16, 'net_amount', 800, 'discount', 0, 'vat_amount', 160, 'amount', 960,
  'description', 'Paint', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N4', 'ORD0000000000N1', 'PRD0000000000N8', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 20, 'fx_price', 5, 'net_amount', 100, 'discount', 0, 'vat_amount', 20, 'amount', 120,
  'description', 'Pallet', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N5', 'ORD0000000000N1', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 10, 'fx_price', 120, 'net_amount', 1200, 'discount', 0, 'vat_amount', 240, 'amount', 1440,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N6', 'ORD0000000000N1', 'PRD0000000000N3', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 10, 'fx_price', 15, 'net_amount', 150, 'discount', 0, 'vat_amount', 30, 'amount', 180,
  'description', 'Nice product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N7', 'ORD0000000000N1', 'PRD0000000000N9', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 20, 'fx_price', 8, 'net_amount', 160, 'discount', 0, 'vat_amount', 32, 'amount', 192,
  'description', 'Basket', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N8', 'ORD0000000000N1', 'PRD0000000000N10', 'VAT05',
jsonb_build_object(
  'unit', 'piece', 'qty', 20, 'fx_price', 20, 'net_amount', 400, 'discount', 0, 'vat_amount', 20, 'amount', 420,
  'description', 'Wine', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N9', 'ORD0000000000N1', 'PRD0000000000N11', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 20, 'fx_price', 16, 'net_amount', 320, 'discount', 0, 'vat_amount', 64, 'amount', 384,
  'description', 'Chocolate', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N10', 'ORD0000000000N2', 'PRD0000000000N12', 'VAT15',
jsonb_build_object(
  'unit', 'piece', 'qty', 2, 'fx_price', 60, 'net_amount', 120, 'discount', 0, 'vat_amount', 18, 'amount', 138,
  'description', 'Souvenir - virtual product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N11', 'ORD0000000000N2', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 3, 'fx_price', 25, 'net_amount', 75, 'discount', 0, 'vat_amount', 15, 'amount', 90,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N12', 'ORD0000000000N3', 'PRD0000000000N2', 'VAT20',
jsonb_build_object(
  'unit', 'hour', 'qty', 1, 'fx_price', 120, 'net_amount', 120, 'discount', 0, 'vat_amount', 24, 'amount', 144,
  'description', 'Very good work!', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N13', 'ORD0000000000N3', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 3, 'fx_price', 166.67, 'net_amount', 500, 'discount', 0, 'vat_amount', 100, 'amount', 600,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N14', 'ORD0000000000N3', 'PRD0000000000N3', 'VAT05',
jsonb_build_object(
  'unit', 'piece', 'qty', 5, 'fx_price', 20, 'net_amount', 100, 'discount', 0, 'vat_amount', 5, 'amount', 105,
  'description', 'Nice product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N15', 'OFF0000000000N4', 'PRD0000000000N2', 'VAT20',
jsonb_build_object(
  'unit', 'hour', 'qty', 1, 'fx_price', 120, 'net_amount', 120, 'discount', 0, 'vat_amount', 24, 'amount', 144,
  'description', 'Very good work!', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N16', 'OFF0000000000N4', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 3, 'fx_price', 166.67, 'net_amount', 500, 'discount', 0, 'vat_amount', 100, 'amount', 600,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N17', 'OFF0000000000N4', 'PRD0000000000N3', 'VAT05',
jsonb_build_object(
  'unit', 'piece', 'qty', 5, 'fx_price', 20, 'net_amount', 100, 'discount', 0, 'vat_amount', 5, 'amount', 105,
  'description', 'Nice product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N18', 'INV0000000000N5', 'PRD0000000000N2', 'VAT20',
jsonb_build_object(
  'unit', 'hour', 'qty', 1, 'fx_price', 120, 'net_amount', 120, 'discount', 0, 'vat_amount', 24, 'amount', 144,
  'description', 'Very good work!', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N19', 'INV0000000000N5', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 3, 'fx_price', 166.67, 'net_amount', 500, 'discount', 0, 'vat_amount', 100, 'amount', 600,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N20', 'INV0000000000N5', 'PRD0000000000N3', 'VAT05',
jsonb_build_object(
  'unit', 'piece', 'qty', 5, 'fx_price', 20, 'net_amount', 100, 'discount', 0, 'vat_amount', 5, 'amount', 105,
  'description', 'Nice product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N21', 'INV0000000000N6', 'PRD0000000000N12', 'VAT15',
jsonb_build_object(
  'unit', 'piece', 'qty', 2, 'fx_price', 60, 'net_amount', 120, 'discount', 0, 'vat_amount', 18, 'amount', 138,
  'description', 'Souvenir - virtual product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N22', 'INV0000000000N6', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 3, 'fx_price', 25, 'net_amount', 75, 'discount', 0, 'vat_amount', 15, 'amount', 90,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N23', 'INV0000000000N7', 'PRD0000000000N5', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 30, 'fx_price', 10, 'net_amount', 300, 'discount', 0, 'vat_amount', 60, 'amount', 360,
  'description', 'Wheel', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N24', 'INV0000000000N7', 'PRD0000000000N6', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 50, 'fx_price', 12, 'net_amount', 600, 'discount', 0, 'vat_amount', 120, 'amount', 720,
  'description', 'Door', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N25', 'INV0000000000N7', 'PRD0000000000N7', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 50, 'fx_price', 16, 'net_amount', 800, 'discount', 0, 'vat_amount', 160, 'amount', 960,
  'description', 'Paint', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N26', 'INV0000000000N7', 'PRD0000000000N8', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 15, 'fx_price', 5, 'net_amount', 75, 'discount', 0, 'vat_amount', 15, 'amount', 90,
  'description', 'Pallet', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N27', 'INV0000000000N7', 'PRD0000000000N1', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 10, 'fx_price', 120, 'net_amount', 1200, 'discount', 0, 'vat_amount', 240, 'amount', 1440,
  'description', 'Big product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N28', 'INV0000000000N7', 'PRD0000000000N3', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 10, 'fx_price', 15, 'net_amount', 150, 'discount', 0, 'vat_amount', 30, 'amount', 180,
  'description', 'Nice product', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N29', 'INV0000000000N7', 'PRD0000000000N9', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 15, 'fx_price', 8, 'net_amount', 120, 'discount', 0, 'vat_amount', 24, 'amount', 144,
  'description', 'Basket', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N30', 'INV0000000000N7', 'PRD0000000000N10', 'VAT05',
jsonb_build_object(
  'unit', 'piece', 'qty', 10, 'fx_price', 20, 'net_amount', 200, 'discount', 0, 'vat_amount', 10, 'amount', 210,
  'description', 'Wine', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N31', 'INV0000000000N7', 'PRD0000000000N11', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 20, 'fx_price', 16, 'net_amount', 320, 'discount', 0, 'vat_amount', 64, 'amount', 384,
  'description', 'Chocolate', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N32', 'WOR0000000000N8', 'PRD0000000000N2', 'VAT20',
jsonb_build_object(
  'unit', 'hour', 'qty', 2, 'fx_price', 130, 'net_amount', 260, 'discount', 0, 'vat_amount', 52, 'amount', 312,
  'description', 'Good work', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO item(
  code, trans_code, product_code, tax_code, item_meta, item_map) 
VALUES('ITM0000000000N33', 'REN0000000000N9', 'PRD0000000000N8', 'VAT20',
jsonb_build_object(
  'unit', 'piece', 'qty', 3, 'fx_price', 12, 'net_amount', 396, 'discount', 0, 'vat_amount', 79.2, 'amount', 475.2,
  'description', 'Pallet', 'deposit', false, 'own_stock', 0, 'action_price', false,
  'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO payment(
  code, paid_date, trans_code, payment_meta, payment_map) 
VALUES('PMT0000000000N1', '2024-12-20', 'BAN0000000000N10',
jsonb_build_object(
  'amount', -4000, 'notes', 'payment two divided...',
  'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO payment(
  code, paid_date, trans_code, payment_meta, payment_map) 
VALUES('PMT0000000000N2', '2024-12-20', 'BAN0000000000N10',
jsonb_build_object(
  'amount', 849, 'notes', '', 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO payment(
  code, paid_date, trans_code, payment_meta, payment_map) 
VALUES('PMT0000000000N3', '2024-12-28', 'BAN0000000000N10',
jsonb_build_object(
  'amount', 228, 'notes', '', 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO payment(
  code, paid_date, trans_code, payment_meta, payment_map) 
VALUES('PMT0000000000N4', '2024-12-18', 'CAS0000000000N11',
jsonb_build_object(
  'amount', -488, 'notes', '', 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PAYMENT', 'PMT0000000000N1', 'LINK_TRANS', 'INV0000000000N7',
jsonb_build_object(
  'amount', 4000, 'qty', 0, 'rate', 1, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PAYMENT', 'PMT0000000000N2', 'LINK_TRANS', 'INV0000000000N5',
jsonb_build_object(
  'amount', 849, 'qty', 0, 'rate', 1, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PAYMENT', 'PMT0000000000N3', 'LINK_TRANS', 'INV0000000000N6',
jsonb_build_object(
  'amount', 228, 'qty', 0, 'rate', 1, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO link(
  link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map) 
VALUES('LINK_PAYMENT', 'PMT0000000000N4', 'LINK_TRANS', 'INV0000000000N7',
jsonb_build_object(
  'amount', 488, 'qty', 0, 'rate', 1, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N1', 'MOVEMENT_TOOL', '2024-12-05T00:00:00', 'WAY0000000000N12', NULL, 'SER0000000000N1', NULL, NULL, NULL,
jsonb_build_object(
  'qty', 0, 'notes', '', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N2', 'MOVEMENT_TOOL', '2024-12-05T00:00:00', 'WAY0000000000N12', NULL, 'SER0000000000N3', NULL, NULL, NULL,
jsonb_build_object(
  'qty', 0, 'notes', 'mobile phone', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N3', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N5', NULL, 'PLA0000000000N4', 'ITM0000000000N1', NULL,
jsonb_build_object(
  'qty', 30, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N4', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N6', NULL, 'PLA0000000000N4', 'ITM0000000000N2', NULL,
jsonb_build_object(
  'qty', 50, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N5', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N7', NULL, 'PLA0000000000N4', 'ITM0000000000N3', NULL,
jsonb_build_object(
  'qty', 50, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N6', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N8', NULL, 'PLA0000000000N4', 'ITM0000000000N4', NULL,
jsonb_build_object(
  'qty', 15, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N7', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N1', NULL, 'PLA0000000000N4', 'ITM0000000000N5', NULL,
jsonb_build_object(
  'qty', 10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N8', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N3', NULL, 'PLA0000000000N4', 'ITM0000000000N6', NULL,
jsonb_build_object(
  'qty', 10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N9', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N9', NULL, 'PLA0000000000N4', 'ITM0000000000N7', NULL,
jsonb_build_object(
  'qty', 15, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N10', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N10', NULL, 'PLA0000000000N4', 'ITM0000000000N8', NULL,
jsonb_build_object(
  'qty', 10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N11', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'DEL0000000000N13', 'PRD0000000000N11', NULL, 'PLA0000000000N4', 'ITM0000000000N9', NULL,
jsonb_build_object(
  'qty', 20, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N12', 'MOVEMENT_INVENTORY', '2024-12-10T00:00:00', 'DEL0000000000N14', 'PRD0000000000N9', NULL, 'PLA0000000000N1', 'ITM0000000000N10', NULL,
jsonb_build_object(
  'qty', -2, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N13', 'MOVEMENT_INVENTORY', '2024-12-10T00:00:00', 'DEL0000000000N14', 'PRD0000000000N10', NULL, 'PLA0000000000N1', 'ITM0000000000N10', NULL,
jsonb_build_object(
  'qty', -2, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N14', 'MOVEMENT_INVENTORY', '2024-12-10T00:00:00', 'DEL0000000000N14', 'PRD0000000000N11', NULL, 'PLA0000000000N1', 'ITM0000000000N10', NULL,
jsonb_build_object(
  'qty', -4, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N15', 'MOVEMENT_INVENTORY', '2024-12-10T00:00:00', 'DEL0000000000N14', 'PRD0000000000N1', NULL, 'PLA0000000000N1', 'ITM0000000000N11', NULL,
jsonb_build_object(
  'qty', -3, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N16', 'MOVEMENT_INVENTORY', '2024-12-10T00:00:00', 'DEL0000000000N15', 'PRD0000000000N1', NULL, 'PLA0000000000N1', 'ITM0000000000N13', NULL,
jsonb_build_object(
  'qty', -3, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N17', 'MOVEMENT_INVENTORY', '2024-12-10T00:00:00', 'DEL0000000000N15', 'PRD0000000000N3', NULL, 'PLA0000000000N1', 'ITM0000000000N14', NULL,
jsonb_build_object(
  'qty', -5, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N18', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N1', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N19', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N1', NULL, 'PLA0000000000N1', NULL, 'MOV0000000000N18',
jsonb_build_object(
  'qty', 10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N20', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N3', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N21', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N3', NULL, 'PLA0000000000N1', NULL, 'MOV0000000000N20',
jsonb_build_object(
  'qty', 10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N22', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N9', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -15, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N23', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N9', NULL, 'PLA0000000000N1', NULL, 'MOV0000000000N22',
jsonb_build_object(
  'qty', 15, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N24', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N10', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N25', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N10', NULL, 'PLA0000000000N1', NULL, 'MOV0000000000N24',
jsonb_build_object(
  'qty', 10, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N26', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N11', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -20, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N27', 'MOVEMENT_INVENTORY', '2024-11-08T00:00:00', 'TRF0000000000N16', 'PRD0000000000N11', NULL, 'PLA0000000000N1', NULL, 'MOV0000000000N26',
jsonb_build_object(
  'qty', 20, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N28', 'MOVEMENT_INVENTORY', '2024-12-01T00:00:00', 'COR0000000000N17', 'PRD0000000000N1', NULL, 'PLA0000000000N1', NULL, NULL,
jsonb_build_object(
  'qty', -2, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N29', 'MOVEMENT_INVENTORY', '2024-12-01T00:00:00', 'COR0000000000N17', 'PRD0000000000N10', NULL, 'PLA0000000000N1', NULL, NULL,
jsonb_build_object(
  'qty', -3, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N30', 'MOVEMENT_HEAD', '2024-12-01T00:00:00', 'FOR0000000000N18', 'PRD0000000000N4', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 5, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N31', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N18', 'PRD0000000000N5', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 20, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N32', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N18', 'PRD0000000000N6', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 20, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N33', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N18', 'PRD0000000000N7', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 30, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N34', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N18', 'PRD0000000000N8', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 1, 'notes', 'demo', 'shared', true, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N35', 'MOVEMENT_HEAD', '2024-12-01T00:00:00', 'FOR0000000000N19', 'PRD0000000000N4', NULL, NULL, NULL, NULL,
jsonb_build_object(
  'qty', 5, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N36', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N19', 'PRD0000000000N5', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 20, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N37', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N19', 'PRD0000000000N6', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 15, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N38', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N19', 'PRD0000000000N7', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 28, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N39', 'MOVEMENT_PLAN', '2024-12-01T00:00:00', 'FOR0000000000N19', 'PRD0000000000N8', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', 1, 'notes', 'demo', 'shared', true, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N40', 'MOVEMENT_INVENTORY', '2024-12-02T00:00:00', 'PRO0000000000N20', 'PRD0000000000N4', NULL, 'PLA0000000000N1', NULL, NULL,
jsonb_build_object(
  'qty', 2, 'notes', 'demo', 'shared', true, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N41', 'MOVEMENT_INVENTORY', '2024-12-01T00:00:00', 'PRO0000000000N20', 'PRD0000000000N5', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -8, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N42', 'MOVEMENT_INVENTORY', '2024-12-01T00:00:00', 'PRO0000000000N20', 'PRD0000000000N6', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -8, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N43', 'MOVEMENT_INVENTORY', '2024-12-01T00:00:00', 'PRO0000000000N20', 'PRD0000000000N7', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -12, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);
INSERT INTO movement(
  code, movement_type, shipping_time, trans_code, product_code, tool_code, place_code, item_code, movement_code, movement_meta, movement_map) 
VALUES('MOV0000000000N44', 'MOVEMENT_INVENTORY', '2024-12-01T00:00:00', 'PRO0000000000N20', 'PRD0000000000N8', NULL, 'PLA0000000000N4', NULL, NULL,
jsonb_build_object(
  'qty', -1, 'notes', 'demo', 'shared', false, 'tags', '[]'::JSONB
), 
'{}'::JSONB);

INSERT INTO config(config_type, data) 
VALUES('CONFIG_SHORTCUT', jsonb_build_object(
  'shortcut_key','google',
  'description', 'Internet URL example',
  'modul', '', 'icon', '', 'method', 'METHOD_GET',
  'func_name', 'search',
  'address', 'https://www.google.com/search',
  'fields', jsonb_build_array(
    jsonb_build_object(
      'field_name', 'q', 'description', 'google search',
      'field_type', 'SHORTCUT_STRING', 'order', 0
    )
  )
));