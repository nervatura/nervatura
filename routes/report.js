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
var PyShell = require('python-shell');

router.use(function (req, res, next) {
  next()});

router.get('/', function(req, res, next) {
  res.redirect('index');});

router.get('/index', function(req, res, next) {
  res.render('report/index.html',{view:"index"});});

router.get('/server', function(req, res, next) {
  var flash;
  const exec = require('child_process').exec;
  exec(req.app.get("host_settings").python2_path+" -V", function(err, stdout, stderr){
      if (err || stderr.indexOf("Python 2.")===-1) {
        flash = req.app.locals.lang.invalid_python_path;}
     res.render('report/server.html',{flash:flash, view:"server"});});});

router.all('/document', function(req, res, next) {
  var orient = "p"; var format = "pdf"; var method = "load_report_xml";
  if (req.query.data || req.body.data){
    format = "xml";}
  if (req.query.html || req.body.html){
    format = "html";}
  if (req.query.landscape || req.body.landscape){
    orient = "l";}
  if (req.query.py || req.body.py){
    method = "create_report_sample"}
    
  const exec = require('child_process').exec;
  exec(req.app.get("host_settings").python2_path+" -V", function(err, stdout, stderr){
    if (err || stderr.indexOf("Python 2.")===-1) {
      res.end(req.app.locals.lang.invalid_python_path);}
    else{
      var ps = new PyShell("pylib.py", {
        args: [method,orient,format,req.app.get("conf").python_script+"/report/sample.xml"],
        pythonPath: req.app.get("host_settings").python2_path,
        scriptPath: req.app.get("conf").python_script,
        mode: 'text', pythonOptions: ['-u']});
      var output = '';
      ps.stdout.on('data', function (data) {
        output += ''+data;});
      ps.end(function (err) {
        if (err) {
          return next(err);}
        else {
          switch (format) {
            case "pdf":
              res.setHeader('Content-Type', 'application/pdf');
              res.end(new Buffer(output, 'base64'));
              break;
            case "xml":
              res.set('Content-Type', 'text/xml');
              res.end(output);
              break;
            default:
              res.set('Content-Type', 'text/html');
              res.end(output);
              break;}}});}});});

router.get('/template', function(req, res, next) {
  if (req.query.py){
    const exec = require('child_process').exec;
    exec(req.app.get("host_settings").python2_path+" -V", function(err, stdout, stderr){
      if (err || stderr.indexOf("Python 2.")===-1) {
        res.end(req.app.locals.lang.invalid_python_path);}
      else{
        var ps = new PyShell("pylib.py", {
          args: ["get_source","create_report_sample"],
          pythonPath: req.app.get("host_settings").python2_path,
          scriptPath: req.app.get("conf").python_script,
          mode: 'text', pythonOptions: ['-u']});
        var output = '';
        ps.stdout.on('data', function (data) {
          output += ''+data;});
        ps.end(function (err) {
          if (err) {
            return next(err);}
          else {
            res.render('report/python.html',{python_code:output});}});}});}
  else {
    res.download(req.app.get("conf").python_script+"/report/sample.xml", 'sample.xml', function(err){
      if(err){return next(err);}});}});

router.get('/client', function(req, res, next) {
  res.render('report/client.html',{view:"client"});});

module.exports = router;