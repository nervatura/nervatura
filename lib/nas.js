/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2019, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var async = require("async");
var json2csv = require('json2csv');

var connect = require('./adapter.js').connect();
var models = require('./adapter.js').models();
var out = require('./tools.js').DataOutput();
var tool = require('./tools.js').NervaTools();

module.exports = function() {

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
          trans.commit(function (cerr) {
            callback(cerr||null);});}
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

function report_install(nstore, params, _callback){
  tool.installReport(nstore,{filename:params.reportkey}, 
    function(err, report_id, reportkey){
      _callback(err, report_id, reportkey);});}

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
        trans.commit(function (cerr) {
        conn.close();
        _callback(cerr||null);});}
      else {
        conn.close();
        _callback(null);}}
    else {
      if (trans){if (trans.rollback){trans.rollback();}}
      if (conn){conn.close();}
      _callback(err);}});}

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
      
  createDatabase: function(nstore, params, _callback) {
    create_database(nstore, params, _callback);},
  createDemo: function(nstore, params, _callback) {
    create_demo(nstore, params, _callback);},
  
  reportInstall: function(nstore, params, _callback) {
    report_install(nstore, params, _callback);},
  reportDelete: function(nstore, params, _callback) {
    report_delete(nstore, params, _callback);},
  getDbsReports: function(nstore, params, _callback) {
    get_dbs_reports(nstore, params, _callback);}
  };
};