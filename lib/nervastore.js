/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var async = require("async");
var jwt = require('jsonwebtoken');

var out = require('./tools.js').DataOutput();
var ntura = require('./models.js');
var connect = require('./adapter.js').connect();
var models = require('./adapter.js').models();

//validators

function valid_fieldvalue(nervatype){
  if (nervatype.substring(0,2) === "ui") {
    return false;}
  else {
    return (["groups", "numberdef", "deffield", "pattern", "fieldvalue"].indexOf(nervatype)===-1);}}

function date_validator(value){
  if (value === null || value === "") {return value;}
  try {
    var ddate = [];
    if (value.indexOf(".")>-1){ddate = value.split(".");}
    else if (value.indexOf("-")>-1){ddate = value.split("-");}
    else if (value.indexOf("/")>-1){ddate = value.split("/");}
    else {return null;}
    if (ddate.length === 3) {
      if (parseInt(ddate[0],10)<1900 || parseInt(ddate[0],10)>2100){
        return null;}
      else if (parseInt(ddate[1],10)<1 || parseInt(ddate[1],10)>12){
        return null;}
      else if (parseInt(ddate[2],10)<1 || parseInt(ddate[2],10)>31){
        return null;}
      var cdate = new Date(parseInt(ddate[0],10),parseInt(ddate[1],10)-1,parseInt(ddate[2],10));
      return out.getISODate(cdate);}
    else {
      return null;}
  } catch (error) {return null;}}

function time_validator(value){
  if (value === null || value === "") {return value;}
  try {
    var tdate = [];
    if (value.indexOf(":")>-1){tdate = value.split(":");}
    else {return null;}
    if (tdate.length === 2) {
      if (parseInt(tdate[0],10)<0 || parseInt(tdate[0],10)>23){
        return null;}
      else if (parseInt(tdate[1],10)<0 || parseInt(tdate[1],10)>59){
        return null;}
      return out.zeroPad(parseInt(tdate[0],10),2)+":"+out.zeroPad(parseInt(tdate[1],10),2);}
    else {
      return null;}

  } catch (error) {return null;}}

function set_password_field(fieldname, value) {
  if (value === null || value === "") {return "";}
  try {
    var key = fieldname+"********"; key = key.substring(0,8);
    var crypted = out.cryptedValue(key, value, "base64");
    return "XXX"+crypted+"XXX";
  } catch (error) {
    return value;}}

function get_password_field(fieldname, value) {
  if (value === null || value === "" || value.substring(0,3) !== "XXX" || value[value.length-3] !== "XXX") {
    return value;}
  try {
    value = value.substr(4,value.length-3);
    var key = fieldname+"********"; key = key.substring(0,8);
    var dec = out.decipherValue(key, value, "base64");
    return dec;
  } catch (error) {
    return value;}}

function get_table_key(nervatype) {
  if (ntura.model.hasOwnProperty(nervatype)){
    if(ntura.model[nervatype]._key.length===1){
      return ntura.model[nervatype]._key[0];}
    else {return "";}}
  else {
    return "";}}

function get_refnumber(self, params, _callback){
  if (typeof params.rettype === "undefined"){
    params.rettype = "refnumber";}
  if (typeof params.use_deleted === "undefined"){
    params.use_deleted = false;}
  var conn = params.conn;
  if (!params.conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  
  async.waterfall([
    function(callback) {
      if (!params.nervatype || !params.ref_id){
        callback(self.lang.ndi_missing_refnumber,null);}
      else {
        var _sql;
        switch (params.nervatype) {
          case "address":
          case "contact":
            _sql = {select:["nt.groupvalue as head_nervatype","t.*"],
              from:params.nervatype+" t", 
              inner_join:["groups nt","on",["t.nervatype","=","nt.id"]],
              where:[["t.id","=",params.ref_id]]};
            if (params.use_deleted!==true){
              _sql.where.push(["and","t.deleted","=","0"]);}
            break;
            
          case "fieldvalue":
          case "setting":
            _sql = {select:["fv.*","nt.groupvalue as head_nervatype"], from:"fieldvalue fv",
              inner_join: [["deffield df","on",["fv.fieldname","=","df.fieldname"]],
                ["groups nt","on",["df.nervatype","=","nt.id"]]],
              where:[["fv.id","=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","fv.deleted","=","0"]);}
            break;
          
          case "item":
          case "payment":
          case "movement":
            _sql = {select:["ti.*","t.transnumber","tt.groupvalue as transtype"],
              from:params.nervatype+" ti", 
              inner_join: [["trans t","on",["ti.trans_id","=","t.id"]], 
                ["groups tt","on",["t.transtype","=","tt.id"]]], 
              where:[["ti.id","=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and",[["t.deleted","=","0"],
                ["or","tt.groupvalue","in",[[],"'cash'","'invoice'","'receipt'"]]]]);}
            break;
          
          case "price":
            _sql = {select:["pr.*","p.partnumber"], from: "price pr",
              inner_join:["product p","on",["pr.product_id","=","p.id"]],
              where:[["pr.id","=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","p.deleted","=","0"]);}
            break;
          
          case "link":
            _sql = {select: ["l.*","nt1.groupvalue as nervatype1","nt2.groupvalue as nervatype2"],
              from:"link l",
              inner_join:[["groups nt1","on",["l.nervatype_1","=","nt1.id"]],
                ["groups nt2","on",["l.nervatype_2","=","nt2.id"]]],
              where: [["l.id","=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","l.deleted","=","0"]);}
            break;
          
          case "rate":
            _sql = {select: ["r.*","rt.groupvalue as rate_type","p.planumber"],
              from:"rate r", inner_join:["groups rt","on",["r.ratetype","=","rt.id"]],
              left_join:["place p","on",["r.place_id","=","p.id"]], 
              where:[["r.id","=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","r.deleted","=","0"]);}
            break;
          
          case "log":
            _sql = {select:["l.*","e.empnumber"], from:"log l", 
              inner_join:["employee e","on",["l.employee_id","=","e.id"]], 
              where:["l.id","=",params.ref_id]}
            break;
                    
          default:
            _sql = {select:["*"], from:params.nervatype, 
              where:[["id","=",params.ref_id]]}
            if (params.use_deleted!==true && ntura.model[params.nervatype].hasOwnProperty("deleted")){
              _sql.where.push(["and","deleted","=","0"]);}
            break;}
        conn.query(models.getSql(self.engine,_sql), [], function (error, data) {
          if (error) {
            callback(error, null);}
          else {
            if (data.rowCount===0) {
              callback(self.lang.ndi_missing_refnumber, null);}
            else {
              callback(null, data.rows[0]);}}});}},
    
    function(data, callback) {
      var index = 1; var _sql;
      switch (params.nervatype) {
        case "address":
        case "contact":
          //ref_nervatype/refnumber~rownumber
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            _sql = {select:["count(*) as count"], from:params.nervatype,
               where:[["nervatype","=",data.nervatype],["and","ref_id","=",data.ref_id],
                 ["and","id","<=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","deleted","=","0"]);}
            conn.query(models.getSql(self.engine,_sql), [], function (error, index_data) {
              if (error) {
                callback(error, null);}
              else {
                if (index_data.rowCount===0) {
                  callback(self.lang.ndi_missing_refnumber, null);}
                else {
                  index = index_data.rows[0].count;
                  if (params.rettype === "index"){
                    callback(null,index);}
                  else {
                    get_refnumber(self, {rettype:"refnumber", nervatype:data.head_nervatype, ref_id:data.ref_id, 
                      use_deleted:params.use_deleted}, function(error, refnumber){
                        if (error) {
                          callback(error, null);}
                        else {
                          callback(null, data.head_nervatype+"/"+refnumber+"~"+index);}});}}}});}
          break;
        
        case "fieldvalue":
        case "setting":
          //refnumber~~fieldname~rownumber
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else if(!data.ref_id || data.head_nervatype==="setting"){
            //setting
            switch (params.rettype) {
              case "index":
                callback(null,index);
                break;
              case "fieldname":
                callback(null,data.fieldname);
                break;
              default:
                //refnumber
                callback(null,data.fieldname);
                break;}}
          else {
            _sql = {select:["count(*) as count"], from:"fieldvalue",
              where:[["fieldname","=","'"+data.fieldname+"'"],["and","ref_id","=",data.ref_id],
                ["and","id","<=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","deleted","=","0"]);}
            conn.query(models.getSql(self.engine,_sql), [], function (error, index_data) {
              if (error) {
                callback(error, null);}
              else {
                if (index_data.rowCount===0) {
                  callback(self.lang.ndi_missing_refnumber, null);}
                else {
                  index = index_data.rows[0].count;
                  if (params.rettype === "index"){
                    callback(null,index);}
                  if (params.rettype === "fieldname"){
                    callback(null,data.fieldname+"~"+index);}
                  else {
                    get_refnumber(self, {rettype:"refnumber", nervatype:data.head_nervatype, ref_id:data.ref_id, 
                      use_deleted:params.use_deleted}, function(error, refnumber){
                        if (error) {
                          callback(error, null);}
                        else {
                          callback(null, refnumber+"~~"+data.fieldname+"~"+index);}});}}}});}
          break;
        
        case "groups":
          //groupname~groupvalue
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            if (params.rettype === "index"){
              callback(null,index);}
            else {
              callback(null, data.groupname+"~"+data.groupvalue);}}
          break;
        
        case "item":
        case "payment":
        case "movement":
          //refnumber~rownumber
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            _sql = {select:["count(*) as count"], from:params.nervatype,
              where:[["trans_id","=",data.trans_id],["and","id","<=",params.ref_id]]}
            if (params.use_deleted!==true){
              _sql.where.push(["and","deleted","=","0"]);}
            conn.query(models.getSql(self.engine,_sql), [], function (error, index_data) {
              if (error) {
                callback(error, null);}
              else {
                if (index_data.rowCount===0) {
                  callback(self.lang.ndi_missing_refnumber, null);}
                else {
                  index = index_data.rows[0].count;
                  if (params.rettype === "index"){
                    callback(null,index);}
                  else {
                    callback(null, data.transnumber+"~"+index);}}}});}
          break;
        
        case "price":
          //partnumber~pricetype~validfrom~curr~qty
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            if (params.rettype === "index"){
              callback(null,index);}
            else {
              var pricetype = "price";
              if (data.discount){pricetype = "discount";}
              var validfrom = data.validfrom;
              if (validfrom.length > 10){
                validfrom = data.validfrom.substr(0,10);}
              callback(null, data.partnumber+"~"+pricetype+"~"+validfrom+"~"+data.curr+"~"+data.qty);}}
          break;      
        
        case "link":
          //nervatype_1~refnumber_1~~nervatype_2~refnumber_2
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            if (params.rettype === "index"){
              callback(null,index);}
            else {
              get_refnumber(self, {rettype:"refnumber", nervatype:data.nervatype1, ref_id:data.ref_id_1, 
                use_deleted:params.use_deleted}, function(error, refnumber){
                  if (error) {callback(error, null);}
                  else {
                    var refnumber_1 = data.nervatype1+"~"+refnumber;
                    get_refnumber(self, {rettype:"refnumber", nervatype:data.nervatype2, ref_id:data.ref_id_2, 
                      use_deleted:params.use_deleted}, function(error, refnumber){
                        if (error) {callback(error, null);}
                        else {
                          callback(null,refnumber_1+"~~"+data.nervatype2+"~"+refnumber);}});}});}}
          break; 
        
        case "rate":
          //ratetype~ratedate~curr~planumber
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            if (params.rettype === "index"){
              callback(null,index);}
            else {
              var ratedate = data.ratedate;
              if (ratedate.length > 10){
                ratedate = data.ratedate.substr(0,10);}
              var refnumber = data.rate_type+"~"+ratedate+"~"+data.curr;
              if (data.planumber){
                refnumber += "~"+data.planumber;}
              callback(null, refnumber);}}
          break;
          
        case "log":
          //empnumber~crdate
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            if (params.rettype === "index"){
              callback(null,index);}
            else {
              callback(null, data.empnumber+"~"+data.crdate);}}
          break;
              
        default:
          //table_key
          if (params.retfield && data.hasOwnProperty(params.retfield)){
            callback(null, data[params.retfield]);}
          else {
            if (params.rettype === "index"){
              callback(null,index);}
            else {
              callback(null, data[get_table_key(params.nervatype)]);}}
          break;}}
  ],
  function(err, refnumber) {
    if(err){if(err.message){err = err.message;}}
    if (!params.conn && conn){
      conn.close();}
    _callback(err, refnumber);});}

function check_fieldvalue(self, params, _callback) {
  var conn = params.conn;
  if (!params.conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  async.waterfall([
    function(callback) {
      var _sql = {select:["ft.groupvalue as fieldtype"], from:"deffield df",
        inner_join:["groups ft","on",["df.fieldtype","=","ft.id"]],
        where:["df.fieldname","=","'"+params.fieldname+"'"]}
      conn.query(models.getSql(self.engine, _sql), [], function (error, data) {
        if (error) {
          callback(error, params.value);}
        else {
          if (data.rowCount===0) {
            callback(self.lang.missing_fieldname, params.value);}
          else {
            callback(null, data.rows[0].fieldtype);}}});},
    
    function(fieldtype, callback) {
      switch (fieldtype) {
        case "bool":
          if (["true", "True", "TRUE", "t", "T", "y", "YES", "yes", 1, "1"].indexOf(params.value)>-1){
            callback(null, "true");}
          else {callback(null, "false");}
          break;
        
        case "integer":
          if (isNaN(parseInt(params.value,10))) {
            callback(self.lang.invalid_value+": "+fieldtype, params.value);}
          else {
            callback(null, parseInt(params.value,10));}
          break;
        
        case "float":
          if (isNaN(parseFloat(params.value))) {
            callback(self.lang.invalid_value+": "+fieldtype, params.value);}
          else {
            callback(null, parseFloat(params.value));}
          break;
        
        case "date":
          var date_value = date_validator(params.value);
          if (date_value === null) {
            callback(self.lang.invalid_value+": "+fieldtype, params.value);}
          else {callback(null, date_value);}
          break;
        
        case "time":
          var time_value = time_validator(params.value);
          if (time_value === null) {
            callback(self.lang.invalid_value+": "+fieldtype, params.value);}
          else {callback(null, time_value);}
          break;
          
        case "password":
          if(params.value !== null){
            if(params.value.toString().length >= 6){
              if (params.value.substring(0,3) !== "XXX" || params.value[params.value.length-3] !== "XXX") {
                params.value = set_password_field(params.fieldname, params.value);}}}
          callback(null, params.value);
          break;
          
        case "string":
        case "valuelist":
        case "notes": 
        case "urlink":
          callback(null, params.value);
          break;
        
        case "customer":
        case "tool":
        case "product":
        case "project":
        case "employee":
        case "place":
        case "transitem":
        case "transmovement":
        case "transpayment":
          var _sql; var prm = [parseInt(params.value,10), params.value];
          if (fieldtype==="transitem" || fieldtype==="transmovement" || fieldtype==="transpayment"){
            _sql = {select:["id"], from:"trans",
            where:[["id","=","?"],["or","transnumber","like","?"]]};}
          else {
            _sql = {select:["id"], from:fieldtype,
            where:[["id","=","?"],["or",get_table_key(fieldtype),"like","?"]]};}
          if (isNaN(parseInt(params.value,10))){
            _sql.where = [get_table_key(fieldtype),"like","?"]
            prm = [params.value];}
          conn.query(models.getSql(self.engine, _sql), prm, 
            function (error, data) {
            if (error) {
              callback(error, params.value);}
            else {
              if (data.rowCount===0) {
                callback(self.lang.invalid_value+": "+fieldtype, params.value);}
              else {
                callback(null, data.rows[0].id);}}});
          break;
              
        default:
          callback(self.lang.invalid_value+": "+fieldtype, params.value);
          break;}
    }
  ],
  function(err, value) {
    if(err){if(err.message){err = err.message;}}
    if (!params.conn && conn){
      conn.close();}
    _callback(err, value, params.ref_id);});}

function get_id_from_refnumber(self, params, _callback){
  if (typeof params.nervatype === "undefined") {params.nervatype = null;}
  if (typeof params.refnumber === "undefined") {params.refnumber = null;}
  if (typeof params.use_deleted === "undefined") {params.use_deleted = false;}
  
  function set_vsql(nervatype, refnumber) {
    var md_1 = {deffield:"fieldname", employee:"empnumber", 
      pattern:"description", project:"pronumber", tool:"serial"}
    var md_2 = {currency:"curr", numberdef:"numberkey", ui_language:"lang", 
      ui_report:"reportkey", ui_menu:"menukey"}
      
    var vsql = null; var err = null; var ref_index = 0;
    if (typeof md_1[nervatype] !== "undefined") {
      vsql = {select:["id"], from:nervatype, where:[[md_1[nervatype],"=","?"]]};
      if (params.use_deleted === false) {
        vsql.where.push(["and","deleted","=","0"]);}}
    else if (typeof md_2[nervatype] !== "undefined") {
      vsql = {select:["id"], from:nervatype, where:[[md_2[nervatype],"=","?"]]};}
    else {
      switch (nervatype) {
        case "address":
        case "contact":
          //ref_nervatype/refnumber~rownumber
          if (refnumber.split("/").length > 1) {
            var ref_nervatype = refnumber.split("/")[0];
            refnumber = refnumber.substr(ref_nervatype.length+1);
            if (refnumber.split("~").length > 1) {
              ref_index = parseInt(refnumber.split("~")[1],10);
              refnumber = refnumber.split("~")[0];
              if (isNaN(ref_index) || ref_index < 1){
                err = self.lang.invalid_refnumber;}
              else {ref_index -= 1;}}
            if (err === null && ['customer', 'employee', 'event', 'place', 'product', 
              'project', 'tool', 'trans'].indexOf(ref_nervatype)>-1){
              vsql = {select:nervatype+".id as id", from:nervatype,
                inner_join:[["groups nt","on",[[nervatype+".nervatype","=","nt.id"],
                 ["and","nt.groupname","=","'nervatype'"],["and","nt.groupvalue","=","'"+ref_nervatype+"'"]]],
                 [ref_nervatype,"on",[[nervatype+".ref_id","=",ref_nervatype+".id"],
                 ["and",ref_nervatype+"."+get_table_key(ref_nervatype),"=","?"]]]]}
              if (params.use_deleted === false){
                vsql.where = [[ref_nervatype+".deleted","=","0"],["and",nervatype+".deleted","=","0"]]}}
            else {err = self.lang.invalid_refnumber;}}
          else {
            err = self.lang.invalid_refnumber;}
          break;
        
        case "barcode":
          //code
          if (params.use_deleted === false) {
            vsql = {select:["barcode.id"], from:"barcode", 
              inner_join:["product","on",["barcode.product_id","=","product.id"]],
              where:[["product.deleted","=","0"],["and","barcode.code","=","?"]]};}
          else {
            vsql = {select:["id"], from:"barcode", where:["code","=","?"]};}
          break;
        
        case "customer":
          //custnumber
          if (params.extra_info){
            vsql = [
              {select:["c.id as id","ct.groupvalue as custtype","c.terms as terms",
                "c.custname as custname","c.taxnumber as taxnumber","addr.zipcode as zipcode",
                "addr.city as city","addr.street as street"],
              from:"customer c", 
              inner_join:["groups ct","on",["c.custtype","=","ct.id"]],
              left_join:[[
                [{select:["*"], from:"address", where:["id","=",[
                  {select:["min(id)"], from:"address", 
                    where: [["deleted","=","0"],
                      ["and","nervatype","=",[
                    {select:["id"], from:"groups", 
                    where:[["groupname","=","'nervatype'"], 
                        ["and","groupvalue","=","'customer'"]]}]], 
                      ["and","ref_id","=",[
                    {select:["min(c.id)"], from:"customer c", 
                    inner_join:["groups ct","on",[["c.custtype","=","ct.id"],["and","groupvalue","=","'own'"]]],
                    where:["c.deleted","=","0"]}]]]}]]}],"addr"],"on",["c.id","=","addr.ref_id"]],
              where:["c.id","=",[
                {select:["min(c.id)"], from:"customer c", 
                inner_join: ["groups ct","on",[["c.custtype","=","ct.id"],["and","groupvalue","=","'own'"]]], 
                where:["c.deleted","=","0"]}]]},
              {union_select:["c.id as id","ct.groupvalue as custtype","c.terms as terms",
                "c.custname as custname","c.taxnumber as taxnumber","addr.zipcode as zipcode","addr.city as city",
                "addr.street as street"], from:"customer c", 
                inner_join:["groups ct","on",["c.custtype","=","ct.id"]],
                left_join:[[
                  [{select:["*"], from:"address", 
                  where:["id","=",
                    [{select:["min(id)"], from:"address", 
                      where: [
                        ["deleted","=","0"],
                        ["and","nervatype","=",[
                          {select:["id"], from:"groups", 
                          where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]],
                        ["and","ref_id","=",[
                          {select:["id"], from:"customer", 
                          where:["custnumber","=","'"+params.refnumber+"'"]}]]]}]]}],"addr"],"on",["c.id","=","addr.ref_id"]],
              where:[["c.custnumber","=","?"]]}];
            if (params.use_deleted === false) {
              vsql[1].where.push(["and","c.deleted","=","0"]);}}
          else {
            vsql = {select:["c.id as id","ct.groupvalue as custtype"], from:"customer c",
              inner_join:["groups ct","on",["c.custtype","=","ct.id"]], 
              where:[["c.custnumber","=","?"]]};
            if (params.use_deleted === false) {
              vsql.where.push(["and","c.deleted","=","0"]);}}
          break;
          
        case "event":
          //calnumber
          vsql = {select:["e.id as id","ntype.groupvalue as ref_nervatype"], 
            from: "event e",
            inner_join:["groups ntype","on",["e.nervatype","=","ntype.id"]],
            where:[["e.calnumber","=","?"]]}
          if (params.use_deleted === false) {
            vsql.where.push(["and","e.deleted","=","0"]);}
          break;
        
        case "groups":
          //groupname~groupvalue
          if (refnumber.split("~").length > 1) {
            var groupname = refnumber.split("~")[0];
            refnumber = refnumber.split("~")[1];
            vsql = {select:"id", from:nervatype, 
              where:[["groupname","=","'"+groupname+"'"],["and","groupvalue","=","?"]]}
            if (params.use_deleted === false) {
              vsql.where.push(["and","deleted","=","0"]);}}
          else {
            err = self.lang.invalid_refnumber;}
          break;
        
        case "fieldvalue":
          //refnumber~~fieldname~rownumber
          if (params.refnumber.split("~~").length > 1) {
            refnumber = params.refnumber.split("~~")[params.refnumber.split("~~").length-1].split("~")[0];
            vsql = {select:["nt.groupvalue as ref_nervatype"],
              from:"deffield df", inner_join: ["groups nt","on",["df.nervatype","=","nt.id"]],
              where: [["df.fieldname","=","?"]]}
            if (params.use_deleted === false) {
              vsql.where.push(["and","df.deleted","=","0"]);}}
          else {
            //setting
            params.nervatype = "setting";
            var rvsql = set_vsql(params.nervatype, params.refnumber);
            return {err:rvsql.err, vsql:rvsql.vsql, refnumber:rvsql.refnumber, 
              ref_index:rvsql.ref_index};}
          break;
            
        case "item":
        case "payment":
        case "movement":
          //refnumber~rownumber
          if (refnumber.split("~").length > 1) {
            ref_index = parseInt(refnumber.split("~")[1],10);
            if (isNaN(ref_index) || ref_index < 1){
              err = self.lang.invalid_refnumber;}
            else {ref_index -= 1;}}
          refnumber = refnumber.split("~")[0];
          if (nervatype === "item"){
            vsql = {
              select:["it.id as id","ttype.groupvalue as transtype","dir.groupvalue as direction",
                "cu.digit as digit","it.qty as qty","it.discount as discount", "it.tax_id as tax_id", 
                "ta.rate as rate"], 
              from:"item it",
              inner_join:[
                ["trans t","on",["it.trans_id","=","t.id"],["and","t.transnumber","=","?"]],
                ["tax ta","on",["it.tax_id","=","ta.id"]],
                ["groups ttype","on",["t.transtype","=","ttype.id"]],
                ["groups dir","on",["t.direction","=","dir.id"]]],
              left_join:["currency cu","on",["t.curr","=","cu.curr"]]};}
          else if (nervatype === "payment"){
            vsql = {
              select:["it.id as id","ttype.groupvalue as transtype","dir.groupvalue as direction"], 
              from:"payment it",
              inner_join:[["trans t","on",[["it.trans_id","=","t.id"],["and","t.transnumber","=","?"]]],
                ["groups ttype","on",["t.transtype","=","ttype.id"]], 
                ["groups dir","on",["t.direction","=","dir.id"]]]}}
          else {
            vsql = {
              select:["it.id as id","ttype.groupvalue as transtype","dir.groupvalue as direction",
                "mt.groupvalue as movetype"],
              from:"movement it",
              inner_join:[["groups mt","on",["it.movetype","=","mt.id"]],
                ["trans t","on",[["it.trans_id","=","t.id"],["and","t.transnumber","=","?"]]],
                ["groups ttype","on",["t.transtype","=","ttype.id"]],
                ["groups dir","on",["t.direction","=","dir.id"]]]};}
          if (params.use_deleted === false) {
            vsql.where = [["t.deleted","=","0"],["and","it.deleted","=","0"]];}
          break;
          
        case "price":
          //partnumber~pricetype~validfrom~curr~qty
          if (refnumber.split("~").length === 5) {
            var pricetype = refnumber.split("~")[1];
            var validfrom = refnumber.split("~")[2]; var curr = refnumber.split("~")[3]; 
            var qty = parseFloat(refnumber.split("~")[4]); 
            if(isNaN(qty)){err = self.lang.invalid_refnumber;}
            if (pricetype === "price") {
              pricetype = ["and","pr.discount","is","null"];}
            else if (pricetype === "discount") {
              pricetype = ["and","pr.discount","is","not null"];}
            else {
              err = self.lang.invalid_refnumber;}
            refnumber = refnumber.split("~")[0];
            vsql = {select:["pr.id as id"], from:"price pr",
              inner_join:["product p","on",["pr.product_id","=","p.id"]],
              where:[["p.partnumber","=","?"],pricetype,["and","pr.curr","=","'"+curr+"'"],
                ["and","pr.validfrom","=","'"+validfrom+"'"],["and","pr.qty","=",qty]]};
            if (params.use_deleted === false) {
              vsql.where.push(["and","p.deleted","=","0"],["and","pr.deleted","=","0"]);}}
          else {
            err = self.lang.invalid_refnumber;}
          break;
        
        case "product":
          //partnumber
          vsql = {
            select:["p.id as id","p.description as description","p.unit as unit",
              "p.tax_id as tax_id","t.rate as rate"], 
            from:"product p", 
            left_join: ["tax t","on",["p.tax_id","=","t.id"]], 
            where:[["p.partnumber","=","?"]]};
          if (params.use_deleted === false) {
            vsql.where.push(["and","p.deleted","=","0"]);}
          break;
        
        case "place":
          //planumber
          vsql = {select:["p.id as id","pt.groupvalue as placetype"], 
            from:"place p", 
            inner_join:["groups pt","on",["p.placetype","=","pt.id"]], 
            where:[["p.planumber","=","?"]]}
          if (params.use_deleted === false) {
            vsql.where.push(["and","p.deleted","=","0"]);}
          break;
          
        case "tax":
          //taxcode
          vsql = {select:["id","rate"], from:"tax", 
            where:[["taxcode","=","?"]]};
          break;
          
        case "trans":
          //transnumber
          if (params.use_deleted === false) {
            vsql = {
              select:["t.id as id","ttype.groupvalue as transtype","dir.groupvalue as direction",
                "cu.digit as digit"], 
              from:"trans t",
              inner_join:[["groups ttype","on",["t.transtype","=","ttype.id"]],
                ["groups dir","on",["t.direction","=","dir.id"]]],
              left_join:["currency cu","on",["t.curr","=","cu.curr"]],
              where:["t.transnumber","=","?"]};}
          else {
            vsql = {
              select:["t.id as id","ttype.groupvalue as transtype","dir.groupvalue as direction",
                "cu.digit as digit"], 
              from:"trans t",
              inner_join:[["groups ttype","on",["t.transtype","=","ttype.id"]],
                ["groups dir","on",["t.direction","=","dir.id"]]],
              left_join:["currency cu","on",["t.curr","=","cu.curr"]], 
              where:[["t.transnumber","=","?"],
                ["and", [["t.deleted","=","0"], 
                  ["or",[["ttype.groupvalue","=","'invoice'"],["and","dir.groupvalue","=","'out'"]]],
                  ["or",[["ttype.groupvalue","=","'receipt'"],["and","dir.groupvalue","=","'out'"]]], 
                  ["or",["ttype.groupvalue","=","'cash'"]]]]]};}
          break;
          
        case "setting":
          //fieldname
          info.ref_nervatype = "setting";
          vsql = {select:["id","'setting' as ref_nervatype"], from:"fieldvalue",
            where:[["ref_id","is","null"],["and","fieldname","=","?"]]};
          if (params.use_deleted === false) {
            vsql.where.push(["and","deleted","=","0"]);}
          break;
          
        case "link":
          //nervatype_1~refnumber_1~~nervatype_2~refnumber_2
          if (refnumber.split("~~").length > 1) {
            params.ref_type_1 = refnumber.split("~~")[0].split("~")[0];
            params.ref_value_1 = refnumber.split("~~")[0].replace(params.ref_type_1+"~","");
            var vsql_1 = set_vsql(params.ref_type_1, params.ref_value_1);
            params.ref_type_2 = refnumber.split("~~")[1].split("~")[0];
            params.ref_value_2 = refnumber.split("~~")[1].replace(params.ref_type_2+"~","");
            return {err:vsql_1.err, vsql:vsql_1.vsql, refnumber:vsql_1.refnumber, 
              ref_index:vsql_1.ref_index};}
          else {
            err = self.lang.invalid_refnumber;}
          break;
          
        case "rate":
          //ratetype~ratedate~curr~planumber
          if (refnumber.split("~").length >= 3) {
            vsql = {select:["r.id as id"], from: "rate r",
              inner_join:["groups rt","on",[["r.ratetype","=","rt.id"],
                ["and","rt.groupvalue","=","?"]]],
              left_join:["place p","on",["r.place_id","=","p.id"]], 
              where:[["r.ratedate","=","'"+refnumber.split("~")[1]+"'"],
                ["and","r.curr","=","'"+refnumber.split("~")[2]+"'"]]}
            if (refnumber.split("~").length>3) {
              vsql.where.push(["and","p.planumber","=","'"+refnumber.split("~")[3]+"'"]);}
            else {
              vsql.where.push(["and","r.place_id","is","null"]);}
            if (params.use_deleted === false) {
              vsql.where.push(["and","r.deleted","=","0"]);}
            refnumber = refnumber.split("~")[0];}
          else {
            err = self.lang.invalid_refnumber;}
          break;
          
        case "log":
          //empnumber~crdate
          if (refnumber.split("~").length > 1) {
            vsql = {select:["l.id as id"], from:"log l",
              inner_join:["employee e","on",["l.employee_id","=","e.id"]],
              where:[["e.empnumber","=","?"],["and","l.crdate","=","'"+refnumber.split("~")[1]+"'"]]};
            refnumber = refnumber.split("~")[0];}
          else {
            err = self.lang.invalid_refnumber;}
          break;
          
        case "ui_audit":
          //usergroup~nervatype~transtype
          if (refnumber.split("~").length > 1) {
            vsql = {select:["au.id as id"], from:"ui_audit au",
              inner_join:[["groups ug","on",[["au.usergroup","=","ug.id"],
                ["and","ug.groupvalue","=","?"]]], 
              ["groups nt","on",["au.nervatype","=","nt.id"]]],
              where:[]};
            if (refnumber.split("~").length > 2) {
              if (refnumber.split("~")[1]==="trans"){
                vsql.inner_join.push(["groups st","on",[["au.subtype","=","st.id"],
                  ["and","st.groupvalue","=","'"+refnumber.split("~")[2]+"'"]]]);}
              else if (refnumber.split("~")[1]==="report"){
                vsql.inner_join.push(["ui_report rp","on",[["au.subtype","=","rp.id"],
                  ["and","rp.reportkey","=","'"+refnumber.split("~")[2]+"'"]]]);}
              else {err = self.lang.invalid_refnumber;}}
            vsql.where.push(["and","nt.groupvalue","=","'"+refnumber.split("~")[1]+"'"]);
            if (refnumber.split("~").length <= 2) {
              vsql.where.push(["and","subtype","is","null"]);}
            refnumber = refnumber.split("~")[0];}
          else {
            err = self.lang.invalid_refnumber;}
          break;
          
        case "ui_menufields":
          //menukey~fieldname
          if (refnumber.split("~").length > 1) {
            vsql = {select:["mf.id as id"], from:"ui_menufields mf",
              inner_join:["ui_menu m","on",[["mf.menu_id","=","m.id"],["and","m.menukey","=","?"]]],
              where:["mf.fieldname","=","'"+refnumber.split("~")[1]+"'"]}
            refnumber = refnumber.split("~")[0];}
          else {
            err = self.lang.invalid_refnumber;}
          break;
          
        case "ui_reportfields":
          //reportkey~fieldname
          if (refnumber.split("~").length > 1) {
            vsql = {select:["rf.id as id"], from:"ui_reportfields rf",
              inner_join:["ui_report r","on",[["rf.report_id","=","r.id"],["and","r.reportkey","=","?"]]],
              where: ["rf.fieldname","=","'"+refnumber.split("~")[1]+"'"]}
            refnumber = refnumber.split("~")[0];}
          else {
            err = self.lang.invalid_refnumber;}
          break;
          
        case "ui_reportsources":
          //reportkey~dataset
          if (refnumber.split("~").length > 1) {
            vsql = {select:["rs.id as id"], from:"ui_reportsources rs",
              inner_join:["ui_report r","on",[["rs.report_id","=","r.id"],["and","r.reportkey","=","?"]]],
              where:["rs.dataset","=","'"+refnumber.split("~")[1]+"'"]}
            refnumber = refnumber.split("~")[0];}
          else {
            err = self.lang.invalid_refnumber;}
          break;
              
        default:
          err = self.lang.invalid_refnumber;
          break;}}
    return {err:err, vsql:models.getSql(self.engine,vsql), refnumber:refnumber, ref_index:ref_index}}
  
  var info = {nervatype:params.nervatype, refnumber:params.refnumber};
  var conn = params.conn;
  if (!conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect, null, info);}}
  async.waterfall([
    function(callback) {
      if(params.nervatype === null || params.refnumber === null) {
        callback(self.lang.invalid_refnumber);}
      else {
        callback(null);}},
    
    function(callback) {
      var rvsql = set_vsql(params.nervatype, params.refnumber);
      callback(rvsql.err, rvsql.vsql, rvsql.refnumber, rvsql.ref_index);},
    
    function(vsql, refnumber, ref_index, callback) {
      if (params.nervatype === "fieldvalue"){
        conn.query(vsql, [refnumber], function (error, data) {
          if (error) {callback(error, null);}
          else {
            if (data.rowCount === 0) {
              callback(self.lang.invalid_refnumber.replace("refnumber",params.nervatype)+params.refnumber, null);}
            else {
              info.ref_nervatype = data.rows[0].ref_nervatype;
              var refnumber = params.refnumber.replace("~~"+refnumber,"").split("~")[0];
              var rsql = set_vsql(data.rows[0].ref_nervatype, refnumber);
              callback(rsql.err, rsql.vsql, rsql.refnumber, rsql.ref_index);}}});}
      else if (params.nervatype === "link"){
        conn.query(vsql, [refnumber], function (error, data) {
          if (error) {callback(error, null);}
          else {
            if (data.rowCount === 0 || ref_index > data.rowCount-1) {
              callback(self.lang.invalid_refnumber.replace("refnumber",params.nervatype)+params.refnumber, null);}
            else {
              params.ref_id_1 = data.rows[ref_index].id;
              var vsql_2 = set_vsql(params.ref_type_2, params.ref_value_2);
              callback(null, vsql_2.vsql, vsql_2.refnumber, vsql_2.ref_index);}}});}
      else {
        callback(null, vsql, refnumber, ref_index);}},
    
    function(vsql, refnumber, ref_index, callback) {
      conn.query(vsql, [refnumber], function (error, data) {
        if (error) {callback(error, null);}
        else {
          if (data.rowCount === 0 || ref_index > data.rowCount-1) {
            callback(self.lang.invalid_refnumber.replace("refnumber",params.nervatype)+params.refnumber, null);}
          else {
            var id = data.rows[ref_index].id;
            switch (params.nervatype) {
              case "customer":
                if (params.extra_info && data.rowCount < 2){
                  callback(self.lang.invalid_refnumber.replace("refnumber",params.nervatype)+params.refnumber, null);}
                else {
                  if (params.extra_info) {
                    data.rows.forEach(function(row) {
                      if (row.custtype === "own"){
                        info.compname = row.custname;
                        info.comptax = row.taxnumber;
                        info.compaddress = ((row.zipcode||"")+" "+(row.city||"")+" "+(row.street||"")).trim();}
                      else {
                        id = row.id;
                        info.custtype = row.custtype;
                        info.terms = row.terms;
                        info.custname = row.custname;
                        info.custtax = row.taxnumber;
                        info.custaddress = ((row.zipcode||"")+" "+(row.city||"")+" "+(row.street||"")).trim();}});}
                  else {
                    info.custtype = data.rows[ref_index].custtype;}}
                break;
              case "product":
                info.description = data.rows[ref_index].description;
                info.unit = data.rows[ref_index].unit;
                info.tax_id = data.rows[ref_index].tax_id;
                info.rate = data.rows[ref_index].rate;
                break;
              case "place":
                info.placetype = data.rows[ref_index].placetype;
                break;
              case "event":
                info.ref_nervatype = data.rows[ref_index].ref_nervatype;
                break;
              case "tax":
                info.rate = data.rows[ref_index].rate;
                break;
              case "item":
                info.transtype = data.rows[ref_index].transtype;
                info.direction = data.rows[ref_index].direction;
                info.digit = data.rows[ref_index].digit;
                info.qty = data.rows[ref_index].qty;
                info.discount = data.rows[ref_index].discount;
                info.tax_id = data.rows[ref_index].tax_id;
                info.rate = data.rows[ref_index].rate;
                break;
              case "movement":
                info.movetype = data.rows[ref_index].movetype;
                info.transtype = data.rows[ref_index].transtype;
                info.direction = data.rows[ref_index].direction;
                break;
              case "trans":
              case "payment":
                info.transtype = data.rows[ref_index].transtype;
                info.direction = data.rows[ref_index].direction;
                break;}
            callback(null, id);}}});},
    
    function(ref_id, callback) {
      if (params.nervatype === "fieldvalue"){
        var fieldname = params.refnumber.split("~~")[params.refnumber.split("~~").length-1];
        var ref_index = 0; var err = null; var vsql;
        if (fieldname.split("~").length > 1) {
          ref_index = parseInt(fieldname.split("~")[1],10);
          if (isNaN(ref_index) || ref_index < 1){
            err = self.lang.invalid_refnumber;}
          else {ref_index -= 1;}
        fieldname = fieldname.split("~")[0];}
        if (err !== null){callback(err, null);}
        else {
          vsql = {select:["id"], from:"fieldvalue", 
            where:[["fieldname","=","?"],["and","ref_id","=","?"]]}
          if (params.use_deleted === false) {
            vsql.where.push(["and","deleted","=","0"]);}
          vsql = models.getSql(self.engine,vsql);
          conn.query(vsql, [fieldname, ref_id], function (error, data) {
            if (error) {callback(error, null);}
            else {
              if (data.rowCount === 0 || ref_index > data.rowCount-1) {
                callback(self.lang.invalid_refnumber.replace("refnumber",params.nervatype)+params.refnumber, null);}
              else {callback(null, data.rows[ref_index].id);}}});}}
      else if (params.nervatype === "link"){
        vsql = {select:["l.id as id"], from:"link l",
          inner_join:[
            ["groups nt1","on",["l.nervatype_1","=","nt1.id"],
              ["and", "nt1.groupname","=","'nervatype'"],["and","nt1.groupvalue","=","?"]],
            ["groups nt2","on",["l.nervatype_2","=","nt2.id"],["and","nt2.groupname","=","'nervatype'"],
              ["and","nt2.groupvalue","=","?"]]],
          where:[["l.ref_id_1","=","?"],["and","l.ref_id_2","=","?"]]}
        if (params.use_deleted === false) {
           vsql.where.push(["and","l.deleted","=","0"]);}
        vsql = models.getSql(self.engine,vsql);
        conn.query(vsql, [params.ref_type_1, params.ref_type_2, params.ref_id_1, ref_id], 
          function (error, data) {
            if (error) {callback(error, null);}
            else {
              if (data.rowCount === 0) {
                callback(self.lang.invalid_refnumber.replace("refnumber",params.nervatype)+params.refnumber, null);}
              else {callback(null, data.rows[0].id);}}});}
      else {
        callback(null, ref_id);}}
  ],
  function(err, ref_id) {
    if(err){if(err.message){err = err.message;}}
    if (!params.conn && conn){
      conn.close();}
    _callback(err, ref_id, info);});}

function get_groups_id(self, params, _callback){
  var _sql; 
  _sql = {select:["*"], from:"groups", 
    where:[["deleted","=","0"]]}
  if (typeof params.groupname !== "undefined") {
    if (Array.isArray(params.groupname)) {
      var groupname_lst ="";
      params.groupname.forEach(function(groupname) {
        groupname_lst += ",'"+groupname+"'";});
      _sql.where.push(["and","groupname","in",[[],groupname_lst.substr(1)]]);}
    else {
      _sql.where.push(["and","groupname","=","'"+params.groupname+"'"]);}}
  var conn = params.conn;
  if (!conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  conn.query(models.getSql(self.engine,_sql), [], function (err, data) {
    if (!params.conn && conn){
      conn.close();}
    var groups = {};
    if(err){if(err.message){err = err.message;}
      _callback(err, groups);}
    else {
      groups.all = data.rows;
      data.rows.forEach(function(row) {
        if (typeof groups[row.groupname] === "undefined"){
          groups[row.groupname] = {};}
        groups[row.groupname][row.groupvalue] = row.id;});
      _callback(null, groups);}});}

function decode_token(self, token){
  return jwt.decode(token, {complete: true}); }

function verify_token(self, token, _callback){
  jwt.verify(token, self.token_settings.key, self.token_settings.options, function(err, decoded) {
    _callback(err, decoded); });}

//connect

function get_data_audit(self, params, _callback){
  //Nervatura data access rights: own,usergroup,all (transfilter)
  //see: employee.usergroup+link+transfilter
  if (self.employee === null){
    _callback(self.lang.ndi_invalid_login);}
  else if (self.employee.usergroup === null){
    _callback(self.lang.missing_usergroup);}
  else {
    var conn;
    if(params.conn){
      conn = params.conn;}
    else {
      conn = get_connect(self);
      if (!conn){return _callback(self.lang.not_connect);}}
    var _sql = {select:["tt.groupvalue as transfilter"],
      from:"employee e",
      inner_join:[
        ["link l","on",[["l.ref_id_1","=","e.usergroup"],["and","l.deleted","=","0"]]],
        ["groups nt1","on",[["l.nervatype_1","=","nt1.id"],["and","nt1.groupname","=","'nervatype'"],
        ["and","nt1.groupvalue","=","'groups'"]]],
        ["groups nt2","on",[["l.nervatype_2","=","nt2.id"],["and","nt2.groupname","=","'nervatype'"],
        ["and","nt2.groupvalue","=","'groups'"]]],
        ["groups tt","on",["l.ref_id_2","=","tt.id"]]],
      where:["e.id","=",self.employee.id]}
    conn.query(models.getSql(self.engine, _sql), [], function (err, data) {
      if(!params.conn){
        conn.close();}
      if(err){if(err.message){err = err.message;}
        _callback(err);}
      else {
        if (data.rowCount === 0) {
          _callback(null, "all");}
        else {_callback(null, data.rows[0].transfilter);}}});}}

function get_object_audit(self, params, _callback){
  //Nervatura objects access rights: disabled,readonly,update,all (inputfilter)
  //see: audit
  if (self.employee === null){
    _callback(self.lang.ndi_invalid_login);}
  else if (self.employee.usergroup === null){
    _callback(self.lang.missing_usergroup);}
  else if (!params.hasOwnProperty("nervatype") && !params.hasOwnProperty("transtype")){
    _callback(self.lang.missing_nervatype);}
  else if (params.nervatype === "sql" || params.nervatype === "fieldvalue"){
    _callback(null, "all");}
  else {
    var _sql;
    _sql = {
      select:["inf.groupvalue as inputfilter"],
      from:"ui_audit a",
      inner_join: [
        ["groups inf","on",["a.inputfilter","=","inf.id"]],
        ["groups nt","on",["a.nervatype","=","nt.id"]]],
      left_join: ["groups st","on",["a.subtype","=","st.id"]],
      where: [["a.usergroup","=",self.employee.usergroup]]}
    
    var typevalue = "";
    if (params.transtype_id) {
      if (Array.isArray(params.transtype_id)) {
        var subtype_lst ="";
        params.transtype_id.forEach(function(subtype) {
          subtype_lst += ","+subtype;});
        _sql.where.push(["and","a.subtype","in",[[],subtype_lst.substr(1)]]);}
      else {
        _sql.where.push(["and","a.subtype","=",params.transtype_id]);}}
    if (params.transtype) {
      if (Array.isArray(params.transtype)) {
        var transtype_lst ="";
        params.transtype.forEach(function(transtype) {
          transtype_lst += ",'"+transtype+"'";});
        typevalue = transtype_lst.substr(1);
        _sql.where.push(["and","st.groupvalue","in",[[],transtype_lst.substr(1)]]);}
      else {
        typevalue = params.transtype;
        _sql.where.push(["and","st.groupvalue","=","'"+params.transtype+"'"]);}}
    
    if (params.nervatype_id) {
      if (Array.isArray(params.nervatype_id)) {
        var nervatype_lst ="";
        params.nervatype_id.forEach(function(nervatype) {
          nervatype_lst += ","+nervatype;});
        _sql.where.push(["and","a.nervatype","in",[[],nervatype_lst.substr(1)]]);}
      else {
        _sql.where.push(["and","a.nervatype","=",params.nervatype_id]);}}
    if (params.nervatype) {
      if (Array.isArray(params.nervatype)) {
        var groupname_lst ="";
        params.nervatype.forEach(function(groupname) {
          groupname_lst += ",'"+groupname+"'";});
        typevalue = groupname_lst.substr(1);
        _sql.where.push(["and","nt.groupvalue","in",[[],groupname_lst.substr(1)]]);}
      else {
        typevalue = params.nervatype;
        _sql.where.push(["and","nt.groupvalue","=","'"+params.nervatype+"'"]);}}
    
    var conn;
    if(params.conn){
      conn = params.conn;}
    else {
      conn = get_connect(self);
      if (!conn){return _callback(self.lang.not_connect);}}
    conn.query(models.getSql(self.engine, _sql), [], function (err, data) {
      if(!params.conn){conn.close();}
      if(err){if(err.message){err = err.message;}
        _callback(err);}
      else {
        if (data.rowCount === 0) {
          _callback(null, "all", typevalue);}
        else {
          var istate = "all";
          data.rows.forEach(function(row) {
            switch (row.inputfilter) {
              case "disabled":
                istate = row.inputfilter;
                break;
              case "readonly":
                if (istate !== "disabled") {
                  istate = row.inputfilter;}
                break;
              case "update":
                if (istate === "all") {
                  istate = row.inputfilter;}
                break;
              default:
                break;}});
          _callback(null, istate, typevalue);}}});}}

function get_database_settings(self, params, _callback){
  var conn = params.conn;
    if(!conn){
      conn = get_connect(self);
      if (!conn){return _callback(self.lang.not_connect);}}
  var _sql = [ 
    {select:["'setting' as stype","fieldname","value","notes as data","id as info"],
      from:"fieldvalue", where:[["deleted","=","0"],["and","ref_id","is","null"]]},
    {union_select:["'pattern' as stype","p.description as fieldname","p.notes as value", 
        "tt.groupvalue as data","p.defpattern as info"],
      from:"pattern p", inner_join:["groups tt","on",["p.transtype","=","tt.id"]],
      where:["p.deleted","=","0"]}]
  conn.query(models.getSql(self.engine, _sql), [], function (err, data) {
    if(!params.conn){conn.close();}
    if(err){if(err.message){err = err.message;}
      _callback(err);}
    else {
      var settings = {setting:{},pattern:{}};
      data.rows.forEach(function(row) {
        if (row.stype === "setting") {
          settings.setting[row.fieldname] = {value:row.value, data:row.data}}
        else {
          if (!settings.pattern.hasOwnProperty(row.data)) {
            settings.pattern[row.data] = [];}
          if (row.info === 1){
            settings.pattern[row.data].unshift({description:row.fieldname, notes:row.value, defpattern:true});}
          else {
            settings.pattern[row.data].push({description:row.fieldname, notes:row.value, defpattern:false});}}});
      _callback(null, settings);}});}
        
function delete_data(self, params, _callback){
  if (typeof params.nervatype === "undefined") {params.nervatype = null;}
  if (typeof params.ref_id === "undefined") {params.ref_id = null;}
  if (typeof params.refnumber === "undefined") {params.refnumber = null;}
  if (typeof params.log_enabled === "undefined") {params.log_enabled = true;}
  if (typeof params.validator === "undefined") {params.validator = {};}
  
  var conn = params.validator.conn; var trans = params.transaction;
  if (!conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  if (!trans){
    trans = connect.beginTransaction({connection:conn, engine:self.engine});} 
  async.waterfall([
    function(callback) {
      //params.refnumber vs params.ref_id
      if (params.nervatype === null) {
        callback(self.lang.missing_nervatype);}
      if (params.ref_id !== null && params.ref_id > 0) {
        callback(null, params.ref_id);}
      else if (params.refnumber !== null) {
        get_id_from_refnumber(self, {nervatype:params.nervatype, refnumber:params.refnumber}, 
          function(err, id, info){
            if (err) {
              callback(err);}
            else if(id === null){
              callback(self.lang.missing_fieldname+": "+params.refnumber);}
            else {callback(null, id);}})}
      else {
        callback(self.lang.missing_fieldname+": ref_id or refnumber");}},
    
    function(ref_id, callback) {
      //check integrity
      if (["address","barcode","contact","event","fieldvalue","item","link","log","movement","pattern",
        "payment", "price","rate"].indexOf(params.nervatype) > -1 ){
        callback(null, ref_id);}
      else if(["numberdef"].indexOf(params.nervatype) > -1) {
        //protected, always false
        callback(self.lang.integrity_error, ref_id);}
      else {
        var _sql;
        switch (params.nervatype) {
          case "currency":
            //(link), place,price,rate,trans
            _sql = {
              select:["sum(co) as sco"], from:[
                [[{select:["count(place.id) as co"], from:"place",
                inner_join:["currency","on",["place.curr","=","currency.curr"]],
                where:[["place.deleted","=","0"],["and","currency.id","=",ref_id]]},
                {union_select:["count(price.id) as co"], from:"price", 
                inner_join:["currency","on",["price.curr","=","currency.curr"]], 
                where:[["price.deleted","=","0"],["and","currency.id","=",ref_id]]}, 
                {union_select:["count(rate.id) as co"], from:"rate", 
                inner_join:["currency","on",["rate.curr","=","currency.curr"]],
                where:[["rate.deleted","=","0"],["and","currency.id","=",ref_id]]},
                {union_select:["count(trans.id) as co"], from:"trans",
                inner_join:["currency","on",["trans.curr","=","currency.curr"]],
                where:[["trans.deleted","=","0"],["and","currency.id","=",ref_id]]}]],"foo"]}
            break;
          case "customer": 
            //(address,contact), event,project,trans,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"trans", where:["customer_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"project", where:["customer_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"event",
                inner_join:[["groups nt","on",["event.nervatype","=","nt.id"]],
                  ["and","nt.groupvalue","=","'customer'"]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[
                  {select:["id"], from:"groups", 
                    where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'customer'"]]}]],
                    ["and","ref_id_2","=",ref_id]]}
              ]],"foo"]}
            break;
          case "deffield": 
            //fieldvalue
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(fieldvalue.id) as co"], from:"fieldvalue",
                inner_join:["deffield","on",["deffield.fieldname","=","fieldvalue.fieldname"]],
                where:[["fieldvalue.deleted","=","0"],["and","deffield.id","=",ref_id]]}
              ]],"foo"]}
            break;
          case "employee": 
            //(address,contact), event,trans,log,link,ui_printqueue,ui_userconfig
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"trans", where:["employee_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"trans", where:["cruser_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"log", where:["employee_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"ui_printqueue", where:["employee_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"ui_userconfig", where:["employee_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"event",
                inner_join:[["groups nt","on",["event.nervatype","=","nt.id"]],["and","nt.groupvalue","=","'employee'"]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link", 
                where:[["nervatype_2","=",[{select:["id"], from:"groups", 
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'employee'"]]}]],
                    ["and","ref_id_2","=",ref_id]]}
              ]],"foo"]}
            break;
          case "groups": 
            //barcode,deffield,employee,event,rate,tool,trans,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"groups",
                where:[["groupname","in",[[],"'nervatype'","'custtype'","'fieldtype'","'logstate'","'movetype'",
                  "'transtype'","'placetype'","'calcmode'","'protype'","'ratetype'","'direction'",
                  "'transtate'","'inputfilter'","'filetype'","'wheretype'","'aggretype'"]],["and","id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[{select:["id"], from:"groups", 
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'groups'"]]}]],
                  ["and","ref_id_2","=",ref_id]]},
                {union_select:["count(*)"], from:"barcode", where:["barcode.barcodetype","=",ref_id]},
                {union_select:["count(*)"], from:"deffield", where:[["deffield.deleted","=","0"],
                  ["and","deffield.subtype","=",ref_id]]},
                {union_select:["count(*)"], from:"employee", where:[["employee.deleted","=","0"],
                  ["and","employee.usergroup","=",ref_id]]},
                {union_select:["count(*)"], from:"employee", where:[["employee.deleted","=","0"],
                  ["and","employee.department","=",ref_id]]},
                {union_select:["count(*)"], from:"event", where:[["event.deleted","=","0"],
                  ["and","event.eventgroup","=",ref_id]]},
                {union_select:["count(*)"], from:"rate", where:[["rate.deleted","=","0"],
                  ["and","rate.rategroup","=",ref_id]]},
                {union_select:["count(*)"], from:"tool", where:[["tool.deleted","=","0"],
                  ["and","tool.toolgroup","=",ref_id]]},
                {union_select:["count(*)"], from:"trans", where:[["trans.deleted","=","0"],
                  ["and","trans.department","=",ref_id]]}
              ]],"foo"]}
            break;
          case "place": 
            //(address,contact), event,movement,place,rate,trans,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"event",
                inner_join:[["groups nt","on",["event.nervatype","=","nt.id"]],
                  ["and","nt.groupvalue","=","'place'"]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[{select:["id"], from:"groups", 
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'place'"]]}]],
                    ["and","ref_id_2","=",ref_id]]},
                {union_select:["count(*) as co"], from:"movement", 
                  where:[["movement.deleted","=","0"],["and","movement.place_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"place", 
                  where:[["place.deleted","=","0"],["and","place.place_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"rate", 
                  where:[["rate.deleted","=","0"],["and","rate.place_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"trans", 
                  where:[["trans.deleted","=","0"],["and","trans.place_id","=",ref_id]]}
              ]],"foo"]}
            break;
          case "product": 
            //address,barcode,contact,event,item,movement,price,tool,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"event",
                inner_join:["groups nt","on",[["event.nervatype","=","nt.id"],
                  ["and","nt.groupvalue","=","'product'"]]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"address",
                inner_join:["groups nt","on",[["address.nervatype","=","nt.id"],
                  ["and","nt.groupvalue","=","'product'"]]],
                where:[["address.deleted","=","0"],["and","address.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"contact",
                inner_join:["groups nt","on",[["contact.nervatype","=","nt.id"],
                  ["and","nt.groupvalue","=","'product'"]]],
                where:[["contact.deleted","=","0"],["and","contact.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[{select:["id"], from:"groups", 
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'product'"]]}]],
                    ["and","ref_id_2","=",ref_id]]},
                {union_select:["count(*) as co"], from:"barcode", where:["barcode.product_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"movement", 
                  where:[["movement.deleted","=","0"],["and","movement.product_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"item", 
                  where:[["item.deleted","=","0"],["and","item.product_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"price", 
                  where:[["price.deleted","=","0"],["and","price.product_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"tool", 
                  where:[["tool.deleted","=","0"],["and","tool.product_id","=",ref_id]]}
              ]],"foo"]}
            break;
          case "project": 
            //(address,contact), event,trans,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"trans", where:["project_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"event",
                inner_join:["groups nt","on",[["event.nervatype","=","nt.id"],
                  ["and","nt.groupvalue","=","'project'"]]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[{select:["id"], from:"groups",
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'project'"]]}]],
                    ["and","ref_id_2","=",ref_id]]}
              ]],"foo"]}
            break;
          case "tax": 
            //item,product
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"item",
                where:[["item.deleted","=","0"],["and","item.tax_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"product",
                where:[["product.deleted","=","0"],["and","product.tax_id","=",ref_id]]}
              ]],"foo"]}
            break;
          case "tool": 
            //(address,contact), event,movement,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"movement", where:["tool_id","=",ref_id]},
                {union_select:["count(*) as co"], from:"event",
                inner_join:["groups nt","on",[["event.nervatype","=","nt.id"],
                  ["and","nt.groupvalue","=","'tool'"]]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[{select:["id"], from:"groups",
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'tool'"]]}]],
                  ["and","ref_id_2","=",ref_id]]}
              ]],"foo"]}
            break;
          case "trans": 
            //(address,contact), event,link
            _sql = {
              select:["sum(co) as sco"],
              from: [[[
                {select:["count(*) as co"], from:"event",
                inner_join:["groups nt","on",[["event.nervatype","=","nt.id"],
                  ["and","nt.groupvalue","=","'trans'"]]],
                where:[["event.deleted","=","0"],["and","event.ref_id","=",ref_id]]},
                {union_select:["count(*) as co"], from:"link",
                where:[["nervatype_2","=",[{select:["id"], from:"groups",
                  where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'trans'"]]}]],
                  ["and","ref_id_2","=",ref_id]]}
              ]],"foo"]}
            break;}
        if(_sql){
          trans.query(models.getSql(self.engine, _sql), [], function (error, data) {
            if (error) {callback(error, ref_id);}
            else {
              if (data.rowCount === 0) {
                callback(null, ref_id);}
              else if (parseInt(data.rows[0].sco,10) > 0) {
                callback(self.lang.integrity_error, ref_id);}
              else {callback(null, ref_id);}}});}
        else {
          if (params.nervatype.substr(0,3) === "ui_" ){
            callback(null, ref_id);}
          else {callback(self.lang.integrity_error, ref_id);}}}},
    
    function(ref_id, callback) {
      //check logical delete
      if (typeof ntura.model[params.nervatype].deleted === "undefined") {
        callback(null, ref_id, false);}
      else {
        var _sql = {
          select:["value"], from:"fieldvalue",
          where:[["ref_id","is","null"],["and","fieldname","=","'not_logical_delete'"]]}
        trans.query(models.getSql(self.engine, _sql), [], function (error, data) {
          if (error) {callback(error, ref_id);}
          else {
            if (data.rowCount === 0) {
              callback(null, ref_id, true);}
            else if (data.rows[0].value === "true") {
              callback(null, ref_id, false);}
            else {callback(null, ref_id, true);}}});}},
    
    function(ref_id, logical_delete, callback) {
      var _sql;
      if (logical_delete) {
        _sql = {update:params.nervatype, set:[[],["deleted","=","1"]], 
          where:["id","=",ref_id]}}
      else {
        _sql = {delete:"", from:params.nervatype, where:["id","=",ref_id]}}
      trans.query(models.getSql(self.engine, _sql), [], function (error, data) {
        if (error) {callback(error, ref_id);}
        else {callback(null, ref_id, logical_delete);}});},
    
    function(ref_id, logical_delete, callback) {
      if (logical_delete === false) {
        //delete all fieldvalue records
        var sel_sql = {
          select:["fv.id"], from:"deffield df", 
          inner_join:[
            ["groups nt","on",[["df.nervatype","=","nt.id"],
              ["and","nt.groupvalue","=","'"+params.nervatype+"'"]]],
            ["fieldvalue fv","on",[["df.fieldname","=","fv.fieldname"],
              ["and","fv.deleted","=","0"],["and","fv.ref_id","=",ref_id]]]]}
        trans.query(models.getSql(self.engine, sel_sql), [], function (error, data) {
          if(data.rowCount>0){
            var del_sql = {delete:"", from:"fieldvalue", 
              where:["id","in",[sel_sql]]}
            trans.query(models.getSql(self.engine, del_sql), [], function (error, data) {
              if (error) {callback(error, ref_id);}
              else {callback(null, ref_id);}});}
          else{
            callback(null, ref_id);}});}
      else {callback(null, ref_id);}},
    
    function(ref_id, callback) {
      if (params.log_enabled === true && self.employee){
        //insert log
        insert_log(self,{conn:trans, logstate:"delete", nervatype:params.nervatype, ref_id:ref_id}, 
          function(err, result){
            callback(err, ref_id);});}
      else {callback(null, ref_id);}}
  ],
  function(err, ref_id) {
    if(err){if(err.message){err = err.message;}}  
    if (err && !params.transaction){
      if(trans){
        if(trans.rollback){
          trans.rollback(function (error) {
            if (!error && !params.validator.conn){
              conn.close();}
            _callback(err, ref_id);});}
        else{
          if (!params.validator.conn){
            conn.close();}
          _callback(err, ref_id);}}
      else{
        _callback(err, ref_id);}}
    else if (!err && !params.transaction){
      if(trans){
        if(trans.commit){
          trans.commit(function (err) {
            if (!err && !params.validator.conn){
              conn.close();}
            _callback(err, ref_id);});}
        else{
          if (!params.validator.conn){
            conn.close();}
          _callback(err, ref_id);}}
      else{
        _callback(err, ref_id);}}
    else {
      _callback(err, ref_id);}});}

function get_connect(self) {
  try {
    return connect.createConnection(
      {connect:self.dbs_config, pool:true, config:self.app_config.pool_config});}
  catch (error) {
    return null;}}

function insert_log(self, params, _callback) {
  var conn = params.conn;
  if (!conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  async.waterfall([
    function(callback) {
      if (self.employee && self.employee !== null && params.nervatype !== "log" && (typeof params.logstate !== "undefined")) {
        if ((typeof params.nervatype === "undefined") || (params.nervatype === null)) {
          params.nervatype = "";}
        var _sql = [{
          select:["id","groupname as fieldname","groupvalue as value"],
          from:"groups",
          where:[["groupname","=","'logstate'"],["and","groupvalue","=","'"+params.logstate+"'"]]},
          {union_select:["id","fieldname","value"], 
           from:"fieldvalue",
           where:[["fieldname","=","'"+"log_"+params.logstate+"'"],
             ["or","fieldname","=","'"+"log_"+params.nervatype+"_"+params.logstate+"'"]]},
          {union_select:["id","groupname as fieldname","groupvalue as value"],
           from:"groups",
           where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'"+params.nervatype+"'"]]}]
        conn.query(models.getSql(self.engine, _sql) , [], 
          function (error, data) {
          if (error) {callback(error,false);}
          else {callback(null,data);}});}
      else {
        return _callback(null,true);}},
    
    function(logdata, callback) {
      var log_enabled = false; var logstate_id = null; var nervatype_id = null;
      logdata.rows.forEach(function(row) {
        if (row.fieldname === "logstate") {logstate_id = row.id;}
        if (row.fieldname === "nervatype") {nervatype_id = row.id;}
        if (row.fieldname === "log_"+params.nervatype+"_"+params.logstate) {
          if(row.value === "true"){log_enabled = true;}}
        if ((row.fieldname === "log_"+params.logstate) &&
          ((logdata.rowCount === 2)&&(params.nervatype === "") || 
          (logdata.rowCount === 3)&&(params.nervatype !== ""))) {
          if(row.value === "true"){log_enabled = true;}}});
      if (log_enabled === true && logstate_id !== null){
        callback(null, logstate_id, nervatype_id);}
      else {
        return _callback(null,true);}},
    
    function(logstate_id, nervatype_id, callback) {
      var _sql = {
        insert_into:["log",[[],"logstate","employee_id","crdate"]], 
        values:[[],logstate_id,self.employee.id,"'"+out.getISODateTime(new Date(),true)+"'"]}
      if (nervatype_id !== null){
        _sql.insert_into[1].push("nervatype");
        _sql.values.push(nervatype_id);}
      if ((typeof params.ref_id !== "undefined")&&(params.ref_id !== null)){
        _sql.insert_into[1].push("ref_id");
        _sql.values.push(params.ref_id);}
      conn.query(models.getSql(self.engine, _sql),[],function (error, data) {
        if (error) {callback(error, false);} else {callback(null, true);}});}
  ],
  function(err, result) {
    if(err){if(err.message){err = err.message;}}
    if (!params.conn && conn){
      conn.close();}
    _callback(err, result);});}

function next_number(self, params, _callback) {
  //Get the next value from the numberdef table (transnumber, custnumber, partnumber etc.)
  if (typeof params.numberkey === "undefined") {params.numberkey = null;}
  if (typeof params.step === "undefined") {params.step = true;}
  if (typeof params.insert_key === "undefined") {params.insert_key = true;}
  
  var conn = params.conn || params.transaction;
  if (!conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  async.waterfall([
    function(callback) {
      if (params.numberkey === null){
        callback(self.lang.missing_required_field+": numberkey");}
      else {
        var _sql = {select:["*"], from:"numberdef", 
          where:["numberkey","=","?"]}
        conn.query(models.getSql(self.engine, _sql), [params.numberkey], 
          function (error, data) {
          if (error) {callback(error);}
          else {callback(null,data);}});}},
    
    function(data, callback) {
      if (data.rowCount === 0) {
        if (params.insert_key) {
          var ndrow = {prefix:params.numberkey.toUpperCase().substr(0,3), isyear:1, sep:"/", curvalue:0};
          var _sql = {
            insert_into:["numberdef",[[],"numberkey","prefix","curvalue","isyear","sep",
              "len","visible","readonly"]],
            values:[[],"'"+params.numberkey+"'","'"+ndrow.prefix+"'",0,1,"'/'",5,1,0]}
          conn.query(models.getSql(self.engine, _sql), [], 
            function (error, data) {
            if (error) {callback(error);}
            else {callback(null,ndrow);}});}
        else {
          callback(self.lang.invalid_value+": "+params.numberkey);}}
      else {callback(null, data.rows[0]);}},
    
    function(ndrow, callback) {
      if (ndrow.isyear === 1) {
        var _sql = {select:["value"], from:"fieldvalue",
          where:[["ref_id","is","null"],["and","fieldname","=","'transyear'"]]}
        conn.query(models.getSql(self.engine, _sql), [], function (error, data) {
          if (error) {callback(error);}
          else {
            var transyear = new Date().getFullYear();
            if (data.rowCount > 0) {
              if (isNaN(parseInt(data.rows[0].value,10)) === false){
                transyear = parseInt(data.rows[0].value,10);}}
            callback(null, ndrow, transyear);}});}
      else {
        callback(null, ndrow, null);}},
    
    function(ndrow, transyear, callback) {
      var retnumber = "";
      if (ndrow.prefix !== null) {
        retnumber += ndrow.prefix+ndrow.sep;}
      if (transyear !== null){
        retnumber += transyear.toString()+ndrow.sep;}
      retnumber += out.zeroPad(ndrow.curvalue+1,5);
      if (params.step === true || params.step === "true" || params.step === 1){
        var _sql = {update:"numberdef", 
          set:[[],["curvalue","=","curvalue+1"]], 
          where:["numberkey","=","'"+params.numberkey+"'"]}
        conn.query(models.getSql(self.engine, _sql), [], function (error, data) {
          if (error) {callback(error);}
          else {callback(null, retnumber);}});}
      else {callback(null, retnumber);}}
  ],
  function(err, retnumber) {
    if(err){if(err.message){err = err.message;}}
    if (!params.conn && !params.transaction){
      conn.close();}
    _callback(err, retnumber);})}

function get_login(self, params, _callback) {
  var validator = {valid:false, message:"", conn:null};
  async.waterfall([
    function(callback) {
      if (params.token) {
        verify_token(self, params.token, function(err, token){
          if(!err && token){
            validator.token = token;
            if(!params.username){
              params.username = token.username || token.email || token.phone_number; }
            if(!params.username){
              err = self.lang.missing_user; }}
          callback(err); })}
      else if (typeof params.username === "undefined") {
        callback(self.lang.missing_user);}
      else {callback(null);}},
    
    function(callback) {
      if (self.dbs_config===null){
        set_engine(self, {database:params.database}, function(err, result){
          if (err){
            callback(err);}
          else {callback(null);}});}
      else {callback(null);}},
    
    function(callback) {
      var conn = get_connect(self); 
      if (!conn){return _callback(self.lang.not_connect);}
      if (params.token) {
        var _sql ={select:["*"], from:"customer",
          where:[["inactive","=","0"],["and","deleted","=","0"],
            ["and","custnumber","=","?"]]}
        conn.query(models.getSql(self.engine, _sql), [params.username], function (error, results) {
          if(!error){
            if (results.rowCount>0) {
              validator.customer = results.rows[0];
              params.username = "admin";
              callback(null, conn); }
            else{
              conn.close();
              callback(self.lang.unknown_user); }}
          else {
            callback(error); }});}
      else {
        callback(null, conn); }
    },

    function(conn, callback) {
      var _sql ={select:["*"], from:"employee",
        where:[["inactive","=","0"],["and","deleted","=","0"],
          ["and","username","=","?"]]}
      conn.query(models.getSql(self.engine, _sql), [params.username], function (error, results) {
        if (error) {
          callback(self.lang.missing_database+params.database);}
        else {
          validator.conn = conn;
          if (results.rowCount>0) {
            validator.employee = results.rows[0]; }
          callback(null); }});},
      
      function(callback) {
        if(!validator.employee && !validator.token){
          callback(self.lang.unknown_user); }
        else {
          if(!validator.token){
            if (typeof params.password === "undefined" || params.password === "" || params.password === null) {
              params.password = null;}
            else {
              params.password = out.createHash(params.password,"hex");}
            if (params.password !== validator.employee.password) {
              callback(self.lang.wrong_password); }
            else {
              callback(null); }}
          else{
            callback(null); }}}
    ],
  function(err) {
    if(err){
      if(err.message){err = err.message;}}
    else{
      if(validator.employee){
        self.employee = validator.employee; }
      validator.valid = true; }
    _callback(err, validator);});};

function update_data(self, params, _callback){
  var values = params.values; 
  var nervatype = params.nervatype;
  if (typeof params.log_enabled === "undefined") {params.log_enabled = true;}
  if (typeof params.validate === "undefined") {params.validate = true;}
  if (typeof params.insert_field === "undefined") {params.insert_field = false;}
  if (typeof params.insert_row === "undefined") {params.insert_row = false;}
  if (typeof params.update_row === "undefined") {params.update_row = true;}
  if (typeof params.validator === "undefined") {params.validator = {};}
 
  var conn = params.validator.conn; var trans = params.transaction;
  if (!conn){
    conn = get_connect(self);
    if (!conn){return _callback(self.lang.not_connect);}}
  if (!trans){
    trans = connect.beginTransaction({connection:conn, engine:self.engine});}
  async.waterfall([
    function(callback) {
      if (typeof ntura.model[nervatype] === "undefined") {
        callback(self.lang.invalid_nervatype+" "+nervatype);}
      else {
        if ((typeof values.id !== "undefined") && (values.id !== null) && (values.id !== "")) {
          //id check
          var _sql = {select:["*"], from:nervatype, 
            where:["id","=","?"]}
          trans.query(models.getSql(self.engine, _sql), [values.id], function (error, data) {
            if (error) {callback(error);}
            else {
              if (data.rowCount>0) {
                if (params.update_row === false) {
                  //readonly record
                  callback(self.lang.disabled_update);}
                else {
                  callback(null, values.id);}}
              else {
                callback(null, null);}}});}
        else {
          callback(null, null);}}},
    
    function(ref_id, callback) {
      if ((ref_id === null) && (params.insert_row === false)){
        //disabled insert
        callback(self.lang.disabled_insert);}
      else {
        callback(null, ref_id);}},
    
    function(ref_id, callback) {
      //check fieldnames
      var check_values={values:{}, fvalues:{}, dvalues:{}, deffield:[], fieldvalue:[]};
      var fieldname_err = null;
      for (var fieldname in values) {
        switch (fieldname) {
          case "id":
          case "__tablename__":
            break;
          default:
            if (typeof ntura.model[nervatype][fieldname] !== "undefined") {
              check_values.values[fieldname] = values[fieldname];}
            else {
              if (valid_fieldvalue(nervatype)){
                check_values.fvalues[fieldname] = values[fieldname];}
              else {
                fieldname_err = self.lang.unknown_fieldname+": "+fieldname;}}
            break;}}
      if (fieldname_err !== null) {callback(fieldname_err);}
      else {callback(null, ref_id, check_values)};
    },
    
    function(ref_id, check_values, callback) {
      if (valid_fieldvalue(nervatype) && (ref_id === null)){
        //add auto deffields
        var _sql = {select:["df.fieldname","ft.groupvalue as fieldtype"],
          from:"deffield df",
          inner_join:[["groups nt","on",[["df.nervatype","=","nt.id"],
            ["and","nt.groupvalue","=","'"+nervatype+"'"]]],
            ["groups ft","on",["df.fieldtype","=","ft.id"]]],
          where:[["df.deleted","=","0"],["and","df.addnew","=","1"],
            ["and","df.visible","=","1"]]}
        trans.query(models.getSql(self.engine, _sql), [], function (error, data) {
          if (error) {
            callback(error);}
          else {
            data.rows.forEach(function(field) {
              if (typeof check_values.fvalues[field.fieldname] === "undefined") {
                switch (field.fieldtype) {
                  case "bool":
                    check_values.fvalues[field.fieldname] = "false";
                    break;
                  case "integer":
                  case "float":
                    check_values.fvalues[field.fieldname] = 0;  
                    break;
                  default:
                    check_values.fvalues[field.fieldname] = "";
                    break;}}});
            callback(null, ref_id, check_values);}});}
      else{callback(null, ref_id, check_values);}
    },
    
    function(ref_id, check_values, callback) {
      if (valid_fieldvalue(nervatype)){
        //check deffields
        var _sql = [
          {select:["df.fieldname as fieldname","fv.id as fieldvalue_id"],
          from:"deffield df",
          inner_join:["groups nt","on",[["df.nervatype","=","nt.id"],
            ["and","nt.groupvalue","=","'"+nervatype+"'"]]],
          left_join:["fieldvalue fv","on",[["df.fieldname","=","fv.fieldname"],
            ["and","fv.deleted","=","0"],["and","fv.ref_id","=",ref_id]]]},
          {union_select:["'fieldtype_string' as fieldname","id as fieldvalue_id"], from:"groups",
            where:[["groupname","=","'fieldtype'"],["and","groupvalue","=","'string'"]]},
          {union_select:["'nervatype_id' as fieldname","id as fieldvalue_id"], from:"groups", 
            where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'"+nervatype+"'"]]}]
        trans.query(models.getSql(self.engine, _sql), [], function (error, data) {
          if (error) {
            callback(error);}
          else {
            data.rows.forEach(function(field) {
              if (typeof check_values.dvalues[field.fieldname] === "undefined") {
                check_values.dvalues[field.fieldname] = [];}
              if (field.fieldvalue_id !== null) {
                check_values.dvalues[field.fieldname].push(field.fieldvalue_id);}});
            callback(null, ref_id, check_values);}});}
      else{callback(null, ref_id, check_values);}
    },
    
    function(ref_id, check_values, callback) {
      if ((Object.keys(check_values.values).length>0) && (params.validate === true)){
        //validate
        var _sql = [
          {select:["g.id as id","g.groupname as groupname","g.groupvalue as groupvalue"],
            from:"groups g", where: ["g.deleted","=","0"]},
          {union_select:["c.id as id","'curr' as groupname","c.curr as groupvalue"],
            from: "currency c"}]
        trans.query(models.getSql(self.engine, _sql), [], function (error, data) {
          if (error) {callback(error);} else {
            var groups ={}; var curr = {};
            data.rows.forEach(function(group) {
              if (group.groupname === "curr") {
                curr[group.groupvalue] = group.id;}
              else {
                groups[group.id] = group;}});
            var req_error = null;
            for (var fieldname in check_values.values) {
              if (!req_error){
                var value = check_values.values[fieldname];
                var field = ntura.model[nervatype][fieldname];
                switch (field.type) {
                  case "integer":
                    if (!field.hasOwnProperty("requires")){
                      if (value === null){
                        if (field.notnull) {
                          check_values.values[fieldname] = 0;}}
                      else if (isNaN(parseInt(value,10))){
                        req_error = self.lang.invalid_value+": "+value+" (integer)";}
                      else {
                        check_values.values[fieldname] = parseInt(value,10);}}
                    break;
                  case "float":
                    if (value === null){
                      if (field.notnull) {
                        check_values.values[fieldname] = 0;}}
                    else if (isNaN(parseFloat(value))){
                      req_error = self.lang.invalid_value+": "+value+" (float)";}
                    else {
                      check_values.values[fieldname] = parseFloat(value);}
                    break;
                  case "date":
                    if ((value === null || value === "") && !field.notnull){
                      check_values.values[fieldname] = null;}
                    else {
                      check_values.values[fieldname] = out.getValidDate(value);}
                    break;
                  case "datetime":
                    if ((value === null || value === "") && !field.notnull){
                      check_values.values[fieldname] = null;}
                    else {
                      check_values.values[fieldname] = out.getValidDateTime(value);}
                    break;
                  case "password":
                  case "string":
                  case "text":
                    if ((value === null || value === "") && field.notnull && field.hasOwnProperty("default")){
                      check_values.values[fieldname] = out.replaceAll(field.default,"'","");}
                    else if (value === "" && !field.notnull){
                      check_values.values[fieldname] = null;}
                    else if (field.hasOwnProperty("length") && value !== null) {
                      if (value.toString().length > field.length){
                        check_values.values[fieldname] = value.substr(0,parseInt(field.length,10));}}
                    break;}
                if ((check_values.values[fieldname] === null) && (field.notnull === true)) {
                  req_error = self.lang.missing_required_field+" "+fieldname;}
                if (!req_error && typeof field.requires !== "undefined") {
                  if (field.requires.hasOwnProperty("bool")){
                    if (["true", "True", "TRUE", "t", "T", "y", "YES", "yes", 1, "1"].indexOf(value)>-1){
                      check_values.values[fieldname] = 1;}
                    else {
                      check_values.values[fieldname] = 0;}}
                  else if (field.requires.hasOwnProperty("curr")){
                    if (value === "" || value === null){
                      check_values.values[fieldname] = null}
                    else {
                      if (typeof curr[value] === "undefined") {
                        req_error = self.lang.invalid_value.replace("value",fieldname)+": "+value;}}}
                  else if (field.requires.hasOwnProperty("min") || field.requires.hasOwnProperty("max")){
                    if (field.type === "integer" || field.type === "float"){
                      if (field.requires.hasOwnProperty("min") && check_values.values[fieldname]<field.requires.min){
                        req_error = self.lang.invalid_value+": "+check_values.values[fieldname]+" ("+fieldname+" min. value:"+field.requires.min+")";}
                      if (field.requires.hasOwnProperty("max") && check_values.values[fieldname]>field.requires.max){
                        req_error = self.lang.invalid_value+": "+check_values.values[fieldname]+" ("+fieldname+" max. value:"+field.requires.max+")";}}}
                  else {
                    if (value === "" || value === null){
                      check_values.values[fieldname] = null}
                    else if (!groups.hasOwnProperty(value)){
                      req_error = self.lang.invalid_value+": "+value;}
                    else {
                      var gvalid;
                      for (var req in field.requires) {
                        if (groups[value].groupname === req) {
                          gvalid = true;}
                        else if (field.requires[req].length>0){
                          if (field.requires[req].indexOf(groups[value].groupvalue) > -1){
                            gvalid = true;}}}
                      if (!gvalid){
                        req_error = self.lang.invalid_value+": "+groups[value].groupvalue;}}}}}}
            if (req_error) {callback(req_error);}
            else {callback(null, ref_id, check_values);}}});}
      else {callback(null, ref_id, check_values);}},
    
    function(ref_id, check_values, callback) {
      if ((Object.keys(check_values.values).length>0) && (nervatype === "fieldvalue")){
        //check_fieldvalue (fieldvalue table)
        if (typeof check_values.values.fieldname === "undefined") {
          callback(self.lang.missing_fieldname);}
        else {
          if (typeof check_values.values.value === "undefined") {
            check_values.values.value = "";}
          var fparams = {conn:trans, fieldname:check_values.values.fieldname, value:check_values.values.value};
          check_fieldvalue(self, fparams, function(err, value){
            if (err) {
              callback(err);}
            else {
              check_values.values.value = value;
              callback(null, ref_id, check_values);}});}}
      else {callback(null, ref_id, check_values);}},
    
    function(ref_id, check_values, callback) {
      if (Object.keys(check_values.values).length>0){
        //update data
        var _sql; var prm = [];
        if (ref_id !== null) {
          _sql = {update:nervatype, set:[[]], where:["id","=",ref_id]}}
        else {
          _sql = {insert_into:[nervatype,[[]]], values:[[]]}}
        for (var fieldname in check_values.values) {
          var pchar = "?"
          if (check_values.values[fieldname] === null){
            pchar = "null";}
          if (ref_id !== null) {
            _sql.set.push([fieldname,"=",pchar]);} 
          else {
            _sql.insert_into[1].push(fieldname);
            _sql.values.push(pchar);}
          if (pchar === "?"){
            prm.push(check_values.values[fieldname]);}}
        trans.query(models.getSql(self.engine, _sql), prm, function (error, data) {
          if (error) {
            callback(error);}
          else {
            if (ref_id === null) {
              if(!data.lastInsertId || data.lastInsertId === null || data.lastInsertId === 0){;
                trans.query(models.getLastID(self.engine,nervatype), [], function (error, data) {
                  if (error) {callback(error);}
                  else {
                    callback(null,parseInt(data.rows[0].id,10), check_values);}});}
              else{
                callback(null, data.lastInsertId, check_values);}}
            else {
              callback(null, ref_id, check_values);}}});}
      else {callback(null, ref_id, check_values);}},
    
    function(ref_id, check_values, callback) {
      if (Object.keys(check_values.fvalues).length>0){
        //update additional data
        //check_fieldvalue
        for (var fieldname in check_values.fvalues) {
          var field_index = 1; var field_id = null; var notes = "";
          var fieldname_err = null;
          var value = check_values.fvalues[fieldname]||"";
          if (fieldname.split("~").length>1){
            field_index = parseInt(fieldname.split("~")[1],10);
            if (isNaN(field_index) || field_index < 1) {field_index = 1;}
            fieldname = fieldname.split("~")[0];}
          if (typeof check_values.dvalues[fieldname] === "undefined") {
            if (params.insert_field === false) {
              fieldname_err = self.lang.missing_insert_field+" "+fieldname;}
            check_values.deffield.push(
              [fieldname, check_values.dvalues.nervatype_id[0], 
              check_values.dvalues.fieldtype_string[0],fieldname]);}
          else if (check_values.dvalues[fieldname].length >= field_index) {
            field_id = check_values.dvalues[fieldname][field_index-1]}
          if (value.toString().split("~").length>1){
            notes = value.toString().split("~")[1];
            value = value.toString().split("~")[0];}
          check_values.fieldvalue.push([field_id, fieldname, value, notes]);}
        if (fieldname_err !== null) {callback(fieldname_err);}
        else {callback(null, ref_id, check_values);}}
      else {callback(null, ref_id, check_values);}},
    
    function(ref_id, check_values, callback) {
      if (check_values.deffield.length>0){
        //create new deffields
        var values_lst = [];
        var sql = models.getSql(self.engine,
          {insert_into:["deffield",[[],"fieldname","nervatype","fieldtype","description",
            "addnew","visible","readonly"]],
          values:[[],"?","?","?","?","0","1","0"]});
        check_values.deffield.forEach(function(prm) {
          values_lst.push(function(callback_){
            trans.query(sql, prm, 
              function (err, data) {
                callback_(err,data);});});});
        async.series(values_lst,function(err,data) {
          callback(err, ref_id, check_values);;});}
      else {callback(null, ref_id, check_values);}},
    
    function(ref_id, check_values, callback) {
      if (check_values.fieldvalue.length>0){
        //check_fieldvalue (!== fielvalue table)
        var values_lst = [];
        check_values.fieldvalue.forEach(function(fieldvalue) {
          var fparams = {conn:trans, fieldname:fieldvalue[1], 
              value:fieldvalue[2], ref_id:check_values.fieldvalue.indexOf(fieldvalue)};
          values_lst.push(function(callback_){
            check_fieldvalue(self, fparams, function(err, value, ref_index){
              if(!err){
                check_values.fieldvalue[ref_index][2] = value;}
              callback_(err,value)});});});
        async.series(values_lst,function(err,data) {
          callback(err, ref_id, check_values);;});}
      else {callback(null, ref_id, check_values);}},
    
    function(ref_id, check_values, callback) {
      if (check_values.fieldvalue.length>0){
        //create new fieldvalues
        var values_lst = [];
        check_values.fieldvalue.forEach(function(fvalues) {
          values_lst.push(function(callback_){
            var _sql;
            if (fvalues[0] === null){
              _sql = {insert_into:["fieldvalue",[[],"fieldname","ref_id","value","notes"]], 
                values:[[],"'"+fvalues[1]+"'",ref_id,"'"+fvalues[2]+"'","'"+fvalues[3]+"'"]}}
            else {
              _sql = {update:"fieldvalue", 
                set:[[],["value","=","'"+fvalues[2]+"'"],["notes","=","'"+fvalues[3]+"'"]], 
                where:["id","=",fvalues[0]]}}
            trans.query(models.getSql(self.engine, _sql), [], function (err, data) {
              callback_(err, data);});});});
        async.series(values_lst,function(err,data) {
          callback(err, ref_id);;});}
      else {callback(null, ref_id);}},
    
    function(ref_id, callback) {
      if (params.log_enabled === true){
        //insert log
        insert_log(self,{conn:trans, logstate:"update", nervatype:nervatype, ref_id:ref_id}, 
          function(err, result){
            callback(err, ref_id);});}
      else {callback(null, ref_id);}}
    
    ],
  function(err, ref_id) {
    if(err){if(err.message){err = err.message;}}   
    if (err && !params.transaction){
      if(trans){
        if(trans.rollback){
          trans.rollback(function (error) {
            if (!error && !params.validator.conn){
              conn.close();}
            _callback(err, ref_id);});}
        else{
          if (!params.validator.conn){
            conn.close();}
          _callback(err, ref_id);}}
      else{
        _callback(err, ref_id);}}
    else if (!err && !params.transaction){
      if(trans){
        if(trans.commit){
          trans.commit(function (err) {
            if (!err && !params.validator.conn){
              conn.close();}
            _callback(err, ref_id);});}
        else{
          if (!params.validator.conn){
            conn.close();}
          _callback(err, ref_id);}}
      else{
        _callback(err, ref_id);}}
    else {
      _callback(err, ref_id);}});}

//local    
function set_sql_params(self, params) {
  if (typeof params.rlimit === "undefined") {params.rlimit = self.limit;}
  if (typeof params.rowlimit === "undefined") {params.rowlimit = self.limit_row;}
  if (typeof params.orderbyStr === "undefined") {params.orderbyStr = "";}
  if (params.paramList !== null) {
    params.paramList.forEach(function(prm) {
      switch(prm.type) {
        case "string":
        case "date":
          if (prm.value.indexOf("'")===-1) {
            prm.value = "'"+prm.value+"'";}
          break;
        default:
          prm.value = prm.value;
      }
      switch(prm.wheretype) {
        case "where":
          params.whereStr = out.replaceAll(params.whereStr,prm.name,prm.value);
          break;
        case "having":
          params.havingStr = out.replaceAll(params.havingStr,prm.name,prm.value);
          break;
        case "in":
          params.sqlStr = out.replaceAll(params.sqlStr,prm.name,prm.value);
          break;
      }});}
  params.sqlStr = params.sqlStr.replace(/@where_str/g, params.whereStr.toString());
  params.sqlStr = params.sqlStr.replace(/@having_str/g, params.havingStr.toString());
  params.sqlStr = params.sqlStr.replace(/@orderby_str/g, params.orderbyStr.toString());
  if (params.rlimit === true) {
    params.sqlStr = params.sqlStr.replace(";", "");
    params.sqlStr = params.sqlStr + " limit " + params.rowlimit.toString();}
  return params.sqlStr.toString();};

function set_engine(self, params, _callback) {
  if (typeof params.check_ndi === "undefined") {params.check_ndi = false;}
  if (typeof params.created === "undefined") {params.created = false;}
  if (typeof params.createdb === "undefined") {params.createdb = false;}
  
  async.waterfall([    
    function(callback) {
      self.storage.getDbsFromAlias(params.database, function(err, database){
        if(err !== null){
          if(err.message === "missing" || err === "missing"){
            err = self.lang.missing_database+" "+params.database;}}
        callback(err,database);});},
    
    function(database, callback) {
      self.ndi_enabled = self.app_config.checkValue(database.settings.ndi_enabled, database.settings.ndi_enabled);
      self.encrypt_password = self.app_config.checkValue(database.settings.encrypt_password, database.settings.encrypt_password);
      var host_restriction = self.app_config.checkValue(database.settings.dbs_host_restriction, database.settings.dbs_host_restriction)
      if (host_restriction!==""){
        if (host_restriction.split(",").indexOf(self.host_ip)===-1) {
          callback("Database "+self.lang.insecure_err);}
        else {callback(null,database);}}
      else {callback(null,database);}},
      
    function(database, callback) {
      self.engine = self.app_config.checkValue(database.engine, database.engine);
      var config = connect.getConfig(self.engine, database.connect, database.settings, self.data_dir, self.app_config);
      if (config.adapter === null){
        callback(self.lang.unsupport_db);}
      else {
        config = (typeof config.uri !== "undefined") ? config.uri : config;
        connect.createConnection({connect:config, pool:false}, function(err, conn){
          if (err) {
            callback(err);}
          else {
            conn.end();
            self.dbs_config = config; callback(null);}});}}
  ],
  function(err) {
    if(err){if(err.message){err = err.message;}}
    _callback(err, (err===null));})}

function set_token_settings(self, params) {
  self.token_settings = { 
    key: params.key, 
    options: params.options || {} 
  }
}

module.exports = function(params) {
  var self = this;
  self.storage = params.storage;
  self.data_dir = params.data_dir;
  self.report_dir = params.report_dir;
  self.lang = params.lang;
  self.ndi_enabled = true;
  self.encrypt_password = null;
  self.app_config = params.conf;
  self.host_ip = params.host_ip;
  self.host_settings = params.host_settings;
  self.dbs_config = null;
  self.limit = false; 
  self.limit_row = 500;
  self.employee = null;
  self.token_settings = params.token_settings || params.conf.token_settings || {};
    
return {
  storage: function(){return self.storage;},
  report_dir: function(){return self.report_dir;},
  lang: function(){return self.lang;},
  ndi_enabled: function(){return self.ndi_enabled;},
  encrypt_password: function(){return self.encrypt_password;},
  dbs_config: function(){return self.dbs_config;},
  app_config: function(){return self.app_config;},
  host_settings: function(){return self.host_settings;},
  employee: function(){return self.employee;},
  engine: function(){return self.engine;},
  token_settings: function(){return self.token_settings;},
  
  valid: {
    //general and validator functions
    getTableKey: function(nervatype){
      return get_table_key(nervatype)},
    getRefnumber: function(params, _callback){
      return get_refnumber(self, params, _callback)},
    //checkFieldValue: function(params, _callback) {
    //  check_fieldvalue(self, params, _callback);},
    getIdFromRefnumber: function(params, _callback) {
      get_id_from_refnumber(self, params, _callback);},
    getGroupsId: function(params, _callback) {
      get_groups_id(self, params, _callback);},
    decodeToken: function(token) {
      return decode_token(self, token); },
    verifyToken: function(token, _callback) {
      return verify_token(self, token, _callback); }
  },
  
  connect: {
    //dbs connect, user login, authentication, data access
    getConnect: function() {
      return get_connect(self)},
    getLogin: function(params, _callback){
      get_login(self, params, _callback);},
    nextNumber: function(params, _callback) {
      next_number(self, params, _callback);},
    updateData: function(params, _callback) {
      update_data(self, params, _callback);},
    deleteData: function(params, _callback) {
      delete_data(self, params, _callback);},
    insertLog: function(params, _callback) {
      insert_log(self, params, _callback);},
    getDataAudit: function(params, _callback) {
      get_data_audit(self, params, _callback)},
    getObjectAudit: function(params, _callback) {
      get_object_audit(self, params, _callback)},
    getDatabaseSettings: function(params, _callback) {
      get_database_settings(self, params, _callback)}
  },
  
  local: {
    //server side data
    setSqlParams: function(params){
      return set_sql_params(self, params);},
    setEngine: function(params, callback) {
      set_engine(self, params, callback);},
    setTokenSettings: function(params){
      return set_token_settings(self, params);}
  }
}};
