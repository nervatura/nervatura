/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var express = require('express');
var router = express.Router();

router.use(function (req, res, next) {
  next()
});

router.get('/', function(req, res, next) {
  res.redirect('ntura/index');
});

router.get('/index', function(req, res, next) {
  res.render('ntura/index.html',{view:"index"});
});

router.get('/ndoc', function(req, res, next) {
  res.render('ntura/index.html',{view:"ndoc"});
});

router.get('/about', function(req, res, next) {
  res.render('ntura/index.html',{view:"about"});
});

router.get('/licenses', function(req, res, next) {
  res.render('ntura/index.html',{view:"licenses"});
});

router.get('/docs/nas', function(req, res, next) {
  res.render('docs/nas.html',{});
});

router.get('/docs/nom', function(req, res, next) {
  res.render('docs/nom.html',{});
});

router.get('/docs/ndi', function(req, res, next) {
  res.render('docs/ndi.html',{});
});

router.get('/docs/npi', function(req, res, next) {
  res.render('docs/npi.html',{});
});

router.get('/docs/ndr', function(req, res, next) {
  res.render('docs/ndr.html',{});
});

router.get('/docs/report', function(req, res, next) {
  res.render('docs/report.html',{});
});
     
module.exports = router;
