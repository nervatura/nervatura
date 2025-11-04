DROP VIEW IF EXISTS link_view;
DROP VIEW IF EXISTS link_map;
DROP VIEW IF EXISTS link_tags;

DROP VIEW IF EXISTS item_shipping;

DROP VIEW IF EXISTS movement_view;
DROP VIEW IF EXISTS movement_map;
DROP VIEW IF EXISTS movement_tags;
DROP VIEW IF EXISTS movement_stock;
DROP VIEW IF EXISTS movement_inventory;
DROP VIEW IF EXISTS movement_waybill;
DROP VIEW IF EXISTS movement_formula;
DROP TABLE IF EXISTS movement;

DROP VIEW IF EXISTS payment_view;
DROP VIEW IF EXISTS payment_map;
DROP VIEW IF EXISTS payment_invoice;
DROP VIEW IF EXISTS payment_tags;
DROP TABLE IF EXISTS payment;

DROP TABLE IF EXISTS link;

DROP VIEW IF EXISTS item_view;
DROP VIEW IF EXISTS item_map;
DROP VIEW IF EXISTS item_tags;
DROP TABLE IF EXISTS item;

DROP VIEW IF EXISTS trans_view;
DROP VIEW IF EXISTS trans_map;
DROP VIEW IF EXISTS trans_tags;
DROP TABLE IF EXISTS trans;

DROP VIEW IF EXISTS price_view;
DROP VIEW IF EXISTS price_map;
DROP VIEW IF EXISTS price_tags;
DROP TABLE IF EXISTS price;

DROP VIEW IF EXISTS project_view;
DROP VIEW IF EXISTS project_addresses;
DROP VIEW IF EXISTS project_contacts;
DROP VIEW IF EXISTS project_events;
DROP VIEW IF EXISTS project_map;
DROP VIEW IF EXISTS project_tags;
DROP TABLE IF EXISTS project;

DROP VIEW IF EXISTS rate_view;
DROP VIEW IF EXISTS rate_map;
DROP VIEW IF EXISTS rate_tags;
DROP TABLE IF EXISTS rate;

DROP VIEW IF EXISTS tool_view;
DROP VIEW IF EXISTS tool_events;
DROP VIEW IF EXISTS tool_map;
DROP VIEW IF EXISTS tool_tags;
DROP TABLE IF EXISTS tool;

DROP VIEW IF EXISTS customer_addresses;
DROP VIEW IF EXISTS customer_contacts;
DROP VIEW IF EXISTS customer_events;
DROP VIEW IF EXISTS customer_view;
DROP VIEW IF EXISTS customer_map;
DROP VIEW IF EXISTS customer_tags;
DROP TABLE IF EXISTS customer;

DROP VIEW IF EXISTS employee_view;
DROP VIEW IF EXISTS employee_events;
DROP VIEW IF EXISTS employee_map;
DROP VIEW IF EXISTS employee_tags;
DROP TABLE IF EXISTS employee;

DROP VIEW IF EXISTS place_view;
DROP VIEW IF EXISTS place_contacts;
DROP VIEW IF EXISTS place_map;
DROP VIEW IF EXISTS place_tags;
DROP TABLE IF EXISTS place;

DROP VIEW IF EXISTS product_view;
DROP VIEW IF EXISTS product_events;
DROP VIEW IF EXISTS product_map;
DROP VIEW IF EXISTS product_tags;
DROP TABLE IF EXISTS product;

DROP VIEW IF EXISTS currency_view;
DROP VIEW IF EXISTS currency_map;
DROP VIEW IF EXISTS currency_tags;
DROP TABLE IF EXISTS currency;

DROP VIEW IF EXISTS tax_view;
DROP VIEW IF EXISTS tax_map;
DROP VIEW IF EXISTS tax_tags;
DROP TABLE IF EXISTS tax;

DROP VIEW IF EXISTS auth_view;
DROP VIEW IF EXISTS auth_map;
DROP TABLE IF EXISTS auth;

DROP VIEW IF EXISTS config_map;
DROP VIEW IF EXISTS config_shortcut;
DROP VIEW IF EXISTS config_message;
DROP VIEW IF EXISTS config_pattern;
DROP VIEW IF EXISTS config_print_queue;
DROP VIEW IF EXISTS config_report;
DROP VIEW IF EXISTS config_data;
DROP TABLE IF EXISTS config;

DROP TABLE IF EXISTS log;
DROP TABLE IF EXISTS usref;

DROP FUNCTION IF EXISTS link_changed();
DROP FUNCTION IF EXISTS trans_invoice_customer_insert();
DROP FUNCTION IF EXISTS trans_invoice_customer_update();
DROP FUNCTION IF EXISTS set_invoice_customer;
DROP FUNCTION IF EXISTS set_trans_code();
DROP FUNCTION IF EXISTS set_new_code();
DROP FUNCTION IF EXISTS set_changed_timestamp();

DROP TYPE IF EXISTS config_type;
DROP TYPE IF EXISTS user_group;
DROP TYPE IF EXISTS customer_type;
DROP TYPE IF EXISTS place_type;
DROP TYPE IF EXISTS product_type;
DROP TYPE IF EXISTS rate_type;
DROP TYPE IF EXISTS price_type;
DROP TYPE IF EXISTS trans_type;
DROP TYPE IF EXISTS direction;
DROP TYPE IF EXISTS movement_type;
DROP TYPE IF EXISTS log_type;
DROP TYPE IF EXISTS link_type;
