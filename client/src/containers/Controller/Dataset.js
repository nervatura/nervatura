const Dataset = {
  currency: () => { return [
    {infoName:"currency_view", infoType:"view"}
  ]},
  customer: () => { return [
    {infoName:"customer", infoType:"table", classAlias:"customer", 
      where:[["deleted","=","0"],["and","id","=","?"]]},
    {infoName:"custtype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'custtype'"]],
        ["and","groupvalue","<>","'own'"]], 
      order:"groupname, groupvalue"},
    {infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",
        [[],"'nervatype'","'fieldtype'","'logstate'","'custtype'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted","=","0"],["and","nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]]], 
      order:"description"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_customer_report'","'transyear'","'log_update'",
        "'log_deleted'","'log_customer_update'","'log_customer_deleted'"]]},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{
        select:["fieldname"], from:"deffield", where:["nervatype","=",
        [[],{select:["id"], from:"groups", 
          where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]]}]],
          ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"},
    {infoName:"address", infoType:"view"},
    {infoName:"contact", infoType:"view"},
    {infoName:"event", infoType:"view"}
  ]},
  deffield: () => { return [
    {infoName:"deffield_view", infoType:"view"},
    {infoName:"fieldtype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'fieldtype'"],
        ["and","groupvalue","not in",[[],"'filter'","'checkbox'","'trans'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"nervatype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'nervatype'"],
        ["and","groupvalue","in",[[],"'customer'","'employee'","'event'","'formula'",
        "'place'","'product'","'project'","'tool'","'trans'"]]], 
      order:"groupname, groupvalue"},
  ]},
  employee: () => { return [
    {infoName:"employee", infoType:"view"},
    {infoName:"usergroup", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'usergroup'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"department", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'department'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'nervatype'","'fieldtype'","'logstate'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted","=","0"],["and","nervatype","=",
        [[],{select:["id"], from:"groups", where:[["groupname","=","'nervatype'"],
          ["and","groupvalue","=","'employee'"]]}]]], 
      order:"description"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_employee_report'","'transyear'","'log_update'",
        "'log_deleted'","'log_employee_update'","'log_employee_deleted'"]]},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{select:["fieldname"], from:"deffield", 
        where:["nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"},
    {infoName:"address", infoType:"view"},
    {infoName:"contact", infoType:"view"},
    {infoName:"event", infoType:"view"}
  ]},
  event: () => { return [
    {infoName:"event", infoType:"view"},
    {infoName:"eventgroup", infoType:"table", classAlias:"groups",
      where:[["deleted","=","0"],["and","groupname","in",[[],"'eventgroup'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"groups", infoType:"table", classAlias:"groups",
      where:[["deleted","=","0"],["and","groupname","in",[[],"'nervatype'","'fieldtype'","'logstate'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted", "=", "0"],["and", "nervatype", "=",[[],
       {select:["id"], from: "groups", where: [["groupname", "=", "'nervatype'"], 
         ["and", "groupvalue", "=", "'event'"]]}]]], 
    order:"description"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue",
      where:[["fieldname","in",[[],"'default_event_report'","'transyear'","'log_update'",
        "'log_deleted'", "'log_event_update'", "'log_event_deleted'"]]]},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue",
      where:[["deleted","=","0"],["and","fieldname","in",[[],{select:["fieldname"], from:"deffield", 
        where:["nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'event'"]]}]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"}
  ]},
  groups: () => { return [
    {infoName:"groups_view", infoType:"view"}
  ]},
  numberdef: () => { return [
    {infoName:"numberdef_view", infoType:"view"}
  ]},
  place: () => { return [
    {infoName:"place", infoType:"table", classAlias:"place", 
      where:[["deleted","=","0"],["and","id","=","?"]]},
    {infoName:"address", infoType:"table", classAlias:"address", 
      where:[["deleted","=","0"],["and","nervatype","=",[[],{select:["id"], 
        from:"groups", where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'place'"]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"placetype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'placetype'"]]], order:"groupname, groupvalue"},
    {infoName:"currency", infoType:"table", classAlias:"currency", where:""},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_place_report'","'transyear'",
        "'log_update'","'log_deleted'","'log_place_update'","'log_place_deleted'"]]},
    {infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'nervatype'","'fieldtype'",
        "'logstate'","'placetype'"]]], order:"groupname, groupvalue"},
    {infoName:"contact", infoType:"view"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted","=","0"],["and","nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'place'"]]}]]],
      order:"description"},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{select:"fieldname", from:"deffield", 
        where:["nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'place'"]]}]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"},
    {infoName:"place_view", infoType:"view"}
  ]},
  printqueue: () => { return [
    {infoName:"server_printers", infoType:"view"},
    {infoName:"items", infoType:"view"}
  ]},
  product: () => { return [
    {infoName:"product", infoType:"table", classAlias:"product", 
      where:[["deleted","=","0"],["and","id","=","?"]]},
    {infoName:"protype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'protype'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"barcodetype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'barcodetype'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"calcmode", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'calcmode'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'nervatype'","'fieldtype'",
        "'logstate'","'protype'","'calcmode'","'barcodetype'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted","=","0"],["and","nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'product'"]]}]]], 
        order:"description"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_product_report'","'transyear'","'log_update'",
      "'log_deleted'","'log_product_update'","'log_product_deleted'","'default_currency'",
      "'default_unit'","'default_taxcode'"]]},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{select:"fieldname", from:"deffield", 
        where:["nervatype","=",[{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'product'"]]}]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"},
    {infoName:"currency", infoType:"table", classAlias:"currency", where:""},
    {infoName:"tax", infoType:"table", classAlias:"tax", where:""},
    {infoName:"event", infoType:"view"},
    {infoName:"barcode", infoType:"view"},
    {infoName:"price", infoType:"view"},
    {infoName:"discount", infoType:"view"}
  ]},
  project: () => { return [
    {infoName:"project", infoType:"view"},
    {infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'nervatype'","'fieldtype'","'logstate'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted","=","0"],["and","nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]]], 
        order:"description"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_project_report'","'transyear'","'log_update'",
        "'log_deleted'","'log_project_update'","'log_project_deleted'"]]},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{select:["fieldname"], from:"deffield", 
        where:["nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"},
    {infoName:"address", infoType:"view"},
    {infoName:"contact", infoType:"view"},
    {infoName:"event", infoType:"view"}
  ]},
  rate: () => { return [
    {infoName:"rate", infoType:"view"},
    {infoName:"currency", infoType:"table", classAlias:"currency", 
      where:""},
    {infoName:"ratetype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'ratetype'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"rategroup", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'rategroup'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_rate_report'","'transyear'","'log_update'",
        "'log_deleted'","'log_rate_update'","'log_rate_deleted'","'default_currency'"]]}
  ]},
  report: () => { return [
    {infoName:"report", infoType:"view"}
  ]},
  setting: () => { return [
    {infoName:"setting_view", infoType:"view"}
  ]},
  tax: () => { return [
    {infoName:"tax_view", infoType:"view"}
  ]},
  template: () => { return [
    {infoName:"template", infoType:"view"},
    {infoName:"template_view", infoType:"view"}
  ]},
  tool: () => { return [
    {infoName:"tool", infoType:"view"},
    {infoName:"toolgroup", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'toolgroup'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'nervatype'","'fieldtype'","'logstate'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:[["deleted","=","0"],["and","nervatype","=",[[],{
        select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'tool'"]]}]]], 
      order:"description"},
    {infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where:["fieldname","in",[[],"'default_tool_report'","'transyear'","'log_update'",
        "'log_deleted'","'log_tool_update'","'log_tool_deleted'"]]},
    {infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{
        select:"fieldname", from:"deffield", where:["nervatype","=",
          [[],{select:["id"], from:"groups", 
            where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'tool'"]]}]]}]],
        ["and","ref_id","=","?"]]},
    {infoName:"deffield_prop", infoType:"view"},
    {infoName:"event", infoType:"view"}
  ]},
  trans: (transtype) => {
    var dataSet = [];
    dataSet.push({infoName:"groups", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","in",[[],"'transtype'","'direction'","'trans'",
        "'nervatype'","'fieldtype'","'logstate'","'placetype'","'movetype'"]]], 
      order:"groupname, groupvalue"});
    dataSet.push({infoName:"transtate", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'transtate'"]], 
      order:"groupname, groupvalue"});
    dataSet.push({infoName:"department", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'department'"]], 
      order:"groupname, groupvalue"});
    dataSet.push({infoName:"paidtype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'paidtype'"]], 
      order:"groupname, groupvalue"});
    
    dataSet.push({infoName:"deffield", infoType:"table", classAlias:"deffield", 
      where:["nervatype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}],
        ["or","nervatype","=",[[],{select:["id"], from:"groups", 
          where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'link'"]]}]]], 
        order:"description"});
    dataSet.push({infoName:"settings", infoType:"table", classAlias:"fieldvalue", 
      where: ["fieldname","in",[[],"'default_currency'","'show_partnumber'","'show_partcustomer'",
      "'invoice_from_inventory'","'log_trans_update'","'log_trans_deleted'","'log_trans_closed'",
      "'default_trans_"+transtype+"_out_report'","'default_trans_"+transtype+"_in_report'",
      "'trans_closed_policy_"+transtype+"_out'","'trans_closed_policy_"+transtype+"_in'",
      "'invoice_copy'","'default_deadline'","'default_paidtype'","'default_bank'",
      "'default_chest'","'default_warehouse'"]]});
    dataSet.push({infoName:"fieldvalue", infoType:"table", classAlias:"fieldvalue", 
      where:[["deleted","=","0"],["and","fieldname","in",[[],{select:"fieldname", from:"deffield df",
        inner_join:["groups g","on",[["df.nervatype","=","g.id"],["and","g.groupvalue","=","'trans'"]]]}]],
      ["and","ref_id","=","?"]]});
    dataSet.push({infoName:"deffield_prop", infoType:"view"});
          
    dataSet.push({infoName:"translink", infoType:"view"});
    dataSet.push({infoName:"cancel_link", infoType:"view"});
    dataSet.push({infoName:"currency", infoType:"table", classAlias:"currency", where:""});
    dataSet.push({infoName:"tax", infoType:"table", classAlias:"tax", where:""});
    dataSet.push({infoName:"pattern", infoType:"table", classAlias:"pattern", 
      where:[["deleted","=","0"],["and","transtype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'transtype'"],["and","groupvalue","=","'"+transtype+"'"]]}]]]});
    dataSet.push({infoName:"delivery_pattern", infoType:"table", classAlias:"pattern", 
      where:[["deleted","=","0"],["and","transtype","=",[[],{select:["id"], from:"groups", 
        where:[["groupname","=","'transtype'"],["and","groupvalue","=","'delivery'"]]}]]]});
    dataSet.push({infoName:"place", infoType:"table", classAlias:"place", 
      where:["deleted","=","0"]});
    dataSet.push({infoName:"direction", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'direction'"],
        ["and","groupvalue","in",[[],"'out'","'in'"]]], order:"groupname, groupvalue"});
          
    dataSet.push({infoName:"trans", infoType:"view"});
    
    dataSet.push({infoName:"item", infoType:"view"});
    dataSet.push({infoName:"tool_movement", infoType:"view"});
    dataSet.push({infoName:"element_count", infoType:"view"});
    
    dataSet.push({infoName:"payment", infoType:"view"});
      
    switch (transtype) {
      case "bank":
      case "cash":
        dataSet.push({infoName:"payment_link", infoType:"view"});
        dataSet.push({infoName:"payment_link_fieldvalue", infoType:"table", classAlias:"fieldvalue", 
          where:[["deleted","=","0"],["and","fieldname","in",[[],"'link_qty'","'link_rate'"]],
          ["and","ref_id","in",[[],{select:["id"], from:"link", 
            where:[["deleted","=","0"],["and","nervatype_1","=",[[],{select:["id"], from:"groups", 
              where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'payment'"]]}]],
            ["and","ref_id_1","in",[[],{select:["id"], from:"payment", 
              where:[["deleted","=","0"],["and","trans_id","=","?"]]}]]]}]]]});
        break;
      
      case "invoice":
        dataSet.push({infoName:"invoice_link", infoType:"view"});
        dataSet.push({infoName:"invoice_link_fieldvalue", infoType:"table", classAlias:"fieldvalue", 
          where:[["deleted","=","0"],["and","fieldname","in",[[],"'link_qty'","'link_rate'"]],
            ["and","ref_id","in",[[],{select:["id"], from:"link", 
              where:[["deleted","=","0"],["and","nervatype_2","=",[[],{select:["id"], from:"groups", 
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]],
            ["and","ref_id_2","=","?"]]}]]]});
        break;
      
      case "formula":
        dataSet.push({infoName:"movement_head", infoType:"view", sqlKey:"movement_formula_head"});
        dataSet.push({infoName:"movement", infoType:"view", sqlKey:"movement_formula"});
        break;
      
      case "production":
        dataSet.push({infoName:"formula_head", infoType:"view", sqlKey:"formula_head"});
        dataSet.push({infoName:"movement_head", infoType:"view", sqlKey:"movement_production_head"});
        dataSet.push({infoName:"movement", infoType:"view", sqlKey:"movement_production"});
        break;
        
      case "delivery":
        dataSet.push({infoName:"place", infoType:"table", classAlias:"place", 
          where:"deleted=0"});
        dataSet.push({infoName:"movement", infoType:"view", sqlKey:"movement_delivery"});
        dataSet.push({infoName:"movement_transfer", infoType:"view"});
        break;
      
      case "inventory":
        dataSet.push({infoName:"movement", infoType:"view", sqlKey:"movement_inventory"});
        break;
      
      case "waybill":
        dataSet.push({infoName:"movement", infoType:"view", sqlKey:"movement_waybill"});
        break;
        
      case "order":
      case "worksheet":
      case "rent":
        dataSet.push({infoName:"transitem_invoice", infoType:"view"});
        dataSet.push({infoName:"transitem_shipping", infoType:"view"});
        dataSet.push({infoName:"shipping_items", infoType:"view"});
        dataSet.push({infoName:"shipping_delivery", infoType:"view"});
        break;
      default:
    }
      
    return dataSet;
  },
  ui_menu: () => { return [
    {infoName:"fieldtype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'fieldtype'"], 
        ["and","groupvalue","in",[[],"'bool'","'date'","'integer'","'float'","'string'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"method", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'method'"]], 
      order:"groupname, groupvalue"},
    {infoName:"ui_menu_view", infoType:"view"}
  ]},
  usergroup: () => { return [
    {infoName:"nervatype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'nervatype'"],
      ["and","groupvalue","in",[[],"'audit'","'customer'","'product'","'price'","'employee'",
        "'tool'","'project'","'event'","'setting'","'trans'","'menu'","'report'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"transtype", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'transtype'"], 
        ["and","groupvalue","in",[[],"'offer'","'order'","'worksheet'","'rent'","'invoice'",
          "'delivery'","'inventory'","'waybill'","'production'","'formula'","'bank'","'cash'"]]], 
      order:"groupname, groupvalue"},
    {infoName:"inputfilter", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'inputfilter'"]], 
      order:"groupname, groupvalue"},
    {infoName:"transfilter", infoType:"table", classAlias:"groups", 
      where:[["deleted","=","0"],["and","groupname","=","'transfilter'"]], 
      order:"groupname, groupvalue"},
    {infoName:"reportkey", infoType:"view"},
    {infoName:"menukey", infoType:"view"},
    {infoName:"usergroup_view", infoType:"view"}
  ]}
}
    
export default Dataset;
