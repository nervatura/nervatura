/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var express = require('express');
var router = express.Router();

var async = require("async");

router.use(function (req, res, next) {
  next()
});

router.get('/create', function(req, res, next) {
  var params = {database:req.query.database, 
    username:req.query.username, password:req.query.password}
  create_demo(params, req, res, function(results){
    res.render('ndi/demo.html',{data:results});});});

router.post('/create', function(req, res, next) {
  req.setTimeout(req.app.settings.conf.long_timeout);
  var params = {database:req.body.database, 
    username:req.body.username, password:req.body.password}
  create_demo(params, req, res, function(results){
    res.set('Content-Type', 'text/json');
    res.send({"id":null, "jsonrpc": "2.0", "result":results});});});

function create_demo(params, req, res, _callback) {
  var nstore = require('../lib/node/nervastore.js')(req, res);
  var data = JSON.parse(JSON.stringify(require('../lib/node/demo.js')));
  var ndi = require('../lib/node/ndi.js')();
  var out = require('../lib/node/tools.js').DataOutput();
  var tool = require('../lib/node/tools.js').NervaTools();
  
  async.waterfall([
    function(callback) {
      var results=[];
      results.push({start_process:true, database:params.database, 
        stamp:out.getISODateTime(new Date(),true)});
      nstore.connect.getLogin(params, function(err, validator){
        params.validator = validator; params.log_enabled = false;
        params.insert_row = true; params.insert_field = true;
        callback(err, results, params);});},
    
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

module.exports = router;