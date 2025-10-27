INSERT INTO auth(id, code, user_name, user_group, disabled, auth_meta) 
VALUES(1, 'USR0000000000N1', 'admin', 'GROUP_ADMIN', false, json_object(
  'tags', json_array()
));

INSERT INTO auth(user_name, user_group, auth_meta, auth_map)
select e.username, 'GROUP_'||upper(ug.groupvalue), 
  json_object('tags', json_array()), json_object('empnumber', e.empnumber)
from bck_employee e
inner join bck_groups ug on e.usergroup = ug.id
where ug.groupvalue in('admin','user','guest') and e.username <> 'admin' and e.username is not null;

INSERT INTO config(config_type, data) 
select 'CONFIG_MAP' as config_type, json_object('field_name', df.fieldname, 'field_type',
  case when ft.groupvalue = 'time' then 'FIELD_DATETIME'
  when ft.groupvalue = 'float' then 'FIELD_NUMBER'
  when ft.groupvalue = 'valuelist' then 'FIELD_ENUM'
  when ft.groupvalue = 'notes' then 'FIELD_MEMO'
  when ft.groupvalue = 'urlink' then 'FIELD_URL'
  when ft.groupvalue = 'filter' then 'FIELD_STRING'
  when ft.groupvalue = 'password' then 'FIELD_STRING'
  when ft.groupvalue = 'transitem' then 'FIELD_TRANS_ITEM'
  when ft.groupvalue = 'transmovement' then 'FIELD_TRANS_MOVEMENT'
  when ft.groupvalue = 'transpayment' then 'FIELD_TRANS_PAYMENT'
  else 'FIELD_'||upper(ft.groupvalue) end,
  'description', df.description, 
  'tags', json('["'||replace(df.valuelist,'|','","')||'"]'),
  'filter', case when st.groupvalue is null then json_array() else json_array('FILTER_'||upper(st.groupvalue)) end) as data
from bck_deffield df
inner join bck_groups nt on df.nervatype = nt.id
inner join bck_groups ft on df.fieldtype = ft.id
left join bck_groups st on df.subtype = st.id
where df.fieldname not in('trans_custinvoice_compname','trans_custinvoice_compaddress','trans_custinvoice_comptax',
'trans_custinvoice_custname','trans_custinvoice_custaddress','trans_custinvoice_custtax','trans_wsdistance',
'trans_wsrepair','trans_wstotal','trans_wsnote','trans_reholiday','trans_rebadtool','trans_reother','trans_rentnote');

INSERT INTO config(config_type, data) 
select 'CONFIG_SHORTCUT' as config_type, json_object(
  'shortcut_key', m.menukey, 'description', m.description, 'modul', m.modul,
  'icon', m.icon, 'func_name', m.funcname, 'address', m.address,
  'method', 'METHOD_'||upper(mm.groupvalue),
  'fields', json(mf.fields)) as data
from bck_ui_menu m
inner join bck_groups mm on m.method = mm.id
left join(
  select mf.menu_id, json_group_array(json_object('field_name', mf.fieldname, 'description', mf.description,
  'field_type', case when ft.groupvalue = 'time' then 'SHORTCUT_DATETIME'
  when ft.groupvalue = 'float' then 'SHORTCUT_NUMBER'
  when ft.groupvalue = 'valuelist' then 'SHORTCUT_ENUM'
  when ft.groupvalue = 'notes' then 'SHORTCUT_MEMO'
  when ft.groupvalue = 'urlink' then 'SHORTCUT_URL'
  else 'SHORTCUT_'||upper(ft.groupvalue) end,
  'order', mf.orderby)) as fields
  from bck_ui_menufields mf
  inner join bck_groups ft on mf.fieldtype = ft.id
  group by mf.menu_id) mf on mf.menu_id = m.id;

INSERT INTO config(config_type, data) 
select 'CONFIG_MESSAGE' as config_type, json_object('section', m.secname, 'key', m.fieldname, 'lang', m.lang, 'value', m.msg) as data
from bck_ui_message m;

INSERT INTO config(config_type, data) 
select 'CONFIG_PATTERN' as config_type, json_object(
  'trans_type', 'TRANS_'||upper(tt.groupvalue), 'description', p.description, 
  'notes', p.notes, 'default_pattern', (p.defpattern) = 1
)
from bck_pattern p
inner join bck_groups tt on p.transtype = tt.id
where p.deleted = 0;

INSERT INTO currency(id, code, currency_meta, currency_map)
select cu.id, upper(cu.curr) as code,
  json_object(
	'description', COALESCE(cu.description, ''), 'digit', COALESCE(cu.digit, 0),
	'cash_round', COALESCE(cu.cround, 0), 'tags', json_array()
  ) as currency_meta,
  COALESCE(fld.md, json_object()) as currency_map
from bck_currency cu
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='currency') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = cu.id;

INSERT INTO customer(id, code, customer_name, customer_type, addresses, contacts, events, customer_meta, customer_map)
select c.id, 'CUS'||unixepoch()||'N'||c.id as code,
  c.custname as customer_name, 
  'CUSTOMER_'||upper(ct.groupvalue) as customer_type,
  COALESCE(addr.addresses, json_array()) as addresses, COALESCE(cont.contacts, json_array()) as contacts,
  COALESCE(evt.events, json_array()) as events,
  json_object(
	'tax_number', COALESCE(c.taxnumber, ''), 'account', COALESCE(c.account, ''), 'tax_free', (c.notax = 1),
	'terms', COALESCE(c.terms, 0), 'credit_limit', COALESCE(c.creditlimit, 0), 'discount', COALESCE(c.discount, 0),
	'notes', COALESCE(c.notes, ''), 'inactive', (c.inactive = 1), 'tags', json_array()
  ) as customer_meta,
  json_patch(COALESCE(fld.md, json_object()), json_object('custnumber', c.custnumber)) as customer_map
from bck_customer c
inner join bck_groups ct on c.custtype = ct.id
left join (select a.ref_id, json_group_array(json_object(
  'country', COALESCE(a.country, ''), 'state', COALESCE(a.state, ''), 'zip_code', COALESCE(a.zipcode, ''), 'city', COALESCE(a.city, ''), 
  'street', COALESCE(a.street, ''), 'notes', COALESCE(a.notes, ''), 'tags', json_array(), 'address_map', json_object())) as addresses
  from bck_address a
  where a.deleted = 0 and a.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='customer')
  group by a.ref_id) addr on addr.ref_id = c.id
left join (select co.ref_id, json_group_array(json_object(
  'first_name', COALESCE(co.firstname, ''), 'surname', COALESCE(co.surname, ''), 'status', COALESCE(co.status, ''), 
  'phone', COALESCE(co.phone, ''), 'mobile', COALESCE(co.mobil, ''), 'email', COALESCE(co.email, ''), 
  'notes', COALESCE(co.notes, ''), 'tags', json_array(), 'contact_map', json_object())) as contacts
  from bck_contact co
  where co.deleted = 0 and co.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='customer')
  group by co.ref_id) cont on cont.ref_id = c.id
left join (select ev.ref_id, json_group_array(json_object(
  'uid', COALESCE(ev.uid, ''), 'subject', COALESCE(ev.subject, ''), 'start_time', ev.fromdate, 'end_time', ev.todate, 'place', COALESCE(ev.place,''), 
  'description', COALESCE(ev.description,''), 
  'tags', case when eg.groupvalue is null then json_array() else json_array(eg.groupvalue) end, 
  'event_map', COALESCE(json(fld.md), json_object()))) as events
  from bck_event ev
  left join bck_groups eg on ev.eventgroup = eg.id
  left join (
    select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
    from bck_fieldvalue fv 
    where fv.deleted = 0 and fv.fieldname in(
      select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='event') and deleted = 0)
    group by fv.ref_id) fld on fld.ref_id = ev.id
  where ev.deleted = 0 and ev.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='customer')
  group by ev.ref_id) evt on evt.ref_id = c.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='customer') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = c.id
where c.deleted = 0;

INSERT INTO employee(id, code, address, contact, events, employee_meta, employee_map)
select e.id, 'EMP'||unixepoch()||'N'||e.id as code,
  COALESCE(json_extract(addr.addresses,'$[0]'), json_object()) as address, 
  COALESCE(json_extract(cont.contacts,'$[0]'), json_object()) as contact,
  COALESCE(evt.events, json_array()) as events,
  json_object(
	'start_date', COALESCE(e.startdate, ''), 'end_date', COALESCE(e.enddate, ''),
	'notes', '', 'inactive', (e.inactive = 1), 'tags', json_array()
  ) as employee_meta,
  json_patch(COALESCE(fld.md, json_object()), json_object('empnumber', e.empnumber)) as employee_map
from bck_employee e
left join (select a.ref_id, json_group_array(json_object(
  'country', COALESCE(a.country, ''), 'state', COALESCE(a.state, ''), 'zip_code', COALESCE(a.zipcode, ''), 'city', COALESCE(a.city, ''), 
  'street', COALESCE(a.street, ''), 'notes', COALESCE(a.notes, ''), 'tags', json_array(), 'address_map', json_object())) as addresses
  from bck_address a
  where a.deleted = 0 and a.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='employee')
  group by a.ref_id) addr on addr.ref_id = e.id
left join (select co.ref_id, json_group_array(json_object(
  'first_name', COALESCE(co.firstname, ''), 'surname', COALESCE(co.surname, ''), 'status', COALESCE(co.status, ''), 
  'phone', COALESCE(co.phone, ''), 'mobile', COALESCE(co.mobil, ''), 'email', COALESCE(co.email, ''), 
  'notes', COALESCE(co.notes, ''), 'tags', json_array(), 'contact_map', json_object())) as contacts
  from bck_contact co
  where co.deleted = 0 and co.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='employee')
  group by co.ref_id) cont on cont.ref_id = e.id
left join (select ev.ref_id, json_group_array(json_object(
  'uid', COALESCE(ev.uid, ''), 'subject', COALESCE(ev.subject, ''), 'start_time', ev.fromdate, 'end_time', ev.todate, 'place', COALESCE(ev.place,''), 
  'description', COALESCE(ev.description,''), 
  'tags', case when eg.groupvalue is null then json_array() else json_array(eg.groupvalue) end, 
  'event_map', COALESCE(json(fld.md), json_object()))) as events
  from bck_event ev
  left join bck_groups eg on ev.eventgroup = eg.id
  left join (
    select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
    from bck_fieldvalue fv 
    where fv.deleted = 0 and fv.fieldname in(
      select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='event') and deleted = 0)
    group by fv.ref_id) fld on fld.ref_id = ev.id
  where ev.deleted = 0 and ev.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='employee')
  group by ev.ref_id) evt on evt.ref_id = e.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='employee') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = e.id
where e.deleted = 0;

INSERT INTO place(id, code, place_name, place_type, currency_code, address, contacts, place_meta, place_map)
select p.id, 'PLA'||unixepoch()||'N'||p.id as code,
  p.description as place_name, 
  'PLACE_'||upper(pt.groupvalue) as place_type, p.curr as currency_code,
  COALESCE(json_extract(addr.addresses,'$[0]'), json_object()) as address, COALESCE(cont.contacts, json_array()) as contacts,
  json_object(
	'notes', COALESCE(p.notes, ''), 'inactive', (p.inactive = 1), 'tags', json_array()
  ) as place_meta,
  json_patch(COALESCE(fld.md, json_object()), json_object('planumber', p.planumber)) as place_map
from bck_place p
inner join bck_groups pt on p.placetype = pt.id
left join (select a.ref_id, json_group_array(json_object(
  'country', COALESCE(a.country, ''), 'state', COALESCE(a.state, ''), 'zip_code', COALESCE(a.zipcode, ''), 'city', COALESCE(a.city, ''), 
  'street', COALESCE(a.street, ''), 'notes', COALESCE(a.notes, ''), 'tags', json_array(), 'address_map', json_object())) as addresses
  from bck_address a
  where a.deleted = 0 and a.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='place')
  group by a.ref_id) addr on addr.ref_id = p.id
left join (select co.ref_id, json_group_array(json_object(
  'first_name', COALESCE(co.firstname, ''), 'surname', COALESCE(co.surname, ''), 'status', COALESCE(co.status, ''), 
  'phone', COALESCE(co.phone, ''), 'mobile', COALESCE(co.mobil, ''), 'email', COALESCE(co.email, ''), 
  'notes', COALESCE(co.notes, ''), 'tags', json_array(), 'contact_map', json_object())) as contacts
  from bck_contact co
  where co.deleted = 0 and co.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='place')
  group by co.ref_id) cont on cont.ref_id = p.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='place') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = p.id
where p.deleted = 0;

INSERT INTO tax(id, code, tax_meta, tax_map)
select tx.id, upper(tx.taxcode) as code,
  json_object(
	'description', COALESCE(tx.description, ''), 'rate_value', COALESCE(tx.rate, 0),
	'inactive', (tx.inactive = 1), 'tags', json_array()
  ) as tax_meta,
  COALESCE(fld.md, json_object()) as tax_map
from bck_tax tx
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='tax') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = tx.id;

INSERT INTO product(id, code, product_name, product_type, tax_code, events, product_meta, product_map)
select p.id, 'PRO'||unixepoch()||'N'||p.id as code,
  p.description as product_name, 
  'PRODUCT_'||upper(pt.groupvalue) as product_type,
  tx.taxcode as tax_code,
  COALESCE(evt.events, json_array()) as events,
  json_object(
	'unit', COALESCE(p.unit, ''),
	'barcode_type', COALESCE(json_extract(json_extract(bar.barcodes,'$[0]'),'$.barcode_type'), ''),
	'barcode', COALESCE(json_extract(json_extract(bar.barcodes,'$[0]'),'$.code'), ''),
	'barcode_qty', COALESCE(json_extract(json_extract(bar.barcodes,'$[0]'),'$.qty'), 0),
	'notes', COALESCE(p.notes, ''), 'inactive', (p.inactive = 1), 'tags', json_array()
  ) as product_meta,
  json_patch(COALESCE(fld.md, json_object()), json_object('partnumber', p.partnumber)) as product_map
from bck_product p
inner join bck_groups pt on p.protype = pt.id
inner join bck_tax tx on p.tax_id = tx.id
left join (select ev.ref_id, json_group_array(json_object(
  'uid', COALESCE(ev.uid, ''), 'subject', COALESCE(ev.subject, ''), 'start_time', ev.fromdate, 'end_time', ev.todate, 'place', COALESCE(ev.place,''), 
  'description', COALESCE(ev.description,''), 
  'tags', case when eg.groupvalue is null then json_array() else json_array(eg.groupvalue) end, 
  'event_map', COALESCE(json(fld.md), json_object()))) as events
  from bck_event ev
  left join bck_groups eg on ev.eventgroup = eg.id
  left join (
    select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
    from bck_fieldvalue fv 
    where fv.deleted = 0 and fv.fieldname in(
      select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='event') and deleted = 0)
    group by fv.ref_id) fld on fld.ref_id = ev.id
  where ev.deleted = 0 and ev.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='product')
  group by ev.ref_id) evt on evt.ref_id = p.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='product') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = p.id
left join (select bc.product_id, json_group_array(json_object(
  'code', COALESCE(bc.code, ''), 'description', COALESCE(bc.description, ''), 
  'barcode_type', COALESCE(upper('BARCODE_'||bt.groupvalue), ''), 'qty', COALESCE(bc.qty, 0)
  )) as barcodes
  from bck_barcode bc
  inner join bck_groups bt on bc.barcodetype = bt.id
  where bc.defcode = 1
  group by bc.product_id) bar on bar.product_id = p.id
where p.deleted = 0;

INSERT INTO price(id, code, price_type, valid_from, valid_to, product_code, currency_code, 
  customer_code, qty, price_meta, price_map)
select pr.id, 'PRC'||unixepoch()||'N'||pr.id as code,
  case when pr.vendorprice = 1 then 'PRICE_VENDOR' else 'PRICE_CUSTOMER' end as price_type,
  pr.validfrom as valid_from, pr.validto as valid_to, p.code as product_code, pr.curr as currency_code,
  lnk.code as customer_code, pr.qty,
  json_object(
	'price_value', pr.pricevalue, 'tags', json_array()
  ) as price_meta, COALESCE(fld.md, json_object()) as price_map
from bck_price pr
inner join product p on pr.product_id = p.id
left join (
  select l.ref_id_1 as price_id, c.code
  from bck_link l
  inner join customer c on l.ref_id_2 = c.id
  where l.deleted = 0 
    and l.nervatype_1 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='price')
    and l.nervatype_2 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='customer')
) lnk on lnk.price_id = pr.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='price') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = pr.id
where pr.deleted = 0;

INSERT INTO project(id, code, project_name, customer_code, addresses, contacts, events, project_meta, project_map)
select p.id, 'PRJ'||unixepoch()||'N'||p.id as code,
  p.description as project_name, c.code as customer_code,
  COALESCE(addr.addresses, json_array()) as addresses, COALESCE(cont.contacts, json_array()) as contacts,
  COALESCE(evt.events, json_array()) as events,
  json_object(
	'start_date', COALESCE(p.startdate, ''), 'end_date', COALESCE(p.enddate, ''),
	'notes', COALESCE(p.notes, ''), 'inactive', (p.inactive = 1), 'tags', json_array()
  ) as project_meta,
  json_patch(COALESCE(fld.md, json_object()), json_object('pronumber', p.pronumber)) as project_map
from bck_project p
left join customer c on p.customer_id = c.id
left join (select a.ref_id, json_group_array(json_object(
  'country', COALESCE(a.country, ''), 'state', COALESCE(a.state, ''), 'zip_code', COALESCE(a.zipcode, ''), 'city', COALESCE(a.city, ''), 
  'street', COALESCE(a.street, ''), 'notes', COALESCE(a.notes, ''), 'tags', json_array(), 'address_map', json_object())) as addresses
  from bck_address a
  where a.deleted = 0 and a.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='project')
  group by a.ref_id) addr on addr.ref_id = p.id
left join (select co.ref_id, json_group_array(json_object(
  'first_name', COALESCE(co.firstname, ''), 'surname', COALESCE(co.surname, ''), 'status', COALESCE(co.status, ''), 
  'phone', COALESCE(co.phone, ''), 'mobile', COALESCE(co.mobil, ''), 'email', COALESCE(co.email, ''), 
  'notes', COALESCE(co.notes, ''), 'tags', json_array(), 'contact_map', json_object())) as contacts
  from bck_contact co
  where co.deleted = 0 and co.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='project')
  group by co.ref_id) cont on cont.ref_id = p.id
left join (select ev.ref_id, json_group_array(json_object(
  'uid', COALESCE(ev.uid, ''), 'subject', COALESCE(ev.subject, ''), 'start_time', ev.fromdate, 'end_time', ev.todate, 'place', COALESCE(ev.place,''), 
  'description', COALESCE(ev.description,''), 
  'tags', case when eg.groupvalue is null then json_array() else json_array(eg.groupvalue) end, 
  'event_map', COALESCE(json(fld.md), json_object()))) as events
  from bck_event ev
  left join bck_groups eg on ev.eventgroup = eg.id
  left join (
    select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
    from bck_fieldvalue fv 
    where fv.deleted = 0 and fv.fieldname in(
      select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='event') and deleted = 0)
    group by fv.ref_id) fld on fld.ref_id = ev.id
  where ev.deleted = 0 and ev.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='project')
  group by ev.ref_id) evt on evt.ref_id = p.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='project') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = p.id
where p.deleted = 0;

INSERT INTO rate(id, code, rate_type, rate_date, place_code, currency_code, rate_meta, rate_map)
select r.id, 'RAT'||unixepoch()||'N'||r.id as code,
  'RATE_'||upper(rt.groupvalue) as rate_type,
  r.ratedate as rate_date, p.code as place_code, r.curr as currency_code,
  json_object(
	'rate_value', r.ratevalue, 'tags', json_array()
  ) as rate_meta, COALESCE(fld.md, json_object()) as rate_map
from bck_rate r
inner join bck_groups rt on r.ratetype = rt.id
left join place p on r.place_id = p.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='rate') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = r.id
where r.deleted = 0;

INSERT INTO tool(id, code, description, product_code, events, tool_meta, tool_map)
select t.id, 'SER'||unixepoch()||'N'||t.id as code,
  t.description as description, p.code as product_code,
  COALESCE(evt.events, json_array()) as events,
  json_object(
	'serial_number', COALESCE(t.serial, ''),
	'notes', COALESCE(t.notes, ''), 'inactive', (t.inactive = 1), 'tags', json_array()
  ) as tool_meta,
  COALESCE(fld.md, json_object()) as tool_map
from bck_tool t
inner join product p on t.product_id = p.id
left join (select ev.ref_id, json_group_array(json_object(
  'uid', COALESCE(ev.uid, ''), 'subject', COALESCE(ev.subject, ''), 'start_time', ev.fromdate, 'end_time', ev.todate, 'place', COALESCE(ev.place,''), 
  'description', COALESCE(ev.description,''), 
  'tags', case when eg.groupvalue is null then json_array() else json_array(eg.groupvalue) end, 
  'event_map', COALESCE(json(fld.md), json_object()))) as events
  from bck_event ev
  left join bck_groups eg on ev.eventgroup = eg.id
  left join (
    select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
    from bck_fieldvalue fv 
    where fv.deleted = 0 and fv.fieldname in(
      select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='event') and deleted = 0)
    group by fv.ref_id) fld on fld.ref_id = ev.id
  where ev.deleted = 0 and ev.nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='tool')
  group by ev.ref_id) evt on evt.ref_id = t.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='tool') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = t.id
where t.deleted = 0;

INSERT INTO trans(id, code, trans_type, trans_date, direction, customer_code, 
  employee_code, project_code, place_code, currency_code, auth_code, trans_meta, trans_map)
select t.id, 
  CASE WHEN upper(tt.groupvalue) = 'INVENTORY' then 'COR'
    WHEN upper(tt.groupvalue) = 'DELIVERY' and upper(gd.groupvalue) = 'TRANSFER' then 'TRF' 
    else substr(upper(tt.groupvalue),1,3)end||unixepoch()||'N'||t.id as code, 
  'TRANS_'||upper(tt.groupvalue) as trans_type, t.transdate as trans_date,
  'DIRECTION_'||upper(gd.groupvalue) as direction, c.code as customer_code,
  e.code as employee_code, p.code as project_code, pl.code as place_code, t.curr as currency_code, a.code as auth_code,
  json_object(
	'due_time', COALESCE(t.duedate, ''), 'ref_number', COALESCE(t.ref_transnumber, ''),
	'paid_type', 'PAID_'||upper(gd.groupvalue), 'tax_free', (t.notax = 1), 'paid', (t.paid = 1),
	'rate', COALESCE(t.acrate, 0), 
	'status', COALESCE('STATUS_'||upper(COALESCE(fld.md, json_object())->>'trans_transcast'),''),
	'trans_state', 'STATE_'||upper(tstat.groupvalue), 'closed', (t.closed = 1),
	'notes', COALESCE(t.notes, ''), 'internal_notes', COALESCE(t.intnotes, ''), 'report_notes', COALESCE(t.fnote, ''),
	'worksheet', json_object(
	  'distance', CAST(COALESCE(COALESCE(fld.md, json_object())->>'trans_wsdistance',0) as float),
	  'repair', CAST(COALESCE(COALESCE(fld.md, json_object())->>'trans_wsrepair',0) as float),
	  'total', CAST(COALESCE(COALESCE(fld.md, json_object())->>'trans_wstotal',0) as float),
	  'notes', COALESCE(COALESCE(fld.md, json_object())->>'trans_wsnote','')
	 ),
	 'rent', json_object(
	  'holiday', CAST(COALESCE(COALESCE(fld.md, json_object())->>'trans_reholiday',0) as float),
	  'bad_tool', CAST(COALESCE(COALESCE(fld.md, json_object())->>'trans_rebadtool',0) as float),
	  'other', CAST(COALESCE(COALESCE(fld.md, json_object())->>'trans_reother',0) as float),
	  'notes', COALESCE(COALESCE(fld.md, json_object())->>'trans_rentnote','')
	 ),
	 'invoice', json_object(
	  'company_name', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_compname',''),
	  'company_address', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_compaddress',''),
	  'company_tax_number', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_comptax',''),
	  'company_account', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_compaccount',''),
	  'customer_name', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_custname',''),
	  'customer_address', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_custaddress',''),
	  'customer_tax_number', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_custtax',''),
	  'customer_account', COALESCE(COALESCE(fld.md, json_object())->>'trans_custinvoice_custaccount','')
	 ),
	'tags', json_array()
  ) as trans_meta,
  json_patch(COALESCE(fld.md, json_object()), json_object('transnumber', t.transnumber)) as trans_map
from bck_trans t
inner join bck_groups tt on t.transtype = tt.id
inner join bck_groups gd on t.direction = gd.id
left join bck_groups pt on t.paidtype = pt.id
left join bck_groups tstat on t.transtate = tstat.id
left join customer c on t.customer_id = c.id
left join employee e on t.employee_id = e.id
left join project p on t.project_id = p.id
left join place pl on t.place_id = pl.id
left join auth a on t.cruser_id = a.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield 
	where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='trans') 
	and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = t.id
where t.deleted = 0;

UPDATE trans set trans_code = tr.code
FROM (select l.ref_id_1 as trans_id, t.code
from bck_link l
inner join trans t on l.ref_id_2 = t.id
where l.deleted = 0 
  and l.nervatype_1 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='trans')
  and l.nervatype_2 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='trans')
) tr
where tr.trans_id = id;

INSERT INTO item(id, code, trans_code, product_code, tax_code, item_meta, item_map)
select i.id, 'ITM'||unixepoch()||'N'||i.id as code,
  t.code as trans_code, p.code as product_code, tx.code as tax_code,
  json_object(
	'unit', i.unit, 'qty', i.qty, 'fx_price', i.fxprice, 'net_amount', i.netamount,
	'discount', i.discount, 'vat_amount', i.vatamount, 'amount', i.amount,
	'description', i.description, 'deposit', (i.deposit = 1),
	'own_stock', i.ownstock, 'action_price', (i.actionprice = 1),
	'tags', json_array()
  ) as item_meta, COALESCE(fld.md, json_object()) as item_map
from bck_item i
inner join trans t on i.trans_id = t.id
inner join product p on i.product_id = p.id
inner join tax tx on i.tax_id = tx.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='item') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = i.id
where i.deleted = 0;

INSERT INTO movement(id, code, movement_type, shipping_time, trans_code, product_code, tool_code, 
  place_code, movement_meta, movement_map)
select mv.id, 'MOV'||unixepoch()||'N'||mv.id as code,
  'MOVEMENT_'||upper(mt.groupvalue) as movement_type, mv.shippingdate as shipping_time,
  t.code as trans_code, p.code as product_code, tl.code as tool_code, pl.code as place_code,
  json_object(
	'qty', mv.qty, 'notes', COALESCE(mv.notes, ''), 'shared', (mv.shared = 1),
	'tags', json_array()
  ) as movement_meta, COALESCE(fld.md, json_object()) as movement_map
from bck_movement mv
inner join bck_groups mt on mv.movetype = mt.id
inner join trans t on mv.trans_id = t.id
left join product p on mv.product_id = p.id
left join tool tl on mv.tool_id = tl.id
left join place pl on mv.place_id = pl.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='movement') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = mv.id
where mv.deleted = 0;

UPDATE movement set item_code = tr.code 
FROM (
  select l.ref_id_1 as movement_id, i.code
  from bck_link l
  inner join item i on l.ref_id_2 = i.id
  where l.deleted = 0 
    and l.nervatype_1 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='movement')
    and l.nervatype_2 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='item')
) tr
where tr.movement_id = id;

UPDATE movement set movement_code = tr.code 
FROM (
  select l.ref_id_2 as movement_id, mv.code
  from bck_link l
  inner join movement mv on l.ref_id_1 = mv.id
  where l.deleted = 0 
    and l.nervatype_1 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='movement')
    and l.nervatype_2 = (select id from bck_groups where groupname = 'nervatype' and groupvalue='movement')
) tr
where tr.movement_id = id;

INSERT INTO payment(id, code, paid_date, trans_code, payment_meta, payment_map)
select pm.id, 'PMT'||unixepoch()||'N'||pm.id as code,
  pm.paiddate as paid_date, t.code as trans_code,
  json_object(
	'amount', pm.amount, 'notes', COALESCE(pm.notes, ''),
	'tags', json_array()
  ) as payment_meta, COALESCE(fld.md, json_object()) as payment_map
from bck_payment pm
inner join trans t on pm.trans_id = t.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='payment') and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = pm.id
where pm.deleted = 0;

INSERT INTO link(id, code, link_type_1, link_code_1, link_type_2, link_code_2, link_meta, link_map)
select l.id, 'LNK'||unixepoch()||'N'||l.id as code,
  'LINK_'||upper(nt1.groupvalue) as link_type_1,
  case when nt1.groupvalue = 'customer' then cu1.code
    when nt1.groupvalue = 'employee' then e1.code
	when nt1.groupvalue = 'item' then i1.code
	when nt1.groupvalue = 'movement' then mv1.code
	when nt1.groupvalue = 'payment' then pm1.code
	when nt1.groupvalue = 'place' then pl1.code
	when nt1.groupvalue = 'product' then p1.code
	when nt1.groupvalue = 'project' then pr1.code
	when nt1.groupvalue = 'tool' then tl1.code
	else t1.code end as link_code_1,
  'LINK_'||upper(nt2.groupvalue) as link_type_2,
  case when nt2.groupvalue = 'customer' then cu2.code
    when nt2.groupvalue = 'employee' then e2.code
	when nt2.groupvalue = 'item' then i2.code
	when nt2.groupvalue = 'movement' then mv2.code
	when nt2.groupvalue = 'payment' then pm2.code
	when nt2.groupvalue = 'place' then pl2.code
	when nt2.groupvalue = 'product' then p2.code
	when nt2.groupvalue = 'project' then pr2.code
	when nt2.groupvalue = 'tool' then tl2.code
	else t2.code end as link_code_2,
  json_object(
	'amount', CAST(COALESCE(COALESCE(fld.md, json_object())->>'link_qty',0) as float),
	'rate', CAST(COALESCE(COALESCE(fld.md, json_object())->>'link_rate',0) as float),
	'tags', json_array()
  ) as link_meta, COALESCE(fld.md, json_object()) as link_map
from bck_link l
inner join bck_groups nt1 on l.nervatype_1 = nt1.id
left join customer cu1 on l.ref_id_1 = cu1.id
left join employee e1 on l.ref_id_1 = e1.id
left join item i1 on l.ref_id_1 = i1.id
left join movement mv1 on l.ref_id_1 = mv1.id
left join payment pm1 on l.ref_id_1 = pm1.id
left join place pl1 on l.ref_id_1 = pl1.id
left join product p1 on l.ref_id_1 = p1.id
left join project pr1 on l.ref_id_1 = pr1.id
left join tool tl1 on l.ref_id_1 = tl1.id
left join trans t1 on l.ref_id_1 = t1.id
inner join bck_groups nt2 on l.nervatype_2 = nt2.id
left join customer cu2 on l.ref_id_2 = cu2.id
left join employee e2 on l.ref_id_2 = e2.id
left join item i2 on l.ref_id_2 = i2.id
left join movement mv2 on l.ref_id_2 = mv2.id
left join payment pm2 on l.ref_id_2 = pm2.id
left join place pl2 on l.ref_id_2 = pl2.id
left join product p2 on l.ref_id_2 = p2.id
left join project pr2 on l.ref_id_2 = pr2.id
left join tool tl2 on l.ref_id_2 = tl2.id
left join trans t2 on l.ref_id_2 = t2.id
left join (
  select fv.ref_id, json_group_object(fv.fieldname, fv.value) as md
  from bck_fieldvalue fv 
  where fv.deleted = 0 and fv.fieldname in(
    select fieldname from bck_deffield 
	where nervatype = (select id from bck_groups where groupname = 'nervatype' and groupvalue='link') 
	and deleted = 0)
  group by fv.ref_id) fld on fld.ref_id = l.id
where l.deleted = 0 
  and l.nervatype_1 in (select id from bck_groups where groupname = 'nervatype' 
    and groupvalue in('customer', 'employee', 'item', 'movement', 'payment', 'place', 'product', 'project', 'tool', 'trans'))
  and l.nervatype_2 in (select id from bck_groups where groupname = 'nervatype' 
    and groupvalue in('customer', 'employee', 'item', 'movement', 'payment', 'place', 'product', 'project', 'tool', 'trans'));