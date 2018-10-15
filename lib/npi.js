/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2018, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var async = require("async");

var connect = require('./adapter.js').connect();
var models = require('./adapter.js').models();

module.exports = function(lang) {

function get_login(nstore, params, _callback) {
  async.waterfall([
    function(callback) {
      nstore.connect.getLogin({database:params.database, username:params.username, 
        password:params.password, token:params.token},
        function(err, validator){
          if (err){
            callback(err, validator);}
          else {
            callback(null, validator);}});},
    
    function(validator, callback) {
      validator.engine = nstore.engine();
      if(validator.employee){
        var _sql = {select:["*"], from:"ui_audit",where:["usergroup","=","?"]}
        validator.conn.query(models.getSql(nstore.engine(), _sql), [validator.employee.usergroup], 
          function (err, results) {
            if (err) {
              validator.valid = false; callback(err, validator);}
            else {
              validator.audit = results.rows; callback(null, validator); }});}
      else {
        callback(null, validator); }},
      
    function(validator, callback) {
      if(validator.employee){
        var _sql = {select:["ref_id_2"], from:"link",
          where:[["ref_id_1","=","?"],["and","link.deleted","=","0"],
            ["and","nervatype_1","in",
              [{select:["id"], from:"groups",
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'groups'"]]}]],
            ["and","nervatype_2","in",
              [{select:["id"], from:"groups",
                where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'groups'"]]}]]]}
        validator.conn.query(models.getSql(nstore.engine(), _sql), [validator.employee.usergroup], 
          function (err, results) {
            if (err) {
              validator.valid = false; callback(err, validator);}
            else {
              if (results.rowCount>0) {
                validator.transfilter = results.rows[0].ref_id_2;}
              else {validator.transfilter = null;}
              callback(null, validator); }});}
      else {
        callback(null, validator); }},
            
    function(validator, callback) {
      var _sql = {select:["*"], from:"groups",
        where:["groupname","in",[[],"'usergroup'","'nervatype'","'transtype'","'inputfilter'",
          "'transfilter'","'department'","'logstate'","'fieldtype'"]]}
      validator.conn.query(models.getSql(nstore.engine(), _sql), [], 
        function (err, results) {
          if (err) {
            validator.valid = false; callback(err, validator);}
          else {
            validator.groups = results.rows; callback(null, validator);}});},
      
    function(validator, callback) {
      var _sql = {select:["*"], from:"ui_menu"}
      validator.conn.query(models.getSql(nstore.engine(), _sql), [], 
        function (err, results) {
          if (err) {
            validator.valid = false; callback(err, validator);}
          else {
            validator.menucmd = results.rows; callback(null, validator);}});},
    
    function(validator, callback) {
      var _sql = {select:["*"], from:"ui_menufields", order_by:["menu_id","orderby"]}
      validator.conn.query(models.getSql(nstore.engine(), _sql), [], 
        function (err, results) {
          if (err) {
            validator.valid = false; callback(err, validator);}
          else {
            validator.menufields = results.rows; callback(null, validator);}});},
    
    function(validator, callback) {
      var _sql = {select:["value"], from:"fieldvalue",
        where:[["ref_id","is","null"],["and","fieldname","=","'log_login'"]]}
      validator.conn.query(models.getSql(nstore.engine(), _sql), [], 
        function (err, results) {
          if (err) {
            validator.valid = false;}
          else {
            if (results.rowCount>0) {
              validator.userlogin = results.rows[0].value;}
            else {validator.userlogin = "false";}}
          callback(err, validator);});}
    ],
  function(err, validator) {
    if(err){if(err.message){err = err.message;}
      validator.message = err;}
    if (validator.conn !== null) {
      validator.conn.close(); validator.conn = null;}
    _callback(validator);});}

function set_data(nstore, login, params, _callback) {
  async.waterfall([
    function(callback) {
      var validator = login;
      if (!validator.conn){
        nstore.connect.getLogin({database:login.database, username:login.username, 
          password:login.password, token:login.token},
          function(err, data){
            if (err){callback(err, data);}
            else {callback(null, data);}});}
      else {
        callback(null, validator);}},
    
    function(validator, callback) {
      switch (params.method) {
        case "function":
          var ntool = require('./tools.js').NervaTools();
          if (typeof ntool[params.functionName] !== "undefined") {
            params.paramList.conn = validator.conn;
            ntool[params.functionName](nstore, params.paramList, function(err, data){
              callback(err, validator, data);});}
          else {callback(lang.unknown_method+" "+params.functionName, validator);}
          break;
        case "update":
          var update_params = {nervatype:params.record.__tablename__, values:params.record, 
            validate:params.validate, insert_row:true, transaction:params.transaction,
            validator:validator};
          nstore.connect.updateData(update_params, function(err, record_id){
            if (err) {
              callback(err, validator);}
            else {
              if (params.record.id === null){
                params.record.id = record_id;}
              callback(null, validator, params.record);}});
          break;
        case "delete":
          var delete_params = {nervatype:params.record.__tablename__, ref_id:params.record.id, 
            transaction:params.transaction, validator:validator};
          nstore.connect.deleteData(delete_params, function(err, id){
            if (err) {
              callback(err, validator);}
            else {
              callback(null, validator, id);}});
          break;
        default:
          var sql = "";
          switch (params.method) {
            case "table":
              if(params.classAlias){
                sql = "select * from "+params.classAlias;}
              else {
                sql = "select * from "+params.tableName;}
              if (params.filterStr!=="" && params.filterStr!==null) {
                sql += " where "+params.filterStr;}
              if (params.orderStr!=="" && params.orderStr!==null) {
                sql += " order by "+params.orderStr;}
              break;
            case "view":
              if (typeof params.orderStr === "undefined") {params.orderStr = "";}
              sql = nstore.local.setSqlParams(params);
              break;
            case "execute":
              sql = nstore.local.setSqlParams({sqlStr:params.sqlStr, whereStr:"", havingStr:"", 
                paramList:params.paramList, orderbyStr:"", rlimit:false});
              break;}
          validator.conn.query(sql, [], function (error, data) {
            if (error) {
              callback(error, validator);}
            else {
              callback(null, validator, data.rows);}});}}
    ],
  function(err, validator, results) {
    if(err){if(err.message){err = err.message;}}
    if (!login.conn && validator){
      if(validator.conn){
        validator.conn.close();}}
    _callback(err, results || validator, params.infoName);});};

function load_dataset(nstore, params, _callback) {
  async.waterfall([    
    function(callback) {
      nstore.connect.getLogin({database:params.login.database, username:params.login.username, 
        password:params.login.password, token:params.login.token}, function(err, validator){
          if (err){
            callback(err, validator);}
          else {callback(null, validator);}});},
      
    function(validator, callback) {
      var results = []; var results_lst = [];
      params.dataSetInfo.forEach(function(recordSetInfo) {
        recordSetInfo.method = recordSetInfo.infoType;
        results_lst.push(function(callback_){
          set_data(nstore, validator, recordSetInfo, function(err, data, info){
            if (err){
              results.push({"infoName":info, "recordSet":err});}
            else {
              results.push({"infoName":info, "recordSet":data});}
            callback_(err);});});});
      async.series(results_lst,function(err) {
        callback(null, validator, results);});
            }
    ],
  function(err, validator, results) {
    if(err){if(err.message){err = err.message;}}
    if (validator.conn !== null){
      validator.conn.close();}
    _callback(err, results || validator);});};

function update_recordset(nstore, params, _callback) {
  async.waterfall([
    function(callback) {
      nstore.connect.getLogin({database:params.login.database, username:params.login.username, 
        password:params.login.password, token:params.login.token}, function(err, validator){
          if (err){
            callback(err, validator);}
          else {callback(null, validator);}});},
        
    function(validator, callback) {
      var results = [];
      var trans = connect.beginTransaction({connection:validator.conn, engine:nstore.engine()});
      var results_lst = [];
      params.recordSet.forEach(function(record) {
        var record_params = {method:params.method, record: record, transaction: trans};
        results_lst.push(function(callback_){
          set_data(nstore, validator, record_params, function(err,data){
            if (err){results.push(err);}
            else {results.push(data);}
            callback_(err);});});});
      async.series(results_lst,function(err) {
        
        if(err){
          if(trans.rollback){
            trans.rollback();}
          callback(err, validator, results);}
        else if(results.length>0){
          if(trans.commit){
            trans.commit(
              function(error){
                callback(error, validator, results);});}
          else{
            callback(err, validator, results);}}
        else {
          callback(err, validator, results);}});}
    ],
  function(err, validator, results) {
    if(err){if(err.message){err = err.message;}}
    if (validator.conn !== null){
      validator.conn.close();}
    _callback(err, results || validator);});
};

function save_dataset(nstore, params, _callback) {
  async.waterfall([
    function(callback) {
      nstore.connect.getLogin({database:params.login.database, username:params.login.username, 
        password:params.login.password, token:params.login.token}, function(err, validator){
          if (err){
            callback(err, validator);}
          else {callback(null, validator);}});},
    
    function(validator, callback) {
      var results = [];
      var trans = connect.beginTransaction({connection:validator.conn, engine:nstore.engine()});
      var results_lst = [];
      params.dataSet.forEach(function(updateSetInfo) {
        switch (updateSetInfo.updateType) {
          case "update":
            updateSetInfo.recordSet.forEach(function(record) {
              var record_params = {method:"update", record: record, transaction: trans};
              results_lst.push(function(callback_){
                set_data(nstore, validator, record_params, function(err,data){
                  if (err){results.push(err);}
                  else {results.push(data);}
                  callback_(err);});});});
            break;
          case "delete":
            updateSetInfo.recordSet.forEach(function(record) {
              var record_params = {method:"delete", record: record, transaction: trans};
              results_lst.push(function(callback_){
                set_data(nstore, validator, record_params, function(err,data){
                  if (err){results.push(err);}
                  else {results.push(data);}
                  callback_(err);});});});
            break;
          case "function":
            var function_params = {method:"function", functionName:updateSetInfo.functionName, 
              paramList:updateSetInfo.paramList}
            results_lst.push(function(callback_){
              set_data(nstore, validator, function_params, function(err, result){
                if(err) {
                  updateSetInfo.value = err;}
                else {
                  updateSetInfo.value = result;}
                callback_(err);});});
            break;}});
      async.series(results_lst,function(err) {
        if(err){
          if(trans.rollback){
            trans.rollback();}
          callback(err, validator, results);}
        else if(results.length>0){
          if(trans.commit){
            trans.commit(
              function(error){
                callback(error, validator, results);});}
          else{
            callback(err, validator, results);}}
        else {
          callback(err, validator, results);}});}
  ],
  function(err, validator, results) {
    if(err){if(err.message){err = err.message;}}
    if (validator.conn !== null){
      validator.conn.close();}
    _callback(err, results || validator);});};

function get_api(nstore, params, _callback){
  
  var _params = params.params;
  switch (params.method) {
    case "getLogin":
    case "getLogin_json":
      get_login(nstore, _params, function(result){
        return _callback({type: "json", id: 1, data: result}, params); });
      break;
    case "loadView":
    case "loadView_json":
      _params.method = "view";
      set_data(nstore, _params.login, _params, function(err, result){
        if(err){ 
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 2, data: result}, params); }});
      break;
    case "loadTable":
    case "loadTable_json":
      _params.method = "table";
      set_data(nstore, _params.login, _params, function(err, result){
        if(err){
           return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 3, data: result}, params); }});
      break;
    case "loadDataSet":
    case "loadDataSet_json":
      load_dataset(nstore, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 4, data: result}, params); }});
      break;
    case "executeSql":
    case "executeSql_json":
      _params.method = "execute";
      set_data(nstore, _params.login, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 5, data: result}, params); }});
      break;
    case "saveRecord":
    case "saveRecord_json":
      _params.method = "update";
      set_data(nstore, _params.login, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 6, data: result}, params); }});
      break;
    case "saveRecordSet":
    case "saveRecordSet_json":
      _params.method = "update";
      update_recordset(nstore, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 7, data: result}, params); }});
      break;
    case "saveDataSet":
    case "saveDataSet_json":
      save_dataset(nstore, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 8, data: result}, params); }});
      break;
    
    case "deleteRecord":
    case "deleteRecord_json":
      _params.method = "delete";
      set_data(nstore, _params.login, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 9, data: result}, params); }});
      break;
    case "deleteRecordSet":
    case "deleteRecordSet_json":
      _params.method = "delete";
      update_recordset(nstore, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 10, data: result}, params); }});
      break;        
    case "callFunction":
    case "callFunction_json":
      _params.method = "function";
      set_data(nstore, _params.login, _params, function(err, result){
        if(err){
          return _callback({type:"error", id:-1, ekey:null, err_msg: lang.log_error, data: err}, params);}
        else {
          return _callback({type: "json", id: 11, data: result}, params); }});
      break;
        
    default:
      return _callback({type:"error", id:-1, ekey:"invalid", err_msg: lang.unknown_method, data: params.method}, params); }
}

return {
  getApi: function(nstore, params, _callback){
    return get_api(nstore, params, _callback);},
  getLogin: function(nstore, params, _callback) {
    get_login(nstore, params, _callback);},
  setData: function(nstore, method, params, _callback){
    params.method = method;
    set_data(nstore, params.login, params, 
    function(err, results, info){
      _callback(err, results);});},
  loadDataSet: function(nstore, params, _callback){
    load_dataset(nstore, params, function(err, results){
      _callback(err, results);});},
  updateRecordSet: function(nstore, method, params, _callback){
    params.method = method;
    update_recordset(nstore, params, function(err, results){
      _callback(err, results);});},
  saveDataSet: function(nstore, params, _callback){
    save_dataset(nstore, params, function(err, results){
      _callback(err, results);});}
  };
};
