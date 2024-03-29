export const Quick = {
  customer: () => ({
    sql: {
      select: ["{CCS}'customer'{SEP}'//'{SEP} {CAS_TEXT}c.id {CAE_TEXT}{CCE} as id",
        "c.custname as label",
        "c.custname as custname", "c.custnumber as custnumber", "addr.city as city",
        "addr.street as street"],
      from: "customer c",
      left_join: [[[{select:["*"], from:"address", 
        where:["id","in",[{select:["min(id) fid"], from:"address a", 
          where:[["a.deleted","=","0"],
            ["and","a.nervatype","=",[{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]]], 
          group_by:["a.ref_id"]}]]}],"addr"],"on",["c.id","=","addr.ref_id"]], 
      where:[["c.deleted","=","0"], 
        ["and","c.id","not in",[{select:["customer.id"], from:"customer", 
          inner_join: ["groups","on",[["customer.custtype","=","groups.id"],
            ["and","groups.groupvalue","=","'own'"]]]}]]]},
    columns: [
      ["custname", "c.custname"],
      ["custnumber", "c.custnumber"],
      ["city", "addr.city"],
      ["street", "addr.street"]]
  }),
  
  employee: () => ({
    sql: {
      select:["{CCS}'employee'{SEP}'//'{SEP} {CAS_TEXT}e.id {CAE_TEXT}{CCE} as id",
        "e.empnumber as label",
        "e.empnumber as empnumber", "pt.groupvalue as usergroup",
        "c.firstname as firstname", "c.surname as surname"],
      from:"employee e",
      inner_join:["groups pt","on",["e.usergroup","=","pt.id"]],
      left_join:["contact c","on",[["e.id","=","c.ref_id"],
        ["and","c.nervatype","=", [{select:["id"], from:"groups", 
          where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]]],
      where:[["e.deleted","=","0"]]},
    columns: [
      ["empnumber", "e.empnumber"],
      ["usergroup", "pt.groupvalue"],
      ["firstname", "c.firstname"],
      ["surname", "c.surname"]
    ]
  }),
  
  payment: () => ({
    sql: {
      select: [
        "{CCS}'payment'{SEP}'/'{SEP} tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "{CCS}t.transnumber{SEP} case when tg.groupvalue='bank' "+
        "then {CCS}' ~ '{SEP} {CAS_TEXT}p.id {CAE_TEXT} {CCE} else '' end {CCE} as label",
        "{CCS}t.transnumber{SEP} case when tg.groupvalue='bank' "+
        "then {CCS}' ~ '{SEP} {CAS_TEXT}p.id {CAE_TEXT} {CCE} else '' end {CCE} as paidnumber",
        "{FMS_DATE}p.paiddate {FME_DATE} as paiddate", "tg.groupvalue as transtype", "dg.groupvalue as direction",
        "pl.curr as curr", "p.amount as amount", "t.id as trans_id"],
      from:"payment p",
      inner_join:[["trans t","on",[["p.trans_id","=","t.id"],["and","t.deleted","=","0"]]],
        ["groups tg","on",["t.transtype","=","tg.id"]], 
        ["groups dg","on",["t.direction","=","dg.id"]], 
        ["place pl","on",["t.place_id","=","pl.id"]]], 
      where:[["p.deleted","=","0"]]},
    columns: [
      ["paidnumber", "{CCS}t.transnumber{SEP} case when tg.groupvalue='bank' "+
        "then {CCS}' ~ '{SEP} {CAS_TEXT}p.id {CAE_TEXT} {CCE} else '' end {CCE}"],
      ["paiddate", "p.paiddate"],
      ["curr", "pl.curr"],
      ["amount", "p.amount"]
    ]
  }),
  
  place: () => ({
    sql: {
      select:["{CCS}'place'{SEP}'/'{SEP}'/'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "{CCS}p.planumber{SEP} case when p.curr is null then '' else {CCS}' | ' {SEP}p.curr {CCE} end{CCE} as label",
        "p.planumber as planumber", "p.description as description", "pt.groupvalue as placetype"],
      from:"place p", 
      inner_join:["groups pt","on",["p.placetype","=","pt.id"]], 
      where:[["p.deleted","=","0"]]} ,
    columns: [
      ["planumber", "p.planumber"],
      ["description", "p.description"],
      ["placetype", "pt.groupvalue"]
    ]
  }),
  
  place_bank: () => ({
    sql: {
      select:["{CCS}'place'{SEP}'/'{SEP}'/'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "{CCS}p.planumber{SEP} case when p.curr is null then '' else {CCS}' | ' {SEP}p.curr {CCE} end{CCE} as label",
        "p.planumber as planumber", "p.description as description", "p.curr as curr"], 
      from:"place p",
      inner_join:["groups pt","on",[["p.placetype","=","pt.id"],["and","pt.groupvalue","=","'bank'"]]],
      where:[["p.deleted","=","0"]]},
    columns: [
      ["planumber", "p.planumber"],
      ["description", "p.description"],
      ["curr", "p.curr"]
    ]
  }),
  
  place_cash: () => ({
    sql: {
      select:["{CCS}'place'{SEP}'/'{SEP}'/'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "{CCS}p.planumber{SEP} case when p.curr is null then '' else {CCS}' | ' {SEP}p.curr {CCE} end{CCE} as label", 
        "p.planumber as planumber", "p.description as description", "p.curr as curr"], 
      from:"place p", 
      inner_join:["groups pt","on",[["p.placetype","=","pt.id"],["and","pt.groupvalue","=","'cash'"]]],
      where:[["p.deleted","=","0"]]} ,
    columns: [
      ["planumber", "p.planumber"],
      ["description", "p.description"],
      ["curr", "p.curr"]
    ]
  }),
  
  place_warehouse: () => ({
    sql: {
      select:["{CCS}'place'{SEP}'/'{SEP}'/'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "p.planumber as label", "p.planumber as planumber", "p.description as description"],
      from:"place p",
      inner_join:["groups pt","on",[["p.placetype","=","pt.id"],
        ["and","pt.groupvalue","=","'warehouse'"]]],
      where:[["p.deleted","=","0"]]} ,
    columns: [
      ["planumber", "p.planumber"],
      ["description", "p.description"]
    ]
  }),
  
  product: () => ({
    sql: {
      select:["{CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "p.partnumber as label",
        "p.partnumber as partnumber", "p.description as description", "pt.groupvalue as protype",
        "p.unit", "p.tax_id"],
      from:"product p",
      inner_join:["groups pt","on",["p.protype","=","pt.id"]],
      where:[["p.deleted","=","0"]]},
    columns: [
      ["partnumber", "p.partnumber"],
      ["description", "p.description"],
      ["protype", "pt.groupvalue"]]
  }),
  
  product_item: () => ({
    sql: {
      select: ["{CCS}'product'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "{CCS}p.partnumber{SEP}' | '{SEP}p.description{CCE} as label",
        "p.partnumber as partnumber", "p.description as description", "pt.groupvalue as protype",
        "p.unit", "p.tax_id", "tx.taxcode as tax"],
      from:"product p",
      inner_join:["groups pt","on",[["p.protype","=","pt.id"],["and","pt.groupvalue","=","'item'"]]],
      left_join:["tax tx","on",["p.tax_id","=","tx.id"]],
      where:[["p.deleted","=","0"]]},
    columns: [
      ["partnumber", "p.partnumber"],
      ["description", "p.description"],
      ["protype", "pt.groupvalue"],
      ["tax", "tx.taxcode"]]
    }),
  
  project: () => ({
    sql: {
      select:["{CCS}'project'{SEP}'//'{SEP} {CAS_TEXT}p.id {CAE_TEXT}{CCE} as id",
        "p.pronumber as label",
        "p.pronumber as pronumber", "p.description as description", "p.startdate as startdate"],
      from:"project p", 
      where:[["p.deleted","=","0"]]} ,
    columns: [
      ["pronumber", "p.pronumber"],
      ["description", "p.description"],
      ["startdate", "p.startdate"]]
  }),
  
  report: (usergroupId) => ({
    sql: {
      select: ["{CCS}'report'{SEP}'//'{SEP} {CAS_TEXT}r.id {CAE_TEXT}{CCE} as id",
        "r.repname as label",
        "r.repname as repname", "r.label as rlabel", "r.description as description", 
        "ft.groupvalue as ftype"],
      from: "ui_report r",
      inner_join:[["groups tg","on",["r.nervatype","=","tg.id"]], 
        ["groups ft","on",["r.filetype","=","ft.id"]]],
      where:[["tg.groupvalue","=","'report'"],["and","ft.groupvalue","in",[[],"'pdf'","'csv'"]], 
        ["and","r.id","not in",[{select:["ur.id"], from:"ui_audit a", 
          inner_join:[["ui_report ur","on",["a.subtype","=","ur.id"]], 
            ["groups inf","on",["a.inputfilter","=","inf.id"]]], 
          where:[["inf.groupvalue","=","'disabled'"],["and","a.usergroup","=", usergroupId]]}]]]},
    columns: [
      ["repname", "r.repname"],
      ["label", "r.label"],
      ["description", "r.description"]]
  }),
  
  servercmd: (usergroupId) => ({
    sql: {
      select:["{CCS}'servercmd'{SEP}'//'{SEP} {CAS_TEXT}m.id {CAE_TEXT}{CCE} as id",
        "m.description as label",
        "m.description as description"],
      from:"ui_menu m",
      inner_join: ["groups st", "on", ["m.method", "=", "st.id"]],
      where:[["st.groupvalue","<>","'printer'"], ["and","m.id","not in",[{select:["um.id"], from:"ui_audit a", 
        inner_join:[["ui_menu um","on",["a.subtype","=","um.id"]], 
          ["groups inf","on",["a.inputfilter","=","inf.id"]]], 
        where:[["inf.groupvalue","=","'disabled'"],["and","a.usergroup","=",usergroupId]]}]]]},
    columns: [
      ["description", "m.description"]]
  }),
  
  tool: () => ({
    sql: {
      select: ["{CCS}'tool'{SEP}'//'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
        "t.serial as label",
        "t.serial as serial", "t.description as description", "p.description as product"],
      from:"tool t",
      inner_join:["product p","on",["t.product_id","=","p.id"]],
      where:[["t.deleted","=","0"]]},
    columns: [
      ["serial", "t.serial"],
      ["description", "t.description"],
      ["product", "p.description"]]
  }),
  
  transitem: () => ({
    sql: {
      select:["{CCS}'trans'{SEP}'/'{SEP} tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
        "t.transnumber as label",
        "t.transnumber as transnumber", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE} as transtype",
        "c.custname as custname", "t.deleted as deleted"],
      from:"trans t",
      inner_join:[["groups tg","on",[["t.transtype","=","tg.id"],["and","tg.groupvalue","in",
        [[],"'invoice'","'receipt'","'order'","'offer'","'worksheet'","'rent'"]]]],
        ["groups dg","on",["t.direction","=","dg.id"]]],
      left_join:["customer c","on",["t.customer_id","=","c.id"]],
      where:[[[["t.deleted","=","0"],["or",[["tg.groupvalue","=","'invoice'"],["and","dg.groupvalue","=","'out'"]]],
        ["or",[["tg.groupvalue","=","'receipt'"],["and","dg.groupvalue","=","'out'"]]]]]]
    },
    columns: [
      ["transnumber", "t.transnumber"],
      ["transtype", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE}"],
      ["custname", "c.custname"]
    ]
  }),
  
  transitem_invoice: () => ({
    sql: {
      select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
        "t.transnumber as label",
        "t.transnumber as transnumber", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE} as transtype",
        "c.custname as custname", "t.deleted as deleted",
        "t.curr", "t.id as trans_id"],
      from:"trans t",
      inner_join:[[["groups tg","on",["t.transtype","=","tg.id"]],
        ["and","tg.groupvalue","in",[[],"'invoice'","'receipt'"]]],
        ["groups dg","on",["t.direction","=","dg.id"]]],
      left_join:["customer c","on",["t.customer_id","=","c.id"]],
      where:[[[["t.deleted","=","0"],["or",[["tg.groupvalue","=","'invoice'"],
        ["and","dg.groupvalue","=","'out'"]]]]]]},
    columns: [
      ["transnumber", "t.transnumber"],
      ["transtype", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE}"],
      ["custname", "c.custname"],
      ["curr","t.curr"]
    ]
  }),
  
  transitem_delivery: () => ({
    sql: {
      select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
        "t.transnumber as label",
        "t.transnumber as transnumber", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE} as transtype",
        "c.custname as custname", "t.deleted as deleted"],
      from:"trans t",
      inner_join:[["groups tg","on",[["t.transtype","=","tg.id"],
        ["and","tg.groupvalue","in",[[],"'order'","'worksheet'","'rent'"]]]],
        ["groups dg","on",["t.direction","=","dg.id"]]],
      left_join:["customer c","on",["t.customer_id","=","c.id"]], 
      where:[["t.deleted","=","0"]]},
    columns: [
      ["transnumber", "t.transnumber"],
      ["transtype", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE}"],
      ["custname", "c.custname"]
    ]
  }),
  
  transmovement: () => ({
    sql: {
      select:["{CCS}'trans'{SEP}'/'{SEP} tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
        "t.transnumber as label",
        "t.transnumber as transnumber", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE} as transtype",
        "{FMS_DATE}t.transdate {FME_DATE} as transdate"],
      from:"trans t",
      inner_join:[["groups tg","on",[["t.transtype","=","tg.id"],
        ["and","tg.groupvalue","in",[[],"'delivery'","'inventory'","'waybill'","'production'","'formula'"]]]],
        ["groups dg","on",["t.direction","=","dg.id"]]],
      where:[["t.deleted","=","0"]]},
    columns: [
      ["transnumber", "t.transnumber"],
      ["transtype", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE}"],
      ["transdate", "{FMS_DATE}t.transdate {FME_DATE}"]]
  }),
  
  transpayment: () => ({
    sql: {
      select:["{CCS}'trans'{SEP}'/'{SEP}tg.groupvalue{SEP}'/'{SEP} {CAS_TEXT}t.id {CAE_TEXT}{CCE} as id",
        "t.transnumber as label",
        "t.transnumber as transnumber", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE} as transtype",
        "p.description as place", "p.curr as curr"],
      from:"trans t",
      inner_join:[["groups tg","on",[["t.transtype","=","tg.id"],
        ["and","tg.groupvalue","in",[[],"'bank'","'cash'"]]]],
        ["groups dg","on",["t.direction","=","dg.id"]],
        ["place p","on",["t.place_id","=","p.id"]]],
      where:[[[["t.deleted","=","0"],["or",["tg.groupvalue","=","'cash'"]]]]] },
    columns: [
      ["transnumber", "t.transnumber"],
      ["transtype", "{CCS}tg.groupvalue{SEP}'-'{SEP}dg.groupvalue{CCE}"],
      ["place", "p.description"],
      ["curr", "p.curr"]]
  })
}