/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var async = require("async");
var fs = require('fs');
var zlib = require('zlib');
var path = require('path');
var xml2js = require('xml2js');
var json2csv = require('json2csv');

var ntura = require('./models.js');
var connect = require('./adapter.js').connect();
var models = require('./adapter.js').models();
var out = require('./tools.js').DataOutput();
var tool = require('./tools.js').NervaTools();

module.exports = function() {

function valid(req){
  if (req.app.get("host_settings").nas_host_restriction.length>0 && req.app.get("host_settings").nas_host_restriction.indexOf(req.ip)===-1){
    return "/login/insecure";}
  else if (req.app.get("host_settings").nas_host_restriction.length===0 && req.app.get("host_settings").all_host_restriction.length>0 
    && req.app.get("host_settings").all_host_restriction.indexOf(req.ip)===-1){
    return "/login/insecure";}
  else if (req.user){
    return "ok";}
  else {return "/login";}}

function get_nstore_prm(req){
  return { conf: req.app.get("conf"), data_dir: req.app.get("data_dir"), report_dir: req.app.get("report_dir"),
    host_ip: req.ip, host_settings: req.app.get("host_settings"), storage: req.app.get("storage"),
    lang: req.app.locals.lang }
}

function valid_setting(nas_login, params, lang){
  switch (params.fieldname) {
    case "nas_auth":
      switch (params.value) {
        case "local":
          return(null);
        case "amazon":
          if(nas_login.amazon.clientID && nas_login.amazon.clientSecret){
              return(null);}
          else {
            return(lang.missing_params+": clientID, clientSecret");}
        case "azure":
          if(nas_login.azure.clientID && nas_login.azure.clientSecret 
            && nas_login.azure.redirectUrl && nas_login.azure.identityMetadata){
              return(null);}
          else {
            return(lang.missing_params+": clientID, clientSecret, redirectUrl, identityMetadata");}
        case "google":
          if(nas_login.google.clientID && nas_login.google.clientSecret){
              return(null);}
          else {
            return(lang.missing_params+": clientID, clientSecret");}  
        default:
          return(lang.invalid_fieldname+" "+params.value+" (local, amazon, azure, google)");}
  
    default:
      return(null);}}

function page_render(params){
  if (!params.data){
    params.data = {};}
  if (!params.dir){
    params.dir = "nas"
    params.data.page_view = params.page;}
  else {
    params.data.view = params.page;}
  if(params.data.view){
    params.data.page_view += "_"+params.data.view;}
  params.res.render(params.dir+"/index.html",params.data);
}
  
function user_list(params){
  params.req.app.settings.storage.getUsers(function(err, users){
    if (err) {return params.next(err);}
    else {
      params.page = "user";
      params.data.users = users;
      params.data.subtitle = params.req.app.locals.lang.label_accounts;
      page_render(params);}});}

function database_list(params){
  params.req.app.settings.storage.getDatabases(function(err, data){
    if (err) {return params.next(err);}
    else {
      params.page = "database";
      params.data.data = data;
      if (!params.data.view){
        params.data.view = "list";
        params.data.subtitle = params.req.app.locals.lang.label_databases;}
      page_render(params);}});}

function import_list(params){
  fs.readdir(params.data.import_dir, function(err, files){
    if (err) {return params.next(err);}
    else {
      files.forEach(function(filename) {
        if(filename.indexOf(".xml")>-1 || filename.indexOf(".data")>-1){
          params.data.filenames.push(filename);}});
      params.data.filenames.unshift("");
      database_list(params);}});}

function create_database(nstore, params, _callback) {
  var conn = params.conn; var trans; var logstr = ""; var results=[];
  if(params.logstr){logstr = params.logstr;}
  async.waterfall([
    function(callback) {
      if(logstr === ""){
        if(params.logtype === "json"){
          results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_database_alias+': '+params.database });
          results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_start_process });}
        else {
          logstr += '<div><span style="font-weight: bold;">'+nstore.lang().log_database_alias+': '
            +params.database+'</span></div><br>';
          logstr += '<div><span>'+nstore.lang().log_start_process+': '
            +out.getISODateTime(new Date(),true)+'</span></div>'; }}
      //check connect
      nstore.local.setEngine({database:params.database}, function(err,result){
        callback(err);});},
        
    function(callback) {
      //drop all tables if exist
      if (!conn){conn = nstore.connect.getConnect();}
      if (!conn){return _callback(nstore.lang().not_connect);}
      trans = connect.beginTransaction({connection:conn, engine:nstore.engine()});
      if(params.logtype === "json"){
        results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_drop_table }); }
      else{
        logstr += '<div><span>'+nstore.lang().log_drop_table+'</span></div>'; }
      var drop_lst = models.dropList(nstore.engine());
      var value_lst = [];
      drop_lst.forEach(function(sql) {
        value_lst.push(function(callback_){
          trans.query(sql, [], function (err, data) {
            callback_(err,data);});});});            
      async.series(value_lst,function(err, data) {
        if(!err && trans.commit){
          trans.commit(function (err) {
            callback(null);});}
        else {
          if(trans.rollback){trans.rollback();}
          callback(null);}});},
    
    function(callback) {
      //create all tables
      if(params.logtype === "json"){
        results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_create_table }); }
      else{
        logstr += '<div><span>'+nstore.lang().log_create_table+'</span></div>'; }
      trans = connect.beginTransaction({connection:conn, engine:nstore.engine()});
      var create_lst = models.modelList(nstore.engine());
      var value_lst = [];
      create_lst.forEach(function(sql) {
        value_lst.push(function(callback_){
          trans.query(sql, [], function (err, data) {
            callback_(err,data);});});});
      async.series(value_lst,function(err, data) {
        callback(err);});},
    
    function(callback) {
      //create indexes
      if(params.logtype === "json"){
        results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_create_index }); }
      else {
        logstr += '<div><span>'+nstore.lang().log_create_index+'</span></div>'; }
      var index_lst = models.indexList(nstore.engine());
      var value_lst = [];
      index_lst.forEach(function(sql) {
        value_lst.push(function(callback_){
          trans.query(sql, [], function (err, data) {
            callback_(err,data);});});});
      async.series(value_lst,function(err,data) {
        callback(err);});},
            
    function(callback) {
      if(!params.empty){
        //insert data
        if(params.logtype === "json"){
          results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_init_data }); }
        else {
          logstr += '<div><span>'+nstore.lang().log_init_data+'</span></div>'; }
        var data_lst = models.dataList(nstore.engine())
        var value_lst = [];
        data_lst.forEach(function(sql) {
          value_lst.push(function(callback_){
            trans.query(sql, [], function (err, data) {
              callback_(err,data);});});});
        async.series(value_lst,function(err, data) {
          callback(err);});}
      else{callback(null);}},
      
    function(callback) {
      if(trans.commit){trans.commit();}
      if(models.compact(nstore.engine()) !== null){
        conn.query(models.compact(nstore.engine()), [], function (err, data) {
          if (!err){
            if(params.logtype === "json"){
              results.push({stamp: out.getISODateTime(new Date(),true), state: 'Rebuilding the database...' }); }
            else {
              logstr += '<div><span>Rebuilding the database...</span></div>'; }}
          callback(err);});}
      else {callback(null);}}
    ],
  function(err) {
    if(err){if(err.message){err = err.message;}}
    if (!err){
      if(!params.logstr){
        if(params.logtype === "json"){
          results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().info_create_ok }); }
        else{
          logstr += '<div><span style="font-weight: bold;">'+nstore.lang().info_create_ok+'</span></div>'; }}}
    else {
      if(trans){if(trans.rollback){trans.rollback();}}
      if(params.logtype === "json"){
        results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_error+': '+err }); }
      else{
        logstr += '<div><span style="color:red;font-weight: bold;">'+nstore.lang().log_error+': '+err+'</span></div>'; }}
    if (conn && !params.conn){conn.close();}
    if(!params.logstr){
      if(params.logtype === "json"){
        results.push({stamp: out.getISODateTime(new Date(),true), state: nstore.lang().log_end_process }); }
      else{
        logstr += '<div><span>'+nstore.lang().log_end_process+': '+out.getISODateTime(new Date(),true)+'</span></div>'; }
      if (!err){
        if(params.logtype !== "json"){
          logstr += '<br><div><span>'+nstore.lang().log_create_demo+': </span><br>';
          logstr += '<a href="/ndi/demo/create?database='+params.database
            +'&username=demo" target="_blank" data-ajax="false" >/ndi/demo/create?database='
            +params.database+'&username=demo</a></div>'; }}}
    _callback(err, (params.logtype === "json") ? results : logstr);});}

function import_database(nstore, params, _callback){
  var conn; var logstr = "";
  if(params.logstr){logstr = params.logstr;}
  async.waterfall([
    function(callback) {
      logstr = '<div><span style="font-weight: bold;">'
        +nstore.lang().log_database_alias+': '+params.database+'</span><br>';
      logstr += '<span>'+nstore.lang().label_export_file+': '
        +params.filename+'</span><br>';
      logstr += '<div><span>'+nstore.lang().log_start_process+': '
        +out.getISODateTime(new Date(),true)+'</span></div>';
      if (!params.filename || params.filename===""){
        callback(nstore.lang().missing_required_field+" "+nstore.lang().label_export_file);}
      else if (!params.database || params.database===""){
        callback(nstore.lang().missing_required_field+" "+nstore.lang().label_database);}
      else {
        //check connect
        nstore.local.setEngine({database:params.database}, function(err,result){
          callback(err);});}},
        
    function(callback) {
      conn = nstore.connect.getConnect();
      if (!conn){
        callback(nstore.lang().not_connect);}
      else {
        logstr += '<div><span>'+nstore.lang().log_load_data+'</span></div>';
        fs.readFile(path.join(params.import_dir, params.filename),function(err, data) {
          callback(err, data);});}},
    
    function(data, callback) {
      if(params.filename.indexOf(".data") > -1){
        zlib.inflate(data, function (err, idata) {
          callback(err, idata);});}
      else {
        callback(null, data);}},
        
    function(data, callback) {
      var parser = new xml2js.Parser();
      parser.parseString(data, function (err, xdata) {
        callback(err, xdata);});},
    
    function(data, callback) {
      var dbs_params = {conn:conn, database:params.database, 
        logstr:logstr, empty:true}
      create_database(nstore, dbs_params, function(err, create_log){
        logstr = create_log;
        callback(null, data);});},
    
    function(data, callback) {
      var ndi = require('./ndi.js')();
      var value_lst = [];
      var ndi_params = {
        validator:{conn:conn}, log_enabled:false, insert_row:true, 
          insert_field:true, check_audit:false, use_deleted:true}
      models.exportList.nom.forEach(function(nom) {
        if(data.data[nom]){
          var items = [];
          data.data[nom].forEach(function(fields) {
            var item = {};
            fields.field.forEach(function(field) {
              item[field.name[0]] = field.value[0];});
            items.push(item);});}
        value_lst.push(function(callback_){
          ndi_params.datatype = nom;
          ndi.updateData(nstore, ndi_params, items, function(err, result){
            if(!err){
              logstr += '<div><span>'+nom+": "+result.data.length+' rows</span></div>';}
            callback_(err);});});});
      async.series(value_lst,function(err) {
        callback(err,data);})},
    
    function(data, callback) {
      var value_lst = [];
      models.exportList.ui_1.forEach(function(ui) {
        var ui_count = 0;
        if(data.data[ui]){
          data.data[ui].forEach(function(fields) {
            var _sql = {}; var prm = []; ui_count +=1;
            _sql = {insert_into:[ui,[[]]], values:[[]]};
            fields.field.forEach(function(field) {
              if(field.value[0] !== ""){
                _sql.insert_into[1].push(field.name[0]);
                _sql.values.push("?");
                prm.push(field.value[0]);}});
            value_lst.push(function(callback_){
              conn.query(models.getSql(nstore.engine(),_sql), prm, function (err, idata) {
                callback_(err);});});});}
          logstr += '<div><span>'+ui+": "+ui_count+' rows</span></div>';});
      async.series(value_lst,function(err) {
        callback(err,data);})},
    
    function(data, callback) {
      var _sql = [
        {select:["'empnumber' as keyname","'employee_id' as refname","empnumber as vkey","id"], from:"employee"},
        {union_select:["'menukey' as keyname","'menu_id' as refname","menukey as vkey","id"], from:"ui_menu"},
        {union_select:["'reportkey' as keyname","'report_id' as refname","reportkey as vkey","id"], from:"ui_report"},
        {union_select:["'usergroup' as keyname","'usergroup' as refname","groupvalue as vkey","id"], 
          from:"groups", where:["groupname","=","'usergroup'"]},
        {union_select:["'nervatype' as keyname","'nervatype' as refname","groupvalue as vkey","id"], 
          from:"groups", where:["groupname","=","'nervatype'"]},
        {union_select:["'transtype' as keyname","'transtype' as refname","groupvalue as vkey","id"], 
          from:"groups", where:["groupname","=","'transtype'"]},
        {union_select:["'inputfilter' as keyname","'inputfilter' as refname","groupvalue as vkey","id"], 
          from:"groups", where:["groupname","=","'inputfilter'"]},
        {union_select:["'fieldtype' as keyname","'fieldtype' as refname","groupvalue as vkey","id"], 
          from:"groups", where:["groupname","=","'fieldtype'"]},
        {union_select:["'wheretype' as keyname","'wheretype' as refname","groupvalue as vkey","id"], 
          from:"groups", where:["groupname","=","'wheretype'"]}]
      conn.query(models.getSql(nstore.engine(),_sql), [], function (err, edata) {
        if (!err){
          edata.rows.forEach(function(row) {
            if (!data[row.keyname]){data[row.keyname] = {};}
            data[row.keyname][row.vkey] = {refname:row.refname, id:row.id};});}
        callback(err,data);});},
  
    function(data, callback) {
      var value_lst = [];
      models.exportList.ui_2.forEach(function(ui) {
        var ui_count = 0;
        if(data.data[ui]){
          data.data[ui].forEach(function(fields) {
            var _sql = {}; var prm = []; ui_count +=1;
            _sql = {insert_into:[ui,[[]]], values:[[]]};
            var last_nervatype;
            fields.field.forEach(function(field) {
              if(field.value[0] !== ""){
              switch (field.name[0]) {
                case "subtype":
                  _sql.insert_into[1].push("subtype");
                  _sql.values.push("?");
                  if(last_nervatype === "trans"){  
                    prm.push(data.transtype[field.value[0]].id);}
                  else if(last_nervatype === "report"){
                    prm.push(data.reportkey[field.value[0]].id);}
                  else{prm.push(null);}
                  break;
                case "empnumber":
                case "menukey":
                case "reportkey":
                case "usergroup":
                case "nervatype":
                case "inputfilter":
                case "fieldtype":
                case "wheretype":
                  _sql.insert_into[1].push(data[field.name[0]][field.value[0]].refname);
                  _sql.values.push("?");
                  prm.push(data[field.name[0]][field.value[0]].id);
                  if(field.name[0]==="nervatype"){
                    last_nervatype = field.value[0];}
                  break;
                default:
                  _sql.insert_into[1].push(field.name[0]);
                  _sql.values.push("?");
                  prm.push(field.value[0]);
                  break;}}});
            value_lst.push(function(callback_){
              conn.query(models.getSql(nstore.engine(),_sql), prm, function (err, idata) {
                callback_(err);});});});}
          logstr += '<div><span>'+ui+": "+ui_count+' rows</span></div>';});
      async.series(value_lst,function(err) {
        callback(err);})},
  
    function(callback) {
      if(models.compact(nstore.engine()) !== null){
        conn.query(models.compact(nstore.engine()), [], function (err, data) {
          if (!err){
            logstr += '<div><span>Rebuilding the database...</span></div>';}
          callback(err);});}
      else {callback(null);}}
  ],
  function(err) {
    if(err){if(err.message){err = err.message;}}
    if (!err){
      logstr += '<div><span style="font-weight: bold;">'+nstore.lang().info_import_ok+'</span></div>';}
    else {
      logstr += '<div><span style="color:red;font-weight: bold;">'+nstore.lang().log_error+': '+err+'</span></div>';}
    if (conn){conn.close();}
    logstr += '<div><span>'+nstore.lang().log_end_process+': '+out.getISODateTime(new Date(),true)+'</span></div>';
    _callback(err, logstr);});}

function export_database(nstore, params, _callback){
  var conn; var timestamp = Date.now(); var logstr = "";
  if(params.logstr){logstr = params.logstr;}
  async.waterfall([
    function(callback) {
      logstr = '<div><span style="font-weight: bold;">'
        +nstore.lang().log_database_alias+': '+params.database+'</span><br>';
      logstr += '<span>'+nstore.lang().label_format+': '
        +params.format+'</span><br>';
      logstr += '<span>Ver.No: '+params.version+'</span></div><br>';
      logstr += '<div><span>'+nstore.lang().log_start_process+': '
        +out.getISODateTime(new Date(),true)+'</span></div>';
      //check connect
      nstore.local.setEngine({database:params.database}, function(err,result){
        callback(err);});},
    
    function(callback) {
      conn = nstore.connect.getConnect();
      if (!conn){
        callback(nstore.lang().not_connect);}
      else {
        var ndi = require('./ndi.js')();
        var value_lst = [];
        var ndi_params = {validator:{conn:conn}, check_audit:false, use_deleted:true}
        var ndi_filter = {output:"json", no_deffield:true, orderby:"id"}
        models.exportList.nom.forEach(function(nom) {
          value_lst.push(function(callback_){
            ndi_params.datatype = nom;
            ndi.getData(nstore, ndi_params, ndi_filter, function (err, result) {
              if(!err){
                logstr += '<div><span>'+result.datatype+' rows: '
                  +result.data.length+'</span></div>';}
              callback_(err,result);});});});
        async.series(value_lst,function(err, results) {
          callback(err, results);});}},
            
    function(nom_data, callback) {
      var value_lst = []; var value_index = 0;
      var ui_lst = models.exportList.ui_1.concat(models.exportList.ui_2);
      ui_lst.forEach(function(ui) {
        value_lst.push(function(callback_){
          var _sql;
          switch (ui) {
            case "ui_audit":
              _sql = {
                select:["ug.groupvalue as usergroup","nt.groupvalue as nervatype", 
                  "case when r.id is null then st.groupvalue else r.reportkey end as subtype", 
                  "ipf.groupvalue as inputfilter","a.supervisor"],
                from:"ui_audit a",
                inner_join:[["groups ug","on",["a.usergroup","=","ug.id"]],
                  ["groups nt","on",["a.nervatype","=","nt.id"]],
                  ["groups ipf","on",["a.inputfilter","=","ipf.id"]]],
                left_join:[["groups st","on",["a.subtype","=","st.id"]],
                  ["ui_report r","on",["a.subtype","=","r.id"]]],
                order_by:["a.id"]}
              break;
            case "ui_menufields":
              _sql = {
                select:["m.menukey","mf.fieldname","mf.description","ft.groupvalue as fieldtype","mf.orderby"],
                from:"ui_menufields mf",
                inner_join:[["ui_menu m","on",["mf.menu_id","=","m.id"]],
                  ["groups ft","on",["mf.fieldtype","=","ft.id"]]],
                order_by:["mf.id"]}
              break;
            case "ui_reportfields":
              _sql = {
                select:["r.reportkey","rf.fieldname","ft.groupvalue as fieldtype","wt.groupvalue as wheretype",
                  "rf.description","rf.orderby","rf.sqlstr","rf.parameter","rf.dataset","rf.defvalue","rf.valuelist"],
                from:"ui_reportfields rf",
                inner_join:[["ui_report r","on",["rf.report_id","=","r.id"]],
                  ["groups ft","on",["rf.fieldtype","=","ft.id"]],
                  ["groups wt","on",["rf.wheretype","=","wt.id"]]],
                order_by:["rf.id"]}
              break;
            case "ui_reportsources":
              _sql = {
                select:["r.reportkey","rs.dataset","rs.sqlstr"],
                from:"ui_reportsources rs",
                inner_join:["ui_report r","on",["rs.report_id","=","r.id"]],
                order_by:["rs.id"]}
              break;
            case "ui_userconfig":
              _sql = {
                select:["e.empnumber","c.section","c.cfgroup","c.cfname","c.cfvalue","c.orderby"],
                from:"ui_userconfig c",
                inner_join:["employee e","on",["c.employee_id","=","e.id"]],
                order_by:["c.id"]}
              break;
            default:
              _sql = {select:["*"], from:ui, order_by:["id"]}
              break;}
          conn.query(models.getSql(nstore.engine(),_sql), [], function (err, data) {
            if(!err){
              var result = {datatype:ui_lst[value_index], data:[]}
              value_index += 1;
              data.rows.forEach(function(row) {
                var item = [];
                for (var field in row) {
                  var item_field = {name:field, value:row[field]};
                  if (ntura.model.hasOwnProperty(result.datatype)){
                    if (ntura.model[result.datatype].hasOwnProperty(field)){
                      if (ntura.model[result.datatype][field].hasOwnProperty("references")){
                        item_field.type = ntura.model[result.datatype][field].references[0];}
                      else {
                        item_field.type = ntura.model[result.datatype][field].type;}}
                    else if (!item_field.type){
                      item_field.type = "reference";}}
                  item.push(item_field);};
                result.data.push(item);});
              logstr += '<div><span>'+result.datatype+' rows: '
                +result.data.length+'</span></div>';
              callback_(null,result);}
            else {callback_(err,null);}});});});
      async.series(value_lst,function(err, results) {
        callback(err, nom_data.concat(results));});},
    
    function(result, callback) {
      var xml_params = {dbs:result, version:params.version,
        timestamp:timestamp}
      params.res.render("nas/export.xml", xml_params, function(err,xml){
        callback(err, xml);});},
    
    function(result, callback) {
      if (params.format === "data"){
        params.bck_filename = params.database+"_"+timestamp+".data"
        zlib.deflate(result, function (err, buffer) {
          if(!err && params.filename === "download"){
            params.res.set('Content-Type', 'application/octet-stream');
            params.res.set('Content-Disposition', 'attachment;filename="'+params.bck_filename+'"');}
          callback(err, buffer);});}
      else {
        params.bck_filename = params.database+"_"+timestamp+".xml"
        if(params.filename === "download") {
          params.res.set('Content-Type', 'text/xml');
          params.res.set('Content-Disposition', 'attachment;filename="'+params.bck_filename+'"');}
        callback(null, result);}},
    
    function(result, callback) {
      if(params.filename !== "download") {
        fs.writeFile(path.join(params.export_dir, params.bck_filename), result, function (err) {
          if(!err){
            logstr += '<br><div><span>'+nstore.lang().label_export+': </span>';
            logstr += params.bck_filename+'</div>';}
          callback(err, result);});}
      else {callback(null, result);}}
  ],
  function(err, result) {
    if(err){if(err.message){err = err.message;}}
    if (conn){conn.close();}
    if(err || params.filename !== "download"){
      if (!err){
        logstr += '<div><span style="font-weight: bold;">'
          +nstore.lang().info_export_ok+'</span></div>';}
      else {
        logstr += '<div><span style="color:red;font-weight: bold;">'
          +nstore.lang().log_error+': '+err+'</span></div>';}
      logstr += '<div><span>'+nstore.lang().log_end_process+': '
        +out.getISODateTime(new Date(),true)+'</span></div>';}
    _callback(err, logstr, result);});}

function report_render(params){
  params.storage.getDatabases(function(err, data){
    if (err && params.flash === null) {
      params.flash = err;}
    params.page = "report";
    if (!params.data.view){
      params.data.view = "list";}
    if(data.length===0){
      data.unshift({doc:{alias:""}});}
    params.data.form.database = data;
    params.data.form.groups = [
      "", "bank", "cash", "customer", "delivery", "employee", "formula", "inventory", "invoice", "offer", 
      "order", "product", "production", "project", "rent", "report", "tool", "waybill", "worksheet"];
    page_render(params);});}

function get_dbs_reports(nstore, params, _callback) {
  var results = {report:{}, label:[]}; var conn;
  if (!params.database){params.database = "";}
  async.waterfall([  
    function(callback) {
      conn = nstore.connect.getConnect();
      if (!conn){
        callback(nstore.lang().not_connect);}
      else {
        if (params.database !== ""){
          var _sql;
          if (params.reportkey){
            _sql = {select:["*"], from:"ui_report", 
              where:["reportkey","=","'"+params.reportkey+"'"]}}
          else {
            _sql = {
              select:["r.id as id","r.reportkey as reportkey","nt.groupvalue as nervatype",
                "tt.groupvalue as transtype","dir.groupvalue as direction","r.repname",
                "r.description","r.label","ft.groupvalue as filetype"],
              from:"ui_report r",
              inner_join:[["groups nt","on",["r.nervatype","=","nt.id"]],
                ["groups ft","on",["r.filetype","=","ft.id"]]],
              left_join:[["groups tt","on",["r.transtype","=","tt.id"]],
                ["groups dir","on",["r.direction","=","dir.id"]]]}}
          conn.query(models.getSql(nstore.engine(), _sql), [], function (error, data) {
            if (error) {
              callback(error);}
            else {
              if (!params.reportkey){
                var reports = {};
                data.rows.forEach(function(report) {
                  reports[report.reportkey] = report;});
                results.report = reports;}
              else if (data.rowCount > 0){
                results.report = data.rows[0];}
              callback(null);}});}
        else {
          callback(null);}}},
        
    function(callback) {
      if ((params.database !== "") && (params.reportkey)){
        var _sql = {select:["*"], from:"ui_message", 
          where:["secname","like","?"]}
        conn.query(models.getSql(nstore.engine(), _sql), [params.reportkey+"%"], 
          function (error, data) {
            if (error) {
              callback(error);}
            else {
              results.label = data.rows;
              callback(null);}});}
      else {
        callback(null);}}
    ],
  function(err) {
    if(err){if(err.message){err = err.message;}}
    if (conn){conn.close();}
    _callback(err, results);});};
    
function report_list(nstore, params, flash){
  get_dbs_reports(nstore, {database:params.form.alias}, function(err, reports){
    if (err !== null){
      report_render({res:params.res, storage:nstore.storage(), 
        data:{subtitle:nstore.lang().label_reports, form:{}, flash:err}});}
    else if (params.form.engine) {
      tool.getReportFiles(nstore, {filters:params.form, dbs_reports:reports.report}, function(err, files){
        if (err !== null && flash === null){
          flash =err;}
        else {
          params.form.files = files;}
        report_render({res:params.res, storage:nstore.storage(), 
          data:{subtitle:nstore.lang().label_reports, form:params.form, flash:flash}});});}
    else {
      report_render({res:params.res, storage:nstore.storage(), 
        data:{subtitle:nstore.lang().label_reports, form:params.form, flash:flash}});}});}

function report_install(nstore, params){
  tool.installReport(nstore,{filename:params.form.update_reportkey}, 
    function(err, report_id, reportkey){
      report_list(nstore, params, err);});}

function report_delete(nstore, params, _callback) {
  if (typeof params.reportkey === "undefined"){
    params.reportkey = "";}
  var conn; var trans;
  async.waterfall([
    function(callback) {
      conn = nstore.connect.getConnect();
      if (!conn){
        callback(nstore.lang().not_connect);}
      else {
        trans = connect.beginTransaction({connection:conn, engine:nstore.engine()});
        var _sql = {select:["*"], from:"ui_report", where:["reportkey","=","?"]}
        trans.query(models.getSql(nstore.engine(), _sql), [params.reportkey], 
          function (err, data) {
            if (err) {callback(err);}
            else {
              if (data.rowCount > 0){
                callback(null, data.rows[0].id);}
              else {
                callback(nstore.lang().missing_reportkey);}}});}},
    
    function(report_id, callback) {
      var _sql = {delete:"", from:"ui_reportfields", 
        where:["report_id","=",report_id]}
      trans.query(models.getSql(nstore.engine(), _sql), [],
        function (err, data) {
          if (err) {callback(err);}
          else {callback(null, report_id);}});},
    
    function(report_id, callback) {
      var _sql = {delete:"", from:"ui_reportsources", 
        where:["report_id","=",report_id]}
      trans.query(models.getSql(nstore.engine(), _sql), [],
        function (err, data) {
          if (err) {callback(err);}
          else {callback(null, report_id);}});},
          
    function(report_id, callback) {
      var _sql = {delete:"", from:"ui_message", 
        where:["secname","like","?"]}
      trans.query(models.getSql(nstore.engine(), _sql), [params.reportkey+"%"],
        function (err, data) {
          if (err) {callback(err);}
          else {callback(null, report_id);}});},
    
    function(report_id, callback) {
      var _sql = {delete:"", from:"ui_report", where:["id","=",report_id]}
      trans.query(models.getSql(nstore.engine(), _sql), [],
        function (err, data) {
          if (err) {callback(err);}
          else {callback(null);}});}
  ],
  function(err) {
    if(err){if(err.message){err = err.message;}}
    if (!err && trans){
      if(trans.commit){
        trans.commit(function (err) {
        if (!err){conn.close();}
        _callback(err);});}
      else {
        conn.close();
        _callback(null);}}
    else {
      if (trans){if (trans.rollback){trans.rollback();}}
      if (conn){conn.close();}
      _callback(err);}});}

function report_update(nstore, params, _callback) {
  var conn; var trans;
  async.waterfall([
    function(callback) {
      conn = nstore.connect.getConnect();
      if (!conn){
        callback(nstore.lang().not_connect);}
      else {
        trans = connect.beginTransaction({connection:conn, engine:nstore.engine()});
        if(params.template){
          var _sql = {update:"ui_report", set:[[],["report","=","?"]], where:["id","=","?"]}
          trans.query(models.getSql(nstore.engine(), _sql), [params.template, params.report_id], 
            function (err, data) {
              callback(err);});}
        else {
          callback(null);}}},
    
    function(callback) {
      var labels = [];
      for (var key in params) {
        if (key.split("_")[0] === "label" && key.split("_").length > 1) {
          labels.push({id:key.split("_")[1], value:params[key]});}}
      if(labels.length > 0){
        var label_lst = [];
        var sql = models.getSql(nstore.engine(), 
          {update:"ui_message", set:[[],["msg","=","?"]], where:["id","=","?"]});
        labels.forEach(function(label) {
          label_lst.push(function(callback_){
            trans.query(sql, [label.value, label.id], function (err, data) {
              callback_(err,data);});});});
        async.series(label_lst,function(err,data) {
          callback(err);;});}
      else {
        callback(null);}}
  ],
  function(err) {
    if(err){if(err.message){err = err.message;}}
    if (!err && trans){
      if(trans.commit){
        trans.commit(function (err) {
        if (!err){conn.close();}
        _callback(err);});}
      else{
        conn.close();
        _callback(null);}}
    else {
      if (trans){if (trans.rollback){trans.rollback();}}
      if (conn){conn.close();}
      _callback(err);}});}
      
function report_edit(nstore, params){
  get_dbs_reports(nstore, {database:params.form.alias, reportkey:params.form.update_reportkey}, 
    function(err, report){
      if (err !== null){
        report_render({res:params.res, storage:nstore.storage(), 
          data:{subtitle:nstore.lang().label_reports, form:params.form, flash:err}});}
      else {
        params.form.report = report.report;
        params.form.labels = report.label;
        report_render({res:params.res, storage:nstore.storage(), 
          data:{view:"edit", subtitle:nstore.lang().label_reports, form:params.form, flash:null}});}});}

function setting_list(params){
  params.req.app.settings.storage.getSettings(function(err, settings){
    if (err) {return params.next(err);}
    else {
      params.page = "setting";
      params.data.settings = settings;
      params.data.subtitle = params.req.app.locals.lang.label_settings;
      page_render(params);}});}

function create_demo(nstore, params, _callback) {
  var data = JSON.parse(JSON.stringify(require('./demo.js')));
  var ndi = require('./ndi.js')(nstore.lang());
  var out = require('./tools.js').DataOutput();
  var tool = require('./tools.js').NervaTools();
  
  async.waterfall([
    function(callback) {
      var results=[];
      results.push({start_process:true, database:params.database, 
        stamp:out.getISODateTime(new Date(),true)});
      var conn = nstore.connect.getConnect();
      if(!conn){
        nstore.connect.getLogin(params, function(err, validator){
        params.validator = validator; params.log_enabled = false;
        params.insert_row = true; params.insert_field = true;
        callback(err, results, params);}); }
      else {
        callback(null, results, params); }
    },
    
    function(results, params, callback) {
      //create 3 departments
      params.datatype = "groups";
      ndi.updateData(nstore, params, data.groups, function(err, result){
        if (!err){
          results.push({start_group:true, end_group:true, 
            datatype:params.datatype, result:result.data});}
        callback(err, results, params);});},
    
    function(results, params, callback) {
      //customer
      //-> def. 4 customer additional data (float,date,valuelist,customer types), 
      //-> create 3 customers, 
      //-> and more create and link to contacts, addresses and events
      async.series([
        function(callback_){
          params.datatype = "deffield";
          ndi.updateData(nstore, params, data.customer.deffield, function(err, result){
          if (!err){
            results.push({start_group:true, section:"customer", datatype:params.datatype, 
              result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "customer";
          ndi.updateData(nstore, params, data.customer.customer, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "address";
          ndi.updateData(nstore, params, data.customer.address, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "contact";
          ndi.updateData(nstore, params, data.customer.contact, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
            
        function(callback_){
          params.datatype = "event";
          ndi.updateData(nstore, params, data.customer.event, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
                          
    function(results, params, callback) {
      //employee
      //-> def. 1 employee additional data (integer type), 
      //->create 1 employee, 
      //->and more create and link to contact, address and event
      async.series([
        function(callback_){
          params.datatype = "deffield";
          ndi.updateData(nstore, params, data.employee.deffield, function(err, result){
          if (!err){
            results.push({start_group:true, section:"employee", datatype:params.datatype, 
              result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "employee";
          ndi.updateData(nstore, params, data.employee.employee, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "address";
          ndi.updateData(nstore, params, data.employee.address, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "contact";
          ndi.updateData(nstore, params, data.employee.contact, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
            
        function(callback_){
          params.datatype = "event";
          ndi.updateData(nstore, params, data.employee.event, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
    
    function(results, params, callback) {
      //product
      //-> def. 3 product additional data (product,integer and valulist types),
      //->create 13 products,
      //->and more create and link to barcodes, events, prices, additional data
      async.series([
        function(callback_){
          params.datatype = "deffield";
          ndi.updateData(nstore, params, data.product.deffield, function(err, result){
          if (!err){
            results.push({start_group:true, section:"product", datatype:params.datatype, 
              result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "product";
          ndi.updateData(nstore, params, data.product.product, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "barcode";
          ndi.updateData(nstore, params, data.product.barcode, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "price";
          ndi.updateData(nstore, params, data.product.price, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
            
        function(callback_){
          params.datatype = "event";
          ndi.updateData(nstore, params, data.product.event, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
    
    function(results, params, callback) {
      //project
      //-> def. 2 project additional data, 
      //->create 1 project, 
      //->and more create and link to contact, address and event
      async.series([
        function(callback_){
          params.datatype = "deffield";
          ndi.updateData(nstore, params, data.project.deffield, function(err, result){
          if (!err){
            results.push({start_group:true, section:"project", datatype:params.datatype, 
              result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "project";
          ndi.updateData(nstore, params, data.project.project, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "address";
          ndi.updateData(nstore, params, data.project.address, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "contact";
          ndi.updateData(nstore, params, data.project.contact, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
            
        function(callback_){
          params.datatype = "event";
          ndi.updateData(nstore, params, data.project.event, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
    
    function(results, params, callback) {
      //tool
      //-> def. 2 tool additional data,
      //->create 3 tools,
      //->and more create and link to event and additional data
      async.series([
        function(callback_){
          params.datatype = "deffield";
          ndi.updateData(nstore, params, data.tool.deffield, function(err, result){
          if (!err){
            results.push({start_group:true, section:"tool", datatype:params.datatype, 
              result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "tool";
          ndi.updateData(nstore, params, data.tool.tool, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "event";
          ndi.updateData(nstore, params, data.tool.event, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
  
    function(results, params, callback) {
      //create +1 warehouse
      params.datatype = "place";
      ndi.updateData(nstore, params, data.place, function(err, result){
        if (!err){
          results.push({start_group:true, end_group:true, 
            datatype:params.datatype, result:result.data});}
        callback(err, results, params);});},
    
    function(results, params, callback) {
      //documents
      //offer, order, invoice, worksheet, rent
      async.series([
        function(callback_){
          params.datatype = "trans";
          ndi.updateData(nstore, params, data.trans_item.trans, function(err, result){
          if (!err){
            results.push({start_group:true, 
              section:"document(offer,order,invoice,rent,worksheet)", 
              datatype:params.datatype, result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "item";
          ndi.updateData(nstore, params, data.trans_item.item, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "link";
          ndi.updateData(nstore, params, data.trans_item.link, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
    
    function(results, params, callback) {
      //payments
      //bank and petty cash
      async.series([
        function(callback_){
          params.datatype = "trans";
          ndi.updateData(nstore, params, data.trans_payment.trans, function(err, result){
          if (!err){
            results.push({start_group:true, 
              section:"payment(bank,petty cash)", 
              datatype:params.datatype, result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "payment";
          ndi.updateData(nstore, params, data.trans_payment.payment, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "link";
          ndi.updateData(nstore, params, data.trans_payment.link, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
  
    function(results, params, callback) {
      //stock control
      //tool movement (for employee)
      //create delivery,stock transfer,correction
      //formula and production
      async.series([
        function(callback_){
          params.datatype = "trans";
          ndi.updateData(nstore, params, data.trans_movement.trans, function(err, result){
          if (!err){
            results.push({start_group:true, 
              section:"stock control(tool movement,delivery,stock transfer,correction,formula,production)", 
              datatype:params.datatype, result:result.data});}
          callback_(err);});},
          
        function(callback_){
          params.datatype = "movement";
          ndi.updateData(nstore, params, data.trans_movement.movement, function(err, result){
          if (!err){
            results.push({datatype:params.datatype, result:result.data});}
          callback_(err);});},
        
        function(callback_){
          params.datatype = "link";
          ndi.updateData(nstore, params, data.trans_movement.link, function(err, result){
          if (!err){
            results.push({end_group:true, datatype:params.datatype, result:result.data});}
          callback_(err);});}  
      ],function(err) {
        callback(err, results, params);});},
  
   function(results, params, callback) {
     //load general reports and other templates
     tool.getReportFiles(nstore, {}, function(err, files){
      if (err){
        callback(err, results, params);}
      else {
        if (files.length>0){
          var report_keys = {}; var report_names = {}; var lrows = [];
          files.forEach(function(report) {
            report_names[report.reportkey] = report.repname;
            report_keys[report.reportkey] = function(callback_){
              tool.installReport(nstore, {conn:params.validator.conn, filename:report.reportkey}, 
              function(err, report_id){
                if(err){
                  lrows.push({datatype:report.reportkey, result:err});}
                else {
                  lrows.push({datatype:report.reportkey, result:report_names[report.reportkey]});}
                callback_(null,report_id);});}});
          async.series(report_keys,function(err, data) {
            if(!err && lrows.length>0){
              lrows[0].start_group = true;
              lrows[0].section = "report template";
              lrows[lrows.length-1].end_group = true;
              lrows.forEach(function(lrow) {
                results.push(lrow);});}
            callback(err, results, params);});}
        else {
          callback(err, results, params);}}});},
    
    function(results, params, callback) {
      //sample menus and menufields
      var menu_lst = {
        groups: function(callback_){
          nstore.valid.getGroupsId({conn:params.validator.conn, groupname:"fieldtype"},
          function(err,groups){
            callback_(err,groups);});}}
      data.menu.ui_menu.forEach(function(menu) {
        menu_lst[menu.menukey] = function(callback_){
          var update_params = {nervatype:"ui_menu", values:menu, 
            validate:true, insert_row:true, validator:params.validator};
          nstore.connect.updateData(update_params, function(err, record_id){
            callback_(err,record_id);});}});
      async.series(menu_lst,function(err, items) {
        if (!err){
          var menufield_lst = [];
          data.menu.ui_menufields.forEach(function(menufield) {
            menufield.menu_id = items[menufield.menu_id];
            menufield.fieldtype = items.groups.fieldtype[menufield.fieldtype];
            menufield_lst.push(function(callback_){
              var update_params = {nervatype:"ui_menufields", values:menufield, 
                validate:true, insert_row:true, validator:params.validator};
              nstore.connect.updateData(update_params, function(err, record_id){
                callback_(err,record_id);});});});
          async.series(menufield_lst,function(err, items) {
            var menu_keys = Object.keys(menu_lst);
            menu_keys.shift();
            var lrow = {datatype:"ui_menu", section:"sample menu",
              start_group:true, end_group:true, result:menu_keys};
            if (err){lrow.result = err;}
            results.push(lrow);
            callback(null, results, params);});}
        else {
          results.push({datatype:"ui_menu", section:"sample menu",
            start_group:true, end_group:true, result:err});
          callback(null, results, params);}});}
              
  ],
  function(err, results, params) {
    if (params.hasOwnProperty("validator")){
      if (params.validator.conn !== null){
        params.validator.conn.close();}}
    if(err){
      results.push({error:err, stamp:out.getISODateTime(new Date(),true)});}
    else {
      results.push({end_process:true, stamp:out.getISODateTime(new Date(),true)});}
    _callback(results);});
}

function get_api(nstore, params, _callback){
  var lang = nstore.lang();

  if(nstore.app_config().check_admin) {
    if(typeof params.code === "undefined"){
      return _callback({type:"error", ekey:"err", err_msg: lang.missing_params, data: "code"}, params); }
    if(params.code !== nstore.app_config().admin_key){
      return _callback({type:"error", ekey:"err", err_msg: lang.wrong_admin_key, data: params.code}, params); }}

  switch (params.method) {
    case "database/create":
      create_database(nstore, {database: params.alias, logtype: "json"}, 
        function(err, logstr){;
          if(err){
            return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
          else{
            if(params.output === "json"){
              return _callback({type: "json", id: params.method, data: logstr}, params); }
            else {
              return _callback({type: "html", tempfile: "data.html", data: {data: logstr, cross_tab: true}}, params); }}});
      break;
    
    case "database/demo":
      if((params.password === "empty") || (params.password === "none")){
        params.password = "";}
      create_demo(nstore, params, function(results){
        if(params.output === "json"){
          return _callback({type: "json", id: params.method, data: results}, params); }
        else {
          return _callback({type: "html", tempfile: "demo.html", data: {data: results}}, params); }});
      break;
    
    case "report/list":
      if(!params.alias){
        return _callback({type:"error", ekey:"err", err_msg: lang.missing_params, data: "alias"}, params); }
      else {
        nstore.local.setEngine({database: params.alias}, function(err, result){
          if(err){
            return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
          else{
            get_dbs_reports(nstore, {database: params.alias}, 
              function(err, reports){;
                if(err){
                  return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
                else{
                  tool.getReportFiles(nstore, 
                    {filters:{ alias:params.alias}, dbs_reports:reports.report}, 
                    function(err, files){
                      if(err){
                        return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
                      else {
                        if(params.output === "csv"){
                          json2csv({data: files}, function(err, csv) {
                            if (err){
                              return _callback({type: "error", ekey: "message", err_msg: err}, params); }
                            else {
                              return _callback({type: "csv", data: csv, filename: params.method}, params); }});}
                        else if(params.output === "json") {
                          return _callback({type: "json", id: params.method, data: files}, params); }
                        else {
                          return _callback({type: "html", 
                            tempfile: "data.html", data: {data: files, cross_tab: true}}, params); }}})}});}})}
      break;
    
    case "report/delete":
      if(!params.alias || !params.reportkey){
        return _callback({type:"error", ekey:"err", err_msg: lang.missing_params, data: "alias, reportkey"}, params); }
      else {
        nstore.local.setEngine({database: params.alias}, function(err, result){
          if(err){
            return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
          else{
            report_delete(nstore, {reportkey: params.reportkey}, function(err){
              if(err){
                return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
              else {
                return _callback({type: "json", id: params.method, 
                  data: {state: lang.data_deleted, reportkey: params.reportkey }}, params); }});}});}
      break;

    case "report/install":
      if(!params.alias || !params.reportkey){
        return _callback({type:"error", ekey:"err", err_msg: lang.missing_params, data: "alias, reportkey"}, params); }
      else {
        nstore.local.setEngine({database: params.alias}, function(err, result){
          if(err){
            return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
          else{
            tool.installReport(nstore,{filename: params.reportkey}, function(err, report_id, reportkey){
              if(err){
                return _callback({type:"error", ekey:"err", err_msg: lang.log_error, data: err}, params); }
              else {
                return _callback({type: "json", id: params.method, 
                  data: {state: lang.update_ok, report_id: report_id }}, params); }});}});}
      break;
  
    default:
      return _callback({type:"error", ekey:"invalid", err_msg: lang.unknown_method, data: params.method}, params);
  }
}
                                  
return {
  getApi: function(nstore, params, _callback){
    return get_api(nstore, params, _callback);},
  validUser: function(req) {
    return valid(req);},
  validSetting: function(conf, params, lang) {
    return valid_setting(conf, params, lang);},
  getNstoreParams(req){
    return get_nstore_prm(req);},
  pageRender: function(params) {
    page_render(params);},
    
  userList: function(params) {
    user_list(params);},
  databaseList: function(params) {
    database_list(params);},
  importList: function(params) {
    import_list(params);},
  settingList: function(params) {
    setting_list(params);},
      
  createDatabase: function(nstore, params, _callback) {
    create_database(nstore, params, _callback);},
  exportDatabase: function(nstore, params, _callback) {
    export_database(nstore, params, _callback);},
  importDatabase: function(nstore, params, _callback) {
    import_database(nstore, params, _callback);},
  createDemo: function(nstore, params, _callback) {
    create_demo(nstore, params, _callback);},
  
  reportRender: function(params) {
    report_render(params);},
  reportList: function(nstore, params, flash) {
    report_list(nstore, params, flash);},
  reportInstall: function(nstore, params) {
    report_install(nstore, params);},
  reportUpdate: function(nstore, params, _callback) {
    report_update(nstore, params, _callback);},
  reportDelete: function(nstore, params, _callback) {
    report_delete(nstore, params, _callback);},
  reportEdit: function(nstore, params) {
    report_edit(nstore, params);},
  getDbsReports: function(nstore, params, _callback) {
    get_dbs_reports(nstore, params, _callback);}
  };
};