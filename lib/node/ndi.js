/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/* global Buffer */

var async = require("async");

var conf = require('./conf.js');
var lang = require('./lang.js')[conf.lang];
var models = require('./adapter.js').models();
var connect = require('./adapter.js').connect();
var ntura = require('./models.js');
var out = require('./tools.js').DataOutput();

module.exports = function() {

function decode_data(data){
  
  function get_param_list(data_str){
    var drows = [];
    if (data_str.trim()!==""){
      var rows = data_str.trim().split("||");
      rows.forEach(function(row) {
        var drow = {};
        var fields = row.split("|");
        fields.forEach(function(field) {
          if (field.indexOf("=")===-1){
            drow[field] = true;}
          else {
            drow[field.substring(0, field.indexOf("="))] = field.substring(field.indexOf("=")+1);}
        });
        if (Object.keys(drow).length>0){
          drows.push(drow);}});}
    return drows;}
  
  var retvalue = {code:"json", output:"json", error:null};
  if (typeof data.jsonrpc !== "undefined"){
    retvalue.params = data.params[0]; retvalue.items = data.params[1];
    switch (data.method) {
    case "getData":
    case "getData_json":
      if (!data.params[0] || !data.params[1]){
        retvalue.error = get_error((data.id || "get"), "missing_params", "params,filters");}
      else {
        retvalue.id = (data.id || "get"); retvalue.params.type = "get"; retvalue.method = get_data;}
      break;
    case "updateData":
    case "updateData_json":
      if (!data.params[0] || !data.params[1]){
        retvalue.error = get_error((data.id || "update"), "missing_params", "params,data");}
      else {
        retvalue.id = (data.id || "update"); retvalue.params.type = "update"; retvalue.method = update_data;}
      break;
    case "deleteData":
    case "deleteData_json":
      if (!data.params[0] || !data.params[1]){
        retvalue.error = get_error((data.id || "delete"), "missing_params", "params,data");}
      else {
        retvalue.id = (data.id || "delete"); retvalue.params.type = "delete"; retvalue.method = update_data;}
      break;
    default:
      retvalue.error = get_error((data.id || null), "message", lang.unknown_method+" "+data.method);}
    if (!retvalue.error){
      if (retvalue.params.type === "get"){
        if (retvalue.items.output==="html" || retvalue.items.output==="csv"){
          retvalue.output = retvalue.items.output;}}}}
  else {
    switch (data.method) {
    case "getData":
      if ((typeof data.params==="undefined") || (typeof data.filters==="undefined")){
        retvalue.error = get_error("get", "missing_params", "params,filters");}
      else {
        retvalue.id = "get"; retvalue.method = get_data;
        retvalue.params = data.params; retvalue.items = data.filters;}
      break;
    case "updateData":
      if ((typeof data.params==="undefined") || (typeof data.data==="undefined")){
        retvalue.error = get_error("update", "missing_params", "params,data");}
      else {
        retvalue.id = "update"; retvalue.method = update_data;
        retvalue.params = data.params; retvalue.items = data.data;}
      break;
    case "deleteData":
      if ((typeof data.params==="undefined") || (typeof data.data==="undefined")){
        retvalue.error = get_error("delete", "missing_params", "params,data");}
      else {
        retvalue.id = "delete"; retvalue.method = update_data;
        retvalue.params = data.params; retvalue.items = data.data;}
      break;}
    if (!retvalue.error){
      if ((data.code==="base64")){
        retvalue.params = new Buffer(retvalue.params, 'base64').toString('utf8');
        retvalue.items = new Buffer(retvalue.items, 'base64').toString('utf8');}
      retvalue.params = get_param_list(retvalue.params)[0];
      retvalue.params.type = retvalue.id;
      retvalue.items = get_param_list(retvalue.items);
      if (retvalue.params.type === "get"){
        retvalue.items = retvalue.items[0];
        if (typeof retvalue.items.output!=="undefined"){
          retvalue.output = retvalue.items.output;}}}}
  return retvalue;}

function get_error(id, ekey, message, data) {
  var err_msg = "";
  if (typeof data==="undefined"){data = null;}
  switch (ekey) {
    case "missing_params":
      err_msg = lang.missing_params+": "+message;
      break;
    default:
      err_msg = message;
      break;}
  return {"id": id, "jsonrpc": "2.0", 
      "error": {"code":ekey,"message":err_msg,"data":data}};}

function update_data(nstore, params, items, _callback){
  if (typeof params.check_audit === "undefined") {params.check_audit = true;}
  if (typeof params.insert_field === "undefined") {params.insert_field = false;}
  if (typeof params.insert_row === "undefined") {params.insert_row = false;}
  var result = {id:params.type+"_"+params.datatype, data:[], new:[]};
  
  async.waterfall([    
    function(callback) {
      if (params.check_audit && nstore.employee()===null && params.validator.validate){
        callback(lang.ndi_invalid_login);}
      else if (typeof params.datatype === "undefined"){
         callback(lang.ndi_invalid_datatype);}
      else if (!items){
        _callback(null, result);}
      else if (items.length === 0){
        _callback(null, result);}
      else {
        var data = {model:{groups:{}, refvalue:{}, required:[]}, items:items, values:[], renew:false,
          checklist:{groups:[], refvalue:{}, nextnumber:{key:[], value:[]}, 
            audit:{nervatype:[],transtype:[]}, settings:{}, extra_info:false, pattern:{}},
          valid_direction :{
            out: ['invoice','receipt','order','offer','worksheet','rent','delivery','waybill','cash'],
            in: ['invoice','order','offer','worksheet','rent','delivery','waybill','cash'],
            transfer: ['delivery','inventory','production','formula','bank']}};
        var err_msg = null;
        if (ntura.model.hasOwnProperty(params.datatype) && params.datatype.substr(0,3)!=="ui_"){
          //check data model: p.key, references, access
          var keys = ntura.model[params.datatype]._key;
          if (ntura.model[params.datatype].hasOwnProperty(keys[0])){
            if (ntura.model[params.datatype][keys[0]].default === "nextnumber"){
              data.checklist.nextnumber.fieldname = keys[0];};}
          var access = ntura.model[params.datatype]._access;
          if (access.length === 1 && access[0] !== "transtype"){
            data.checklist.audit.nervatype.push(access[0]);}
          if (params.type === "update"){
            if (keys.indexOf("rownumber")>-1){
              data.renew = true;}
            for (var fieldname in ntura.model[params.datatype]) {
              if (fieldname.substr(0,1) !== "_" ){
                var field = ntura.model[params.datatype][fieldname];
                if (field.notnull){
                  var req = {name:fieldname};
                  if (field.hasOwnProperty("refname")){
                    req.refname = field.refname;}
                  else {
                    req.refname = fieldname;}
                  if (field.hasOwnProperty("default")){
                    req.default = field.default;}
                  data.model.required.push(req);}
                if (field.hasOwnProperty("refname") || field.hasOwnProperty("references")) {
                  if (field.hasOwnProperty("references")){
                    if (field.references[0] !== params.datatype){
                      if (field.references[0] === "groups"){
                        if (Object.keys(field.requires).length>1){
                          data.model.groups[fieldname.replace("_","")] = Object.keys(field.requires);}
                        else {
                          data.model.groups[fieldname.replace("_","")] = Object.keys(field.requires)[0];}}
                      else {
                        data.model.refvalue[field.refname] = field.references[0];}}}
                  else {
                    data.model.refvalue[field.refname] = null;}}}}} 
          data.items.forEach(function(item) {
            //item required fields
            if (params.datatype === "fieldvalue"){
              if (!item.hasOwnProperty("fieldname")){
                  err_msg = lang.missing_required_field+" fieldname";}}
            else if (!data.checklist.nextnumber.hasOwnProperty("fieldname")) {
              keys.forEach(function(key) {
                if (key !== "rownumber" && !item.hasOwnProperty(key)){
                  err_msg = lang.missing_required_field+" "+key;}});}
            if (!err_msg){
              //create id
              if (keys.indexOf("rownumber")>-1 && !item.hasOwnProperty("rownumber") && 
                params.datatype !== "fieldvalue"){
                item.id = "";}
              else if (data.checklist.nextnumber.hasOwnProperty("fieldname") && (!item.hasOwnProperty(keys[0]) || item[keys[0]] === "")){
                //counter-type id (numberdef)
                item.id = "";
                if (params.datatype === "trans"){
                  if (!item.hasOwnProperty("transtype")){
                    err_msg = lang.missing_required_field+" transtype";}
                  else if (!item.hasOwnProperty("direction")){
                    err_msg = lang.missing_required_field+" direction";}
                  else if (!data.valid_direction.hasOwnProperty(item.direction)){
                    err_msg = lang.invalid_fieldname.replace("fieldname","direction")+" "+item.direction;}
                  else if(data.valid_direction[item.direction].indexOf(item.transtype)===-1){
                    err_msg = lang.invalid_fieldname.replace("fieldname","transtype/direction")+" "+item.transtype+"-"+item.direction;}
                  else {
                    if (item.transtype==="cash" || item.transtype==="waybill"){
                      data.checklist.nextnumber.key.push(item.transtype);}
                    else {
                      data.checklist.nextnumber.key.push(item.transtype+"_"+item.direction);}
                    data.checklist.audit.transtype.push(item.transtype);
                    if (item.transtype === "invoice"){
                      data.checklist.extra_info = true;}}}
                else {
                  data.checklist.nextnumber.key.push(keys[0]);}}
              else {
                switch (params.datatype) {
                  case "address":
                  case "contact":
                    if (item.hasOwnProperty("rownumber") && item.rownumber !== "" && item.rownumber !== null) {
                      item.id = item.nervatype+"/"+item.refnumber+"~"+item.rownumber;}
                    if (data.checklist.audit.nervatype.indexOf(item.nervatype)===-1){
                      data.checklist.audit.nervatype.push(item.nervatype);}
                    break;
                  case "fieldvalue":
                    if (item.hasOwnProperty("refnumber") && item.refnumber !== "" 
                      && item.refnumber !== null && item.hasOwnProperty("rownumber") 
                      && item.rownumber !== "" && item.rownumber !== null) {
                        item.id = item.refnumber+"~~"+item.fieldname+"~"+item.rownumber;}
                    else if (item.hasOwnProperty("refnumber") && item.refnumber !== "" && item.refnumber !== null) {
                      item.id = item.refnumber+"~~"+item.fieldname;
                      item.fieldname = item.fieldname.split("~")[0];}
                    else {
                      //setting
                      item.fieldname = item.fieldname.split("~")[0];
                      item.id = item.fieldname;
                      if (data.checklist.audit.nervatype.indexOf("setting")===-1){
                        data.checklist.audit.nervatype.push("setting");}}
                    break;           
                  case "groups":
                    item.id = item.groupname+"~"+item.groupvalue;
                    break;       
                  case "item":
                  case "movement":
                  case "payment":
                    if (item.hasOwnProperty("rownumber") && item.rownumber !== "" && item.rownumber !== null) {
                      item.id = item.transnumber+"~"+item.rownumber;}
                    break;
                  case "link":
                    item.id = item.nervatype1+"~"+item.refnumber1+"~~"+item.nervatype2+"~"+item.refnumber2;
                    if (data.checklist.audit.nervatype.indexOf(item.nervatype1)===-1){
                      data.checklist.audit.nervatype.push(item.nervatype1);}
                    if (data.checklist.audit.nervatype.indexOf(item.nervatype2)===-1){
                      data.checklist.audit.nervatype.push(item.nervatype2);}
                    break;
                  case "log":
                    item.id = item.empnumber+"~"+out.getValidDateTime(item.crdate);
                    break;
                  case "price":
                    item.id = item.partnumber+"~"+item.pricetype+"~"+item.validfrom+"~"+item.curr+"~"+item.qty;
                    break;
                  case "rate":
                    item.id = item.ratetype+"~"+item.ratedate+"~"+item.curr+"~"+item.planumber;
                    break;   
                  default:
                    item.id = item[keys[0]];
                    break;}}}
            if (!err_msg && params.type === "update"){
              switch (params.datatype) {
                case "event":
                  if (item.id === ""){
                    if (!item.hasOwnProperty("nervatype")){
                      err_msg = lang.missing_required_field+" nervatype";}
                    else {
                      if (data.checklist.audit.nervatype.indexOf(item.nervatype)===-1){
                        data.checklist.audit.nervatype.push(item.nervatype);}}}
                  break;
                case "price":
                  data.checklist.settings.default_currency = null;
                  break;
                case "product":
                  data.checklist.settings.default_taxcode = null;
                  data.checklist.settings.default_unit = null;
                  break;
                case "trans":
                  data.checklist.settings.default_currency = null;
                  data.checklist.settings.default_paidtype = null;
                  data.checklist.settings.default_deadline = null;
                  data.checklist.settings.audit_control = null;
                  data.checklist.groups.transtate = {new:null, ok:null};
                  break;}}
            if (!err_msg && params.type === "update"){
              for (var fieldname in item) {
                //check groups and reference value
                if (data.model.groups.hasOwnProperty(fieldname)){
                  if(item[fieldname] !== "" && item[fieldname] !== null){
                    if (Array.isArray(data.model.groups[fieldname])){
                      data.model.groups[fieldname].forEach(function(groupname) {
                        if (typeof data.checklist.groups[groupname] === "undefined"){
                          data.checklist.groups[groupname] = {};}});
                      if (typeof data.checklist.groups.all === "undefined"){
                        data.checklist.groups.all = {};}
                      if (typeof data.checklist.groups.all[fieldname] === "undefined"){
                        data.checklist.groups.all[fieldname] = {};}
                      if(item[fieldname] !== "" && item[fieldname] !== null){
                        data.checklist.groups.all[fieldname][item[fieldname]] = null;}}
                    else {
                      if (typeof data.checklist.groups[data.model.groups[fieldname]] === "undefined"){
                        data.checklist.groups[data.model.groups[fieldname]] = {};}
                      if(item[fieldname] !== "" && item[fieldname] !== null){
                        data.checklist.groups[data.model.groups[fieldname]][item[fieldname]] = null;}}}}
                if (data.model.refvalue.hasOwnProperty(fieldname)){
                  if (data.model.refvalue[fieldname] === null){
                    var vtype;
                    if (params.datatype !== "fieldvalue") {
                      if (fieldname === "refnumber"){
                        vtype = item.nervatype;}
                      else if (fieldname === "refnumber1"){
                        vtype = item.nervatype1;}
                      else if (fieldname === "refnumber2"){
                        vtype = item.nervatype2;}}
                    if (vtype){
                      if (typeof data.checklist.refvalue[vtype] === "undefined"){
                        data.checklist.refvalue[vtype] = {};}
                      if(item[fieldname] !== "" && item[fieldname] !== null){
                        data.checklist.refvalue[vtype][item[fieldname]] = null;}}}
                  else {
                    if (typeof data.checklist.refvalue[data.model.refvalue[fieldname]] === "undefined"){
                      data.checklist.refvalue[data.model.refvalue[fieldname]] = {};}
                    if(item[fieldname] !== "" && item[fieldname] !== null){
                      data.checklist.refvalue[data.model.refvalue[fieldname]][item[fieldname]] = null;}}}}}});}
        else {
          err_msg = lang.ndi_invalid_datatype+": "+params.datatype;}
      callback(err_msg, data);}},
    
    function(data, callback) {
      //database settings
      if (Object.keys(data.checklist.settings).length>0){
        nstore.connect.getDatabaseSettings({conn:params.validator.conn}, function(err, settings){
          if (err){
            callback(err, data);}
          else {
            if (data.checklist.extra_info){
              data.checklist.pattern = settings.pattern;}
            var keys = Object.keys(data.checklist.settings);
            keys.forEach(function(skey) {
              if (settings.setting.hasOwnProperty(skey)) {
                if (settings.setting[skey].value !== null && settings.setting[skey].value !== "") {
                  data.checklist.settings[skey] = settings.setting[skey].value;
                  switch (skey) {
                    case "default_taxcode":
                      if (!data.checklist.refvalue.hasOwnProperty("tax")){
                        data.checklist.refvalue.tax = {};}
                      data.checklist.refvalue.tax[settings.setting[skey].value] = null;
                      break;
                    case "default_paidtype":
                      if (!data.checklist.groups.hasOwnProperty("paidtype")){
                        data.checklist.groups.paidtype = {};}
                      data.checklist.groups.paidtype[settings.setting[skey].value] = null;
                      break;
                    case "default_currency":
                    case "default_unit":
                    case "audit_control":
                      break;}}}});
            callback(null, data);}});}
      else {callback(null, data);}},
    
    function(data, callback) {
      //check id values
      var refvalues = []; 
      data.items.forEach(function(item) {
        if (item.id !== ""){
          refvalues.push(function(callback_){
            nstore.valid.getIdFromRefnumber({conn:params.validator.conn,
              nervatype:params.datatype, refnumber:item.id, 
              use_deleted:params.hasOwnProperty("use_deleted")}, function(err, id, info){
              var refnumber_err = null;
              var ref_item = data.items.filter(
                function(item){return(item.id===info.refnumber)});
              if (ref_item.length>0){
                if (err || id === null) {
                  if (!data.renew && params.type === "update"){
                    result.data.push(ref_item[0].id);}
                  ref_item[0].id = "";
                  switch (info.nervatype) {
                    case "fieldvalue":
                      if (!refnumber_err && !info.ref_nervatype){
                        refnumber_err = lang.invalid_fieldname+" "+ref_item[0]["fieldname"];}
                      else if(!refnumber_err && info.ref_nervatype === "setting" 
                        && ref_item[0].hasOwnProperty("refnumber") && ref_item[0]["refnumber"] !== ""
                        && ref_item[0]["refnumber"] !== null){
                        refnumber_err = lang.invalid_fieldname.replace("fieldname","refnumber")+" "+ref_item[0]["refnumber"];}
                      else if (ref_item[0].hasOwnProperty("refnumber") && ref_item[0]["refnumber"] !== ""
                        && ref_item[0]["refnumber"] !== null){
                        data.renew = true;
                        ref_item[0].nervatype = info.ref_nervatype;
                        if (data.checklist.audit.nervatype.indexOf(info.ref_nervatype)===-1){
                          data.checklist.audit.nervatype.push(info.ref_nervatype);}
                        if (typeof data.checklist.refvalue[info.ref_nervatype] === "undefined"){
                          data.checklist.refvalue[info.ref_nervatype] = {};}
                        data.checklist.refvalue[info.ref_nervatype][ref_item[0]["refnumber"]] = null;}
                      break;
                    case "event":
                      if (!refnumber_err && !ref_item[0].hasOwnProperty("nervatype")){
                        refnumber_err = lang.missing_required_field+" nervatype";}
                      else {
                        if (data.checklist.audit.nervatype.indexOf(ref_item[0].nervatype)===-1){
                          data.checklist.audit.nervatype.push(ref_item[0].nervatype);}}
                      break;
                    case "trans":
                      if (!ref_item[0].hasOwnProperty("transtype")){
                        refnumber_err = lang.missing_required_field+" transtype";}
                      else if (!ref_item[0].hasOwnProperty("direction")){
                        refnumber_err = lang.missing_required_field+" direction";}
                      else if (!data.valid_direction.hasOwnProperty(ref_item[0].direction)){
                        refnumber_err = lang.invalid_fieldname.replace("fieldname","direction")+" "+ref_item[0].direction;}
                      else if(data.valid_direction[ref_item[0].direction].indexOf(ref_item[0].transtype)===-1){
                        refnumber_err = lang.invalid_fieldname.replace("fieldname","transtype/direction")+" "+ref_item[0].transtype+"-"+ref_item[0].direction;}
                      else {
                        if (ref_item[0].transtype==="cash" || ref_item[0].transtype==="waybill"){
                          data.checklist.nextnumber.key.push(ref_item[0].transtype);}
                        else {
                          data.checklist.nextnumber.key.push(ref_item[0].transtype+"_"+ref_item[0].direction);}
                        data.checklist.audit.transtype.push(ref_item[0].transtype);
                        if (ref_item[0].transtype === "invoice"){
                          data.checklist.extra_info = true;}}
                      break;}}
                else {
                  result.data.push(ref_item[0].id);
                  ref_item[0].id = id;
                  switch (info.nervatype) {
                    case "customer":
                      if (info.custtype === "own"){
                        if (!refnumber_err && params.type === "delete"){
                          refnumber_err = lang.protected_data+": "+info.refnumber;}
                        if (ref_item[0].hasOwnProperty("custtype")){
                          refnumber_err = lang.protected_data+": own";}}
                      break;
                    case "event":
                      if (data.checklist.audit.nervatype.indexOf(info.ref_nervatype)===-1){
                        data.checklist.audit.nervatype.push(info.ref_nervatype);}
                      break;
                    case "fieldvalue":
                      if (data.checklist.audit.nervatype.indexOf(info.ref_nervatype)===-1){
                        data.checklist.audit.nervatype.push(info.ref_nervatype);}
                      if (ref_item[0].hasOwnProperty("refnumber")){
                        ref_item[0].nervatype = info.ref_nervatype;
                        if (typeof data.checklist.refvalue[info.ref_nervatype] === "undefined"){
                          data.checklist.refvalue[info.ref_nervatype] = {};}
                        data.checklist.refvalue[info.ref_nervatype][ref_item[0]["refnumber"]] = null;}
                      break;
                    case "item":
                      if (!ref_item[0].hasOwnProperty("qty")){
                        ref_item[0].qty = info.qty;}
                      if (!ref_item[0].hasOwnProperty("discount")){
                        ref_item[0].discount = info.discount;}
                      if (!ref_item[0].hasOwnProperty("taxcode")){
                        ref_item[0].tax_id = info.tax_id;
                        ref_item[0].rate_value = info.rate;}
                      break;
                    case "place":
                      if (ref_item[0].hasOwnProperty("placetype") && ref_item[0].placetype !== info.placetype){
                        refnumber_err = lang.disabled_readonly_type+": placetype";}
                      break;
                    case "movement":
                      if (ref_item[0].hasOwnProperty("movetype") && ref_item[0].movetype !== info.movetype){
                        refnumber_err = lang.disabled_readonly_type+": movetype";}
                      break;
                    case "trans":
                      if (ref_item[0].hasOwnProperty("transtype") && ref_item[0].transtype !== info.transtype){
                        refnumber_err = lang.disabled_readonly_type+": transtype";}
                      else if (ref_item[0].hasOwnProperty("direction") && ref_item[0].direction !== info.direction){
                        refnumber_err = lang.disabled_readonly_type+": direction";}
                      if (data.checklist.audit.transtype.indexOf(info.transtype)===-1){
                        data.checklist.audit.transtype.push(info.transtype);}
                      if (info.transtype === "invoice"){
                          data.checklist.extra_info = true;}
                      break;}}}
              else {
                refnumber_err=err;}
              callback_(refnumber_err);});});}});
      if (refvalues.length>0){
        async.series(refvalues,function(err) {
          callback(err, data);});}
      else {
        callback(null, data);}},
    
    function(data, callback) {
      if (Object.keys(data.checklist.refvalue).length>0){
        //check all refnumber
        var refvalues = [];
        var nervatype_lst = Object.keys(data.checklist.refvalue);
        nervatype_lst.forEach(function(nervatype) {
          var refnumber_lst = Object.keys(data.checklist.refvalue[nervatype]);
          refnumber_lst.forEach(function(refnumber) {
            refvalues.push(function(callback_){
              nstore.valid.getIdFromRefnumber({conn:params.validator.conn,
                nervatype:nervatype, refnumber:refnumber, 
                use_deleted:params.hasOwnProperty("use_deleted"), extra_info:data.checklist.extra_info}, function(err, id, info){
                if (err) {callback_(err);;}
                else if(id === null){
                  callback_(lang.missing_required_field+" "+info.refnumber);}
                else {
                  data.checklist.refvalue[info.nervatype][info.refnumber] = {id:id};
                  switch (info.nervatype) {
                    case "customer":
                      if (data.checklist.extra_info){
                        data.checklist.refvalue.customer[info.refnumber].compname = info.compname;
                        data.checklist.refvalue.customer[info.refnumber].comptax = info.comptax;
                        data.checklist.refvalue.customer[info.refnumber].compaddress = info.compaddress; 
                        data.checklist.refvalue.customer[info.refnumber].custname = info.custname;
                        data.checklist.refvalue.customer[info.refnumber].terms = info.terms;
                        data.checklist.refvalue.customer[info.refnumber].custtax = info.custtax;
                        data.checklist.refvalue.customer[info.refnumber].custaddress = info.custaddress;}
                      break;
                    case "product":
                      data.checklist.refvalue[info.nervatype][info.refnumber].description = info.description;
                      data.checklist.refvalue[info.nervatype][info.refnumber].unit = info.unit;
                      data.checklist.refvalue[info.nervatype][info.refnumber].tax_id = info.tax_id;
                      if (info.rate !== null){
                        data.checklist.refvalue[info.nervatype][info.refnumber].rate = info.rate;}
                      else {
                        data.checklist.refvalue[info.nervatype][info.refnumber].rate = 0;}
                      break;
                    case "tax":
                      data.checklist.refvalue[info.nervatype][info.refnumber].rate = info.rate;
                      break;
                    case "trans":
                      if (["item","payment","movement"].indexOf(params.datatype)>-1){
                        data.checklist.refvalue[info.nervatype][info.refnumber].digit = info.digit;
                        if (data.checklist.audit.transtype.indexOf(info.transtype)===-1){
                          data.checklist.audit.transtype.push(info.transtype);}}
                      break;}
                  callback_(null);}});});});});
        async.series(refvalues,function(err) {
          callback(err, data);});}
      else {callback(null, data);}},
    
    function(data, callback) {
      if (params.check_audit && data.checklist.audit.nervatype.length+data.checklist.audit.transtype.length>0){
        //input access
        var access_type = {conn:params.validator.conn};
        if (data.checklist.audit.nervatype.length>0){
          access_type.nervatype = data.checklist.audit.nervatype;}
        if (data.checklist.audit.transtype.length>0){
          access_type.transtype = data.checklist.audit.transtype;}
        nstore.connect.getObjectAudit(access_type, function(err, state, typevalue){
          if (err){
            callback(err, data);}
          else if ((state === "disabled" || state === "readonly") && params.type === "update"){
            callback(lang.disabled_readonly_type+": "+typevalue, data);}
          else if (state !== "all" && params.type === "delete"){
            callback(lang.disabled_readonly_type+": "+typevalue, data);}
          else {
            callback(null, data);}});}
      else {callback(null, data);}},
      
    function(data, callback) {
      if (Object.keys(data.checklist.groups).length>0){
        //check groups id
        var groups_err = null;
        nstore.valid.getGroupsId({conn:params.validator.conn, groupname:Object.keys(data.checklist.groups)}, 
        function(err, groups){
          if (err && !groups_err) {groups_err = err;}
          else {
            for (var groupname in data.checklist.groups) {
              if (groups.hasOwnProperty(groupname)){
                if (groupname === "all"){
                  for (var fieldname in data.checklist.groups.all) {
                    if (data.model.groups.hasOwnProperty(fieldname)){
                      var groupnames = data.model.groups[fieldname];
                      for (var groupvalue in data.checklist.groups.all[fieldname]) {
                        var group_item = groups.all.filter(function(item){
                          return(item.groupvalue===groupvalue && groupnames.indexOf(item.groupname)>-1)});
                        if (group_item.length>0){
                          data.checklist.groups.all[fieldname][groupvalue] = group_item[0].id;}
                        else if(groupvalue === ""){
                          data.checklist.groups.all[fieldname][groupvalue] = null;}
                        else {
                          groups_err = lang.invalid_fieldname.replace("fieldname",fieldname)+groupvalue;}}}}}
                else {
                  for (var groupvalue in data.checklist.groups[groupname]) {
                    if(groupvalue === ""){
                      data.checklist.groups[groupname][groupvalue] = null;}
                    else if (typeof groups[groupname][groupvalue] === "undefined"){
                      groups_err = lang.invalid_fieldname.replace("fieldname",groupname)+groupvalue;}
                    else {
                      data.checklist.groups[groupname][groupvalue] = groups[groupname][groupvalue];}}}}
              else {
                groups_err = lang.invalid_fieldname.replace("fieldname",groupname);}}}
          callback(groups_err, data);});}
      else {callback(null, data);}},
    
    function(data, callback) {
      if (data.checklist.nextnumber.key.length>0){
        var nextnumbers = [];
        if (!data.trans){
         data.trans = connect.beginTransaction({connection:params.validator.conn, engine:nstore.engine()});}
        data.checklist.nextnumber.key.forEach(function(numberkey) {
          nextnumbers.push(function(callback_){
            nstore.connect.nextNumber({numberkey:numberkey,insert_key:false,transaction:data.trans}, 
            function(err, value){
              if (!err){
                data.checklist.nextnumber.value.push(value);}
              callback_(err);});});});
        async.series(nextnumbers,function(err) {
          callback(err, data);});}
      else {
        callback(null, data);}},
    
    function(data, callback) {
      if (params.type === "update"){
        var values; var value_error=null;
        data.items.forEach(function(item) {
          values = {};
          if (item.id === "" && !params.hasOwnProperty("insert_row")){
            value_error = lang.disabled_insert;}
          switch (params.datatype) {
            case "customer":
              if (params.check_audit && !value_error && item.custtype==="own") {
                value_error = lang.invalid_value+": own";}
              break;
            case "deffield":
              if (!value_error && item.id==="") {
                if(item.valuelist){
                  item.valuelist = out.replaceAll(item.valuelist,"~","|");}
                if (item.description === "" || item.description === null || 
                  !item.hasOwnProperty("description")){
                  item.description = item.fieldname;}}
              break;
            case "event":
              if (item.hasOwnProperty("nervatype") && !item.hasOwnProperty("refnumber")){
                value_error = lang.missing_required_field+" refnumber";}
              if (!item.hasOwnProperty("nervatype") && item.hasOwnProperty("refnumber")){
                value_error = lang.missing_required_field+" nervatype";}
              break;
            case "fieldvalue":
            case "setting":
              if (item.fieldtype){delete item.fieldtype}
              if (item.description){delete item.description}
              break;
            case "item":
              if (!value_error && item.id==="") {
                if (item.hasOwnProperty("partnumber")){
                  if (item.description === "" || item.description === null || 
                    !item.hasOwnProperty("description")){
                    item.description = data.checklist.refvalue.product[item.partnumber].description;}
                  if (item.unit === "" || item.unit === null || !item.hasOwnProperty("unit")){
                    item.unit = data.checklist.refvalue.product[item.partnumber].unit;}
                  if (item.taxcode === "" || item.taxcode === null || !item.hasOwnProperty("taxcode")){
                    item.tax_id = data.checklist.refvalue.product[item.partnumber].tax_id;
                    item.rate_value = data.checklist.refvalue.product[item.partnumber].rate;}}
                else {
                  value_error = lang.missing_required_field+" partnumber";}}
              if (item.hasOwnProperty("taxcode")){
                item.rate_value = data.checklist.refvalue.tax[item.taxcode].rate;}
              if (!value_error && item.hasOwnProperty("inputmode")){
                if (!item.hasOwnProperty("inputvalue") && !value_error){
                  value_error = lang.missing_required_field+" inputvalue";}
                else {
                  if (!item.hasOwnProperty("qty")){
                    item.qty = 0;}
                  if (!item.hasOwnProperty("discount")){
                    item.discount = 0;}
                  var digit = data.checklist.refvalue.trans[item.transnumber].digit;
                  switch (item.inputmode) {
                    case "fxprice":
                      item.fxprice = parseFloat(item.inputvalue);
                      item.netamount = out.Round(item.fxprice*(1-item.discount/100)*item.qty,digit);
                      item.vatamount = out.Round(item.netamount*item.rate_value,digit);
                      item.amount = item.netamount + item.vatamount;
                      break;
                    case "netamount":
                      item.netamount = parseFloat(item.inputvalue);
                      if (item.qty === 0) {
                        item.fxprice = 0;
                        item.vatamount = 0;}
                      else {
                        item.fxprice = out.Round(item.netamount/(1-item.discount/100)/item.qty,digit);
                        item.vatamount = out.Round(item.netamount*item.rate_value,digit);}
                      item.amount = item.netamount + item.vatamount;
                      break;
                    case "amount":
                      item.amount = parseFloat(item.inputvalue);
                      if (item.qty === 0) {
                        item.fxprice = 0;
                        item.netamount = 0;
                        item.vatamount = 0;}
                      else {
                        item.netamount = out.Round(item.amount/(1+item.rate_value),digit);
                        item.vatamount = item.amount - item.netamount;
                        item.fxprice = out.Round(item.netamount/(1-item.discount/100)/item.qty,digit);}
                      break;
                    default:
                      if (!value_error) {
                        value_error = lang.invalid_value.replace("value","inputmode")+": "+item.inputmode;}
                      break;}}}
              break;
            case "log":
              if (!value_error){
                if (!item.hasOwnProperty("logstate")){
                  value_error = lang.missing_required_field+" logstate";}
                else if (item.id==="" && item.logstate !== "login" && item.logstate !== "logout" 
                  && !item.hasOwnProperty("nervatype")){
                    value_error = lang.missing_required_field+" nervatype";}
                else if ((item.logstate === "login" || item.logstate === "logout") && item.hasOwnProperty("nervatype")){
                  value_error = lang.invalid_value+": nervatype ";}
                else if (item.hasOwnProperty("nervatype") && !item.hasOwnProperty("refnumber")){
                  value_error = lang.missing_required_field+" refnumber";}
                else if (!item.hasOwnProperty("nervatype") && item.hasOwnProperty("refnumber")){
                  value_error = lang.missing_required_field+" nervatype";}}
            case "movement":
              if (!value_error){
                if (item.id === ""){
                  if (["store","inventory","formula","production"].indexOf(item.movetype) > -1 
                    && !item.hasOwnProperty("partnumber")){
                    value_error = lang.missing_required_field+" partnumber";}
                  else if (["store","inventory","production"].indexOf(item.movetype) > -1 
                    && !item.hasOwnProperty("planumber")){
                    value_error = lang.missing_required_field+" planumber";}
                  else if (item.movetype === "tool" && !item.hasOwnProperty("serial")){
                    value_error = lang.missing_required_field+" serial";}
                  else if(!item.hasOwnProperty("shippingdate")) {
                    item.shippingdate = out.getISODateTime();}}}
              break;
            case "place":
              if (!value_error){
                if (item.id === ""){
                  if ((item.placetype === "bank" || item.placetype === "cash") && !item.hasOwnProperty("curr")){
                    value_error = lang.missing_required_field+" curr";}}}
              break;
            case "price":
              if (!value_error){
                if (!item.hasOwnProperty("pricetype")){
                  value_error = lang.missing_required_field+" pricetype";}
                else if(item.pricetype === "price"){
                  item.discount = null;}
                else if(item.pricetype === "discount"){
                  if (item.id === "" && !item.hasOwnProperty("calcmode")){
                    value_error = lang.missing_required_field+" calcmode";}
                  else if (item.id === "" && !item.hasOwnProperty("discount")){
                    item.discount = 0;}}}
              break;
            case "product":
              if (!value_error && item.id === ""){
                if (!item.hasOwnProperty("unit")){
                  if (data.checklist.settings.default_unit!==null){
                    item.unit = data.checklist.settings.default_unit;}
                  else {
                    value_error = lang.missing_required_field+" unit";}}
                if (!item.hasOwnProperty("taxcode")){
                  if (data.checklist.settings.default_taxcode!==null){
                    item.tax_id = data.checklist.refvalue.tax[data.checklist.settings.default_taxcode].id;}
                  else if(!value_error) {
                    value_error = lang.missing_required_field+" taxcode";}}}
              break;
            case "trans":
              if (!value_error){
                if (item.hasOwnProperty("username")){
                  if (params.check_audit){
                    value_error = lang.disabled_readonly_type+": username";}
                  else {
                    item.cruser_id = data.checklist.refvalue.employee[item.username].id;
                    delete item.username;}}
                else if (item.id === ""){
                  if (!item.hasOwnProperty("custnumber") && 
                    ["offer","order","worksheet","rent","invoice"].indexOf(item.transtype)>-1){
                      value_error = lang.missing_required_field+" custnumber";}
                  else if (!item.hasOwnProperty("planumber") && ["bank","cash"].indexOf(item.transtype)>-1){
                    value_error = lang.missing_required_field+" planumber";}
                  else {
                    if (["offer","order","worksheet","rent","invoice"].indexOf(item.transtype)>-1){
                      if (!item.hasOwnProperty("curr")) {
                        if (data.checklist.settings.default_currency!==null){
                          item.curr = data.checklist.settings.default_currency;}
                        else {
                          value_error = lang.missing_required_field+" curr";}}
                      if (!item.hasOwnProperty("paidtype")) {
                        if (data.checklist.settings.default_paidtype!==null){
                          item.paidtype = data.checklist.settings.default_paidtype;}
                        else {
                          value_error = lang.missing_required_field+" paidtype";}}}
                    if (!item.hasOwnProperty("crdate")) {
                      item.crdate = out.getISODate();}
                    if (!item.hasOwnProperty("transdate")) {
                      item.transdate = out.getISODate();}
                    if (!item.hasOwnProperty("duedate")) {
                      if (item.transtype === "invoice" && item.direction === "out"){
                        //set invoice duedate
                        var duedate = new Date(); var terms = 0;
                        if (data.checklist.refvalue.customer[item.custnumber].terms > 0){
                          terms = data.checklist.refvalue.customer[item.custnumber].terms;}
                        else if (data.checklist.settings.default_deadline!==null){
                          terms = parseInt(data.checklist.settings.default_deadline,10);}
                        duedate.setDate(duedate.getDate()+terms);
                        item.duedate = out.getISODateTime(duedate,false);}
                      else if (["offer","order","worksheet","rent","invoice"].indexOf(item.transtype)>-1){
                        item.duedate = out.getISODateTime();}}
                    if (data.checklist.pattern.hasOwnProperty(item.transtype)){
                      if (data.checklist.pattern[item.transtype][0].defpattern){
                        item.fnote = data.checklist.pattern[item.transtype][0].notes;}}
                    if (data.checklist.settings.audit_control === "true"){
                      item.transtate = "new";}
                    else {
                      item.transtate = "ok";}
                    if (!item.hasOwnProperty("cruser_id")) {
                      item.cruser_id = nstore.employee().id;}
                    item.trans_transcast = "normal";}}}
              if (!value_error && item.transtype === "invoice" && item.hasOwnProperty("custnumber")){
                item.trans_custinvoice_compname = data.checklist.refvalue.customer[item.custnumber].compname;
                item.trans_custinvoice_comptax = data.checklist.refvalue.customer[item.custnumber].comptax;
                item.trans_custinvoice_compaddress = data.checklist.refvalue.customer[item.custnumber].compaddress; 
                item.trans_custinvoice_custname= data.checklist.refvalue.customer[item.custnumber].custname;
                item.trans_custinvoice_custtax = data.checklist.refvalue.customer[item.custnumber].custtax;
                item.trans_custinvoice_custaddress = data.checklist.refvalue.customer[item.custnumber].custaddress;}
              break;}
          if (!value_error){
            if (item.id==="" && data.checklist.nextnumber.hasOwnProperty("fieldname")){
              var fieldnumber = data.checklist.nextnumber.fieldname;
              if (!item.hasOwnProperty(fieldnumber) || item[fieldnumber] === ""){
                if (data.checklist.nextnumber.value.length >0) {
                  item[fieldnumber] = data.checklist.nextnumber.value.shift();
                  result.data.push(item[fieldnumber]);}}}
            for (var fieldname in item) {
              if (data.model.groups.hasOwnProperty(fieldname) && item[fieldname] !== "" && item[fieldname] !== null){
                if (Array.isArray(data.model.groups[fieldname])){
                  values[fieldname] = data.checklist.groups.all[fieldname][item[fieldname]];}
                else {
                  if (fieldname === "nervatype1"){
                    values.nervatype_1 = data.checklist.groups.nervatype[item.nervatype1];}
                  else if (fieldname === "nervatype2"){
                    values.nervatype_2 = data.checklist.groups.nervatype[item.nervatype2];}
                  else {
                    values[data.model.groups[fieldname]] = data.checklist.groups[data.model.groups[fieldname]][item[fieldname]];}}}
              else if (data.model.refvalue.hasOwnProperty(fieldname) && item[fieldname] !== "" && item[fieldname] !== null){
                if (data.model.refvalue[fieldname] === null){
                  if (fieldname === "refnumber"){
                    values.ref_id = data.checklist.refvalue[item.nervatype][item[fieldname]].id;}
                  else if (fieldname === "refnumber1"){
                    values.ref_id_1 = data.checklist.refvalue[item.nervatype1][item[fieldname]].id;}
                  else if (fieldname === "refnumber2"){
                    values.ref_id_2 = data.checklist.refvalue[item.nervatype2][item[fieldname]].id;}}
                else {
                  values[data.model.refvalue[fieldname]+"_id"] = data.checklist.refvalue[data.model.refvalue[fieldname]][item[fieldname]].id;}}
              else {
                switch (fieldname) {
                  case "rownumber":
                  case "inputmode":
                  case "inputvalue":
                  case "rate_value":
                  case "pricetype":
                  case "refnumber":
                    break;
                  case "nervatype":
                    if (params.datatype !== "fieldvalue"){
                      values[fieldname] = item[fieldname];}
                    break;
                  
                  case "empnumber":
                    if (params.datatype === "employee"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "calnumber":
                    if (params.datatype === "event"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "custnumber":
                    if (params.datatype === "customer"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "pronumber":
                    if (params.datatype === "project"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "planumber":
                    if (params.datatype === "place"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "taxcode":
                    if (params.datatype === "tax"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "partnumber":
                    if (params.datatype === "product"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "serial":
                    if (params.datatype === "tool"){
                      values[fieldname] = item[fieldname];}
                    break;
                  case "transnumber":
                    if (params.datatype === "trans"){
                      values[fieldname] = item[fieldname];}
                    break;
                  default:
                    values[fieldname] = item[fieldname];
                    break;}}}
            data.model.required.forEach(function(field) {
              if ((item.id === "") && (!item.hasOwnProperty(field.refname)) || 
                item[field.refname]==="" || item[field.refname]===null){
                  if (field.hasOwnProperty("default")){
                    if (typeof field.default === "string"){
                      field.default = out.replaceAll(field.default,"'","");}
                    values[field.refname] = field.default;}
                  else {
                    if (!value_error){
                      value_error = lang.missing_required_field+" "+field.refname;}}}});
            if (!value_error){
              data.values.push(values);}}});
        callback(value_error, data);}
      else {
        data.items.forEach(function(item) {
          if (item.id !== ""){
            data.values.push(item.id);}});
        callback(null, data);}},
            
    function(data, callback) {
      if (data.values.length>0){
        if (!data.trans){
         data.trans = connect.beginTransaction({connection:params.validator.conn, engine:nstore.engine()});}
        var value_lst = [];
        if (params.type === "update"){
          data.values.forEach(function(values) {
            value_lst.push(function(callback_){
              var update_params = {nervatype:params.datatype, values:values, 
                validate:true, insert_row:params.insert_row, insert_field:params.insert_field,
                transaction:data.trans, validator:params.validator};
              nstore.connect.updateData(update_params, function(err, record_id){
                if (!err){
                  var ref_item = data.items.filter(
                    function(item){
                      return(item.id===record_id)});
                    if (data.renew && ref_item.length===0){
                      result.new.push(record_id);}}
                callback_(err);});});});}
        else {
          data.values.forEach(function(id) {
            value_lst.push(function(callback_){
              var delete_params = {nervatype:params.datatype, ref_id:id, 
                transaction:data.trans, validator:params.validator};
              nstore.connect.deleteData(delete_params, function(err, id){
                callback_(err);});});});}
        async.series(value_lst,function(err) {
          callback(err, data);});}
      else {
        callback(null,data);}},
  
  function(data, callback) {
    if (params.type === "update" && result.new.length>0){
      var new_lst = [];
      result.new.forEach(function(id) {
        new_lst.push(function(callback_){
          var prm = {nervatype:params.datatype, ref_id:id, 
            conn:data.trans, use_deleted:params.hasOwnProperty("use_deleted")}
          nstore.valid.getRefnumber(prm, function(err,refnumber){
            if (!err){
              result.data.push(refnumber);}
            callback_(err);});});});
      async.series(new_lst,function(err) {
        callback(err, data);});}
    else {
      callback(null,data);}}
    
  ],
  function(err,data) {
    if (err){
      if(err){if(err.message){err = err.message;}}
      if (data){
        if (data.trans){
          if(data.trans.rollback){
            data.trans.rollback(function(error){
              _callback(err);});}
          else {_callback(err);}}
        else {_callback(err);}}
      else {_callback(err);}}
    else {
      if (data){
        if (data.trans){
          if(data.trans.commit){
            data.trans.commit(function(error){
              _callback(error, result);});}
          else {_callback(null, result);}}
        else {
          _callback(null, result);}}
      else {
        _callback(null, result);}}});}

function get_data(nstore, params, filter, _callback){
  if (typeof filter.output === "undefined"){
    filter.output = "json";}
  if (typeof params.check_audit === "undefined"){
    params.check_audit = true;}
  
  var represents = {
    represent_rownumber:{
      label:"rownumber*", rettype:"index", ref_id_field:"represent_rownumber", 
      nervatype_field:null, function:get_refnumber},
    represent_refnumber:{
      label:"refnumber*", rettype:"refnumber", ref_id_field:"represent_refnumber", 
      nervatype_field:null, function:get_refnumber},
    represent_refnumber_nervatype:{
      label:"refnumber*", rettype:"refnumber", ref_id_field:"represent_refnumber_nervatype", 
      nervatype_field:"nervatype", function:get_refnumber},
    represent_refnumber1:{
      label:"refnumber1*", rettype:"refnumber", ref_id_field:"represent_refnumber1", 
      nervatype_field:"nervatype1", function:get_refnumber},
    represent_refnumber2:{
      label:"refnumber2*", rettype:"refnumber", ref_id_field:"represent_refnumber2", 
      nervatype_field:"nervatype2", function:get_refnumber},
    represent_fieldvalue:{
      label:"value*", rettype:"refnumber", ref_id_field:"represent_fieldvalue", 
      nervatype_field:"fieldtype", function:get_refnumber},
    represent_fieldname:{
      label:"fieldname*", rettype:"fieldname", ref_id_field:"represent_fieldname", 
      nervatype_field:null, function:get_refnumber}}
      
  function get_refnumber(params, callback){
    //row,repkey,use_deleted
    try {
      var ref_id = params.row[represents[params.repkey].ref_id_field];
      if (!ref_id){
        return callback(null, params.row, params.repkey, null);}
      var nervatype;
      if (represents[params.repkey].nervatype_field === null){
        nervatype = params.datatype;}
      else {
        nervatype = params.row[represents[params.repkey].nervatype_field];}
      var prm = {conn:params.conn, rettype:represents[params.repkey].rettype,
        nervatype:nervatype, ref_id:ref_id, use_deleted:params.use_deleted};
      if (params.repkey === "represent_fieldvalue"){
        switch (nervatype) {
          case "customer":
          case "employee":
          case "place":
          case "product":
          case "project":
          case "tool":
          case "trans":
          case "event":
            nstore.valid.getRefnumber(prm, function(err,refnumber){
              callback(err, params.row, params.repkey, refnumber);});
            break;
          case "password":
            callback(null, params.row, params.repkey, ref_id);
            break;
          default:
            callback(null, params.row, params.repkey, ref_id);
            break;}}
      else {
        nstore.valid.getRefnumber(prm, function(err,refnumber){
          callback(err, params.row, params.repkey, refnumber);});}} 
    catch (error) {
      callback(null, params.row, params.repkey, lang.ndi_missing_refnumber);}}
  
  async.waterfall([    
    function(callback) {
      if (params.check_audit && nstore.employee()===null && params.validator.validate){
        callback(lang.ndi_invalid_login);}
      else if (["html","xml","json","csv"].indexOf(filter.output)===-1){
        callback(lang.invalid_fieldname+": "+filter.output);}
      else if (filter.hasOwnProperty("cross_tab") && filter.output!=="html"){
        callback(lang.ndi_cross_format);}
      else {
        var data = {data_audit:"all", item_str:"", fld_value:"", items:[],
          select_str:"", from_str:"", where_str:"", orderby_str:"", limit_str:"", orderby:"",
          fvalue_sql:[], repfields:{}, labels:{}, columns:{}, fieldvalues:{}}
        if (typeof filter.where !== "undefined" && filter.where !== ""){
          data.where_str = "and " + filter.where;}
        if (typeof filter.orderby !== "undefined" && filter.orderby !== ""){
          data.orderby = " order by " + filter.orderby;}
        if (typeof filter.limit !== "undefined" && params.datatype !== "sql"){
          data.limit_str=" limit " + filter.limit;}
        callback(null, data);}},
    
    function(data, callback) {
      if (params.check_audit){
        var access = "setting"; 
        if(ntura.model[params.datatype]){
          ntura.model[params.datatype]._access;
          if (access.length === 1) {
            access = access[0];}}
        if (access === "transtype"){
          //data access rights
          nstore.connect.getDataAudit({conn:params.validator.conn}, function(err, state){
              if (err){
                callback(err);}
              else {
                data.data_audit = state;
                callback(null, data);}});}
        else {
          //objects access
          var access_type = {conn:params.validator.conn, nervatype:access, transtype:null}; var err_msg = null;
          switch (params.datatype) {
            case "address":
            case "contact":
            case "fieldvalue":
              if (typeof filter.nervatype !== "undefined"){
                if (access_type.nervatype.indexOf(filter.nervatype) === -1){
                  err_msg = lang.invalid_nervatype;}
                else {
                  access_type.nervatype = filter.nervatype;}}
              break;
            
            case "event":
              if (typeof filter.nervatype !== "undefined"){
                if (access_type.nervatype.indexOf(filter.nervatype) === -1){
                  err_msg = lang.invalid_nervatype;}}
              break;
            
            case "link":
              if (filter.hasOwnProperty("nervatype1") || filter.hasOwnProperty("nervatype1")){
                var link_acc = "";
                if (filter.hasOwnProperty("nervatype1")){
                  if (access_type.nervatype.indexOf(filter.nervatype1) === -1){
                    err_msg = lang.invalid_nervatype;}
                  else {
                    if (filter.hasOwnProperty("nervatype2")){
                      link_acc = "'"+filter.nervatype1+"',";}
                    else {
                      link_acc = filter.nervatype1;}}}
                if (filter.hasOwnProperty("nervatype2")){
                  if (access_type.nervatype.indexOf(filter.nervatype2) === -1){
                    err_msg = lang.invalid_nervatype;}
                  else {
                    if (filter.hasOwnProperty("nervatype1")){
                      link_acc += "'"+filter.nervatype2+"'";}
                    else {
                      link_acc = filter.nervatype2;}}}
                access_type.nervatype = link_acc;}
              break;}
          if (err_msg){
            callback(err_msg);}
          else {
            nstore.connect.getObjectAudit(access_type, function(err, state){
              if (err){
                callback(err);}
              else if (state === "disabled"){
                callback(lang.disabled_type+": "+access_type.nervatype);}
              else {
                callback(null, data);}});}}}
    else {callback(null, data);}},
    
    function(data, callback) {
      var fields_ = {}; var nervanumber;
      switch (params.datatype) {
        case "sql":
          data.item_str = "item";
          if (typeof filter.sql === "undefined" || filter.sql === ""){
            return callback(lang.missing_sql);}
          if (filter.sql.trim().toLowerCase().split(" ")[0]!=="select" && 
            filter.sql.trim().toLowerCase().split(" ")[0]!=="update" && 
            filter.sql.trim().toLowerCase().split(" ")[0]!=="insert"){
            return callback(lang.valid_sql);}
          break;
        
        case "address":
          data.item_str = "address"; data.fld_value = "address";
          if (typeof filter.nervatype !== "undefined"){
            nervanumber = nstore.valid.getTableKey(filter.nervatype);
            fields_ = {"nervatype":"g.groupvalue", "refnumber":"nerv."+nervanumber, 
              "represent_rownumber":"address.id", "country":"address.country", "state":"address.state", 
              "zipcode":"address.zipcode", "city":"address.city", "street":"address.street", 
              "notes":"address.notes"};
            data._sql = {
              select:["g.groupvalue as nervatype",
                "nerv."+nstore.valid.getTableKey(filter.nervatype)+" as refnumber", 
                "address.id as represent_rownumber","address.country","address.state",
               "address.zipcode","address.city","address.street","address.notes"],
              from:"address",
              inner_join:[["groups g","on",[["address.nervatype","=","g.id"],
                ["and","g.groupvalue","=","'"+filter.nervatype+"'"]]],
                [filter.nervatype+" nerv","on",[["address.ref_id","=","nerv.id"]]]],
              left_join:[]}
            if (params.use_deleted){
              data._sql.inner_join[1][2].push(["and","nerv.deleted","=","0"]);}}
          else {
            fields_ = {"nervatype":"g.groupvalue", "represent_refnumber_nervatype":"address.ref_id", 
              "represent_rownumber":"address.id", "country":"address.country", "state":"address.sate", 
              "zipcode":"address.zipcode", "city":"address.city", "street":"address.street", 
              "notes":"address.notes"};
            data._sql = {
              select:["g.groupvalue as nervatype","address.ref_id as represent_refnumber_nervatype",
                "address.id as represent_rownumber","address.country","address.state",
                "address.zipcode","address.city","address.street","address.notes"],
              from:"address", inner_join:["groups g","on",["address.nervatype","=","g.id"]],
              left_join:[]}}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("address.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "address.deleted";
            data._sql.select.push("address.deleted");}
          else {
            data._sql.where = ["address.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str =  " order by nervatype, id ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " address.id");
          data.orderby_str = data.orderby_str.replace(" id", " address.id"); 
          break;
        
        case "barcode":
          data.item_str = "barcode"; data.fld_value = "barcode";
          fields_ = {"code":"barcode.code", "partnumber":"p.partnumber", "barcodetype":"g.groupvalue", 
            "description":"barcode.description", "qty":"barcode.qty", "defcode":"barcode.defcode"};
          data._sql = {
            select:["barcode.code as code","p.partnumber","g.groupvalue as barcodetype",
              "barcode.description","barcode.qty","barcode.defcode"],
            from:"barcode", 
            inner_join:[["groups g","on",["barcode.barcodetype","=","g.id"]],
              ["product p","on",["barcode.product_id","=","p.id"]]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("p.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];}
          else {
            data._sql.where = ["p.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = "order by id ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " p.id");
          data.where_str = data.where_str.replace(" barcodetype", " g.groupvalue");
          data.orderby_str = data.orderby_str.replace(" id", " p.id"); 
          break;
        
        case "contact":
          data.item_str = "contact"; data.fld_value = "contact";
          if (typeof filter.nervatype !== "undefined"){
            nervanumber = nstore.valid.getTableKey(filter.nervatype);
            fields_ = {"nervatype":"g.groupvalue", "refnumber":"nerv."+nervanumber, "represent_rownumber":"contact.id", 
              "firstname":"contact.firstname", "surname":"contact.surname", "status":"contact.status", "phone":"contact.phone", 
              "fax":"contact.fax", "mobil":"contact.mobil", "email":"contact.email", "notes":"contact.notes"};
            data._sql = {
              select:["g.groupvalue as nervatype","nerv."+nstore.valid.getTableKey(filter.nervatype)+" as refnumber",
                "contact.id as represent_rownumber","contact.firstname","contact.surname",
                "contact.status","contact.phone","contact.fax","contact.mobil","contact.email","contact.notes"],
              from:"contact",
              inner_join:[["groups g","on",[["contact.nervatype","=","g.id"],
                ["and","g.groupvalue","=","'"+filter.nervatype+"'"]]],
                [filter.nervatype+" nerv","on",[["contact.ref_id","=","nerv.id"]]]],
              left_join:[]}
            if (params.use_deleted){
              data._sql.inner_join[1][2].push(["and","nerv.deleted","=","0"]);}}
          else {
            fields_ = {"nervatype":"g.groupvalue", "represent_refnumber_nervatype":"contact.ref_id", 
              "represent_rownumber":"contact.id","firstname":"contact.firstname", "surname":"contact.surname", 
              "status":"contact.status", "phone":"contact.phone","fax":"contact.fax", "mobil":"contact.mobil",
              "email":"contact.email", "notes":"contact.notes"};
            data._sql = {
              select:["g.groupvalue as nervatype","contact.ref_id as represent_refnumber_nervatype",
                "contact.id as represent_rownumber","contact.firstname","contact.surname", 
                "contact.status","contact.phone","contact.fax","contact.mobil","contact.email","contact.notes"],
              from:"contact", inner_join:["groups g","on",["contact.nervatype","=","g.id"]],
              left_join:[]}}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("contact.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "contact.deleted";
            data._sql.select.push("contact.deleted");}
          else {
            data._sql.where = ["contact.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = "order by nervatype, id ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " contact.id");
          data.orderby_str = data.orderby_str.replace(" id", " contact.id");
          break;
        
        case "currency":
          data.item_str = "currency"; data.fld_value = "currency";
          data._sql = {
            select:["curr","description","digit","defrate","cround"],
            from:"currency", inner_join:[], left_join:[], where:["1","=","1"]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("id");}
          if (data.orderby === "") {
            data.orderby_str = "order by curr ";}
          else {
            data.orderby_str = data.orderby;}
          break;
        
        case "customer":
          data.item_str = "customer"; data.fld_value = "customer";
          fields_ = {"custnumber":"customer.custnumber", "custtype":"g.groupvalue", "custname":"customer.custname", 
            "taxnumber":"customer.taxnumber", "account":"customer.account", "notax":"customer.notax", 
            "terms":"customer.terms", "creditlimit":"customer.creditlimit", "discount":"customer.discount", 
            "notes":"customer.notes", "inactive":"customer.inactive"};
          data._sql = {
            select:["customer.custnumber","g.groupvalue as custtype","customer.custname",
              "customer.taxnumber","customer.account","customer.notax","customer.terms",
              "customer.creditlimit","customer.discount","customer.notes","customer.inactive"],
            from:"customer", inner_join:["groups g","on",["customer.custtype","=","g.id"]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("customer.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "customer.deleted";
            data._sql.select.push("customer.deleted");}
          else {
            data._sql.where = ["customer.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = "order by id ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " customer.id");
          data.orderby_str = data.orderby_str.replace(" id", " customer.id");
          break;
        
        case "deffield":
          data.item_str = "deffield";
          fields_ = {"fieldname":"df.fieldname", "nervatype":"g.groupvalue", "subtype":"sg.groupvalue", 
            "fieldtype":"fg.groupvalue", "description":"df.description", "valuelist":"df.valuelist", 
            "addnew":"df.addnew", "visible":"df.visible", "readonly":"df.readonly"};
          data._sql = {
            select:["df.fieldname","g.groupvalue as nervatype","sg.groupvalue as subtype",
              "fg.groupvalue as fieldtype","df.description","df.valuelist","df.addnew",
              "df.visible","df.readonly"],
            from:"deffield df", 
            inner_join:[["groups g","on",["df.nervatype","=","g.id"]],
              ["groups fg","on",["df.fieldtype","=","fg.id"]]]}
          if (params.use_deleted){
            data._sql.select.push("df.deleted");
            data._sql.left_join = ["groups sg","on",["df.subtype","=","sg.id"]];
            data._sql.where = ["1","=","1"];
            fields_.deleted = "df.deleted";}
          else {
            data._sql.left_join = ["groups sg","on",[["df.subtype","=","sg.id"],
              ["and","sg.deleted","=","0"]]];
            data._sql.where = ["df.deleted","=","0"];}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("df.id");}
          if (data.orderby === "") {
            data.orderby_str = "order by fieldname ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " df.id");
          data.orderby_str = data.orderby_str.replace(" id", " df.id");
          break;
        
        case "employee":
          data.item_str = "employee"; data.fld_value = "employee";
          fields_ = {"empnumber":"employee.empnumber", "username":"employee.username", "usergroup":"g.groupvalue", 
            "startdate":"employee.startdate", "enddate":"employee.enddate", "department":"dg.groupvalue", 
            "inactive":"employee.inactive"};
          data._sql = {
            select:["employee.empnumber","employee.username","g.groupvalue as usergroup",
              "{FMS_DATE}employee.startdate{FME_DATE} as startdate",
              "{FMS_DATE}employee.enddate{FME_DATE} as enddate","dg.groupvalue as department","employee.inactive"],
            from:"employee", inner_join:["groups g","on",["employee.usergroup","=","g.id"]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("employee.id");}
          if (params.use_deleted){
            data._sql.select.push("employee.password","employee.registration_key","employee.deleted");
            data._sql.left_join = ["groups dg","on",["employee.department","=","dg.id"]];
            data._sql.where = ["1","=","1"];
            fields_.password = "employee.password";
            fields_.registration_key = "employee.registration_key";
            fields_.deleted = "employee.deleted";}
          else {
            data._sql.left_join =["groups dg","on",[["employee.department","=","dg.id"],
              ["and","dg.deleted","=","0"]]];
            data._sql.where = ["employee.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = "order by id ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " employee.id");
          data.orderby_str = data.orderby_str.replace(" id", " employee.id");
          break;
        
        case "event":
          data.item_str = "event"; data.fld_value = "event";
          if (typeof filter.nervatype !== "undefined"){
            nervanumber = nstore.valid.getTableKey(filter.nervatype);
            fields_ = {"calnumber":"event.calnumber", "nervatype":"g.groupvalue", 
              "refnumber":"nerv."+nervanumber, "uid":"event.uid", "eventgroup":"eg.groupvalue", 
              "fromdate":"event.fromdate", "todate":"event.todate", "subject":"event.subject", 
              "place":"event.place", "description":"event.description"};
            data._sql = {
              select:["event.calnumber","g.groupvalue as nervatype", "nerv."+nervanumber+" as refnumber",
                "event.uid","eg.groupvalue as eventgroup",
                "{FMS_DATETIME}event.fromdate{FME_DATETIME} as fromdate",
                "{FMS_DATETIME}event.todate{FME_DATETIME} as todate",
                "event.subject","event.place","event.description"],
              from:"event", 
              inner_join:[["groups g","on",[["event.nervatype","=","g.id"],
                ["and","g.groupvalue","=","'"+filter.nervatype+"'"]]]],
              left_join:["groups eg","on",["event.eventgroup","=","eg.id"]]}
            if (params.use_deleted){
              data._sql.inner_join.push([filter.nervatype+" nerv","on",["event.ref_id","=","nerv.id"]]);}
            else {
              data._sql.inner_join.push([filter.nervatype+" nerv","on",
                ["event.ref_id","=","nerv.id"],["and","nerv.deleted","=","0"]]);}}
          else {
            fields_ = {"calnumber":"event.calnumber", "nervatype":"g.groupvalue", "represent_refnumber_nervatype":"event.ref_id", 
              "uid":"event.uid", "eventgroup":"eg.groupvalue", "fromdate":"event.fromdate", "todate":"event.todate", 
              "subject":"event.subject", "place":"event.place", "description":"event.description"};
            data._sql = {
              select:["event.calnumber","g.groupvalue as nervatype","event.ref_id as represent_refnumber_nervatype",
                "event.uid","eg.groupvalue as eventgroup",
                "{FMS_DATETIME}event.fromdate{FME_DATETIME} as fromdate",
                "{FMS_DATETIME}event.todate{FME_DATETIME} as todate",
                "event.subject","event.place","event.description"],
              from:"event", inner_join:["groups g","on",["event.nervatype","=","g.id"]]}
            if (params.use_deleted){
              data._sql.left_join = ["groups eg","on",["event.eventgroup","=","eg.id"]];
              data.from_str = " from event \
                inner join groups g on event.nervatype=g.id left join groups eg on (event.eventgroup=eg.id) ";}
            else {
              data._sql.left_join = ["groups eg","on",[["event.eventgroup","=","eg.id"],
                ["and","eg.deleted","=","0"]]];}}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("event.id");}
          if (params.use_deleted){
            data._sql.select.push("event.deleted");
            data._sql.where = ["1","=","1"];
            fields_.deleted = "event.deleted";}
          else {
            data._sql.where = ["event.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = "order by nervatype, calnumber ";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " event.id");
          data.orderby_str = data.orderby_str.replace(" id", " event.id");
          break;
        
        case "fieldvalue":
          data.item_str = "fieldvalue";
          fields_ = {"nervatype":"g.groupvalue", "represent_fieldname":"fv.id",
            "represent_refnumber_nervatype":"fv.ref_id", "description":"df.description",
            "fieldtype":"fg.groupvalue", "represent_fieldvalue":"fv.value", "notes":"fv.notes"};
          data._sql = {
            select:["g.groupvalue as nervatype","fv.id as represent_fieldname",
              "fv.ref_id as represent_refnumber_nervatype","df.description","fg.groupvalue as fieldtype",
              "fv.value as represent_fieldvalue","fv.notes"],
            from:"fieldvalue fv",
            inner_join:[]}
          if (params.use_deleted){
            data._sql.inner_join.push(["deffield df","on",["fv.fieldname","=","df.fieldname"]]);
            data._sql.where = [["1","=","1"]];
            data._sql.select.push("fv.deleted");
            data.select_str += ", fv.deleted ";}
          else {
            data._sql.inner_join.push(["deffield df","on",[["fv.fieldname","=","df.fieldname"],
              ["and","df.deleted","=","0"]]]);
            data._sql.where = [["fv.deleted","=","0"]];}
          data._sql.inner_join.push(["groups fg","on",["df.fieldtype","=","fg.id"]],
            ["groups g","on",["df.nervatype","=","g.id"]]);
          if (filter.nervatype){
            data._sql.where.push("and","g.groupvalue","=","'"+filter.nervatype+"'");}
          if (data.orderby === "") {
            data.orderby_str = " order by g.groupvalue, fv.ref_id, fv.fieldname, fv.id ";}
          else {
            data.orderby_str = data.orderby;}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("fv.id");}
          data.where_str = data.where_str.replace(" id", " fv.id");
          data.orderby_str = data.orderby_str.replace(" id", " fv.id");
          break;
        
        case "groups":
          data.item_str = "groups"; data.fld_value = "groups";
          fields_ = {"groupname":"groups.groupname", "groupvalue":"groups.groupvalue", "description":"groups.description", 
            "inactive":"groups.inactive"};
          data._sql = {
            select:["groups.groupname","groups.groupvalue","groups.description","groups.inactive"],
            from:"groups"}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("groups.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "groups.deleted";
            data._sql.select.push("groups.deleted");}
          else {
            data._sql.where = ["groups.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by groupname, groupvalue";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " groups.id");
          data.orderby_str = data.orderby_str.replace(" id", " groups.id");
          break;
          
        case "item":
          data.item_str = "item"; data.fld_value = "item";
          fields_ = {"transnumber":"t.transnumber", "represent_rownumber":"item.id", "partnumber":"p.partnumber", 
            "unit":"item.unit", "qty":"item.qty", "fxprice":"item.fxprice", "netamount":"item.netamount", 
            "discount":"item.discount", "taxcode":"tx.taxcode", "vatamount":"item.vatamount", 
            "amount":"item.amount", "description":"item.description", "deposit":"item.deposit", 
            "ownstock":"item.ownstock", "actionprice":"item.actionprice"};
          data._sql = {
            select:["t.transnumber","item.id as represent_rownumber","p.partnumber","item.unit",
              "item.qty","item.fxprice","item.netamount","item.discount","tx.taxcode as taxcode",
              "item.vatamount","item.amount","item.description","item.deposit","item.ownstock",
              "item.actionprice"],
            from:"item", 
            inner_join:[["trans t","on",["item.trans_id","=","t.id"]],
              ["product p","on",["item.product_id","=","p.id"]],
              ["tax tx","on",["item.tax_id","=","tx.id"]],
              ["groups g","on",["t.transtype","=","g.id"]],
              ["groups gdir","on",["t.direction","=","gdir.id"]]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("item.id");}
          if (params.use_deleted){
            data._sql.where = [["1","=","1"]];
            fields_.deleted = "item.deleted";
            data._sql.select.push("item.deleted");}
          else {
            data._sql.where = [["item.deleted","=","0"],
              ["and",[["t.deleted","=","0"],
                ["or",[["g.groupvalue","=","'invoice'"],["and","gdir.groupvalue","=","'out'"]]],
                ["or",[["g.groupvalue","=","'receipt'"],["and","gdir.groupvalue","=","'out'"]]]]]];}
          if (nstore.employee()){
            data._sql.where.push(["and","t.transtype","not in",[{
              select:["subtype"], from:"ui_audit",
              inner_join:[["groups nt","on",[
                ["ui_audit.nervatype","=","nt.id"],["and","nt.groupvalue","=","'trans'"]]],
                ["groups ifa","on",[["ui_audit.inputfilter","=","ifa.id"],["and","ifa.groupvalue","=","'disabled'"]]]],
              where:["usergroup","=",nstore.employee().usergroup]}]]);}
          if (data.orderby === "") {
            data.orderby_str = " order by transnumber, id";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " item.id");
          data.orderby_str = data.orderby_str.replace(" id", " item.id");
          if (data.data_audit !== "all"){
            if (data.data_audit === "usergroup"){
              data._sql.where.push(["and","t.cruser_id","in",
                [{select:["id"], from:"employee", where:["usergroup","=",nstore.employee().usergroup]}]]);}
            else if (data.data_audit === "own") {
              data._sql.where.push(["and","t.cruser_id","=",nstore.employee().id]);}}
          break;
        
        case "link":
          data.item_str = "link"; data.fld_value = "link";
          fields_ = {"nervatype1":"g1.groupvalue", "represent_refnumber1":"link.ref_id_1", 
            "nervatype2":"g2.groupvalue", "represent_refnumber2":"link.ref_id_2"};
          data._sql = {
            select:["g1.groupvalue as nervatype1","link.ref_id_1 as represent_refnumber1",
              "g2.groupvalue as nervatype2","link.ref_id_2 as represent_refnumber2"],
            from:"link", 
            inner_join:[["groups g1","on",["link.nervatype_1","=","g1.id"]],
              ["groups g2","on",["link.nervatype_2","=","g2.id"]]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("link.id");}
          if (params.use_deleted){
            data._sql.where = [["1","=","1"]];
            fields_.deleted = "link.deleted";
            data._sql.select.push("link.deleted");}
          else {
            data._sql.where = [["link.deleted","=","0"]];}
          if (filter.hasOwnProperty("nervatype1")) {
            data._sql.where.push(["and","g1.groupvalue","=","'"+filter.nervatype1+"'"]);}
          if (filter.hasOwnProperty("nervatype2")) {
            data._sql.where.push(["and","g2.groupvalue","=","'"+filter.nervatype2+"'"]);}
          if (data.orderby === "") {
            data.orderby_str = " order by g1.groupvalue, g2.groupvalue";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " link.id");
          data.orderby_str = data.orderby_str.replace(" id", " link.id");
          break;
        
        case "log":
          data.item_str = "log"; data.fld_value = "log";
          params.use_deleted = true;
          fields_ = {"empnumber":"e.empnumber", "crdate":"log.crdate", "logstate":"lg.groupvalue", 
            "nervatype":"g.groupvalue", "represent_refnumber_nervatype":"log.ref_id"};
          data._sql = {
            select:["e.empnumber as empnumber","{FMS_DATE}log.crdate{FME_DATE} as crdate",
              "lg.groupvalue as logstate",
              "case when g.groupvalue is null then '' else g.groupvalue end as nervatype",
              "log.ref_id as represent_refnumber_nervatype"],
            from:"log", 
            left_join:["groups g","on",["log.nervatype","=","g.id"]],
            inner_join:[["groups lg","on",["log.logstate","=","lg.id"]],
              ["employee e","on",["log.employee_id","=","e.id"]]],
            where:["1","=","1"]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("log.id");}
          if (data.orderby === "") {
            data.orderby_str = " order by id";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " log.id");
          data.orderby_str = data.orderby_str.replace(" id", " log.id");
          break;
        
        case "movement":
          data.item_str = "movement"; data.fld_value = "movement";
          fields_ = {"transnumber":"t.transnumber", "represent_rownumber":"movement.id",
            "movetype":"g.groupvalue", "partnumber":"p.partnumber", "serial":"tl.serial", 
            "planumber":"pl.planumber", "shippingdate":"movement.shippingdate", 
            "qty":"movement.qty", "notes":"movement.notes", "shared":"movement.shared"};
          data._sql = {
            select:["t.transnumber","movement.id as represent_rownumber","g.groupvalue as movetype",
              "p.partnumber","tl.serial","pl.planumber",
              "{FMS_DATETIME}movement.shippingdate{FME_DATETIME} as shippingdate","movement.qty",
              "movement.notes","movement.shared"],
            from:"movement",
            inner_join:[["groups g","on",["movement.movetype","=","g.id"]]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("movement.id");}
          if (params.use_deleted){
            data._sql.inner_join.push(["trans t","on",["movement.trans_id","=","t.id"]]);
            data._sql.left_join = [
              ["product p","on",["movement.product_id","=","p.id"]],
              ["tool tl","on",["movement.tool_id","=","tl.id"]],
              ["place pl","on",["movement.place_id","=","pl.id"]]];
            data._sql.where = [["1","=","1"]];
            fields_.deleted ="movement.deleted";
            data._sql.select.push("movement.deleted");}
          else {
            data._sql.inner_join.push(["trans t","on",[["movement.trans_id","=","t.id"],
              ["and","t.deleted","=","0"]]]);
            data._sql.left_join = [
              ["product p","on",[["movement.product_id","=","p.id"],["and","p.deleted","=","0"]]],
              ["tool tl","on",[["movement.tool_id","=","tl.id"],["and","tl.deleted","=","0"]]],
              ["place pl","on",[["movement.place_id","=","pl.id"],["and","pl.deleted","=","0"]]]];
            data._sql.where = [["movement.deleted","=","0"]];}
          if (nstore.employee()){
            data._sql.where.push(["and","t.transtype","not in",[{
              select:["subtype"], from:"ui_audit",
              inner_join:[["groups nt","on",[
                ["ui_audit.nervatype","=","nt.id"],["and","nt.groupvalue","=","'trans'"]]],
                ["groups ifa","on",[["ui_audit.inputfilter","=","ifa.id"],["and","ifa.groupvalue","=","'disabled'"]]]],
              where:["usergroup","=",nstore.employee().usergroup]}]]);}
          if (data.orderby === "") {
            data.orderby_str = " order by transnumber, id";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " movement.id");
          data.orderby_str = data.orderby_str.replace(" id", " movement.id");
          if (data.data_audit !== "all"){
            if (data.data_audit === "usergroup"){
              data._sql.where.push(["and","t.cruser_id","in",
                [{select:["id"], from:"employee", where:["usergroup","=",nstore.employee().usergroup]}]]);}
            else if (data.data_audit === "own"){
              data._sql.where.push(["and","t.cruser_id","=",nstore.employee().id]);}}
          break;
        
        case "numberdef":
          data.item_str = "numberdef";
          data._sql = {
            select:["numberkey","prefix","curvalue","isyear","sep","len","description",
              "visible","readonly","orderby"],
            from:"numberdef", where:["1","=","1"]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("id");}
          if (data.orderby === "") {
            data.orderby_str = " order by id";}
          else {
            data.orderby_str = data.orderby;}
          break;
        
        case "pattern":
          data.item_str = "pattern";
          fields_ = {"description":"pattern.description", "transtype":"g.groupvalue", "notes":"pattern.notes", 
            "defpattern":"pattern.defpattern"};
          data._sql = {
            select:["pattern.description","g.groupvalue as transtype","pattern.notes","pattern.defpattern"],
            from:"pattern", inner_join:["groups g","on",["pattern.transtype","=","g.id"]]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("pattern.id");}
          if (params.use_deleted){
            data._sql.where = [["1","=","1"]];
            fields_.deleted = "pattern.deleted";
            data._sql.select.push("pattern.deleted");}
          else {
            data._sql.where = [["pattern.deleted","=","0"]];}
          if (nstore.employee()){
            data._sql.where.push(["and","pattern.transtype","not in",[{
              select:["subtype"], from:"ui_audit",
              inner_join:[["groups nt","on",[
                ["ui_audit.nervatype","=","nt.id"],["and","nt.groupvalue","=","'trans'"]]],
                ["groups ifa","on",[["ui_audit.inputfilter","=","ifa.id"],["and","ifa.groupvalue","=","'disabled'"]]]],
              where:["usergroup","=",nstore.employee().usergroup]}]]);}
          if (data.orderby === "") {
            data.orderby_str = " order by id";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " pattern.id");
          data.orderby_str = data.orderby_str.replace(" id", " pattern.id");
          break;
        
        case "payment":
          data.item_str = "payment"; data.fld_value = "payment";
          fields_ = {"transnumber":"t.transnumber", "represent_rownumber":"payment.id", 
            "paiddate":"payment.paiddate", "amount":"payment.amount", "notes":"payment.notes"};
          data._sql = {
            select:["t.transnumber","payment.id as represent_rownumber",
              "{FMS_DATE}payment.paiddate{FME_DATE} as paiddate",
              "payment.amount","payment.notes"],
            from:"payment", 
            inner_join:[["trans t","on",["payment.trans_id","=","t.id"]],
              ["groups g","on",["t.transtype","=","g.id"]]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("payment.id");}
          if (params.use_deleted){
            data._sql.where = [["1","=","1"]];
            fields_.deleted = "payment.deleted";
            data._sql.select.push("payment.deleted");}
          else {
            data._sql.where = [["payment.deleted","=","0"],["and",[["t.deleted","=","0"],["or",["g.groupvalue","=","'cash'"]]]]];}
          if (data.orderby === "") {
            data.orderby_str = " order by transnumber, id";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " payment.id");
          if (nstore.employee()){
            data._sql.where.push(["and","t.transtype","not in",[{
              select:["subtype"], from:"ui_audit",
              inner_join:[["groups nt","on",[
                ["ui_audit.nervatype","=","nt.id"],["and","nt.groupvalue","=","'trans'"]]],
                ["groups ifa","on",[["ui_audit.inputfilter","=","ifa.id"],["and","ifa.groupvalue","=","'disabled'"]]]],
              where:["usergroup","=",nstore.employee().usergroup]}]]);}
          data.orderby_str = data.orderby_str.replace(" id", " payment.id");
          if (data.data_audit !== "all"){
            if (data.data_audit === "usergroup"){
              data._sql.where.push(["and","t.cruser_id","in",
                [{select:["id"], from:"employee", where:["usergroup","=",nstore.employee().usergroup]}]]);}
            else if (data.data_audit === "own"){
              data._sql.where.push(["and","t.cruser_id","=",nstore.employee().id]);}}
          break;
        
        case "place":
          data.item_str = "place"; data.fld_value = "place";
          fields_ = {"planumber":"place.planumber", "placetype":"g.groupvalue", "description":"place.description", 
            "curr":"place.curr", "defplace":"place.defplace", "notes":"place.notes", "inactive":"place.inactive"};
          data._sql = {
            select:["place.planumber","g.groupvalue as placetype","place.description",
              "place.curr","place.defplace","place.notes","place.inactive"],
            from:"place", inner_join:["groups g","on",["place.placetype","=","g.id"]],
            left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("place.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "place.deleted";
            data._sql.select.push("place.deleted");}
          else {
            data._sql.where = ["place.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by planumber";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " place.id");
          data.orderby_str = data.orderby_str.replace(" id", " place.id");
          break;
        
        case "price":
          data.item_str = "price"; data.fld_value = "price";
          fields_ = {"partnumber":"p.partnumber", "pricetype":"case when discount is null then 'price' else 'discount' end", 
            "validfrom":"price.validfrom", "validto":"price.validto", "curr":"price.curr", "qty":"price.qty", 
            "pricevalue":"price.pricevalue", "discount":"price.discount", "calcmode":"g.groupvalue", 
            "vendorprice":"price.vendorprice"}
          data._sql = {
            select:["p.partnumber","case when discount is null then 'price' else 'discount' end as pricetype",
              "{FMS_DATE}price.validfrom{FME_DATE} as validfrom",
              "{FMS_DATE}price.validto{FME_DATE} as validto",
              "price.curr","price.qty","price.pricevalue","price.discount",
              "case when discount is null then null else g.groupvalue end as calcmode","price.vendorprice"],
            from:"price", inner_join:[], left_join:["groups g","on",["price.calcmode","=","g.id"]]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("price.id");}
          if (params.use_deleted){
            data._sql.inner_join = ["product p","on",["price.product_id","=","p.id"]];
            data._sql.where = ["1","=","1"];
            fields_.deleted = "price.deleted";
            data._sql.select.push("price.deleted");}
          else {
            data._sql.inner_join = ["product p","on",[["price.product_id","=","p.id"],["and",["p.deleted","=","0"]]]];
            data._sql.where = ["price.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by partnumber, validfrom";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " price.id");
          data.orderby_str = data.orderby_str.replace(" id", " price.id");
          break;
        
        case "product":
          data.item_str = "product"; data.fld_value = "product";
          fields_ = {"partnumber":"product.partnumber", "description":"product.description", 
            "protype":"g.groupvalue", "unit":"product.unit", "taxcode":"tax.taxcode", 
            "notes":"product.notes", "webitem":"product.webitem", "inactive":"product.inactive"}
          data._sql = {
            select:["product.partnumber","product.description","g.groupvalue as protype",
              "product.unit","tax.taxcode","product.notes","product.webitem","product.inactive"],
            from:"product", 
            inner_join:[["groups g","on",["product.protype","=","g.id"]],
              ["tax","on",["product.tax_id","=","tax.id"]]], left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("product.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "product.deleted";
            data._sql.select.push("product.deleted");}
          else {
            data._sql.where = ["product.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by partnumber";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " product.id");
          data.orderby_str = data.orderby_str.replace(" id", " product.id");
          break;
        
        case "project":
          data.item_str = "project"; data.fld_value = "project";
          fields_ = {"pronumber":"project.pronumber", "description":"project.description", 
            "custnumber":"c.custnumber", "startdate":"project.startdate", "enddate":"project.enddate", 
            "notes":"project.notes", "inactive":"project.inactive"}
          data._sql = {
            select:["project.pronumber","project.description","c.custnumber",
              "{FMS_DATE}project.startdate{FME_DATE} as startdate",
              "{FMS_DATE}project.enddate{FME_DATE} as enddate","project.notes","project.inactive"],
            from:"project", inner_join:[], left_join:[]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("project.id");}
          if (params.use_deleted){
            data._sql.left_join = ["customer c","on",["project.customer_id","=","c.id"]];
            data._sql.where = ["1","=","1"];
            fields_.deleted = "project.deleted";
            data._sql.select.push("project.deleted");}
          else {
            data._sql.left_join = ["customer c","on",[["project.customer_id","=","c.id"],["and","c.deleted","=","0"]]];
            data._sql.where = ["project.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by pronumber";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " project.id");
          data.orderby_str = data.orderby_str.replace(" id", " project.id");
          break;
        
        case "rate":
          data.item_str = "rate"; data.fld_value = "rate";
          fields_ = {"ratetype":"g.groupvalue", "ratedate":"rate.ratedate", "curr":"rate.curr", 
            "planumber":"p.planumber", "rategroup":"rg.groupvalue", "ratevalue":"rate.ratevalue"}
          data._sql = {
            select:["g.groupvalue as ratetype","{FMS_DATE}ratedate{FME_DATE} as ratedate",
              "rate.curr","p.planumber","rg.groupvalue as rategroup","rate.ratevalue"],
            from:"rate", inner_join:["groups g","on",["rate.ratetype","=","g.id"]]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("rate.id");}
          if (params.use_deleted){
            data._sql.left_join = [["place p","on",["rate.place_id","=","p.id"]],
              ["groups rg","on",["rategroup","=","rg.id"]]];
            data._sql.where = ["1","=","1"];
            fields_.deleted = "rate.deleted";
            data._sql.select.push("rate.deleted");}
          else {
            data._sql.left_join = [["place p","on",[["rate.place_id","=","p.id"],["and","p.deleted","=","0"]]],
              ["groups rg","on",[["rategroup","=","rg.id"],["and","rg.deleted","=","0"]]]];
            data._sql.where = ["rate.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by rategroup, ratedate, curr";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " rate.id");
          data.orderby_str = data.orderby_str.replace(" id", " rate.id");
          break;
        
        case "tax":
          data.item_str = "tax"; data.fld_value = "tax";
          data._sql = {
            select:["taxcode","description","rate","inactive"],
            from:"tax", inner_join:[], left_join:[],
            where:["1","=","1"]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("tax.id");}
          if (data.orderby === "") {
            data.orderby_str = " order by tax.id";}
          else {
            data.orderby_str = data.orderby;}
          break;
        
        case "tool":
          data.item_str = "tool"; data.fld_value = "tool";
          fields_ = {"serial":"tool.serial", "description":"tool.description", "partnumber":"p.partnumber", 
            "toolgroup":"g.groupvalue", "notes":"tool.notes", "inactive":"tool.inactive"};
          data._sql = {
            select:["tool.serial","tool.description","p.partnumber","g.groupvalue as toolgroup",
              "tool.notes","tool.inactive"],
            from:"tool", inner_join:["product p","on",["tool.product_id","=","p.id"]], 
            left_join:["groups g","on",["tool.toolgroup","=","g.id"]]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("tool.id");}
          if (params.use_deleted){
            data._sql.where = ["1","=","1"];
            fields_.deleted = "tool.deleted";
            data._sql.select.push("tool.deleted");}
          else {
            data._sql.where = ["tool.deleted","=","0"];}
          if (data.orderby === "") {
            data.orderby_str = " order by serial";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " tool.id");
          data.orderby_str = data.orderby_str.replace(" id", " tool.id");
          break;
        
        case "trans":
          data.item_str = "trans"; data.fld_value = "trans";
          fields_ = {"transnumber":"trans.transnumber", "transtype":"g.groupvalue", "direction":"gdir.groupvalue", 
            "ref_transnumber":"trans.ref_transnumber", "crdate":"trans.crdate", "transdate":"trans.transdate", 
            "duedate":"trans.duedate", "custnumber":"c.custnumber", "empnumber":"e.empnumber", 
            "department":"dg.groupvalue", "pronumber":"p.pronumber", "planumber":"pl.planumber", 
            "paidtype":"pg.groupvalue", "curr":"trans.curr", "notax":"trans.notax", "paid":"trans.paid", 
            "acrate":"trans.acrate", "notes":"trans.notes", "intnotes":"trans.intnotes", "fnote":"trans.fnote", 
            "transtate":"sg.groupvalue", "closed":"trans.closed", "username":"cruser.username"}
          data._sql = {
            select:["trans.transnumber","g.groupvalue as transtype","gdir.groupvalue as direction",
              "trans.ref_transnumber","{FMS_DATE}trans.crdate{FME_DATE} as crdate",
              "{FMS_DATE}trans.transdate{FME_DATE} as transdate",
              "{FMS_DATETIME}trans.duedate{FME_DATETIME} as duedate","c.custnumber",
              "e.empnumber","dg.groupvalue as department","p.pronumber","pl.planumber",
              "pg.groupvalue as paidtype","trans.curr","trans.notax","trans.paid","trans.acrate",
              "trans.notes","trans.intnotes","trans.fnote","sg.groupvalue as transtate",
              "trans.closed","cruser.username"],
            from:"trans", inner_join:[["groups g","on",["trans.transtype","=","g.id"]],
              ["groups gdir","on",["trans.direction","=","gdir.id"]],
              ["employee cruser","on",["trans.cruser_id","=","cruser.id"]],
              ["groups sg","on",["trans.transtate","=","sg.id"]]],
            left_join:[["groups pg","on",["trans.paidtype","=","pg.id"]]]}
          if (filter.hasOwnProperty("show_id") || !filter.hasOwnProperty("cross_tab")){
            data._sql.select.unshift("trans.id");}
          if (params.use_deleted){
            data._sql.left_join.push(
              ["customer c","on",["trans.customer_id","=","c.id"]],
              ["employee e","on",["trans.employee_id","=","e.id"]],
              ["groups dg","on",["trans.department","=","dg.id"]],
              ["project p","on",["trans.project_id","=","p.id"]],
              ["place pl","on",["trans.place_id","=","pl.id"]]);
            data._sql.where = [["1","=","1"]];
            fields_.deleted = "trans.deleted";
            data._sql.select.push("trans.deleted");}
          else {
            data._sql.left_join.push(
              ["customer c","on",[["trans.customer_id","=","c.id"],["and","c.deleted","=","0"]]],
              ["employee e","on",[["trans.employee_id","=","e.id"],["and","e.deleted","=","0"]]],
              ["groups dg","on",[["trans.department","=","dg.id"],["and","dg.deleted","=","0"]]],
              ["project p","on",[["trans.project_id","=","p.id"],["and","p.deleted","=","0"]]],
              ["place pl","on",[["trans.place_id","=","pl.id"],["and","pl.deleted","=","0"]]]);
            data._sql.where = [[["trans.deleted","=","0"],
                ["or",[["g.groupvalue","=","'invoice'"],["and","gdir.groupvalue","=","'out'"]]],
                ["or",[["g.groupvalue","=","'receipt'"],["and","gdir.groupvalue","=","'out'"]]],
                ["or",[["g.groupvalue","=","'cash'"]]]]];}
          if (nstore.employee()){
            data._sql.where.push(["and","trans.transtype","not in",[{
              select:["subtype"], from:"ui_audit",
              inner_join:[["groups nt","on",[
                ["ui_audit.nervatype","=","nt.id"],["and","nt.groupvalue","=","'trans'"]]],
                ["groups ifa","on",[["ui_audit.inputfilter","=","ifa.id"],["and","ifa.groupvalue","=","'disabled'"]]]],
              where:["usergroup","=",nstore.employee().usergroup]}]]);}
          if (data.orderby === "") {
            data.orderby_str = " order by transnumber";}
          else {
            data.orderby_str = data.orderby;}
          data.where_str = data.where_str.replace(" id", " trans.id");
          data.orderby_str = data.orderby_str.replace(" id", " trans.id");
          if (data.data_audit !== "all"){
            if (data.data_audit === "usergroup"){
              data._sql.where.push(["and","trans.cruser_id","in",
                [{select:["id"], from:"employee", where:["usergroup","=",nstore.employee().usergroup]}]]);}
            else if (data.data_audit === "own"){
              data._sql.where.push(["and","trans.cruser_id","=",nstore.employee().id]);}}
          break;
            
        default:
          return callback(lang.ndi_invalid_datatype+": "+params.datatype);
          break;}
          
      for (var field in fields_) {
        data.where_str = data.where_str.replace(" "+field, " "+fields_[field]).replace(","+field, ","+fields_[field]);
        data.orderby_str = data.orderby_str.replace(" "+field, " "+fields_[field]).replace(","+field, ","+fields_[field]);}
            
      callback(null, data);},
    
    function(data, callback) {
      if ((params.datatype !== "sql") && (data.fld_value !== "") && (typeof filter.no_deffield === "undefined")){
        var colname = "description";
        var _sql = {
          select:["df.fieldname as fieldname","ft.groupvalue as fieldtype",
            "df.description as description","df.valuelist as valuelist"],
          from:"deffield df",
          inner_join:[["groups nt","on",[["df.nervatype","=","nt.id"],["and","nt.groupvalue","=","'"+data.fld_value+"'"]]],
            ["groups ft","on",["df.fieldtype","=","ft.id"]]],
          where:[["df.deleted","=","0"],["and","df.visible","=","1"],["and","df.readonly","=","0"]]}
        params.validator.conn.query(models.getSql(nstore.engine(), _sql), [], function (error, deffields) {
          if (error) {callback(error);}
          else {
            for (var idx = 0; idx < deffields.rows.length; idx++) {
              var field = deffields.rows[idx]; var _sql;
              if (!filter.hasOwnProperty("cross_tab")){
                if (data.where_str.indexOf(field[colname])>-1){
                  return callback(lang.ndi_deffields_1.replace("@colname",field[colname]));}
                if (data.orderby_str.indexOf(field[colname])>-1) {
                  return callback(lang.ndi_deffields_1+field[colname]);}
                _sql = {
                  select:[], 
                  union_select:["fieldvalue.id as id","fieldvalue.ref_id as ref_id",
                    "fieldvalue.fieldname as fieldname","'"+field.description+"' as description",
                    "'"+field.fieldtype+"' as fieldtype"],
                  from:"fieldvalue", inner_join:[],
                  where:[["fieldvalue.deleted","=","0"],["and","fieldname","=","'"+field.fieldname+"'"]]}}
              else {
                data.labels[field.fieldname] = field.description;
                if(!data._sql.left_join){data._sql.left_join = [];}
                data._sql.left_join.push(
                  ["fieldvalue lj_"+idx,"on",[[data.fld_value+".id","=","lj_"+idx+".ref_id"],
                    ["and","lj_"+idx+".fieldname","=","'"+field.fieldname+"'"],
                    ["and","lj_"+idx+".deleted","=","0"]]]);}
              switch (field.fieldtype) {
                case "date":
                  if (!filter.hasOwnProperty("cross_tab")){
                    delete _sql.inner_join; _sql.union_select.push(["fieldvalue.value as value"]);}
                  else {
                    data._sql.select.push("lj_"+idx+".value as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "{CAS_DATE}lj_"+idx+".value{CAE_DATE}");
                    data.orderby_str = data.orderby_str.replace(field[colname], "{CAS_DATE}lj_"+idx+".value{CAE_DATE}");}
                  break;
                case "float":
                  if (!filter.hasOwnProperty("cross_tab")){
                    delete _sql.inner_join; _sql.union_select.push(["fieldvalue.value as value"]);}
                  else {
                    data._sql.select.push("lj_"+idx+".value as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "{CAS_FLOAT}lj_"+idx+".value{CAE_FLOAT}");
                    data.orderby_str = data.orderby_str.replace(field[colname], "{CAS_FLOAT}lj_"+idx+".value{CAE_FLOAT}");}
                  break;
                case "integer":
                  if (!filter.hasOwnProperty("cross_tab")){
                    delete _sql.inner_join; _sql.union_select.push(["fieldvalue.value as value"]);}
                  else {
                    data._sql.select.push("lj_"+idx+".value as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "{CAS_INT}lj_"+idx+".value{CAE_INT}");
                    data.orderby_str = data.orderby_str.replace(field[colname], "{CAS_INT}lj_"+idx+".value{CAE_INT}");}
                  break;
                case "customer":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["customer","on",["{CAS_INT}value{CAE_INT}","=","customer.id"]]; 
                    _sql.union_select.push(["customer.custnumber as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["customer rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id"]]);
                    data._sql.select.push("rlj_"+idx+".custname as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".custname");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".custname");}
                  break;
                case "tool":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["tool","on",["{CAS_INT}value{CAE_INT}","=","tool.id"]]; 
                    _sql.union_select.push(["tool.serial as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["tool rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id"]]);
                    data._sql.select.push("rlj_"+idx+".serial as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".serial");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".serial");}
                  break;
                case "trans":
                case "transitem":
                case "transmovement":
                case "transpayment":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["trans","on",["{CAS_INT}value{CAE_INT}","=","trans.id"]]; 
                    _sql.union_select.push(["trans.transnumber as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["trans rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id"]]);
                    data._sql.select.push("rlj_"+idx+".transnumber as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".transnumber");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".transnumber");}
                  break;
                case "product":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["product","on",["{CAS_INT}value{CAE_INT}","=","product.id"]]; 
                    _sql.union_select.push(["product.partnumber as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["product rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id "]]);
                    data._sql.select.push("rlj_"+idx+".partnumber as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".partnumber");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".partnumber");}
                  break;
                case "project":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["project","on",["{CAS_INT}value{CAE_INT}","=","project.id"]]; 
                    _sql.union_select.push(["project.pronumber as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["project rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id"]]);
                    data._sql.select.push("rlj_"+idx+".pronumber as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".pronumber");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".pronumber");}
                  break;
                case "employee":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["employee","on",["{CAS_INT}value{CAE_INT}","=","employee.id"]]; 
                    _sql.union_select.push(["employee.empnumber as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["employee rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id"]]);
                    data._sql.select.push("rlj_"+idx+".empnumber as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".empnumber");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".empnumber");}
                  break;
                case "place":
                  if (!filter.hasOwnProperty("cross_tab")){
                    _sql.inner_join = ["place","on",["{CAS_INT}value{CAE_INT}","=","place.id"]]; 
                    _sql.union_select.push(["place.planumber as value"]);}
                  else {
                    if(!data._sql.left_join){data._sql.left_join = [];}
                    data._sql.left_join.push(["place rlj_"+idx,"on",
                      ["{CAS_INT}lj_"+idx+".value{CAE_INT}","=","rlj_"+idx+".id"]]);
                    data._sql.select.push("rlj_"+idx+".planumber as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "rlj_"+idx+".planumber");
                    data.orderby_str = data.orderby_str.replace(field[colname], "rlj_"+idx+".planumber");}
                  break;
                default:
                  if (!filter.hasOwnProperty("cross_tab")){
                    delete _sql.inner_join; _sql.union_select.push(["fieldvalue.value as value"]);}
                  else {
                    data._sql.select.push("lj_"+idx+".value as "+field.fieldname);
                    data.where_str = data.where_str.replace(field[colname], "lj_"+idx+".value");
                    data.orderby_str = data.orderby_str.replace(field[colname], "lj_"+idx+".value");}
                  break;}
              if (!filter.hasOwnProperty("cross_tab")){
                if (_sql){
                  _sql.union_select.push(["fieldvalue.notes as notes"]);
                  data.fvalue_sql.push(_sql);}}}}
          callback(null, data);});}
      else {
        callback(null, data);}},
    
    function(data, callback) {
      var sql;
      if (params.datatype === "sql"){
        sql = filter.sql;}
      else {
        sql = models.getSql(nstore.engine(),data._sql)+" "+models.getSql(nstore.engine(),data.where_str)
          +" "+models.getSql(nstore.engine(),data.orderby_str)+" "+data.limit_str;}
      params.validator.conn.query(sql, [], function (error, items) {
        if (error) {callback(error);}
        else {
          data.items = items.rows;
          callback(null, data);}});},
    
    function(data, callback) {
      if (data.items.length > 0 && data.fvalue_sql.length>0) {
        var ids = [[]];
        data.items.forEach(function(row) {
          ids.push(row.id);});
        for (var i = 0; i < data.fvalue_sql.length; i++) {
          data.fvalue_sql[i].where.push(["and","ref_id","in",ids]);
          if(i === 0){
            data.fvalue_sql[i].select = data.fvalue_sql[i].union_select;
            delete data.fvalue_sql[i].union_select;}
          else {
            delete data.fvalue_sql[i].select;}}
        data.fvalue_sql[data.fvalue_sql.length-1].order_by = ["ref_id","fieldname","id"];
        params.validator.conn.query(models.getSql(nstore.engine(),data.fvalue_sql), [], function (error, items) {
          if (error) {callback(error);}
          else {
            items.rows.forEach(function(item) {
              if (typeof data.fieldvalues[item.ref_id] === "undefined"){
                data.fieldvalues[item.ref_id] = {};}
              if (typeof data.fieldvalues[item.ref_id][item.fieldname] === "undefined"){
                data.fieldvalues[item.ref_id][item.fieldname] = [];}
              data.fieldvalues[item.ref_id][item.fieldname].push(item);});
            callback(null, data);}});}
      else {
        callback(null, data);}},
    
    function(data, callback) {
      if (data.items.length > 0) {
        var rep_cols = [];
        var fields = Object.keys(data.items[0]);
        fields.forEach(function(field) {
          if (field.indexOf("represent_") > -1){
            if (typeof represents[field] !== "undefined") {
              rep_cols.push(field);}}});
        if (rep_cols.length > 0) {
          var rep_lst = [];
          data.items.forEach(function(item) {
            rep_cols.forEach(function(field) {
              rep_lst.push(function(callback_){
                represents[field].function({conn:params.validator.conn, row:item, repkey:field, 
                  datatype:params.datatype, use_deleted: (typeof params.use_deleted !== "undefined")}, 
                  function(err, row, repkey, refnumber){
                    row[repkey] = refnumber;
                    if (err) {
                      row["__deleted__"] = true;}
                    if (typeof represents[repkey].label !== "undefined") {
                       if (filter.output==="json" || filter.output==="xml"){
                         data.repfields[repkey] = represents[repkey].label.replace("*","");}
                       else {
                         data.repfields[repkey] = represents[repkey].label;}}
                    callback_(null,refnumber);});});});});
          if (rep_lst.length>0){
            async.series(rep_lst,function(err, refnumbers) {
              callback(null, data);});}}
        else {
          callback(null, data);}}
      else {
        callback(null, data);}},
    
    function(data, callback) {
      if (data.items.length > 0) {
        if (typeof filter.columns !== "undefined" && filter.columns !== ""){
          var columns = filter.columns.split(",");
          columns.forEach(function(column) {
            if (typeof data.items[0][column] !== "undefined"){
              data.columns[column] = true;}});
          if (data.columns.length === 0) {
            data.columns = data.items[0];}}
        else {
          data.columns = data.items[0];}
        if (typeof filter.header !== "undefined" && filter.header !== ""){
          var headers = filter.header.split(",");
          for (var ci = 0; ci < headers.length; ci++) {
            if (typeof Object.keys(data.columns)[ci] !== "undefined") {
              data.labels[Object.keys(data.columns)[ci]] = headers[ci];}}}
        callback(null, data);}
      else {
        callback(null, data);}},
    
    function(data, callback) {
      var result = {id:"get_"+params.datatype,
        datatype:data.item_str, show_id:filter.hasOwnProperty("show_id"), 
        cross_tab:filter.hasOwnProperty("cross_tab"), data:[]};
      if (data.items.length > 0) {
        if (filter.hasOwnProperty("cross_tab")){
          data.items.forEach(function(row) {
            if (!row.hasOwnProperty("__deleted__")) {
              var item = {};
              for (var field in row) {
                if (data.columns.hasOwnProperty(field)) {
                  if (data.labels.hasOwnProperty(field)) {
                    item[data.labels[field]] = row[field];}
                  else if (data.repfields.hasOwnProperty(field)) {
                    item[data.repfields[field]] = row[field];}
                  else {
                    item[field] = row[field];}}}
              result.data.push(item);}});}
        else {
          data.items.forEach(function(row) {
            if (!row.hasOwnProperty("__deleted__")) {
              var item = []; var id;
              if (filter.hasOwnProperty("show_id") && row.hasOwnProperty("id")){
                id = row.id;}
              for (var field in row) {
                if (field !== "id"){
                  var item_field = {name:field, value:row[field]};
                  if (id && (filter.output==="csv" || filter.output==="html")){
                    item_field = {id:id, name:field, value:row[field]};}
                  if (data.repfields.hasOwnProperty(field)) {
                    item_field.name = data.repfields[field];
                    if (represents[field].hasOwnProperty("rettype")) {
                      item_field.type = represents[field].rettype;}}
                  if (ntura.model.hasOwnProperty(params.datatype)){
                    if (ntura.model[params.datatype].hasOwnProperty(field)){
                      if (ntura.model[params.datatype][field].hasOwnProperty("references")){
                        item_field.type = ntura.model[params.datatype][field].references[0];}
                      else {
                        item_field.type = ntura.model[params.datatype][field].type;}}
                    else if (!item_field.type){
                      item_field.type = "reference";}}
                  if (data.labels.hasOwnProperty(field)) {
                    item_field.label = data.labels[field];}
                  item.push(item_field);}}
              if (id && (filter.output==="json" || filter.output==="xml")){
                item.unshift({name:"id", value:row.id, type:"id"});}
              if (data.fieldvalues.hasOwnProperty(row.id)) {
                for (var fieldname in data.fieldvalues[row.id]) {
                  for (var index = 0; index < data.fieldvalues[row.id][fieldname].length; index++) {
                    var fieldvalue = data.fieldvalues[row.id][fieldname][index];
                    if (id && (filter.output==="csv" || filter.output==="html")){
                      item.push({id:id, name:fieldname, value:fieldvalue.value, type:fieldvalue.fieldtype,
                        label:fieldvalue.description, index:index+1, data:fieldvalue.notes});}
                    else {
                      item.push({name:fieldname, value:fieldvalue.value, type:fieldvalue.fieldtype,
                        label:fieldvalue.description, index:index+1, data:fieldvalue.notes});}}}}
              if (filter.output==="csv"){
                item.forEach(function(field) {
                  result.data.push(field);});}
              else {result.data.push(item);}}});}}
      callback(null, result);}
    
    ],
  function(err, result) {
    if(err){if(err.message){err = err.message;}}
    //if (params.validator.conn !== null){
    //  params.validator.conn.close();}
    if (err !== null){
      _callback(err);}
    else {_callback(null, result);}});}
  
return {
  getError: function(id, ekey, message, data){
    return get_error(id, ekey, message, data);},
  decodeData: function(data){
	  return decode_data(data);},
  getData: function(nstore, params, filter, _callback) {
    get_data(nstore, params, filter, _callback);},
  updateData: function(nstore, params, data, _callback) {
    params.type = "update"
    update_data(nstore, params, data, _callback);},
  deleteData: function(nstore, params, data, _callback) {
    params.type = "delete";
    update_data(nstore, params, data, _callback);}
  };
};