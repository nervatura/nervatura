/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var express = require('express');
var router = express.Router();
var async = require("async");

router.use(function (req, res, next) {
  next()
});

router.all('/getVernum', function(req, res, next) {
  res.send(req.app.settings.version);});

router.post('/call/jsonrpc', function (req, res, next) {
  index(req.body, req, res);});

router.post('/call/jsonrpc2', function (req, res, next) {
  index(req.body, req, res);});
  
router.post('/jsonrpc', function (req, res, next) {
  index(req.body, req, res);});

router.post('/jsonrpc2', function (req, res, next) {
  index(req.body, req, res);});
  
router.post('/', function (req, res, next) {
  index(req.body, req, res);});

function index(data, req, res) {
  if (req.app.get("host_settings").npi_host_restriction.length>0 && req.app.get("host_settings").npi_host_restriction.indexOf(req.ip)===-1){
    res.send(getError("host_restrict", "NPI "+req.app.locals.lang.insecure_err));}
  else if (req.app.get("host_settings").npi_host_restriction.length===0 && req.app.get("host_settings").all_host_restriction.length>0 
    && req.app.get("host_settings").all_host_restriction.indexOf(req.ip)===-1){
    res.send(getError("host_restrict", "NPI "+req.app.locals.lang.insecure_err));}
  else {
    var nstore = require('../lib/node/nervastore.js')(req, res);
    var npi = require('../lib/node/npi.js')();
    switch (data.method) {
      case "getLogin":
      case "getLogin_json":
        npi.getLogin(nstore, data.params, function(result){res.send(getResult(1, result));});
        break;
      case "loadView":
      case "loadView_json":
        npi.setData(nstore, "view", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(2, result));}});
        break;
      case "loadTable":
      case "loadTable_json":
        npi.setData(nstore, "table", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(3, result));}});
        break;
      case "loadDataSet":
      case "loadDataSet_json":
        npi.loadDataSet(nstore, data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(4, result));}});
        break;
      case "executeSql":
      case "executeSql_json":
        npi.setData(nstore, "execute", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(5, result));}});
        break;
      case "saveRecord":
      case "saveRecord_json":
        npi.setData(nstore, "update", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(6, result));}});
        break;
      case "saveRecordSet":
      case "saveRecordSet_json":
        npi.updateRecordSet(nstore, "update", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(7, result));}});
        break;
      case "saveDataSet":
      case "saveDataSet_json":
        npi.saveDataSet(nstore, data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(8, result));}});
        break;
      
      case "deleteRecord":
      case "deleteRecord_json":
        npi.setData(nstore, "delete", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(9, result));}});
        break;
      case "deleteRecordSet":
      case "deleteRecordSet_json":
        npi.updateRecordSet(nstore, "delete", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(10, result));}});
        break;        
      case "callFunction":
      case "callFunction_json":
        npi.setData(nstore, "function", data.params, function(err, result){
          if(err){res.send(getError(null, err));}
          else {res.send(getResult(11, result));}});
        break;
          
      default:
        res.send(getError("not_found", data.method));
        break;}}};

function getResult(id, result) {
  return {"id": id, "jsonrpc": "2.0", "result": result};}

function getError(code, info) {
  var jdata = {"id": null, "jsonrpc": "2.0", "error": {}};
  switch (code) {
    case "host_restrict":
      jdata.error.code = 0;
      jdata.error.message = "Method not found";
      jdata.error.data = info;
      break;
    case "not_found":
      jdata.error.code = -32601;
      jdata.error.message = "Method not found";
      jdata.error.data = info;
      break;
    default:
      jdata.error.code = null;
      jdata.error.message = "Error result";
      jdata.error.data = info;
      break;}
  return jdata;}
 
module.exports = router;
