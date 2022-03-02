export const Queries = ({ getText }) => {
  return {
    
    customer: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]]]}},
              
      CustomerView: {
        columns: {custnumber:true, custname:true, address:true},
        label: getText("customer_view"),
        fields: {
          custnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("customer_custnumber"), sqlstr:'c.custnumber '},
          custname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          taxnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("customer_taxnumber"), sqlstr:'c.taxnumber '},
          custtype: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("customer_custtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          account: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("customer_account"), sqlstr:'c.account '},
          notax: {fieldtype:'bool', wheretype:'where', orderby:5, 
            label:getText("customer_notax"), sqlstr:'c.notax '},
          terms: {fieldtype:'float', wheretype:'where', orderby:6, 
            label:getText("customer_terms"), sqlstr:'c.terms '},
          creditlimit: {fieldtype:'float', wheretype:'where', aggretype:'sum', orderby:7, 
            label:getText("customer_creditlimit"), sqlstr:'c.creditlimit '},
          discount: {fieldtype:'float', wheretype:'where', orderby:8, 
            label:getText("customer_discount"), sqlstr:'c.discount '},
          notes: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("customer_notes"), sqlstr:'c.notes '},
          inactive: {fieldtype:'bool', wheretype:'where', orderby:10, 
            label:getText("customer_inactive"), sqlstr:'c.inactive '},
          address: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("customer_address"), 
            sqlstr:"{CCS}case when addr.city is null then '' else addr.city end {SEP} ' ' {SEP} case when addr.street is null then '' else addr.street end{CCE} "}
        },
        sql: {
          select:["{CCS}'customer'{SEP}'//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as id","c.id as row_id",
            "c.custnumber","c.custname","case when mst.msg is null then tg.groupvalue else mst.msg end as custtype",
            "c.taxnumber","c.account","c.notax",
            "{FMS_FLOAT}c.terms{FME_FLOAT} as terms","c.terms as export_terms",
            "{FMS_FLOAT}c.creditlimit{FME_FLOAT} as creditlimit","c.creditlimit as export_creditlimit",
            "c.discount","c.notes","c.inactive",
            "{CCS}case when addr.city is null then '' else addr.city end {SEP} ' ' {SEP} case when addr.street is null then '' else addr.street end{CCE} as address"],
          from:"customer c",
          inner_join:["groups tg","on",["c.custtype","=","tg.id"]],
            left_join:[["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'custtype'"],
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
          [[[{select:["*"], from:"address",
            where:["id","in",[{select:["min(id) fid"], from:"address a",
              where:[["a.deleted","=","0"],["and","a.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]]],
              group_by:["a.ref_id"]}]]}],"addr"],"on",["c.id","=","addr.ref_id"]]],
          where:[["c.deleted","=","0"],["and","c.id","not in",[{select:["customer.id"], from:"customer", inner_join:["groups","on",[["customer.custtype","=","groups.id"],["and","groups.groupvalue","=","'own'"]]]}]]]}},
      
      CustomerFieldsView: {
        columns: {custname:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        fields: {
          custnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("customer_custnumber"), sqlstr:'c.custnumber'},
          custname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'customer'{SEP}'//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as id","fv.id as row_id",
            "fv.fieldname", "c.custnumber", "c.custname", "df.description as fielddef", "'fieldvalue' as form", "fg.groupvalue as fieldtype",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"],
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]],
              ["and","df.visible","=","1"],["and","df.deleted","=","0"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["customer c","on",["fv.ref_id","=","c.id"]]],
          left_join:[
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and","c.deleted","=","0"],
            ["and","c.id","not in",[{select:["customer.id"], from:"customer", inner_join:["groups","on",[["customer.custtype","=","groups.id"],["and","groups.groupvalue","=","'own'"]]]}]]]}},
      
      CustomerContactView: {
        columns: {custname:true, firstname:true, surname:true, phone:true},
        label: getText("contact_view"),
        fields: {
          custnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("customer_custnumber"), sqlstr:'c.custnumber '},
          custname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          firstname: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("contact_firstname"), sqlstr:'co.firstname'},
          surname: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("contact_surname"), sqlstr:' co.surname'},
          status: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("contact_status"), sqlstr:' co.status'},
          phone: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("contact_phone"), sqlstr:' co.phone'},
          fax: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("contact_fax"), sqlstr:' co.fax'},
          mobil: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("contact_mobil"), sqlstr:' co.mobil'},
          email: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("contact_email"), sqlstr:' co.email'},
          notes: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("contact_notes"), sqlstr:' co.notes'}
          },
        sql: {
          select:["{CCS}'customer'{SEP}'//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as id","co.id as row_id",
            "c.custnumber","c.custname","co.firstname","co.surname","co.status","co.phone","co.fax",
            "co.mobil","co.email","co.notes","'contact' as form","co.id as form_id"],
          from:"contact co",
          inner_join:["customer c","on",["co.ref_id","=","c.id"]],
          where:[["co.deleted","=","0"],["and","c.deleted","=","0"], 
            ["and","co.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]],
            ["and","c.id","not in",[{select:["customer.id"], from:"customer",
              inner_join:["groups","on",[["customer.custtype","=","groups.id"],
              ["and","groups.groupvalue","=","'own'"]]]}]]]}},
      
      CustomerAddressView: {
        columns: {custname:true, city:true, street:true},
        label: getText("address_view"),
        fields: {
          custnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("customer_custnumber"), sqlstr:'c.custnumber '},
          custname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          country: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("address_country"), sqlstr:'a.country'},
          state: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("address_state"), sqlstr:'a.state'},
          zipcode: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("address_zipcode"), sqlstr:'a.zipcode '},
          city: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("address_city"), sqlstr:'a.city'},
          street: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("address_street"), sqlstr:'a.street'},
          notes: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("address_notes"), sqlstr:'a.notes'}
        },
        sql: {
          select:["{CCS}'customer'{SEP}'//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as id","a.id as row_id",
            "c.custnumber","c.custname","a.country","a.state","a.zipcode","a.city","a.street","a.notes",
            "'address' as form","a.id as form_id"],
          from:"address a",
          inner_join:["customer c","on",["a.ref_id","=","c.id"]],
          where:[["a.deleted","=","0"],["and","c.deleted","=","0"],
            ["and","a.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]],
            ["and","c.id","not in",[{select:["customer.id"], from:"customer", 
              inner_join:["groups","on",[["customer.custtype","=","groups.id"],
              ["and","groups.groupvalue","=","'own'"]]]}]]]}},
      
      CustomerEvents: {
        columns: {custname:true, fromdate:true, subject:true},
        label: getText("event_view"),
        edit: "event",
        fields: {
          custnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("customer_custnumber"), sqlstr:'c.custnumber '},
          custname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          calnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("event_calnumber"), sqlstr:'e.calnumber'},
          eventgroup: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("event_group"), sqlstr:'eg.groupvalue'},
          fromdate: {fieldtype:'date', wheretype:'where', orderby:4, 
            label:getText("event_fromdate"), 
            sqlstr:'{FMS_DATE}e.fromdate {FME_DATE}'},
          fromtime: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("event_fromtime"), 
            sqlstr:'{FMS_TIME}e.fromdate {FME_TIME}'},
          todate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("event_todate"), 
            sqlstr:'{FMS_DATE}e.todate {FME_DATE}'},
          totime: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("event_totime"), 
            sqlstr:'{FMS_TIME}e.todate {FME_TIME}'},
          subject: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("event_subject"), sqlstr:'e.subject'},
          place: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("event_place"), sqlstr:'e.place'},
          description: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("event_description"), sqlstr:'e.description'}

        },
        sql: {
          select:["{CCS}'event'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","e.id as row_id",
            "c.custnumber","c.custname","e.calnumber","eg.groupvalue as eventgroup",
            "{FMS_DATE}e.fromdate {FME_DATE} as fromdate",
            "{FMS_TIME}e.fromdate {FME_TIME} as fromtime",
            "{FMS_DATE}e.todate {FME_DATE} as todate",
            "{FMS_TIME}e.todate {FME_TIME} as totime",
            "e.subject", "e.place", "e.description"],
          from:"event e",
          inner_join:["customer c","on",["e.ref_id","=","c.id"]],
          left_join:["groups eg","on",["e.eventgroup","=","eg.id"]],
          where:[["e.deleted","=","0"],["and","c.deleted","=","0"],
            ["and","e.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]],
            ["and","c.id","not in",[{select:["customer.id"], from:"customer", 
              inner_join:["groups","on",[["customer.custtype","=","groups.id"],
              ["and","groups.groupvalue","=","'own'"]]]}]]]}}};
    },
    
    employee: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]}},
        
      EmployeeView: {
        columns: {empnumber:true, firstname:true, surname:true, username:true},
        label: getText("employee_view"),
        fields: {
          empnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("employee_empnumber"), sqlstr:'e.empnumber '},
          firstname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("contact_firstname"), sqlstr:'c.firstname '},
          surname: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("contact_surname"), sqlstr:'c.surname '},
          username: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("employee_username"), sqlstr:'e.username '},
          startdate: {fieldtype:'date', wheretype:'where', orderby:4, 
            label:getText("employee_startdate"), sqlstr:'e.startdate '},
          enddate: {fieldtype:'date', wheretype:'where', orderby:5, 
            label:getText("employee_enddate"), sqlstr:'e.enddate '},
          status: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("contact_status"), sqlstr:'c.status '},
          phone: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("contact_phone"), sqlstr:'c.phone '},
          mobil: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("contact_mobil"), sqlstr:'c.mobil '},
          email: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("contact_email"), sqlstr:'c.email '},
          zipcode: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("address_zipcode"), sqlstr:'a.zipcode '},
          city: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("address_city"), sqlstr:'a.city '},
          street: {fieldtype:'string', wheretype:'where', orderby:12, 
            label:getText("address_street"), sqlstr:'a.street '},
          usergroup: {fieldtype:'string', wheretype:'where', orderby:13, 
            label:getText("employee_usergroup"), sqlstr:'ug.groupvalue '},
          department: {fieldtype:'string', wheretype:'where', orderby:14, 
            label:getText("employee_department"), sqlstr:'dg.groupvalue '},
          inactive: {fieldtype:'bool', wheretype:'where', orderby:15, 
            label:getText("employee_inactive"), sqlstr:'e.inactive '},
          notes: {fieldtype:'string', wheretype:'where', orderby:16, 
            label:getText("contact_notes"), sqlstr:'c.notes '}
        },
        sql: {
          select:["{CCS}'employee'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","e.id as row_id",
            "e.empnumber","c.firstname","c.surname","e.username",
            "{FMS_DATE}e.startdate {FME_DATE} as startdate","{FMS_DATE}e.enddate {FME_DATE} as enddate",
            "c.status","c.phone","c.mobil","c.email","a.zipcode","a.city","a.street",
            "ug.groupvalue as usergroup","dg.groupvalue as department","c.notes","e.inactive"],
          from:"employee e",
          inner_join:["groups ug","on",["e.usergroup","=","ug.id"]],
          left_join:[
            ["contact c","on",[["e.id","=","c.ref_id"],
              ["and","c.nervatype","=",[{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]],
            ["address a","on",[["e.id","=","a.ref_id"],
              ["and","a.nervatype","=",[{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]],
            ["groups dg","on",["e.department","=","dg.id"]]],
          where:[["e.deleted","=","0"]]}},
      
      EmployeeFieldsView: {
        columns: {empnumber:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        fields: {
          empnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("employee_empnumber"), sqlstr:'e.empnumber '},
          firstname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("contact_firstname"), sqlstr:'c.firstname '},
          surname: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("contact_surname"), sqlstr:'c.surname '},
          username: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("employee_username"), sqlstr:'e.username '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'employee'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","fv.id as row_id",
            "e.empnumber", "c.firstname", "c.surname", "e.username", "df.description as fielddef", "'fieldvalue' as form", "fg.groupvalue as fieldtype",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"],
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]], 
              ["and","df.visible","=","1"],["and","df.deleted","=","0"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["employee e","on",["fv.ref_id","=","e.id"]]],
          left_join:[
            ["contact c","on",[["e.id","=","c.ref_id"],
              ["and","c.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]],
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and","e.deleted","=","0"]]}},
      
      EmployeeEvents: {
        columns: {empnumber:true, fromdate:true, subject:true},
        label: getText("event_view"),
        edit: "event",
        fields: {
          empnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("employee_empnumber"), sqlstr:'em.empnumber '},
          firstname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("contact_firstname"), sqlstr:'c.firstname '},
          surname: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("contact_surname"), sqlstr:'c.surname '},
          username: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("employee_username"), sqlstr:'em.username '},
          calnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("event_calnumber"), sqlstr:'e.calnumber'},
          eventgroup: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("event_group"), sqlstr:'eg.groupvalue'},
          fromdate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("event_fromdate"), 
            sqlstr:'{FMS_DATE}e.fromdate {FME_DATE}'},
          fromtime: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("event_fromtime"), 
            sqlstr:'{FMS_TIME}e.fromdate {FME_TIME}'},
          todate: {fieldtype:'date', wheretype:'where', orderby:8, 
            label:getText("event_todate"), 
            sqlstr:'{FMS_DATE}e.todate {FME_DATE}'},
          totime: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("event_totime"), 
            sqlstr:'{FMS_TIME}e.todate {FME_TIME}'},
          subject: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("event_subject"), sqlstr:'e.subject'},
          place: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("event_place"), sqlstr:'e.place'},
          description: {fieldtype:'string', wheretype:'where', orderby:12, 
            label:getText("event_description"), sqlstr:'e.description'}

        },
        sql: {
          select:["{CCS}'event'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","e.id as row_id",
            "em.empnumber","c.firstname","c.surname","em.username",
            "e.calnumber","eg.groupvalue as eventgroup",
            "{FMS_DATE}e.fromdate {FME_DATE} as fromdate",
            "{FMS_TIME}e.fromdate {FME_TIME} as fromtime",
            "{FMS_DATE}e.todate {FME_DATE} as todate",
            "{FMS_TIME}e.todate {FME_TIME} as totime",
            "e.subject","e.place","e.description"],
          from:"event e",
          inner_join:["employee em","on",["e.ref_id","=","em.id"]],
          left_join:[
            ["contact c","on",[["em.id","=","c.ref_id"],
              ["and","c.nervatype","=",[{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]],
            ["groups eg","on",["e.eventgroup","=","eg.id"]]],
          where:[["e.deleted","=","0"],["and","em.deleted","=","0"],
            ["and","e.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]}}};
    },
    
    product: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'product'"]]}]]]}},

      ProductView: {
        columns: {partnumber:true, protype:true, description:true},
        label: getText("product_view"),
        fields: {
          partnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          protype: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_protype"), 
            sqlstr:'case when ms.msg is null then g.groupvalue else ms.msg end '},
          description: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("product_description"), sqlstr:'p.description '},
          unit: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("product_unit"), sqlstr:'p.unit '},
          tax: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("product_tax"), sqlstr:'t.taxcode '},
          webitem: {fieldtype:'bool', wheretype:'where', orderby:5, 
            label:getText("product_webitem"), sqlstr:'p.webitem  '},
          inactive: {fieldtype:'bool', wheretype:'where', orderby:6, 
            label:getText("product_inactive"), sqlstr:'p.inactive  '},
          notes: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("product_notes"), sqlstr:'p.notes '}
        },
        sql: {
          select:["{CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","p.id as row_id",
            "p.partnumber","case when ms.msg is null then g.groupvalue else ms.msg end as protype",
            "p.description","p.unit","t.taxcode as tax","p.notes",
            "p.webitem as webitem","p.inactive"],
          from:"product p", 
          inner_join:[["groups g","on",["p.protype","=","g.id"]],
            ["tax t","on",["p.tax_id","=","t.id"]]],
          left_join:["ui_message ms","on",[["ms.fieldname","=","g.groupvalue"],["and","ms.secname","=","'protype'"],
            ["and","ms.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
          where:[["p.deleted","=","0"]]}},
      
      ProductFieldsView: {
        columns: {custname:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        fields: {
          partnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber'},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_description"), sqlstr:'p.description '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","fv.id as row_id",
            "p.partnumber","p.description","df.description as fielddef","'fieldvalue' as form", "fg.groupvalue as fieldtype",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"], 
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'product'"]]}]], 
              ["and","df.visible","=","1"],["and","df.deleted","=","0"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["product p","on",["fv.ref_id","=","p.id"]]],
          left_join:[
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and","p.deleted","=","0"]]}},
      
      ProductBarcodeView: {
        columns: {partnumber:true, description:true, barcode:true},
        label: getText("barcode_view"),
        fields: {
          partnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          partname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_description"), sqlstr:'p.description '},
          unit: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("product_unit"), sqlstr:'p.unit '},
          description: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("barcode_description"), sqlstr:'b.description '},
          barcodetype: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("barcode_barcodetype"), sqlstr:'g.description '},
          qty: {fieldtype:'float', wheretype:'where', orderby:5, 
            label:getText("barcode_qty"), sqlstr:'b.qty '},
          defcode: {fieldtype:'bool', wheretype:'where', orderby:6, 
            label:getText("barcode_defcode"), sqlstr:'b.defcode '},
          barcode: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("barcode_code"), sqlstr:'b.id '}
        },
        sql: {
          select:["{CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","b.id as row_id",
            "p.partnumber","p.description as partname","p.unit","b.code as barcode",
            "b.description","g.description as barcodetype",
            "{FMS_FLOAT}b.qty{FME_FLOAT} as qty","b.qty as export_qty",
            "b.defcode as defcode",
            "'barcode' as form","b.id as form_id"],
          from:"barcode b", 
          inner_join:[
            ["product p","on",["b.product_id","=","p.id"]],["groups g","on",["b.barcodetype","=","g.id"]]],
          where:[["p.deleted","=","0"]]}},
      
      ProductPriceView: {
        columns: {partnumber:true, validfrom:true, curr:true, pricevalue:true},
        label: getText("price_view"),
        fields: {
          partnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_description"), sqlstr:'p.description '},
          unit: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("product_unit"), sqlstr:'p.unit '},
          vendor: {fieldtype:'bool', wheretype:'where', orderby:3, 
            label:getText("price_vendor"), sqlstr:'pr.vendorprice '},
          custname: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("price_custname"), sqlstr:'c.custname '},
          validfrom: {fieldtype:'date', wheretype:'where', orderby:5, 
            label:getText("price_validfrom"), sqlstr:'pr.validfrom '},
          validto: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("price_validto"), sqlstr:'pr.validto '},
          curr: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("price_curr"), sqlstr:'pr.curr '},
          qty: {fieldtype:'float', wheretype:'where', orderby:8, 
            label:getText("price_qty"), sqlstr:'pr.qty '},
          pricevalue: {fieldtype:'float', wheretype:'where', orderby:9, 
            label:getText("price_pricevalue"), sqlstr:'pr.pricevalue '}
        },
        sql: {
          select:["{CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","pr.id as row_id",
            "p.partnumber","p.description","p.unit","pr.vendorprice as vendor","c.custname",
            "{FMS_DATE}pr.validfrom {FME_DATE} as validfrom","{FMS_DATE}pr.validto {FME_DATE} as validto",
            "pr.curr",
            "{FMS_FLOAT}pr.qty{FME_FLOAT} as qty","pr.qty as export_qty",
            "pr.pricevalue","'price' as form","pr.id as form_id"],
          from:"price pr",
          inner_join:["product p","on",["pr.product_id","=","p.id"]],
          left_join:[
            ["link ln0","on",[
              ["ln0.nervatype_1","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'price'"]]}]],
              ["and","ln0.ref_id_1","=","pr.id"],["and","ln0.deleted","=","0"], 
              ["and","ln0.nervatype_2","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]]]],
            ["customer c","on",[["ln0.ref_id_2","=","c.id"],["and","c.deleted","=","0"]]]],
          where:[["p.deleted","=","0"],["and","pr.deleted","=","0"],["and","pr.discount","is null"]]}},
      /*
      ProductDiscountView: {
        columns: {partnumber:true, validfrom:true, curr:true, discount:true},
        label: getText("discount_view,
        fields: {
          partnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("product_partnumber, sqlstr:'p.partnumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_description, sqlstr:'p.description '},
          unit: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("product_unit, sqlstr:'p.unit '},
          vendor: {fieldtype:'bool', wheretype:'where', orderby:3, 
            label:getText("price_vendor, sqlstr:'pr.vendorprice '},
          custname: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("price_custname, sqlstr:'c.custname '},
          validfrom: {fieldtype:'date', wheretype:'where', orderby:5, 
            label:getText("price_validfrom, sqlstr:'pr.validfrom '},
          validto: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("price_validto, sqlstr:'pr.validto '},
          curr: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("price_curr, sqlstr:'pr.curr '},
          calcmode: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("price_calcmode, sqlstr:'case when ms.msg is null then g.description else ms.msg end '},
          qty: {fieldtype:'float', wheretype:'where', orderby:9, 
            label:getText("price_qty, sqlstr:'pr.qty '},
          pricevalue: {fieldtype:'float', wheretype:'where', orderby:10, 
            label:getText("price_limit, sqlstr:'pr.pricevalue '},
          discount: {fieldtype:'float', wheretype:'where', orderby:11, 
            label:getText("price_discount, sqlstr:'pr.discount '}
        },
        sql: "select {CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id, \
            p.partnumber, p.description, p.unit, pr.vendorprice as export_vendor, c.custname, \
            {FMS_DATE}pr.validfrom {FME_DATE} as validfrom, {FMS_DATE}pr.validto {FME_DATE} as validto, \
            pr.curr, 'discount' as form, pr.id as form_id, \
            case when ms.msg is null then g.description else ms.msg end as calcmode, \
            "{FMS_FLOAT}pr.qty{FME_FLOAT} as qty","pr.qty as export_qty",
            pr.pricevalue, pr.discount, \
            case when pr.vendorprice = 1 then \
              '<div align=\"center\" width=\"100&#37;\"><a class=\"ui-btn ui-btn-icon-notext ui-icon-check ui-state-disabled ui-btn-b\" style=\"background-color:#838B83;border-style:none;\">YES</a></div>' \
              else \
                '<div align=\"center\" width=\"100&#37;\"><a class=\"ui-btn ui-btn-icon-notext ui-icon-delete ui-state-disabled ui-btn-b\" style=\"background-color:#EEE8CD;border-style:none;\">NO</a></div>' end as vendor \
          from price pr \
          inner join product p on pr.product_id = p.id \
          left join groups g on pr.calcmode = g.id \
            left join ui_message ms on ms.fieldname = g.groupvalue and  ms.secname = 'calcmode' \
              and ms.lang = (select value from fieldvalue where fieldname = 'default_lang') \
          left join link ln0 on ln0.nervatype_1 = ( \
            select id from groups where groupname='nervatype' and groupvalue='price') and ln0.ref_id_1 = pr.id \
            and ln0.deleted=0 and ln0.nervatype_2 = ( \
              select id from groups where groupname='nervatype' and groupvalue='customer') \
            left join customer c on ln0.ref_id_2 = c.id and c.deleted=0 \
          where p.deleted = 0 and pr.deleted = 0 and pr.discount is not null  "
      },*/
      
      ProductEvents: {
        columns: {partname:true, fromdate:true, subject:true},
        label: getText("event_view"),
        edit: "event",
        fields: {
          partnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          partname: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_description"), sqlstr:'p.description '},
          calnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("event_calnumber"), sqlstr:'e.calnumber'},
          eventgroup: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("event_group"), sqlstr:'eg.groupvalue'},
          fromdate: {fieldtype:'date', wheretype:'where', orderby:4, 
            label:getText("event_fromdate"), 
            sqlstr:'{FMS_DATE}e.fromdate {FME_DATE}'},
          fromtime: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("event_fromtime"), 
            sqlstr:'{FMS_TIME}e.fromdate {FME_TIME}'},
          todate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("event_todate"), 
            sqlstr:'{FMS_DATE}e.todate {FME_DATE}'},
          totime: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("event_totime"), 
            sqlstr:'{FMS_TIME}e.todate {FME_TIME}'},
          subject: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("event_subject"), sqlstr:'e.subject'},
          place: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("event_place"), sqlstr:'e.place'},
          description: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("event_description"), sqlstr:'e.description'}

        },
        sql: {
          select:["{CCS}'event'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","e.id as row_id",
            "p.partnumber","p.description as partname","e.calnumber","eg.groupvalue as eventgroup",
            "{FMS_DATE}e.fromdate {FME_DATE} as fromdate",
            "{FMS_TIME}e.fromdate {FME_TIME} as fromtime",
            "{FMS_DATE}e.todate {FME_DATE} as todate",
            "{FMS_TIME}e.todate {FME_TIME} as totime",
            "e.subject","e.place","e.description"],
          from:"event e",
          inner_join:["product p","on",["e.ref_id","=","p.id"]],
          left_join:["groups eg","on",["e.eventgroup","=","eg.id"]],
          where:[["e.deleted","=","0"],["and","p.deleted","=","0"], 
            ["and","e.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'product'"]]}]]]}}};
    },
    
    project: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]]]}},
        
      ProjectView: {
        columns: {pronumber:true, description:true, startdate:true},
        label: getText("project_view"),
        fields: {
          pronumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("project_description"), sqlstr:'p.description '},
          customer: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("project_customer"), sqlstr:'c.custname '},
          startdate: {fieldtype:'date', wheretype:'where', orderby:3, 
            label:getText("project_startdate"), sqlstr:'cast(p.startdate as date) '},
          enddate: {fieldtype:'date', wheretype:'where', orderby:4, 
            label:getText("project_enddate"), sqlstr:'cast(p.enddate as date) '},
          inactive: {fieldtype:'bool', wheretype:'where', orderby:5, 
            label:getText("project_inactive"), sqlstr:'p.inactive'},
          notes: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("project_notes"), sqlstr:'p.notes '}
        },
        sql: {
          select:["{CCS}'project'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","p.id as row_id",
            "p.pronumber","p.description","c.custname as export_customer",
            "{CCS}'customer//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as customer",
            "{FMS_DATE}p.startdate {FME_DATE} as startdate","{FMS_DATE}p.enddate {FME_DATE} as enddate",
            "p.inactive","p.notes"],
          from:"project p",
          left_join:["customer c","on",["p.customer_id","=","c.id"]],
          where:[["p.deleted","=","0"]]}},
      
      ProjectFieldsView: {
        columns: {pronumber:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        fields: {
          pronumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber'},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("project_description"), sqlstr:'p.description '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'project'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","fv.id as row_id",
            "p.pronumber","p.description","fv.fieldname","df.description as fielddef","'fieldvalue' as form", "fg.groupvalue as fieldtype",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"], 
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]], 
              ["and","df.visible","=","1"],["and","df.deleted","=","0"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["project p","on",["fv.ref_id","=","p.id"]]],
          left_join:[
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and","p.deleted","=","0"]]}},
      
      ProjectContactView: {
        columns: {pronumber:true, firstname:true, surname:true, phone:true},
        label: getText("contact_view"),
        fields: {
          pronumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("project_description"), sqlstr:'p.description '},
          firstname: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("contact_firstname"), sqlstr:'co.firstname'},
          surname: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("contact_surname"), sqlstr:' co.surname'},
          status: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("contact_status"), sqlstr:' co.status'},
          phone: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("contact_phone"), sqlstr:' co.phone'},
          fax: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("contact_fax"), sqlstr:' co.fax'},
          mobil: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("contact_mobil"), sqlstr:' co.mobil'},
          email: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("contact_email"), sqlstr:' co.email'},
          notes: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("contact_notes"), sqlstr:' co.notes'}
          },
        sql: {
          select:["{CCS}'project'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","co.id as row_id",
            "p.pronumber","p.description","co.firstname","co.surname",
            "co.status","co.phone","co.fax","co.mobil","co.email","co.notes"],
          from:"contact co",
          inner_join:["project p","on",["co.ref_id","=","p.id"]],
          where:[["co.deleted","=","0"],["and","p.deleted","=","0"], 
            ["and","co.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]]]}},
      
      ProjectAddressView: {
        columns: {pronumber:true, city:true, street:true},
        label: getText("address_view"),
        fields: {
          pronumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("project_description"), sqlstr:'p.description '},
          country: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("address_country"), sqlstr:'a.country'},
          state: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("address_state"), sqlstr:'a.state'},
          zipcode: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("address_zipcode"), sqlstr:'a.zipcode '},
          city: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("address_city"), sqlstr:'a.city'},
          street: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("address_street"), sqlstr:'a.street'},
          notes: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("address_notes"), sqlstr:'a.notes'}
        },
        sql: {
          select:["{CCS}'project'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","a.id as row_id",
            "p.pronumber","p.description","a.country","a.state","a.zipcode",
            "a.city","a.street","a.notes"],
          from:"address a",
          inner_join:["project p","on",["a.ref_id","=","p.id"]],
          where:[["a.deleted","=","0"],["and","p.deleted","=","0"], 
            ["and","a.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]]]}},
      
      ProjectEvents: {
        columns: {pronumber:true, fromdate:true, subject:true},
        label: getText("event_view"),
        edit: "event",
        fields: {
          pronumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber '},
          pdescription: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("project_description"), sqlstr:'p.description '},
          calnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("event_calnumber"), sqlstr:'e.calnumber'},
          eventgroup: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("event_group"), sqlstr:'eg.groupvalue'},
          fromdate: {fieldtype:'date', wheretype:'where', orderby:4, 
            label:getText("event_fromdate"), 
            sqlstr:'{FMS_DATE}e.fromdate {FME_DATE}'},
          fromtime: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("event_fromtime"), 
            sqlstr:'{FMS_TIME}e.fromdate {FME_TIME}'},
          todate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("event_todate"), 
            sqlstr:'{FMS_DATE}e.todate {FME_DATE}'},
          totime: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("event_totime"), 
            sqlstr:'{FMS_TIME}e.todate {FME_TIME}'},
          subject: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("event_subject"), sqlstr:'e.subject'},
          place: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("event_place"), sqlstr:'e.place'},
          description: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("event_description"), sqlstr:'e.description'}

        },
        sql: {
          select:["{CCS}'event'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","e.id as row_id",
            "p.pronumber","p.description as pdescription","e.calnumber","eg.groupvalue as eventgroup",
            "{FMS_DATE}e.fromdate {FME_DATE} as fromdate",
            "{FMS_TIME}e.fromdate {FME_TIME} as fromtime",
            "{FMS_DATE}e.todate {FME_DATE} as todate",
            "{FMS_TIME}e.todate {FME_TIME} as totime",
            "e.subject","e.place","e.description"],
          from:"event e",
          inner_join:["project p","on",["e.ref_id","=","p.id"]],
          left_join:["groups eg","on",["e.eventgroup","=","eg.id"]],
          where:[["e.deleted","=","0"],["and","p.deleted","=","0"],
            ["and","e.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]]]}},
      
      ProjectTrans: {
        columns: {pronumber:true, transtype:true, transnumber:true, transdate:true},
        label: getText("document_view"),
        fields: {
          pronumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("project_description"), sqlstr:'p.description '},
          transtype: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("document_direction"), 
            sqlstr:'case when msd.msg is null then dg.groupvalue else msd.msg end '},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          transdate: {fieldtype:'date', wheretype:'where', orderby:5, 
            label:getText("document_transdate2"), sqlstr:'t.transdate '},
          curr: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("document_curr"), sqlstr:'t.curr '},
          amount: {fieldtype:'float', wheretype:'having', orderby:7, 
            label:getText("item_amount"), sqlstr:'sum(i.amount) '},
          custname: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("customer_custnumber"), sqlstr:'c.custname '}

        },
        sql: {
          select:["{CCS}'project'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id","t.id as row_id",
            "p.pronumber","p.description",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when msd.msg is null then dg.groupvalue else msd.msg end as direction",
            "t.transnumber as export_transnumber","t.transdate","t.curr",
            "{FMS_FLOAT}sum(i.amount){FME_FLOAT} as amount","sum(i.amount) as export_amount",
            "c.custname",
            "{CCS}'trans/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as transnumber"],
          from:"project p", 
          inner_join:[
            ["trans t","on",["p.id","=","t.project_id"]],
            ["groups tg","on",["t.transtype","=","tg.id"]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["item i","on",[["t.id","=","i.trans_id"],["and","i.deleted","=","0"]]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"],
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"],
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["customer c","on",["t.customer_id","=","c.id"]]],
          where:[["p.deleted","=","0"],["and", [["t.deleted","=","0"],
            ["or",[["tg.groupvalue","=","'invoice'"],["and","dg.groupvalue","=","'out'"]]],
              ["or",[["tg.groupvalue","=","'receipt'"],["and","dg.groupvalue","=","'out'"]]]]]],
          group_by:["p.id","p.pronumber","p.description","tg.groupvalue","dg.groupvalue","mst.msg, msd.msg",
            "t.id","t.transnumber","t.transdate","c.custname","t.curr"], 
          having:[["sum(i.amount)","<>","0"]]}}};
    },
    
    rate: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'rate'"]]}]]]}},

      RateView: {
        columns: {ratetype:true, ratedate:true, curr:true, ratevalue:true},
        label: getText("rate_view"),
        actions_new: {
          action: "loadEditor", ntype: "rate", ttype: null},
        fields: {
          ratetype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("rate_ratetype"), sqlstr:'rtype.groupvalue '},
          ratedate: {fieldtype:'date', wheretype:'where', orderby:1, 
            label:getText("rate_ratedate"), sqlstr:'r.ratedate '},
          curr: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("rate_curr"), sqlstr:'r.curr '},
          ratevalue: {fieldtype:'float', wheretype:'where', orderby:3, 
            label:getText("rate_ratevalue"), sqlstr:'r.ratevalue '},
          rategroup: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("rate_rategroup"), sqlstr:'rgroup.groupvalue '},
          planumber: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("rate_planumber"), sqlstr:'p.planumber'}
        },
        sql:{
          select:["{CCS}'rate'{SEP}'//'{SEP} {CAS_TEXT}r.id {CAE_TEXT}{CCE} as id","r.id as row_id",
            "rtype.groupvalue as ratetype","{FMS_DATE}r.ratedate {FME_DATE} as ratedate","r.curr",
            "{FMS_FLOAT}r.ratevalue{FME_FLOAT} as ratevalue","r.ratevalue as export_ratevalue",
            "rgroup.groupvalue as rategroup","p.planumber as export_planumber",
            "{CCS}'place//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as planumber"],
            from:"rate r",
            inner_join:["groups as rtype","on",["r.ratetype","=","rtype.id"]],
            left_join:[
              ["groups as rgroup","on",["r.rategroup","=","rgroup.id"]],
              ["place p","on",["r.place_id","=","p.id"]]],
            where:[["r.deleted","=","0"]]}}};
    },
    
    tool: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'tool'"]]}]]]}},

      ToolView: {
        columns: {serial:true, description:true, product:true},
        label: getText("tool_view"),
        fields: {
          serial: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("tool_serial"), sqlstr:'t.serial '},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("tool_description"), sqlstr:'t.description '},
          product: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("tool_product"), sqlstr:'p.description '},
          toolgroup: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("tool_toolgroup"), sqlstr:'tg.groupvalue '},
          //state: {fieldtype:'string', wheretype:'where', orderby:4, 
          //  label:getText("tool_state, sqlstr:'ssel.state '},
          inactive: {fieldtype:'bool', wheretype:'where', orderby:5, 
            label:getText("tool_inactive"), sqlstr:'t.inactive'},
          notes: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("tool_notes"), sqlstr:'t.notes '}
        },
        sql: {
          select:["{CCS}'tool'{SEP}'//'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","t.id as row_id",
            "t.serial","t.description","p.description as export_product",
            "{CCS}'product//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as product",
            //"case when ssel.state is null then '***************' else ssel.state end  as state",
            "tg.groupvalue as toolgroup","t.inactive","t.notes"],
          from:"tool t",
          inner_join:["product p","on",["t.product_id","=","p.id"]],
          left_join:[
            ["groups tg","on",["t.toolgroup","=","tg.id"]],
            [[[{select:["mv.tool_id"//, 
              //"case when t.direction=dir_in.id then owncust.custname "+
              //"when c.custname is not null then c.custname "+
              //"when e.empnumber is not null then e.empnumber "+
              //"else ltc.custname end as state"
              ],
              from:[
                "movement mv"//,
                //[[{select:["id"], from:"groups", 
                //  where:[["groupname","=","'direction'"],["and","groupvalue","=","'in'"]]}],"dir_in,"],
                //[[{select:["custname"], from:"customer", 
                //  where:["id","in",[
                //    {select:["min(customer.id)"], from:"customer",
                //    inner_join:["groups","on",[["customer.custtype","=","groups.id"],
                //      ["and","groups.groupvalue","=","'own'"]]]}]]}],"owncust"]
              ],
              inner_join:["trans t","on",["mv.trans_id","=","t.id"]],
              left_join:[
                ["customer c","on",["t.customer_id","=","c.id"]],
                ["employee e","on",["t.employee_id","=","e.id"]],
                ["link lnk","on",[["t.id","=","lnk.ref_id_1"],["and","lnk.deleted","=","0"],
                  ["and","lnk.nervatype_1","=",[{select:["id"], from:"groups", 
                    where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]],
                  ["and","lnk.nervatype_2","=",[{select:["id"], from:"groups", 
                    where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]]]],
                ["trans lt","on",["lnk.ref_id_2","=","lt.id"]], 
                ["customer ltc","on",["lt.customer_id","=","ltc.id"]]],
              where:["mv.id","in",[{
                select:["max(id)"], from:"movement mv",
                inner_join:[[[{
                  select:["tool_id","max(shippingdate) as ldate"], from:"movement mv",
                  inner_join:["trans t","on",["mv.trans_id","=","t.id"]],
                  where:[["mv.deleted","=","0"],["and","tool_id","is not","null"],["and","t.deleted","=","0"], 
                    ["and","{CAS_DATE}shippingdate {CAE_DATE}","<=","{CUR_DATE}"]],
                  group_by:["tool_id"]}],"lst_date"],"on",[["mv.tool_id","=","lst_date.tool_id"],["and","mv.shippingdate","=","lst_date.ldate"]]],
                group_by:["mv.tool_id"]}]]}],"ssel"],"on",["t.id","=","ssel.tool_id"]]],
          where:[["t.deleted","=","0"],["and",["t.toolgroup","not in",[
            {select:["id"], from:"groups", where:[["groupname","=","'toolgroup'"],
              ["and","groupvalue","=","'printer'"]]}],
              ["or","t.toolgroup","is null"]]]]}},
      
      ToolFieldsView: {
        columns: {serial:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        fields: {
          serial: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("tool_serial"), sqlstr:'t.serial'},
          description: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("tool_description"), sqlstr:'t.description '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'tool'{SEP}'//'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","fv.id as row_id",
            "fv.fieldname","t.serial","t.description","df.description as fielddef","'fieldvalue' as form", "fg.groupvalue as fieldtype",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"], 
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'tool'"]]}]], 
              ["and","df.visible","=","1"],["and","df.deleted","=","0"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["tool t","on",["fv.ref_id","=","t.id"]]],
          left_join:[
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and","t.deleted","=","0"]]}
      },
      
      ToolEvents: {
        columns: {serial:true, fromdate:true, subject:true},
        label: getText("event_view"),
        edit: "event",
        fields: {
          serial: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("tool_serial"), sqlstr:'t.serial'},
          pdescription: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("tool_description"), sqlstr:'t.description '},
          calnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("event_calnumber"), sqlstr:'e.calnumber'},
          eventgroup: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("event_group"), sqlstr:'eg.groupvalue'},
          fromdate: {fieldtype:'date', wheretype:'where', orderby:4, 
            label:getText("event_fromdate"), 
            sqlstr:'{FMS_DATE}e.fromdate {FME_DATE}'},
          fromtime: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("event_fromtime"), 
            sqlstr:'{FMS_TIME}e.fromdate {FME_TIME}'},
          todate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("event_todate"), 
            sqlstr:'{FMS_DATE}e.todate {FME_DATE}'},
          totime: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("event_totime"), 
            sqlstr:'{FMS_TIME}e.todate {FME_TIME}'},
          subject: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("event_subject"), sqlstr:'e.subject'},
          place: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("event_place"), sqlstr:'e.place'},
          description: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("event_description"), sqlstr:'e.description'}

        },
        sql: {
          select:["{CCS}'event'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id","e.id as row_id",
            "t.serial","t.description as pdescription","e.calnumber","eg.groupvalue as eventgroup",
            "{FMS_DATE}e.fromdate {FME_DATE} as fromdate",
            "{FMS_TIME}e.fromdate {FME_TIME} as fromtime",
            "{FMS_DATE}e.todate {FME_DATE} as todate",
            "{FMS_TIME}e.todate {FME_TIME} as totime",
            "e.subject","e.place","e.description"],
          from:"event e",
          inner_join:["tool t","on",["e.ref_id","=","t.id"]],
          left_join:["groups eg","on",["e.eventgroup","=","eg.id"]],
          where:[["e.deleted","=","0"],["and","t.deleted","=","0"],
            ["and","e.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'tool'"]]}]]]}}};
    },
    
    transitem: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],
            ["and",[["df.visible","=","1"],["or","df.fieldname",
            "in",[[],"'trans_wsdistance'","'trans_wsrepair'","'trans_wstotal'","'trans_wsnote'",
              "'trans_reholiday'","'trans_rebadtool'","'trans_reother'","'trans_rentnote'"]]]], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]]]}},
      
      TransItemHeadView: {
        columns: {transtype:true, transnumber:true, transdate:true, custname: true},
        label: getText("document_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:'case when msd.msg is null then dg.groupvalue else msd.msg end '},
          transcast: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transcast"), 
            sqlstr:'case when msc.msg is null then fv.value else msc.msg end '},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          ref_transnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("document_ref_transnumber"), sqlstr:'t.ref_transnumber '},
          crdate: {fieldtype:'date', wheretype:'where', orderby:5, 
            label:getText("document_crdate"), sqlstr:'t.crdate '},
          transdate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("document_transdate"), sqlstr:'t.transdate '},
          duedate: {fieldtype:'date', wheretype:'where', orderby:7, 
            label:getText("document_duedate"), 
            sqlstr:'{FMS_DATE}t.duedate {FME_DATE} '},
          custname: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          empnumber: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("employee_empnumber"), sqlstr:'e.empnumber '},
          department: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("document_department"), sqlstr:'deg.groupvalue '},
          pronumber: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("project_pronumber"), sqlstr:'p.pronumber '},
          paidtype: {fieldtype:'string', wheretype:'where', orderby:12, 
            label:getText("document_paidtype"), 
            sqlstr:'case when msp.msg is null then ptg.groupvalue else msp.msg end '},
          curr: {fieldtype:'string', wheretype:'where', orderby:13, 
            label:getText("document_curr"), sqlstr:'t.curr '},
          netamount: {fieldtype:'float', wheretype:'where', orderby:14, 
            label:getText("item_netamount"), sqlstr:'irow.netamount '},
          vatamount: {fieldtype:'float', wheretype:'where', orderby:15, 
            label:getText("item_vatamount"), sqlstr:'irow.vatamount '},
          amount: {fieldtype:'float', wheretype:'where', orderby:16, 
            label:getText("item_amount"), sqlstr:'irow.amount '},
          paid: {fieldtype:'bool', wheretype:'where', orderby:17, 
            label:getText("document_paid"), sqlstr:'t.paid '},
          acrate: {fieldtype:'float', wheretype:'where', orderby:18, 
            label:getText("document_acrate"), sqlstr:'t.acrate '},
          notes: {fieldtype:'string', wheretype:'where', orderby:19, 
            label:getText("document_notes"), sqlstr:'t.notes '},
          intnotes: {fieldtype:'string', wheretype:'where', orderby:20, 
            label:getText("document_intnotes"), sqlstr:'t.intnotes '},
          transtate: {fieldtype:'string', wheretype:'where', orderby:21, 
            label:getText("document_transtate"), 
            sqlstr:'case when mss.msg is null then sg.groupvalue else mss.msg end '},
          closed: {fieldtype:'bool', wheretype:'where', orderby:22, 
            label:getText("document_closed"), sqlstr:'t.closed '},
          deleted: {fieldtype:'bool', wheretype:'where', orderby:23, 
            label:getText("document_deleted"), sqlstr:'t.deleted '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","t.id as row_id",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when msd.msg is null then dg.groupvalue else msd.msg end as direction",
            "case when msc.msg is null then fv.value else msc.msg end as transcast",
            "t.transnumber","t.ref_transnumber",
            "{FMS_DATE}t.crdate {FME_DATE} as crdate",
            "{FMS_DATE}t.transdate {FME_DATE} as transdate",
            "{FMS_DATE}t.duedate {FME_DATE} as duedate",
            "{CCS}'customer//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as custname", "c.custname as export_custname",
            "{CCS}'employee//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as empnumber", "e.empnumber as export_empnumber",
            "deg.groupvalue as department",
            "{CCS}'project//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as pronumber", "p.pronumber as export_pronumber",
            "case when msp.msg is null then ptg.groupvalue else msp.msg end as paidtype","t.curr",
            "case when irow.netamount is null then '0' else {FMS_FLOAT}irow.netamount{FME_FLOAT} end as netamount",
            "case when irow.netamount is null then 0 else irow.netamount end as export_netamount",
            "case when irow.vatamount is null then '0' else {FMS_FLOAT}irow.vatamount{FME_FLOAT} end as vatamount",
            "case when irow.vatamount is null then 0 else irow.vatamount end as export_vatamount",
            "case when irow.amount is null then '0' else {FMS_FLOAT}irow.amount{FME_FLOAT} end as amount",
            "case when irow.amount is null then 0 else irow.amount end as export_amount",
            "t.paid","t.acrate","t.notes","t.intnotes",
            "case when mss.msg is null then sg.groupvalue else mss.msg end as transtate",
            "t.closed","t.deleted"],
          from:"trans t",
          inner_join:[
            ["groups tg","on",[["t.transtype","=","tg.id"],
              ["and","tg.groupvalue","in",[[],"'invoice'","'receipt'","'order'","'offer'","'worksheet'","'rent'"]]]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["groups ptg","on",["t.paidtype","=","ptg.id"]],
            ["groups sg","on",["t.transtate","=","sg.id"]]],
          left_join:[
            ["customer c","on",["t.customer_id","=","c.id"]],
            ["employee e","on",["t.employee_id","=","e.id"]],
            ["groups deg","on",["t.department","=","deg.id"]],
            ["project p","on",["t.project_id","=","p.id"]],
            ["fieldvalue fv","on",[["t.id","=","fv.ref_id"],["and","fv.fieldname","=","'trans_transcast'"]]],
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"],
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msp","on",[["msp.fieldname","=","ptg.groupvalue"],["and","msp.secname","=","'paidtype'"], 
              ["and","msp.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message mss","on",[["mss.fieldname","=","sg.groupvalue"],["and","mss.secname","=","'transtate'"], 
              ["and","mss.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msc","on",[["msc.fieldname","=","fv.value"],["and","msc.secname","=","'trans_transcast'"], 
              ["and","msc.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            [[[{select:["trans_id","sum(netamount) as netamount","sum(vatamount) as vatamount","sum(amount) as amount"],
              from:"item", where:["deleted","=","0"], group_by:["trans_id"]}],"irow"],"on",["t.id","=","irow.trans_id"]]],
          where:[[[["t.deleted","=","0"],["or",[["tg.groupvalue","=","'invoice'"],["and","dg.groupvalue","=","'out'"]]],
            ["or",[["tg.groupvalue","=","'receipt'"],["and","dg.groupvalue","=","'out'"]]]]]]}},
      
      TransItemFieldsView: {
        columns: {transtype:true, transnumber:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:'case when msd.msg is null then dg.groupvalue else msd.msg end '},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","fv.id as row_id",
            "'fieldvalue' as form","fv.fieldname","fg.groupvalue as fieldtype",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when msd.msg is null then dg.groupvalue else msd.msg end as direction", "t.transnumber",
            "df.description as fielddef",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"],
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]], 
              ["and","df.deleted","=","0"],
              ["and",[["df.visible","=","1"],["or","df.fieldname","in",[[],"'trans_reholiday'","'trans_rebadtool'","'trans_reother'",
                "'trans_rentnote'","'trans_wsdistance'","'trans_wsrepair'","'trans_wstotal'","'trans_wsnote'"]]]]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["trans t","on",["fv.ref_id","=","t.id"]],
            ["groups tg","on",[["t.transtype","=","tg.id"],
              ["and","tg.groupvalue","in",[[],"'invoice'","'receipt'","'order'","'offer'","'worksheet'","'rent'"]]]],
            ["groups dg","on",["t.direction","=","dg.id"]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and",[["t.deleted","=","0"],
            ["or",[["tg.groupvalue","=","'invoice'"],["and","dg.groupvalue","=","'out'"]]],
            ["or",[["tg.groupvalue","=","'receipt'"],["and",["dg.groupvalue","=","'out'"]]]]]]]}},
      
      TransItemRowView: {
        columns: {transnumber:true, curr:true, description:true, amount:true},
        label: getText("item_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:'case when msd.msg is null then dg.groupvalue else msd.msg end '},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          transdate: {fieldtype:'date', wheretype:'where', orderby:3, 
            label:getText("item_transdate"), sqlstr:'t.transdate '},
          partnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("item_description"), sqlstr:'i.description '},
          unit: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("item_unit"), sqlstr:'i.unit '},
          curr: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("document_curr"), sqlstr:'t.curr '},
          qty: {fieldtype:'float', wheretype:'where', orderby:8, 
            label:getText("item_qty"), sqlstr:'i.qty '},
          fxprice: {fieldtype:'float', wheretype:'where', orderby:9, 
            label:getText("item_fxprice"), sqlstr:'i.fxprice '},
          netamount: {fieldtype:'float', wheretype:'where', orderby:10, 
            label:getText("item_netamount"), sqlstr:'i.netamount '},
          discount: {fieldtype:'float', wheretype:'where', orderby:11, 
            label:getText("item_discount"), sqlstr:'i.discount '},
          taxcode: {fieldtype:'string', wheretype:'where', orderby:12, 
            label:getText("item_taxcode"), sqlstr:'tax.taxcode '},
          vatamount: {fieldtype:'float', wheretype:'where', orderby:13, 
            label:getText("item_vatamount"), sqlstr:'i.vatamount '},
          amount: {fieldtype:'float', wheretype:'where', orderby:14, 
            label:getText("item_amount"), sqlstr:'i.amount '},
          deposit: {fieldtype:'bool', wheretype:'where', orderby:15, 
            label:getText("item_deposit"), sqlstr:'i.deposit '},
          actionprice: {fieldtype:'bool', wheretype:'where', orderby:16, 
            label:getText("item_actionprice"), sqlstr:'i.actionprice '},
          ownstock: {fieldtype:'float', wheretype:'where', orderby:17, 
            label:getText("item_ownstock"), sqlstr:'i.ownstock '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","i.id as row_id",
            "'item' as form", "i.id as form_id",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when msd.msg is null then dg.groupvalue else msd.msg end as direction", "t.transnumber",
            "{FMS_DATE}t.transdate {FME_DATE} as transdate", "t.curr",
            "{CCS}'product//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as partnumber", "p.partnumber as export_partnumber",
            "i.description","i.unit",
            "{FMS_FLOAT}i.qty{FME_FLOAT} as qty","i.qty as export_qty",
            "{FMS_FLOAT}i.fxprice{FME_FLOAT} as fxprice","i.fxprice as export_fxprice",
            "{FMS_FLOAT}i.netamount{FME_FLOAT} as netamount","i.netamount as export_netamount",
            "i.discount as discount","tax.taxcode",
            "{FMS_FLOAT}i.vatamount{FME_FLOAT} as vatamount","i.vatamount as export_vatamount",
            "{FMS_FLOAT}i.amount{FME_FLOAT} as amount","i.amount as export_amount",
            "i.deposit","i.actionprice",
            "{FMS_FLOAT}i.ownstock{FME_FLOAT} as ownstock","i.ownstock as export_ownstock"],
          from:"item i", 
          inner_join:[
            ["trans t","on",["i.trans_id","=","t.id"]],
            ["groups tg","on",["t.transtype","=","tg.id"]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["product p","on",["i.product_id","=","p.id"]],
            ["tax","on",["i.tax_id","=","tax.id"]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]]],
          where:[["i.deleted","=","0"],["and",[["t.deleted","=","0"],
            ["or",[["tg.groupvalue","=","'invoice'"],["and","dg.groupvalue","=","'out'"]]],
            ["or",[["tg.groupvalue","=","'receipt'"],["and","dg.groupvalue","=","'out'"]]]]]]}}};
    },
    
    transmovement: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]]]}},
      
      InventoryView: {
        columns: {warehouse:true, partnumber:true, unit:true, sqty:true},
        label: getText("inventory_view"),
        readonly: true,
        fields: {
          warehouse: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("inventory_warehouse"), sqlstr:'pl.description '},
          partnumber: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          unit: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("product_unit"), sqlstr:'p.unit '},
          sqty: {fieldtype:'float', wheretype:'having', orderby:3, 
            label:getText("inventory_sqty"), sqlstr:'sum(mv.qty) '},
          posdate: {fieldtype:'date', wheretype:'having', orderby:4, 
            label:getText("inventory_posdate"), sqlstr:'{CAS_DATE}max(mv.shippingdate){CAE_DATE} '},
          description: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("product_description"), sqlstr:'p.description '}
        },
        sql: {
          select:["{CCS} {CAS_TEXT}pl.id {CAE_TEXT}{SEP}'/'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{SEP}'/'{SEP} mv.notes{CCE} as row_id",
            "pl.description as warehouse","p.partnumber as export_partnumber",
            "{CCS}'product//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as partnumber",
            "p.description","p.unit","mv.notes as pgroup",
            "{FMS_FLOAT}sum(mv.qty){FME_FLOAT} as sqty","sum(mv.qty) as export_sqty",
            "{FMS_DATE}max(mv.shippingdate) {FME_DATE} as posdate"],
          from:"movement mv", 
          inner_join:[
            ["groups g","on",[["mv.movetype","=","g.id"],["and","g.groupvalue","=","'inventory'"]]],
            ["place pl","on",["mv.place_id","=","pl.id"]],
            ["product p","on",["mv.product_id","=","p.id"]],
            ["trans t","on",[["mv.trans_id","=","t.id"],["and","t.deleted","=","0"]]],
            ["groups tg","on",["t.transtype","=","tg.id"]],
            ["groups dg","on",["t.direction","=","dg.id"]]],
          where:[["mv.deleted","=","0"],["and","t.deleted","=","0"]],
          group_by:["pl.id","pl.description","p.id","p.partnumber","p.description","p.unit","mv.notes"],
          having:[["sum(mv.qty)","<>","0"]]}},
      
      MovementView: {
        columns: {transtype:true, shippingdate:true, partnumber:true, qty:true},
        label: getText("movement_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:'case when msd.msg is null then dg.groupvalue else msd.msg end '},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          shippingdate: {fieldtype:'date', wheretype:'where', orderby:3, 
            label:getText("movement_shippingdate"), sqlstr:'{CAS_DATE}mt.shippingdate{CAE_DATE} '},
          warehouse: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("inventory_warehouse"), sqlstr:'pt.description  '},
          partnumber: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber  '},
          description: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("product_description"), sqlstr:'p.description  '},
          unit: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("product_unit"), sqlstr:'p.unit  '},
          pgroup: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("movement_batchnumber"), sqlstr:'mt.notes '},
          qty: {fieldtype:'float', wheretype:'where', orderby:9, 
            label:getText("movement_qty"), sqlstr:'mt.qty '},
          refnumber: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("movement_refnumber"), 
            sqlstr:'case when it.transnumber is null then '+
              'case when vt1.transnumber is null then case when vt2.transnumber is null then '+
              'tl.transnumber end     else vt1.transnumber end else it.transnumber end '},
          refcustomer: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("movement_refcustomer"), 
            sqlstr:'case when c1.custname is null then c2.custname else c1.custname end '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP}"+
            "case when it.id is null then {CAS_TEXT}t.id {CAE_TEXT} else {CAS_TEXT}it.id {CAE_TEXT} end{CCE} as id",
            "mt.id as row_id",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when msd.msg is null then dg.groupvalue else msd.msg end as direction",
            "t.transnumber","{FMS_DATE}mt.shippingdate {FME_DATE} as shippingdate",
            "pt.description as warehouse",
            "{CCS}'product//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as partnumber",
            "p.partnumber as export_partnumber","p.description","p.unit","mt.notes as pgroup","mt.qty",
            "{FMS_FLOAT}mt.qty{FME_FLOAT} as qty","mt.qty as export_qty",
            "case when it.transnumber is null then "+
              "case when vt1.transnumber is null then "+
                "case when vt2.transnumber is null then tl.transnumber end "+
              "else vt1.transnumber end else it.transnumber end as refnumber",
            "case when c1.custname is null then c2.custname else c1.custname end as refcustomer"],
          from:"movement mt",
          inner_join:[
            ["trans t","on",["mt.trans_id","=","t.id"]], 
            ["groups gm","on",[["mt.movetype","=","gm.id"],["and","gm.groupvalue","=","'inventory'"]]], 
            ["groups tg","on",["t.transtype","=","tg.id"]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["product p","on",["mt.product_id","=","p.id"]],
            ["place pt","on",["mt.place_id","=","pt.id"]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["link iln","on",[["iln.nervatype_1","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'movement'"]]}]],
              ["and","iln.ref_id_1","=","mt.id"],
              ["and","iln.nervatype_2","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'item'"]]}]]]],
            ["item i","on",["iln.ref_id_2","=","i.id"]],
            ["trans it","on",["i.trans_id","=","it.id"]],
            ["customer c1","on",["it.customer_id","=","c1.id"]],
            ["link tln","on",["tln.nervatype_1","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}],
              ["and","tln.ref_id_1","=","mt.trans_id"],["and","tln.nervatype_2","=",[{select:["id"], 
                from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]],
                ["and","tln.deleted","=","0"]]],
            ["trans tl","on",["tln.ref_id_2","=","tl.id"]], 
            ["customer c2","on",["tl.customer_id","=","c2.id"]],
            ["link pln1","on",[["pln1.nervatype_1","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'movement'"]]}]],
              ["and","pln1.ref_id_1","=","mt.id"],["and","pln1.nervatype_2","=",[{select:["id"], 
                from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'movement'"]]}]]]],
            ["movement mv1","on",["pln1.ref_id_2","=","mv1.id"]], 
            ["trans vt1","on",["mv1.trans_id","=","vt1.id"]],
            ["link pln2","on",["pln2.nervatype_1","=", [{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'movement'"]]}],
              ["and","pln2.ref_id_2","=","mt.id"],["and","pln2.nervatype_2","=",[{select:["id"], 
                from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'movement'"]]}]]]],
            ["movement mv2","on",["pln2.ref_id_1","=","mv2.id"]], 
            ["trans vt2","on",["mv2.trans_id","=","vt2.id"]]],
          where:[["mt.deleted","=","0"],["and","t.deleted","=","0"]]}},
      
      ToolMovement: {
        columns: {transnumber:true, direction:true, shippingdate:true, serial:true},
        label: getText("toolmovement_view"),
        edit: "trans",
        fields: {
          transnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          crdate: {fieldtype:'date', wheretype:'where', orderby:1, 
            label:getText("document_crdate"), sqlstr:'t.crdate '},
          direction: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_direction"), 
            sqlstr:'case when msd.msg is null then dg.groupvalue else msd.msg end '},
          refnumber: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("movement_refnumber"), sqlstr:'lt.transnumber '},
          empnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("employee_empnumber"), sqlstr:'e.empnumber '},
          custname: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("customer_custname"), sqlstr:'c.custname '},
          shippingdate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("movement_shippingdate"), 
            sqlstr:'{CAS_DATE}mv.shippingdate {CAE_DATE} '},
          serial: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("tool_serial"), sqlstr:'tl.serial '},
          description: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("tool_description"), sqlstr:'tl.description '},
          mvnotes: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("movement_mvnotes"), sqlstr:'mv.notes '},
          closed: {fieldtype:'bool', wheretype:'where', orderby:10, 
            label:getText("document_closed"), sqlstr:'t.closed '},
          transtate: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("document_transtate"), 
            sqlstr:'case when mss.msg is null then sg.groupvalue else mss.msg end '},
          notes: {fieldtype:'string', wheretype:'where', orderby:12, 
            label:getText("document_notes"), sqlstr:'t.notes '},
          intnotes: {fieldtype:'string', wheretype:'where', orderby:13, 
            label:getText("document_intnotes"), sqlstr:'t.intnotes '}
        },
        sql: {
          select: ["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","mv.id as row_id",
            "t.transnumber","{FMS_DATE}t.crdate{FME_DATE} as crdate","t.crdate as export_crdate",
            "case when msd.msg is null then dg.groupvalue else msd.msg end as direction",
            "lt.transnumber as export_refnumber","e.empnumber as export_empnumber","c.custname as export_custname",
            "{CCS}'trans/'{SEP}ltg.groupvalue{SEP}'/'"+
              "{SEP} {CAS_TEXT}lt.id {CAE_TEXT}{CCE} as refnumber",
            "{CCS}'employee//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as empnumber",
            "{CCS}'customer//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as custname",
            "{FMS_DATETIME}mv.shippingdate {FME_DATETIME} as shippingdate",
            "{CCS}'tool//'{SEP} {CAS_TEXT}tl.id {CAE_TEXT}{CCE} as serial",
            "tl.serial as export_serial","tl.description","mv.notes as mvnotes",
            "case when mss.msg is null then sg.groupvalue else mss.msg end as transtate",
            "t.notes","t.intnotes","t.closed"],
          from:"trans t",
          inner_join:[
            ["movement mv","on",["t.id","=","mv.trans_id"]],
            ["tool tl","on",["mv.tool_id","=","tl.id"]],
            ["groups tg","on",[["t.transtype","=","tg.id"],["and","tg.groupvalue","=","'waybill'"]]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["groups sg","on",["t.transtate","=","sg.id"]]],
          left_join:[
            ["employee e","on",["t.employee_id","=","e.id"]],
            ["customer c","on",["t.customer_id","=","c.id"]],
            ["ui_message mss","on",[["mss.fieldname","=","sg.groupvalue"],["and","mss.secname","=","'transtate'"], 
              ["and","mss.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["link lnk","on",[["t.id","=","lnk.ref_id_1"],["and","lnk.deleted","=","0"],
              ["and","lnk.nervatype_1","=",[{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]],
              ["and","lnk.nervatype_2","=",[{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]]]], 
            ["trans lt","on",["lnk.ref_id_2","=","lt.id"]], 
            ["groups ltg","on",["lt.transtype","=","ltg.id"]]], 
          where:[["t.deleted","=","0"],["and","mv.deleted","=","0"]]}},
      
      FormulaView: {
        columns: {transnumber:true, type:true, partnumber:true, qty:true},
        label: getText("formula_view"),
        edit: "trans",
        fields: {
          transnumber: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          type: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("formula_type"), 
            sqlstr:"case when mt.groupvalue = 'head' then 'in' else 'out' end "},
          partnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("product_partnumber"), sqlstr:'p.partnumber '},
          description: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("product_description"), sqlstr:'p.description '},
          unit: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("product_unit"), sqlstr:'p.unit '},
          qty: {fieldtype:'float', wheretype:'where', orderby:5, 
            label:getText("movement_qty"), sqlstr:'mv.qty '},
          batch_no: {fieldtype:'string', wheretype:'where', orderby:6, 
            label:getText("movement_batchnumber"), sqlstr:'mv.notes '},
          planumber: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("inventory_warehouse"), sqlstr:'pl.planumber '},
          shared: {fieldtype:'bool', wheretype:'where', orderby:8, 
            label:getText("formula_shared"), sqlstr:'mv.shared '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
            "mv.id as row_id","t.transnumber",
            "case when mt.groupvalue = 'head' then 'in' else 'out' end as type",
            "{CCS}'product//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as partnumber", 
            "p.partnumber as export_partnumber","p.description","p.unit",
            "{FMS_FLOAT}mv.qty{FME_FLOAT} as qty","mv.qty as export_qty",
            "mv.notes as batch_no","pl.planumber","mv.shared"],
          from:"trans t",
          inner_join:[
            ["movement mv","on",[["t.id","=","mv.trans_id"],["and","mv.deleted","=","0"]]],
          ["groups tg","on",[["t.transtype","=","tg.id"],["and","tg.groupvalue","=","'formula'"]]],
          ["groups mt","on",["mv.movetype","=","mt.id"]],
          ["product p","on",["mv.product_id","=","p.id"]]],
          left_join:["place pl","on",["mv.place_id","=","pl.id"]],
          where:[["t.deleted","=","0"]]}},
      
      MovementFieldsView: {
        columns: {transtype:true, transnumber:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:"case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end "},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP} tg.groupvalue{SEP}'/'{SEP} case when delt.ref_id is null then {CAS_TEXT}t.id {CAE_TEXT} else {CAS_TEXT}delt.ref_id {CAE_TEXT} end{CCE} as id",
            "fv.id as row_id","fv.fieldname", "'fieldvalue' as form","fg.groupvalue as fieldtype",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end as direction", "t.transnumber",
            "df.description as fielddef",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join: [
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"],
              ["and","df.nervatype","=",[{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]], 
              ["and","df.deleted","=","0"],["and","df.visible","=","1"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["trans t","on",["fv.ref_id","=","t.id"]],
            ["groups tg","on",[["t.transtype","=","tg.id"],["and","tg.groupvalue","in",[[],"'delivery'","'inventory'","'waybill'","'production'","'formula'"]]]],
            ["groups dg","on",["t.direction","=","dg.id"]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]],
            [[[{select:["mt.trans_id","min(it.id) as ref_id"], from:"movement mt",
              inner_join:[
                ["link iln","on",
                  ["iln.nervatype_1","=",[[{select:["id"], from:"groups", 
                    where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'movement'"]]}],
                    ["and","iln.ref_id_1","=","mt.id"], 
                    ["and","iln.nervatype_2","=",[{select:["id"], from:"groups", 
                    where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'item'"]]}]]]]],
              ["item i","on",["iln.ref_id_2","=","i.id"]]],
              left_join:["trans it","on",["i.trans_id","=","it.id"]], group_by:["mt.trans_id"]}],"delt"],"on",["fv.ref_id","=","delt.trans_id"]]],
          where:[["fv.deleted","=","0"],["and","t.deleted","=","0"]]}}};
    },
    
    transpayment: () => {return {
      options: {
        deffield_sql: {
          select:["df.id","df.fieldname","g.groupvalue as fieldtype","df.description"],
          from:"deffield df", 
          inner_join:["groups g","on",["g.id","=","df.fieldtype"]],
          where: [["df.deleted","=","0"],["and","df.visible","=","1"], 
            ["and","df.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]]]}},
      
      PaymentView: {
        columns: {transnumber:true, paiddate:true, place:true, curr:true, amount:true},
        label: getText("payment_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:"case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end "},
          transcast: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transcast"), 
            sqlstr:'case when msc.msg is null then fv.value else msc.msg end '},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          ref_transnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("document_ref_transnumber"), sqlstr:'t.ref_transnumber '},
          crdate: {fieldtype:'date', wheretype:'where', orderby:5, 
            label:getText("document_crdate"), sqlstr:'t.crdate '},
          paiddate: {fieldtype:'date', wheretype:'where', orderby:6, 
            label:getText("payment_paiddate"), sqlstr:'pm.paiddate '},
          place: {fieldtype:'string', wheretype:'where', orderby:7, 
            label:getText("payment_place"), sqlstr:'pc.description '},
          curr: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("payment_curr"), sqlstr:'pc.curr '},
          amount: {fieldtype:'float', wheretype:'where', orderby:9, 
            label:getText("payment_amount"), 
            sqlstr:"case when dg.groupvalue='out' then -pm.amount else pm.amount end "},
          description: {fieldtype:'string', wheretype:'where', orderby:10, 
            label:getText("payment_description"), sqlstr:'pm.notes '},
          empnumber: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("employee_empnumber"), sqlstr:'e.empnumber '},
          transtate: {fieldtype:'string', wheretype:'where', orderby:12, 
            label:getText("document_transtate"), 
            sqlstr:'case when mss.msg is null then sg.groupvalue else mss.msg end '},
          closed: {fieldtype:'bool', wheretype:'where', orderby:13, 
            label:getText("document_closed"), sqlstr:'t.closed '},
          deleted: {fieldtype:'bool', wheretype:'where', orderby:14, 
            label:getText("document_deleted"), sqlstr:'t.deleted '},
          notes: {fieldtype:'string', wheretype:'where', orderby:15, 
            label:getText("document_notes"), sqlstr:'t.notes '},
          intnotes: {fieldtype:'string', wheretype:'where', orderby:16, 
            label:getText("document_intnotes"), sqlstr:'t.intnotes '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id","pm.id as row_id",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end as direction",
            "case when msc.msg is null then fv.value else msc.msg end as transcast",
            "t.transnumber","t.ref_transnumber","{FMS_DATE}t.crdate {FME_DATE} as crdate",
            "{FMS_DATE}pm.paiddate {FME_DATE} as paiddate",
            "{CCS}'place//'{SEP} {CAS_TEXT}pc.id {CAE_TEXT}{CCE} as place, pc.description as export_place","pc.curr",
            "{FMS_FLOAT}case when dg.groupvalue='out' then -pm.amount else pm.amount end{FME_FLOAT} as amount",
            "case when dg.groupvalue='out' then -pm.amount else pm.amount end as export_amount",
            "pm.notes as description",
            "{CCS}'employee//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as empnumber", "e.empnumber as export_empnumber",
            "case when mss.msg is null then sg.groupvalue else mss.msg end as transtate",
            "t.notes","t.intnotes","t.closed","t.deleted as deleted"],
          from:"trans t",
          inner_join:[
            ["groups tg","on",["t.transtype","=","tg.id"]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["place pc","on",["t.place_id","=","pc.id"]],
            ["groups sg","on",["t.transtate","=","sg.id"]],
            ["payment pm","on",[["t.id","=","pm.trans_id"],["and","pm.deleted","=","0"]]]],
          left_join:[
            ["employee e","on",["t.employee_id","=","e.id"]],
            ["fieldvalue fv","on",[["t.id","=","fv.ref_id"],["and","fv.fieldname","=","'trans_transcast'"]]],
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message mss","on",[["mss.fieldname","=","sg.groupvalue"],["and","mss.secname","=","'transtate'"], 
              ["and","mss.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msc","on",[["msc.fieldname","=","fv.value"],["and","msc.secname","=","'trans_transcast'"], 
              ["and","msc.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]]],
          where:[[[["t.deleted","=","0"],["or",["tg.groupvalue","=","'cash'"]]]]]}},
      
      PaymentFieldsView: {
        columns: {transtype:true, direction:true, transnumber:true, fielddef:true, deffield_value:true},
        label: getText("fields_view"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:"case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end "},
          transnumber: {fieldtype:'string', wheretype:'where', orderby:2, 
            label:getText("document_transnumber"), sqlstr:'t.transnumber '},
          fielddef: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("fields_fielddef"), sqlstr:'df.description '},
          deffield_value: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("fields_value"), sqlstr:'fv.value '},
          notes: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("fields_notes"), sqlstr:'fv.notes '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
            "fv.id as row_id","fv.fieldname", "'fieldvalue' as form","fg.groupvalue as fieldtype",
            "case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end as direction", "t.transnumber",
            "df.description as fielddef",
            "case when fg.groupvalue in ('bool') then fv.value "+
            "when fg.groupvalue in ('integer') then {CCS}{FMS_INT}fv.value{FME_INT}{CCE} "+
            "when fg.groupvalue in ('float') then {CCS}{FMS_FLOAT}fv.value{FME_FLOAT}{CCE} "+
            "when fg.groupvalue in ('customer') then {CCS}'customer//'{SEP} {CAS_TEXT}rf_customer.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('tool') then {CCS}'tool//'{SEP} {CAS_TEXT}rf_tool.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('product') then {CCS}'product//'{SEP} {CAS_TEXT}rf_product.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then "+
              "{CCS}'trans//'{SEP} {CAS_TEXT}rf_trans.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('project') then {CCS}'project//'{SEP} {CAS_TEXT}rf_project.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('employee') then {CCS}'employee//'{SEP} {CAS_TEXT}rf_employee.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('place') then {CCS}'place//'{SEP} {CAS_TEXT}rf_place.id {CAE_TEXT}{CCE} "+
            "when fg.groupvalue in ('urlink') then {CCS}'url/'{SEP}fv.value{CCE} "+
            "when fg.groupvalue in ('password') then '**********' "+
            "when fg.groupvalue in ('notes', 'string', 'valuelist','date','time') then fv.value else null end as deffield_value, "+
            "case when fg.groupvalue in ('bool','float','integer','date','time','urlink','notes', 'string', 'valuelist') then fv.value "+
            "when fg.groupvalue in ('customer') then rf_customer.custname "+
            "when fg.groupvalue in ('tool') then rf_tool.serial "+
            "when fg.groupvalue in ('product') then rf_product.partnumber "+
            "when fg.groupvalue in ('trans', 'transitem', 'transmovement', 'transpayment') then rf_trans.transnumber "+
            "when fg.groupvalue in ('project') then rf_project.pronumber "+
            "when fg.groupvalue in ('employee') then rf_employee.empnumber "+
            "when fg.groupvalue in ('place') then rf_place.planumber "+
            "when fg.groupvalue in ('password') then '**********' else null end as export_deffield_value, "+
            "fv.notes"],
          from:"fieldvalue fv",
          inner_join:[
            ["deffield df","on",[["fv.fieldname","=","df.fieldname"],
              ["and","df.nervatype","=",[{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]], 
              ["and","df.deleted","=","0"],["and","df.visible","=","1"]]],
            ["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["trans t","on",["fv.ref_id","=","t.id"]],
            ["groups tg","on",[["t.transtype","=","tg.id"],["and","tg.groupvalue","in",[[],"'bank'","'cash'"]]]],
            ["groups dg","on",["t.direction","=","dg.id"]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            [[[{select:["fieldname"], from:"deffield df",
              inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","in",[[],"'integer'","'float'"]]]]}],"rf_number"],"on",["fv.fieldname","=","rf_number.fieldname"]],
            [[[{select:["fieldname"], from:"deffield df",
            inner_join:["groups fg","on",[["df.fieldtype","=","fg.id"],["and","fg.groupvalue","=","'date'"]]]}],"rf_date"],"on",["fv.fieldname","=","rf_date.fieldname"]],
            ["customer rf_customer","on",["fv.value","=","{CAS_TEXT}rf_customer.id {CAE_TEXT}"]],
            ["tool rf_tool","on",["fv.value","=","{CAS_TEXT}rf_tool.id {CAE_TEXT}"]],
            ["trans rf_trans","on",["fv.value","=","{CAS_TEXT}rf_trans.id {CAE_TEXT}"]],
            ["product rf_product","on",["fv.value","=","{CAS_TEXT}rf_product.id {CAE_TEXT}"]],
            ["project rf_project","on",["fv.value","=","{CAS_TEXT}rf_project.id {CAE_TEXT}"]],
            ["employee rf_employee","on",["fv.value","=","{CAS_TEXT}rf_employee.id {CAE_TEXT}"]],
            ["place rf_place","on",["fv.value","=","{CAS_TEXT}rf_place.id {CAE_TEXT}"]]],
          where:[["fv.deleted","=","0"],["and",[["t.deleted","=","0"],["or",["tg.groupvalue","=","'cash'"]]]],
            ["and","fv.ref_id","in",[{select:["trans_id"], from:"payment"}]]]}},
      
      PaymentInvoiceView: {
        columns: {paidnumber:true, paiddate:true, pcurr:true, paidamount:true, invnumber:true},
        label: getText("payment_invoices"),
        edit: "trans",
        fields: {
          transtype: {fieldtype:'string', wheretype:'where', orderby:0, 
            label:getText("document_transtype"), 
            sqlstr:'case when mst.msg is null then tg.groupvalue else mst.msg end '},
          direction: {fieldtype:'string', wheretype:'where', orderby:1, 
            label:getText("document_direction"), 
            sqlstr:"case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end "},
          paiddate: {fieldtype:'date', wheretype:'where', orderby:2, 
            label:getText("payment_paiddate"), sqlstr:'p.paiddate '},
          place: {fieldtype:'string', wheretype:'where', orderby:3, 
            label:getText("payment_place"), sqlstr:'pa.description '},
          paidnumber: {fieldtype:'string', wheretype:'where', orderby:4, 
            label:getText("payment_paidnumber"), sqlstr:'tp.transnumber '},
          pcurr: {fieldtype:'string', wheretype:'where', orderby:5, 
            label:getText("payment_pcurr"), sqlstr:'pa.curr '},
          paidamount: {fieldtype:'float', wheretype:'where', orderby:6, 
            label:getText("payment_amount"), sqlstr:'{CAS_FLOAT}af.value {CAE_FLOAT} '},
          prate: {fieldtype:'float', wheretype:'where', calc:"avg", orderby:7, 
            label:getText("payment_rate"), sqlstr:'{CAS_FLOAT}rf.value {CAE_FLOAT} '},
          invnumber: {fieldtype:'string', wheretype:'where', orderby:8, 
            label:getText("payment_invnumber"), sqlstr:'inv.transnumber '},
          icurr: {fieldtype:'string', wheretype:'where', orderby:9, 
            label:getText("payment_icurr"), sqlstr:'inv.curr '},
          invamount: {fieldtype:'float', wheretype:'where', orderby:10, 
            label:getText("payment_iamount"), sqlstr:'irow.amount '},
          pnotes: {fieldtype:'string', wheretype:'where', orderby:11, 
            label:getText("payment_description"), sqlstr:'p.notes '}
        },
        sql: {
          select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
            "ln.id as row_id","case when mst.msg is null then tg.groupvalue else mst.msg end as transtype",
            "case when tg.groupvalue='bank' then '' else "+
              "case when msd.msg is null then dg.groupvalue else msd.msg end end as direction",
            "{FMS_DATE}p.paiddate {FME_DATE} as paiddate","pa.description as place",
            "t.transnumber as paidnumber","pa.curr as pcurr",
            "{FMS_FLOAT}{CAS_FLOAT}af.value {CAE_FLOAT}{FME_FLOAT} as paidamount",
            "{CAS_FLOAT}af.value {CAE_FLOAT} as export_paidamount",
            "{FMS_FLOAT}{CAS_FLOAT}rf.value {CAE_FLOAT}{FME_FLOAT} as prate",
            "{CAS_FLOAT}rf.value {CAE_FLOAT} as export_prate",
            "{CCS}'trans/'{SEP}itg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}inv.id {CAE_TEXT}{CCE} as invnumber",
            "inv.transnumber as export_invnumber","inv.curr as icurr",
            "{FMS_FLOAT}irow.amount{FME_FLOAT} as invamount","irow.amount as export_invamount",
            "p.notes as pnotes"],
          from:"link ln",
          inner_join:[
            ["payment p","on",[["ln.ref_id_1","=","p.id"],["and","p.deleted","=","0"]]],
            ["trans t","on",["p.trans_id","=","t.id"]],
            ["groups tg","on",["t.transtype","=","tg.id"]],
            ["groups dg","on",["t.direction","=","dg.id"]],
            ["place pa","on",["t.place_id","=","pa.id"]],
            ["trans inv","on",["ln.ref_id_2","=","inv.id"]],
            ["groups itg","on",["inv.transtype","=","itg.id"]],
            [[[{select:["trans_id","sum(amount) as amount"],
              from:"item", where:["deleted","=","0"], group_by:["trans_id"]}],"irow"],"on",["inv.id","=","irow.trans_id"]],
            ["fieldvalue af","on",[["ln.id","=","af.ref_id"],["and","af.fieldname","=","'link_qty'"]]],
            ["fieldvalue rf","on",[["ln.id","=","rf.ref_id"],["and","rf.fieldname","=","'link_rate'"]]]],
          left_join:[
            ["ui_message mst","on",[["mst.fieldname","=","tg.groupvalue"],["and","mst.secname","=","'transtype'"], 
              ["and","mst.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]],
            ["ui_message msd","on",[["msd.fieldname","=","dg.groupvalue"],["and","msd.secname","=","'direction'"], 
              ["and","msd.lang","=",[{select:["value"], from:"fieldvalue", where:["fieldname","=","'default_lang'"]}]]]]],                  
          where:[["ln.deleted","=","0"],["and","ln.nervatype_1","=",[{select:["id"], 
            from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'payment'"]]}]],
            ["and","ln.nervatype_2","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]]]}}};
    }
  }
}