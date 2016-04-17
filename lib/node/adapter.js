/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

exports.connect = function() {
  var db = require('any-db');
  var begin = require('any-db-transaction');
  
return {
  getConfig: function(engine, connect, data_dir){
    var config = {adapter:engine}
    for (var key in connect) {
      var kvalue = connect[key].toString().trim();
      if (kvalue.length > 0) {
        if(kvalue[0] === "$" && process.env[kvalue.substr(1)]){
          connect[key] = process.env[kvalue.substr(1)];}}}
    switch (engine) {
        case "sqlite":
          config.adapter = "sqlite3";
          config.database = data_dir+"/database/"+connect.dbname+".db";
          break;
        case "mysql":
        case "postgres":
        case "mssql":
          if (connect.host.indexOf("/")>-1){
            config.host = connect.host.split("/")[0];
            config.instanceName = connect.host.split("/")[1];}
          else {
            config.host = connect.host;}
          if (connect.port!==0 && connect.port!==null && connect.port!==""){
            config.port = connect.port;}
          config.database = connect.dbname;
          if (connect.username!==null && connect.username !=="") {
            config.user = connect.username;
            config.password = connect.password;}
          break;
        default:
          config.adapter = null;}
      //console.log(config);
      return config;},
  
  createConnection: function(params, _callback) {
    try {
      if (params.pool) {
        return db.createPool(params.connect, params.config);}
      else {
        return db.createConnection(params.connect, _callback);}} 
    catch (error) {
      if (_callback){_callback(error)}
      else {return null;}}},
      
  beginTransaction: function(params, _callback) {
    try{
      if(params.engine === "mssql"){
        //begin(params.connection, {autoRollback: false, commit:"COMMIT TRAN", rollback:"ROLLBACK TRAN"}, "BEGIN TRAN", _callback);
        //disabled
        return params.connection;}
      else {
        return begin(params.connection, _callback);}}
    catch (error) {
      if (_callback){_callback(error)}
      else {return null;}}}
}}

exports.models = function() {
  
var ntura = require('./models.js');
var types = {
  id: {
    base: "INTEGER PRIMARY KEY AUTOINCREMENT",
    postgres: "SERIAL PRIMARY KEY",
    mysql: "INT AUTO_INCREMENT NOT NULL, PRIMARY KEY (id)",
    mssql: "INT IDENTITY PRIMARY KEY"},
  integer: {
    base: "INTEGER",
    mysql: "INT",
    mssql: "INT"},
  float: {
    base: "DOUBLE",
    postgres: "FLOAT8",
    mssql: "DECIMAL(19,4)"},
  string: {
    base: "CHAR",
    postgres: "VARCHAR",
    mysql: "VARCHAR",
    mssql: "VARCHAR"},
  password: {
    base: "CHAR",
    postgres: "VARCHAR",
    mysql: "VARCHAR",
    mssql: "VARCHAR"},
  text: {
    base: "TEXT",
    mysql: "LONGTEXT",
    mssql: "VARCHAR(max)"},
  date: {
    base: "DATE"},
  datetime: {
    base: "TIMESTAMP",
    mysql: "DATETIME",
    mssql: "DATETIME2"},
  reference: {
    base: "INTEGER REFERENCES foreign_key ON DELETE ",
    postgres: "INTEGER REFERENCES foreign_key ON DELETE ",
    mysql: "INT, INDEX index_name (field_name), FOREIGN KEY (field_name) REFERENCES foreign_key ON DELETE ",
    mssql: "INT NULL, CONSTRAINT constraint_name FOREIGN KEY (field_name) REFERENCES foreign_key ON DELETE "}}

return {
  getDataType: function(engine, dtype){
    if(types[dtype][engine]){
      return types[dtype][engine];}
    else {
      return types[dtype].base;}},
  
  getLastID: function(engine, table){
    switch (engine) {
      case "postgres":
        return "select currval('"+table+"_id_seq') as id";
      case "mssql":
        return "select ident_current('"+table+"') as id"
      default:
        return null;}},
  
  getSql: function(engine,_sql) {
    
    function engine_type(sql){
      switch (engine) {
        case "sqlite":
          sql = sql.replace(/{CAS_INT}/g, "cast(");//cast as integer - start
          sql = sql.replace(/{CAE_INT}/g, " as integer)");//cast as integer - end
          sql = sql.replace(/{CAS_FLOAT}/g, "cast(");//cast as float - start
          sql = sql.replace(/{CAE_FLOAT}/g, " as double)");//cast as float - end
          sql = sql.replace(/{CAS_DATE}/g, "");//cast as date - start
          sql = sql.replace(/{CAE_DATE}/g, "");//cast as date - end
          sql = sql.replace(/{FMS_DATE}/g, "");//format to iso date - start
          sql = sql.replace(/{FME_DATE}/g, "");//format to iso date - end
          sql = sql.replace(/{FMS_DATETIME}/g, "substr(");//format to iso datetime - start
          sql = sql.replace(/{FME_DATETIME}/g, ",1,16)");//format to iso datetime - end
          break;
        case "mysql":
          sql = sql.replace(/{CAS_INT}/g, "cast(");
          sql = sql.replace(/{CAE_INT}/g, " as signed)");
          sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
          sql = sql.replace(/{CAE_FLOAT}/g, " as decimal)");
          sql = sql.replace(/{CAS_DATE}/g, "cast(");
          sql = sql.replace(/{CAE_DATE}/g, " as date)");
          sql = sql.replace(/{FMS_DATE}/g, "date_format(");
          sql = sql.replace(/{FME_DATE}/g, ", '%Y-%m-%d')");
          sql = sql.replace(/{FMS_DATETIME}/g, "date_format(");
          sql = sql.replace(/{FME_DATETIME}/g, ", '%Y-%m-%d %H:%i')");
          break;
        case "mssql":
          sql = sql.replace(/{CAS_INT}/g, "cast(");
          sql = sql.replace(/{CAE_INT}/g, " as int)");
          sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
          sql = sql.replace(/{CAE_FLOAT}/g, " as float)");
          sql = sql.replace(/{CAS_DATE}/g, "cast(");
          sql = sql.replace(/{CAE_DATE}/g, " as date)");
          sql = sql.replace(/{FMS_DATE}/g, "convert(varchar(10),");
          sql = sql.replace(/{FME_DATE}/g, ", 120)");
          sql = sql.replace(/{FMS_DATETIME}/g, "convert(varchar(16),");
          sql = sql.replace(/{FME_DATETIME}/g, ", 120)");
          break;
        case "postgres":
          sql = sql.replace(/{CAS_INT}/g, "cast(");
          sql = sql.replace(/{CAE_INT}/g, " as integer)");
          sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
          sql = sql.replace(/{CAE_FLOAT}/g, " as float8)");
          sql = sql.replace(/{CAS_DATE}/g, "cast(");
          sql = sql.replace(/{CAE_DATE}/g, " as date)");
          sql = sql.replace(/{FMS_DATE}/g, "to_char(");
          sql = sql.replace(/{FME_DATE}/g, ", 'YYYY-MM-DD')");
          sql = sql.replace(/{FMS_DATETIME}/g, "to_char(");
          sql = sql.replace(/{FME_DATETIME}/g, ", 'YYYY-MM-DD HH24:MI')");
          break;}
      return sql;}
    
    function sql_decode(data,key){
      var sql = "";
      if (Array.isArray(data)){
        var sep=" ", start_br="", end_br="";
        if (data.length>0){
          if ((key === "select")|| (key === "union_select") || (key === "order_by") 
            || (key === "group_by") || (data[0].length === 0)){
              sep = ", ";}}
        data.forEach(function(element) {
          if ((typeof element === "undefined") || (element === null)){element = "null";}
          if(element.length === 0){
            if(key !== "set"){
              start_br = "("; end_br=")";}}
          else if((data.length === 2) && ((element === "and") || (element === "or"))){
            sql += element+" ("; end_br=")";}
          else if(key && (data.length === 1) && (typeof data[0] === "object")){
            sql += " ("+sql_decode(element,key)+")";}
          else{
            sql += sep+sql_decode(element,key);}});
        if(sep === ", "){
          sql = sql.substr(2);}
        if(key && (data.indexOf("on")>-1)){
          sql = key.replace("_"," ")+sql;}
        return start_br+sql.toString().trim()+end_br;}
      else if (typeof data === "object"){
        for (var _key in data) {
          if (data.hasOwnProperty(_key)) {
            if ((_key === "inner_join") || (_key === "left_join")){
              sql += " "+sql_decode(data[_key],_key);}
            else {
              sql += " "+_key.replace("_"," ")+" "+sql_decode(data[_key],_key);}}}
        return sql;}
      else {
        if (data === "?" && engine === "postgres"){
          prm_count += 1; data = "$"+prm_count;}
        return data;}}
    var prm_count = 0;
    return engine_type(sql_decode(_sql));},
  
  exportList: {
    nom: ["groups","numberdef","currency","tax","deffield","pattern","customer","employee",
      "place","product","barcode","price","project","rate","tool","trans","event",
      "item","movement","payment","address","contact","link","log","fieldvalue"],
    ui_1: ["ui_menu","ui_language","ui_message","ui_report"],
    ui_2: ["ui_audit","ui_menufields","ui_reportfields","ui_reportsources","ui_userconfig"]},
      
  dropList:  function(engine) {
    var drop_lst = [
      "pattern","movement","payment","item","trans","barcode","price","tool", "product","tax", "rate",
      "place","currency","project","customer","event","contact","address","numberdef", "log","fieldvalue",
      "deffield","ui_audit","link","ui_userconfig","ui_printqueue","employee", "ui_reportsources",
      "ui_reportfields","ui_report", "ui_message","ui_menufields","ui_menu","ui_language","groups"]
    var rows = [];
    if (engine === "mysql"){
      rows.push("SET FOREIGN_KEY_CHECKS=0;");}
    drop_lst.forEach(function(model) {
      if (engine === "mssql"){
        rows.push("DROP TABLE "+model+";");}
      else {
        rows.push("DROP TABLE IF EXISTS "+model+";");}});
    if (engine === "mysql"){
      rows.push("SET FOREIGN_KEY_CHECKS=1;");}
    return rows;},
    
  modelList:  function(engine) {
    var rows = [];
    for (var table in ntura.model) {
      var cstr = "CREATE TABLE "+table+"("
      for (var fieldname in ntura.model[table]) {
        if (fieldname.substr(0,1) !== "_" ){
          cstr += fieldname;
          var field = ntura.model[table][fieldname];
          if (field.references){
            var reference = types.reference[engine];
            if (!reference){
              reference = types.reference.base;}
            reference = reference.replace("foreign_key",field.references[0]+"(id)");
            reference = reference.replace("field_name",fieldname).replace("field_name",fieldname);
            reference = reference.replace("index_name",fieldname+"__idx");
            reference = reference.replace("constraint_name",table+"__"+fieldname+"__constraint");
            if ((engine === "mssql") && (field.references.length === 3)){
              reference += field.references[2];}
            else {reference += field.references[1];}
            cstr += " "+reference;}
          else {
            if (types[field.type]){
              if (types[field.type][engine]){
                cstr += " "+types[field.type][engine];}
              else {
                cstr += " "+types[field.type]["base"];}}
            if (field.length){
              cstr += "("+field.length+")";}}
          if (field.notnull && !field.references){
            cstr += " NOT NULL";}
          if ((typeof field.default !== "undefined") && (field.default !== "nextnumber")){
            cstr += " DEFAULT "+field.default;}    
          cstr += ", ";}}
      cstr += ");"; cstr = cstr.replace(", );",");");
      rows.push(cstr);}
    return rows;},
    
  indexList:  function(engine) {
    var rows = [];
    for (var ikey in ntura.index) {
      var ifield = ntura.index[ikey];
      var istr = "CREATE INDEX ";
      if (ifield.unique) {
        istr = "CREATE UNIQUE INDEX ";}
      istr += ikey + " ON " + ifield.model + "(";
      ifield.fields.forEach(function(fkey) {
        istr += fkey + ", ";});
      istr += ");"; istr = istr.replace(", );",");");
      rows.push(istr);}
    return rows;},
  
  compact: function(engine) {
    switch (engine) {
      case "sqlite":
        return "vacuum";
      case "postgres":
        return "vacuum";
      default:
        return null;}},
    
  dataList:  function(engine,datarows,ref_id) {
    var rows = [];
    if (!datarows){
      datarows = ntura.data;}
    if (!ref_id){
      ref_id = {};}
    
    function set_id(mname, key_fields, values, id){
      if (!ref_id[mname]){
        ref_id[mname] = {};}
      switch (key_fields.length) {
        case 1:
          ref_id[mname][values[key_fields[0]]] = id;
          break;
        case 2:
          if (!ref_id[mname][values[key_fields[0]]]){
            ref_id[mname][values[key_fields[0]]] = {};}
          ref_id[mname][values[key_fields[0]]][values[key_fields[1]]] = id;
          break;
        case 3:
          if (!ref_id[mname][values[key_fields[0]]]){
            ref_id[mname][values[key_fields[0]]] = {};}
          if (!ref_id[mname][values[key_fields[0]]][values[key_fields[1]]]){
            ref_id[mname][values[key_fields[0]]][values[key_fields[1]]] = {};}
          ref_id[mname][values[key_fields[0]]][values[key_fields[1]]][values[key_fields[2]]] = id;
          break;
        default:
          break;}}
          
    function get_id(values){
      switch (values.length) {
        case 2:
          return ref_id[values[0]][values[1]];
        case 3:
          return ref_id[values[0]][values[1]][values[2]];
        case 4:
          return ref_id[values[0]][values[1]][values[2]][values[3]];
        default:
          return 0;}}
    
    for (var mname in datarows) {
      var key_fields = ntura.model[mname]._key;
      var drows = datarows[mname]; var insert_id = "";
      if(engine === "mssql"){
        insert_id = "SET IDENTITY_INSERT "+mname+" ON; ";}
      for (var i = 0; i < drows.length; i++) {
        set_id(mname, key_fields, drows[i], i+1);
        var fields = "id, "; var values = i+1 + ", ";
        for (var field in drows[i]) {
          fields += field + ", ";
          if (Array.isArray(drows[i][field])){
            values += get_id(drows[i][field]) + ", ";}
          else {
            switch (ntura.model[mname][field].type) {
              case "string":
              case "password":
              case "text":
              case "date":
              case "datetime":
                values += "'"+drows[i][field] + "', ";
                break;
              default:
                values += drows[i][field] + ", ";
                break;}}}
        fields += ") "; fields = fields.replace(", ) ",") ");
        values += ");"; values = values.replace(", );",");");
        rows.push(insert_id+"INSERT INTO "+mname+"("+fields+" VALUES("+values);}}
    
    switch (engine){
      case "postgres":
        //update all sequences
        var tables = Object.keys(ntura.model);
        tables.forEach(function(table) {
          rows.push("SELECT setval('"+table+"_id_seq', (SELECT max(id) FROM "+table+"))");});
        break;}
    return rows;}

}}


  








