/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/* global Buffer */

var express = require('express');
var router = express.Router();

var lang;

router.use(function (req, res, next) {
  lang = req.app.locals.lang;
  next()});
  
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
  
router.get('/updateData', function (req, res, next) {
  req.query.method = "updateData";
  index(req.query, req, res);});

router.get('/deleteData', function (req, res, next) {
  req.query.method = "deleteData";
  index(req.query, req, res);});

router.get('/getData', function (req, res, next) {
  req.query.method = "getData";
  index(req.query, req, res);});
  
function index(data, req, res) {
  var ndi = require('../lib/node/ndi.js')();
  var ndi_data = ndi.decodeData(data);
  if (ndi_data.error){
    res.send(ndi_data.error);}
  else {
    var nstore = require('../lib/node/nervastore.js')(req, res);
    nstore.connect.getLogin(
      {database:ndi_data.params.database, username:ndi_data.params.username, password:ndi_data.params.password},
      function(err, validator){
        if (err){
          res.send(ndi.getError("login", "message", err));}
        else {
          if (nstore.encrypt_data() && nstore.encrypt_password()!==null){
            var out = require('../lib/node/tools.js').DataOutput();
            ndi_data.items = out.decipherValue(nstore.encrypt_password(), ndi_data.items, "hex");}
          ndi_data.params.validator = validator;
          ndi_data.method(nstore, ndi_data.params, ndi_data.items, function(err, result){
            if (ndi_data.params.validator.conn !== null){
              ndi_data.params.validator.conn.close();}
            if (err){
              res.send(ndi.getError(ndi_data.id, "message", err));}
            else {
              switch (ndi_data.output) {
                case "html":
                  res.set('Content-Type', 'text/html');
                  res.render("ndi/data.html",result);
                  break;
                
                case "xml":
                  res.set('Content-Type', 'text/xml');
                  res.render("ndi/data.xml",result);
                  break;
                
                case "csv":
                  var json2csv = require('json2csv');
                  res.set('Content-Type', 'text/csv');
                  res.set('Content-Disposition', 'attachment;filename="NDIExport.csv');
                  var fields = ["name","value","type","label","index","data"];
                  if (result.show_id) {fields.unshift("id");}
                  json2csv({data: result.data, fields:fields}, function(err, csv) {
                    if (err){
                      res.send(ndi.getError(ndi_data.id, "message", err));}
                    else {
                      res.send(csv);}});
                  break;
                
                default:
                  //json
                  res.set('Content-Type', 'text/json');
                  result = {"id":result.id, "jsonrpc": "2.0", "result":result.data};
                  res.send(result);
                  break;}}});}});}}
 
module.exports = router;