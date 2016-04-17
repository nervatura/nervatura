/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var async = require("async");
var express = require("express");
var router = express.Router();

var ntura = require("../lib/node/models.js");
var models = require("../lib/node/adapter.js").models();
  
router.use(function (req, res, next) {
  next()
});

router.get("/", function(req, res, next) {
  index(req, res);});

router.get("/index", function(req, res, next) {
  index(req, res);});

router.post("/index", function(req, res, next) {
  index(req, res, req.body);});

function index(req, res, data) {
  var nstore = require("../lib/node/nervastore.js")(req, res);
  var response = {
  title:"NDI Wizard", 
    data:{
      nom:"address", database:"demo", username:"demo", password:"", code:"", use_deleted:"false",
      output:"html", cross_tab:"false", show_id:"false", no_deffield:"false", insert_row:"true", insert_field:"true",
      update_rows:2, delete_rows:2, selected_tab:"view"},
    url_code: ["base64"],
    view_output: ["html","json","xml","csv"],
    lst_no_deffield: ["groups","fieldvalue","numberdef","deffield","pattern"],
    view_nervatype: {
      address: ["customer","employee","event","place","product","project","tool","trans"],
      contact: ["customer","employee","event","place","product","project","tool","trans"],
      event: ["customer","employee","place","product","project","tool","trans"],
      fieldvalue: ["address","barcode","contact","currency","customer","employee","event","item","link", 
        "log","movement","payment","price","place","product","project","rate","tax",
        "tool","trans","setting"],
      link: ["address","barcode","contact","currency","customer","employee","event","groups","item", 
        "movement","payment","price","place","product","project","rate","tax","tool","trans"]},
    lst_nom: [
    "address","barcode","contact","currency","customer","deffield","employee","event","fieldvalue",
    "groups","item","link","log","movement","numberdef","pattern","payment","place","price",
    "product","project","rate","tax","tool","trans","sql"]};
  if (typeof data !== "undefined"){
    response.data = data;}
  
  async.waterfall([
    function(callback) {
      response.data.fields_lst = create_fieldlist(response.data.nom);
      nstore.connect.getLogin({
        database:response.data.database, username:response.data.username, password:response.data.password}, 
        function(err, validator){
          if (err){
            callback(err, validator);}
          else{
            callback(null, validator);}});},
    
    function(validator, callback) {
      if (["address","barcode","contact","currency","customer","employee","event","groups", 
           "item","link","log","movement","price","place","product","project","rate","tax",
           "tool","trans"].indexOf(response.data.nom) >-1){
        var _sql = {
          select:["df.fieldname as fieldname","ft.groupvalue as fieldtype",
            "df.description as description","df.valuelist as valuelist"],
          from:"deffield df",
          inner_join:[["groups nt","on",[["df.nervatype","=","nt.id"],
            ["and","nt.groupvalue","=","'"+response.data.nom+"'"]]],
            ["groups ft","on",["df.fieldtype","=","ft.id"]]],
          where:[["df.deleted","=","0"],["and","df.visible","=","1"],["and","df.readonly","=","0"]]}
        validator.conn.query(models.getSql(nstore.engine(), _sql), [], 
        function (error, data) {
          if (error) {
            callback(error, validator);}
          else {
            data.rows.forEach(function(field) {
              if (field.valuelist){field.valuelist = field.valuelist.split("|");}
              response.data.fields_lst.push({fieldname:field.fieldname, label:field.description,
                fieldtype:field.fieldtype, mtype:"deffield", fieldcat:1, values:field.valuelist});});
            callback(null, validator);}});}
      else {callback(null, validator);}}
  ],
  function(err, validator) {
    if(err){if(err.message){err = err.message;}}
    if (validator.conn !== null) {
      validator.conn.close(); validator.conn = null;}
    
    response.flash = err;
    res.render("ndi/wizard.html",response);});}

function get_groupname(groupname){
  var values = [];
  ntura.data.groups.forEach(function(group) {
    if (group.groupname === groupname){
      values.push(group.groupvalue);}});
  return values;}

function create_fieldlist(mname) {
  var fields_lst = [];
  for (var fieldname in ntura.model[mname]) {
    var field = ntura.model[mname][fieldname];
    var fvalue = {fieldname:fieldname, label:fieldname, fieldtype:field.type, fieldcat:1, 
      values:[], mtype:"basefield", visible:true};
    
    if (fieldname.substr(0,1) === "_" || fieldname === "id" || fieldname === "deleted"){
      fvalue.visible = false;}
    if (field.type === "integer"){
      if (typeof field.requires !== "undefined"){
        if (field.requires.bool === true){
          fvalue.fieldtype = "bool";}}}
    
    switch (mname) {
      case "address":
      case "contact":
        if (fieldname === "nervatype"){
          fvalue.fieldcat = 0;
          fvalue.fieldtype = "valuelist";
          fvalue.values = ["customer","employee","event","place","product","project","tool","trans"];}
        else if (fieldname === "ref_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"refnumber", label:"refnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"rownumber", label:"rownumber", fieldtype:"integer", 
            mtype:"basefield", fieldcat:0, values:[], visible:true});}
        break;
    
      case "barcode":
        if (fieldname === "code"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "product_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"partnumber", label:"partnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "barcodetype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("barcodetype");}
        break;
        
      case "currency":
        if (fieldname === "curr"){
          fvalue.fieldcat = 0;}
        break;
      
      case "customer":
        if (fieldname === "custnumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "custtype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("custtype");
          delete fvalue.values[fvalue.values.indexOf("own")];}
        break;
        
      case "deffield":
        if (fieldname === "fieldname"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "nervatype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = [
            "address","barcode","contact","currency","customer","employee","event","item","link","log",
            "movement","payment","price","place","product","project","rate","tax","tool","trans","setting"];}
        else if (fieldname === "subtype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = [
            "own","company","private","other","bank","cash","warehouse","item","service","printer",
            "invoice","receipt","order","offer","worksheet","rent","delivery","inventory","waybill",
            "production","formula","bank","cash"]}
        else if (fieldname === "fieldtype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("fieldtype");
          delete fvalue.values[fvalue.values.indexOf("trans")];
          delete fvalue.values[fvalue.values.indexOf("checkbox")];}
        break;
      
      case "employee":
        if (fieldname === "empnumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "usergroup"){
          fvalue.fieldtype = "string";}
        else if (fieldname === "department"){
          fvalue.fieldtype = "string";}
        else if (fieldname === "password"){
          fvalue.visible = false;}
        else if (fieldname === "registration_key"){
          fvalue.visible = false;}
        break;
      
      case "event":
        if (fieldname === "calnumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "nervatype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = ["customer","employee","place","product","project","tool","trans"];}
        else if (fieldname === "ref_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"refnumber", label:"refnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "eventgroup"){
          fvalue.fieldtype = "string";}
        break;
      
      case "fieldvalue":
        if (fieldname === "fieldname"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "ref_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"refnumber", label:"refnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"rownumber", label:"rownumber", fieldtype:"integer", 
            mtype:"basefield", fieldcat:0, values:[], visible:true});}
        else if (fieldname === "fieldtype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("fieldtype");
          delete fvalue.values[fvalue.values.indexOf("trans")];
          delete fvalue.values[fvalue.values.indexOf("checkbox")];}
        break;
        
      case "groups":
        if (fieldname === "groupname"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "groupvalue"){
          fvalue.fieldcat = 0;}
        break;
      
      case "item":
        if (fieldname === "trans_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"transnumber", label:"transnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"rownumber", label:"rownumber", fieldtype:"integer", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"inputmode", label:"inputmode", fieldtype:"valuelist", fieldcat:1, 
            mtype:"basefield", values:["fxprice","netamount","amount"], visible:true});
          fields_lst.push({fieldname:"inputvalue", label:"inputvalue", fieldtype:"float", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "product_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"partnumber", label:"partnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "tax_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"taxcode", label:"taxcode", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (["fxprice","netamount","vatamount","amount"].indexOf(fieldname)>-1){
          fvalue.visible = false;}
        break;
      
      case "link":
        if (fieldname === "nervatype_1"){
          fvalue.fieldname = "nervatype1";
          fvalue.label = "nervatype1";
          fvalue.fieldcat = 0;
          fvalue.fieldtype = "valuelist";
          fvalue.values = ["address","barcode","contact","currency","customer","employee","event","groups","item", 
            "movement","payment","price","place","product","project","rate","tax","tool","trans"];}
        else if (fieldname === "nervatype_2"){
          fvalue.fieldname = "nervatype2";
          fvalue.label = "nervatype2";
          fvalue.fieldcat = 0;
          fvalue.fieldtype = "valuelist";
          fvalue.values = ["address","barcode","contact","currency","customer","employee","event","groups","item", 
            "movement","payment","price","place","product","project","rate","tax","tool","trans"];}
        else if (fieldname === "ref_id_1"){
          fvalue.fieldname = "refnumber1";
          fvalue.label = "refnumber1";
          fvalue.fieldcat = 0;
          fvalue.fieldtype = "string";}
        else if (fieldname === "ref_id_2"){
          fvalue.fieldname = "refnumber2";
          fvalue.label = "refnumber2";
          fvalue.fieldcat = 0;
          fvalue.fieldtype = "string";}
        break;
        
      case "log":
        if (fieldname === "employee_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"empnumber", label:"empnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "crdate"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "nervatype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = ["address","barcode","contact","currency","customer","employee","deffield",
            "event","groups","item","link","movement","payment","price","place","product","project", 
            "rate","tax","tool","trans","setting"];}
        else if (fieldname === "logstate"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("logstate");}
        else if (fieldname === "ref_id"){
          fvalue.fieldname = "refnumber";
          fvalue.label = "refnumber";
          fvalue.fieldtype = "string";}
        break;
      
      case "movement":
        if (fieldname === "trans_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"transnumber", label:"transnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"rownumber", label:"rownumber", fieldtype:"integer", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "movetype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("movetype");}
        else if (fieldname === "product_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"partnumber", label:"partnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "tool_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"serial", label:"serial", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "place_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"planumber", label:"planumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        break;
        
      case "numberdef":
        if (fieldname === "numberkey"){
          fvalue.fieldcat = 0;}
        break;
      
      case "pattern":
        if (fieldname === "description"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "transtype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("transtype");}
        break;
        
      case "payment":
        if (fieldname === "trans_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"transnumber", label:"transnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"rownumber", label:"rownumber", fieldtype:"integer", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});}
        break;
        
      case "place":
        if (fieldname === "planumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "placetype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("placetype");
          delete fvalue.values[fvalue.values.indexOf("store")];}
        break;
        
      case "price":
        if (fieldname === "product_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"partnumber", label:"partnumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});
          fields_lst.push({fieldname:"pricetype", label:"pricetype", fieldtype:"valuelist", fieldcat:0, 
            mtype:"basefield", values:["price","discount"], visible:true});}
        else if (fieldname === "validfrom"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "curr"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "qty"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "calcmode"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("calcmode");}
        break;
      
      case "product":
        if (fieldname === "partnumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "protype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("protype");}
        else if (fieldname === "tax_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"taxcode", label:"taxcode", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        break;
        
      case "project":
        if (fieldname === "pronumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "customer_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"custnumber", label:"custnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        break;
      
      case "rate":
        if (fieldname === "ratetype"){
          fvalue.fieldcat = 0;
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("ratetype");}
        else if (fieldname === "ratedate"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "curr"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "place_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"planumber", label:"planumber", fieldtype:"string", fieldcat:0, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "rategroup"){
          fvalue.fieldtype = "string";}
        break;
      
      case "tax":
        if (fieldname === "taxcode"){
          fvalue.fieldcat = 0;}
        break;
        
      case "tool":
        if (fieldname === "serial"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "product_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"partnumber", label:"partnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "toolgroup"){
          fvalue.fieldtype = "string";}
        break;
        
      case "trans":
        if (fieldname === "transnumber"){
          fvalue.fieldcat = 0;}
        else if (fieldname === "transtype"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("transtype");
          delete fvalue.values[fvalue.values.indexOf("filing")];
          delete fvalue.values[fvalue.values.indexOf("store")];}
        else if (fieldname === "direction"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("direction");
          delete fvalue.values[fvalue.values.indexOf("return")];}
        else if (fieldname === "paidtype"){
          fvalue.fieldtype = "string";}
        else if (fieldname === "department"){
          fvalue.fieldtype = "string";}
        else if (fieldname === "transtate"){
          fvalue.fieldtype = "valuelist";
          fvalue.values = get_groupname("transtate");}
        else if (fieldname === "customer_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"custnumber", label:"custnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "employee_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"empnumber", label:"empnumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "project_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"pronumber", label:"pronumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "place_id"){
          fvalue.visible = false;
          fields_lst.push({fieldname:"planumber", label:"planumber", fieldtype:"string", fieldcat:1, 
            mtype:"basefield", values:[], visible:true});}
        else if (fieldname === "cruser_id"){
          fvalue.visible = false;}
        break;}
    
    if (fvalue.visible){
      fields_lst.push(fvalue);}}
    return fields_lst;}

module.exports = router;